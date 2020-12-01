# :zap: Giant Swarm Release v13.0.0 for Azure :zap:

This is the first tenant cluster release to support Kubernetes 1.18 and node pools on Azure.

A node pool is a subset of the Kubernetes nodes. They enable having pools of nodes with different configurations (like a different instance size) within one cluster.
After cluster creation with 1 node pools, additional node pools can be freely added and removed from the cluster.

If you have access to the Control Plane API you can manage your clusters directly from there.
The clusters that you create are now represented by [Cluster API](https://cluster-api.sigs.k8s.io/) CRDs ([Custom Resource Definition](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)).
Using [our kubectl plugin](https://github.com/giantswarm/kubectl-gs/) you can easily create the Custom Resources required to create a cluster.

## Change details

### kubernetes [1.18.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v11812)
### app-operator [2.7.0](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#270---2020-11-09)
### azure-operator [5.0.0](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md#500---2020-12-01)
### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)
### cert-exporter [1.3.0](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#130---2020-09-17)
### chart-operator [2.5.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#250---2020-11-09)
### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02318---2020-10-21)
### containerlinux [2605.8.0](https://www.flatcar-linux.org/releases/#release-2605.8.0)
### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)
### external-dns [1.5.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#150---2020-10-07)
### kube-state-metrics [1.3.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#130---2020-11-25)
### node-exporter [1.7.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#170---2020-11-26)
### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#192---2020-08-21)
