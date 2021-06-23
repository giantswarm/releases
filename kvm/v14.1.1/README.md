# :zap: Giant Swarm Release v14.1.1 for KVM :zap:

This release fixes a rare bug encountered when deleting clusters using host volumes.

## Change details

### kvm-operator [3.17.3](https://github.com/giantswarm/kvm-operator/releases/tag/v3.17.3)

#### Fixed

- Avoid panic during deletion of clusters with host volumes.
