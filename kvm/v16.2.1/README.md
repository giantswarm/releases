# :zap: Giant Swarm Release v16.2.1 for KVM :zap:

This release updates several preinstalled apps running in the workload cluster to reduce Kubernetes API server load and etcd memory usage in certain installations.

## Change details


### kvm-operator [3.18.6](https://github.com/giantswarm/app-operator/releases/tag/v3.18.6)

#### Fixed
- Fix error rendering multiple registry mirrors in the configmap.
