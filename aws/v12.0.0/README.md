# :zap: Giant Swarm Release v12.0.0 for AWS :zap:

This is the first release to support Kubernetes 1.17 on AWS.

This release provides new aws-operator and cluster-operator versions with reliability improvements and picks up upgrades to many components.

## Change details

### Kubernetes [v1.17.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.9)

#### Known issues

- Former feature gate `AttachVolumeLimit`, which was disabled in Giant Swarm tenant clusters, is now always active. As a result, limits for Persistent Volumes will be set to a wrong value in tenant clusters. In cases where many Persistent Volumes are used in a cluster, this may lead to a situation where the scheduler falsely assigns a Pod to a node, assuming that a volume can be attached to this node, while in fact it cannot. [kubernetes#92799](https://github.com/kubernetes/kubernetes/issues/92799)

#### Significant changes

- CVE-2020-8557 (Medium): Node-local denial of service via container /etc/hosts file. See [kubernetes#93032](https://github.com/kubernetes/kubernetes/issues/93032) for more details.

For a complete list of changes, known issues, removals and deprecations, please check the [Kubernetes 1.17 release notes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md).

### aws-operator [v8.7.5](https://github.com/giantswarm/aws-operator/blob/master/CHANGELOG.md#875---2020-07-30)

- Adjusted `MAX_PODS` setting for master and worker nodes to max IPs per Elastic Network Interface (ENI).

### calico [v3.15.1](https://github.com/projectcalico/calico/releases/tag/v3.15.1)

- Fix issue with service IP advertisement breaking host service connectivity

Complete release notes can be found at [docs.projectcalico.org/v3.15/release-notes](https://docs.projectcalico.org/v3.15/release-notes/)

### chart-operator [v0.13.2](https://github.com/giantswarm/chart-operator/blob/v0.13.2/CHANGELOG.md#v0132-2020-06-23)

- Add metrics for Helm releases with a mismatched namespace.
- Calculate md5sum from go struct.

### cluster-autoscaler [v1.17.3](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.17.3) (Giant Swarm app [v1.17.3](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md#1173---2020-07-30))

- Switch leader election mechanism to use lease objects, includes RBAC rule update
- Nodes with small difference in available memory will now be considered similar for the purposes of balancing node pool sizes. This should increase the reliability of node pool balancing.
- Fixed a bug where Cluster Autoscaler incorrectly didn't take into account resource requests of initContainers.
- Allow marking pods with the `cluster-autoscaler.kubernetes.io/daemonset-pod` annotation. Marked pods will be treated as daemon set pods.
- Add TTL and batching to launch config cache mechanism.
- improved logging when a taint blocks scale-up
- Invalidate node instances cache after deleting failed and unregistered nodes
- Allow for greater node memory mismatch on comparison for similarity (bumped to 256K from 128K)
- Fix case when ASG size could be descresed twice
- Allow arbitrary placeholder AWS instance names

### cluster-operator [v2.3.2](https://github.com/giantswarm/cluster-operator/blob/master/CHANGELOG.md#232---2020-07-31)

- Fixes a problem where the handling of a "basedomain not found" error would prevent the proper removal of `G8sControlPlane` and `MachineDeployment` CRs when deleting a cluster.

### kiam [v3.6](https://github.com/uswitch/kiam/blob/master/CHANGELOG.md#v36) (Giant Swarm app [v1.3.1](https://github.com/giantswarm/kiam-app/blob/master/CHANGELOG.md#131---2020-07-23))

Check the [kiam changelog](https://github.com/uswitch/kiam/blob/master/CHANGELOG.md#v36) for details.

### kube-state-metrics [v1.9.7](https://github.com/kubernetes/kube-state-metrics/blob/master/CHANGELOG.md#v197--2020-05-24) (Giant Swarm app [v1.1.1](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#111---2020-07-22))

- Switch mutatingWebhookConfiguration to use v1 api

Check the [kube-state-metrics changelog](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.9.7) for more details.

### metrics-server [v0.3.6](https://github.com/kubernetes-sigs/metrics-server/releases/tag/v0.3.6) (Giant Swarm app [v1.1.1](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#111---2020-07-23))

- Fix: Don't break metric storage when duplicate pod metrics encountered

Check the [metrics-server changelog](https://github.com/kubernetes-sigs/metrics-server/releases) for more details.

### node-exporter [v1.0.1](https://github.com/prometheus/node_exporter/blob/master/CHANGELOG.md#101--2020-06-15) (Giant Swarm app [v1.3.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#130---2020-07-23))

- Several changes regarding metrics and labels.

Check the [node-exporter changelog](https://github.com/prometheus/node_exporter/blob/master/CHANGELOG.md#101--2020-06-15) for detais.
