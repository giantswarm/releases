# :zap: Giant Swarm Release v13.0.1 for Azure :zap:

This is the first tenant cluster release to support Kubernetes 1.18 and node pools on Azure.

A node pool is a subset of the Kubernetes nodes. They enable having pools of nodes with different configurations (like a different instance size) within one cluster.
After cluster creation with 1 node pools, additional node pools can be freely added and removed from the cluster.

If you have access to the Control Plane API you can manage your clusters directly from there.
The clusters that you create are now based on [Cluster API](https://cluster-api.sigs.k8s.io/) and [Cluster API Azure](https://capz.sigs.k8s.io/) CRDs ([Custom Resource Definition](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)) being reconciled  by Giant Swarm Controllers implementation.
Using [our kubectl plugin](https://github.com/giantswarm/kubectl-gs/) you can easily create the Custom Resources required to create a cluster.
Follow our [guide](https://docs.giantswarm.io/guides/creating-clusters-via-crs-on-azure/) for cluster creation and deletion to have a peak at the cluster management via the Control Plane API.

Please note that some endpoints for the clusters management via the Giant Swarm API will change from `v4` to `v5`, follow the [GS API specs](https://docs.giantswarm.io/api/#tag/clusters) for more details.

Also note the version is v13.0.1 due to problems found with the reconciliation of apps in the unannounced and unused v13.0.0 release.

## Change details
Please analyse the release notes of all components upgrades below:
### kubernetes [1.18.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v11812)
- Check changelog starting from [1.18.0](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md)
### app-operator [2.7.0](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#270---2020-11-09)
- Check changelog starting from [2.4.0](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#240---2020-10-23)
### azure-operator [5.0.0](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md#500---2020-12-01)
- Check changelog starting from [5.0.0-alpha1](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md#500---2020-12-01)
### cluster-operator [0.23.19](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02319---2020-12-03)
- Check changelog starting from [0.23.18](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02318---2020-10-21)
### calico 3.15.1 -> [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)
- Check changelog starting from [3.15.2](https://github.com/projectcalico/calico/releases/tag/v3.15.2)
### cert-exporter [1.3.0](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#130---2020-09-17)
### chart-operator 2.3.5 -> [2.5.1](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#251---2020-12-01)
- Check changelog starting from [2.4.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#240---2020-10-29)
### containerlinux [2605.8.0](https://www.flatcar-linux.org/releases/#release-2605.8.0)
- Check changelog starting from [2512.3.0](https://www.flatcar-linux.org/releases/#release-2512.3.0)
### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)
- Check changelog starting from [3.4.10](https://github.com/etcd-io/etcd/releases/tag/v3.4.10)
### kube-state-metrics [1.3.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#130---2020-11-25)
- Check changelog starting from [1.1.1](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#130---2020-11-25)
### node-exporter [1.7.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#170---2020-11-26)
- Check changelog starting from [1.3.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#170---2020-11-26)
### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#192---2020-08-21)
- Check changelog starting from [1.9.0](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#192---2020-08-21)
