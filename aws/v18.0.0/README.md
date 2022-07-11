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



