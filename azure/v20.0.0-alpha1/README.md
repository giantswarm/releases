# :zap: Giant Swarm Release v20.0.0-alpha1 for Azure :zap:

This release provides initial support for creating clusters with Cluster API for Azure (CAPZ).

> **_Warning:_** This is an **`alpha preview release`** intended only for testing cluster creation. Upgrading to or from this version is not supported.

## Change details

Clusters are created in a similar way as regular Giant Swarm clusters by using the `v20.0.0-alpha1` release with the `kubectl gs` command:

```
kubectl gs template cluster --provider azure --release v20.0.0-alpha1 --organization giantswarm --description 'test' --output cluster.yaml
kubectl gs template nodepool  --provider azure --release "v20.0.0-alpha1" --organization giantswarm --cluster-name "hc27f" --description "np1" --nodes-min 3 --nodes-max 10 --output nodepool.yaml
```

At the moment, only `MachinePool` and `AzureMachinePool` Cluster API custom resources are supported.
