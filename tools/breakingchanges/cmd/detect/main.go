package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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
	var changedFiles []string

	// Try to get changed files from GitHub PR API first (most reliable in CI)
	prNumber := os.Getenv("PR_NUMBER")
	githubToken := os.Getenv("GITHUB_TOKEN")

	if prNumber != "" && githubToken != "" {
		fmt.Printf("Fetching changed files from PR #%s via GitHub API...\n", prNumber)
		files, err := fetchPRChangedFiles(prNumber, githubToken)
		if err != nil {
			fmt.Printf("Warning: Failed to fetch PR files from API: %v\n", err)
		} else {
			changedFiles = files
			fmt.Printf("Found %d changed files from PR\n", len(changedFiles))
		}
	}

	// Fallback to git diff if API fetch failed or not in CI
	if len(changedFiles) == 0 {
		fmt.Println("Using git diff to detect changed files...")
		cmd := exec.Command("git", "diff", "--name-only", "origin/master", "HEAD")
		cmd.Dir = repoRoot
		output, err := cmd.Output()
		if err != nil {
			// If that fails (shallow clone), try HEAD~1
			cmd = exec.Command("git", "diff", "--name-only", "HEAD~1", "HEAD")
			cmd.Dir = repoRoot
			output, err = cmd.Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run git diff: %w", err)
			}
		}
		changedFiles = strings.Split(string(output), "\n")
	}

	// Parse changed files and extract unique release directories
	releaseMap := make(map[string]bool) // Use map to deduplicate
	validProviders := map[string]bool{
		"azure": true, "capa": true, "cloud-director": true,
		"eks": true, "vsphere": true,
	}

	fmt.Println("Analyzing changed files for release directories...")
	for _, file := range changedFiles {
		if file == "" {
			continue
		}

		// Match pattern: provider/vX.Y.Z/...
		parts := strings.Split(file, "/")
		if len(parts) < 2 {
			continue
		}

		provider := parts[0]
		version := parts[1]

		// Check if it's a valid provider directory
		if !validProviders[provider] {
			continue
		}

		// Check if it's a version directory (starts with 'v')
		if !strings.HasPrefix(version, "v") {
			continue
		}

		// Skip archived releases
		if version == "archived" {
			continue
		}

		// Create unique key for this release
		releaseKey := provider + "/" + version
		if !releaseMap[releaseKey] {
			fmt.Printf("  Detected release: %s\n", releaseKey)
			releaseMap[releaseKey] = true
		}
	}

	fmt.Printf("Total unique releases detected: %d\n", len(releaseMap))

	// Load release data for each changed release
	for releaseKey := range releaseMap {
		parts := strings.Split(releaseKey, "/")
		provider := parts[0]
		version := parts[1]

		releasePath := filepath.Join(repoRoot, provider, version)

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
			Version:  version,
			Path:     releasePath,
			README:   string(readme),
			Diff:     string(diff),
			YAML:     string(yaml),
		})
	}

	return releases, nil
}

// fetchPRChangedFiles fetches the list of changed files from a GitHub PR
func fetchPRChangedFiles(prNumber, token string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/giantswarm/releases/pulls/%s/files", prNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var prFiles []struct {
		Filename string `json:"filename"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&prFiles); err != nil {
		return nil, err
	}

	files := make([]string, len(prFiles))
	for i, f := range prFiles {
		files[i] = f.Filename
	}

	return files, nil
}
