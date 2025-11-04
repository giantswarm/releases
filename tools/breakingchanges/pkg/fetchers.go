package breakingchanges

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// fetchURL fetches content from a URL with optional GitHub authentication
func (d *Detector) fetchURL(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	if d.githubToken != "" && strings.Contains(url, "github.com") {
		req.Header.Set("Authorization", "token "+d.githubToken)
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// fetchComponentMappings fetches and parses the devctl component mappings
func (d *Detector) fetchComponentMappings(ctx context.Context) error {
	if d.componentMappingDone {
		return nil // Already fetched
	}

	fmt.Println("Fetching component mappings from devctl...")

	devctlURL := "https://raw.githubusercontent.com/giantswarm/devctl/main/pkg/release/changelog/changelog.go"
	content, err := d.fetchURL(ctx, devctlURL)
	if err != nil {
		fmt.Printf("Warning: Failed to fetch devctl mappings: %v\n", err)
		d.componentMappingDone = true
		return nil // Non-fatal, continue without mappings
	}

	d.parseComponentMappings(content)
	fmt.Printf("✓ Loaded %d component mappings from devctl\n", len(d.componentMappings))
	d.componentMappingDone = true
	return nil
}

// parseComponentMappings parses Go source to extract component mappings
func (d *Detector) parseComponentMappings(content string) {
	// Pattern: "component-name": { Tag: "...", Changelog: "...", }
	componentPattern := regexp.MustCompile(`(?s)"([^"]+)":\s*\{[^}]*?Tag:\s*"([^"]*)"[^}]*?Changelog:\s*"([^"]*)"`)
	matches := componentPattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			componentName := match[1]
			tagURL := match[2]
			changelogURL := match[3]

			// Extract repository name from the URL
			repoPattern := regexp.MustCompile(`github\.com/giantswarm/([^/]+)`)
			if repoMatch := repoPattern.FindStringSubmatch(tagURL); len(repoMatch) >= 2 {
				d.componentMappings[componentName] = ComponentMapping{
					RepositoryName: repoMatch[1],
					ChangelogURL:   changelogURL,
					TagURL:         tagURL,
				}
			}
		}
	}
}

// fetchExternalChangelogs fetches Flatcar and Kubernetes changelogs
func (d *Detector) fetchExternalChangelogs(ctx context.Context, changes []VersionChange) (map[string]string, error) {
	result := make(map[string]string)

	for _, change := range changes {
		if change.ChangelogURL == "" {
			continue
		}

		// Only fetch for external dependencies
		if change.Component != "Flatcar" && change.Component != "Kubernetes" {
			continue
		}

		fmt.Printf("Fetching %s changelog from %s...\n", change.Component, change.ChangelogURL)

		content, err := d.fetchURL(ctx, change.ChangelogURL)
		if err != nil {
			fmt.Printf("Warning: Failed to fetch %s changelog: %v\n", change.Component, err)
			continue
		}

		// Process based on component type
		if change.Component == "Flatcar" && strings.Contains(change.ChangelogURL, "releases-stable.json") {
			content = d.processFlatcarChangelog(content, change.FromVersion, change.ToVersion)
		} else if change.Component == "Kubernetes" {
			content = d.processKubernetesChangelog(content, change.FromVersion, change.ToVersion)
		}

		// Limit size to avoid token overload
		if len(content) > 50000 {
			content = content[:50000] + "\n\n[... truncated for size ...]"
		}

		result[change.Component] = content
		fmt.Printf("✓ Fetched %s changelog (%d bytes)\n", change.Component, len(content))
	}

	return result, nil
}

// processFlatcarChangelog extracts Flatcar release notes from JSON
func (d *Detector) processFlatcarChangelog(content, fromVersion, toVersion string) string {
	releaseNotes, err := d.extractFlatcarReleaseNotes(content, fromVersion, toVersion)
	if err != nil {
		fmt.Printf("Warning: Failed to parse Flatcar JSON: %v\n", err)
		return content
	}
	fmt.Printf("✓ Extracted Flatcar release notes v%s -> v%s (%d bytes)\n", fromVersion, toVersion, len(releaseNotes))
	return releaseNotes
}

// processKubernetesChangelog extracts breaking changes from K8s changelog
func (d *Detector) processKubernetesChangelog(content, fromVersion, toVersion string) string {
	breakingChanges := d.extractK8sBreakingChanges(content, fromVersion, toVersion)
	if breakingChanges != "" {
		fmt.Printf("✓ Extracted Kubernetes breaking changes (%d bytes)\n", len(breakingChanges))
		return breakingChanges
	}
	return content
}

// fetchComponentDiffs fetches diffs for Giant Swarm components
func (d *Detector) fetchComponentDiffs(ctx context.Context, changes []VersionChange) (map[string]string, error) {
	result := make(map[string]string)

	for _, change := range changes {
		// Skip external dependencies (Flatcar, Kubernetes) - handled separately
		if change.Component == "Flatcar" || change.Component == "Kubernetes" {
			continue
		}

		d.fetchSingleComponentDiff(ctx, change, result)
	}

	return result, nil
}

// fetchSingleComponentDiff fetches a single component's diff and its dependencies
func (d *Detector) fetchSingleComponentDiff(ctx context.Context, change VersionChange, result map[string]string) {
	// Resolve actual repository name from devctl mappings
	repoName := change.Component
	changelogTemplate := ""
	if mapping, ok := d.componentMappings[change.Component]; ok {
		repoName = mapping.RepositoryName
		changelogTemplate = mapping.ChangelogURL
		fmt.Printf("  → Resolved %s to repo %s\n", change.Component, repoName)
	}

	// Always build diff URL with resolved repository name
	diffURL := fmt.Sprintf("https://github.com/giantswarm/%s/compare/v%s...v%s.diff",
		repoName, change.FromVersion, change.ToVersion)

	fmt.Printf("Fetching diff for %s (v%s -> v%s)...\n", change.Component, change.FromVersion, change.ToVersion)

	diff, err := d.fetchURL(ctx, diffURL)
	if err != nil {
		d.tryFetchChangelog(ctx, change.Component, change.ToVersion, repoName, changelogTemplate, result)
		return
	}

	result[change.Component] = diff
	fmt.Printf("✓ Fetched diff for %s (%d bytes)\n", change.Component, len(diff))

	// Check for Chart.yaml dependency changes
	d.fetchDependencyDiffs(ctx, diff, change.Component, repoName, result)
}

// tryFetchChangelog attempts to fetch the changelog if diff fails
func (d *Detector) tryFetchChangelog(ctx context.Context, componentName, version, repoName, changelogTemplate string, result map[string]string) {
	fmt.Printf("Warning: Failed to fetch diff for %s, trying CHANGELOG\n", componentName)

	var changelogURL string
	if changelogTemplate != "" {
		changelogURL = strings.ReplaceAll(changelogTemplate, "{{.Version}}", version)
	} else {
		changelogURL = fmt.Sprintf("https://raw.githubusercontent.com/giantswarm/%s/master/CHANGELOG.md", repoName)
	}

	changelog, err := d.fetchURL(ctx, changelogURL)
	if err == nil {
		result[componentName+"_CHANGELOG"] = changelog
		fmt.Printf("✓ Fetched CHANGELOG for %s\n", componentName)
	}
}

// isLikelyUpstreamChart heuristically determines if a chart is an upstream chart based on version numbers
// Upstream charts typically have high major versions (v10+, v20+) while Giant Swarm charts are usually v0-v5
func isLikelyUpstreamChart(change VersionChange) bool {
	// Parse major version from ToVersion
	versionPattern := regexp.MustCompile(`^v?(\d+)\.`)
	if match := versionPattern.FindStringSubmatch(change.ToVersion); len(match) >= 2 {
		// If major version >= 10, it's likely an upstream chart
		var majorVersion int
		fmt.Sscanf(match[1], "%d", &majorVersion)
		return majorVersion >= 10
	}
	return false
}

// fetchDependencyDiffs fetches diffs for Chart.yaml dependencies
func (d *Detector) fetchDependencyDiffs(ctx context.Context, diff, parentComponent, parentRepoName string, result map[string]string) {
	dependencyChanges := d.extractChartDependencies(diff, parentComponent)

	for _, depChange := range dependencyChanges {
		fmt.Printf("  → Detected dependency: %s (v%s -> v%s)\n", depChange.Component, depChange.FromVersion, depChange.ToVersion)

		// Resolve dependency name through devctl mappings
		depRepoName := depChange.Component
		if mapping, ok := d.componentMappings[depChange.Component]; ok {
			depRepoName = mapping.RepositoryName
			fmt.Printf("  → Resolved %s to repo %s\n", depChange.Component, depRepoName)
		} else {
			// Not in devctl mappings - check if it looks like an upstream chart
			if isLikelyUpstreamChart(depChange) {
				fmt.Printf("  → Skipping %s (likely upstream chart based on version numbers)\n", depChange.Component)
				continue
			}
			// Try standard Giant Swarm repo pattern
			// (some core charts like 'cluster' might not be in devctl but still exist)
			fmt.Printf("  → %s not in devctl mappings, trying standard Giant Swarm repo\n", depChange.Component)
		}

		// Skip circular/self-references
		if depRepoName == parentRepoName {
			fmt.Printf("  → Skipping %s (circular reference to parent component)\n", depChange.Component)
			continue
		}

		// Fetch the dependency diff
		depDiffURL := fmt.Sprintf("https://github.com/giantswarm/%s/compare/v%s...v%s.diff",
			depRepoName, depChange.FromVersion, depChange.ToVersion)

		depDiff, err := d.fetchURL(ctx, depDiffURL)
		if err != nil {
			fmt.Printf("  Warning: Failed to fetch dependency diff for %s: %v\n", depChange.Component, err)
			continue
		}

		result[depChange.Component+"_dependency"] = depDiff
		fmt.Printf("  ✓ Fetched dependency diff for %s (%d bytes)\n", depChange.Component, len(depDiff))
	}
}
