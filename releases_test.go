package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/versionbundle"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sigs.k8s.io/yaml"
)

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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
				fmt.Println(release.ObjectMeta)
				result := validator.Validate(release)
				if len(result.Errors) > 0 {
					t.Errorf("invalid release: %#v", release)
					for i, err := range result.Errors {
						t.Errorf("validation error %d: %#v", i, err)
					}
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
