# :zap: Giant Swarm Release v31.0.0 for VMWare Cloud Director :zap:

<< Add description here >>

## Changes compared to v30.1.3

### Components

- Flatcar from 4152.2.1 to [4152.2.3](https://www.flatcar-linux.org/releases/#release-4152.2.3)
- Kubernetes from v1.30.11 to [v1.31.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.9)

### Apps

- cert-manager from 3.9.0 to 3.9.1
- cilium from 0.31.5 to 1.2.0
- coredns from 1.24.0 to 1.25.0
- etcd-defrag from 1.0.2 to 1.0.4
- observability-bundle from 1.11.0 to 1.16.0
- observability-policies from 0.0.1 to 0.0.2
- security-bundle from 1.10.1 to 1.11.0
- teleport-kube-agent from 0.10.4 to 0.10.5
- vertical-pod-autoscaler from 5.4.0 to 5.5.0
- vertical-pod-autoscaler-crd from 3.2.0 to 3.3.0


### cert-manager [v3.9.0...v3.9.1](https://github.com/giantswarm/cert-manager-app/compare/v3.9.0...v3.9.1)

#### Added

- Added Vertical Pod Autoscaler support for `controller` pods.
- Added renovate configutarion

### cilium [v0.31.5...v1.2.0](https://github.com/giantswarm/cilium-app/compare/v0.31.5...v1.2.0)

#### Changed

- Re-enable Cilium agent and operator metrics port.
- Add resource requests and limits to Hubble UI and Relay.
- Add resource requests and limits to Cilium Operator.
- Upgrade Cilium to [v1.17.4](https://github.com/cilium/cilium/releases/tag/v1.17.4).
- Cilium v1.17.4 disables kubernetes api connectivity check for liveness probes. (Upstream PR: https://github.com/cilium/cilium/pull/38703)
- Upgrade Cilium to [v1.17.3](https://github.com/cilium/cilium/releases/tag/v1.17.3).
- Upgrade Cilium to [v1.17.2](https://github.com/cilium/cilium/releases/tag/v1.17.2).
- Remove cleanup kube-proxy patch.
- Identity computation label exclusion list regular expressions. Remove `controller-uid`, since this is excluded by default now.
- Upgrade Cilium to [v1.17.0](https://github.com/cilium/cilium/releases/tag/v1.17.0).
- Use upstream default value for `prometheus.metrics`.
- Enable Envoy Proxy in standalone DaemonSet.

### coredns [v1.24.0...v1.25.0](https://github.com/giantswarm/coredns-app/compare/v1.24.0...v1.25.0)

#### Changed

- Update `coredns` image to [1.12.1](https://github.com/coredns/coredns/releases/tag/v1.12.1).

### etcd-defrag [v1.0.2...v1.0.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.2...v1.0.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.27.0. ([#29](https://github.com/giantswarm/etcd-defrag-app/pull/29))
- Chart: Update dependency ahrtr/etcd-defrag to v0.26.0. ([#22](https://github.com/giantswarm/etcd-defrag-app/pull/22))

### observability-bundle [v1.11.0...v1.16.0](https://github.com/giantswarm/observability-bundle/compare/v1.11.0...v1.16.0)

#### Added

- Add support for enabling pre-configured custom resources in KSM
- Add metrics containing labels for Crossplane resources

#### Changed

- Upgrade alloy-app from 0.10.0 to 0.11.0
  - This bumps the version of Alloy from 1.8.3 to 1.9.0
- Upgrade `alloy-app` from 0.9.0 to 0.10.0
  - This bumps the version of Alloy from 1.7.1 to 1.8.3
- Reconfigure Flux-related part of the KSM to use wildcards instead of hardcoded versions.
- Rename Flux-related metrics produced by the KSM.
- Upgrade `kube-prometheus-stack` to 72.3.0
  - Bumps prometheus-operator to 0.82.0
  - Bumps prometheus-operator CRDs to 0.82.0
- Upgrade `kube-prometheus-stack` to 72.3.0
  - Bumps prometheus-operator to 0.82.0
- Upgrade `kube-prometheus-stack` from 69.5.1 to 70.1.1
  - Bumps prometheus-operator to 0.81.0
  - Bumps prometheus to 3.2.1

### observability-policies [v0.0.1...v0.0.2](https://github.com/giantswarm/observability-policies-app/compare/v0.0.1...v0.0.2)

#### Changed

- Add Cluster Role to allow latest Kyverno versions to work (https://github.com/giantswarm/giantswarm/issues/33416)
- Switch `.Values.disabled` to `.Values.enabled` to follow best practices.

### security-bundle [v1.10.1...v1.11.0](https://github.com/giantswarm/security-bundle/compare/v1.10.1...v1.11.0)

#### Added

- Add `policy-api-crds` app to manage Policy API CRDs.

#### Changed

- Update `trivy` (app) to v0.13.4.
- Update `cloudnative-pg` (app) to v0.0.7.
- Update `starboard-exporter` (app) to v0.8.1.
- Update `kyverno-policy-operator` (app) to v0.0.11.
- Update `cloudnative-pg` (app) to v0.0.9.

### teleport-kube-agent [v0.10.4...v0.10.5](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.4...v0.10.5)

#### Added

- Set Home URL in chart metadata.

### vertical-pod-autoscaler [v5.4.0...v5.5.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.4.0...v5.5.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v10.1.0. ([#350](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/350))
- Chart: Update Helm release vertical-pod-autoscaler to v10.2.0. ([#351](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/351))
- Chart: Update Helm release vertical-pod-autoscaler to v10.0.1. ([#346](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/346))

### vertical-pod-autoscaler-crd [v3.2.0...v3.3.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.2.0...v3.3.0)

#### Changed

- Chart: Sync to upstream. ([#140](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/140))
- Chart: Sync to upstream. ([#136](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/136))
