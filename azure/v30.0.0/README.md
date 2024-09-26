# :zap: Giant Swarm Release v30.0.0 for Azure :zap:

## Changes compared to v29.0.0

### Components

- cluster-azure from v1.0.0 to v1.1.0
- Flatcar from v3815.2.5 to [v3975.2.0](https://www.flatcar.org/releases#release-3975.2.0)
- Kubernetes from v1.29.7 to [v1.30.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md)

### cluster-azure [v1.0.0...v1.1.0](https://github.com/giantswarm/cluster-azure/compare/v1.0.0...v1.1.0)

#### Changed

- Apps: Use `catalog` from Release CR.

### Apps

- azure-cloud-controller-manager from v1.29.8-gs1 to v1.30.6-gs1
- azure-cloud-node-manager from v1.29.8-gs1 to v1.30.6-gs1
- cert-exporter from v2.9.1 to v2.9.2
- cilium from v0.25.1 to v0.27.0
- node-exporter from v1.19.0 to v1.20.0
- observability-bundle from v1.5.3 to v1.6.1
- security-bundle from v1.8.0 to v1.8.2
- vertical-pod-autoscaler from v5.2.4 to v5.3.0
- vertical-pod-autoscaler-crd from v3.1.0 to v3.1.1

### azure-cloud-controller-manager [v1.29.8-gs1...v1.30.6-gs1](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v1.29.8-gs1...v1.30.6-gs1)

#### Changed

- Chart: Update to upstream v1.30.6. ([#87](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/87))

### azure-cloud-node-manager [v1.29.8-gs1...v1.30.6-gs1](https://github.com/giantswarm/azure-cloud-node-manager-app/compare/v1.29.8-gs1...v1.30.6-gs1)

#### Changed

- Chart: Update to upstream v1.30.6. ([#77](https://github.com/giantswarm/azure-cloud-node-manager-app/pull/77))

### cert-exporter [v2.9.1...v2.9.2](https://github.com/giantswarm/cert-exporter/compare/v2.9.1...v2.9.2)

#### Added

- Chart: Add VPA and resources configuration for deployment and daemonset. ([#382](https://github.com/giantswarm/cert-exporter/pull/382))

### cilium [v0.25.1...v0.27.0](https://github.com/giantswarm/cilium-app/compare/v0.25.1...v0.27.0)

#### Changed

- Upgrade Cilium to [v1.16.1](https://github.com/cilium/cilium/releases/tag/v1.16.1).
- Disable digest in all images.
- Improve security defaults for:
  - Hubble UI
  - Hubble Relay
  - Cilium Operator

### node-exporter [v1.19.0...v1.20.0](https://github.com/giantswarm/node-exporter-app/compare/v1.19.0...v1.20.0)

#### Changed

- Synced with upstream chart v4.38.0 (node-exporter 1.8.2).

### observability-bundle [v1.5.3...v1.6.1](https://github.com/giantswarm/observability-bundle/compare/v1.5.3...v1.6.1)

#### Added

- Add `alloy` v0.4.0 as `alloyMetrics`.

#### Changed

- Disable usage reporting to GrafanaLabs by:
  - Bumping `alloyLogs` and `alloyMetrics` to v0.4.1.
  - Bumping `grafanaAgent` to v0.4.6.
- Bump `alloyLogs` to v0.4.0.

### security-bundle [v1.8.0...v1.8.2](https://github.com/giantswarm/security-bundle/compare/v1.8.0...v1.8.2)

#### Changed

- Update `cloudnative-pg` (app) to v0.0.6.
- Update `trivy-operator` (app) to v0.10.0.
- Update `kyverno-policy-operator` (app) to v0.0.8.
- Update `kyverno` (app) to v0.17.16.

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
