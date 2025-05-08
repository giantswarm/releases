package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/release-operator/v4/api/v1alpha1"
	"sigs.k8s.io/yaml"
)

// ReleaseInfo combines the release with its parsed semver and file path
type ReleaseInfo struct {
	Path    string
	Version *semver.Version
	Release *v1alpha1.Release
	InUse   bool
}

// GrafanaResponse represents the structure of Grafana's API response
type GrafanaResponse struct {
	Results map[string]struct {
		Frames []struct {
			Schema struct {
				Fields []struct {
					Labels map[string]string `json:"labels"`
				} `json:"fields"`
			} `json:"schema"`
		} `json:"frames"`
	} `json:"results"`
}

func main() {
	var providerPath string
	var dryRun bool
	var verbose bool
	var skipUsageCheck bool
	var grafanaAPIKey string

	flag.StringVar(&providerPath, "provider", "", "Path to the provider directory (required)")
	flag.BoolVar(&dryRun, "dry-run", false, "Run in dry-run mode without making changes")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&skipUsageCheck, "skip-usage-check", false, "Skip checking if releases are in use")
	flag.StringVar(&grafanaAPIKey, "grafana-api-key", os.Getenv("GRAFANA_API_KEY"), "Grafana API key for checking release usage")
	flag.Parse()

	if providerPath == "" {
		log.Fatalf("Error: -provider flag is required")
	}

	if _, err := os.Stat(providerPath); os.IsNotExist(err) {
		log.Fatalf("Provider directory %s does not exist", providerPath)
	}

	if !skipUsageCheck && grafanaAPIKey == "" {
		log.Printf("Warning: No Grafana API key provided. Usage checks will be skipped.")
		skipUsageCheck = true
	}

	releases, err := findReleases(providerPath, verbose)
	if err != nil {
		log.Fatalf("Error finding releases: %v", err)
	}

	if len(releases) == 0 {
		log.Printf("No releases found in %s", providerPath)
		return
	}

	if !skipUsageCheck {
		checkReleasesInUse(releases, grafanaAPIKey, filepath.Base(providerPath), verbose)
	}

	if verbose {
		log.Printf("Found %d total releases", len(releases))

		activeCount := 0
		inUseCount := 0
		for _, r := range releases {
			if isActive(r.Release) {
				activeCount++
			}
			if r.InUse {
				inUseCount++
			}
		}
		log.Printf("%d active, %d deprecated, %d in use", activeCount, len(releases)-activeCount, inUseCount)
	}

	deprecatedReleases, keptReleases := deprecateReleases(releases, skipUsageCheck, verbose)

	if len(deprecatedReleases) == 0 {
		log.Println("No releases needed to be deprecated")
		return
	}

	log.Printf("Would deprecate %d releases, keeping %d releases", len(deprecatedReleases), len(keptReleases))

	if verbose {
		log.Println("Keeping these releases active:")
		for _, release := range keptReleases {
			inUseStr := ""
			if release.InUse {
				inUseStr = " (in use)"
			}
			log.Printf("  - %s%s", filepath.Base(release.Path), inUseStr)
		}
	}

	for _, releaseInfo := range deprecatedReleases {
		filePath := filepath.Join(releaseInfo.Path, "release.yaml")

		log.Printf("Would deprecate: %s", filePath)

		if dryRun {
			continue
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			continue
		}

		newContent := strings.Replace(string(content), "state: active", "state: deprecated", 1)
		newContent = strings.Replace(newContent, "  state: active", "  state: deprecated", 1)

		if newContent != string(content) {
			if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
				log.Printf("Error writing file %s: %v", filePath, err)
			} else {
				log.Printf("Deprecated: %s", filePath)
			}
		} else {
			log.Printf("No 'state: active' found in %s", filePath)
		}
	}
}

// checkReleasesInUse queries Grafana for all active releases in a single API call
func checkReleasesInUse(releases []*ReleaseInfo, apiKey string, provider string, verbose bool) {
	// Convert provider name for API if needed
	apiProvider := provider
	if provider == "azure" {
		apiProvider = "capz"
	}

	if verbose {
		log.Printf("Querying Grafana API for active %s releases...", provider)
	}

	query := map[string]interface{}{
		"from": "now-5m",
		"to":   "now",
		"queries": []map[string]interface{}{
			{
				"refId": "A",
				"expr":  fmt.Sprintf(`sum(aggregation:giantswarm:cluster_release_version{provider="%s", release_version=~".*", installation=~".*", cluster_type=~".*", customer=~".*"}) by (release_version)`, apiProvider),
				"datasource": map[string]string{
					"uid":  "000000006",
					"type": "prometheus",
				},
			},
		},
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		log.Printf("Error creating Grafana query: %v", err)
		return
	}

	req, err := http.NewRequest("POST", "https://giantswarm.grafana.net/api/ds/query", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error querying Grafana API: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Grafana API returned non-OK status: %s", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var grafanaResp GrafanaResponse
	if err := json.Unmarshal(body, &grafanaResp); err != nil {
		log.Printf("Error parsing Grafana response: %v", err)
		return
	}

	// Extract active versions
	activeVersions := make(map[string]bool)
	resultsA, hasResults := grafanaResp.Results["A"]
	if hasResults {
		for _, frame := range resultsA.Frames {
			for _, field := range frame.Schema.Fields {
				if field.Labels != nil {
					if version, hasVersion := field.Labels["release_version"]; hasVersion {
						activeVersions[version] = true
					}
				}
			}
		}
	}

	if verbose {
		log.Printf("Found %d active versions for %s", len(activeVersions), provider)
		for v := range activeVersions {
			log.Printf("  - %s is in use", v)
		}
	}

	// Mark releases that are in use
	for _, release := range releases {
		version := strings.TrimPrefix(filepath.Base(release.Path), "v")
		if activeVersions[version] {
			release.InUse = true
			if verbose {
				log.Printf("Release %s is in use by clusters", filepath.Base(release.Path))
			}
		}
	}
}

// isActive checks if a release is in the active state
func isActive(release *v1alpha1.Release) bool {
	return release.Spec.State == v1alpha1.StateActive
}

// findReleases locates all release directories and parses their version info
func findReleases(providerPath string, verbose bool) ([]*ReleaseInfo, error) {
	var releases []*ReleaseInfo

	entries, err := os.ReadDir(providerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read provider directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirName := entry.Name()
		versionStr := strings.TrimPrefix(dirName, "v")

		// Try to parse as semver
		version, err := semver.NewVersion(versionStr)
		if err != nil {
			if verbose {
				log.Printf("Skipping directory %s: not a valid semver", dirName)
			}
			continue
		}

		yamlPath := filepath.Join(providerPath, dirName, "release.yaml")
		if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
			if verbose {
				log.Printf("Skipping directory %s: no release.yaml found", dirName)
			}
			continue
		}

		data, err := os.ReadFile(yamlPath)
		if err != nil {
			log.Printf("Error reading %s: %v", yamlPath, err)
			continue
		}

		var release v1alpha1.Release
		if err := yaml.Unmarshal(data, &release); err != nil {
			log.Printf("Error parsing YAML in %s: %v", yamlPath, err)
			continue
		}

		releases = append(releases, &ReleaseInfo{
			Path:    filepath.Join(providerPath, dirName),
			Version: version,
			Release: &release,
			InUse:   false,
		})
	}

	return releases, nil
}

// deprecateReleases applies the deprecation rules to the releases
func deprecateReleases(releases []*ReleaseInfo, skipUsageCheck bool, verbose bool) ([]*ReleaseInfo, []*ReleaseInfo) {
	var activeReleases []*ReleaseInfo
	var deprecatedReleases []*ReleaseInfo = []*ReleaseInfo{}
	var keptReleases []*ReleaseInfo = []*ReleaseInfo{}

	for _, release := range releases {
		if isActive(release.Release) {
			activeReleases = append(activeReleases, release)
		}
	}

	if verbose {
		log.Printf("Processing %d active releases", len(activeReleases))
	}

	if len(activeReleases) == 0 {
		return deprecatedReleases, keptReleases
	}

	// Sort all releases by version (newest first)
	sort.Slice(activeReleases, func(i, j int) bool {
		return activeReleases[i].Version.GreaterThan(activeReleases[j].Version)
	})

	// Group releases by major version
	majorGrouped := make(map[uint64][]*ReleaseInfo)
	var allMajors []uint64

	for _, release := range activeReleases {
		major := release.Version.Major()
		if _, exists := majorGrouped[major]; !exists {
			allMajors = append(allMajors, major)
		}
		majorGrouped[major] = append(majorGrouped[major], release)
	}

	// Sort majors (newest first)
	sort.Slice(allMajors, func(i, j int) bool {
		return allMajors[i] > allMajors[j]
	})

	if verbose {
		log.Printf("Found releases with major versions: %v", allMajors)
	}

	// RULE 1: For each major version, find the latest minor/patch and deprecate all others
	for _, major := range allMajors {
		releases := majorGrouped[major]

		latestRelease := releases[0]

		if verbose {
			log.Printf("For major version v%d, latest version is %s",
				major, latestRelease.Version.String())
		}

		keptReleases = append(keptReleases, latestRelease)

		// Deprecate all other releases in this major version (all predecessors)
		// unless they are in use
		for _, rel := range releases[1:] {
			if !skipUsageCheck && rel.InUse {
				if verbose {
					log.Printf("Keeping %s active because it's in use by clusters",
						filepath.Base(rel.Path))
				}
				keptReleases = append(keptReleases, rel)
			} else {
				if verbose {
					log.Printf("Deprecating %s because it's a predecessor within major version v%d (latest is %s)",
						filepath.Base(rel.Path), major, latestRelease.Version.String())
				}
				deprecatedReleases = append(deprecatedReleases, rel)
			}
		}
	}

	// RULE 2: If there are more than 3 major versions, deprecate all from 3+ versions back
	// unless they are in use
	if len(allMajors) > 3 {
		// We want to keep the top three major versions
		keepMajors := map[uint64]bool{}

		for i := 0; i < 3 && i < len(allMajors); i++ {
			keepMajors[allMajors[i]] = true
		}

		if verbose {
			var majorsToKeep []uint64
			for i := 0; i < 3 && i < len(allMajors); i++ {
				majorsToKeep = append(majorsToKeep, allMajors[i])
			}
			log.Printf("Keeping major versions %v, deprecating older majors", majorsToKeep)
		}

		// Move any releases from older major versions from kept to deprecated
		// unless they are in use
		var stillKept []*ReleaseInfo
		for _, release := range keptReleases {
			major := release.Version.Major()
			if !keepMajors[major] && (skipUsageCheck || !release.InUse) {
				if verbose {
					log.Printf("Deprecating %s because major v%d is too old (keeping only the top 3 major versions)",
						filepath.Base(release.Path), major)
				}
				deprecatedReleases = append(deprecatedReleases, release)
			} else {
				stillKept = append(stillKept, release)
			}
		}
		keptReleases = stillKept
	}

	return deprecatedReleases, keptReleases
}
