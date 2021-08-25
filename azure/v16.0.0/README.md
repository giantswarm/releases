# :zap: Giant Swarm Release v16.0.0 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.8.2-dev](https://github.com/giantswarm/azure-operator/releases/tag/v5.8.2-dev)

Not found


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



### containerlinux [2905.2.2](https://www.flatcar-linux.org/releases/#release-2905.2.2)

_Changes since **Stable**_ **2905.2.1**

**Security fixes**



* Linux ([CVE-2021-34556](https://nvd.nist.gov/vuln/detail/CVE-2021-34556), [CVE-2021-35477](https://nvd.nist.gov/vuln/detail/CVE-2021-35477), [CVE-2021-38205](https://nvd.nist.gov/vuln/detail/CVE-2021-38205))
* Go ([CVE-2021-36221](https://nvd.nist.gov/vuln/detail/CVE-2021-36221))
* Systemd ([CVE-2020-13529](https://nvd.nist.gov/vuln/detail/CVE-2020-13529), [CVE-2021-33910](https://nvd.nist.gov/vuln/detail/CVE-2021-33910))

**Bug Fixes**



* Fixed `pam.d` sssd LDAP auth with sudo ([coreos-overlay#1170](https://github.com/kinvolk/coreos-overlay/pull/1170))
* Let network-cleanup.service finish before entering rootfs ([coreos-overlay#1182](https://github.com/kinvolk/coreos-overlay/pull/1182))

**Changes**



* Switched to zstd for the initramfs ([coreos-overlay#1136](https://github.com/kinvolk/coreos-overlay/pull/1136))
* Embedded new subkey in flatcar-install ([coreos-overlay#1180](https://github.com/kinvolk/coreos-overlay/pull/1180))

**Updates**



* Linux ([5.10.59](https://lwn.net/Articles/866302/))
* Systemd ([247.9](https://github.com/systemd/systemd-stable/releases/tag/v247.9))
* Go ([1.16.7](https://golang.org/doc/devel/release#go1.16.minor))
* portage-utils ([0.90](https://github.com/gentoo/portage-utils/releases/tag/v0.90))


### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

#### Added
- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.
#### Changed
- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.



### external-dns [2.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.5.0)

#### Changed
- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.



