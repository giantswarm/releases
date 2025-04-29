# :zap: Giant Swarm Release v29.0.0 for vSphere :zap:

## Changes compared to v28.0.1

### Components

- cluster-vsphere from v0.65.1 to v0.65.2.
- Flatcar from v3815.2.5 to [v3975.2.2](https://www.flatcar.org/releases#release-3975.2.2)
- Kubernetes from v1.28.15 to [v1.29.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md)

### cluster-vsphere [v0.65.1...v0.65.2](https://github.com/giantswarm/cluster-vsphere/compare/v0.65.1...v0.65.2)

#### Changed

- Fix `kube-vip` static pod manifest for Kubernetes `1.29` onwards.

### Apps

- cert-exporter from v2.9.1 to v2.9.2
- coredns from v1.21.0 to v1.22.0
- node-exporter from v1.19.0 to v1.20.0
- observability-bundle from v1.5.3 to v1.6.2
- security-bundle from v1.8.0 to v1.8.2
- teleport-kube-agent from v0.9.2 to v0.10.3
- vertical-pod-autoscaler from v5.2.4 to v5.3.0
- vertical-pod-autoscaler-crd from v3.1.0 to v3.1.1

### cert-exporter [v2.9.1...v2.9.2](https://github.com/giantswarm/cert-exporter/compare/v2.9.1...v2.9.2)

#### Added

- Chart: Add VPA and resources configuration for deployment and daemonset. ([#382](https://github.com/giantswarm/cert-exporter/pull/382))

### coredns [v1.21.0...v1.22.0](https://github.com/giantswarm/coredns-app/compare/v1.21.0...v1.22.0)

#### Changed

- Update `coredns` image to [1.11.3](https://github.com/coredns/coredns/releases/tag/v1.11.3).

#### Removed

- Removed legacy Giant Swarm monitoring labels as coredns is monitored through a prometheus-operator generated servicemonitor.

### node-exporter [v1.19.0...v1.20.0](https://github.com/giantswarm/node-exporter-app/compare/v1.19.0...v1.20.0)

#### Changed

- Synced with upstream chart v4.38.0 (node-exporter 1.8.2).

### observability-bundle [v1.5.3...v1.6.2](https://github.com/giantswarm/observability-bundle/compare/v1.5.3...v1.6.2)

#### Added

- Add `alloy` v0.4.0 as `alloyMetrics`.

#### Changed

- Fixed `alloyMetrics` catalog
- Disable usage reporting to GrafanaLabs by:
  - Bumping `alloyLogs` and `alloyMetrics` to v0.4.1.
  - Bumping `grafanaAgent` to v0.4.6.

### security-bundle [v1.8.0...v1.8.2](https://github.com/giantswarm/security-bundle/compare/v1.8.0...v1.8.2)

#### Changed

- Update `cloudnative-pg` (app) to v0.0.6.
- Update `trivy-operator` (app) to v0.10.0.
- Update `kyverno-policy-operator` (app) to v0.0.8.
- Update `kyverno` (app) to v0.17.16.

### teleport-kube-agent [v0.9.2...v0.10.3](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.9.2...v0.10.3)

#### Changed

- Disable JAMF components on chart templates
- Fix issues with templates
- Change ownership to Team Shield
- Added small fix on `podSecurityContext` for `seccompProfile`.
- Upgraded to Teleport `version 16`

### vertical-pod-autoscaler [v5.2.4...v5.3.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.2.4...v5.3.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.9.0. ([#314](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/314))
- Chart: Consume `global.imageRegistry`. ([#315](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/315))

#### Removed

- Chart: Do not override `crds.image.tag`. ([#316](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/316))

### vertical-pod-autoscaler-crd [v3.1.0...v3.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.1.0...v3.1.1)

#### Changed

- Chart: Improve `Chart.yaml`. ([#110](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/110))
- Repository: Some chores. ([#111](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/111))
