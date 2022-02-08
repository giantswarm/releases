# :zap: Giant Swarm Release v20.0.0-alpha1 for Azure :zap:

This release provides initial support for creating clusters with Cluster API for Azure (CAPZ).

> **_Warning:_** This is an **`alpha preview release`** intended only for testing cluster creation. Upgrading to or from this version is not supported.

> **_Warning:_** There is a breaking change if `kubectl` is used manage machine pools. The `MachinePool` resource was moved to a new Kubernetes API Group named `machinepools.cluster.x-k8s.io` from `machinepools.exp.cluster.x-k8s.io`. The full resource path has to be specified when using `kubectl-gs` in order to access the machine pools created with the old API group. More details are available [below](#machine-pools-with-kubectl-gs).

```
kubectl get machinepools.exp.cluster.x-k8s.io -A
```

## Change details

Clusters are created in a similar way as regular Giant Swarm clusters by using the `v20.0.0-alpha1` release with the `kubectl gs` command:

```
kubectl gs template cluster --provider azure --release v20.0.0-alpha1 --organization giantswarm --description 'test' --output cluster.yaml
kubectl gs template nodepool  --provider azure --release "v20.0.0-alpha1" --organization giantswarm --cluster-name "hc27f" --description "np1" --nodes-min 3 --nodes-max 10 --output nodepool.yaml
```

At the moment, only `MachinePool` and `AzureMachinePool` Cluster API custom resources are supported.

### Machine pools with kubectl-gs

There is a breaking change that you should be aware of if you use `kubectl` to manage `MachinePools`.
The `MachinePool` and `AzureMachinePool` CRDs have been moved to a new kubernetes API Group.
The `MachinePool` object has moved from `machinepools.exp.cluster.x-k8s.io` to `machinepools.cluster.x-k8s.io`, and the `AzureMachinePool` from `azuremachinepools.exp.infrastructure.cluster.x-k8s.io` to `azuremachinepools.infrastructure.cluster.x-k8s.io`.

This means that when you use `kubectl get machinepools` you no longer will see your `MachinePools`, because they were created using the old API Group and `kubectl` defaults to the new API Group.
To get the same behavior that we used to have you need to specify the old API Group for `MachinePools`

```
kubectl get machinepools.exp.cluster.x-k8s.io -A
```

This is not the case for `AzureMachinePools`. Using `kubectl get azuremachinepools` you will see `AzureMachinePools` using the old API Group.
