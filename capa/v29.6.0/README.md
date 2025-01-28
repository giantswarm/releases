# :zap: Giant Swarm Release v29.6.0 for CAPA :zap:

Most notable change in this release is the reduction of IAM permissions on the worker nodes instance profile, aiming at improving the general security of the clusters. Additional changes include reducing the size of the ETCD volume to 50GB targetting costs saving initiatives, as well as improvements for the `node-termination-handler` application for smoother upgrades and operations. Several components such as Flatcar or Kubernetes have also been updated to the latest available versions.

## Changes compared to v29.5.0

### Components

- cluster-aws from v2.5.0 to v2.6.0
- Flatcar from v3975.2.2 to [v4081.2.1](https://www.flatcar.org/releases#release-4081.2.1)
- Kubernetes from v1.29.12 to [v1.29.13](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md#changelog-since-v12912)

### cluster-aws [v2.5.0...v2.6.0](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...v2.6.0)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools
- Use reduced IAM permissions on worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.

#### Fixed

- Explicitly set aws-node-termination-handler queue region so crash-loops are avoided, allowing faster startup

### Apps

- aws-nth-bundle from v1.2.0 to v1.2.1
- aws-pod-identity-webhook from v1.17.0 to v1.18.0
- cilium from v0.25.1 to v0.25.2
- prometheus-blackbox-exporter from v0.4.2 to v0.5.0
- security-bundle from v1.8.2 to v1.9.1
- vertical-pod-autoscaler from v5.3.0 to v5.3.1
- vertical-pod-autoscaler-crd from v3.1.1 to v3.1.2

### aws-nth-bundle [v1.2.0...v1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)

#### Added

- Forward proxy settings to `aws-node-termination-handler-app` as environment variables

### aws-pod-identity-webhook [v1.17.0...v1.18.0](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.17.0...v1.18.0)

#### Changed

- Update `securityContext` to be compliant.

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
