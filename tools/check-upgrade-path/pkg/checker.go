package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v3"
)

type UpgradePathChecker struct {
	Provider         string
	Version          string
	WorkingDir       string
	NextMajorVersion string
}

type Release struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Apps       []App       `yaml:"apps"`
		Components []Component `yaml:"components"`
		State      string      `yaml:"state"`
		Date       string      `yaml:"date"`
	} `yaml:"spec"`
}

type App struct {
	Name      string   `yaml:"name"`
	Version   string   `yaml:"version"`
	Catalog   string   `yaml:"catalog,omitempty"`
	DependsOn []string `yaml:"dependsOn,omitempty"`
}

type Component struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Catalog string `yaml:"catalog,omitempty"`
}

type VersionWarning struct {
	ItemType         string // "app" or "component"
	Name             string
	CurrentVersion   string
	NextMajorVersion string
	IsDowngrade      bool
}

func (c *UpgradePathChecker) Check() ([]VersionWarning, error) {
	// Parse the current release version to extract major
	currentMajor, err := extractMajor(c.Version)
	if err != nil {
		return nil, fmt.Errorf("invalid version format: %w", err)
	}

	// Find the next major version
	nextMajor := currentMajor + 1
	c.NextMajorVersion = fmt.Sprintf("v%d.0", nextMajor)

	// Check if this is the latest major (no next major to compare against)
	nextMajorPath := filepath.Join(c.WorkingDir, c.Provider)
	entries, err := os.ReadDir(nextMajorPath)
	if err != nil {
		return nil, fmt.Errorf("error reading provider directory: %w", err)
	}

	// Find the smallest unarchived minor version in the next major
	var nextMajorReleases []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		// Check if this is a release directory for the next major
		if strings.HasPrefix(entry.Name(), fmt.Sprintf("v%d.", nextMajor)) {
			releasePath := filepath.Join(nextMajorPath, entry.Name(), "release.yaml")
			if _, err := os.Stat(releasePath); err == nil {
				// Check if it's not archived
				release, err := parseRelease(releasePath)
				if err != nil {
					continue
				}
				if release.Spec.State != "archived" {
					nextMajorReleases = append(nextMajorReleases, entry.Name())
				}
			}
		}
	}

	if len(nextMajorReleases) == 0 {
		// No next major found, this is the latest major, no need to check
		return nil, nil
	}

	// Sort to get the smallest version
	sort.Strings(nextMajorReleases)
	smallestNextMajor := nextMajorReleases[0]
	c.NextMajorVersion = smallestNextMajor

	// Load both releases
	currentReleasePath := filepath.Join(c.WorkingDir, c.Provider, c.Version, "release.yaml")
	nextMajorReleasePath := filepath.Join(c.WorkingDir, c.Provider, smallestNextMajor, "release.yaml")

	currentRelease, err := parseRelease(currentReleasePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing current release: %w", err)
	}

	nextMajorRelease, err := parseRelease(nextMajorReleasePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing next major release: %w", err)
	}

	// Compare versions
	var warnings []VersionWarning

	// Compare apps
	nextMajorApps := make(map[string]string)
	for _, app := range nextMajorRelease.Spec.Apps {
		nextMajorApps[app.Name] = app.Version
	}

	for _, app := range currentRelease.Spec.Apps {
		if nextMajorVersion, exists := nextMajorApps[app.Name]; exists {
			if isNewerVersion(app.Version, nextMajorVersion) {
				warnings = append(warnings, VersionWarning{
					ItemType:         "App",
					Name:             app.Name,
					CurrentVersion:   app.Version,
					NextMajorVersion: nextMajorVersion,
					IsDowngrade:      true,
				})
			}
		}
	}

	// Compare components
	nextMajorComponents := make(map[string]string)
	for _, comp := range nextMajorRelease.Spec.Components {
		nextMajorComponents[comp.Name] = comp.Version
	}

	for _, comp := range currentRelease.Spec.Components {
		if nextMajorVersion, exists := nextMajorComponents[comp.Name]; exists {
			if isNewerVersion(comp.Version, nextMajorVersion) {
				warnings = append(warnings, VersionWarning{
					ItemType:         "Component",
					Name:             comp.Name,
					CurrentVersion:   comp.Version,
					NextMajorVersion: nextMajorVersion,
					IsDowngrade:      true,
				})
			}
		}
	}

	return warnings, nil
}

func parseRelease(path string) (*Release, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var release Release
	if err := yaml.Unmarshal(data, &release); err != nil {
		return nil, err
	}

	return &release, nil
}

func extractMajor(version string) (int, error) {
	// Remove 'v' prefix if present
	version = strings.TrimPrefix(version, "v")

	// Split by '.' and get the first part
	parts := strings.Split(version, ".")
	if len(parts) < 1 {
		return 0, fmt.Errorf("invalid version format: %s", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid major version: %s", parts[0])
	}

	return major, nil
}

func isNewerVersion(v1, v2 string) bool {
	// Parse versions using semver
	ver1, err1 := semver.NewVersion(v1)
	ver2, err2 := semver.NewVersion(v2)

	// If either version can't be parsed, fall back to string comparison
	if err1 != nil || err2 != nil {
		return v1 > v2
	}

	return ver1.GreaterThan(ver2)
}
