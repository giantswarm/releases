# :zap: Giant Swarm Release v32.2.0 for Azure :zap:

<< Add description here >>

## Changes compared to v32.1.0

### Components

- cluster-azure from v3.0.0 to v4.4.0
- Flatcar from v4230.2.4 to [v4459.2.0](https://www.flatcar.org/releases/#release-4459.2.0)
- Kubernetes from v1.32.9 to [v1.32.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.10)

### cluster-azure [v3.0.0...v4.4.0](https://github.com/giantswarm/cluster-azure/compare/v3.0.0...v4.4.0)

#### Changed

- Make `global.baseDomain` and `global.managementCluster` required values. These values will be passed to the chart when deploying it from the `cluster-app-installation-values` ConfigMap in the default namespace.
- Extract required values to its own central file to avoid repeating the `required` keyword and error messages. This is normally done automatically by a Kyverno policy.
- Chart: Update `cluster` to v4.4.0.
- Install External DNS CRDs and watch default sources.
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
