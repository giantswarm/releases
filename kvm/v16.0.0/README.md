# :zap: Giant Swarm Release v16.0.0 for KVM :zap:

<< Add description here >>

## Change details


### kubernetes [1.21.4](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.4)

#### Feature
- Kubernetes is now built with Golang 1.16.7 ([#104201](https://github.com/kubernetes/kubernetes/pull/104201), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Disable aufs module for gce clusters ([#103831](https://github.com/kubernetes/kubernetes/pull/103831), [@lizhuqi](https://github.com/lizhuqi)) [SIG Cloud Provider]
- Fix kube-apiserver metric reporting for the deprecated watch path of /api/<version>/watch/... ([#104190](https://github.com/kubernetes/kubernetes/pull/104190), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- Fix the code is leaking the defaulting between unrelated pod instances. ([#103284](https://github.com/kubernetes/kubernetes/pull/103284), [@kebe7jun](https://github.com/kebe7jun)) [SIG CLI]
- Fix: Provide IPv6 support for internal load balancer ([#103794](https://github.com/kubernetes/kubernetes/pull/103794), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: cleanup outdated routes ([#102935](https://github.com/kubernetes/kubernetes/pull/102935), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: ignore not a VMSS error for VMAS nodes in reconcileBackendPools ([#103997](https://github.com/kubernetes/kubernetes/pull/103997), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: return empty VMAS name if using standalone VM ([#103470](https://github.com/kubernetes/kubernetes/pull/103470), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed a bug that scheduler extenders are not called on preemptions ([#103019](https://github.com/kubernetes/kubernetes/pull/103019), [@ordovicia](https://github.com/ordovicia)) [SIG Scheduling]
- Fixes an issue cleaning up CertificateSigningRequest objects with an unparseable `status.certificate` field ([#103948](https://github.com/kubernetes/kubernetes/pull/103948), [@liggitt](https://github.com/liggitt)) [SIG Apps and Auth]
- Fixes issue with websocket-based watches of Service objects not closing correctly on timeout ([#102541](https://github.com/kubernetes/kubernetes/pull/102541), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.19 → v0.0.22
#### Removed
_Nothing has changed._



