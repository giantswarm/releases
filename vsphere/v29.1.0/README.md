# :zap: Giant Swarm Release v29.1.0 for vSphere :zap:

## Changes compared to v29.0.0

### Components

- cluster-vsphere from v0.65.2 to v0.66.0

### cluster-vsphere [v0.65.2...v0.66.0](https://github.com/giantswarm/cluster-vsphere/compare/v0.65.2...v0.66.0)

#### Changed

- Use Renovate to update `kube-vip` static pod manifest.
- Updated `giantswarm/cluster` to `v1.6.0`.
- Update `kubectl` image used by IPAM job to `1.29.9`.
- Use init-container to prepare `/etc/hosts` file for `kube-vip`.

### Apps

- cert-exporter from v2.9.2 to v2.9.3
- observability-bundle from v1.6.2 to v1.8.0

### cert-exporter [v2.9.2...v2.9.3](https://github.com/giantswarm/cert-exporter/compare/v2.9.2...v2.9.3)

#### Changed

- Chart: Enable `global.podSecurityStandards.enforced`. ([#420](https://github.com/giantswarm/cert-exporter/pull/420))

### observability-bundle [v1.6.2...v1.8.0](https://github.com/giantswarm/observability-bundle/compare/v1.6.2...v1.8.0)

#### Changed

- Upgrade `prometheus-agent` from v0.6.9 to v0.7.0.
  - Adds extraArgs to be able to use nice features like wal truncation
- upgrade `kube-prometheus-stack` from 61.0.0 to 65.1.1
  - prometheus-operator CRDs from 0.73.0 to 0.75.0
  - prometheus-operator from 0.75.0 to 0.77.1
  - prometheus upgraded from 2.53.0 to 2.54.1
  - grafana from 8.2.0 to 8.5.0
  - thanos ruler upgraded from 0.35.1 to 0.36.1
  - prometheus-node-exporter upgraded from 1.8.1 to 1.8.2
- Add missing depends on annotation on alloy-metrics and alloy-logs to make sure they are deployed after the prometheus-operator-crds.
- Upgrade `alloyLogs` to v0.6.1
  - Allow passing PodLogs via helm chart values
  - Upgrade to Alloy v1.4.2 which fixes a bug with component reload/evaluation and keeping Alloy up-to-date
  - Fixes an issue with CiliumNetworkPolicy preventing Alloy to run in clustering mode
