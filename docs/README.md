# Release Workflows Documentation

This directory contains documentation for the automated GitHub workflows that manage releases in the giantswarm/releases repository.

## Documentation Structure

### [Creating and Updating Releases](workflows-creating-and-updating-releases.md)
Documentation for workflows that create and maintain releases:
- **Create Release PR** (`create-release.yaml`) - Creates new major, minor, and patch releases
- **Update Release PR** (`update-release.yaml`) - Handles slash commands for updating release PRs
- **Bump Open Releases** (`bump-open-releases.yaml`) - Weekly updates to all open release PRs

These workflows automate the release drafting process and provide commands for updating components and apps.

### [Deprecating and Archiving Releases](workflows-deprecating-and-archiving-releases.md)
Documentation for workflows that manage the end-of-life process for releases:
- **Deprecate Releases** (`deprecate-releases.yaml`) - Marks outdated releases as deprecated
- **Archive Releases** (`archive-releases.yaml`) - Moves deprecated releases to archived directories

These workflows ensure safe transitions by validating against active usage metrics from Grafana before making changes.

### [Pinning Component Versions](workflows-pinning-versions.md)
Documentation for pinning specific component or app versions in releases:
- **Pin Version** (`pin-version.yaml`) - Pins component/app versions via PR comments
- Prevents automatic version bumps for specified duration
- Updates both current release and `requests.yaml` for future releases

Use this when you need to keep a specific version temporarily or long-term.

## Quick Reference

### Workflow Triggers

| Workflow | Schedule | Manual Trigger |
|----------|----------|----------------|
| Create Release PR | 1st of month 6:00 AM UTC | ✅ |
| Update Release PR | On PR comment | ❌ |
| Bump Open Releases | Monday 9:00 AM UTC | ✅ |
| Deprecate Releases | Monday 6:00 AM UTC | ✅ |
| Archive Releases | Monday 10:00 AM UTC | ✅ |

## External Resources

- [CAPI Release Drafting Guide](https://intranet.giantswarm.io/docs/product/releases/capi/capi-release-drafting/) - Complete guide for drafting CAPI releases
- [CAPI Releases Dashboard](https://giantswarm.grafana.net/d/be9a0bh8mbwn4e/capi-releases) - Live metrics for release usage
