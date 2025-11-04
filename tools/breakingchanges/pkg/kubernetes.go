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
	
	// Extract relevant subsections
	sb.WriteString(extractK8sSections(versionSection))
	sb.WriteString(extractK8sKeywordLines(versionSection))
	
	return sb.String()
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
func extractK8sSections(changelog string) string {
	var sb strings.Builder

	sections := []string{
		"API Change",
		"API Changes",
		"Removed",
		"Deprecation",
		"Deprecated",
		"Breaking Change",
		"Breaking Changes",
		"Urgent Upgrade Notes",
		"Known Issues",
		"No Longer Supported",
	}

	for _, section := range sections {
		// Pattern: ## Section or ### Section followed by content until next heading
		pattern := regexp.MustCompile(fmt.Sprintf(`(?s)(###?\s+%s.*?)(?:##|$)`, regexp.QuoteMeta(section)))
		matches := pattern.FindAllStringSubmatch(changelog, -1)

		for _, match := range matches {
			if len(match) > 1 {
				sb.WriteString(match[1])
				sb.WriteString("\n\n")
			}
		}
	}

	return sb.String()
}

// extractK8sKeywordLines finds individual lines with breaking change keywords
func extractK8sKeywordLines(changelog string) string {
	var sb strings.Builder

	lines := strings.Split(changelog, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "BREAKING") ||
			strings.Contains(line, "removed") ||
			strings.Contains(line, "deprecated") ||
			(strings.HasPrefix(line, "-") && (strings.Contains(line, "API") || strings.Contains(line, "removed"))) {
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
