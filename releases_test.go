package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/versionbundle"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sigs.k8s.io/yaml"
)

type kustomizationFile struct {
	CommonAnnotations map[string]string `yaml:"commonAnnotations"`
	Resources         []string          `yaml:"resources"`
	Transformers      []string          `yaml:"transformers"`
}

type requestException struct {
	Version string `yaml:"version"`
	Reason  string `yaml:"reason"`
}

type versionRequest struct {
	Name       string             `yaml:"name"`
	Version    string             `yaml:"version"`
	Exceptions []requestException `yaml:"except,omitempty" json:"except,omitempty"`
}

type releaseRequest struct {
	Name     string           `yaml:"name"`
	Requests []versionRequest `yaml:"requests"`
}

type requestsFile struct {
	Releases []releaseRequest `yaml:"releases"`
}

func findReleases(provider string, archived bool) ([]v1alpha1.Release, error) {
	path := provider
	if archived {
		path = filepath.Join(provider, "archived")
	}
	releaseDirectories, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	var releases []v1alpha1.Release
	for _, releaseDirectory := range releaseDirectories {
		if !releaseDirectory.IsDir() || releaseDirectory.Name() == "archived" {
			continue
		}
		releaseFilename := filepath.Join(path, releaseDirectory.Name(), "release.yaml")
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

func componentListSatisfiesRequest(request versionRequest, componentList []v1alpha1.ReleaseSpecComponent) (bool, string, error) {
	var unsatisfied string
	for _, component := range componentList {
		if component.Name == request.Name {
			actualMatchesRequest, err := versionMatches(component.Version, request.Version)
			if err != nil {
				return false, "", microerror.Mask(err)
			}

			if actualMatchesRequest {
				return true, "", nil
			}

			unsatisfied = fmt.Sprintf("requested: %s:%s actual: %s", request.Name, request.Version, component.Version)
			break // No need to keep searching for this component
		}
	}
	return false, unsatisfied, nil
}

// versionMatches compares the given version with the given semver
// constraint pattern and returns whether it matches.
func versionMatches(version string, pattern string) (bool, error) {
	c, err := semver.NewConstraint(pattern)
	if err != nil {
		return false, fmt.Errorf("release names for requests must be valid semver constraints: %s", err)
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		return false, fmt.Errorf("release names must be valid semver: %s", err)
	}

	// if c.Check(v) {
	// 	//fmt.Printf("%s matches %s", version, pattern)
	// 	return true, nil

	// }
	// return false, nil
	return c.Check(v), nil
}

func findMatchingRequests(release string, requests []releaseRequest) ([]versionRequest, error) {
	// fmt.Printf("finding matches for %s\n", release)
	var requestList []versionRequest
	for _, request := range requests {

		match, err := versionMatches(request.Name, release)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if match {
			for _, component := range request.Requests {
				exception := false
				if component.Exceptions != nil {
					// Check the exclusions for this component to see if our release is there
					for _, e := range component.Exceptions {
						exception, err = versionMatches(request.Name, e.Version)
						if err != nil {
							return nil, microerror.Mask(err)
						}
					}
				}

				if !exception {
					requestList = append(requestList, component)
				}
			}
		}
	}
	return requestList, nil
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

			providerRequests := []releaseRequest{}
			{
				var providerRequestsFile requestsFile
				providerRequestsData, err := ioutil.ReadFile(filepath.Join(tc.provider, "requests.yaml"))
				if err != nil {
					t.Fatal(err)
				}
				err = yaml.UnmarshalStrict(providerRequestsData, &providerRequestsFile)
				if err != nil {
					t.Fatal(err)
				}
				for _, release := range providerRequestsFile.Releases {
					providerRequests = append(providerRequests, release)
				}
			}

			// for k, r := range providerRequests {
			// 	fmt.Printf("provider requests: k:%d r:%v#!", k, r)
			// }

			releases, err := findReleases(tc.provider, false)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}

			crd := v1alpha1.NewReleaseCRD()

			for _, crdVersion := range crd.Spec.Versions {
				var v apiextensions.CustomResourceValidation
				// Convert the CRD validation into the version-independent form.
				err := v1.Convert_v1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crdVersion.Schema, &v, nil)
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
					releaseNotesData, err := ioutil.ReadFile(filepath.Join(tc.provider, release.Name, "README.md"))
					if err != nil {
						t.Errorf("missing file for %s release %s: %s", tc.provider, release.Name, err)
					}
					releaseNotesLines := strings.Split(string(releaseNotesData), "\n")
					if len(releaseNotesLines) == 0 || !strings.Contains(releaseNotesLines[0], strings.TrimPrefix(release.Name, "v")) {
						t.Errorf("expected release notes for %s release %s to contain the release version on the first line", tc.provider, release.Name)
					}
				}

				// Check that the README links to the release.
				if !strings.Contains(readmeContent, fmt.Sprintf("https://github.com/giantswarm/releases/tree/master/%s/%s", tc.provider, release.Name)) {
					t.Errorf("expected link in README.md to %s release %s", tc.provider, release.Name)
				}

				// Check that this release contains all requested component versions
				{
					requests, err := findMatchingRequests(release.Name, providerRequests)
					if err != nil {
						t.Fatal(microerror.Mask(err))
					}

					unsatisfiedRequests := map[string][]string{}
					for _, request := range requests {
						satisfied, unsatisfied, err := componentListSatisfiesRequest(request, release.Spec.Components)
						if err != nil {
							t.Fatal(microerror.Mask(err))
						}
						if !satisfied {
							// fmt.Printf("%#v\n", unsatisfied)
							unsatisfiedRequests[release.Name] = append(unsatisfiedRequests[release.Name], unsatisfied)
						}
					}
					if len(unsatisfiedRequests) > 0 {
						// fmt.Printf("%#v\n", unsatisfiedRequests)
						msg := fmt.Sprintf("Release %s does not meet the requested version requirements: %s", release.Name, strings.Join(unsatisfiedRequests[release.Name], ", "))
						t.Errorf(msg)
					}

				}

			}

			archived, err := findReleases(tc.provider, true)
			if err != nil {
				t.Fatal(microerror.Mask(err))
			}

			for _, release := range archived {
				// Check that the README links to the release.
				if !strings.Contains(readmeContent, fmt.Sprintf("https://github.com/giantswarm/releases/tree/master/%s/archived/%s", tc.provider, release.Name)) {
					t.Errorf("expected link in README.md to archived %s release %s", tc.provider, release.Name)
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
