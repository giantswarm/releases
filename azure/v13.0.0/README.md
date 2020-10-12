# :zap: Giant Swarm Release v13.0.0 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.0.0-alpha1](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0-alpha1)

Not found


### containerlinux [2512.5.0](https://www.flatcar-linux.org/releases/#release-2512.5.0)

Changes:
- Update public key to include a [new subkey](https://www.flatcar-linux.org/security/image-signing-key/)

Updates:
- Linux [4.19.145](https://lwn.net/Articles/831367/)



### kubernetes [1.18.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.9)

#### Bug or Regression
- "unbound immediate PersistentVolumeClaims" causes UnschedulableAndUnresolvable status rather than an Error in the scheduler. ([#93892](https://github.com/kubernetes/kubernetes/pull/93892), [@ahg-g](https://github.com/ahg-g)) [SIG Apps and Storage]
- $ kubectl get event
  LAST SEEN   TYPE     REASON              OBJECT                        MESSAGE
  <unknown>   Normal   Scheduled           pod/nginx-6c975b59f8-gvmjr    Successfully assigned default/nginx-6c975b59f8-gvmjr to minikube
  
  $ kubectl describe pod xxx
  ......
  Events:
    Type    Reason     Age        From               Message
    ----    ------     ----       ----               -------
    Normal  Scheduled  <unknown>  default-scheduler  Successfully assigned default/nginx-6c975b59f8-gvmjr to minikube
  ...... ([#94226](https://github.com/kubernetes/kubernetes/pull/94226), [@ingvagabund](https://github.com/ingvagabund)) [SIG CLI]
- Azure: fix a bug that kube-controller-manager would panic if wrong Azure VMSS name is configured ([#94306](https://github.com/kubernetes/kubernetes/pull/94306), [@knight42](https://github.com/knight42)) [SIG Cloud Provider]
- Fix a concurrent map writes error in kubelet ([#93773](https://github.com/kubernetes/kubernetes/pull/93773), [@knight42](https://github.com/knight42)) [SIG Node]
- Fix calling AttachDisk on a previously attached EBS volume ([#93567](https://github.com/kubernetes/kubernetes/pull/93567), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider, Storage and Testing]
- Fix: incorrect max azure disk max count ([#92331](https://github.com/kubernetes/kubernetes/pull/92331), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed bug in reflector that couldn't recover from "Too large resource version" errors with API servers 1.17.0-1.18.5 ([#94316](https://github.com/kubernetes/kubernetes/pull/94316), [@janeczku](https://github.com/janeczku)) [SIG API Machinery]
- Fixed the EndpointSliceController to correctly create endpoints for IPv6-only pods.
  
  Fixed the EndpointController to allow IPv6 headless services, if the IPv6DualStack
  feature gate is enabled, by specifying `ipFamily: IPv6` on the service. (This already
  worked with the EndpointSliceController.) ([#91399](https://github.com/kubernetes/kubernetes/pull/91399), [@danwinship](https://github.com/danwinship)) [SIG Apps and Network]
- Fixes a bug evicting pods after a taint with a limited tolerationSeconds toleration is removed from a node ([#93722](https://github.com/kubernetes/kubernetes/pull/93722), [@liggitt](https://github.com/liggitt)) [SIG Apps and Node]
- Fixes an issue that can result in namespaced custom resources being orphaned when their namespace is deleted, if the CRD defining the custom resource is removed concurrently with namespaces being deleted, then recreated. ([#93790](https://github.com/kubernetes/kubernetes/pull/93790), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Fixing race condition with EndpointSlice controller garbage collection. ([#91311](https://github.com/kubernetes/kubernetes/pull/91311), [@robscott](https://github.com/robscott)) [SIG Apps, Network and Testing]
- If firstTimestamp is not set use eventTime when printing event ([#94252](https://github.com/kubernetes/kubernetes/pull/94252), [@ingvagabund](https://github.com/ingvagabund)) [SIG CLI]
- Kube-apiserver: fixed a bug returning inconsistent results from list requests which set a field or label selector and set a paging limit ([#94002](https://github.com/kubernetes/kubernetes/pull/94002), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- Pod Affinity/AntiAffinity label selectors are now validated in the pod affinity score plugin ([#93758](https://github.com/kubernetes/kubernetes/pull/93758), [@damemi](https://github.com/damemi)) [SIG Scheduling]
- Scheduler bugfix: Scheduler doesn't lose pod information when nodes are quickly recreated. This could happen when nodes are restarted or quickly recreated reusing a nodename. ([#93964](https://github.com/kubernetes/kubernetes/pull/93964), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling and Testing]
- The EndpointSlice controller now waits for EndpointSlice and Node caches to be synced before starting. ([#94086](https://github.com/kubernetes/kubernetes/pull/94086), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Upon successful authorization check, an impersonated user is added to the system:authenticated group.
  system:anonymous when impersonated is added to the system:unauthenticated group. ([#94409](https://github.com/kubernetes/kubernetes/pull/94409), [@tkashem](https://github.com/tkashem)) [SIG API Machinery and Testing]
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



