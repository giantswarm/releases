# :zap: Giant Swarm Release v14.2.1 for AWS :zap:

This release fixes the issue which prevented clusters being created on some management clusters with the AWS v14.2.0 release.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

## Change details

### aws-operator [10.3.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.3.1)

#### Added
- Backport China Flatcar AMIs.

### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.

#### Changed
- Update Kubernetes dependencies to 1.18 versions.
- Reconcile `CertConfig`s based on their `cert-operator.giantswarm.io/version` label.

#### Removed
- Stop using the `VersionBundle` version.

#### Added
- Add network policy resource.
- Added lookup for nodepool clusters in other namespaces than `default`.
