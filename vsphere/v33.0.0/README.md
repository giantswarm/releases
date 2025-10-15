# :zap: Giant Swarm Release v33.0.0 for vSphere :zap:

## Changes compared to v32.0.0

### Components

- cluster-vsphere from v2.0.0 to v3.2.0
- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)
- Kubernetes from v1.32.9 to [v1.33.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.33.md#v1.33.5)

### cluster-vsphere [v2.0.0...v3.2.0](https://github.com/giantswarm/cluster-vsphere/compare/v2.0.0...v3.2.0)

#### Changed

- Chart: Update `cluster` to v4.2.0.
- Chart: Update `cluster` to v4.1.0.
- Chart: Update `cluster` to v4.0.3.
- Chart: Update `cluster` to v4.0.2.
- Chart: Update `cluster` to v4.0.1.

### Apps

- cert-exporter from v2.9.9 to v2.9.12
- cert-manager from v3.9.2 to v3.9.3
- cilium from v1.3.0 to v1.3.1
- coredns from v1.27.0 to v1.28.2
- etcd-defrag from v1.0.8 to v1.2.1
- etcd-k8s-res-count-exporter from v1.10.7 to v1.10.9
- k8s-audit-metrics from v0.10.6 to v0.10.8
- node-exporter from v1.20.5 to v1.20.7
- observability-bundle from v2.2.2 to v2.3.2
- vertical-pod-autoscaler from v6.0.1 to v6.1.1
- vertical-pod-autoscaler-crd from v4.0.1 to v4.1.1

### cert-exporter [v2.9.9...v2.9.12](https://github.com/giantswarm/cert-exporter/compare/v2.9.9...v2.9.12)

#### Changed

- Go: Update dependencies.
- Chart: Add value to toggle creation of Daemonset resources.
- Go: Update dependencies.

### cert-manager [v3.9.2...v3.9.3](https://github.com/giantswarm/cert-manager-app/compare/v3.9.2...v3.9.3)

#### Changed

- Fix missing targetPort in `cainjector-service`

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### coredns [v1.27.0...v1.28.2](https://github.com/giantswarm/coredns-app/compare/v1.27.0...v1.28.2)

#### Changed

- Update `coredns` image to [1.13.1](https://github.com/coredns/coredns/releases/tag/v1.13.1).
- Add value to toggle creation of controlplane deployment.
- Update `coredns` image to [1.13.0](https://github.com/coredns/coredns/releases/tag/v1.13.0).

### etcd-defrag [v1.0.8...v1.2.1](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.8...v1.2.1)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.34.0. ([#62](https://github.com/giantswarm/etcd-defrag-app/pull/62))
- Chart: Update dependency ahrtr/etcd-defrag to v0.33.0. ([#60](https://github.com/giantswarm/etcd-defrag-app/pull/60))
- Update Kyverno API to v2 for policy exceptions
- Chart: Update dependency ahrtr/etcd-defrag to v0.32.0. ([#57](https://github.com/giantswarm/etcd-defrag-app/pull/57))

### etcd-k8s-res-count-exporter [v1.10.7...v1.10.9](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.7...v1.10.9)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.6...v0.10.8](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.6...v0.10.8)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### node-exporter [v1.20.5...v1.20.7](https://github.com/giantswarm/node-exporter-app/compare/v1.20.5...v1.20.7)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### observability-bundle [v2.2.2...v2.3.2](https://github.com/giantswarm/observability-bundle/compare/v2.2.2...v2.3.2)

#### Added

- Add KSM metrics for cloudnative-pg Cluster objects

#### Changed

- Update alloy-app to 0.15.0
  - Bumps alloy to 1.11.0

#### Fixed

- Update alloy-app to 0.15.1
  - Bumps alloy to 1.11.2

### vertical-pod-autoscaler [v6.0.1...v6.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.0.1...v6.1.1)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.1.1. ([#375](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/375))
- Chart: Update Helm release vertical-pod-autoscaler to v11.1.0. ([#372](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/372))

### vertical-pod-autoscaler-crd [v4.0.1...v4.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.0.1...v4.1.1)

#### Changed

- Chart: Sync to upstream. ([#166](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/166))
- Chart: Sync to upstream. ([#164](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/164))
