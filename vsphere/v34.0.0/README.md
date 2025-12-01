# :zap: Giant Swarm Release v34.0.0 for vSphere :zap:

<< Add description here >>

## Changes compared to v33.0.1

### Components

- cluster-vsphere from v3.2.0 to v4.0.0
- Flatcar from v4230.2.3 to [v4459.2.1](https://www.flatcar.org/releases/#release-4459.2.1)
- Kubernetes from v1.33.5 to [v1.34.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.2)
- os-tooling from v1.26.1 to v1.26.2

### cluster-vsphere [v3.2.0...v4.0.0](https://github.com/giantswarm/cluster-vsphere/compare/v3.2.0...v4.0.0)

#### Added

- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- Add `"helm.sh/resource-policy": keep` annotation to `VSphereCluster` CR so that it doesn't get removed by Helm when uninstalling this chart. The CAPI controllers will take care of removing it, following the expected deletion order.
- Add `"helm.sh/resource-policy": keep` annotation to the provider secret. This is to ensure that it isn't removed by Helm, thus leading to a race condition when deleting the cluster as the vSphere cleaner needs it to clean up resources in vSphere.

#### Changed

- Chart: Update `cluster` to v5.0.0.
- Chart: Update `cluster` to v4.6.0.
- Chart: Update `cluster` to v4.5.1.
- Chart: Update `cluster` to v4.5.0.
- Chart: Update `cluster` to v4.4.0.
- Chart: Update `cluster` to v4.3.0.

### Apps

- cert-exporter from v2.9.12 to v2.9.14
- cert-manager from v3.9.3 to v3.9.4
- cilium from v1.3.1 to v1.3.2
- etcd-defrag from v1.2.1 to v1.2.3
- etcd-k8s-res-count-exporter from v1.10.9 to v1.10.11
- k8s-audit-metrics from v0.10.8 to v0.10.10
- node-exporter from v1.20.7 to v1.20.9
- observability-policies from v0.0.2 to v0.0.3
- Added priority-classes v0.3.0
- security-bundle from v1.13.1 to v1.15.0

### cert-exporter [v2.9.12...v2.9.14](https://github.com/giantswarm/cert-exporter/compare/v2.9.12...v2.9.14)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### cert-manager [v3.9.3...v3.9.4](https://github.com/giantswarm/cert-manager-app/compare/v3.9.3...v3.9.4)

#### Added

- Add E2E tests using apptest-framework for automated PR testing across multiple providers (CAPA, CAPV, CAPZ, CAPVCD).
  - **Basic test suite**: Validates fresh installations
  - **Upgrade test suite**: Tests upgrade scenarios and certificate reconciliation
- Add certificate issuance integration test to cluster-test-suites.

#### Changed

- Upgrade cert-manager to v1.18.2.

### cilium [v1.3.1...v1.3.2](https://github.com/giantswarm/cilium-app/compare/v1.3.1...v1.3.2)

#### Changed

- Upgrade Cilium to [v1.18.4](https://github.com/cilium/cilium/releases/tag/v1.18.4).

### etcd-defrag [v1.2.1...v1.2.3](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.1...v1.2.3)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.36.0. ([#69](https://github.com/giantswarm/etcd-defrag-app/pull/69))
- Chart: Update dependency ahrtr/etcd-defrag to v0.35.0. ([#64](https://github.com/giantswarm/etcd-defrag-app/pull/64))

### etcd-k8s-res-count-exporter [v1.10.9...v1.10.11](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.9...v1.10.11)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.8...v0.10.10](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.8...v0.10.10)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### node-exporter [v1.20.7...v1.20.9](https://github.com/giantswarm/node-exporter-app/compare/v1.20.7...v1.20.9)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### observability-policies [v0.0.2...v0.0.3](https://github.com/giantswarm/observability-policies-app/compare/v0.0.2...v0.0.3)

#### Fixed

- Missing RBAC for kyverno-report-controller

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).

### security-bundle [v1.13.1...v1.15.0](https://github.com/giantswarm/security-bundle/compare/v1.13.1...v1.15.0)

#### Added

- Add `kubescape` (app) version v0.0.4.

#### Changed

- Update `kyverno` (app) to v0.21.1.
- Update `kyverno-crds` (app) to v1.15.0.
- Update `kyverno` (app) to v0.20.1.
- Update `kyverno-crds` (app) to v1.14.0.
- Update `kyverno-policies` (app) to v0.24.0.
- Update `reports-server` (app) to v0.0.3.
