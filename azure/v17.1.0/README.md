# :zap: Giant Swarm Release v17.1.0 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.18.1-dev](https://github.com/giantswarm/azure-operator/releases/tag/v5.18.1-dev)

Not found


### containerlinux [3033.2.4](https://www.flatcar-linux.org/releases/#release-3033.2.4)

Containerlinux release "3033.2.4" was not found in the changelog


### kubernetes [1.22.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.8)

#### API Change
- Fixes a regression in v1beta1 PodDisruptionBudget handling of "strategic merge patch"-type API requests for the `selector` field. Prior to 1.21, these requests would merge `matchLabels` content and replace `matchExpressions` content. In 1.21, patch requests touching the `selector` field started replacing the entire selector. This is consistent with server-side apply and the v1 PodDisruptionBudget behavior, but should not have been changed for v1beta1. ([#108141](https://github.com/kubernetes/kubernetes/pull/108141), [@liggitt](https://github.com/liggitt)) [SIG Auth and Testing]
#### Feature
- Kubernetes is now built with Golang 1.16.15 ([#108564](https://github.com/kubernetes/kubernetes/pull/108564), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Bump sigs.k8s.io/apiserver-network-proxy/konnectivity-client to v0.0.30, fixing goroutine leaks in kube-apiserver. ([#108439](https://github.com/kubernetes/kubernetes/pull/108439), [@andrewsykim](https://github.com/andrewsykim)) [SIG API Machinery, Auth and Cloud Provider]
- Fix static pod restarts in cases where the container is not present. ([#108189](https://github.com/kubernetes/kubernetes/pull/108189), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Fixes a bug where a partial EndpointSlice update could cause node name information to be dropped from endpoints that were not updated. ([#108202](https://github.com/kubernetes/kubernetes/pull/108202), [@robscott](https://github.com/robscott)) [SIG Network]
- Fixes a regression in the kubelet restarting static pods. ([#108303](https://github.com/kubernetes/kubernetes/pull/108303), [@rphillips](https://github.com/rphillips)) [SIG Node and Testing]
- Increase Azure ACR credential provider timeout ([#108209](https://github.com/kubernetes/kubernetes/pull/108209), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.27 → v0.0.30
#### Removed
_Nothing has changed._



### metrics-server [1.6.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.6.0)

#### Changed
- Updated metrics-server version to 0.5.2.



### net-exporter [1.12.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.12.0)

#### Changed
- Use parameter for CoreDNS namespace (defaulted to kube-system)



### node-exporter [1.10.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.10.0)

#### Changed
- Disable the fibrechannel collector.
- Disable the tapestats collector.



### azure-scheduled-events [0.7.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.7.0)

#### Added
- Add Vertical Pod Autoscaler CR.



### vertical-pod-autoscaler [2.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.1.2)

#### Fixed
- Fixed default value for admission controller PDB.



### vertical-pod-autoscaler-crd [1.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v1.0.1)

#### Added
- Add cluster singleton restriction so app can only be installed once.



