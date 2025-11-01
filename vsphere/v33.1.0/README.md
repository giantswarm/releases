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

- cert-manager from v3.9.3 to v3.9.4
- etcd-defrag from v1.2.1 to v1.2.2
- security-bundle from v1.13.1 to v1.14.0

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

### security-bundle [v1.13.1...v1.14.0](https://github.com/giantswarm/security-bundle/compare/v1.13.1...v1.14.0)

#### Changed

- Update `kyverno` (app) to v0.20.1.
- Update `kyverno-crds` (app) to v1.14.0.
- Update `kyverno-policies` (app) to v0.24.0.
- Update `reports-server` (app) to v0.0.3.
