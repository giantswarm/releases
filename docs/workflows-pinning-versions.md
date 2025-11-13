# Pinning Component Versions in Releases

This document describes how to pin specific component or app versions in release PRs using the `/pin-version` command.

## Overview

By default, the release system automatically bumps apps and components to their latest versions. Sometimes you need to pin a specific version (not the latest) for a release, either temporarily for one release or for multiple releases.

The `/pin-version` command provides an interface for teams to pin component versions via PR comments without manually editing `requests.yaml`.

## Use Cases

- **Security/Stability**: Pin a component to a known-good version while testing a newer version
- **Incremental Updates**: Pin multiple components and update them one at a time
- **Version Compatibility**: Keep a component at a specific version due to dependencies
- **Testing**: Pin a version for this release only, then allow auto-updates in future releases

## Command Syntax

```bash
/pin-version --component NAME@VERSION [--until VERSION] [--provider PROVIDER] [--reason "DESCRIPTION"]
/pin-version --app NAME@VERSION [--until VERSION] [--provider PROVIDER] [--reason "DESCRIPTION"]
```

### Parameters

| Parameter | Required | Description | Example |
|-----------|----------|-------------|---------|
| `--component` | Yes* | Component name and version to pin | `--component flatcar@4152.2.3` |
| `--app` | Yes* | App name and version to pin | `--app aws-ebs-csi-driver@3.0.5` |
| `--until` | No | Pin until this version (default: current release only) | `--until v32.0.0` |
| `--provider` | No** | Provider to update (for consolidated PRs) | `--provider aws` |
| `--reason` | No | Human-readable reason for the pin | `--reason "Testing before general rollout"` |

\* Either `--component` or `--app` is required (not both)  
\*\* Required for consolidated release PRs, optional for single-provider PRs

### Pin Duration with `--until`

| `--until` Value | Description | Example Use Case |
|-----------------|-------------|------------------|
| Not specified | Pin only this release | One-time exception for testing |
| `vX.Y.0` (minor) | Pin until this minor version | Keep consistent through v31.1.x patches |
| `vX.0.0` (major) | Pin until this major version | Long-term pin across all v31.x.x releases |

## How It Works

### 1. Pin Updates the Release

First, the command updates the current release's `release.yaml` to use the pinned version:

```bash
/pin-version --component flatcar@4152.2.3
```

This is equivalent to running:
```bash
/update-release --component flatcar@4152.2.3
```

### 2. Pin Updates requests.yaml

Then, based on the `--until` flag, it updates the provider's `requests.yaml` file:

#### No `--until` (Default): This Release Only
Creates a version-specific constraint that applies only to this exact release:

```yaml
releases:
- name: "v31.1.2"  # Exact version match
  requests:
  - name: flatcar
    version: "4152.2.3"  # Pinned version
- name: ">= 31.0.0"  # Other releases not affected
  requests:
  - name: flatcar
    version: ">= 4100.0.0"  # Auto-updates as usual
```

**Result**: Only v31.1.2 uses the pinned version. v31.1.3, v31.2.0, and all future releases auto-update.

#### `--until v31.2.0`: Until Minor Version
Creates a constraint for all patch releases in this minor series:

```yaml
releases:
- name: ">= 31.1.0 < 31.2.0"  # All 31.1.x releases
  requests:
  - name: flatcar
    version: "4152.2.3"  # Pinned for entire 31.1.x series
- name: ">= 31.0.0 < 31.1.0"  # Earlier minors not affected
  requests:
  - name: flatcar
    version: ">= 4100.0.0"
- name: ">= 31.2.0"  # v31.2.0 and later not affected
  requests:
  - name: flatcar
    version: ">= 4200.0.0"
```

**Result**: All v31.1.x releases use the pinned version. v31.2.0 and later can use newer versions.

#### `--until v32.0.0`: Until Major Version
Creates a constraint for all releases in this major series:

```yaml
releases:
- name: ">= 31.0.0 < 32.0.0"  # All 31.x.x releases
  requests:
  - name: flatcar
    version: "4152.2.3"  # Pinned for entire 31.x series
- name: ">= 32.0.0"  # v32.0.0 and later not affected
  requests:
  - name: flatcar
    version: ">= 4200.0.0"
```

**Result**: All v31.x.x releases use the pinned version. v32.0.0 and later can use newer versions.

## Examples

### Example 1: Pin Component for This Release Only

Pin Flatcar to 4152.2.3 for this release, but allow future releases to auto-update:

```bash
/pin-version --component flatcar@4152.2.3 --reason "Testing new Flatcar before general rollout"
```

**Result**: Only v31.1.2 will use Flatcar 4152.2.3. Future releases (31.1.3, 31.2.0, etc.) will auto-update to latest.

### Example 2: Pin App for Minor Series

Pin AWS EBS CSI Driver for all 31.1.x releases:

```bash
/pin-version --app aws-ebs-csi-driver@3.0.4 --until v31.2.0 --reason "Version 3.0.5 has known issues"
```

**Result**: All 31.1.x releases will use 3.0.4, but 31.2.0 and later can use newer versions.

### Example 3: Pin Component for Major Series

Pin Kubernetes for all 31.x.x releases:

```bash
/pin-version --component kubernetes@1.31.11 --until v32.0.0 --reason "Staying on 1.31.x for this major release series"
```

**Result**: All v31.x.x releases will use Kubernetes 1.31.11. v32.0.0 can use Kubernetes 1.32.x.

### Example 4: Pin in Consolidated PR

For consolidated PRs with multiple providers, specify which provider:

```bash
/pin-version --provider aws --component flatcar@4152.2.3 --until v32.0.0
```

### Example 5: Pin Multiple Components

Run separate commands to pin multiple components:

```bash
/pin-version --component flatcar@4152.2.3 --until v32.0.0
/pin-version --component kubernetes@1.31.11 --until v32.0.0
/pin-version --app cilium@1.2.2 --reason "Testing new version"
```

## Unpinning Versions

To unpin a version and return to auto-updates:

### Option 1: Manual Edit
Edit the `requests.yaml` file to remove or update the version constraint:

```yaml
# Before (pinned)
- name: "v31.1.2"
  requests:
  - name: flatcar
    version: "4152.2.3"

# After (unpinned, delete the entire release block or update the constraint)
# Just remove the specific release entry
```

### Option 2: Use `/update-release` with New Version
Simply update to a newer version, which will override the pin in `release.yaml`:

```bash
/update-release --component flatcar@4200.0.0
```

Note: This updates the current release but doesn't remove the `requests.yaml` constraint for future releases.

### Option 3: Create Follow-up PR
The most reliable way to unpin for future releases is to create a follow-up PR that modifies `requests.yaml` directly.

## Viewing Current Pins

To see what's currently pinned:

1. **For this release**: Check the PR's `release.yaml` file
2. **For future releases**: Check the provider's `requests.yaml` file

Example:
```bash
# View current release
cat capa/v31.1.2/release.yaml

# View requests (pins for future releases)
cat capa/requests.yaml
```

## Best Practices

### 1. Always Provide a Reason
Use `--reason` to document why you're pinning a version:

```bash
/pin-version --component flatcar@4152.2.3 --until v32.0.0 --reason "CVE-2024-1234 not yet fixed in 4200.x"
```

This creates a comment in the PR that explains the decision.

### 2. Use Minimal Pin Duration
Start without `--until` (this release only) and expand only if needed:

- **No `--until`**: Best for temporary pins or testing
- **`--until vX.Y.0`**: Use when you know the pin should last through patch releases
- **`--until vX.0.0`**: Use sparingly, only for fundamental compatibility issues

### 3. Review Pins Regularly
Pins can accumulate over time. Periodically review `requests.yaml` to remove outdated constraints.

### 4. Document Breaking Changes
If pinning due to breaking changes, document it in the release notes:

```bash
/update-readme "⚠️ Flatcar pinned to 4152.2.3 due to known issues in 4200.x series"
```

### 5. Test Pin Before Expanding Duration
Test without `--until` first, then expand the duration if the pin works well:

```bash
# Step 1: Test in this release
/pin-version --component flatcar@4152.2.3

# Step 2: After validation, expand duration in next release PR
/pin-version --component flatcar@4152.2.3 --until v32.0.0
```

## Workflow Integration

### When Creating a Release

1. Release PR is created with latest versions (via devctl)
2. Review the versions in `release.yaml`
3. If any version needs pinning, use `/pin-version`
4. The bot updates both `release.yaml` and `requests.yaml`

### Weekly Bump Process

The weekly bump process (`bump-open-releases.yaml`) respects pins:

- Components/apps pinned in `requests.yaml` won't be auto-updated
- Only unpinned components will bump to latest
- This ensures your pins persist across automatic updates

## Troubleshooting

### Pin Not Applied

**Symptom**: Version didn't change after `/pin-version` command

**Solutions**:
- Check if the version is already at that level
- Ensure component/app name is spelled correctly
- Verify the version exists for that component
- Check workflow logs for errors

### Pin Affects Wrong Releases

**Symptom**: Pin applied to releases it shouldn't affect

**Solutions**:
- Review the `requests.yaml` version constraint syntax
- Check if the `--until` version is correct
- Consider using no `--until` flag for isolation to this release only

### Conflicting Constraints

**Symptom**: Multiple `requests.yaml` entries conflict

**Solutions**:
- `requests.yaml` entries are evaluated top-to-bottom
- More specific constraints should come before general ones
- Remove outdated/overlapping constraints

### Pin Ignored by Weekly Bump

**Symptom**: Pinned version got updated by automatic bump

**Solutions**:
- Verify the pin exists in `requests.yaml` (not just `release.yaml`)
- Check the version constraint matches the release version
- Ensure the constraint range is correct

## Related Commands

- `/update-release`: Update components/apps without pinning for future releases
- `/update-readme`: Update release description
- `/update-announcement`: Update announcement text

## Technical Details

### Version Constraint Syntax

The `requests.yaml` file uses semantic version constraints:

| Constraint | Meaning | Example |
|------------|---------|---------|
| `1.2.3` | Exact version | `version: "1.2.3"` |
| `>= 1.2.3` | Greater than or equal | `version: ">= 1.2.3"` |
| `>= 1.2.0 < 1.3.0` | Range (minor series) | `version: ">= 1.2.0 < 1.3.0"` |
| `>= 1.0.0 < 2.0.0` | Range (major series) | `version: ">= 1.0.0 < 2.0.0"` |

### Name Constraint Syntax

Release name constraints match against release versions:

| Constraint | Matches | Example |
|------------|---------|---------|
| `"v31.1.2"` | Exact release | Only v31.1.2 |
| `">= 31.1.0"` | Greater or equal | v31.1.0, v31.1.1, v31.2.0, v32.0.0, etc. |
| `">= 31.1.0 < 31.2.0"` | Minor series | v31.1.0, v31.1.1, v31.1.2, etc. |
| `">= 31.0.0 < 32.0.0"` | Major series | All v31.x.x releases |

### Evaluation Order

`requests.yaml` entries are evaluated top-to-bottom. The first matching entry wins:

```yaml
releases:
- name: "v31.1.2"          # Most specific - checked first
  requests:
  - name: flatcar
    version: "4152.2.3"

- name: ">= 31.1.0 < 31.2.0"  # More specific - checked second
  requests:
  - name: flatcar
    version: ">= 4152.0.0"

- name: ">= 31.0.0"       # General fallback - checked last
  requests:
  - name: flatcar
    version: ">= 4100.0.0"
```

## Future Enhancements

Potential improvements being considered:

1. **Automatic Unpin**: Add expiration dates for pins
2. **Pin Reporting**: Dashboard showing all active pins
3. **Pin Validation**: Check if pinned versions have known CVEs
4. **Bulk Pin Operations**: Pin multiple components in one command

## Feedback

This feature is new. Please provide feedback on:
- Command syntax and options
- Pin duration behavior
- Documentation clarity
- Additional features needed

Open issues in the [releases repository](https://github.com/giantswarm/releases/issues).
