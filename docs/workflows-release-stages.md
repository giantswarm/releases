# Release Stages

This document describes the release stages lifecycle and the `/stage` slash command workflow (`release-stages.yaml`).

## Overview

Release PRs follow a staged lifecycle managed via PR labels. Stages coordinate handoffs between Team Tenet (core testing) and product teams (component bumps, testing, reviews).

## Stages

| Stage | Label | Who can make changes | Purpose |
|-------|-------|---------------------|---------|
| Development | `stage/development` | Everyone | Tenet is reviewing and testing core functionality. PR is not yet ready for product team input. |
| Active | `stage/active` | Everyone | Product teams can review, bump components, run tests, and add changes. |
| Freeze | `stage/freeze` | Team Tenet only | Only bugfixes and critical patches. Non-Tenet `/update-release` commands are rejected. |
| Release & Publish | _(PR merged)_ | — | Release is finalized and published. |

## Stage Transitions

Valid transitions:

```
Development → Active → Freeze → (merge)
```

Legacy PRs without a stage label can transition directly to Active:

```
(none) → Active
```

All other transitions are rejected with an error comment.

## Commands

### `/stage active`

Moves the release from Development to Active.

**What happens:**
- `stage/development` label is removed, `stage/active` is added
- Team reviewers are requested: `team-rocket`, `team-phoenix`, `team-shield`, `team-atlas`
- Slack notification is sent (consolidated releases only)
- PR body stage dashboard is updated

### `/stage freeze`

Moves the release from Active to Freeze.

**What happens:**
- `stage/active` label is removed, `stage/freeze` is added
- Slack notification is sent (consolidated releases only)
- PR body stage dashboard is updated
- Weekly automatic bump job (`bump-open-releases.yaml`) skips this PR
- Non-Tenet users are blocked from running `/update-release`

## Stage Gating

### Development (soft gate)

The `stage/development` label is informational only. It signals that Tenet is working on the release, but does not block anyone from making changes. The Slack announcement says "Development stage" instead of "ready for product teams".

### Active (no gate)

All users can make changes during the Active stage.

### Freeze (hard gate)

During Freeze:
- **`/update-release`**: Only Team Tenet members can run this command. Non-members receive a rejection comment with a `-1` reaction.
- **Weekly bump**: The `bump-open-releases.yaml` workflow excludes PRs with the `stage/freeze` label from its matrix entirely.
- **`/pin-version`**: Not gated (pinning is always allowed).

## Stage Dashboard

Release PRs include a stage dashboard in the PR body (between `<!-- STAGE_DASHBOARD_START -->` and `<!-- STAGE_DASHBOARD_END -->` markers):

```markdown
| Stage | Status |
|-------|--------|
| Development | :white_check_mark: Done |
| Active | :arrow_right: Current |
| Freeze | :hourglass: Pending |
| Release & Publish | :hourglass: Pending |
```

The dashboard is automatically updated when the stage changes.

## Slack Notifications

Slack notifications are sent **only for consolidated (monthly) releases** — individual/provider-specific releases are skipped since provider teams only watch their own PRs.

Detection: consolidated releases have branch pattern `release-vX.Y.Z`, while individual releases have `release-vX.Y.Z-provider`.

### On PR creation (Development)

> "A new workload cluster release PR has been created and is in the Development stage. Team Tenet is reviewing and testing core functionality. Product teams will be notified when the release moves to the Active stage."

### On `/stage active`

> "Release VERSION is now in the Active stage. Product teams can now review, bump components, and run tests."

### On `/stage freeze`

> "Release VERSION has entered Freeze. Only bugfixes and critical patches are allowed."

## Workflow Details

### Trigger

```yaml
on:
  issue_comment:
    types: [created]
```

Matches comments starting with `/stage` on PRs in `giantswarm/releases`.

### Feedback Pattern

Follows the same pattern as `pin-version.yaml`:
1. Eyes reaction on comment (processing started)
2. Validate transition
3. Execute changes
4. +1 reaction and summary comment (success) or -1 reaction and error comment (failure)

## Labels

Created via `gh label create`:

| Label | Color | Description |
|-------|-------|-------------|
| `stage/development` | `#BFD4F2` (light blue) | Release stage: Development (Tenet testing) |
| `stage/active` | `#0E8A16` (green) | Release stage: Active (open for product team changes) |
| `stage/freeze` | `#D93F0B` (red-orange) | Release stage: Freeze (bugfixes only) |

## Related Files

- `.github/workflows/release-stages.yaml`: Stage transition workflow
- `.github/workflows/create-release.yaml`: Adds `stage/development` label on PR creation
- `.github/workflows/update-release.yaml`: Freeze gating for `/update-release`
- `.github/workflows/bump-open-releases.yaml`: Excludes frozen PRs from weekly bump
- `.github/pr-body-templates/consolidated-release-instructions.md`: Stage command docs
- `.github/pr-body-templates/individual-release-instructions.md`: Stage command docs
