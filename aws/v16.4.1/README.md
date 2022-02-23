# :zap: Giant Swarm Release v16.4.1 for AWS :zap:

This release allows one replica of `coredns` to run on the control plane nodes for clusters without any node pools.

## Change details


### kubernetes [1.21.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.9)

#### Feature
- Kube-apiserver: when merging lists, Server Side Apply now prefers the order of the submitted request instead of the existing persisted object ([#107569](https://github.com/kubernetes/kubernetes/pull/107569), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
#### Bug or Regression
- An inefficient lock in EndpointSlice controller metrics cache has been reworked. Network programming latency may be significantly reduced in certain scenarios, especially in clusters with a large number of Services. ([#107169](https://github.com/kubernetes/kubernetes/pull/107169), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Client-go: fix that paged list calls with ResourceVersionMatch set would fail once paging kicked in. ([#107336](https://github.com/kubernetes/kubernetes/pull/107336), [@fasaxc](https://github.com/fasaxc)) [SIG API Machinery]
- Fix a panic when using invalid output format in kubectl create secret command ([#107345](https://github.com/kubernetes/kubernetes/pull/107345), [@rikatz](https://github.com/rikatz)) [SIG CLI]
- Fixes a rare race condition handling requests that timeout ([#107460](https://github.com/kubernetes/kubernetes/pull/107460), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Mount-utils: Detect potential stale file handle ([#107040](https://github.com/kubernetes/kubernetes/pull/107040), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage]
#### Other (Cleanup or Flake)
- Updates konnectivity-network-proxy to v0.0.27. This includes a memory leak fix for the network proxy ([#107188](https://github.com/kubernetes/kubernetes/pull/107188), [@rata](https://github.com/rata)) [SIG API Machinery and Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.39.0 → v0.39.3](https://github.com/google/cadvisor/compare/v0.39.0...v0.39.3)
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.22 → v0.0.27
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.1
#### Removed
_Nothing has changed._



### coredns [1.8.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.8.0)

#### Changed
- Add deployment to run one replica of coredns in master nodes (for clusters with no node pools).



