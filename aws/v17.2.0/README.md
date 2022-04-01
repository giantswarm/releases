# :zap: Giant Swarm Release v17.2.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [11.2.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.2.0)

#### Added
 - Add annotation to ASG to make cluster-autoscaler work when scaling from zero replicas.

#### Changed
- Bumped k8scc to 13.2.0 to enable VPA for kube-proxy.
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
- Fix Azurefile volumeid collision issue in csi migration ([#107575](https://github.com/kubernetes/kubernetes/pull/107575), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix: delete non existing Azure disk issue ([#107406](https://github.com/kubernetes/kubernetes/pull/107406), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: ignore the case when comparing azure tags in service annotation (azure) ([#107580](https://github.com/kubernetes/kubernetes/pull/107580), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]

#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.27 → v0.0.30
- k8s.io/utils: bdf08cb → 6203023
#### Removed
_Nothing has changed._



### containerlinux [3033.2.4](https://www.flatcar-linux.org/releases/#release-3033.2.4)

New **Stable** Release **3033.2.4**

**Changes since Stable-3033.2.3**

#### Security fixes
- Linux
  - [CVE-2022-25636](https://nvd.nist.gov/vuln/detail/CVE-2022-25636)
  - [CVE-2022-24448](https://nvd.nist.gov/vuln/detail/CVE-2022-24448)
  - [CVE-2022-0617](https://nvd.nist.gov/vuln/detail/CVE-2022-0617)
  - [CVE-2022-24959](https://nvd.nist.gov/vuln/detail/CVE-2022-24959)
  - [CVE-2022-0492](https://nvd.nist.gov/vuln/detail/CVE-2022-0492)
  - [CVE-2022-0516](https://nvd.nist.gov/vuln/detail/CVE-2022-0516)
  - [CVE-2022-0435](https://nvd.nist.gov/vuln/detail/CVE-2022-0435)
  - [CVE-2022-0487](https://nvd.nist.gov/vuln/detail/CVE-2022-0487)
  - [CVE-2022-25375](https://nvd.nist.gov/vuln/detail/CVE-2022-25375)
  - [CVE-2022-25258](https://nvd.nist.gov/vuln/detail/CVE-2022-25258)
  - [CVE-2022-0847](https://nvd.nist.gov/vuln/detail/CVE-2022-0847)
- Go
  - [CVE-2022-24921](https://nvd.nist.gov/vuln/detail/CVE-2022-24921)
  - [CVE-2022-23806](https://nvd.nist.gov/vuln/detail/CVE-2022-23806)
  - [CVE-2022-23772](https://nvd.nist.gov/vuln/detail/CVE-2022-23772)
  - [CVE-2022-23773](https://nvd.nist.gov/vuln/detail/CVE-2022-23773)
- systemd ([CVE-2021-3997](https://nvd.nist.gov/vuln/detail/CVE-2021-3997))
- containerd ([CVE-2022-23648](https://nvd.nist.gov/vuln/detail/CVE-2022-23648))
- openssl ([CVE-2022-0778](https://nvd.nist.gov/vuln/detail/CVE-2022-0778))
- ignition
  - [CVE-2020-14040](https://nvd.nist.gov/vuln/detail/CVE-2020-14040)
- expat
  - [CVE-2022-25235](https://nvd.nist.gov/vuln/detail/CVE-2022-25235)
  - [CVE-2022-25236](https://nvd.nist.gov/vuln/detail/CVE-2022-25236)
  - [CVE-2022-25313](https://nvd.nist.gov/vuln/detail/CVE-2022-25313)
  - [CVE-2022-25314](https://nvd.nist.gov/vuln/detail/CVE-2022-25314)
  - [CVE-2022-25315](https://nvd.nist.gov/vuln/detail/CVE-2022-25315)

#### Bug fixes
- Reverted the Linux kernel commit which broke networking on AWS instances which use Intel 82559 NIC (c4/m4) ([Flatcar#665](https://github.com/flatcar-linux/Flatcar/issues/665), [coreos-overlay#1720](https://github.com/flatcar-linux/coreos-overlay/pull/1720))
- Disabled the systemd-networkd settings ManageForeignRoutes and ManageForeignRoutingPolicyRules by default to ensure that CNIs like Cilium don’t get their routes or routing policy rules discarded on network reconfiguration events ([Flatcar#620](https://github.com/flatcar-linux/Flatcar/issues/620)).
- Prevented hitting races when creating filesystems in Ignition, these races caused boot failures like fsck[1343]: Failed to stat /dev/disk/by-label/ROOT: No such file or directory when creating a btrfs root filesystem ([ignition#35](https://github.com/flatcar-linux/ignition/pull/35))
- Reverted the Linux kernel change to forbid xfrm id 0 for IPSec state because it broke Cilium ([Flatcar#626](https://github.com/flatcar-linux/Flatcar/issues/626), [coreos-overlay#1682](https://github.com/flatcar-linux/coreos-overlay/pull/1682))

#### Changes
- Added support for switching back to CGroupsV1 without requiring a reboot. Create `/etc/flatcar-cgroupv1` through ignition. ([coreos-overlay#1666](https://github.com/flatcar-linux/coreos-overlay/pull/1666))

#### Updates
- Linux ([5.10.107](https://lwn.net/Articles/888522) (from 5.10.102, includes [5.10.103](https://lwn.net/Articles/886570), [5.10.104](https://lwn.net/Articles/887220), [5.10.105](https://lwn.net/Articles/887639), [5.10.106](https://lwn.net/Articles/888115)))
- Go ([1.17.8](https://go.googlesource.com/go/+/refs/tags/go1.17.8))
- systemd ([249.10](https://github.com/systemd/systemd-stable/releases/tag/v249.10))
- ca-certificates ([3.76](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_76.html))
- containerd ([1.5.10](https://github.com/containerd/containerd/releases/tag/v1.5.10))
- openssl ([1.1.1n](https://www.openssl.org/news/changelog.html#openssl-111))
- expat ([2.4.6](https://github.com/libexpat/libexpat/blob/R_2_4_6/expat/Changes))



### kiam [2.3.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.3.0)

#### Added
- Add VerticalPodAutoscaler CR.



### kiam-watchdog [0.6.0](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.6.0)

#### Added
- Add VerticalPodAutoscaler CR.



### aws-ebs-csi-driver [2.9.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.9.0)

#### Added
- Add VerticalPodAutoscaler CR.



### chart-operator [2.20.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.20.1)

#### Changed
- Use `apptestctl` to install CRDs in integration tests to avoid hitting GitHub rate limits.

#### Fixed
- Fix `status` resource to use Helm release status if it exists.



### cert-operator [2.0.0](https://github.com/giantswarm/cert-operator/releases/tag/v2.0.0)

#### Changed
- Use v1beta1 CAPI CRDs.
- Bump `giantswarm/apiextensions` to `v6.0.0`.
- Bump `giantswarm/exporterkit` to `v1.0.0`.
- Bump `giantswarm/microendpoint` to `v1.0.0`.
- Bump `giantswarm/microerror` to `v0.4.0`.
- Bump `giantswarm/microkit` to `v1.0.0`.
- Bump `giantswarm/micrologger` to `v0.6.0`.
- Bump `giantswarm/k8sclient` to `v7.0.1`.
- Bump `giantswarm/operatorkit` to `v7.0.1`.
- Bump k8s dependencies to `v0.22.2`.
- Bump `controller-runtime` to `v0.10.3`.
- Use `apptestctl` to install CRDs in integration tests to avoid hitting GitHub rate limits.


