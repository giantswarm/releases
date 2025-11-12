package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/giantswarm/releases/tools/check-upgrade-path/pkg"
)

func main() {
	var (
		providers  = flag.String("providers", "", "Comma-separated list of provider names (aws,azure,vsphere,cloud-director,eks)")
		version    = flag.String("version", "", "Release version to check (e.g., v30.1.0)")
		outputFile = flag.String("output", "upgrade-path-warning.md", "Output file for the warning report")
		workingDir = flag.String("workdir", ".", "Working directory (repository root)")
	)
	flag.Parse()

	if *providers == "" || *version == "" {
		fmt.Fprintf(os.Stderr, "Error: both -providers and -version are required\n")
		flag.Usage()
		os.Exit(1)
	}

	// Parse provider list
	providerList := strings.Split(*providers, ",")
	if len(providerList) == 0 {
		fmt.Fprintf(os.Stderr, "Error: at least one provider is required\n")
		os.Exit(1)
	}

	// Collect warnings from all providers
	var allWarnings []pkg.ProviderWarnings
	nextMajorVersion := ""

	for _, provider := range providerList {
		provider = strings.TrimSpace(provider)
		if provider == "" {
			continue
		}

		// Translate provider name to directory name
		providerDir := provider
		if provider == "aws" {
			providerDir = "capa"
		}

		checker := &pkg.UpgradePathChecker{
			Provider:   providerDir,
			Version:    *version,
			WorkingDir: *workingDir,
		}

		warnings, err := checker.Check()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Error checking upgrade path for %s: %v\n", provider, err)
			continue
		}

		if len(warnings) > 0 {
			// Store the next major version from the first provider
			if nextMajorVersion == "" {
				nextMajorVersion = checker.NextMajorVersion
			}

			allWarnings = append(allWarnings, pkg.ProviderWarnings{
				Provider: provider,
				Warnings: warnings,
			})
		}
	}

	if len(allWarnings) > 0 {
		report := pkg.GenerateConsolidatedReport(allWarnings, *version, nextMajorVersion)
		outputPath := filepath.Join(*workingDir, *outputFile)

		if err := os.WriteFile(outputPath, []byte(report), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing report: %v\n", err)
			os.Exit(1)
		}

		totalWarnings := 0
		for _, pw := range allWarnings {
			totalWarnings += len(pw.Warnings)
		}

		fmt.Printf("⚠️  Found %d upgrade path warnings across %d provider(s). Report written to %s\n",
			totalWarnings, len(allWarnings), outputPath)
		os.Exit(0)
	}

	fmt.Println("✅ No upgrade path issues detected")
	os.Exit(0)
}
