# :zap: Giant Swarm Release v32.0.0 for vSphere :zap:

> **WARNING:** With Flatcar 4230.2.0, cgroups v1 backwards compatibility has been removed. This means that enabling legacy cgroups v1 is no longer supported and nodes still using them will fail to update.

## Changes compared to v31.1.1

### Components

- cluster-vsphere from v1.5.1 to v2.0.0
- Flatcar from v4152.2.3 to [v4230.2.2](https://www.flatcar-linux.org/releases/#release-4230.2.2)
- Kubernetes from v1.31.11 to [v1.32.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.9)

### cluster-vsphere [v1.5.1...v2.0.0](https://github.com/giantswarm/cluster-vsphere/compare/v1.5.1...v2.0.0)

#### Changed

- Chart: Update `cluster` to v3.0.1.
  - **BREAKING CHANGE:** Cgroups v1 is not supported anymore. The `.internal.advancedConfiguration.cgroupsv1` and `.global.nodePools.().cgroupsv1` flags have been removed.
  - Chart: Simplify containerd configuration by using a single config file for both control-plane and worker nodes.
- Chart: Update `cluster` to v2.6.2.
- Update `kube-vip` static pod manifest to `v1.0.0`.
- Chart: Update `cluster` to v2.6.1.

### Apps

- capi-node-labeler from v1.1.2 to v1.1.3
- cert-exporter from v2.9.8 to v2.9.9
- cert-manager from v3.9.1 to v3.9.2
- cilium from v1.2.2 to v1.3.0
- coredns from v1.26.0 to v1.27.0
- etcd-defrag from v1.0.6 to v1.0.8
- etcd-k8s-res-count-exporter from v1.10.6 to v1.10.7
- k8s-audit-metrics from v0.10.5 to v0.10.6
- k8s-dns-node-cache from v2.9.0 to v2.9.1
- metrics-server from v2.6.0 to v2.7.0
- node-exporter from v1.20.4 to v1.20.5
- observability-bundle from v2.0.0 to v2.2.2
- vertical-pod-autoscaler from v5.5.1 to v6.0.1
- vertical-pod-autoscaler-crd from v3.3.1 to v4.0.1

### capi-node-labeler [v1.1.2...v1.1.3](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.1.2...v1.1.3)

#### Changed

- Go: Update dependencies.

### cert-exporter [v2.9.8...v2.9.9](https://github.com/giantswarm/cert-exporter/compare/v2.9.8...v2.9.9)

#### Changed

- Go: Update dependencies.

### cert-manager [v3.9.1...v3.9.2](https://github.com/giantswarm/cert-manager-app/compare/v3.9.1...v3.9.2)

#### Changed

- Add `alloy` ingress rules for cainjector metrics ingestion.

### cilium [v1.2.2...v1.3.0](https://github.com/giantswarm/cilium-app/compare/v1.2.2...v1.3.0)

#### Changed

- Upgrade Cilium to [v1.18.1](https://github.com/cilium/cilium/releases/tag/v1.18.1).
- Improve the k8s service host autodiscovery mechanism
- Upgrade Cilium to [v1.17.7](https://github.com/cilium/cilium/releases/tag/v1.17.7).

### coredns [v1.26.0...v1.27.0](https://github.com/giantswarm/coredns-app/compare/v1.26.0...v1.27.0)

#### Changed

- Updated E2E tests to use apptest-framework v1.14.0
- Update `coredns` image to [1.12.3](https://github.com/coredns/coredns/releases/tag/v1.12.3).

### etcd-defrag [v1.0.6...v1.0.8](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.6...v1.0.8)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.31.0. ([#52](https://github.com/giantswarm/etcd-defrag-app/pull/52))
- Chart: Update dependency ahrtr/etcd-defrag to v0.30.0. ([#46](https://github.com/giantswarm/etcd-defrag-app/pull/46))

### etcd-k8s-res-count-exporter [v1.10.6...v1.10.7](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.6...v1.10.7)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.5...v0.10.6](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.5...v0.10.6)

#### Changed

- Go: Update dependencies.

### k8s-dns-node-cache [v2.9.0...v2.9.1](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.0...v2.9.1)

#### Changed

- Update PolicyException apiVersion to v2.

### metrics-server [v2.6.0...v2.7.0](https://github.com/giantswarm/metrics-server-app/compare/v2.6.0...v2.7.0)

#### Changed

- Chart: Update PolicyExceptions to v2.

### node-exporter [v1.20.4...v1.20.5](https://github.com/giantswarm/node-exporter-app/compare/v1.20.4...v1.20.5)

#### Changed

- Go: Update dependencies.

### observability-bundle [v2.0.0...v2.2.2](https://github.com/giantswarm/observability-bundle/compare/v2.0.0...v2.2.2)

#### Added

- Add KSM metrics for IRSAClaim objects

#### Changed

- Upgrade `kube-prometheus-stack-app` to 18.1.0
  - Add relabeling rules from `cluster-api-monitoring-app` so that `cluster_id` label points to the workload cluster name as expected in some alert definitions
- Upgrade `kube-prometheus-stack` to 77.0.1
  - Bumps prometheus-operator and CRDs to 0.85.0
- Update alloy-app to 0.13.0
- Upgrade `kube-prometheus-stack` to 76.4.0
  - Bumps prometheus-operator and CRDs to 0.84.1
  - Bumps prometheus to 3.5.0
- Update alloy-app to 0.12.1
  - Bumps alloy to 1.10.1

### vertical-pod-autoscaler [v5.5.1...v6.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.5.1...v6.0.1)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.0.1. ([#370](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/370))
- Chart: Update Helm release vertical-pod-autoscaler to v11.0.0. ([#362](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/362))

### vertical-pod-autoscaler-crd [v3.3.1...v4.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.3.1...v4.0.1)

#### Changed

- Chart: Sync to upstream. ([#162](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/162))
- Chart: Sync to upstream. ([#154](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/154))
