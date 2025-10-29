# :zap: Giant Swarm Release v33.1.0 for Azure :zap:

<< Add description here >>

## Changes compared to v33.0.1

### Components

- cluster-azure from v4.2.0 to v4.4.0
- Flatcar from v4230.2.3 to [v4230.2.4](https://www.flatcar-linux.org/releases/#release-4230.2.4)
- os-tooling from v1.26.1 to v1.26.2

### cluster-azure [v4.2.0...v4.4.0](https://github.com/giantswarm/cluster-azure/compare/v4.2.0...v4.4.0)

#### Changed

- Make `global.baseDomain` and `global.managementCluster` required values. These values will be passed to the chart when deploying it from the `cluster-app-installation-values` ConfigMap in the default namespace.
- Extract required values to its own central file to avoid repeating the `required` keyword and error messages. This is normally done automatically by a Kyverno policy.
- Chart: Update `cluster` to v4.4.0.
- Install External DNS CRDs and watch default sources.
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
