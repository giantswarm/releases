# :zap: Giant Swarm Release v14.0.0 for Azure :zap:

This is the first workload cluster release to support Kubernetes 1.19 on Azure.

With recent [added support by `Azure`](https://azure.microsoft.com/en-us/updates/germany-west-central-availability-zones-now-generally-available/),
with this release `Availability Zones` are supported for all workload clusters running in the `Germany West Central`
region as well.

Starting from this release, Azure workload clusters include by default a new application named `azure-scheduled-events`
that leverages the [Azure scheduled events](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/scheduled-events)
feature to automatically drain a `Kubernetes` node when the underlying virtual instance is about to be terminated.
This ensures the workload running on the node is handled gracefully when the `cluster autoscaler` scales down a node pool.

Please note that with version `1.19` there a few breaking changes in the `Kubernetes` APIs.
Please refer to the [upstream documentation](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#urgent-upgrade-notes)
and feel free to get in touch with your solutions engineer for any concern you might have.

## Change details

### kubernetes [1.19.7](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.7)

Please check the official changelog for what's new in [`kubernetes 1.19`](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#changelog-since-v1180).

### app-operator [3.0.0](https://github.com/giantswarm/app-operator/releases/tag/v3.0.0)

#### Changed
- Enable mutating and validating webhooks in app-admission-controller for
tenant app CRs.
#### Added
- Make resync period configurable for use in integration tests.
- Pause App CR reconciliation when it has
  `app-operator.giantswarm.io/paused=true` annotation.
- Print difference between the current chart and desired chart.

### cluster-operator [0.23.22](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.20)

#### Added
- Check existence of chart tarball for release CR apps in catalog.

### azure-operator [5.3.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.3.0)

#### Added

- Enable Azure termination events for Node Pool VMSSes.
- Enable availability zones on `Germany West Central`.

### containerlinux [2605.11.0](https://www.flatcar-linux.org/releases/#release-2605.11.0)

Security fixes:

* Linux [CVE-2020-27815](https://www.openwall.com/lists/oss-security/2020/11/30/5)
* Linux [CVE-2020-29568](https://nvd.nist.gov/vuln/detail/CVE-2020-29568)
* Linux [CVE-2020-29569](https://nvd.nist.gov/vuln/detail/CVE-2020-29569)
* Linux [CVE-2020-29661](https://nvd.nist.gov/vuln/detail/CVE-2020-29661)
* Linux [CVE-2020-29660](https://nvd.nist.gov/vuln/detail/CVE-2020-29660)
* Linux [CVE-2020-27830](https://nvd.nist.gov/vuln/detail/CVE-2020-27830)
* Linux [CVE-2020-28588](https://nvd.nist.gov/vuln/detail/CVE-2020-28588)

Bug fixes:

*   networkd: avoid managing MAC addresses for veth devices ([kinvolk/init#33](https://github.com/kinvolk/init/pull/33))
*   The sysctl `net.ipv4.conf.*.rp_filter` is set to `0` for the Cilium CNI plugin to work ([Flatcar#181](https://github.com/kinvolk/Flatcar/issues/181))
*   Package downloads in the developer container now use the correct URL again ([Flatcar#298](https://github.com/kinvolk/Flatcar/issues/298))

Changes:

*   The sysctl default config file is now applied under the prefix 60 which allows for custom sysctl config files to take effect when they start with a prefix of 70, 80, or 90 ([baselayout#13](https://github.com/kinvolk/baselayout/pull/13))
*   Containerd CRI plugin got enabled by default, only the containerd socket path needs to be specified as kubelet parameter for Kubernetes 1.20 to use containerd instead of Docker ([Flatcar#283](https://github.com/kinvolk/Flatcar/issues/283))
*   For users with a custom update server a machine alias setting in update-engine allows to give human-friendly names to client instances ([update-engine#8](https://github.com/kinvolk/update_engine/pull/8))

Updates:

*   Linux ([5.4.87](https://lwn.net/Articles/841900/))

### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

Please check the [official changelog](https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.4.md#v3414-2020-11-25)
for the list of changes included in this release.

### cluster-autoscaler [1.19.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.1)

Updated to the `1.19.1` version that is the suggested upstream version to run for `1.19` `Kubernetes` clusters.

### cert-exporter [1.6.0](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#160---2021-01-27)

#### Added
- Add exceptions in NetworkPolicies to allow DNS to work correctly through port 53.

### external-dns [1.6.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#160---2021-01-27)

#### Changed
- Upgrade upstream external-dns from v0.7.4 to v0.7.6.
