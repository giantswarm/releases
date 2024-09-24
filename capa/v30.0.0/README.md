# :zap: Giant Swarm Release v30.0.0 for CAPA :zap:

## Changes compared to v29.1.0

### Components

- cluster-aws from v2.0.0 to v2.1.0
- Kubernetes from v1.29.8 to [v1.30.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md)

### cluster-aws [v2.0.0...v2.1.0](https://github.com/giantswarm/cluster-aws/compare/v2.0.0...v2.1.0)

#### Changed

- Do not allow additional properties in most values in order to avoid unnoticed typos.
- Validate that machine pool availability zones belong to the selected region.
- CI: Bump release version.
- Apps: Use `catalog` from Release CR.

#### Removed

- Remove unused kubectl image Helm value.

### Apps

- cilium from v0.25.1 to v0.27.0
- cloud-provider-aws from v1.29.3-gs1 to v1.30.3-gs1
- cluster-autoscaler from v1.29.3-gs1 to v1.30.2-gs1
- security-bundle from v1.8.1 to v1.8.2
- vertical-pod-autoscaler from v5.2.4 to v5.3.0
- vertical-pod-autoscaler-crd from v3.1.0 to v3.1.1

### cilium [v0.25.1...v0.27.0](https://github.com/giantswarm/cilium-app/compare/v0.25.1...v0.27.0)

#### Changed

- Upgrade Cilium to [v1.16.1](https://github.com/cilium/cilium/releases/tag/v1.16.1).
- Disable digest in all images.
- Improve security defaults for:
  - Hubble UI
  - Hubble Relay
  - Cilium Operator

### cloud-provider-aws [v1.29.3-gs1...v1.30.3-gs1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.29.3-gs1...v1.30.3-gs1)

#### Changed

- Chart: Update to upstream v1.30.3.

### cluster-autoscaler [v1.29.3-gs1...v1.30.2-gs1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.29.3-gs1...v1.30.2-gs1)

#### Changed

- Update `PolicyExceptions` apiVersion to `v2beta1`. ([#282](https://github.com/giantswarm/cluster-autoscaler-app/pull/282))
- Chart: Update to upstream v1.30.2. ([#284](https://github.com/giantswarm/cluster-autoscaler-app/pull/284))

### security-bundle [v1.8.1...v1.8.2](https://github.com/giantswarm/security-bundle/compare/v1.8.1...v1.8.2)

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
