package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/giantswarm/releases/tools/check-upgrade-path/pkg"
)

func main() {
	var (
		provider   = flag.String("provider", "", "Provider name (aws, azure, vsphere, cloud-director, eks)")
		version    = flag.String("version", "", "Release version to check (e.g., v30.1.0)")
		outputFile = flag.String("output", "upgrade-path-warning.md", "Output file for the warning report")
		workingDir = flag.String("workdir", ".", "Working directory (repository root)")
	)
	flag.Parse()

	if *provider == "" || *version == "" {
		fmt.Fprintf(os.Stderr, "Error: both -provider and -version are required\n")
		flag.Usage()
		os.Exit(1)
	}

	// Translate provider name to directory name
	providerDir := *provider
	if *provider == "aws" {
		providerDir = "capa"
	}

	checker := &pkg.UpgradePathChecker{
		Provider:   providerDir,
		Version:    *version,
		WorkingDir: *workingDir,
	}

	warnings, err := checker.Check()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking upgrade path: %v\n", err)
		os.Exit(1)
	}

	if len(warnings) > 0 {
		report := pkg.GenerateReport(warnings, *version, checker.NextMajorVersion)
		outputPath := filepath.Join(*workingDir, *outputFile)

		if err := os.WriteFile(outputPath, []byte(report), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing report: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("⚠️  Found %d upgrade path warnings. Report written to %s\n", len(warnings), outputPath)
		os.Exit(0)
	}

	fmt.Println("✅ No upgrade path issues detected")
	os.Exit(0)
}
