package breakingchanges

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Etcd version shipped with kubeadm is not set in release.yaml — kubeadm bakes
// it in via the `SupportedEtcdVersion` map in cmd/kubeadm/app/constants/constants.go.
// We resolve the bundled etcd version for a given Kubernetes tag and, when it
// changes across the upgrade, pull the relevant etcd CHANGELOG section.

// resolveEtcdVersionForK8s returns the etcd version kubeadm bundles for the
// given Kubernetes version (e.g. "1.34.1" -> "3.6.4"). The trailing image tag
// suffix (e.g. "-0") is stripped. Lookups are cached.
func (d *Detector) resolveEtcdVersionForK8s(ctx context.Context, k8sVersion string) (string, error) {
	k8sVersion = strings.TrimPrefix(k8sVersion, "v")
	if cached, ok := d.etcdVersionCache[k8sVersion]; ok {
		return cached, nil
	}

	minor := parseK8sMinorVersion(k8sVersion)
	if minor == 0 {
		return "", fmt.Errorf("could not parse minor version from %q", k8sVersion)
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/v%s/cmd/kubeadm/app/constants/constants.go", k8sVersion)
	content, err := d.fetchURL(ctx, url)
	if err != nil {
		return "", fmt.Errorf("fetch kubeadm constants for k8s v%s: %w", k8sVersion, err)
	}

	etcdVersion, err := parseSupportedEtcdVersion(content, minor)
	if err != nil {
		return "", err
	}

	d.etcdVersionCache[k8sVersion] = etcdVersion
	return etcdVersion, nil
}

// parseSupportedEtcdVersion extracts the etcd version for a given Kubernetes
// minor from the text of kubeadm's constants.go. Handles the standard map
// layout: `SupportedEtcdVersion = map[uint8]string{ 34: "3.6.4-0", ... }`.
func parseSupportedEtcdVersion(constantsGo string, k8sMinor int) (string, error) {
	blockRe := regexp.MustCompile(`(?s)SupportedEtcdVersion\s*=\s*map\[uint8\]string\s*\{(.*?)\}`)
	blockMatch := blockRe.FindStringSubmatch(constantsGo)
	if len(blockMatch) < 2 {
		return "", fmt.Errorf("SupportedEtcdVersion map not found in constants.go")
	}

	entryRe := regexp.MustCompile(`(?m)^\s*(\d+)\s*:\s*"([^"]+)"`)
	for _, match := range entryRe.FindAllStringSubmatch(blockMatch[1], -1) {
		minor, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		if minor == k8sMinor {
			return stripEtcdImageSuffix(match[2]), nil
		}
	}

	return "", fmt.Errorf("no etcd version listed for Kubernetes minor %d", k8sMinor)
}

// stripEtcdImageSuffix removes the image-tag rev suffix kubeadm appends (e.g.
// "3.6.4-0" -> "3.6.4"). Etcd CHANGELOGs use the unsuffixed form.
func stripEtcdImageSuffix(v string) string {
	if idx := strings.Index(v, "-"); idx > 0 {
		return v[:idx]
	}
	return v
}

// fetchEtcdChangelog returns a filtered slice of the etcd CHANGELOG covering
// the versions between fromVer (exclusive) and toVer (inclusive). If the minor
// changes (e.g. 3.5 -> 3.6), the returned text is prefixed with a strong
// warning — etcd minor bumps typically include storage-format or API changes.
func (d *Detector) fetchEtcdChangelog(ctx context.Context, fromVer, toVer string) (string, error) {
	fromVer = strings.TrimPrefix(fromVer, "v")
	toVer = strings.TrimPrefix(toVer, "v")
	cacheKey := fmt.Sprintf("%s->%s", fromVer, toVer)
	if cached, ok := d.etcdCache[cacheKey]; ok {
		return cached, nil
	}

	fromMinor, err := parseEtcdMinor(fromVer)
	if err != nil {
		return "", err
	}
	toMinor, err := parseEtcdMinor(toVer)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# etcd bundled by kubeadm: v%s -> v%s\n\n", fromVer, toVer))

	if fromMinor != toMinor {
		sb.WriteString("## ⚠️ ETCD MINOR-VERSION BUMP ⚠️\n\n")
		sb.WriteString(fmt.Sprintf("kubeadm switches the bundled etcd from the %s line to the %s line across this Kubernetes upgrade.\n", fromMinor, toMinor))
		sb.WriteString("Minor etcd bumps are high-risk: storage format, WAL, and default behaviour changes are routine at this boundary. Treat this as a release-blocker concern until verified in a test cluster.\n\n")
	}

	// Pull the target minor's CHANGELOG — that's where the new entries live.
	targetURL := fmt.Sprintf("https://raw.githubusercontent.com/etcd-io/etcd/main/CHANGELOG/CHANGELOG-%s.md", toMinor)
	targetBody, err := d.fetchURL(ctx, targetURL)
	if err != nil {
		return "", fmt.Errorf("fetch etcd CHANGELOG-%s.md: %w", toMinor, err)
	}

	// If crossing minors, include the full CHANGELOG-<toMinor>.md up to toVer.
	// Otherwise, only include entries strictly between fromVer and toVer.
	var filtered string
	if fromMinor != toMinor {
		filtered = extractEtcdUpTo(targetBody, toVer)
	} else {
		filtered = extractEtcdBetween(targetBody, fromVer, toVer)
	}
	if strings.TrimSpace(filtered) == "" {
		filtered = fmt.Sprintf("No CHANGELOG entries found between v%s and v%s. Review the full changelog manually: %s\n", fromVer, toVer, targetURL)
	}
	sb.WriteString(filtered)

	result := sb.String()
	d.etcdCache[cacheKey] = result
	return result, nil
}

// parseEtcdMinor extracts the "MAJOR.MINOR" prefix of an etcd version.
// e.g. "3.6.4" -> "3.6".
func parseEtcdMinor(version string) (string, error) {
	re := regexp.MustCompile(`^(\d+\.\d+)`)
	match := re.FindStringSubmatch(version)
	if len(match) < 2 {
		return "", fmt.Errorf("could not parse etcd minor from %q", version)
	}
	return match[1], nil
}

// etcdReleaseHeaderRe matches etcd CHANGELOG patch release headers, e.g.
// "## v3.6.4 (2025-07-31)" or "## v3.6.11 (TBC)".
var etcdReleaseHeaderRe = regexp.MustCompile(`(?m)^##\s+v(\d+\.\d+\.\d+)\b`)

// extractEtcdBetween returns CHANGELOG sections for releases strictly greater
// than fromVer and up to (and including) toVer. Assumes standard etcd format
// where releases are listed newest-first.
func extractEtcdBetween(changelog, fromVer, toVer string) string {
	sections := splitEtcdSections(changelog)
	var sb strings.Builder
	for _, sec := range sections {
		if compareSemver(sec.version, fromVer) > 0 && compareSemver(sec.version, toVer) <= 0 {
			sb.WriteString(sec.body)
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// extractEtcdUpTo returns CHANGELOG sections for releases up to (and including)
// toVer. Used when the minor differs from the source so we include every entry
// in the target minor up to the bundled patch.
func extractEtcdUpTo(changelog, toVer string) string {
	sections := splitEtcdSections(changelog)
	var sb strings.Builder
	for _, sec := range sections {
		if compareSemver(sec.version, toVer) <= 0 {
			sb.WriteString(sec.body)
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

type etcdSection struct {
	version string
	body    string
}

// splitEtcdSections breaks an etcd CHANGELOG into per-release sections.
func splitEtcdSections(changelog string) []etcdSection {
	indexes := etcdReleaseHeaderRe.FindAllStringSubmatchIndex(changelog, -1)
	if len(indexes) == 0 {
		return nil
	}

	var sections []etcdSection
	for i, idx := range indexes {
		start := idx[0]
		end := len(changelog)
		if i+1 < len(indexes) {
			end = indexes[i+1][0]
		}
		version := changelog[idx[2]:idx[3]]
		sections = append(sections, etcdSection{
			version: version,
			body:    strings.TrimSpace(changelog[start:end]),
		})
	}
	return sections
}

// compareSemver compares two dotted versions (e.g. "3.6.4", "3.5.21").
// Returns negative if a < b, zero if equal, positive if a > b.
func compareSemver(a, b string) int {
	aParts := strings.Split(a, ".")
	bParts := strings.Split(b, ".")
	n := len(aParts)
	if len(bParts) > n {
		n = len(bParts)
	}
	for i := 0; i < n; i++ {
		var ai, bi int
		if i < len(aParts) {
			ai, _ = strconv.Atoi(aParts[i])
		}
		if i < len(bParts) {
			bi, _ = strconv.Atoi(bParts[i])
		}
		if ai != bi {
			return ai - bi
		}
	}
	return 0
}
