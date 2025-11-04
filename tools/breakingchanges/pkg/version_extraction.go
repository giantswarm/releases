package breakingchanges

import (
	"fmt"
	"regexp"
)

// extractVersionChanges parses the README to find all version changes
func extractVersionChanges(readme string) []VersionChange {
	var changes []VersionChange

	// Extract Flatcar and Kubernetes FIRST (before generic pattern) so they get proper changelog URLs
	changes = append(changes, extractFlatcarVersion(readme)...)
	changes = append(changes, extractKubernetesVersion(readme)...)
	changes = append(changes, extractGenericComponentVersions(readme)...)
	changes = append(changes, extractComponentCompareLinks(readme)...)

	return deduplicateVersionChanges(changes)
}

// extractFlatcarVersion extracts Flatcar version changes
func extractFlatcarVersion(readme string) []VersionChange {
	var changes []VersionChange

	// Pattern for Flatcar changes - ALWAYS use JSON API
	flatcarPattern := regexp.MustCompile(`Flatcar from v?([\d.]+) to \[v?([\d.]+)\]`)
	if match := flatcarPattern.FindStringSubmatch(readme); len(match) >= 3 {
		changes = append(changes, VersionChange{
			Component:    "Flatcar",
			FromVersion:  match[1],
			ToVersion:    match[2],
			ChangelogURL: "https://www.flatcar.org/releases-json/releases-stable.json",
		})
		return changes
	}

	// Try without link/brackets
	flatcarPatternNoLink := regexp.MustCompile(`Flatcar from v?([\d.]+) to .*?v?([\d.]+)`)
	if match := flatcarPatternNoLink.FindStringSubmatch(readme); len(match) >= 3 {
		changes = append(changes, VersionChange{
			Component:    "Flatcar",
			FromVersion:  match[1],
			ToVersion:    match[2],
			ChangelogURL: "https://www.flatcar.org/releases-json/releases-stable.json",
		})
	}

	return changes
}

// extractKubernetesVersion extracts Kubernetes version changes
func extractKubernetesVersion(readme string) []VersionChange {
	var changes []VersionChange

	// Pattern for Kubernetes changes with changelog link
	k8sPattern := regexp.MustCompile(`Kubernetes from v?([\d.]+) to \[v?([\d.]+)\]\((https://[^)]+)\)`)
	if match := k8sPattern.FindStringSubmatch(readme); len(match) >= 4 {
		if changelogURL := buildK8sChangelogURL(match[2]); changelogURL != "" {
			changes = append(changes, VersionChange{
				Component:    "Kubernetes",
				FromVersion:  match[1],
				ToVersion:    match[2],
				ChangelogURL: changelogURL,
			})
			return changes
		}
	}

	// Try without link - build K8s changelog URL
	k8sPatternNoLink := regexp.MustCompile(`Kubernetes from v?([\d.]+) to .*?v?([\d.]+)`)
	if match := k8sPatternNoLink.FindStringSubmatch(readme); len(match) >= 3 {
		if changelogURL := buildK8sChangelogURL(match[2]); changelogURL != "" {
			changes = append(changes, VersionChange{
				Component:    "Kubernetes",
				FromVersion:  match[1],
				ToVersion:    match[2],
				ChangelogURL: changelogURL,
			})
		}
	}

	return changes
}

// buildK8sChangelogURL constructs the raw GitHub changelog URL for a K8s version
func buildK8sChangelogURL(version string) string {
	versionParts := regexp.MustCompile(`^(\d+\.\d+)`).FindStringSubmatch(version)
	if len(versionParts) < 2 {
		return ""
	}
	majorMinor := versionParts[1]
	return fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/master/CHANGELOG/CHANGELOG-%s.md", majorMinor)
}

// extractGenericComponentVersions extracts version changes for Giant Swarm components
func extractGenericComponentVersions(readme string) []VersionChange {
	var changes []VersionChange

	// Pattern for component version changes
	// Example: "cluster-aws from v5.0.0 to v6.0.0"
	componentPattern := regexp.MustCompile(`(?m)^-\s+(\S+)\s+from\s+v?([\d.]+)\s+to\s+.*?v?([\d.]+)`)
	matches := componentPattern.FindAllStringSubmatch(readme, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			changes = append(changes, VersionChange{
				Component:   match[1],
				FromVersion: match[2],
				ToVersion:   match[3],
				CompareURL:  fmt.Sprintf("https://github.com/giantswarm/%s/compare/v%s...v%s", match[1], match[2], match[3]),
			})
		}
	}

	return changes
}

// extractComponentCompareLinks extracts version changes from compare links
func extractComponentCompareLinks(readme string) []VersionChange {
	var changes []VersionChange

	// Pattern: "### cluster-aws [v3.4.0...v3.6.0](https://github.com/giantswarm/cluster-aws/compare/v3.4.0...v3.6.0)"
	comparePattern := regexp.MustCompile(`###\s+(\S+)\s+\[v?([\d.]+)\.\.\.v?([\d.]+)\]\((https://github\.com/[^)]+)\)`)
	matches := comparePattern.FindAllStringSubmatch(readme, -1)

	for _, match := range matches {
		if len(match) >= 5 {
			changes = append(changes, VersionChange{
				Component:   match[1],
				FromVersion: match[2],
				ToVersion:   match[3],
				CompareURL:  match[4],
			})
		}
	}

	return changes
}

// deduplicateVersionChanges removes duplicate version changes, keeping first occurrence
func deduplicateVersionChanges(changes []VersionChange) []VersionChange {
	seen := make(map[string]bool)
	var deduplicated []VersionChange

	for _, change := range changes {
		key := fmt.Sprintf("%s:%s:%s", change.Component, change.FromVersion, change.ToVersion)
		if !seen[key] {
			seen[key] = true
			deduplicated = append(deduplicated, change)
		}
	}

	return deduplicated
}

