package breakingchanges

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// extractK8sBreakingChanges extracts breaking change sections from K8s changelog
// It filters to only include changes relevant to the upgrade (fromVersion -> toVersion)
func (d *Detector) extractK8sBreakingChanges(changelog, fromVersion, toVersion string) string {
	var sb strings.Builder

	// Parse version numbers
	fromMinor := parseK8sMinorVersion(fromVersion)
	toMinor := parseK8sMinorVersion(toVersion)

	sb.WriteString(fmt.Sprintf("# Kubernetes Upgrade: v%s → v%s\n\n", fromVersion, toVersion))

	// If minor version is the same, this is just a patch upgrade
	if fromMinor == toMinor {
		sb.WriteString(fmt.Sprintf("**Note**: This is a patch upgrade within Kubernetes 1.%d. Breaking changes are unlikely.\n\n", toMinor))
		sb.WriteString(extractK8sPatchChanges(changelog, toVersion))
	} else {
		// Extract sections only for the target minor version
		sb.WriteString(fmt.Sprintf("**Note**: Upgrading from Kubernetes 1.%d to 1.%d. Reviewing changes in 1.%d release.\n\n", fromMinor, toMinor, toMinor))
		sb.WriteString(extractK8sMinorVersionChanges(changelog, toMinor))
	}

	result := sb.String()
	if len(result) == 0 {
		return changelog // Fallback to full changelog
	}

	return result
}

// parseK8sMinorVersion extracts minor version from K8s version string
// e.g., "1.33.5" -> 33, "v1.34.1" -> 34
func parseK8sMinorVersion(version string) int {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return 0
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return minor
}

// extractK8sMinorVersionChanges extracts changes for a specific minor version
func extractK8sMinorVersionChanges(changelog string, minorVersion int) string {
	var sb strings.Builder

	// Find the section for this minor version (e.g., "# v1.34.0")
	versionHeader := fmt.Sprintf("# v1.%d.0", minorVersion)
	versionIdx := strings.Index(changelog, versionHeader)
	if versionIdx == -1 {
		return fmt.Sprintf("Could not find section for Kubernetes 1.%d in changelog.\n", minorVersion)
	}

	// Find the next major version header to know where this section ends
	nextVersionPattern := regexp.MustCompile(fmt.Sprintf(`\n# v1\.%d\.`, minorVersion-1))
	endIdx := len(changelog)
	if match := nextVersionPattern.FindStringIndex(changelog[versionIdx+1:]); match != nil {
		endIdx = versionIdx + 1 + match[0]
	}

	// Extract just this version's section
	versionSection := changelog[versionIdx:endIdx]

	// PRIORITY 1: Extract "Urgent Upgrade Notes" section first (most critical)
	urgentNotes := extractK8sUrgentUpgradeNotes(versionSection)
	if urgentNotes != "" {
		sb.WriteString("## ⚠️ URGENT UPGRADE NOTES ⚠️\n\n")
		sb.WriteString(urgentNotes)
		sb.WriteString("\n\n")
	}

	// PRIORITY 2: Extract structured sections (Deprecation, Breaking Changes, etc.)
	// Note: extractK8sSections prioritizes Deprecation section over generic API Changes
	sb.WriteString(extractK8sSections(versionSection))

	return sb.String()
}

// extractK8sUrgentUpgradeNotes specifically extracts the "Urgent Upgrade Notes" section
func extractK8sUrgentUpgradeNotes(changelog string) string {
	// Look for "## Urgent Upgrade Notes" section
	// Structure is typically:
	// ## Urgent Upgrade Notes
	// ### (No, really, you MUST read this before you upgrade)
	// - actual notes...

	pattern := regexp.MustCompile(`(?s)##\s+Urgent Upgrade Notes.*?(?:##\s+[A-Z]|$)`)
	matches := pattern.FindStringSubmatch(changelog)

	if len(matches) > 0 {
		return strings.TrimSpace(matches[0])
	}

	return ""
}

// extractK8sPatchChanges extracts changes for a patch release
func extractK8sPatchChanges(changelog, version string) string {
	var sb strings.Builder

	// Find the section for this specific patch version
	version = strings.TrimPrefix(version, "v")
	versionHeader := fmt.Sprintf("# v%s", version)
	versionIdx := strings.Index(changelog, versionHeader)
	if versionIdx == -1 {
		return "No specific section found for this patch version. Patch releases rarely contain breaking changes.\n"
	}

	// Find the next version header
	nextVersionPattern := regexp.MustCompile(`\n# v\d+\.\d+\.\d+`)
	endIdx := len(changelog)
	if match := nextVersionPattern.FindStringIndex(changelog[versionIdx+1:]); match != nil {
		endIdx = versionIdx + 1 + match[0]
	}

	// Extract just this version's section
	versionSection := changelog[versionIdx:endIdx]

	// For patch releases, only look for critical keywords
	lines := strings.Split(versionSection, "\n")
	foundAny := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "BREAKING") ||
			strings.Contains(line, "ACTION REQUIRED") ||
			strings.Contains(line, "Urgent") {
			sb.WriteString(line)
			sb.WriteString("\n")
			foundAny = true
		}
	}

	if !foundAny {
		sb.WriteString("No breaking changes or urgent notes found in this patch release.\n")
	}

	return sb.String()
}

// extractK8sSections finds sections related to breaking changes
// Prioritizes the most critical sections: Deprecation and explicit breaking changes
func extractK8sSections(changelog string) string {
	var sb strings.Builder

	// Priority order: Most critical sections first
	// Focus on Deprecation (removals) and explicit breaking changes
	prioritySections := []string{
		"Deprecation", // Removals and deprecations - CRITICAL
		"Deprecated",
		"Breaking Change", // Explicit breaking changes
		"Breaking Changes",
		"Removed", // Already removed items
		"No Longer Supported",
	}

	// Lower priority: API changes (many are additive, not breaking)
	secondarySections := []string{
		"API Change",
		"API Changes",
	}

	// Extract priority sections first
	for _, section := range prioritySections {
		pattern := regexp.MustCompile(fmt.Sprintf(`(?s)(###?\s+%s.*?)(?:##|$)`, regexp.QuoteMeta(section)))
		matches := pattern.FindAllStringSubmatch(changelog, -1)

		for _, match := range matches {
			if len(match) > 1 {
				sb.WriteString(match[1])
				sb.WriteString("\n\n")
			}
		}
	}

	// Extract secondary sections (but limit to avoid noise)
	// Only extract first few API Changes to avoid overwhelming the context
	for _, section := range secondarySections {
		pattern := regexp.MustCompile(fmt.Sprintf(`(?s)(###?\s+%s.*?)(?:##|$)`, regexp.QuoteMeta(section)))
		matches := pattern.FindAllStringSubmatch(changelog, -1)

		if len(matches) > 0 && len(matches[0]) > 1 {
			// Only include the section header and a note
			sb.WriteString(fmt.Sprintf("### %s\n", section))
			sb.WriteString("(Full API Changes section available in changelog - focus on Deprecation section for breaking changes)\n\n")
			break // Only add note once
		}
	}

	return sb.String()
}
