package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Provider struct {
	id         string // ID: aws, eks, azure, vsphere, cloud-director
	repository string // Cluster Chart Repository: cluster-aws, cluster-eks, cluster-azure, cluster-vsphere, cluster-cloud-director
	directory  string // Releases Directory: capa, eks, azure, vsphere, cloud-director
}

var (
	SharedAppNameRe              = regexp.MustCompile(`^appName:\s*(\S+)`)
	SharedAppConfigKeyRe         = regexp.MustCompile(`^configKey:\s*(\S+)`)
	SharedAppDependsOnRe         = regexp.MustCompile(`^dependsOn:\s*$`)
	SharedAppDependencyRe        = regexp.MustCompile(`^-\s+(\S.*)$`)
	SharedAppProviderConditionRe = regexp.MustCompile(`\{\{-?\s*if\s+(eq|ne)\s+\$?\.?Values\.providerIntegration\.provider\s+"([^"]+)"\s*\}\}`)

	TemplateAppSeparatorRe  = regexp.MustCompile(`(?m)^---\s*$`)
	TemplateAppNameRe       = regexp.MustCompile(`set\s+\$\s+"appName"\s+"([^"]+)"`)
	TemplateAppKindRe       = regexp.MustCompile(`(?m)^kind:\s*HelmRelease\s*$`)
	TemplateAppDependsOnRe  = regexp.MustCompile(`^(\s*)dependsOn:\s*$`)
	TemplateAppDependencyRe = regexp.MustCompile(`-\s*name:\s*\{\{\s*include\s+"resource\.default\.name"\s+[.$]\s*\}\}-([\w-]+)`)
)

type App struct {
	name         string
	dependencies map[string]bool
	source       string // "shared:<file>" or "template:<file>"
}

var VersionRe = regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)$`)

type ReleaseDocument struct {
	Spec ReleaseSpec `yaml:"spec"`
}

type ReleaseSpec struct {
	Apps []ReleaseApp `yaml:"apps"`
}

type ReleaseApp struct {
	Name         string   `yaml:"name"`
	Dependencies []string `yaml:"dependsOn"`
}

// Apps allowed to be missing from a release, keyed by releases directory (e.g. "capa" for AWS).
var KnownExceptions = map[string]map[string]bool{
	"vsphere": {"kamaji-etcd": true},
}

func main() {
	base := flag.String("base", defaultBase(), "Directory containing the `releases`, `cluster` and `cluster-<provider>` repositories as siblings.")
	only := flag.String("providers", "", "Comma-separated list of providers to check (default: auto-discover all `cluster-<provider>` repositories found in base).")
	version := flag.String("version", "", "Release to check, e.g. v35.0.0 (default: the latest/highest vX.Y.Z directory found for each provider)")
	flag.Parse()

	providers, err := discoverProviders(*base)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR]", err)
		os.Exit(2)
	}

	if *only != "" {
		wanted := map[string]bool{}
		for provider := range strings.SplitSeq(*only, ",") {
			wanted[strings.TrimSpace(provider)] = true
		}
		filtered := providers[:0]
		for _, provider := range providers {
			if wanted[provider.id] {
				filtered = append(filtered, provider)
			}
		}
		providers = filtered
	}

	if len(providers) == 0 {
		fmt.Fprintln(os.Stderr, "[ERROR] No providers found.")
		os.Exit(2)
	}

	sharedDir := filepath.Join(*base, "cluster", "helm", "cluster", "files", "helmreleases")

	failed := false
	for _, provider := range providers {
		ok, err := checkProvider(*base, sharedDir, provider, *version)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] %s (%s): %v\n", provider.id, provider.repository, err)
			failed = true
			continue
		}
		if !ok {
			failed = true
		}
	}

	if failed {
		os.Exit(1)
	}
}

func defaultBase() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return filepath.Join(home, "Documents")
}

func discoverProviders(base string) ([]Provider, error) {
	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, fmt.Errorf("reading base dir %s: %w", base, err)
	}

	var providers []Provider
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), "cluster-") || !entry.IsDir() {
			continue
		}

		repository := entry.Name()
		id := strings.TrimPrefix(repository, "cluster-")

		values := filepath.Join(base, repository, "helm", repository, "values.yaml")
		if _, err := os.Stat(values); err != nil {
			continue
		}

		directory := id
		if id == "aws" {
			directory = "capa"
		}

		providers = append(providers, Provider{
			id:         id,
			repository: repository,
			directory:  directory,
		})
	}

	sort.Slice(providers, func(i, j int) bool { return providers[i].id < providers[j].id })
	return providers, nil
}

// --- STEP 1: Provider Chart Values ---

func readProviderValues(base string, provider Provider) (providerID string, configKeys map[string]bool, err error) {
	path := filepath.Join(base, provider.repository, "helm", provider.repository, "values.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var document map[string]any
	if err := yaml.Unmarshal(data, &document); err != nil {
		return "", nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	cluster, ok := document["cluster"].(map[string]any)
	if !ok {
		return "", nil, fmt.Errorf("%s: missing top-level 'cluster' key", path)
	}

	providerIntegration, ok := cluster["providerIntegration"].(map[string]any)
	if !ok {
		return "", nil, fmt.Errorf("%s: missing 'cluster.providerIntegration'", path)
	}

	providerID, _ = providerIntegration["provider"].(string)
	if providerID == "" {
		return "", nil, fmt.Errorf("%s: missing/empty 'cluster.providerIntegration.provider'", path)
	}

	configKeys = map[string]bool{}
	if apps, ok := providerIntegration["apps"].(map[string]any); ok {
		for key, value := range apps {
			if app, ok := value.(map[string]any); ok {
				if enable, _ := app["enable"].(bool); enable {
					configKeys[key] = true
				}
			}
		}
	}

	return providerID, configKeys, nil
}

// --- STEP 2: Shared Cluster Chart HelmRelease Files ---

func kebabToCamel(kebab string) string {
	parts := strings.Split(kebab, "-")
	if len(parts) == 1 {
		return kebab
	}

	camel := parts[0]
	for _, part := range parts[1:] {
		if part == "" {
			continue
		}
		camel += strings.ToUpper(part[:1]) + part[1:]
	}

	return camel
}

func loadSharedHelmReleases(sharedDir, providerID string) (map[string]App, error) {
	files, err := filepath.Glob(filepath.Join(sharedDir, "*.yaml"))
	if err != nil {
		return nil, err
	}

	apps := map[string]App{}
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", path, err)
		}

		lines := strings.Split(string(data), "\n")

		var appName, configKey string
		for _, line := range lines {
			if match := SharedAppNameRe.FindStringSubmatch(line); match != nil {
				appName = match[1]
			}
			if match := SharedAppConfigKeyRe.FindStringSubmatch(line); match != nil {
				configKey = match[1]
			}
		}

		if appName == "" {
			continue
		}
		if configKey == "" {
			configKey = kebabToCamel(appName)
		}

		dependencies := map[string]bool{}
		for i := range len(lines) {
			if !SharedAppDependsOnRe.MatchString(lines[i]) {
				continue
			}

			j := i - 1
			for j >= 0 && strings.TrimSpace(lines[j]) == "" {
				j--
			}
			if j >= 0 {
				if condition := SharedAppProviderConditionRe.FindStringSubmatch(lines[j]); condition != nil {
					operator, value := condition[1], condition[2]
					if operator == "eq" && providerID != value || operator == "ne" && providerID == value {
						continue
					}
				}
			}

			for k := i + 1; k < len(lines); k++ {
				dependency := SharedAppDependencyRe.FindStringSubmatch(lines[k])
				if dependency == nil {
					break
				}
				dependencies[strings.TrimSpace(dependency[1])] = true
			}
		}

		apps[configKey] = App{
			name:         appName,
			dependencies: dependencies,
			source:       "shared:" + filepath.Base(path),
		}
	}

	return apps, nil
}

// --- STEP 3: Provider-specific HelmRelease CRs ---

func findTemplateHelmReleases(base string, provider Provider) ([]App, error) {
	templates := filepath.Join(base, provider.repository, "helm", provider.repository, "templates")

	var apps []App
	err := filepath.Walk(templates, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		content := string(data)
		documents := TemplateAppSeparatorRe.Split(content, -1)
		var appName string

		for _, document := range documents {
			if match := TemplateAppNameRe.FindStringSubmatch(document); match != nil {
				appName = match[1]
			}

			if !TemplateAppKindRe.MatchString(document) {
				continue
			}

			dependencies := extractDependencies(document)
			apps = append(apps, App{
				name:         appName,
				dependencies: dependencies,
				source:       "template:" + filepath.Base(path),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func extractDependencies(document string) map[string]bool {
	dependencies := map[string]bool{}

	lines := strings.Split(document, "\n")
	for i, line := range lines {
		dependsOn := TemplateAppDependsOnRe.FindStringSubmatch(line)
		if dependsOn == nil {
			continue
		}

		indent := len(dependsOn[1])

		var block []string
		for k := i + 1; k < len(lines); k++ {
			nextLine := lines[k]
			if strings.TrimSpace(nextLine) == "" {
				continue
			}

			currentIndent := len(nextLine) - len(strings.TrimLeft(nextLine, " "))
			if currentIndent < indent {
				break
			}
			if currentIndent == indent && !strings.HasPrefix(strings.TrimSpace(nextLine), "-") {
				break
			}

			block = append(block, nextLine)
		}

		blockText := strings.Join(block, "\n")
		for _, dependency := range TemplateAppDependencyRe.FindAllStringSubmatch(blockText, -1) {
			dependencies[dependency[1]] = true
		}
	}

	return dependencies
}

// --- STEP 5: Release ---

func latestReleaseVersion(directory string) (string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", directory, err)
	}

	var versions []string
	for _, entry := range entries {
		if entry.IsDir() && VersionRe.MatchString(entry.Name()) {
			versions = append(versions, entry.Name())
		}
	}

	if len(versions) == 0 {
		return "", fmt.Errorf("no vX.Y.Z directories found under %s", directory)
	}

	sort.Slice(versions, func(i, j int) bool { return compareVersions(versions[i], versions[j]) < 0 })
	return versions[len(versions)-1], nil
}

func compareVersions(a, b string) int {
	pa := VersionRe.FindStringSubmatch(a)
	pb := VersionRe.FindStringSubmatch(b)

	for i := 1; i <= 3; i++ {
		na, _ := strconv.Atoi(pa[i])
		nb, _ := strconv.Atoi(pb[i])

		if na != nb {
			if na < nb {
				return -1
			}
			return 1
		}
	}

	return 0
}

func loadReleaseApps(base string, provider Provider, requestedVersion string) (version string, apps map[string][]string, err error) {
	directory := filepath.Join(base, "releases", provider.directory)

	if requestedVersion == "" {
		version, err = latestReleaseVersion(directory)
		if err != nil {
			return "", nil, err
		}
	} else {
		if !VersionRe.MatchString(requestedVersion) {
			return "", nil, fmt.Errorf("invalid -version %q, expected format vX.Y.Z", requestedVersion)
		}

		if info, statErr := os.Stat(filepath.Join(directory, requestedVersion)); statErr != nil || !info.IsDir() {
			return "", nil, fmt.Errorf("release version %s not found under %s", requestedVersion, directory)
		}

		version = requestedVersion
	}

	path := filepath.Join(directory, version, "release.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var document ReleaseDocument
	if err := yaml.Unmarshal(data, &document); err != nil {
		return "", nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	apps = map[string][]string{}
	for _, app := range document.Spec.Apps {
		apps[app.Name] = app.Dependencies
	}

	return version, apps, nil
}

// --- STEP 4 & 6: Build expected list & compare ---

func checkProvider(base, sharedDir string, provider Provider, requestedVersion string) (bool, error) {
	fmt.Printf("[INFO] %s (%s / releases/%s)\n", provider.id, provider.repository, provider.directory)

	providerID, configKeys, err := readProviderValues(base, provider)
	if err != nil {
		return false, err
	}

	sharedApps, err := loadSharedHelmReleases(sharedDir, providerID)
	if err != nil {
		return false, err
	}

	ok := true

	var apps []App
	for configKey := range configKeys {
		app, found := sharedApps[configKey]
		if !found {
			fmt.Printf("[WARN] Enabled app %q has no matching shared HelmRelease file.\n", configKey)
			ok = false
			continue
		}
		apps = append(apps, app)
	}

	templateApps, err := findTemplateHelmReleases(base, provider)
	if err != nil {
		return false, err
	}

	for _, app := range templateApps {
		if app.name == "" {
			fmt.Printf("[WARN] Found a HelmRelease CR in %s with no resolvable name.\n", app.source)
			ok = false
			continue
		}
		apps = append(apps, app)
	}

	appsByName := map[string][]App{}
	for _, app := range apps {
		appsByName[app.name] = append(appsByName[app.name], app)
	}

	var names []string
	for name := range appsByName {
		names = append(names, name)
	}
	sort.Strings(names)

	appMap := map[string]map[string]bool{}
	for _, name := range names {
		duplicates := appsByName[name]
		first := duplicates[0]
		for _, other := range duplicates[1:] {
			if !sameSet(first.dependencies, other.dependencies) {
				fmt.Printf("[WARN] App %q dependencies differs between %s (%s) and %s (%s)\n", name, first.source, setToString(first.dependencies), other.source, setToString(other.dependencies))
				ok = false
			}
		}
		appMap[name] = first.dependencies
	}

	version, releaseApps, err := loadReleaseApps(base, provider, requestedVersion)
	if err != nil {
		return false, err
	}

	fmt.Printf("[INFO] Version: %s | Provider Apps: %d | Release Apps: %d\n", version, len(appMap), len(releaseApps))

	exceptions := KnownExceptions[provider.directory]

	var missing []string
	for name := range appMap {
		if _, found := releaseApps[name]; !found && !exceptions[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)

	if len(missing) > 0 {
		fmt.Printf("[WARN] Missing from release: %v\n", missing)
		ok = false
	}

	var additional []string
	for name := range releaseApps {
		if _, found := appMap[name]; !found {
			additional = append(additional, name)
		}
	}
	sort.Strings(additional)

	if len(additional) > 0 {
		fmt.Printf("[WARN] Additional in release: %v\n", additional)
		ok = false
	}

	var mismatches []string
	for name, dependencies := range appMap {
		releaseDependencies, found := releaseApps[name]
		if !found {
			continue
		}

		releaseDependencySet := map[string]bool{}
		for _, dependency := range releaseDependencies {
			releaseDependencySet[dependency] = true
		}

		if !sameSet(dependencies, releaseDependencySet) {
			mismatches = append(mismatches, fmt.Sprintf("- %s: expected %s, got %s", name, setToString(dependencies), setToString(releaseDependencySet)))
		}
	}
	sort.Strings(mismatches)

	if len(mismatches) > 0 {
		fmt.Println("[WARN] Dependency mismatches:")
		for _, mismatch := range mismatches {
			fmt.Println(mismatch)
		}
		ok = false
	}

	if ok {
		fmt.Println("[INFO] Exact match!")
	}

	return ok, nil
}

func sameSet(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func setToString(s map[string]bool) string {
	var items []string
	for k := range s {
		items = append(items, k)
	}
	sort.Strings(items)
	return "[" + strings.Join(items, ", ") + "]"
}
