package breakingchanges

import (
	"fmt"
	"regexp"
	"strings"
)

// extractK8sBreakingChanges extracts breaking change sections from K8s changelog
func (d *Detector) extractK8sBreakingChanges(changelog string) string {
	var sb strings.Builder

	// Extract relevant sections
	sb.WriteString(extractK8sSections(changelog))

	// Also look for lines with specific keywords
	sb.WriteString(extractK8sKeywordLines(changelog))

	result := sb.String()
	if len(result) == 0 {
		return changelog // Fallback to full changelog
	}

	return "# Kubernetes Breaking Changes and Deprecations\n\n" + result
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
