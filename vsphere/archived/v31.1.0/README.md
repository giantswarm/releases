# :zap: Giant Swarm Release v31.1.0 for vSphere :zap:

## Changes compared to v31.0.0

### Components

- cluster-vsphere from v1.4.0 to v1.5.0
- Kubernetes from v1.31.9 to [v1.31.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.11)

### cluster-vsphere [v1.4.0...v1.5.0](https://github.com/giantswarm/cluster-vsphere/compare/v1.4.0...v1.5.0)

#### Changed

- Chart: update `cluster` to v2.5.0.
- Chart: Update `kube-vip` to v0.9.2.

### Apps

- capi-node-labeler from v1.1.1 to v1.1.2
- cert-exporter from v2.9.7 to v2.9.8
- cilium from v1.2.1 to v1.2.2
- coredns from v1.25.0 to v1.26.0
- etcd-defrag from v1.0.5 to v1.0.6
- etcd-k8s-res-count-exporter from v1.10.5 to v1.10.6
- k8s-audit-metrics from v0.10.4 to v0.10.5
- k8s-dns-node-cache from v2.8.1 to v2.9.0
- node-exporter from v1.20.3 to v1.20.4
- security-bundle from v1.11.0 to v1.12.0
- teleport-kube-agent from v0.10.5 to v0.10.6


### capi-node-labeler [v1.1.1...v1.1.2](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.1.1...v1.1.2)

#### Changed

- Go: Update dependencies.

### cert-exporter [v2.9.7...v2.9.8](https://github.com/giantswarm/cert-exporter/compare/v2.9.7...v2.9.8)

#### Changed

- Go: Update dependencies.

### cilium [v1.2.1...v1.2.2](https://github.com/giantswarm/cilium-app/compare/v1.2.1...v1.2.2)

#### Changed

- Upgrade Cilium to [v1.17.6](https://github.com/cilium/cilium/releases/tag/v1.17.6).
- Updated E2E tests to use apptest-framework v1.14.0
- Increase Cilium operator resource limits.

#### Removed

- Remove deprecated "partial" mode from Kube Proxy Replacement options.

### coredns [v1.25.0...v1.26.0](https://github.com/giantswarm/coredns-app/compare/v1.25.0...v1.26.0)

#### Changed

- Update `coredns` image to [1.12.2](https://github.com/coredns/coredns/releases/tag/v1.12.2).

### etcd-defrag [v1.0.5...v1.0.6](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.5...v1.0.6)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.29.0. ([#43](https://github.com/giantswarm/etcd-defrag-app/pull/43))

### etcd-k8s-res-count-exporter [v1.10.5...v1.10.6](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.5...v1.10.6)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.4...v0.10.5](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.4...v0.10.5)

#### Changed

- Go: Update dependencies.

### k8s-dns-node-cache [v2.8.1...v2.9.0](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.8.1...v2.9.0)

#### Changed

- Upgrade application to version 1.26.4 (includes coredns 1.11.3)
- Increase ServiceMonitor's scrapping interval to 1m.
- Remove obsolete PSPs

### node-exporter [v1.20.3...v1.20.4](https://github.com/giantswarm/node-exporter-app/compare/v1.20.3...v1.20.4)

#### Changed

- Go: Update to v1.24.5.

### security-bundle [v1.11.0...v1.12.0](https://github.com/giantswarm/security-bundle/compare/v1.11.0...v1.12.0)

#### Changed

- Update `trivy-operator` (app) to v0.11.1.
- Update `trivy` (app) to v0.14.0.
- Update `falco` (app) to v0.10.1.
- Update `cloudnative-pg` (app) to v0.0.10.
- Update `starboard-exporter` (app) to v0.8.2.
- Updated E2E tests to use apptest-framework v1.14.0

### teleport-kube-agent [v0.10.5...v0.10.6](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.5...v0.10.6)

#### Changed

- AppVersion upgrade to 17.5.4
