# :zap: Giant Swarm Release v32.0.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v31.1.0

### Components

- Flatcar from v4152.2.3 to [v4230.2.1](https://www.flatcar-linux.org/releases/#release-4230.2.1)
- Kubernetes from v1.31.11 to [v1.32.8](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.8)

### Apps

- coredns from v1.26.0 to v1.27.0
- etcd-defrag from v1.0.6 to v1.0.7
- vertical-pod-autoscaler from v5.5.1 to v6.0.0
- vertical-pod-autoscaler-crd from v3.3.1 to v4.0.0


### coredns [v1.26.0...v1.27.0](https://github.com/giantswarm/coredns-app/compare/v1.26.0...v1.27.0)

#### Changed

- Updated E2E tests to use apptest-framework v1.14.0
- Update `coredns` image to [1.12.3](https://github.com/coredns/coredns/releases/tag/v1.12.3).

### etcd-defrag [v1.0.6...v1.0.7](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.6...v1.0.7)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.30.0. ([#46](https://github.com/giantswarm/etcd-defrag-app/pull/46))

### vertical-pod-autoscaler [v5.5.1...v6.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.5.1...v6.0.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.0.0. ([#362](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/362))

### vertical-pod-autoscaler-crd [v3.3.1...v4.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.3.1...v4.0.0)

#### Changed

- Chart: Sync to upstream. ([#154](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/154))
