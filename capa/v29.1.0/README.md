# :zap: Giant Swarm Release v29.1.0 for CAPA :zap:

## Changes compared to v29.0.0

### Components

- Flatcar from v3815.2.5 to [v3975.2.0](https://www.flatcar.org/releases#release-3975.2.0)
- Kubernetes from v1.29.7 to [v1.29.8](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md#changelog-since-v1297)

### Apps

- cert-exporter from v2.9.1 to v2.9.2
- node-exporter from v1.19.0 to v1.20.0
- observability-bundle from v1.5.2 to v1.6.2
- security-bundle from v1.8.0 to v1.8.1

### cert-exporter [v2.9.1...v2.9.2](https://github.com/giantswarm/cert-exporter/compare/v2.9.1...v2.9.2)

#### Added

- Chart: Add VPA and resources configuration for deployment and daemonset. ([#382](https://github.com/giantswarm/cert-exporter/pull/382))

### node-exporter [v1.19.0...v1.20.0](https://github.com/giantswarm/node-exporter-app/compare/v1.19.0...v1.20.0)

#### Changed

- Synced with upstream chart v4.38.0 (node-exporter 1.8.2).

### observability-bundle [v1.5.2...v1.6.2](https://github.com/giantswarm/observability-bundle/compare/v1.5.2...v1.6.2)

#### Added

- Add `alloy` v0.4.0 as `alloyMetrics`.

#### Changed

- Disable usage reporting to GrafanaLabs by:
  - Bumping `alloyLogs` and `alloyMetrics` to v0.4.1.
  - Bumping `grafanaAgent` to v0.4.6.
- Bump `alloyLogs` to v0.4.0.
- Rename `alloy-logs` app to camel case `alloyLogs`.

### security-bundle [v1.8.0...v1.8.1](https://github.com/giantswarm/security-bundle/compare/v1.8.0...v1.8.1)

#### Changed

- Update `trivy-operator` (app) to v0.9.1.
