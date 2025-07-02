package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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

	releasesByName := make(map[string]*ReleaseInfo)
	for _, releaseInfo := range releases {
		if releaseInfo.Release != nil {
			releasesByName[releaseInfo.Release.ObjectMeta.Name] = releaseInfo
		}
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
			if verbose {
				log.Printf("Dry run: Simulating diff update for %s", releaseInfo.Path)
			}
			if err := updateReleaseDiffState(releaseInfo, releasesByName, verbose, dryRun); err != nil {
				log.Printf("Error during dry-run diff processing for %s: %v", releaseInfo.Path, err)
			}
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
				if err := updateReleaseDiffState(releaseInfo, releasesByName, verbose, dryRun); err != nil {
					log.Printf("Error updating diff file for %s: %v", releaseInfo.Path, err)
				}
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
	// Use ALL releases (not just active) to ensure complete upgrade paths
	majorGrouped := make(map[uint64][]*ReleaseInfo)
	latestInMajor := make(map[uint64]*ReleaseInfo)
	latestActiveinMajor := make(map[uint64]*ReleaseInfo) // Track latest active separately
	var allMajorNumbers []uint64

	// First pass: find latest release in each major (including deprecated)
	for _, r := range releases {
		major := r.Version.Major()
		if _, exists := latestInMajor[major]; !exists || r.Version.GreaterThan(latestInMajor[major].Version) {
			latestInMajor[major] = r
		}
		if isActive(r.Release) {
			if latestActive, exists := latestActiveinMajor[major]; !exists || r.Version.GreaterThan(latestActive.Version) {
				latestActiveinMajor[major] = r
			}
		}
	}

	// Build list of all major numbers
	for major := range latestInMajor {
		allMajorNumbers = append(allMajorNumbers, major)
	}

	// Group active releases by major
	for _, r := range activeReleases {
		major := r.Version.Major()
		majorGrouped[major] = append(majorGrouped[major], r)
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
	// Check ALL releases (both active and deprecated) that are in use
	oldestMajorInUse := uint64(0)
	foundOldestMajorInUse := false
	majorsWithInUseReleases := make(map[uint64]bool)

	for _, r := range releases {
		if r.InUse {
			major := r.Version.Major()
			majorsWithInUseReleases[major] = true
			if !foundOldestMajorInUse || major < oldestMajorInUse {
				oldestMajorInUse = major
				foundOldestMajorInUse = true
			}
		}
	}

	if foundOldestMajorInUse {
		if verbose {
			log.Printf("Oldest major version with an in-use release: v%d. Ensuring upgrade path from there.", oldestMajorInUse)
			majorsList := make([]uint64, 0, len(majorsWithInUseReleases))
			for major := range majorsWithInUseReleases {
				majorsList = append(majorsList, major)
			}
			sort.Slice(majorsList, func(i, j int) bool { return majorsList[i] < majorsList[j] })
			log.Printf("Major versions with in-use releases: %v", majorsList)
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

		// Additionally, ensure that any major version with in-use releases has its latest kept
		// This is crucial for upgrade paths from deprecated but still in-use releases
		for major := range majorsWithInUseReleases {
			if _, exists := latestInMajor[major]; exists {
				if !majorsToKeepLatest[major] {
					if verbose {
						log.Printf("Marking latest of v%d to keep (has in-use releases)", major)
					}
				}
				majorsToKeepLatest[major] = true
			}
		}
	} else {
		// Fallback: When Grafana API fails, ensure upgrade path continuity
		// If we're keeping releases from recent majors, ensure we can upgrade from older majors
		if len(allMajorNumbers) > 0 && len(majorsToKeepLatest) > 0 {
			// Find oldest major we're keeping by default (latest 3)
			oldestKeptMajor := uint64(0)
			for major := range majorsToKeepLatest {
				if oldestKeptMajor == 0 || major < oldestKeptMajor {
					oldestKeptMajor = major
				}
			}

			// Find oldest available major
			oldestAvailableMajor := allMajorNumbers[0]
			for _, major := range allMajorNumbers {
				if major < oldestAvailableMajor {
					oldestAvailableMajor = major
				}
			}

			if verbose {
				log.Printf("API failed - using fallback upgrade path logic. Oldest available: v%d, oldest kept: v%d", oldestAvailableMajor, oldestKeptMajor)
			}

			// Ensure upgrade path from oldest available to oldest kept
			for _, major := range allMajorNumbers {
				if major >= oldestAvailableMajor && major <= oldestKeptMajor {
					if !majorsToKeepLatest[major] {
						if verbose {
							log.Printf("Marking latest of v%d to keep (upgrade path continuity fallback)", major)
						}
					}
					majorsToKeepLatest[major] = true
				}
			}
		}
	}

	for major, keep := range majorsToKeepLatest {
		if keep {
			// Prefer latest active release in this major, but use overall latest if no active release exists
			var releaseToKeep *ReleaseInfo
			if latestActive, hasActive := latestActiveinMajor[major]; hasActive {
				releaseToKeep = latestActive
			} else {
				// No active release in this major, need to reactivate the latest deprecated one
				releaseToKeep = latestInMajor[major]
				if verbose {
					log.Printf("Need to reactivate %s (latest in v%d major, required for upgrade path)", filepath.Base(releaseToKeep.Path), major)
				}
			}

			isAlreadyKept := false
			for _, kr := range keptReleases {
				if kr.Path == releaseToKeep.Path {
					isAlreadyKept = true
					break
				}
			}
			if !isAlreadyKept {
				if verbose {
					log.Printf("Marking %s to keep (latest of a required major v%d)", filepath.Base(releaseToKeep.Path), major)
				}
				keptReleases = append(keptReleases, releaseToKeep)
			}
		}
	}

	// Build list of releases to keep active and releases to deprecate
	keptPaths := make(map[string]bool)
	for _, r := range keptReleases {
		keptPaths[r.Path] = true
	}

	// Check active releases - deprecate those not in kept list
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

// updateReleaseDiffState modifies the corresponding .diff file to reflect
// the new deprecated state for the current release (right side)
// and the actual state of the previous release (left side).
// It also handles transforming "no-pipe" state lines into "pipe" format
// *if a state change is necessary*.
func updateReleaseDiffState(
	releaseInfoCurrent *ReleaseInfo,
	releasesByName map[string]*ReleaseInfo,
	verbose bool,
	dryRun bool,
) error {
	releaseDirPath := releaseInfoCurrent.Path
	diffFilePath := filepath.Join(releaseDirPath, "release.diff")

	originalDiffContent, err := os.ReadFile(diffFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			if verbose {
				log.Printf("Diff file %s does not exist, skipping update.", diffFilePath)
			}
			return nil
		}
		return fmt.Errorf("failed to read diff file %s: %w", diffFilePath, err)
	}

	if len(originalDiffContent) == 0 {
		if verbose {
			log.Printf("Diff file %s is empty, skipping update.", diffFilePath)
		}
		return nil
	}

	var processedLines []string
	var modifiedInDiff bool // Flag to track if any line in the file was changed

	metaNameRegex := regexp.MustCompile(`^\s*name:\s*(\S+)\s*\|\s*name:\s*(\S+)\s*$`)
	stateLineWithPipeRegex := regexp.MustCompile(`^(\s*state:\s*\S+\s*)\|(\s*state:\s*\S+\s*)$`)
	stateLineNoPipeRegex := regexp.MustCompile(`^(\s*state:\s*)(\S+)(.*?)(\S+)\s*$`)
	stateValueCaptureRegex := regexp.MustCompile(`(state:\s*)(\S+)`)

	var previousReleaseNameFromDiff string
	var previousReleaseActualState v1alpha1.ReleaseState
	previousReleaseStateKnown := false

	// --- Pre-scan diff to find the name of the previous release ---
	tempScanner := bufio.NewScanner(bytes.NewReader(originalDiffContent))
	for tempScanner.Scan() {
		line := tempScanner.Text()
		if matches := metaNameRegex.FindStringSubmatch(line); len(matches) == 3 {
			previousReleaseNameFromDiff = matches[1]
			if prevRelInfo, ok := releasesByName[previousReleaseNameFromDiff]; ok && prevRelInfo.Release != nil {
				previousReleaseActualState = prevRelInfo.Release.Spec.State
				previousReleaseStateKnown = true
				if verbose {
					log.Printf("Pre-scan: Previous release from diff is '%s', its actual state is '%s'.", previousReleaseNameFromDiff, previousReleaseActualState)
				}
			} else {
				if verbose {
					log.Printf("Pre-scan: Previous release '%s' (from diff) not found in release map or has no parsed release data.", previousReleaseNameFromDiff)
				}
			}
			break
		}
	}
	if !previousReleaseStateKnown && verbose {
		log.Printf("Pre-scan: Could not determine previous release name/state from diff's metadata.name line.")
	}
	// --- End pre-scan ---

	scanner := bufio.NewScanner(bytes.NewReader(originalDiffContent))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		processedLine := line

		if verbose {
			log.Printf("Scanning diff file %s, Line %d: [%s]", diffFilePath, lineNumber, line)
		}

		if pipeMatches := stateLineWithPipeRegex.FindStringSubmatch(line); len(pipeMatches) == 3 {
			lineModifiedThisIteration := false
			if verbose {
				log.Printf("Line %d for %s MATCHED by stateLineWithPipeRegex.", lineNumber, diffFilePath)
			}
			leftPartOriginal := pipeMatches[1]
			rightPartOriginal := pipeMatches[2]

			finalLeftPart := leftPartOriginal
			finalRightPart := rightPartOriginal

			if previousReleaseStateKnown {
				targetLeftStateString := string(previousReleaseActualState)
				if subMatchesLeft := stateValueCaptureRegex.FindStringSubmatch(leftPartOriginal); len(subMatchesLeft) == 3 {
					currentValLeft := subMatchesLeft[2]
					if currentValLeft != targetLeftStateString {
						finalLeftPart = stateValueCaptureRegex.ReplaceAllString(leftPartOriginal, "${1}"+targetLeftStateString)
						lineModifiedThisIteration = true
						if verbose {
							log.Printf("(Pipe) Left part target state '%s', current in diff '%s'. MODIFIED.", targetLeftStateString, currentValLeft)
						}
					} else {
						if verbose {
							log.Printf("(Pipe) Left part target state '%s', current in diff '%s'. No change needed.", targetLeftStateString, currentValLeft)
						}
					}
				} else {
					if verbose {
						log.Printf("(Pipe) Could not parse 'state: value' from left part: [%s]", leftPartOriginal)
					}
				}
			}

			targetRightStateString := string(v1alpha1.StateDeprecated)
			if subMatchesRight := stateValueCaptureRegex.FindStringSubmatch(rightPartOriginal); len(subMatchesRight) == 3 {
				currentValRight := subMatchesRight[2]
				if currentValRight != targetRightStateString {
					finalRightPart = stateValueCaptureRegex.ReplaceAllString(rightPartOriginal, "${1}"+targetRightStateString)
					lineModifiedThisIteration = true
					if verbose {
						log.Printf("(Pipe) Right part target state '%s', current in diff '%s'. MODIFIED.", targetRightStateString, currentValRight)
					}
				} else {
					if verbose {
						log.Printf("(Pipe) Right part target state '%s', current in diff '%s'. No change needed.", targetRightStateString, currentValRight)
					}
				}
			} else {
				if verbose {
					log.Printf("(Pipe) Could not parse 'state: value' from right part: [%s]", rightPartOriginal)
				}
			}

			if lineModifiedThisIteration {
				processedLine = finalLeftPart + "|" + finalRightPart
				modifiedInDiff = true
				if verbose {
					log.Printf("(Pipe) Diff line modification PROPOSED for %s (line %d):\n    Old: %s\n    New: %s", diffFilePath, lineNumber, line, processedLine)
				}
			}

		} else if noPipeMatches := stateLineNoPipeRegex.FindStringSubmatch(line); len(noPipeMatches) == 5 {
			// noPipeMatches: [0]Full, [1]Prefix ("  state: "), [2]Val1, [3]MiddlePart, [4]Val2
			if verbose {
				log.Printf("Line %d for %s PRELIMINARY MATCH by stateLineNoPipeRegex (5 groups).", lineNumber, diffFilePath)
				log.Printf("Group 1 (Prefix): [%s]", noPipeMatches[1])
				log.Printf("Group 2 (Val1 from diff):   [%s]", noPipeMatches[2])
				log.Printf("Group 3 (Middle): [%s]", noPipeMatches[3])
				log.Printf("Group 4 (Val2 from diff):   [%s]", noPipeMatches[4])
			}

			indentAndPrefix := noPipeMatches[1]
			val1FromDiff := noPipeMatches[2]
			val2FromDiff := noPipeMatches[4] // Assuming Val2 is the one corresponding to the current release

			targetLeftState := val1FromDiff // Default to what's in the diff
			if previousReleaseStateKnown {
				targetLeftState = string(previousReleaseActualState)
			}
			targetRightState := string(v1alpha1.StateDeprecated)

			// Only reformat and change if the target states are different from what's parsed,
			// OR if the values parsed from the diff (val1FromDiff, val2FromDiff) are not identical
			// (which implies it might be a malformed line that needs fixing to pipe format anyway if it was intended to be two distinct values).
			// The primary driver for reformatting a "no-pipe" line should be a change in state values.
			if targetLeftState != val1FromDiff || targetRightState != val2FromDiff {
				if verbose {
					log.Printf("(No-Pipe) State change detected. Val1InDiff: %s, TargetLeft: %s. Val2InDiff: %s, TargetRight: %s. REFORMATTING.",
						val1FromDiff, targetLeftState, val2FromDiff, targetRightState)
				}

				leftSideString := indentAndPrefix + targetLeftState
				rightSideString := "state: " + targetRightState
				const leftPartTargetWidth = 40
				paddingLength := leftPartTargetWidth - len(leftSideString)
				if paddingLength < 1 {
					paddingLength = 1
				}
				padding := strings.Repeat(" ", paddingLength)
				rightLeadingSpaces := "         "

				processedLine = leftSideString + padding + "|" + rightLeadingSpaces + rightSideString
				modifiedInDiff = true
				if verbose {
					log.Printf("(No-Pipe) Line %d REFORMATTED to pipe format:\n    Old: %s\n    New: %s", lineNumber, line, processedLine)
				}
			} else {
				if verbose {
					log.Printf("(No-Pipe) Line %d: Parsed values ('%s', '%s') match target states ('%s', '%s'). No reformatting needed.",
						lineNumber, val1FromDiff, val2FromDiff, targetLeftState, targetRightState)
				}
			}

		} else {
			if verbose && strings.Contains(line, "state:") {
				log.Printf("Line %d for %s containing 'state:' DID NOT MATCH stateLineWithPipeRegex.", lineNumber, diffFilePath)
				attemptedNoPipeMatches := stateLineNoPipeRegex.FindStringSubmatch(line)
				log.Printf("For stateLineNoPipeRegex (pattern: `%s`):", stateLineNoPipeRegex.String())
				log.Printf("FindStringSubmatch on line [%s] returned %d elements (expected 5 for a structural match).", line, len(attemptedNoPipeMatches))
				if len(attemptedNoPipeMatches) > 0 {
					for i, m := range attemptedNoPipeMatches {
						log.Printf("  Submatch[%d]: [%s]", i, m)
					}
				} else {
					log.Printf("  No submatches found by stateLineNoPipeRegex.")
				}
			}
		}
		processedLines = append(processedLines, processedLine)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning diff file %s: %w", diffFilePath, err)
	}

	if modifiedInDiff {
		log.Printf("Would update state line(s) in diff file: %s", diffFilePath)
		if dryRun {
			if verbose {
				log.Printf("Dry run: Skipping actual write to diff file %s", diffFilePath)
			}
			return nil
		}

		newDiffContentStr := strings.Join(processedLines, "\n")
		originalEndsWithNewline := len(originalDiffContent) > 0 && originalDiffContent[len(originalDiffContent)-1] == '\n'
		if len(newDiffContentStr) > 0 {
			currentEndsWithNewline := newDiffContentStr[len(newDiffContentStr)-1] == '\n'
			if originalEndsWithNewline && !currentEndsWithNewline {
				newDiffContentStr += "\n"
			} else if !originalEndsWithNewline && currentEndsWithNewline {
				newDiffContentStr = newDiffContentStr[:len(newDiffContentStr)-1]
			}
		} else if originalEndsWithNewline {
			newDiffContentStr = "\n"
		}

		if err := os.WriteFile(diffFilePath, []byte(newDiffContentStr), 0644); err != nil {
			return fmt.Errorf("failed to write modified diff file %s: %w", diffFilePath, err)
		}
		log.Printf("Updated state line(s) in diff file: %s", diffFilePath)
	} else {
		if verbose {
			log.Printf("No 'state' line update was ultimately needed or performed for diff file %s based on detailed line processing.", diffFilePath)
		} else if !dryRun {
			log.Printf("No 'state' line update needed for diff file %s", diffFilePath)
		}
	}
	return nil
}
