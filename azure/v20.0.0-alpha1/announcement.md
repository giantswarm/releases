**Workload cluster release v20.0.0-alpha1 for Azure** is now available. It provides initial support for creating clusters with Cluster API for Azure (CAPZ). Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-azure/releases/azure-v20.0.0-alpha1/).

> **_Warning:_** This is an **`alpha preview release`** intended only for testing cluster creation. Upgrading to or from this version is not supported.

> **_Warning:_** There is a breaking change if `kubectl` is used manage machine pools. The `MachinePool` resource was moved to a new Kubernetes API Group named `machinepools.cluster.x-k8s.io` from `machinepools.exp.cluster.x-k8s.io`. The full resource path has to be specified when using `kubectl-gs` in order to access the machine pools created with the old API group. More details are available in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-azure/releases/azure-v20.0.0-alpha1/README.md#machine-pools-with-kubectl-gs)

```
kubectl get machinepools.exp.cluster.x-k8s.io -A
```

Clusters can be created with `kubectl-gs` by following the [documentation](https://docs.giantswarm.io/ui-api/kubectl-gs/template-cluster/#flags-capz).
