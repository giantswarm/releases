# :zap: Giant Swarm Release v34.1.0 for Azure :zap:

<< Add description here >>

## Changes compared to v34.0.0

### Components

- Flatcar from v4459.2.2 to [v4459.2.3](https://www.flatcar.org/releases/#release-4459.2.3)
- os-tooling from v1.26.3 to v1.26.4

### Apps

- etcd-defrag from v1.2.3 to v1.2.4
- observability-bundle from v2.5.0 to v2.6.0
- security-bundle from v1.16.1 to v1.17.0

### etcd-defrag [v1.2.3...v1.2.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.3...v1.2.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.37.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/78))

### observability-bundle [v2.5.0...v2.6.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.6.0)

#### Added

- Add KSM metrics for Gateway API resources

### security-bundle [v1.16.1...v1.17.0](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v1.17.0)

#### Changed

- Update `kyverno` (app) to v0.23.0.
- Update `kyverno-crds` (app) to v1.16.0.
- Update `reports-server` (app) to v0.1.0.
- Update `cloudnative-pg` (app) to v0.0.13.
- Update `kubescape` (app) to v0.0.5.
- Update `starboard-exporter` (app) to v1.0.2.
