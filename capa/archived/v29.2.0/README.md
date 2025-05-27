# :zap: Giant Swarm Release v29.2.0 for CAPA :zap:

## Changes compared to v29.1.0

### Components

- cluster-aws from v2.0.0 to v2.2.0
- Flatcar from v3975.2.0 to [v3975.2.1](https://www.flatcar.org/releases#release-3975.2.1)
- Kubernetes from v1.29.8 to [v1.29.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md#changelog-since-v1298)

### cluster-aws [v2.0.0...v2.2.0](https://github.com/giantswarm/cluster-aws/compare/v2.0.0...v2.2.0)

#### ⚠️ Breaking change

- Do not allow additional properties in the following fields in order to avoid unnoticed typos:

  - `global.connectivity.network`
  - `global.connectivity.network.pods`
  - `global.connectivity.network.services`
  - `global.connectivity.subnets[]`
  - `global.connectivity.topology`
  - `global.controlPlane`
  - `global.controlPlane.additionalSecurityGroups[]`
  - `global.controlPlane.machineHealthCheck`
  - `global.controlPlane.oidc`
  - `global.providerSpecific`
  - `global.providerSpecific.instanceMetadataOptions`

If you were using values like `global.controlPlane.containerdVolumeSizeGB` and `global.controlPlane.kubeletVolumeSizeGB`, please move to the new `.global.controlPlane.libVolumeSizeGB` which defines the size of disk volume used for `/var/lib` mount point.

#### Added

- Allow to enable `auditd` through `global.components.auditd.enabled` helm value.
- Chart: Support multiple service account issuers.\
  This is used for example in the migration from Vintage AWS clusters to CAPA. Multiple issuers were previously supported only through internal chart values (this change removes `internal.migration.irsaAdditionalDomain`). The internal annotation `aws.giantswarm.io/irsa-additional-domain` on AWSMachineTemplate objects is changed to plural `aws.giantswarm.io/irsa-trust-domains` on the AWSCluster object.

#### Changed

- Chart: Update `cluster` to v1.4.1.
- Set provider specific configuration for cilium CNI ENI values.
- Do not allow additional properties in most values in order to avoid unnoticed typos.
- Validate that machine pool availability zones belong to the selected region.
- CI: Bump release version.
- Apps: Use `catalog` from Release CR.

#### Removed

- Remove Cilium app deprecated values.
- Remove unused kubectl image Helm value.

### Apps

- aws-pod-identity-webhook from v1.16.0 to v1.17.0
- coredns from v1.21.0 to v1.22.0
- observability-bundle from v1.6.1 to v1.6.2
- security-bundle from v1.8.1 to v1.8.2
- teleport-kube-agent from v0.9.2 to v0.10.3
- vertical-pod-autoscaler from v5.2.4 to v5.3.0
- vertical-pod-autoscaler-crd from v3.1.0 to v3.1.1

### aws-pod-identity-webhook [v1.16.0...v1.17.0](https://github.com/giantswarm/aws-pod-identity-webhook-app/compare/v1.16.0...v1.17.0)

#### Changed

- Fix VPA being ineffective due to referring to a non-existing `Deployment` name

### coredns [v1.21.0...v1.22.0](https://github.com/giantswarm/coredns-app/compare/v1.21.0...v1.22.0)

#### Changed

- Update `coredns` image to [1.11.3](https://github.com/coredns/coredns/releases/tag/v1.11.3).

#### Removed

- Removed legacy Giant Swarm monitoring labels as coredns is monitored through a prometheus-operator generated servicemonitor.

### observability-bundle [v1.6.1...v1.6.2](https://github.com/giantswarm/observability-bundle/compare/v1.6.1...v1.6.2)

#### Changed

- Fixed `alloyMetrics` catalog

### security-bundle [v1.8.1...v1.8.2](https://github.com/giantswarm/security-bundle/compare/v1.8.1...v1.8.2)

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
