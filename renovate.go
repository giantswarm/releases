package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/giantswarm/release-operator/v3/api/v1alpha1"
	"gopkg.in/yaml.v2"
)

type Release struct {
	Version          string `yaml:"version"`
	IsDeprecated     bool   `yaml:"isDeprecated"`
	ReleaseTimestamp string `json:"releaseTimestamp"`
	ChangelogUrl     string `json:"changelogUrl"`
	IsStable         bool   `json:"isStable"`
}

type Releases struct {
	Releases     []Release `json:"releases"`
	SourceUrl    string    `json:"sourceUrl"`
	ChangelogUrl string    `json:"changelogUrl"`
	Homepage     string    `json:"homepage"`
}

type GiantSwarmRelease struct {
	Metadata struct {
		Name string `json:"name"`
	}
	Spec v1alpha1.ReleaseSpec `json:"spec"`
}

func main() {
	directories := []string{"capa", "azure"}
	for _, dir := range directories {
		releases := Releases{
			SourceUrl:    "https://github.com/giantswarm/releases",
			ChangelogUrl: "https://github.com/giantswarm/releases/blob/main/README.md",
			Homepage:     "https://giantswarm.io",
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory %s: %v\n", dir, err)
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				releaseFile := filepath.Join(dir, f.Name(), "release.yaml")
				if _, err := os.Stat(releaseFile); os.IsNotExist(err) {
					continue
				}

				data, err := os.ReadFile(releaseFile)
				if err != nil {
					continue
				}

				var gsRelease = GiantSwarmRelease{}
				if err := yaml.Unmarshal(data, &gsRelease); err != nil {
					continue
				}

				// Use regex to extract MAJOR.MINOR.PATCH from the version
				versionRegex := regexp.MustCompile(`(?:provider-)?(\d+\.\d+\.\d+)`)
				matches := versionRegex.FindStringSubmatch(gsRelease.Metadata.Name)
				if len(matches) < 2 { // matches[0] is the full match, matches[1] is the first capture group
					continue
				}

				// Set additional fields not in YAML
				release := Release{
					Version:          matches[1],
					IsDeprecated:     gsRelease.Spec.State == "deprecated",
					ChangelogUrl:     fmt.Sprintf("https://github.com/giantswarm/releases/blob/master/%s/v%s/README.md", dir, matches[1]),
					IsStable:         !strings.Contains(matches[1], "alpha") || !strings.Contains(matches[1], "beta") || !strings.Contains(matches[1], "rc"),
					ReleaseTimestamp: gsRelease.Spec.Date.String(),
				}

				releases.Releases = append(releases.Releases, release)
			}
		}
		jsonData, err := json.MarshalIndent(releases, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling JSON: %v\n", err)
			return
		}

		if err := os.WriteFile(fmt.Sprintf("%s/releases.json", dir), jsonData, 0644); err != nil {
			fmt.Printf("Error writing JSON to file: %v\n", err)
		}
	}

}
