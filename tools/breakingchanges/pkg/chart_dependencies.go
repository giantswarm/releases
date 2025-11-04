package breakingchanges

import "regexp"

// extractChartDependencies parses Chart.yaml changes from a diff to find dependency version changes
func (d *Detector) extractChartDependencies(diff, parentComponent string) []VersionChange {
	var dependencies []VersionChange

	// Look for Chart.yaml changes in the diff
	chartYamlPattern := regexp.MustCompile(`(?s)diff --git a/.*?Chart\.yaml.*?@@.*?@@(.*?)(?:diff --git|$)`)
	matches := chartYamlPattern.FindAllStringSubmatch(diff, -1)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		chartContent := match[1]
		dependencies = append(dependencies, extractDependencyPattern1(chartContent)...)
		dependencies = append(dependencies, extractDependencyPattern2(chartContent)...)
	}

	return dependencies
}

// extractDependencyPattern1 looks for pattern: - name: X followed by - version: Y
func extractDependencyPattern1(chartContent string) []VersionChange {
	var dependencies []VersionChange

	// Pattern: - name: cluster followed by - version: 1.2.3 and + version: 2.0.0
	depPattern := regexp.MustCompile(`(?m)^\s*-\s*name:\s*(\S+)\s*\n.*?-\s*version:\s*["]?v?([\d.]+)["]?\s*\n.*?\+\s*version:\s*["]?v?([\d.]+)["]?`)
	matches := depPattern.FindAllStringSubmatch(chartContent, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			dependencies = append(dependencies, VersionChange{
				Component:   match[1],
				FromVersion: match[2],
				ToVersion:   match[3],
			})
		}
	}

	return dependencies
}

// extractDependencyPattern2 looks for alternative pattern when name and version are separate
func extractDependencyPattern2(chartContent string) []VersionChange {
	var dependencies []VersionChange

	// Alternative pattern for when name and version are on separate lines
	altPattern := regexp.MustCompile(`(?m)-\s*version:\s*["]?v?([\d.]+)["]?(?:.*?\n){0,3}.*?\+\s*version:\s*["]?v?([\d.]+)["]?`)
	altMatches := altPattern.FindAllStringSubmatch(chartContent, -1)

	// Try to find the dependency name near these changes
	if len(altMatches) == 0 {
		return dependencies
	}

	namePattern := regexp.MustCompile(`name:\s*(\S+)`)
	nameMatches := namePattern.FindAllStringSubmatch(chartContent, -1)

	for i, altMatch := range altMatches {
		if len(altMatch) >= 3 && i < len(nameMatches) && len(nameMatches[i]) >= 2 {
			depName := nameMatches[i][1]

			// Avoid duplicates
			if !isDuplicate(dependencies, depName) {
				dependencies = append(dependencies, VersionChange{
					Component:   depName,
					FromVersion: altMatch[1],
					ToVersion:   altMatch[2],
				})
			}
		}
	}

	return dependencies
}

// isDuplicate checks if a dependency already exists in the list
func isDuplicate(dependencies []VersionChange, name string) bool {
	for _, existing := range dependencies {
		if existing.Component == name {
			return true
		}
	}
	return false
}
