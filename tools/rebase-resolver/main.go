// rebase-resolver auto-resolves conflicts in */releases.json and
// */kustomization.yaml files left behind by `git rebase origin/master`.
//
// Both files are ordered lists keyed by semver. A conflict between two
// release PRs is almost always a bookkeeping clash, not a semantic one, so
// the correct merged result is the union of entries sorted by semver.
//
// The tool reads the "ours" (stage 2) and "theirs" (stage 3) versions from
// the git index, merges them, writes the result, and `git add`s the file.
// Any conflict in a file this tool doesn't know how to resolve is left
// alone and reported on stderr with a non-zero exit code.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/blang/semver/v4"
	"gopkg.in/yaml.v3"
)

type releaseEntry struct {
	Version          string `json:"version"`
	IsDeprecated     bool   `json:"isDeprecated"`
	ReleaseTimestamp string `json:"releaseTimestamp"`
	ChangelogUrl     string `json:"changelogUrl"`
	IsStable         bool   `json:"isStable"`
}

type releasesFile struct {
	Releases []releaseEntry `json:"releases"`
}

func main() {
	conflicted, err := conflictedFiles()
	if err != nil {
		fail("list conflicted files: %v", err)
	}
	if len(conflicted) == 0 {
		return
	}

	var unhandled []string
	for _, path := range conflicted {
		base := filepath.Base(path)
		switch base {
		case "releases.json":
			if err := resolveReleasesJSON(path); err != nil {
				fail("resolve %s: %v", path, err)
			}
		case "kustomization.yaml":
			if err := resolveKustomization(path); err != nil {
				fail("resolve %s: %v", path, err)
			}
		default:
			unhandled = append(unhandled, path)
		}
	}

	if len(unhandled) > 0 {
		fmt.Fprintln(os.Stderr, "unresolved conflicts:")
		for _, p := range unhandled {
			fmt.Fprintln(os.Stderr, "  "+p)
		}
		os.Exit(2)
	}
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func conflictedFiles() ([]string, error) {
	out, err := exec.Command("git", "diff", "--name-only", "--diff-filter=U").Output()
	if err != nil {
		return nil, err
	}
	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			files = append(files, line)
		}
	}
	return files, nil
}

// readStage returns the given git index stage of path:
//
//	1 = common ancestor, 2 = ours, 3 = theirs
func readStage(path string, stage int) ([]byte, error) {
	var stderr bytes.Buffer
	cmd := exec.Command("git", "show", fmt.Sprintf(":%d:%s", stage, path))
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git show :%d:%s: %v: %s", stage, path, err, stderr.String())
	}
	return out, nil
}

func gitAdd(path string) error {
	return exec.Command("git", "add", "--", path).Run()
}

// resolveReleasesJSON merges the "ours" and "theirs" copies of
// releases.json by union-of-entries (keyed on version), sorted ascending
// by semver, and writes the result back to disk.
func resolveReleasesJSON(path string) error {
	ours, err := readStage(path, 2)
	if err != nil {
		return err
	}
	theirs, err := readStage(path, 3)
	if err != nil {
		return err
	}

	merged, err := mergeReleasesJSON(ours, theirs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, merged, 0644); err != nil {
		return err
	}
	return gitAdd(path)
}

func mergeReleasesJSON(ours, theirs []byte) ([]byte, error) {
	var a, b releasesFile
	if err := json.Unmarshal(ours, &a); err != nil {
		return nil, fmt.Errorf("parse ours: %v", err)
	}
	if err := json.Unmarshal(theirs, &b); err != nil {
		return nil, fmt.Errorf("parse theirs: %v", err)
	}

	byVersion := make(map[string]releaseEntry)
	for _, r := range a.Releases {
		byVersion[r.Version] = r
	}
	// On same-version conflict, prefer "theirs" (the PR's version) since
	// that's the author's intent and rebase semantics.
	for _, r := range b.Releases {
		byVersion[r.Version] = r
	}

	merged := make([]releaseEntry, 0, len(byVersion))
	for _, r := range byVersion {
		merged = append(merged, r)
	}
	sort.Slice(merged, func(i, j int) bool {
		return semverLess(merged[i].Version, merged[j].Version)
	})

	out := releasesFile{Releases: merged}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(out); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// resolveKustomization merges the "ours" and "theirs" copies of
// kustomization.yaml. Only the `resources` list is re-merged; all other
// fields are taken from "ours" (master), because kustomization
// scaffolding changes are made separately from release adds and should
// not be re-derived here.
func resolveKustomization(path string) error {
	ours, err := readStage(path, 2)
	if err != nil {
		return err
	}
	theirs, err := readStage(path, 3)
	if err != nil {
		return err
	}

	merged, err := mergeKustomization(ours, theirs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, merged, 0644); err != nil {
		return err
	}
	return gitAdd(path)
}

// mergeKustomization rewrites only the top-level `resources:` list block
// of "ours", replacing it with the semver-sorted union of resources from
// both sides. Every byte outside that block is preserved verbatim, so
// comments, formatting, and the quirk of column-zero list items are
// retained — avoiding cosmetic diffs on every rebase.
func mergeKustomization(ours, theirs []byte) ([]byte, error) {
	oursResources, err := extractResources(ours)
	if err != nil {
		return nil, fmt.Errorf("parse ours: %v", err)
	}
	theirsResources, err := extractResources(theirs)
	if err != nil {
		return nil, fmt.Errorf("parse theirs: %v", err)
	}

	union := sortedUnion(append(oursResources, theirsResources...))
	return rewriteResourcesBlock(ours, union)
}

func extractResources(data []byte) ([]string, error) {
	var doc struct {
		Resources []string `yaml:"resources"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	return doc.Resources, nil
}

// rewriteResourcesBlock locates the top-level `resources:` key in a
// kustomization.yaml document and replaces the contiguous run of list
// items below it with the provided resources. It preserves the original
// list-item indentation (whatever prefix like "- " or "  - " the file
// already uses) and leaves the rest of the file byte-for-byte identical.
func rewriteResourcesBlock(doc []byte, resources []string) ([]byte, error) {
	// Keep the trailing-newline status of the input so the output matches.
	hasTrailingNL := bytes.HasSuffix(doc, []byte("\n"))
	content := doc
	if hasTrailingNL {
		content = doc[:len(doc)-1]
	}
	lines := strings.Split(string(content), "\n")

	startIdx := -1
	for i, l := range lines {
		// Top-level `resources:` only (no leading whitespace).
		if l == "resources:" || strings.HasPrefix(l, "resources:") && !isIndented(l) {
			startIdx = i
			break
		}
	}
	if startIdx == -1 {
		return nil, fmt.Errorf("resources: key not found")
	}

	// Determine the item prefix from the first existing list item, or
	// default to "- " (column-zero) which matches the repo convention.
	itemPrefix := "- "
	endIdx := startIdx + 1
	for endIdx < len(lines) {
		l := lines[endIdx]
		trimmed := strings.TrimLeft(l, " \t")
		if !strings.HasPrefix(trimmed, "- ") && trimmed != "-" {
			break
		}
		if endIdx == startIdx+1 {
			// Record the exact prefix the file uses.
			itemPrefix = l[:len(l)-len(trimmed)] + "- "
		}
		endIdx++
	}

	var b bytes.Buffer
	for i := 0; i <= startIdx; i++ {
		b.WriteString(lines[i])
		b.WriteByte('\n')
	}
	for _, r := range resources {
		b.WriteString(itemPrefix)
		b.WriteString(r)
		b.WriteByte('\n')
	}
	for i := endIdx; i < len(lines); i++ {
		b.WriteString(lines[i])
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}
	if hasTrailingNL {
		b.WriteByte('\n')
	}
	return b.Bytes(), nil
}

func isIndented(s string) bool {
	return len(s) > 0 && (s[0] == ' ' || s[0] == '\t')
}

func sortedUnion(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, s := range items {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool {
		return semverLess(out[i], out[j])
	})
	return out
}

// semverLess compares two version strings lexicographically by semver.
// It strips an optional leading "v" and falls back to string comparison
// for anything that doesn't parse — this keeps non-semver resource
// entries (if they ever appear) from panicking the tool.
func semverLess(a, b string) bool {
	av, aok := parseSemver(a)
	bv, bok := parseSemver(b)
	switch {
	case aok && bok:
		return av.LT(bv)
	case aok:
		return true
	case bok:
		return false
	default:
		return a < b
	}
}

func parseSemver(s string) (semver.Version, bool) {
	v, err := semver.ParseTolerant(strings.TrimPrefix(s, "v"))
	if err != nil {
		return semver.Version{}, false
	}
	return v, true
}
