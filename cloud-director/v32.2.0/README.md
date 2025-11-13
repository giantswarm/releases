# :zap: Giant Swarm Release v32.2.0 for VMware Cloud Director :zap:

<< Add description here >>

## Changes compared to v32.1.0

### Components

- Added cluster-aws v6.2.0
- cluster-cloud-director from v1.0.0 to v2.4.0
- Flatcar from v4230.2.4 to [v4459.2.0](https://www.flatcar.org/releases/#release-4459.2.0)
- Kubernetes from v1.32.9 to [v1.32.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.10)

### cluster-aws [v6.2.0](https://github.com/giantswarm/cluster-aws/releases/tag/v6.2.0)

#### Changed

- Change default consolidation time to 6 hours to avoid constant node rolling.
- Rename `capa-karpenter-taint-remover` app.

### cluster-cloud-director [v1.0.0...v2.4.0](https://github.com/giantswarm/cluster-cloud-director/compare/v1.0.0...v2.4.0)

#### Changed

- Chart: Update `cluster` to v4.4.0.
- Chart: Update `cluster` to v4.3.0.
- Chart: Update `cluster` to v4.2.0.
- Chart: Update `cluster` to v4.1.0.
- Chart: Update `cluster` to v4.0.3.
- Chart: Update `cluster` to v4.0.2.
- Chart: Update `cluster` to v4.0.1.

### Apps

- etcd-defrag from v1.2.2 to v1.2.3
- security-bundle from v1.14.0 to v1.15.0

### etcd-defrag [v1.2.2...v1.2.3](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.2...v1.2.3)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.36.0. ([#69](https://github.com/giantswarm/etcd-defrag-app/pull/69))

### security-bundle [v1.14.0...v1.15.0](https://github.com/giantswarm/security-bundle/compare/v1.14.0...v1.15.0)

#### Added

- Add `kubescape` (app) version v0.0.4.

#### Changed

- Update `kyverno` (app) to v0.21.1.
- Update `kyverno-crds` (app) to v1.15.0.
