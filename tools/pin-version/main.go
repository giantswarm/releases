package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Request struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Release struct {
	Name     string    `yaml:"name"`
	Requests []Request `yaml:"requests"`
}

type RequestsFile struct {
	Releases []Release `yaml:"releases"`
}

func main() {
	var (
		requestsFile     string
		releaseVersion   string
		untilVersion     string
		componentName    string
		componentVersion string
	)

	flag.StringVar(&requestsFile, "requests-file", "", "Path to requests.yaml file")
	flag.StringVar(&releaseVersion, "release-version", "", "Current release version (e.g., 31.1.2)")
	flag.StringVar(&untilVersion, "until", "", "Pin until this version (optional)")
	flag.StringVar(&componentName, "component-name", "", "Component or app name to pin")
	flag.StringVar(&componentVersion, "component-version", "", "Version to pin")
	flag.Parse()

	if requestsFile == "" || releaseVersion == "" || componentName == "" || componentVersion == "" {
		log.Fatal("Required flags: -requests-file, -release-version, -component-name, -component-version")
	}

	// Read the requests.yaml file
	data, err := os.ReadFile(requestsFile)
	if err != nil {
		log.Fatalf("Failed to read requests file: %v", err)
	}

	var requestsData RequestsFile
	if err := yaml.Unmarshal(data, &requestsData); err != nil {
		log.Fatalf("Failed to parse requests file: %v", err)
	}

	// Determine the name constraint
	nameConstraint := determineNameConstraint(releaseVersion, untilVersion)

	// Find or create the release entry
	releaseIndex := -1
	for i, release := range requestsData.Releases {
		if release.Name == nameConstraint {
			releaseIndex = i
			break
		}
	}

	newRequest := Request{
		Name:    componentName,
		Version: componentVersion,
	}

	if releaseIndex >= 0 {
		// Update existing entry
		componentIndex := -1
		for i, req := range requestsData.Releases[releaseIndex].Requests {
			if req.Name == componentName {
				componentIndex = i
				break
			}
		}

		if componentIndex >= 0 {
			// Update existing component
			requestsData.Releases[releaseIndex].Requests[componentIndex] = newRequest
		} else {
			// Add new component at the beginning
			requestsData.Releases[releaseIndex].Requests = append([]Request{newRequest}, requestsData.Releases[releaseIndex].Requests...)
		}
	} else {
		// Create new release entry at the beginning
		newRelease := Release{
			Name:     nameConstraint,
			Requests: []Request{newRequest},
		}
		requestsData.Releases = append([]Release{newRelease}, requestsData.Releases...)
	}

	// Write back to file
	output, err := yaml.Marshal(&requestsData)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}

	if err := os.WriteFile(requestsFile, output, 0644); err != nil {
		log.Fatalf("Failed to write requests file: %v", err)
	}

	fmt.Printf("Updated %s with pin: %s@%s for constraint %s\n", requestsFile, componentName, componentVersion, nameConstraint)
}

// determineNameConstraint creates the name constraint based on release version and until version
func determineNameConstraint(releaseVersion, untilVersion string) string {
	// Remove 'v' prefix if present
	releaseVersion = strings.TrimPrefix(releaseVersion, "v")

	// If no until version, pin only this exact release
	if untilVersion == "" {
		return releaseVersion
	}

	// Parse release version components
	releaseParts := strings.Split(releaseVersion, ".")
	if len(releaseParts) != 3 {
		log.Fatalf("Invalid release version format: %s (expected X.Y.Z)", releaseVersion)
	}

	// Parse until version
	untilVersion = strings.TrimPrefix(untilVersion, "v")
	untilParts := strings.Split(untilVersion, ".")
	if len(untilParts) != 3 {
		log.Fatalf("Invalid until version format: %s (expected X.Y.Z)", untilVersion)
	}

	releaseMajor, _ := strconv.Atoi(releaseParts[0])
	releaseMinor, _ := strconv.Atoi(releaseParts[1])

	untilMajor, _ := strconv.Atoi(untilParts[0])
	untilMinor, _ := strconv.Atoi(untilParts[1])

	// Determine the constraint range
	if releaseMajor < untilMajor {
		// Pin across major version boundaries
		return fmt.Sprintf(">= %s.%s.0 < %s.0.0", releaseParts[0], releaseParts[1], untilParts[0])
	} else if releaseMajor == untilMajor && releaseMinor < untilMinor {
		// Pin within same major, across minor versions
		return fmt.Sprintf(">= %s.%s.0 < %s.%s.0", releaseParts[0], releaseParts[1], untilParts[0], untilParts[1])
	} else {
		// Until version is same or earlier - just pin this release
		return releaseVersion
	}
}
