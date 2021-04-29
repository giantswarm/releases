# :zap: Giant Swarm Release v14.0.0 for KVM :zap:

This release upgrades Kubernetes to 1.19. It also includes other minor component updates summarized below.

## Change details


### kvm-operator [3.16.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.16.0)

#### Changed

- Update k8scloudconfig to use calico-crd-installer.


### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

#### Security
- Restrict ingress to only expose the status endpoint.



### kubernetes [1.19.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.9)

#### Failing Test
- Fix handing special characters in the volume path on Windows ([#99137](https://github.com/kubernetes/kubernetes/pull/99137), [@yujuhong](https://github.com/yujuhong)) [SIG Storage]
#### Bug or Regression
- Count pod overhead against an entity's ResourceQuota ([#99600](https://github.com/kubernetes/kubernetes/pull/99600), [@gjkim42](https://github.com/gjkim42)) [SIG API Machinery and Node]
- EndpointSlice controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100114](https://github.com/kubernetes/kubernetes/pull/100114), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSliceMirroring controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100144](https://github.com/kubernetes/kubernetes/pull/100144), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Fixed bug that caused cAdvisor to incorrectly detect single-socket multi-NUMA topology. ([#99209](https://github.com/kubernetes/kubernetes/pull/99209), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixes kubelet to retrieve number of sockets from cAdvisor MachineInfo, instead of assuming it to be equal to number of NUMA nodes. ([#99771](https://github.com/kubernetes/kubernetes/pull/99771), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#98140](https://github.com/kubernetes/kubernetes/pull/98140), [@CKchen0726](https://github.com/CKchen0726)) [SIG Apps and Node]
- Kubelet now cleans up orphaned volume directories automatically ([#95301](https://github.com/kubernetes/kubernetes/pull/95301), [@lorenz](https://github.com/lorenz)) [SIG Node and Storage]
- Resolves spurious `Failed to list *v1.Secret` or `Failed to list *v1.ConfigMap` messages in kubelet logs. ([#99538](https://github.com/kubernetes/kubernetes/pull/99538), [@liggitt](https://github.com/liggitt)) [SIG Auth and Node]
- Using NUMA nodes instead of sockets for CPU manager hints. ([#99276](https://github.com/kubernetes/kubernetes/pull/99276), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- We will no longer automatically delete all data when a failure is detected during creation of the volume data file on a CSI volume. Now we will only remove the data file and volume path. ([#96021](https://github.com/kubernetes/kubernetes/pull/96021), [@huffmanca](https://github.com/huffmanca)) [SIG Storage]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.37.4 → v0.37.5](https://github.com/google/cadvisor/compare/v0.37.4...v0.37.5)
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.9 → v0.0.15
#### Removed
_Nothing has changed._



### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Changed
- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed
- Set docker.io as the default registry

### net-exporter [1.10.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.0)

#### Changed

- Add label selector for pods to help lower memory usage.
