# :zap: Giant Swarm Release v14.2.2 for AWS :zap:

This release fixes the issue which caused Kubernetes nodes to lose network connectivity in certain situations by reverting Flatcar Container Linux to an older version.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

## Change details

- revert Flatcar Container Linux to `v2605.12.0`.
