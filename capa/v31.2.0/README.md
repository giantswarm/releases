# :zap: Giant Swarm Release v31.2.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v31.1.0

### Components

- Flatcar from v4152.2.3 to [v4230.2.1](https://www.flatcar-linux.org/releases/#release-4230.2.1)
- Kubernetes from v1.31.11 to [v1.31.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.12)

### Apps

- etcd-defrag from v1.0.6 to v1.0.7
- vertical-pod-autoscaler from v5.5.1 to v6.0.0
- vertical-pod-autoscaler-crd from v3.3.1 to v4.0.0


### etcd-defrag [v1.0.6...v1.0.7](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.6...v1.0.7)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.30.0. ([#46](https://github.com/giantswarm/etcd-defrag-app/pull/46))

### vertical-pod-autoscaler [v5.5.1...v6.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.5.1...v6.0.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.0.0. ([#362](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/362))

### vertical-pod-autoscaler-crd [v3.3.1...v4.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.3.1...v4.0.0)

#### Changed

- Chart: Sync to upstream. ([#154](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/154))
