# :zap: Giant Swarm Release v14.2.1 for AWS :zap:

This release fixes the issue which prevented clusters being created on some management clusters with the AWS v14.2.0 release.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

## Change details

- Upgraded cert-operator to v1.0.1
- Upgraded aws-operator to v10.3.1
