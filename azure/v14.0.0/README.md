# :zap: Giant Swarm Release v14.0.0 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.19.7](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.7)

TODO

### app-operator [3.0.0](https://github.com/giantswarm/app-operator/releases/tag/v3.0.0)

#### Changed
- Enable mutating and validating webhooks in app-admission-controller for
tenant app CRs.
#### Added
- Make resync period configurable for use in integration tests.
- Pause App CR reconciliation when it has
  `app-operator.giantswarm.io/paused=true` annotation.
- Print difference between the current chart and desired chart.



### cluster-operator [0.23.20](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.20)

Not found


### azure-operator [5.2.1-dev](https://github.com/giantswarm/aws-operator/releases/tag/v5.2.1-dev)

Not found


### containerlinux [2605.11.0](https://www.flatcar-linux.org/releases/#release-2605.11.0)

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



### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

Not found


### cluster-autoscaler [1.19.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.1)

#### Changed
- Updated cluster-autoscaler to version `1.19.1`.



