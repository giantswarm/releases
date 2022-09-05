# :zap: Giant Swarm Release v19.0.0 for Azure :zap:

<< Add description here >>

## Change details


### containerlinux [3227.2.2](https://www.flatcar-linux.org/releases/#release-3227.2.2)

_Note: The ARM64 AWS AMI of the Stable release has an unknown issue of corrupted images which we are still investigating. We will release the AMI as soon as we have resolved the issue. Follow [#840](https://github.com/flatcar-linux/Flatcar/issues/840) for more information_

_Changes since **Stable 3227.2.1**_

#### Security fixes:

- Linux ([CVE-2022-1679](https://nvd.nist.gov/vuln/detail/CVE-2022-1679), [CVE-2022-2585](https://nvd.nist.gov/vuln/detail/CVE-2022-2585), [CVE-2022-2586](https://nvd.nist.gov/vuln/detail/CVE-2022-2586), [CVE-2022-2588](https://nvd.nist.gov/vuln/detail/CVE-2022-2588), [CVE-2022-26373](https://nvd.nist.gov/vuln/detail/CVE-2022-26373), [CVE-2022-36946](https://nvd.nist.gov/vuln/detail/CVE-2022-36946))

#### Bug fixes:

- AWS: added EKS support for version 1.22 and 1.23. ([coreos-overlay#2110](https://github.com/flatcar-linux/coreos-overlay/pull/2110), [Flatcar#829](https://github.com/flatcar-linux/Flatcar/issues/829))
- VMWare: excluded `wireguard` (and others) from `systemd-networkd` management. ([init#80](https://github.com/flatcar-linux/init/pull/80))

#### Changes:

- The new image signing subkey was added to the public key embedded into `flatcar-install` (the old expired on 10th August 2022), only an updated `flatcar-install` script can verify releases signed with the new key ([init#79](https://github.com/flatcar-linux/init/pull/79))

#### Updates:

- Linux ([5.15.63](https://lwn.net/Articles/906061) (includes [5.15.62](https://lwn.net/Articles/905533), [5.15.61](https://lwn.net/Articles/904959), [5.15.60](https://lwn.net/Articles/904461), [5.15.59](https://lwn.net/Articles/903688))
- ca-certificates ([3.82](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_82.html))



### kubernetes [1.24.3](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.3)

#### Bug or Regression
- Fix a bug on endpointslices tests comparing the wrong metrics ([#110920](https://github.com/kubernetes/kubernetes/pull/110920), [@jluhrsen](https://github.com/jluhrsen)) [SIG Apps and Network]
- Fix a bug that caused the wrong result length when using --chunk-size and --selector together ([#110735](https://github.com/kubernetes/kubernetes/pull/110735), [@Abirdcfly](https://github.com/Abirdcfly)) [SIG API Machinery and Testing]
- Fix bug that prevented the job controller from enforcing activeDeadlineSeconds when set ([#110544](https://github.com/kubernetes/kubernetes/pull/110544), [@harshanarayana](https://github.com/harshanarayana)) [SIG Apps]
- Fix image pulling failure when IMDS is unavailable in kubelet startup ([#110523](https://github.com/kubernetes/kubernetes/pull/110523), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix printing resources with int64 fields ([#110572](https://github.com/kubernetes/kubernetes/pull/110572), [@sanchezl](https://github.com/sanchezl)) [SIG API Machinery]
- Fix unnecessary recreation of placeholder EndpointSlice ([#110732](https://github.com/kubernetes/kubernetes/pull/110732), [@jluhrsen](https://github.com/jluhrsen)) [SIG Apps and Network]
- Fixed a regression introduced in 1.24.0 where Azure load balancers were not kept up to date with the state of cluster nodes. In particular, nodes that are not in the ready state and are not newly created (i.e. not having the `node.cloudprovider.kubernetes.io/uninitialized` taint) now get removed from Azure load balancers. ([#109931](https://github.com/kubernetes/kubernetes/pull/109931), [@ricky-rav](https://github.com/ricky-rav)) [SIG Cloud Provider]
- Kubeadm: fix error adding extra prefix unix:// to CRI endpoints that were missing URL scheme ([#110634](https://github.com/kubernetes/kubernetes/pull/110634), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubeadm: fix the bug that configurable KubernetesVersion not respected during kubeadm join ([#111021](https://github.com/kubernetes/kubernetes/pull/111021), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



