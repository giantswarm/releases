# :zap: Giant Swarm Release v13.1.0 for Azure :zap:

This is the next minor workload cluster release for Node Pools Clusters.

Release brings two new features that can be enabled by customers:

- Automatic Termination Unhealthy Nodes feature enables customers to discard unhealthy kubernetes nodes in favor of new ones added to the cluster. Follow [this guide](https://docs.giantswarm.io/basics/automatic-termination-of-bad-nodes/) to learn how to enable and disable this feature via annotations on Cluster CR.

- [Node Pools Auto Scaling](https://docs.giantswarm.io/basics/nodepools/#autoscaling) feature allows customers to set up minimum and maximum number of nodes per node pool used by autoscaler app installed by default.


## Change details


### azure-operator [5.1.0](https://github.com/giantswarm/aws-operator/releases/tag/v5.1.0)

#### Changed
- Replaced Cluster `ProviderInfrastructureReady` with upstream `InfrastructureReady` condition.
#### Fixed
- Fix incorrect (too early) `Upgrading` condition transition from `True` to `False`.
#### Added
- Add Cluster autoscaler-related tags to Node Pools VMSS.
- Add `terminate-unhealthy-node` feature to automatically terminate bad and unhealthy nodes in a Cluster, disabled by default.
- Add Cluster `ControlPlaneReady` condition to comply with upstream Cluster API implementation.
- Add AzureMachine `Ready`, `SubnetReady` and `VMSSReady` conditions to comply with upstream Cluster API implementation.
- Add MachinePool `Creating` condition to comply with upstream Cluster API implementation.

### kubernetes [1.18.14](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.14)

#### Feature
- Add a new flag to set priority for the kubelet on Windows nodes so that workloads cannot overwhelm the node there by disrupting kubelet process. ([#96158](https://github.com/kubernetes/kubernetes/pull/96158), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node]
#### Bug or Regression
- Avoid GCE API calls when initializing GCE CloudProvider for Kubelets.
  Avoid unnecessary GCE API calls when adding IP alises or reflecting them in Node object in GCE cloud provider. ([#96863](https://github.com/kubernetes/kubernetes/pull/96863), [@tosi3k](https://github.com/tosi3k)) [SIG Apps, Cloud Provider and Network]
- Bump node-problem-detector version to v0.8.5 to fix OOM detection in with Linux kernels 5.1+ ([#96716](https://github.com/kubernetes/kubernetes/pull/96716), [@tosi3k](https://github.com/tosi3k)) [SIG Cloud Provider, Scalability and Testing]
- Exposes and sets a default timeout for the SubjectAccessReview client for DelegatingAuthorizationOptions ([#96152](https://github.com/kubernetes/kubernetes/pull/96152), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Fix memory leak in kube-apiserver when underlying time goes forth and back. ([#96266](https://github.com/kubernetes/kubernetes/pull/96266), [@chenyw1990](https://github.com/chenyw1990)) [SIG API Machinery]
- Fix pull image error from multiple ACRs using azure managed identity ([#96355](https://github.com/kubernetes/kubernetes/pull/96355), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: resize Azure disk issue when it's in attached state ([#96705](https://github.com/kubernetes/kubernetes/pull/96705), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a bug that prevents kubectl to validate CRDs with schema using x-kubernetes-preserve-unknown-fields on object fields.
  Fix kubectl SchemaError on CRDs with schema using x-kubernetes-preserve-unknown-fields on array types. ([#96563](https://github.com/kubernetes/kubernetes/pull/96563), [@gautierdelorme](https://github.com/gautierdelorme)) [SIG API Machinery and Testing]
- Fixed kubelet creating extra sandbox for pods with RestartPolicyOnFailure after all containers succeeded ([#92614](https://github.com/kubernetes/kubernetes/pull/92614), [@tnqn](https://github.com/tnqn)) [SIG Node and Testing]
- Metric names for CSI and flexvolume drivers will include the driver name as well as the CSI plugin name. ([#96474](https://github.com/kubernetes/kubernetes/pull/96474), [@mattcary](https://github.com/mattcary)) [SIG Instrumentation and Storage]
- New Azure instance types do now have correct max data disk count information. ([#94340](https://github.com/kubernetes/kubernetes/pull/94340), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Cloud Provider and Storage]


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
