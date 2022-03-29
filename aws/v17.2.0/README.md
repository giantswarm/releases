# :zap: Giant Swarm Release v17.2.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [11.0.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.0.0)

#### Changed
- Update CAPI dependencies.
- Allow resource limits/requests to be passed as values.
- Switch `gp2` to `gp3` volumes.
- Allow etcd volume IOPS and Throughput to be set.



### cluster-operator [4.0.1](https://github.com/giantswarm/cluster-operator/releases/tag/v4.0.1)

#### Changed
- Update CAPI dependencies.

#### Fixed
- Only list apps from cluster namespace.



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



### containerlinux [3033.2.4](https://www.flatcar-linux.org/releases/#release-3033.2.4)

New **Stable** Release **3033.2.4**

**Changes since Stable-3033.2.3**

#### Security fixes
- Linux ([CVE-2022-25636](https://nvd.nist.gov/vuln/detail/CVE-2022-25636))
- Go ([CVE-2022-24921](https://nvd.nist.gov/vuln/detail/CVE-2022-24921))
- systemd ([CVE-2021-3997](https://nvd.nist.gov/vuln/detail/CVE-2021-3997))
- containerd ([CVE-2022-23648](https://nvd.nist.gov/vuln/detail/CVE-2022-23648))
- openssl ([CVE-2022-0778](https://nvd.nist.gov/vuln/detail/CVE-2022-0778))

#### Bug fixes
- Reverted the Linux kernel commit which broke networking on AWS instances which use Intel 82559 NIC (c4/m4) ([Flatcar#665](https://github.com/flatcar-linux/Flatcar/issues/665), [coreos-overlay#1720](https://github.com/flatcar-linux/coreos-overlay/pull/1720))

#### Changes
- Added support for switching back to CGroupsV1 without requiring a reboot. Create `/etc/flatcar-cgroupv1` through ignition. ([coreos-overlay#1666](https://github.com/flatcar-linux/coreos-overlay/pull/1666))

#### Updates
- Linux ([5.10.107](https://lwn.net/Articles/888522) (from 5.10.102, includes [5.10.103](https://lwn.net/Articles/886570), [5.10.104](https://lwn.net/Articles/887220), [5.10.105](https://lwn.net/Articles/887639), [5.10.106](https://lwn.net/Articles/888115)))
- Go ([1.17.8](https://go.googlesource.com/go/+/refs/tags/go1.17.8))
- systemd ([249.10](https://github.com/systemd/systemd-stable/releases/tag/v249.10))
- ca-certificates ([3.76](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_76.html))
- containerd ([1.5.10](https://github.com/containerd/containerd/releases/tag/v1.5.10))
- openssl ([1.1.1n](https://www.openssl.org/news/changelog.html#openssl-111))



### kiam [2.3.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.3.0)

#### Added
- Add VerticalPodAutoscaler CR.



### kiam-watchdog [0.6.0](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.6.0)

#### Added
- Add VerticalPodAutoscaler CR.



### aws-ebs-csi-driver [2.9.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.9.0)

#### Added
- Add VerticalPodAutoscaler CR.



