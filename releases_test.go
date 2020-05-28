package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/versionbundle"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sigs.k8s.io/yaml"
)

type kustomizationFile struct {
	CommonAnnotations map[string]string `yaml:"commonAnnotations"`
	Resources         []string          `yaml:"resources"`
}

func findReleases(provider string) ([]v1alpha1.Release, error) {
	releaseDirectories, err := ioutil.ReadDir(provider)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	var releases []v1alpha1.Release
	for _, releaseDirectory := range releaseDirectories {
		if !releaseDirectory.IsDir() {
			continue
		}
		releaseFilename := filepath.Join(provider, releaseDirectory.Name(), "release.yaml")
		data, err := ioutil.ReadFile(releaseFilename)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		var release v1alpha1.Release
		err = yaml.Unmarshal(data, &release)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if releaseDirectory.Name() != release.Name {
			return nil, fmt.Errorf("%s release %s is in directory %s which doesn't match its name", provider, release.Name, releaseDirectory)
		}
		releases = append(releases, release)
	}
	return releases, nil
}

// To reuse versionbundle.ValidateIndexReleases, the slice of Releases must first be
// converted into a slice of versionbundle.IndexRelease.
func releasesToIndex(releases []v1alpha1.Release) []versionbundle.IndexRelease {
	var indexReleases []versionbundle.IndexRelease
	for _, release := range releases {
		var apps []versionbundle.App
		for _, app := range release.Spec.Apps {
			indexApp := versionbundle.App{
				App:              app.Name,
				ComponentVersion: app.ComponentVersion,
				Version:          app.Version,
			}
			apps = append(apps, indexApp)
		}
		var authorities []versionbundle.Authority
		for _, component := range release.Spec.Components {
			indexAuthority := versionbundle.Authority{
				Name:    component.Name,
				Version: component.Version,
			}
			authorities = append(authorities, indexAuthority)
		}
		indexRelease := versionbundle.IndexRelease{
			Active:      release.Spec.State == "active",
			Apps:        apps,
			Authorities: authorities,
			Date:        release.Spec.Date.Time,
			Version:     release.Name,
		}
		indexReleases = append(indexReleases, indexRelease)
	}
	return indexReleases
}

func Test_Releases(t *testing.T) {
	testCases := []struct {
		provider string
		name     string
	}{
		{
			provider: "aws",
			name:     "case 1: aws releases are valid",
		},
		{
			provider: "azure",
			name:     "case 2: azure releases are valid",
		},
		{
			provider: "kvm",
			name:     "case 3: kvm releases are valid",
		},
	}

	// Load the README so we can check links for each release.
	var readmeContent string
	{
		readmeContentBytes, err := ioutil.ReadFile("README.md")
		if err != nil {
			t.Fatal(err)
		}
		readmeContent = string(readmeContentBytes)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerResources := map[string]bool{}
			{
				var providerKustomization kustomizationFile
				providerKustomizationData, err := ioutil.ReadFile(filepath.Join(tc.provider, "kustomization.yaml"))
				if err != nil {
					t.Fatal(err)
				}
				err = yaml.UnmarshalStrict(providerKustomizationData, &providerKustomization)
				if err != nil {
					t.Fatal(err)
				}
				for _, resource := range providerKustomization.Resources {
					providerResources[resource] = false
				}
			}

			releases, err := findReleases(tc.provider)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}

			crd := v1alpha1.NewReleaseCRD()
			var v apiextensions.CustomResourceValidation
			// Convert the CRD validation into the version-independent form.
			err = v1beta1.Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crd.Spec.Validation, &v, nil)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}

			validator, _, err := validation.NewSchemaValidator(&v)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}

			// Ensure that releases satisfy OpenAPI validation.
			for _, release := range releases {
				// Check that the release is registered in the main provider kustomization.yaml resources.
				if _, ok := providerResources[release.Name]; !ok {
					t.Errorf("release %s not registered in %s/kustomization.yaml", release.Name, tc.provider)
				}
				providerResources[release.Name] = true
				result := validator.Validate(release)
				if len(result.Errors) > 0 {
					t.Errorf("invalid release: %#v", release)
					for i, err := range result.Errors {
						t.Errorf("validation error %d: %#v", i, err)
					}
				}
			}

			for _, release := range releases {
				// Check that the release-specific kustomization.yaml file points to the release manifest.
				{
					releaseKustomizationData, err := ioutil.ReadFile(filepath.Join(tc.provider, release.Name, "kustomization.yaml"))
					if err != nil {
						t.Errorf("missing file for %s release %s: %s", tc.provider, release.Name, err)
					}
					var releaseKustomization kustomizationFile
					err = yaml.UnmarshalStrict(releaseKustomizationData, &releaseKustomization)
					if len(releaseKustomization.Resources) != 1 || releaseKustomization.Resources[0] != "release.yaml" {
						t.Errorf("kustomization.yaml for %s release %s should contain only one resource, \"release.yaml\"", tc.provider, release.Name)
					}
				}

				// Check that the version in the first line of the release notes is correct.
				{
					releaseNotesData, err := ioutil.ReadFile(filepath.Join(tc.provider, release.Name, "release-notes.md"))
					if err != nil {
						t.Errorf("missing file for %s release %s: %s", tc.provider, release.Name, err)
					}
					releaseNotesLines := strings.Split(string(releaseNotesData), "\n")
					if len(releaseNotesLines) == 0 || !strings.Contains(releaseNotesLines[0], strings.TrimPrefix(release.Name, "v")) {
						t.Errorf("expected release notes for %s release %s to contain the release version on the first line", tc.provider, release.Name)
					}
				}

				// Check that the README links to the release.
				if !strings.Contains(readmeContent, fmt.Sprintf("https://github.com/giantswarm/releases/blob/master/%s/%s/release-notes.md", tc.provider, release.Name)) {
					t.Errorf("expected link in README.md to %s release %s", tc.provider, release.Name)
				}
			}

			// Check for extra resources in provider kustomization.yaml that don't have a corresponding release.
			for release, processed := range providerResources {
				if !processed {
					t.Errorf("release %s registered in %s/kustomization.yaml resources but not found", release, tc.provider)
				}
			}

			// Ensure that releases are unique.
			indexReleases := releasesToIndex(releases)
			err = versionbundle.ValidateIndexReleases(indexReleases)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}
		})
	}
}
