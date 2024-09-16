# :zap: Giant Swarm Release v29.1.0 for Azure :zap:

## Changes compared to v29.0.0

### Apps

- coredns from v1.21.0 to v1.22.0
- cert-exporter from v2.9.1 to v2.9.2
- observability-bundle from v1.5.3 to v1.6.1

### coredns [1.22.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.22.0)

#### Changed

- Update `coredns` image to [1.11.3](https://github.com/coredns/coredns/releases/tag/v1.11.3).

#### Removed

- Removed legacy Giant Swarm monitoring labels as coredns is monitored through a prometheus-operator generated servicemonitor.

### cert-exporter [2.9.2](https://github.com/giantswarm/cert-exporter/releases/tag/v2.9.2)

#### Added

- Chart: Add VPA and resources configuration for deployment and daemonset. ([#382](https://github.com/giantswarm/cert-exporter/pull/382))

### observability-bundle [1.6.1](https://github.com/giantswarm/observability-bundle/releases/tag/v1.6.1)

#### Fixed
- Disable usage reporting to GrafanaLabs by:
  - Bumping `alloyLogs` and `alloyMetrics` to v0.4.1.
  - Bumping `grafanaAgent` to v0.4.6.
