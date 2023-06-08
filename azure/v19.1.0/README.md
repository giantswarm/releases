# :zap: Giant Swarm Release v19.1.0 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.24.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.11)

#### Feature
- Kubelet TCP and HTTP probes are more effective using networking resources: conntrack entries, sockets, ... 
  This is achieved by reducing the TIME-WAIT state of the connection to 1 second, instead of the defaults 60 seconds. This allows kubelet to free the socket, and free conntrack entry and ephemeral port associated. ([#115143](https://github.com/kubernetes/kubernetes/pull/115143), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Kubernetes is now built with Go 1.19.6 ([#115831](https://github.com/kubernetes/kubernetes/pull/115831), [@cpanato](https://github.com/cpanato)) [SIG Release and Testing]
#### Bug or Regression
- Fix the regression that introduced 34s timeout for DELETECOLLECTION calls ([#115482](https://github.com/kubernetes/kubernetes/pull/115482), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fixed bug which caused the status of Indexed Jobs to only be updated when there are newly completed indexes. The completed indexes are now updated if the .status.completedIndexes has values outside of the [0, .spec.completions> range ([#115457](https://github.com/kubernetes/kubernetes/pull/115457), [@danielvegamyhre](https://github.com/danielvegamyhre)) [SIG Apps]
- Golang.org/x/net updates to v0.7.0 to fix CVE-2022-41723 ([#115789](https://github.com/kubernetes/kubernetes/pull/115789), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Security and Storage]
- The Kubernetes API server now correctly detects and closes existing TLS connections when its client certificate file for kubelet authentication has been rotated. ([#115580](https://github.com/kubernetes/kubernetes/pull/115580), [@enj](https://github.com/enj)) [SIG API Machinery, Node and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- golang.org/x/net: 1e63c2f → v0.7.0
- golang.org/x/sys: v0.3.0 → v0.5.0
- golang.org/x/term: v0.3.0 → v0.5.0
- golang.org/x/text: v0.5.0 → v0.7.0
#### Removed
_Nothing has changed._



### chart-operator [2.35.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.35.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



