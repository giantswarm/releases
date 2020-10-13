# :zap: Giant Swarm Release v13.0.0-alpha2 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.17.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.12)

#### Bug or Regression
- Azure: fix a bug that kube-controller-manager would panic if wrong Azure VMSS name is configured ([#94306](https://github.com/kubernetes/kubernetes/pull/94306), [@knight42](https://github.com/knight42)) [SIG Cloud Provider]
- Correctly handle resetting cpuacct in a live container ([#94041](https://github.com/kubernetes/kubernetes/pull/94041), [@andyzhangx](https://github.com/andyzhangx)) [SIG Node and Windows]
- Fix a concurrent map writes error in kubelet ([#93773](https://github.com/kubernetes/kubernetes/pull/93773), [@knight42](https://github.com/knight42)) [SIG Node]
- Fix calling AttachDisk on a previously attached EBS volume ([#93567](https://github.com/kubernetes/kubernetes/pull/93567), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider, Storage and Testing]
- Fix: incorrect max azure disk max count ([#92331](https://github.com/kubernetes/kubernetes/pull/92331), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixes a bug evicting pods after a taint with a limited tolerationSeconds toleration is removed from a node ([#93722](https://github.com/kubernetes/kubernetes/pull/93722), [@liggitt](https://github.com/liggitt)) [SIG Apps and Node]
- Fixes an issue that can result in namespaced custom resources being orphaned when their namespace is deleted, if the CRD defining the custom resource is removed concurrently with namespaces being deleted, then recreated. ([#93790](https://github.com/kubernetes/kubernetes/pull/93790), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Kube-apiserver: fixed a bug returning inconsistent results from list requests which set a field or label selector and set a paging limit ([#94002](https://github.com/kubernetes/kubernetes/pull/94002), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- The EndpointSlice controller now waits for EndpointSlice and Node caches to be synced before starting. ([#94086](https://github.com/kubernetes/kubernetes/pull/94086), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- The audit event sourceIPs list will now always end with the IP that sent the request directly to the API server. ([#87167](https://github.com/kubernetes/kubernetes/pull/87167), [@tallclair](https://github.com/tallclair)) [SIG API Machinery and Auth]
- Upon successful authorization check, an impersonated user is added to the system:authenticated group.
  system:anonymous when impersonated is added to the system:unauthenticated group. ([#94410](https://github.com/kubernetes/kubernetes/pull/94410), [@tkashem](https://github.com/tkashem)) [SIG API Machinery and Testing]
- Use NLB Subnet CIDRs instead of VPC CIDRs in Health Check SG Rules ([#93515](https://github.com/kubernetes/kubernetes/pull/93515), [@t0rr3sp3dr0](https://github.com/t0rr3sp3dr0)) [SIG Cloud Provider]
#### Other (Cleanup or Flake)
- Fixes the flooding warning messages about setting volume ownership for configmap/secret volumes ([#92878](https://github.com/kubernetes/kubernetes/pull/92878), [@jvanz](https://github.com/jvanz)) [SIG Instrumentation, Node and Storage]
- Update CNI plugins to v0.8.7 ([#94367](https://github.com/kubernetes/kubernetes/pull/94367), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Network, Node, Release and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/evanphx/json-patch: [162e562 → v4.9.0+incompatible](https://github.com/evanphx/json-patch/compare/162e562...v4.9.0)
#### Removed
- github.com/jessevdk/go-flags: [v1.4.0](https://github.com/jessevdk/go-flags/tree/v1.4.0)



### azure-operator [5.0.0-alpha2](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0-alpha2)

Not found


