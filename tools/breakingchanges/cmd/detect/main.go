package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	breakingchanges "github/giantswarm/releases/tools/breakingchanges/pkg"
)

func main() {
	ctx := context.Background()

	// Get environment variables
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Println("ANTHROPIC_API_KEY not set")
		os.Exit(1)
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Println("GITHUB_TOKEN not set")
		os.Exit(1)
	}

	prNumber := os.Getenv("PR_NUMBER")
	headSHA := os.Getenv("HEAD_SHA")

	// Initialize detector
	detector, err := breakingchanges.NewDetector(apiKey, githubToken)
	if err != nil {
		fmt.Printf("Failed to initialize detector: %v\n", err)
		os.Exit(1)
	}

	// Get the repository root (we're running from tools/breakingchanges)
	repoRoot := "../.."

	// Detect changed releases in the PR
	releases, err := detectChangedReleases(repoRoot)
	if err != nil {
		fmt.Printf("Failed to detect changed releases: %v\n", err)
		os.Exit(1)
	}

	if len(releases) == 0 {
		fmt.Println("No release changes detected")
		os.Exit(0)
	}

	fmt.Printf("Analyzing %d release(s)...\n", len(releases))

	// Analyze each release and track which provider each finding came from
	var allFindings []breakingchanges.FindingWithProvider
	
	for _, release := range releases {
		fmt.Printf("Analyzing %s...\n", release.Path)

		findings, err := detector.AnalyzeRelease(ctx, release)
		if err != nil {
			fmt.Printf("Warning: Failed to analyze %s: %v\n", release.Path, err)
			continue
		}

		for _, finding := range findings {
			allFindings = append(allFindings, breakingchanges.FindingWithProvider{
				Finding:  finding,
				Provider: release.Provider,
				Version:  release.Version,
			})
		}
	}

	// Generate report with provider grouping
	report := breakingchanges.GenerateReportWithProviders(allFindings, prNumber, headSHA)

	// Write report to file
	err = os.WriteFile("breaking-changes-report.md", []byte(report), 0644)
	if err != nil {
		fmt.Printf("Failed to write report: %v\n", err)
		os.Exit(1)
	}

	// Set GitHub Actions output
	if len(allFindings) > 0 {
		fmt.Println("::set-output name=has_findings::true")
	} else {
		fmt.Println("::set-output name=has_findings::false")
	}

	fmt.Printf("Analysis complete. Found %d potential breaking changes.\n", len(allFindings))
}

func detectChangedReleases(repoRoot string) ([]breakingchanges.Release, error) {
	var releases []breakingchanges.Release

	providers := []string{"aws", "azure", "capa", "cloud-director", "eks", "kvm", "vsphere"}

	for _, provider := range providers {
		providerPath := filepath.Join(repoRoot, provider)

		entries, err := os.ReadDir(providerPath)
		if err != nil {
			continue // Provider directory might not exist
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			// Check if it's a version directory (starts with 'v')
			if !strings.HasPrefix(entry.Name(), "v") {
				continue
			}

			// Skip archived releases
			if entry.Name() == "archived" {
				continue
			}

			releasePath := filepath.Join(providerPath, entry.Name())

			// Check if README exists (indicator of a real release)
			readmePath := filepath.Join(releasePath, "README.md")
			if _, err := os.Stat(readmePath); os.IsNotExist(err) {
				continue
			}

			// Read files
			readme, _ := os.ReadFile(readmePath)
			diff, _ := os.ReadFile(filepath.Join(releasePath, "release.diff"))
			yaml, _ := os.ReadFile(filepath.Join(releasePath, "release.yaml"))

			releases = append(releases, breakingchanges.Release{
				Provider: provider,
				Version:  entry.Name(),
				Path:     releasePath,
				README:   string(readme),
				Diff:     string(diff),
				YAML:     string(yaml),
			})
		}
	}

	return releases, nil
}
