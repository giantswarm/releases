# :zap: Giant Swarm Release v30.1.0 for VMware Cloud Director :zap:

## Changes compared to v30.0.0

### Components

- cluster-cloud-director from v0.65.0 to v0.66.0
- Kubernetes from v1.30.10 to [v1.30.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md)
- os-tooling from v1.23.1 to v1.24.0

### cluster-cloud-director [v0.65.0...v0.66.0](https://github.com/giantswarm/cluster-cloud-director/compare/v0.65.0...v0.66.0)

#### Changed

- Chart: Update `cluster` to v2.2.0.

### os-tooling [v1.23.1...v1.24.0](https://github.com/giantswarm/capi-image-builder/compare/v1.23.1...v1.24.0)

#### Added

- Added nvidia_runtime to allow running of GPU workloads

### Apps

- capi-node-labeler from v1.0.1 to v1.0.2
- cert-exporter from v2.9.4 to v2.9.5
- cilium from v0.31.0 to v0.31.1
- etcd-defrag from v1.0.1 to v1.0.2
- etcd-kubernetes-resources-count-exporter from v1.10.1 to v1.10.3
- k8s-audit-metrics from v0.10.1 to v0.10.2
- net-exporter from v1.21.0 to v1.22.0
- node-exporter from v1.20.1 to v1.20.2
- observability-bundle from v1.9.0 to v1.11.0
- security-bundle from v1.9.1 to v1.10.0
- teleport-kube-agent from v0.10.3 to v0.10.4

### capi-node-labeler [v1.0.1...v1.0.2](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.0.1...v1.0.2)

#### Changed

- Go: Update dependencies.

### cert-exporter [v2.9.4...v2.9.5](https://github.com/giantswarm/cert-exporter/compare/v2.9.4...v2.9.5)

#### Changed

- Go: Update dependencies.

### cilium [v0.31.0...v0.31.1](https://github.com/giantswarm/cilium-app/compare/v0.31.0...v0.31.1)

#### Changed

- Upgrade Cilium to [v1.16.7](https://github.com/cilium/cilium/releases/tag/v1.16.7).

### etcd-defrag [v1.0.1...v1.0.2](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.1...v1.0.2)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.25.0. ([#17](https://github.com/giantswarm/etcd-defrag-app/pull/17))

### etcd-kubernetes-resources-count-exporter [v1.10.1...v1.10.3](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.1...v1.10.3)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.1...v0.10.2](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.1...v0.10.2)

#### Changed

- Go: Update dependencies.

### net-exporter [v1.21.0...v1.22.0](https://github.com/giantswarm/net-exporter/compare/v1.21.0...v1.22.0)

#### Changed

- Narrow down CiliumNetworkPolicy to allow desired traffic only.

#### Removed

- Remove NetworkPolicy resource and rely on CiliumNetworkPolicy only.

### node-exporter [v1.20.1...v1.20.2](https://github.com/giantswarm/node-exporter-app/compare/v1.20.1...v1.20.2)

#### Changed

- Go: Update dependencies.

### observability-bundle [v1.9.0...v1.11.0](https://github.com/giantswarm/observability-bundle/compare/v1.9.0...v1.11.0)

#### Changed

- prometheus-operator will not check promql syntax for prometheusRules that are labelled `observability.giantswarm.io/rule-type: logs`
- Upgrade `alloy` to chart 0.9.0.
  - Bumps `alloy` from to 1.5.1 to 1.7.1
- Upgrade `kube-prometheus-stack` from 66.2.1 to 69.5.1
  - Bumps prometheus-operator to 0.80.1
  - Bumps prometheus to 3.0.1

### security-bundle [v1.9.1...v1.10.0](https://github.com/giantswarm/security-bundle/compare/v1.9.1...v1.10.0)

#### Added

- Add e2e tests for the `security-bundle` and all is components

#### Changed

- Update `kyverno` (app) to v0.19.0.
- Update `kyverno-crds` (app) to v1.13.0.
- Update `kyverno-policies` (app) to v0.23.0.
- Update `edgedb` (app) to v0.1.0.
- Update `falco` (app) to v0.10.0.
- Update `trivy` (app) to v0.13.2.

### teleport-kube-agent [v0.10.3...v0.10.4](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.3...v0.10.4)

#### Added

- Add headless service on `diag` port 3000.

#### Changed

- Migrated to ABS
