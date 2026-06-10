// rebase-resolver auto-resolves conflicts in */releases.json and
// */kustomization.yaml files left behind by `git rebase origin/master`.
//
// Both files are ordered lists keyed by semver. A conflict between two
// release PRs is almost always a bookkeeping clash, not a semantic one,
// so the correct merged result is a 3-way merge against the common
// ancestor: the union of entries minus anything that either side
// removed relative to the ancestor, sorted by semver.
//
// The tool reads the "base" (stage 1, common ancestor), "ours"
// (stage 2) and "theirs" (stage 3) versions from the git index, merges
// them, writes the result, and `git add`s the file. Any conflict in a
// file this tool doesn't know how to resolve is left alone and
// reported on stderr with a non-zero exit code.
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

// releasesFile mirrors the on-disk shape of every <provider>/releases.json
// in this repo. The trailing metadata fields after `releases` are part of
// the file and must round-trip; otherwise every rebase silently strips
// them.
type releasesFile struct {
	Releases     []releaseEntry `json:"releases"`
	SourceUrl    string         `json:"sourceUrl,omitempty"`
	ChangelogUrl string         `json:"changelogUrl,omitempty"`
	Homepage     string         `json:"homepage,omitempty"`
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

// readStageOptional returns the stage's contents if present, or nil if
// the stage doesn't exist (the typical case for stage 1 when the file
// was added on both sides without a common ancestor).
func readStageOptional(path string, stage int) ([]byte, error) {
	cmd := exec.Command("git", "show", fmt.Sprintf(":%d:%s", stage, path))
	cmd.Stderr = nil // swallow "exists on disk but not in the index" noise
	out, err := cmd.Output()
	if err != nil {
		// Treat any error as "stage absent". The downstream merge falls
		// back to 2-way union, which is the right behavior when there's
		// no ancestor to compare against.
		return nil, nil
	}
	return out, nil
}

func gitAdd(path string) error {
	return exec.Command("git", "add", "--", path).Run()
}

// resolveReleasesJSON 3-way-merges the ancestor, ours, and theirs copies
// of releases.json and writes the result back to disk.
func resolveReleasesJSON(path string) error {
	ancestor, err := readStageOptional(path, 1)
	if err != nil {
		return err
	}
	ours, err := readStage(path, 2)
	if err != nil {
		return err
	}
	theirs, err := readStage(path, 3)
	if err != nil {
		return err
	}

	merged, err := mergeReleasesJSON(ancestor, ours, theirs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, merged, 0644); err != nil {
		return err
	}
	return gitAdd(path)
}

// mergeReleasesJSON 3-way merges the releases list. An entry that
// existed in the ancestor but is missing from either ours or theirs is
// treated as deliberately deleted and dropped from the result. Versions
// added on either side are included. On a same-version collision
// between ours and theirs, theirs wins (the PR's intent).
//
// The trailing metadata fields (sourceUrl, changelogUrl, homepage) are
// taken from ours so they round-trip when otherwise stripped by the
// JSON encoder.
func mergeReleasesJSON(ancestor, ours, theirs []byte) ([]byte, error) {
	var a, o, t releasesFile
	if len(ancestor) > 0 {
		if err := json.Unmarshal(ancestor, &a); err != nil {
			return nil, fmt.Errorf("parse ancestor: %v", err)
		}
	}
	if err := json.Unmarshal(ours, &o); err != nil {
		return nil, fmt.Errorf("parse ours: %v", err)
	}
	if err := json.Unmarshal(theirs, &t); err != nil {
		return nil, fmt.Errorf("parse theirs: %v", err)
	}

	aMap := indexReleases(a.Releases)
	oMap := indexReleases(o.Releases)
	tMap := indexReleases(t.Releases)

	versions := make(map[string]struct{}, len(oMap)+len(tMap))
	for v := range oMap {
		versions[v] = struct{}{}
	}
	for v := range tMap {
		versions[v] = struct{}{}
	}

	merged := make([]releaseEntry, 0, len(versions))
	for v := range versions {
		_, inA := aMap[v]
		oe, inO := oMap[v]
		te, inT := tMap[v]

		if inA && (!inO || !inT) {
			// Deleted by one side relative to the ancestor; honor the
			// deletion. This is what makes archive PRs land cleanly.
			continue
		}
		if inT {
			merged = append(merged, te)
		} else {
			merged = append(merged, oe)
		}
	}
	sort.Slice(merged, func(i, j int) bool {
		return semverLess(merged[i].Version, merged[j].Version)
	})

	out := releasesFile{
		Releases:     merged,
		SourceUrl:    o.SourceUrl,
		ChangelogUrl: o.ChangelogUrl,
		Homepage:     o.Homepage,
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(out); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func indexReleases(rs []releaseEntry) map[string]releaseEntry {
	m := make(map[string]releaseEntry, len(rs))
	for _, r := range rs {
		m[r.Version] = r
	}
	return m
}

// resolveKustomization 3-way-merges the ancestor, ours, and theirs
// copies of kustomization.yaml and writes the result.
func resolveKustomization(path string) error {
	ancestor, err := readStageOptional(path, 1)
	if err != nil {
		return err
	}
	ours, err := readStage(path, 2)
	if err != nil {
		return err
	}
	theirs, err := readStage(path, 3)
	if err != nil {
		return err
	}

	merged, err := mergeKustomization(ancestor, ours, theirs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, merged, 0644); err != nil {
		return err
	}
	return gitAdd(path)
}

// mergeKustomization rewrites only the top-level `resources:` list block
// of "ours", replacing it with the 3-way-merged result. Every byte
// outside that block is preserved verbatim, so comments, formatting,
// and the column-zero list style are retained.
func mergeKustomization(ancestor, ours, theirs []byte) ([]byte, error) {
	var ancestorResources []string
	if len(ancestor) > 0 {
		r, err := extractResources(ancestor)
		if err != nil {
			return nil, fmt.Errorf("parse ancestor: %v", err)
		}
		ancestorResources = r
	}
	oursResources, err := extractResources(ours)
	if err != nil {
		return nil, fmt.Errorf("parse ours: %v", err)
	}
	theirsResources, err := extractResources(theirs)
	if err != nil {
		return nil, fmt.Errorf("parse theirs: %v", err)
	}

	merged := merge3Strings(ancestorResources, oursResources, theirsResources)
	return rewriteResourcesBlock(ours, merged)
}

// merge3Strings 3-way merges three sets of strings. Items present in
// the ancestor but missing from either ours or theirs are dropped
// (honoring the deletion). Items added on either side are included.
// The result is sorted by semver.
func merge3Strings(ancestor, ours, theirs []string) []string {
	aSet := toSet(ancestor)
	oSet := toSet(ours)
	tSet := toSet(theirs)

	union := make(map[string]struct{}, len(oSet)+len(tSet))
	for s := range oSet {
		union[s] = struct{}{}
	}
	for s := range tSet {
		union[s] = struct{}{}
	}

	out := make([]string, 0, len(union))
	for s := range union {
		if _, inA := aSet[s]; inA {
			if _, inO := oSet[s]; !inO {
				continue
			}
			if _, inT := tSet[s]; !inT {
				continue
			}
		}
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool {
		return semverLess(out[i], out[j])
	})
	return out
}

func toSet(items []string) map[string]struct{} {
	m := make(map[string]struct{}, len(items))
	for _, s := range items {
		m[s] = struct{}{}
	}
	return m
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
