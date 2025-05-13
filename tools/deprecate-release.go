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
	var grafanaAPIKey string

	flag.StringVar(&providerPath, "provider", "", "Path to the provider directory (required)")
	flag.BoolVar(&dryRun, "dry-run", false, "Run in dry-run mode without making changes")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.StringVar(&grafanaAPIKey, "grafana-api-key", os.Getenv("GRAFANA_API_KEY"), "Grafana API key for checking release usage")
	flag.Parse()

	if providerPath == "" {
		log.Fatalf("Error: -provider flag is required")
	}

	if _, err := os.Stat(providerPath); os.IsNotExist(err) {
		log.Fatalf("Provider directory %s does not exist", providerPath)
	}

	if grafanaAPIKey == "" {
		log.Fatalf("Error: Grafana API key is required for usage checks, but none was provided. " +
			"Please set the GRAFANA_API_KEY environment variable or provide it via configuration. ")
	}

	releases, err := findReleases(providerPath, verbose)
	if err != nil {
		log.Fatalf("Error finding releases: %v", err)
	}

	if len(releases) == 0 {
		log.Printf("No releases found in %s", providerPath)
		return
	}

	checkReleasesInUse(releases, grafanaAPIKey, filepath.Base(providerPath), verbose)

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

	deprecatedReleases, keptReleases := deprecateReleases(releases, verbose)

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

// checkReleasesInUse queries Grafana for all active releases
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

	// Active versions only
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

	// Compare releases that are in use from directories
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
func deprecateReleases(releases []*ReleaseInfo, verbose bool) ([]*ReleaseInfo, []*ReleaseInfo) {
	var activeReleases []*ReleaseInfo
	var deprecatedReleases = make([]*ReleaseInfo, 0)
	var keptReleases = make([]*ReleaseInfo, 0)

	// Filter to only active releases
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

	// Sort all active releases by version (newest first for easier processing)
	sort.Slice(activeReleases, func(i, j int) bool {
		return activeReleases[i].Version.GreaterThan(activeReleases[j].Version)
	})

	// Identify all releases explicitly in use
	inUseReleasePaths := make(map[string]bool)
	for _, r := range activeReleases {
		if r.InUse {
			if verbose {
				log.Printf("Marking %s to keep (in use by customer)", filepath.Base(r.Path))
			}
			keptReleases = append(keptReleases, r)
			inUseReleasePaths[r.Path] = true
		}
	}

	// Group by major, find latest in each major
	majorGrouped := make(map[uint64][]*ReleaseInfo)
	latestInMajor := make(map[uint64]*ReleaseInfo)
	var allMajorNumbers []uint64

	for _, r := range activeReleases {
		major := r.Version.Major()
		majorGrouped[major] = append(majorGrouped[major], r)
		if _, exists := latestInMajor[major]; !exists {
			latestInMajor[major] = r
			allMajorNumbers = append(allMajorNumbers, major)
		}
	}

	// Sort majors: oldest to newest
	sort.Slice(allMajorNumbers, func(i, j int) bool { return allMajorNumbers[i] < allMajorNumbers[j] })

	if verbose {
		log.Printf("Found active releases across major versions: %v", allMajorNumbers)
		for major, latest := range latestInMajor {
			log.Printf("  Latest for v%d: %s", major, filepath.Base(latest.Path))
		}
	}

	// Determine which major versions need their latest release kept
	majorsToKeepLatest := make(map[uint64]bool)

	// Keep latest 3 majors
	numMajorsToKeepByDefault := 3
	if len(allMajorNumbers) > 0 {
		for i := 0; i < numMajorsToKeepByDefault && i < len(allMajorNumbers); i++ {
			majorToKeep := allMajorNumbers[len(allMajorNumbers)-1-i]
			majorsToKeepLatest[majorToKeep] = true
			if verbose {
				log.Printf("Marking latest of v%d to keep (one of %d most recent by default)", majorToKeep, numMajorsToKeepByDefault)
			}
		}
	}

	// Keep latest of majors for upgrade paths from any in-use versions
	oldestMajorInUse := uint64(0)
	foundOldestMajorInUse := false

	for _, r := range keptReleases {
		major := r.Version.Major()
		if !foundOldestMajorInUse || major < oldestMajorInUse {
			oldestMajorInUse = major
			foundOldestMajorInUse = true
		}
	}

	if foundOldestMajorInUse {
		if verbose {
			log.Printf("Oldest major version with an in-use release: v%d. Ensuring upgrade path from there.", oldestMajorInUse)
		}
		// Ensure all majors from oldest in use up to the newest available major have their latest kept
		for _, major := range allMajorNumbers {
			if major >= oldestMajorInUse {
				if !majorsToKeepLatest[major] {
					if verbose {
						log.Printf("Marking latest of v%d to keep (part of upgrade path from v%d)", major, oldestMajorInUse)
					}
				}
				majorsToKeepLatest[major] = true
			}
		}
	}

	for major, keep := range majorsToKeepLatest {
		if keep {
			latest := latestInMajor[major]
			isAlreadyKept := false
			for _, kr := range keptReleases {
				if kr.Path == latest.Path {
					isAlreadyKept = true
					break
				}
			}
			if !isAlreadyKept {
				if verbose {
					log.Printf("Marking %s to keep (latest of a required major v%d)", filepath.Base(latest.Path), major)
				}
				keptReleases = append(keptReleases, latest)
			}
		}
	}

	// Any active release not in keptReleases is deprecated
	keptPaths := make(map[string]bool)
	for _, r := range keptReleases {
		keptPaths[r.Path] = true
	}

	for _, r := range activeReleases {
		if !keptPaths[r.Path] {
			if verbose {
				log.Printf("Marking %s for deprecation (not in kept list)", filepath.Base(r.Path))
			}
			deprecatedReleases = append(deprecatedReleases, r)
		}
	}

	// Prevent duplicates
	finalKeptReleases := make([]*ReleaseInfo, 0)
	finalKeptPaths := make(map[string]bool)
	for _, r := range keptReleases {
		if !finalKeptPaths[r.Path] {
			finalKeptReleases = append(finalKeptReleases, r)
			finalKeptPaths[r.Path] = true
		}
	}

	return deprecatedReleases, finalKeptReleases
}
