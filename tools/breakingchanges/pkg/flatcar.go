package breakingchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type flatcarRelease struct {
	ReleaseNotes  string              `json:"release_notes"`
	ReleaseDate   string              `json:"release_date"`
	MajorSoftware map[string][]string `json:"major_software"`
	Channel       string              `json:"channel"`
}

// extractFlatcarReleaseNotes parses Flatcar's JSON and extracts release notes
func (d *Detector) extractFlatcarReleaseNotes(jsonContent, fromVersion, toVersion string) (string, error) {
	var releases map[string]flatcarRelease

	if err := json.Unmarshal([]byte(jsonContent), &releases); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Strip "v" prefix if present for lookup
	lookupVersion := strings.TrimPrefix(toVersion, "v")

	// Get the target release
	release, ok := releases[lookupVersion]
	if !ok {
		return "", fmt.Errorf("version %s not found in releases (tried %s)", toVersion, lookupVersion)
	}

	// Build comprehensive release notes
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# Flatcar Release Notes: v%s → v%s\n\n", fromVersion, toVersion))
	sb.WriteString(fmt.Sprintf("**IMPORTANT**: This upgrade spans from v%s to v%s. ", fromVersion, toVersion))
	sb.WriteString("Breaking changes may exist in ANY intermediate version.\n\n")

	// Fetch channel-level warnings from the HTML page (only relevant to this version range)
	fmt.Println("  → Fetching Flatcar channel warnings from HTML page...")
	fromVersionStripped := strings.TrimPrefix(fromVersion, "v")
	channelWarnings, err := d.fetchFlatcarChannelWarnings(context.Background(), fromVersionStripped, lookupVersion)
	if err != nil {
		fmt.Printf("  Warning: Could not fetch channel warnings: %v\n", err)
	} else if channelWarnings != "" {
		fmt.Printf("  ✓ Found relevant Flatcar channel warnings (%d bytes)\n", len(channelWarnings))
		sb.WriteString("## ⚠️ CHANNEL-LEVEL WARNINGS (relevant to this upgrade)\n\n")
		sb.WriteString(channelWarnings)
		sb.WriteString("\n\n")
	} else {
		fmt.Println("  ℹ️  No Flatcar channel warnings relevant to this upgrade")
	}

	sb.WriteString(fmt.Sprintf("## Target Version: %s\n\n", toVersion))
	sb.WriteString(fmt.Sprintf("Release Date: %s\n\n", release.ReleaseDate))

	if len(release.MajorSoftware) > 0 {
		sb.WriteString("### Major Software Versions\n\n")
		for name, versions := range release.MajorSoftware {
			sb.WriteString(fmt.Sprintf("- %s: %s\n", name, strings.Join(versions, ", ")))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("### Release Notes\n\n")
	sb.WriteString(release.ReleaseNotes)
	sb.WriteString("\n\n")

	return sb.String(), nil
}

// fetchFlatcarChannelWarnings fetches channel-level warnings from Flatcar HTML page
// and filters them to only show warnings relevant to the upgrade path (fromVersion -> toVersion)
func (d *Detector) fetchFlatcarChannelWarnings(ctx context.Context, fromVersion, toVersion string) (string, error) {
	htmlURL := "https://www.flatcar.org/releases"
	html, err := d.fetchURL(ctx, htmlURL)
	if err != nil {
		return "", err
	}

	// Extract warning sections from the HTML
	// Pattern: <h4 class="announcement">⚠️ ... ⚠️</h4><p class="announcement">...</p>
	warningPattern := regexp.MustCompile(`(?s)<h4[^>]*>\s*⚠️\s*(.*?)\s*⚠️\s*</h4>\s*<p[^>]*>(.*?)</p>`)
	matches := warningPattern.FindAllStringSubmatch(html, -1)

	if len(matches) == 0 {
		return "", nil
	}

	// Parse version numbers for comparison
	fromMajor := parseFlatcarMajor(fromVersion)
	toMajor := parseFlatcarMajor(toVersion)

	// Filter warnings relevant to this upgrade
	var sb strings.Builder
	for _, match := range matches {
		if len(match) > 2 {
			title := strings.TrimSpace(match[1])
			content := strings.TrimSpace(match[2])

			// Clean HTML entities and tags
			title = cleanHTMLEntities(title)
			content = cleanHTMLEntities(content)
			content = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(content, "")

			// Try to extract version numbers from the warning content
			// e.g., "in Stable 4230.2.0 release", "Flatcar 4081"
			versionPattern := regexp.MustCompile(`\b(\d{4}(?:\.\d+){0,2})\b`)
			versionMatches := versionPattern.FindAllString(content+" "+title, -1)

			// Check if any mentioned version falls within our upgrade range
			isRelevant := false
			if len(versionMatches) > 0 {
				for _, verStr := range versionMatches {
					verMajor := parseFlatcarMajor(verStr)
					// Include warning if the version it mentions is between from and to (exclusive of from)
					if verMajor > fromMajor && verMajor <= toMajor {
						isRelevant = true
						break
					}
				}
			} else {
				// If no version found in warning, include it to be safe
				isRelevant = true
			}

			if isRelevant {
				sb.WriteString(fmt.Sprintf("⚠️ **%s** ⚠️\n\n", title))
				sb.WriteString(content)
				sb.WriteString("\n\n")
			}
		}
	}

	return sb.String(), nil
}

// parseFlatcarMajor extracts the major version number from a Flatcar version string
// e.g., "4230.2.3" -> 4230
func parseFlatcarMajor(version string) int {
	// Remove any "v" prefix
	version = strings.TrimPrefix(version, "v")

	// Split by dots and take first part
	parts := strings.Split(version, ".")
	if len(parts) == 0 {
		return 0
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}

	return major
}

// cleanHTMLEntities removes common HTML entities
func cleanHTMLEntities(s string) string {
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&rsquo;", "'")
	s = strings.ReplaceAll(s, "&ldquo;", "\"")
	s = strings.ReplaceAll(s, "&rdquo;", "\"")
	return s
}
