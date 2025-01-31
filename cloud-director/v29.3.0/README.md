# :zap: Giant Swarm Release v29.3.0 for VMware Cloud Director :zap:

## Changes compared to v29.2.0

### Components

- cluster-cloud-director from v0.64.0 to v0.64.2
- Flatcar from v3975.2.2 to [v4081.2.1](https://www.flatcar.org/releases#release-4081.2.1)
- Kubernetes from v1.29.12 to [v1.29.13](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md#changelog-since-v12912)

### cluster-cloud-director [v0.64.0...v0.64.2](https://github.com/giantswarm/cluster-cloud-director/compare/v0.64.0...v0.64.2)

#### Added

- Add `components.containerd` to the schema and values.

#### Changed

- Make CPI helmrelease catalog configurable.

### Apps

- cilium from v0.25.1 to v0.25.2
- prometheus-blackbox-exporter from v0.4.2 to v0.5.0
- security-bundle from v1.8.2 to v1.9.1
- vertical-pod-autoscaler from v5.3.0 to v5.3.1
- vertical-pod-autoscaler-crd from v3.1.1 to v3.1.2

### cilium [v0.25.1...v0.25.2](https://github.com/giantswarm/cilium-app/compare/v0.25.1...v0.25.2)

#### Changed

- Upgrade cilium to [v1.15.13](https://github.com/cilium/cilium/releases/tag/v1.15.13).

### prometheus-blackbox-exporter [v0.4.2...v0.5.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.4.2...v0.5.0)

#### Changed

- Harden security context to pass PSS compliance.

#### Removed

- Remove PSP resources.

### security-bundle [v1.8.2...v1.9.1](https://github.com/giantswarm/security-bundle/compare/v1.8.2...v1.9.1)

#### Breaking changes

**Note:** When upgrading to this security-bundle version with Falco enabled, the Falco App will fail to upgrade due to a breaking change in the upstream chart. To finish the upgrade, disable, then re-enable the Falco App by setting `apps.falco.enabled=[false|true]` [in the security-bundle user values Config Map](https://github.com/giantswarm/security-bundle/tree/main?tab=readme-ov-file#configuring).

#### Changed

- Update `trivy-operator` (app) to v0.10.3.
- Update `trivy` (app) to v0.13.1.
- Update `kyverno` (app) to v0.18.1.
- Update `kyverno-crds` (app) to v1.12.0.
- Update `kyverno-policies` (app) to v0.21.0.
- Update `starboard-exporter` (app) to v0.8.0.
- Update `falco` (app) to v0.9.1.

### vertical-pod-autoscaler [v5.3.0...v5.3.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.3.0...v5.3.1)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.9.1. ([#333](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/333))

### vertical-pod-autoscaler-crd [v3.1.1...v3.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.1.1...v3.1.2)

#### Changed

- Chart: Sync to upstream. ([#124](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/124))
