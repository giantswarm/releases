# :zap: Giant Swarm Release v19.1.0 for AWS :zap:

<< Add description here >>

## Change details


### cert-operator [3.2.1](https://github.com/giantswarm/cert-operator/releases/tag/v3.2.1)

#### Fixed
- Fix rule names of PolicyException.



### cluster-operator [5.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.8.0)

#### Added
- Add ENI mode for Cilium on AWS.
- Consider new control-plane label.
#### Changed
- Propagate `global.podSecurityStandards.enforced` value set to `true` for PSS migration
- Rename function for better readbility.



### etcd [3.5.9](https://github.com/etcd-io/etcd/releases/tag/v3.5.9)

#### etcd server
- Fix [LeaseTimeToLive API may return keys to clients which have no read permission on the keys](https://github.com/etcd-io/etcd/pull/15815).
#### Dependencies
- Compile binaries using [go 1.19.9](https://github.com/etcd-io/etcd/pull/15822).



### kubernetes [1.24.17](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.17)

#### Feature
- Kubernetes is now built with Go 1.20.7 ([#119837](https://github.com/kubernetes/kubernetes/pull/119837), [@jeremyrickard](https://github.com/jeremyrickard)) [SIG Apps, Cloud Provider, Node, Release, Storage and Testing]
#### Bug or Regression
- Fixed a bug where clusters that use KMS v1 with skewed API servers on versions v1.24 and v1.25 would see internal errors when attempting to read encrypted data via the v1.24 API servers. ([#119387](https://github.com/kubernetes/kubernetes/pull/119387), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### app-operator [6.8.0](https://github.com/giantswarm/app-operator/releases/tag/v6.8.0)

#### Added
- Add Service Monitor by default to make it complain with the latest monitoring improvements 



### aws-operator [14.21.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.21.0)




### cert-manager [3.3.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.3.0)

⚠️ Attention: Major release [3.0.0](#300---2023-07-26) contains breaking changes in user values! Please make yourself familiar with its changelog! ⚠️
#### Added
- Add NetworkPolicies for controller and cainjector. ([#354](https://github.com/giantswarm/cert-manager-app/pull/354))



### cilium [0.12.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.12.0)

#### Added
- Support creating `CiliumNetworkPolicy` manifests that allow egress requests to DNS and proxy hosts
#### Changed
- Add missing conditional for PSP rendering of default-policies installer job



### external-dns [2.39.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.39.0)

#### Changed
- Replace monitoring labels with ServiceMonitor ([#296](https://github.com/giantswarm/external-dns-app/pull/296)).
- Update ATS to 0.4.1 and python deps ([#297](https://github.com/giantswarm/external-dns-app/pull/297)).



### node-exporter [1.17.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.17.1)

#### Changed
- fix apparmor annotation



### observability-bundle [0.8.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.8.0)

#### Changed
- Upgrade `prometheus-agent` to 0.6.0.
- Upgrade `prometheus-operator-app` and `prometheus-operator-crd` to 6.0.0.



### cilium-servicemonitors [0.1.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.2)

#### Changed
- Drop metrics with high cardinality.
- Increase scrape interval to 60s.



### aws-ebs-csi-driver [2.25.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.25.0)

#### Changed
- Updated ebs-csi-driver to `v1.21.0` and updated sidecar images.



### cert-exporter [2.6.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.6.0)




### cluster-autoscaler [1.24.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3)

#### Changed
- Change ScaleDownUtilizationThreshold default from 0.5 to 0.7
- Update cluster-autoscaler to version `1.24.3`.



### net-exporter [1.17.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.17.0)

#### Changed
- Add security context values to make chart comply to PodSecurityStandard restricted profile.



### vertical-pod-autoscaler [4.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v4.0.0)

#### Changed
WARNING: this version requires Cilium to run because of the dependency on the CiliumNetworkPolicy CRD
- Upgrade dependency chart to 9.2.0.
- Adjusted the resource and limits to accomodate larger clusters by default
- Adjusted the admission controller to give it more QPS against the API
- Adjusted the updater to give it more QPS against the API
- Adjusted the recommender to give it
  - more QPS against the API
  - doubling the memory in case of an OOMKilled event
  - Using the 95% percentile for the calculation of the CPU usage: should allow to scale up more precisely to account for spikes in CPU consumption of the workload
  - Adjusted the resource and limits to accomodate larger clusters by default
  - Calculating recommendations only for workloads which do have a VPA custom resource, instead of all workloads
  - Removed standard network policies to decrease maintenance burden
  - Fixed Cilium Network Policy to allow CRD jobs execution
  - Added Cilium Network Policy weight for an early execution
  - Disabled VPA for the updater pod otherwise it keeps on getting re-scheduled because the memory consumption varies a lot between reconsiling resources and idle
  - Disabled VPA for the recommender pod otherwise it keeps on getting re-scheduled because the memory consumption varies a lot between reconsiling resources and idle



### etcd-kubernetes-resources-count-exporter [1.4.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.4.0)

#### Changed
- Add Max memory (default 500Mi) for VPA.



