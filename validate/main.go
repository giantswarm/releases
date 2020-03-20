package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/versionbundle"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sigs.k8s.io/yaml"
)

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

func validateIndividualReleases(releases []v1alpha1.Release) error {
	crd := v1alpha1.NewReleaseCRD()
	var v apiextensions.CustomResourceValidation
	err := v1beta1.Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crd.Spec.Validation, &v, nil)
	if err != nil {
		return microerror.Mask(err)
	}

	validator, _, err := validation.NewSchemaValidator(&v)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, release := range releases {
		result := validator.Validate(release)
		if len(result.Errors) > 0 {
			return microerror.Mask(errors.New(fmt.Sprintf("invalid release: %#v", result.Errors)))
		}
	}

	return nil
}

func validateCombinedReleases(releases []v1alpha1.Release) error {
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
				Name:     component.Name,
				Provider: component.Name,
				Version:  component.Version,
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
	err := versionbundle.ValidateIndexReleases(indexReleases)
	if err != nil {
		return microerror.Mask(err)
	}
	return nil
}

func main() {
	files := []string{
		"aws.yaml",
		"azure.yaml",
		"kvm.yaml",
	}
	for _, file := range files {
		documents, err := scanDocuments(file)
		if err != nil {
			log.Fatal(err)
		}
		releases, err := parseReleases(documents)
		if err != nil {
			log.Fatal(err)
		}
		err = validateIndividualReleases(releases)
		if err != nil {
			log.Fatal(err)
		}
		err = validateCombinedReleases(releases)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf(color.GreenString("Verified given releases are valid.\n"))
}
