# :zap: Giant Swarm Release v36.0.0 for  :zap:

## Changes compared to v35.1.0

### Components

- Kubernetes from v1.35.7 to [v1.36.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.4)

### Apps

- node-exporter from v1.20.11 to v1.20.13
- observability-bundle from v3.3.0 to v3.3.1
- security-bundle from v2.2.0 to v2.3.0
- teleport-kube-agent from v0.11.1 to v0.12.0

### node-exporter [v1.20.11...v1.20.13](https://github.com/giantswarm/node-exporter-app/compare/v1.20.11...v1.20.13)

#### Changed

- CircleCI: Do not override app version.
- Chart: Move PolicyException to `kube-system` namespace.

### observability-bundle [v3.3.0...v3.3.1](https://github.com/giantswarm/observability-bundle/compare/v3.3.0...v3.3.1)

#### Changed

- Values: Generate schema for Alloy PodLogs CRDs.
- Values: Add Cilium as dependency for Alloy apps & Kube Prometheus Stack.
- Values: Update Alloy apps to v0.21.2.

#### Removed

- Values: Remove unused catalog.

### security-bundle [v2.2.0...v2.3.0](https://github.com/giantswarm/security-bundle/compare/v2.2.0...v2.3.0)

#### Changed

- Update `kyverno-policy-operator` (app) to v0.2.3.
- Update `policy-api` (app) to v0.0.9.
- Update `starboard-exporter` (app) to v1.2.3.
- Update `trivy` (app) to v0.17.0.
- Update `trivy-operator` (app) to v0.13.3.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Values: Make Kyverno and Kyverno Policy Operator depend on Cilium.

### teleport-kube-agent [v0.11.1...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.11.1...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
