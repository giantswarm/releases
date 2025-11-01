# :zap: Giant Swarm Release v33.1.0 for Azure :zap:

<< Add description here >>

## Changes compared to v33.0.1

### Components

- cluster-azure from v4.2.0 to v4.4.0
- Flatcar from v4230.2.3 to [v4230.2.4](https://www.flatcar.org/releases/#release-4230.2.4)
- os-tooling from v1.26.1 to v1.26.2

### cluster-azure [v4.2.0...v4.4.0](https://github.com/giantswarm/cluster-azure/compare/v4.2.0...v4.4.0)

#### Changed

- Make `global.baseDomain` and `global.managementCluster` required values. These values will be passed to the chart when deploying it from the `cluster-app-installation-values` ConfigMap in the default namespace.
- Extract required values to its own central file to avoid repeating the `required` keyword and error messages. This is normally done automatically by a Kyverno policy.
- Chart: Update `cluster` to v4.4.0.
- Install External DNS CRDs and watch default sources.
- Chart: Update `cluster` to v4.3.0.

### Apps

- cert-exporter from v2.9.12 to v2.9.13
- cert-manager from v3.9.3 to v3.9.4
- etcd-defrag from v1.2.1 to v1.2.2
- etcd-k8s-res-count-exporter from v1.10.9 to v1.10.10
- k8s-audit-metrics from v0.10.8 to v0.10.9
- node-exporter from v1.20.7 to v1.20.8
- security-bundle from v1.13.1 to v1.14.0

### cert-exporter [v2.9.12...v2.9.13](https://github.com/giantswarm/cert-exporter/compare/v2.9.12...v2.9.13)

#### Changed

- Go: Update dependencies.

### cert-manager [v3.9.3...v3.9.4](https://github.com/giantswarm/cert-manager-app/compare/v3.9.3...v3.9.4)

#### Added

- Add E2E tests using apptest-framework for automated PR testing across multiple providers (CAPA, CAPV, CAPZ, CAPVCD).
  - **Basic test suite**: Validates fresh installations
  - **Upgrade test suite**: Tests upgrade scenarios and certificate reconciliation
- Add certificate issuance integration test to cluster-test-suites.

#### Changed

- Upgrade cert-manager to v1.18.2.

### etcd-defrag [v1.2.1...v1.2.2](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.1...v1.2.2)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.35.0. ([#64](https://github.com/giantswarm/etcd-defrag-app/pull/64))

### etcd-k8s-res-count-exporter [v1.10.9...v1.10.10](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.9...v1.10.10)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.8...v0.10.9](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.8...v0.10.9)

#### Changed

- Go: Update dependencies.

### node-exporter [v1.20.7...v1.20.8](https://github.com/giantswarm/node-exporter-app/compare/v1.20.7...v1.20.8)

#### Changed

- Go: Update dependencies.

### security-bundle [v1.13.1...v1.14.0](https://github.com/giantswarm/security-bundle/compare/v1.13.1...v1.14.0)

#### Changed

- Update `kyverno` (app) to v0.20.1.
- Update `kyverno-crds` (app) to v1.14.0.
- Update `kyverno-policies` (app) to v0.24.0.
- Update `reports-server` (app) to v0.0.3.
