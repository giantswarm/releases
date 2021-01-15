# :zap: Giant Swarm Release v13.1.0 for Azure :zap:

This is the next minor workload cluster release for Node Pools Clusters.

Release brings two new features that can be enabled by customers:

- Automatic Termination Unhealthy Nodes feature enables customers to discard unhealthy kubernetes nodes in favor of new ones added to the cluster. Follow [this guide](https://docs.giantswarm.io/basics/automatic-termination-of-bad-nodes/) to learn how to enable and disable this feature via annotations on Cluster CR.

- [Node Pools Auto Scaling](https://docs.giantswarm.io/basics/nodepools/#autoscaling) feature allows customers to set up minimum and maximum number of nodes per node pool used by autoscaler app installed by default.


## Change details


### azure-operator [5.2.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.2.0)

Main changes since previous release:

#### Fixed
- Ensure that availability zones are kept unchanged during migration from 12.x to 13.x.

#### Added
- Add `terminate-unhealthy-node` feature to automatically terminate bad and unhealthy nodes in a Cluster, disabled by default.
- Ensure autoscaler annotations during migration from 12.x to 13.x.
- Add Cluster `ControlPlaneReady` condition to comply with upstream Cluster API implementation.
- Add AzureMachine `Ready`, `SubnetReady` and `VMSSReady` conditions to comply with upstream Cluster API implementation.
- Add MachinePool `Creating` condition to comply with upstream Cluster API implementation.

For a detailed list of all changes, please refer to the [changelog](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md).

### kubernetes [1.18.15](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.15)

Upgraded from `1.18.12`, please refer to the upstream changelog for a list of changes:

- Changelog since `1.18.14`: [link](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#changelog-since-v11814)
- Changelog since `1.18.13`: [link](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#changelog-since-v11813) 
- Changelog since `1.18.12`: [link](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#changelog-since-v11812)


### containerlinux [2605.10.0](https://www.flatcar-linux.org/releases/#release-2605.10.0)

Security fixes:

*   Linux [CVE-2020-29661](https://nvd.nist.gov/vuln/detail/CVE-2020-29661), [CVE-2020-29660](https://nvd.nist.gov/vuln/detail/CVE-2020-29660), [CVE-2020-27830](https://nvd.nist.gov/vuln/detail/CVE-2020-27830), [CVE-2020-28588](https://nvd.nist.gov/vuln/detail/CVE-2020-28588)

Bug fixes:

*   The sysctl `net.ipv4.conf.*.rp_filter` is set to `0` for the Cilium CNI plugin to work ([Flatcar#181](https://github.com/kinvolk/Flatcar/issues/181))
*   Package downloads in the developer container now use the correct URL again ([Flatcar#298](https://github.com/kinvolk/Flatcar/issues/298))

Changes:

*   The sysctl default config file is now applied under the prefix 60 which allows for custom sysctl config files to take effect when they start with a prefix of 70, 80, or 90 ([baselayout#13](https://github.com/kinvolk/baselayout/pull/13))
*   Containerd CRI plugin got enabled by default, only the containerd socket path needs to be specified as kubelet parameter for Kubernetes 1.20 to use containerd instead of Docker ([Flatcar#283](https://github.com/kinvolk/Flatcar/issues/283))
*   For users with a custom update server a machine alias setting in update-engine allows to give human-friendly names to client instances ([update-engine#8](https://github.com/kinvolk/update_engine/pull/8))

Updates:

*   Linux ([5.4.83](https://lwn.net/Articles/839875/))


Please analyse the release notes of all other updated components below:

### app-operator [3.0.0](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#300---2021-01-05)
### cert-exporter [1.5.0](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#150---2021-01-05)
### chart-operator [2.6.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#260---2020-12-21)
### cluster-autoscaler [1.18.3](https://github.com/giantswarm/cluster-autoscaler-app/blob/release-v1.18.x/CHANGELOG.md#1183---2020-11-03)
### metrics-server [1.2.1](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#121---2020-12-10)
### node-exporter [1.7.1](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#171---2020-12-11)
