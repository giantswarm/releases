# :zap: Giant Swarm Release v33.1.0 for vSphere :zap:

<< Add description here >>

## Changes compared to v33.0.1

### Components

- cluster-vsphere from v3.2.0 to v3.4.0
- Flatcar from v4230.2.3 to [v4230.2.4](https://www.flatcar.org/releases/#release-4230.2.4)
- os-tooling from v1.26.1 to v1.26.2

### cluster-vsphere [v3.2.0...v3.4.0](https://github.com/giantswarm/cluster-vsphere/compare/v3.2.0...v3.4.0)

#### Changed

- Chart: Update `cluster` to v4.4.0.
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
