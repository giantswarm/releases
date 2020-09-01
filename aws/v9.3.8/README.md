# :zap: Giant Swarm Release v9.3.8 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [5.7.6](https://github.com/giantswarm/aws-operator/releases/tag/v5.7.6)

#### Changed
- No notable changes.



### kubernetes [1.16.14](https://github.com/kubernetes/kubernetes/releases/tag/v1.16.14)

#### Bug or Regression
- Do not add nodes labeled with kubernetes.azure.com/managed=false to backend pool of load balancer. ([#93034](https://github.com/kubernetes/kubernetes/pull/93034), [@matthias50](https://github.com/matthias50)) [SIG Cloud Provider]
- Fix instance not found issues when an Azure Node is recreated in a short time ([#93316](https://github.com/kubernetes/kubernetes/pull/93316), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix: don't use docker config cache if it's empty ([#92330](https://github.com/kubernetes/kubernetes/pull/92330), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: initial delay in mounting azure disk & file ([#93052](https://github.com/kubernetes/kubernetes/pull/93052), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a performance issue applying json patches to deeply nested objects ([#93813](https://github.com/kubernetes/kubernetes/pull/93813), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, CLI, Cloud Provider and Cluster Lifecycle]
- Fixed a regression in kubeadm manifests for kube-scheduler and kube-controller-manager which caused continuous restarts because of failing health checks ([#93208](https://github.com/kubernetes/kubernetes/pull/93208), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Fixes a regression in kube-apiserver causing 500 errors from the `/readyz` endpoint ([#93645](https://github.com/kubernetes/kubernetes/pull/93645), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG API Machinery]
#### Other (Cleanup or Flake)
- Build: Update Debian base images
  - debian-base:v2.1.3
  - debian-iptables:v12.1.2
  - debian-hyperkube-base:v1.1.3 ([#93927](https://github.com/kubernetes/kubernetes/pull/93927), [@justaugustus](https://github.com/justaugustus)) [SIG API Machinery, Cluster Lifecycle and Release]
- Kubernetes is now built with go1.13.15 ([#93956](https://github.com/kubernetes/kubernetes/pull/93956), [@justaugustus](https://github.com/justaugustus)) [SIG Release and Testing]
- Update Golang to v1.13.14
  - Update bazel to 2.2.0
  - Update repo-infra to 0.0.8 (to support go1.14.6 and go1.13.14)
    - Includes:
      - bazelbuild/bazel-toolchains@3.4.0
      - bazelbuild/rules_go@v0.22.8 ([#93235](https://github.com/kubernetes/kubernetes/pull/93235), [@justaugustus](https://github.com/justaugustus)) [SIG API Machinery, Release and Testing]
#### Dependencies
#### Added
- github.com/jessevdk/go-flags: [v1.4.0](https://github.com/jessevdk/go-flags/tree/v1.4.0)
#### Changed
- github.com/evanphx/json-patch: [v4.2.0+incompatible → 162e562](https://github.com/evanphx/json-patch/compare/v4.2.0...162e562)
#### Removed
_Nothing has changed._



### calico [3.10.4](https://github.com/projectcalico/calico/releases/tag/v3.10.4)

#### Other changes
- Use mv to install CNI binaries instead of cp cni-plugin #852 (@caseydavenport)



### etcd [3.3.25](https://github.com/etcd-io/etcd/releases/tag/v3.3.25)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.3.23...v3.3.25) and [v3.3 upgrade guide](https://github.com/etcd-io/etcd/blob/master/Documentation/upgrades/upgrade_3_3.md) for any breaking changes.
#### Security
- A [log warning](https://github.com/etcd-io/etcd/pull/12242) is added when etcd use any existing directory that has a permission different than 700 on Linux and 777 on Windows.



### containerlinux [2512.3.0](https://www.flatcar-linux.org/releases/#release-2512.3.0)

Security fixes:

* Bind: fixes for [CVE-2020-8616](https://nvd.nist.gov/vuln/detail/CVE-2020-8616), [CVE-2020-8617](https://nvd.nist.gov/vuln/detail/CVE-2020-8617), [CVE-2020-8620](https://nvd.nist.gov/vuln/detail/CVE-2020-8620), [CVE-2020-8621](https://nvd.nist.gov/vuln/detail/CVE-2020-8621), [CVE-2020-8622](https://nvd.nist.gov/vuln/detail/CVE-2020-8622), [CVE-2020-8623](https://nvd.nist.gov/vuln/detail/CVE-2020-8623), [CVE-2020-8624](https://nvd.nist.gov/vuln/detail/CVE-2020-8624)

Bug fixes:

* The static IP address configuration in the initramfs works again in the format `ip=<ip>::<gateway>:<netmask>:<hostname>:<iface>:none[:<dns1>[:<dns2>]]` ([flatcar-linux/bootengine#15](https://github.com/flatcar-linux/bootengine/pull/15))
* app-admin/{kubelet, etcd, flannel}-wrapper: don't overwrite the user supplied –insecure-options argument ([flatcar-linux/coreos-overlay#426](https://github.com/flatcar-linux/coreos-overlay/pull/426))
* etcd-wrapper: Adjust data dir permissions ([flatcar-linux/coreos-overlay#536](https://github.com/flatcar-linux/coreos-overlay/pull/536))

Changes:

* Vultr support in Ignition ([flatcar-linux/ignition#13](https://github.com/flatcar-linux/ignition/pull/13))
* VMware OVF settings default to ESXi 6.5 and Linux 3.x

Updates:

* Linux [4.19.140](https://lwn.net/Articles/829107/)
* bind-tools [9.11.22](https://ftp.isc.org/isc/bind9/cur/9.11/RELEASE-NOTES-bind-9.11.22.txt)
* etcd-wrapper [3.3.24](https://github.com/etcd-io/etcd/releases/tag/v3.3.24)
* Git [2.26.2](https://raw.githubusercontent.com/git/git/v2.26.2/Documentation/RelNotes/2.26.2.txt)


