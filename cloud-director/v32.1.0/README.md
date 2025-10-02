# :zap: Giant Swarm Release v32.1.0 for VMware Cloud Director :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- cluster-cloud-director from v1.0.0 to v2.0.0
- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### cluster-cloud-director [v1.0.0...v2.0.0](https://github.com/giantswarm/cluster-cloud-director/compare/v1.0.0...v2.0.0)

#### Changed

- Chart: Update `cluster` to v4.0.1.

### Apps

- cert-exporter from v2.9.9 to v2.9.10
- cilium from v1.3.0 to v1.3.1
- etcd-defrag from v1.0.8 to v1.1.0
- etcd-k8s-res-count-exporter from v1.10.7 to v1.10.8
- k8s-audit-metrics from v0.10.6 to v0.10.7
- node-exporter from v1.20.5 to v1.20.6
- vertical-pod-autoscaler from v6.0.1 to v6.1.0
- vertical-pod-autoscaler-crd from v4.0.1 to v4.1.0

### cert-exporter [v2.9.9...v2.9.10](https://github.com/giantswarm/cert-exporter/compare/v2.9.9...v2.9.10)

#### Changed

- Go: Update dependencies.

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### etcd-defrag [v1.0.8...v1.1.0](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.8...v1.1.0)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Chart: Update dependency ahrtr/etcd-defrag to v0.32.0. ([#57](https://github.com/giantswarm/etcd-defrag-app/pull/57))

### etcd-k8s-res-count-exporter [v1.10.7...v1.10.8](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.7...v1.10.8)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.6...v0.10.7](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.6...v0.10.7)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### node-exporter [v1.20.5...v1.20.6](https://github.com/giantswarm/node-exporter-app/compare/v1.20.5...v1.20.6)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### vertical-pod-autoscaler [v6.0.1...v6.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.0.1...v6.1.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.1.0. ([#372](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/372))

### vertical-pod-autoscaler-crd [v4.0.1...v4.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.0.1...v4.1.0)

#### Changed

- Chart: Sync to upstream. ([#164](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/164))
