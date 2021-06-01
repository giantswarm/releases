# :zap: Giant Swarm Release v14.1.0 for KVM :zap:

This release introduces a new feature that allows KVM clusters to run behind a proxy. It also adds support for host volumes.

## Change details


### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

#### Changed
- Dropped ensuring cluster CRDs from controllers.



### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases.
- Emit events when config maps and secrets referenced in App CRs are updated.



### kvm-operator [3.17.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.17.1)

#### Changed
- Reconcile only deployments that are managed by kvm-operator



