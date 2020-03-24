package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/versionbundle"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sigs.k8s.io/yaml"
)

// sigs.k8s.io/yaml can't handle multi-part YAML documents (separated by "---")
// so documents such as aws.yaml must be split on boundaries and parsed
// individually.
func scanDocuments(filename string) ([]string, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	scanner := bufio.NewScanner(data)
	var document string
	var documents []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "---" {
			document += line + "\n"
			continue
		}
		if document == "" {
			continue
		}
		documents = append(documents, document)
		document = ""
	}
	return documents, nil
}

// Given a slice of strings defining Release CRs as YAML, this function
// simply parses the YAML and returns a slice of Releases.
func parseReleases(documents []string) ([]v1alpha1.Release, error) {
	var releases []v1alpha1.Release
	for _, document := range documents {
		var release v1alpha1.Release
		err := yaml.Unmarshal([]byte(document), &release)
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
		filename string
		name     string
	}{
		{
			filename: "aws.yaml",
			name:     "case 1: aws.yaml is valid",
		},
		{
			filename: "azure.yaml",
			name:     "case 2: azure.yaml is valid",
		},
		{
			filename: "kvm.yaml",
			name:     "case 3: kvm.yaml is valid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			documents, err := scanDocuments(tc.filename)
			if err != nil {
				t.Fatal(err)
			}

			releases, err := parseReleases(documents)
			if err != nil {
				t.Fatal(err)
			}

			crd := v1alpha1.NewReleaseCRD()
			var v apiextensions.CustomResourceValidation
			// Convert the CRD validation into the version-independent form.
			err = v1beta1.Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crd.Spec.Validation, &v, nil)
			if err != nil {
				t.Fatal(err)
			}

			validator, _, err := validation.NewSchemaValidator(&v)
			if err != nil {
				t.Fatal(err)
			}

			// Ensure that releases satisfy OpenAPI validation.
			for _, release := range releases {
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
				t.Error(err)
			}
		})
	}
}
