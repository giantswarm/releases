# :zap: Giant Swarm Release v18.0.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.23.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.8)

#### Feature
- Kubernetes is now built with Golang 1.17.11 ([#110423](https://github.com/kubernetes/kubernetes/pull/110423), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- EndpointSlices marked for deletion are now ignored during reconciliation. ([#110483](https://github.com/kubernetes/kubernetes/pull/110483), [@aryan9600](https://github.com/aryan9600)) [SIG Apps and Network]
- Fixed a kubelet issue that could result in invalid pod status updates to be sent to the api-server where pods would be reported in a terminal phase but also report a ready condition of true in some cases. ([#110480](https://github.com/kubernetes/kubernetes/pull/110480), [@bobbypage](https://github.com/bobbypage)) [SIG Node and Testing]
- Pods will now post their readiness during termination. ([#110417](https://github.com/kubernetes/kubernetes/pull/110417), [@aojea](https://github.com/aojea)) [SIG Network, Node and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### etcd [3.5.4](https://github.com/etcd-io/etcd/releases/tag/v3.5.4)

Not found


### aws-ebs-csi-driver [2.16.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.16.0)

#### Changed
- Bump aws-ebs-csi-driver version to `v1.8.0`.



### coredns [1.11.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.11.0)

#### Changed
- Update `coredns` to upstream version [1.9.3](https://coredns.io/2022/05/27/coredns-1.9.3-release/).



### cluster-autoscaler [1.23.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.23.1)

#### Changed
- Update cluster-autoscaler to version `1.23.1`.
#### [1.22.2-gs7] - 2022-06-27
#### Changed
- Enable balance similar nodepools by default
#### Fixed
- Ignore labels to consider nodepools similar groups
#### Added
- Support to add extra arguments
#### [1.22.2-gs6] - 2022-04-07
#### [1.22.2-gs5] - 2022-04-06
#### Added
- Support cloud provider alias names (GCP -> GCE)
#### [1.22.2-gs4] - 2022-03-08
#### Fixed
- Updated to correct cluster-autoscaler version
- Use GS-built 1.22 image to deliver upstream unreleased fix https://github.com/kubernetes/autoscaler/pull/4600
#### [1.22.2-gs3] - 2022-02-07
#### Added
- Added support for specifying `balance-similar-node-groups` flag
#### [1.22.2-gs2] - 2022-01-14
#### Fixed
- Fix RBAC for version 1.22.
#### [1.22.2-gs1] - 2022-01-11
#### Changed
- Updated cluster-autoscaler to version `1.22.2`.
#### [1.21.2-gs1] - 2021-12-21
#### Changed
- Updated cluster-autoscaler to version `1.21.2`.
#### [1.21.0-gs2] - 2021-09-29
- Fix RBAC for cluster autoscaler 1.21.
#### [1.21.0-gs1] - 2021-09-13
#### Changed
- Updated cluster-autoscaler to version `1.21.0`.



### metrics-server [1.8.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.8.0)

#### Changed
- Updated metrics-server version to 0.6.1.



### vertical-pod-autoscaler [2.4.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.1)

#### Fixed
- Correct selector in admission controller PDB



