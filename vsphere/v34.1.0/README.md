# :zap: Giant Swarm Release v34.1.0 for vSphere :zap:

## Changes compared to v34.0.0

### Components

- Flatcar from v4459.2.2 to [v4459.2.3](https://www.flatcar.org/releases/#release-4459.2.3)
- Kubernetes from v1.34.3 to [v1.34.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.4)
- os-tooling from v1.26.3 to v1.26.4

### Apps

- cert-exporter from v2.9.15 to v2.9.16
- cilium from v1.3.4 to v1.4.0
- etcd-defrag from v1.2.3 to v1.2.4
- etcd-k8s-res-count-exporter from v1.10.12 to v1.10.13
- k8s-dns-node-cache from v2.9.1 to v2.9.2
- observability-bundle from v2.5.0 to v2.6.0
- security-bundle from v1.16.1 to v1.17.0

### cert-exporter [v2.9.15...v2.9.16](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.9.16)

#### Changed

- Go: Update dependencies.

### cilium [v1.3.4...v1.4.0](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.4.0)

#### Changed

- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.

### etcd-defrag [v1.2.3...v1.2.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.3...v1.2.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.37.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/78))

### etcd-k8s-res-count-exporter [v1.10.12...v1.10.13](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.12...v1.10.13)

#### Changed

- Migrate to App Build Suite (ABS) for Helm chart building.
- Go: Update dependencies.

### k8s-dns-node-cache [v2.9.1...v2.9.2](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.1...v2.9.2)

#### Changed

- Upgrade application to version 1.26.7 (includes coredns 1.13.1)

### observability-bundle [v2.5.0...v2.6.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.6.0)

#### Added

- Add KSM metrics for Gateway API resources

### security-bundle [v1.16.1...v1.17.0](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v1.17.0)

#### Changed

- Update `kyverno` (app) to v0.23.0.
- Update `kyverno-crds` (app) to v1.16.0.
- Update `reports-server` (app) to v0.1.0.
- Update `cloudnative-pg` (app) to v0.0.13.
- Update `kubescape` (app) to v0.0.5.
- Update `starboard-exporter` (app) to v1.0.2.
