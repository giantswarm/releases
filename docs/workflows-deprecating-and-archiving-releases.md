# Deprecating and Archiving Releases

This document describes the GitHub workflows responsible for deprecating and archiving releases in the giantswarm/releases repository.

## Overview

These workflows automate the end-of-life process for releases by:
- **Deprecating** outdated releases that are no longer actively maintained
- **Archiving** deprecated releases that are no longer in use
- Ensuring safe transitions by validating against active usage metrics

## Workflows

### 1. Deprecate Releases (`deprecate-releases.yaml`)

**Purpose**: Automatically marks outdated releases as deprecated across all CAPI providers.

#### Trigger Schedule
- **Weekly**: Every Monday at 6:00 AM UTC
- **Manual**: Via workflow dispatch
- **On Push**: To master/main branch (for testing workflow changes)

#### How It Works

1. **Build Deprecation Tool**
   - Compiles `tools/deprecate-release.go` which contains the deprecation logic
   - Requires Go 1.25+

2. **Process Each Provider**
   - Iterates through: `azure`, `capa`, `vsphere`, `cloud-director`, `eks`
   - Runs the deprecation tool with Grafana API key for usage validation
   - The tool determines which releases should be deprecated based on:
     - Whether they're currently in use (checked via Grafana metrics)
     - Whether they're the latest of supported major versions
     - Whether they're required for upgrade paths

3. **Update Configuration Files**
   - Modifies `release.yaml` files to set `state: deprecated`
   - Updates `releases.json` files to set `isDeprecated: true`
   - Only commits release-related changes (other files are reset)

4. **Create Pull Request**
   - Opens a PR titled "CAPI: Deprecate outdated releases"
   - Groups deprecated releases by provider and major version in a structured table
   - Includes link to [CAPI Releases Dashboard](https://giantswarm.grafana.net/d/be9a0bh8mbwn4e/capi-releases) for verification
   - Auto-labels with: `release-maintenance`, `automated`
   - Includes `/skip-ci` directive since changes don't affect CI tests

#### Example PR Body

```markdown
## Automated Release Deprecation for CAPI

This PR automatically marks outdated releases across all CAPI providers as deprecated.
A release is kept active if it meets any of the following criteria:

- Currently in use
- Latest of supported major versions
- Required for upgrade path

### Releases Deprecated in this PR

#### CAPZ

| Major Version | Releases Deprecated |
|---------------|---------------------|
| v29           | v29.0.0, v29.1.0    |
| v30           | v30.0.0             |
```

#### Dependencies
- **Secrets Required**:
  - `GRAFANA_API_KEY`: For checking active usage metrics
  - `GITHUB_TOKEN`: For creating PRs

---

### 2. Archive Releases (`archive-releases.yaml`)

**Purpose**: Moves deprecated releases to the `archived/` directory once they're confirmed to be no longer in use.

#### Trigger Schedule
- **Weekly**: Every Monday at 10:00 AM UTC
- **Manual**: Via workflow dispatch
- **On Push**: To master/main branch (for testing workflow changes)

#### How It Works

1. **Scan for Deprecated Releases**
   - Searches each provider directory for releases with `state: deprecated`
   - Checks if already archived (skips if in `archived/` directory)

2. **Validate Safety to Archive**
   - For each deprecated release:
     - Creates a temporary Git branch
     - Simulates moving the release to `archived/`
     - Runs `tools/check-release-archive.sh` to validate against Grafana metrics
     - Confirms the release is not actively in use
   - Only proceeds if validation passes
   - Cleans up temporary branch after validation

3. **Move Safe Releases**
   - Moves validated releases from `<provider>/<version>` to `<provider>/archived/<version>`
   - Uses `git mv` to preserve history

4. **Update References**
   - **kustomization.yaml**: Removes archived releases from the resources list
   - **README.md**: Updates GitHub links to point to new archived location
     - Example: `tree/master/azure/v29.0.0` → `tree/master/azure/archived/v29.0.0`

5. **Create Pull Request**
   - Opens a PR titled "CAPI: Archive deprecated releases"
   - Groups archived releases by provider and major version
   - Auto-labels with: `release-maintenance`, `automated`
   - Includes `/skip-ci` directive

#### Safety Mechanisms

- **Double Validation**: Releases must first be deprecated, then confirmed unused
- **Grafana Integration**: Real-time usage data prevents archiving active releases
- **Temporary Branch Testing**: Validates the move before committing
- **Selective Processing**: Only archives releases that pass all safety checks

#### Example PR Body

```markdown
## Automated Release Archiving for CAPI

This PR archives deprecated releases across all CAPI providers.
Archived releases are moved from `<provider>/<version>` to `<provider>/archived/<version>`.

### Releases Archived in this PR

#### CAPZ

| Major Version | Releases Archived |
|---------------|-------------------|
| v28           | v28.0.0, v28.1.0, v28.2.0 |
| v29           | v29.0.0           |
```

#### Dependencies
- **Secrets Required**:
  - `GRAFANA_API_KEY`: For validating releases are not in use
  - `GITHUB_TOKEN`: For creating PRs
- **Script Dependencies**:
  - `tools/check-release-archive.sh`: Validates safety of archiving

---

## Workflow Sequence

The typical end-of-life process for a release follows this pattern:

```
Active Release
    ↓
    │ (Time passes, newer releases available)
    ↓
Deprecated (by deprecate-releases.yaml)
    ↓
    │ (Confirmed no longer in use)
    ↓
Archived (by archive-releases.yaml)
```

### Timeline Example

- **Week 1**: Release v29.0.0 is the latest
- **Week 10**: Release v30.0.0 is released, v29.0.0 still active
- **Week 20**: Release v31.0.0 is released
  - `deprecate-releases.yaml` marks v29.0.0 as deprecated (no longer needed for upgrade paths)
- **Week 24**: v29.0.0 confirmed unused by all installations
  - `archive-releases.yaml` moves v29.0.0 to `archived/` directory

---

## Monitoring & Verification

### Grafana Dashboard
All release lifecycle decisions are based on the [CAPI Releases Dashboard](https://giantswarm.grafana.net/d/be9a0bh8mbwn4e/capi-releases):
- Shows which releases are actively running on customer installations
- Provides 6-hour rolling window of usage data
- Used by both deprecation and archival workflows

### Manual Verification
Before merging PRs created by these workflows:
1. Review the dashboard to confirm usage patterns
2. Check that deprecated releases have suitable upgrade paths
3. Verify archived releases show zero usage for at least 4 weeks

---

## Troubleshooting

### Release Still in Use
If a release appears in deprecation PR but shouldn't be deprecated:
- Check Grafana for recent usage spikes
- Verify it's not required for a critical upgrade path
- Close the PR and manually adjust `tools/deprecate-release.go` criteria

### Failed Archival
If `archive-releases.yaml` fails to archive a release:
- Check `tools/check-release-archive.sh` output in workflow logs
- Verify Grafana API key is valid
- Confirm release is actually deprecated in `release.yaml`

### Untracked Files After Workflow
If workflow leaves untracked files (like compiled binaries):
- This is expected and handled by "Keep only release files changes" step
- Only `release.yaml` and `releases.json` changes are committed
- Binaries are cleaned up before commit

---

## Configuration

### Provider List
Both workflows process these providers:
```bash
azure capa vsphere cloud-director eks
```

To add a new provider, update the `PROVIDERS` variable in both workflow files.

### Scheduling
- **Deprecation**: Runs first (Monday 6:00 AM UTC)
- **Archival**: Runs 4 hours later (Monday 10:00 AM UTC)
- This ensures newly deprecated releases don't get immediately archived

---

## Related Files

- `tools/deprecate-release.go`: Core deprecation logic
- `tools/check-release-archive.sh`: Safety validation for archiving
- `<provider>/releases.json`: Provider-specific release metadata
- `<provider>/<version>/release.yaml`: Individual release configuration

