# :zap: Giant Swarm Release v16.2.0 for KVM :zap:

This release updates several preinstalled apps running in the workload cluster to reduce Kubernetes API server load and etcd memory usage in certain installations.

## Change details


### app-operator [5.6.0](https://github.com/giantswarm/app-operator/releases/tag/v5.6.0)

#### Changed
- Get tarball URL for chart CRs from index.yaml for better community app catalog support.
#### Fixed
- Fix error handling in chart CR watcher when chart CRD not installed.



### cert-operator [1.3.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.3.0)

#### Changed
- Use `RenewSelf` instead of `LookupSelf` to prevent expiration of `Vault token`.



### cert-exporter [2.1.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.1.0)

#### Changed
- Make exporter's monitor flags configurable.



### kvm-operator [3.18.5](https://github.com/giantswarm/kvm-operator/releases/tag/v3.18.5)

#### Fixed
- Update k8scloudconfig to v10.16.1 for calico permissions fix.
