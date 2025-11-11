# Creating and Updating Releases

This document describes the GitHub workflows responsible for creating new releases and updating existing release PRs in the giantswarm/releases repository.

## Overview

These workflows automate:
- **Creating** new major, minor, and patch releases
- **Updating** existing release PRs with component and app version bumps
- **Bumping** apps and components in open release PRs on a weekly schedule

## Workflows

### 1. Create Release PR (`create-release.yaml`)

**Purpose**: Creates new release PRs for CAPI workload clusters across all providers.

#### Trigger Schedule
- **Scheduled**: First day of every month at 6:00 AM UTC
- **Manual**: Via workflow dispatch with customizable options

#### Manual Trigger Options

```yaml
Inputs:
  - release_type: major | minor | patch
  - version: Specific version (e.g., v30.0.1) - overrides release_type
  - provider_all: Include all providers (default: true)
  - provider_aws: Include AWS/CAPA
  - provider_azure: Include Azure/CAPZ
  - provider_cloud_director: Include Cloud Director/CAPVCD
  - provider_eks: Include EKS
  - provider_vsphere: Include vSphere/CAPV
```

---

### Release Types

#### Major Releases (X.0.0)

**Characteristics**:
- New Kubernetes versions
- Significant platform upgrades
- Breaking changes
- Attempted monthly (every 1st of the month)
- If an unreleased major exists, a minor release is created instead

**Pre-flight Checks**:
1. **Unreleased Major Check**: Skips if there's an open (unmerged) PR for the next major version
2. **Kubernetes Version**: Verifies new upstream K8s version exists
3. **Alignment Check**: Determines if all providers are on the same major version
4. **Existing PRs**: Skips if PRs already exist for the next major version

**Flow**:
```
Dispatcher → Major Release Decider
    ↓
    ├─ Aligned? → Consolidated PR (all providers in one PR)
    └─ Misaligned? → Individual PRs (one per provider)
```

**Example Branch Names**:
- Consolidated: `release-v32.0.0`
- Individual: `release-v32.0.0-aws`, `release-v32.0.0-azure`, etc.

---

#### Minor Releases (X.Y.0)

**Characteristics**:
- Component updates
- Bug fixes
- Feature improvements
- No breaking changes
- Scheduled monthly (runs when no unreleased major exists)

**Pre-flight Checks**:
1. **Existing PRs**: Checks for open minor release PRs
2. **Alignment Check**: Determines if all providers are on the same minor version

**Flow**:
```
Dispatcher → Minor Release Decider
    ↓
    ├─ Aligned? → Consolidated PR
    └─ Misaligned? → Individual PRs
```

---

#### Patch Releases (X.Y.Z)

**Characteristics**:
- Targeted security fixes
- Critical bug fixes
- Specific component updates
- **Manual only** (not scheduled)
- **Does not automatically bump** any components

**Pre-flight Checks**:
- Only runs on manual trigger with patch type or specific version

**Flow**:
```
Manual Trigger → Patch Release Decider
    ↓
    ├─ 0 or 2+ providers? → Consolidated PR
    └─ 1 provider? → Individual PR
```

**Important**: Patch releases require manual component updates via `/update-release` commands.

---

### Workflow Architecture

#### Job Structure

```
┌─────────────┐
│ Dispatcher  │  Determines release type, providers, and flags
└──────┬──────┘
       │
       ├──────────────────┬───────────────────┬──────────────────┐
       ↓                  ↓                   ↓                  ↓
┌──────────────┐   ┌─────────────┐   ┌──────────────┐   ┌──────────┐
│Major Release │   │Minor Release│   │Patch Release │   │...more   │
│   Decider    │   │   Decider   │   │   Decider    │   │deciders  │
└──────┬───────┘   └──────┬──────┘   └──────┬───────┘   └──────────┘
       │                  │                  │
       ├──────────┬───────┼──────────┬───────┼──────────┬──────────┐
       ↓          ↓       ↓          ↓       ↓          ↓          ↓
┌─────────┐  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐
│Consol.  │  │Indiv.  │ │Consol. │ │Indiv.  │ │Consol. │ │Indiv.  │
│ Major   │  │ Major  │ │ Minor  │ │ Minor  │ │ Patch  │ │ Patch  │
└─────────┘  └────────┘ └────────┘ └────────┘ └────────┘ └────────┘
       │          │         │          │         │          │
       └──────────┴─────────┴──────────┴─────────┴──────────┘
                           ↓
                    ┌─────────────┐
                    │ Summarize   │
                    │  Releases   │
                    └─────────────┘
                           ↓
                    ┌─────────────┐
                    │  Announce   │  (Scheduled runs only)
                    │  to Slack   │
                    └─────────────┘
```

---

### Key Steps in Release Creation

#### 1. Determine Next Release Version
Uses `tools/determine-next-release.sh` to calculate:
- Next version number based on provider's current latest
- Base version for comparison
- Provider acronym for PR title

#### 2. Install DevCtl
Uses `.github/actions/install-devctl` to:
- Fetch latest `devctl` release from giantswarm/devctl
- Download platform-specific binary
- Make available for release creation

#### 3. Filter Providers
Uses `tools/filter-providers-for-release.sh` to:
- Exclude providers that already have the target version
- Skip providers where release would be redundant

#### 4. Create Release Files
Executes `devctl release create` with:
- `--base`: Previous version to compare against
- `--name`: New version to create
- `--provider`: Target provider (aws, azure, etc.)
- `--bumpall`: Automatically bump all components to latest
- `--overwrite`: Replace existing files if present
- `--output markdown`: Generate markdown change summary

**Retry Logic**: Up to 3 attempts with 15-second delays

#### 5. Update README
Runs `tools/update-readme.sh` to:
- Add new version to provider's release table
- Update links and version references

#### 6. Analyze Breaking Changes
- Checks out the PR branch
- Runs `tools/breakingchanges/cmd/detect`
- Uses Anthropic API to analyze component changes
- Posts AI-generated breaking changes report to PR

#### 7. Create Pull Request
Creates PR with:
- **Title**: `"[Provider]: Release vX.Y.Z."`
- **Body**: Instructions for updating the release
- **Labels**: Release type + provider(s)
- **Branch**: `release-vX.Y.Z[-provider]`

---

### PR Body Templates

PRs include instructions for updating releases using slash commands:

#### Consolidated Release PR
```markdown
## Updating this Release

For consolidated releases, you **must** specify which provider you are updating.

Examples:
- `/update-release --provider aws --component flatcar@4152.2.3`
- `/update-release --provider azure --app azuredisk-csi-driver@1.32.9`
- `/update-readme --provider aws "Description here"`
```

#### Individual Release PR
```markdown
## Updating this Release

Provider is auto-detected from PR files.

Examples:
- `/update-release --component flatcar@4152.2.3`
- `/update-release --app aws-ebs-csi-driver@3.0.5`
- `/update-readme "Description here"`
```

---

### 2. Update Release PR (`update-release.yaml`)

**Purpose**: Responds to slash commands in PR comments to update release files.

#### Trigger
- Issue comment containing: `/update-release`, `/update-readme`, or `/update-announcement`
- Only on pull requests
- Only in giantswarm/releases repository

#### Supported Commands

##### `/update-release`
Updates components and/or apps in the release.

**Syntax**:
```bash
/update-release [--provider PROVIDER] [OPTIONS]

Options:
  --component NAME@VERSION    Update specific component
  --app NAME@VERSION         Update specific app
  --app NAME@VERSION@@dep1#dep2  Update app with dependencies
```

**Examples**:
```bash
# Single component update
/update-release --component flatcar@4152.2.3

# Multiple updates
/update-release --component cluster-aws@4.0.2 --app karpenter-bundle@2.2.0

# With provider (consolidated PRs)
/update-release --provider azure --component cluster-azure@2.4.1

# Bump all to latest (no specific version)
/update-release
```

**Behavior**:
- **With specific versions**: Uses `--update-existing --regenerate-readme` (targeted update)
- **Without versions**: Uses `--bumpall` (updates everything to latest)
- **Preserves descriptions**: Backs up and restores README/announcement descriptions

##### `/update-readme`
Updates the README.md description (line 3).

**Syntax**:
```bash
/update-readme [--provider PROVIDER] "Description text"
```

**Examples**:
```bash
/update-readme "This release includes Kubernetes 1.32 support and security fixes."
/update-readme --provider aws "CAPA-specific changes here."
```

##### `/update-announcement`
Updates the announcement.md with a brief summary.

**Syntax**:
```bash
/update-announcement [--provider PROVIDER] "Brief summary"
```

**Examples**:
```bash
/update-announcement "Critical security fixes and stability improvements."
/update-announcement --provider vsphere "CAPV workload cluster updates."
```

---

### Update Workflow Process

#### 1. React to Comment
- Adds 👀 (eyes) reaction to show processing has started
- Prevents duplicate runs

#### 2. Parse Command
- Extracts provider from `--provider` flag OR
- Auto-detects from PR's changed files (looks for azure/, capa/, vsphere/, etc.)
- Parses devctl arguments
- Decodes descriptions (handles special characters via base64)

#### 3. Execute Update
For component/app updates:
- Backs up existing descriptions
- Runs `devctl release create` with appropriate flags
- Generates change summary table
- Restores preserved descriptions

For description updates:
- Updates README.md line 3 OR
- Updates announcement.md line 1 with proper format

#### 4. Commit Changes
- Configures Git with bot identity
- Stages all changes
- Commits with message: `"chore: Update release files via comment"`
- Pushes to PR branch

#### 5. Post Results
- Adds ✅ (thumbs up) reaction for success
- Posts detailed comment with:
  - Change tables showing version updates
  - Collapsible sections for large outputs
  - Error messages if something failed

---

### Change Summary Format

The bot posts structured comments showing what changed:

#### Small Changes (expanded)
```markdown
### User-requested changes for AWS

| APP NAME | CURRENT APP VERSION | DESIRED APP VERSION |
|----------|---------------------|---------------------|
| aws-ebs-csi-driver | 3.0.4 | **3.0.5** |
| karpenter-bundle | 2.1.0 | **2.2.0** |
```

#### Large Changes (collapsed)
```markdown
<details><summary>User-requested changes for AZURE</summary>

[... tables with 5+ rows ...]

</details>
```

---

### 3. Bump Open Releases (`bump-open-releases.yaml`)

**Purpose**: Automatically updates all open release PRs weekly with latest component versions.

#### Trigger Schedule
- **Weekly**: Every Monday at 9:00 AM UTC
- **Manual**: Via workflow dispatch (can specify single PR)

#### How It Works

1. **Find Open Release PRs**
   - Searches for PRs with branch pattern: `release-v[0-9]+\.[0-9]+\.[0-9]+`
   - Manual runs can target specific PR number

2. **Post Update Command**
   - Comments `/update-release` on each PR
   - Uses `TAYLORBOT_GITHUB_ACTION` token to trigger workflows
   - No specific components = bump all to latest

3. **Trigger Update Workflow**
   - Comment triggers `update-release.yaml` workflow
   - Each PR processes independently
   - Failures don't block other PRs (fail-fast: false)

4. **Generate Summary**
   - Reports how many PRs were updated
   - Links to individual PRs for detailed results

#### Example Summary
```markdown
## Weekly Apps and Components Bump Summary

Posted `/update-release` comment on 3 open release PR(s)

The update-release workflow will now process each PR and bump apps 
and components to their latest versions.

Check individual PR comments for detailed results.
```

---

## Scheduled Release Announcements

For scheduled (non-manual) releases that successfully create PRs:

### Slack Notification
Sent to configured Slack channel with:
- Release version
- PR title and link
- Link to CAPI Release Drafting Guide
- Metadata for tracking

**Conditions**:
- Only for scheduled runs (not manual)
- Only if PR was successfully created
- Requires `SLACK_BOT_TOKEN` and `SLACK_CHANNEL_ID` secrets

---

## Environment Variables & Secrets

### Required Secrets
- `TAYLORBOT_GITHUB_ACTION`: GitHub token with workflow permissions
- `ANTHROPIC_API_KEY`: For AI-powered breaking change analysis
- `GRAFANA_API_KEY`: For checking K8s version availability
- `SLACK_BOT_TOKEN`: For Slack notifications (optional)
- `SLACK_CHANNEL_ID`: Target Slack channel (optional)

### Environment Variables
- `OPSCTL_GITHUB_TOKEN`: Authentication for opsctl operations
- `DEVCTL_GITHUB_TOKEN`: Authentication for devctl operations
- `DEVCTL_UNSAFE_FORCE_VERSION`: Forces specific devctl version

---

## Provider Configuration

### Provider Names

The workflows use different naming conventions in different contexts:

| Directory | DevCtl Flag | Display Name |
|-----------|-------------|--------------|
| `capa/` | `aws` | CAPA |
| `azure/` | `azure` | CAPZ |
| `vsphere/` | `vsphere` | CAPV |
| `cloud-director/` | `cloud-director` | CAPVCD |
| `eks/` | `eks` | EKS |

**Translation**: The workflows automatically translate between `capa` (directory) and `aws` (devctl) as needed.

---

## Helper Scripts

### `tools/determine-next-release.sh`
Calculates next version number based on:
- Provider's current releases
- Desired release type (major/minor/patch)
- Optional version override

**Outputs**:
```bash
version=v32.0.0
base=v31.2.1
provider_acronym=CAPA
```

### `tools/filter-providers-for-release.sh`
Filters out providers that:
- Already have the target release version
- Would create duplicate releases

**Outputs**:
```bash
included_providers=aws azure vsphere
excluded_providers=eks
```

### `tools/check-new-k8s-version-exists.sh`
Verifies new Kubernetes version is available for major releases.

### `tools/update-readme.sh`
Updates README.md with new release version in the appropriate table.

---

## Troubleshooting

### Release Creation Failed

**DevCtl Errors**:
- Check `DEVCTL_GITHUB_TOKEN` has correct permissions
- Verify base version exists in provider directory
- Review devctl logs in workflow output

**Version Conflicts**:
- Check if version already exists
- Use `--overwrite` flag (already included)
- Manually delete conflicting release directory

### Update Command Not Working

**Provider Detection Failed**:
- Explicitly specify `--provider` flag
- Verify PR has files in valid provider directories
- Check PR branch name follows pattern

**Changes Not Applied**:
- Check if requested version already exists
- Review devctl output in bot comment
- Verify app/component names are correct

### Breaking Change Analysis Failed

**Common Issues**:
- `ANTHROPIC_API_KEY` invalid or expired
- Branch checkout failed
- `tools/breakingchanges/` directory issues

**Workaround**:
- Breaking change analysis is optional
- PR creation continues even if analysis fails
- Manually analyze changes in PR diff

---

## Best Practices

### Creating Releases
1. **Major releases**: Wait for quarterly schedule unless urgent
2. **Minor releases**: Use monthly schedule for consistency
3. **Patch releases**: Only for targeted fixes, explicitly specify components

### Updating Releases
1. **Batch updates**: Update multiple components in single command
2. **Test individually**: For major changes, update one provider at a time
3. **Preserve descriptions**: Always add README/announcement text early

### Reviewing PRs
1. **Check breaking changes**: Review AI analysis before merging
2. **Verify versions**: Ensure component versions are correct
3. **Test upgrades**: Validate upgrade paths work as expected

---

## Related Files

- `.github/actions/install-devctl/`: DevCtl installation action
- `tools/determine-next-release.sh`: Version calculation
- `tools/filter-providers-for-release.sh`: Provider filtering
- `tools/update-readme.sh`: README updates
- `tools/check-new-k8s-version-exists.sh`: K8s version validation
- `tools/breakingchanges/`: Breaking change detection

