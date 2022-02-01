# :zap: Giant Swarm Release v20.0.0-alpha1 for Azure :zap:

We are happy to announce our first Giant Swarm ClusterAPI alpha release.
Using this release to create cluster means using all the power from ClusterAPI but inside the Giant Swarm product.

## Change details

This is an alpha release, but you can already create clusters in a similar fashion as your regular Giant Swarm clusters.

```
kubectl gs template cluster --provider azure --release v20.0.0-alpha1 --organization giantswarm --description 'test' --output cluster.yaml
kubectl gs template nodepool  --provider azure --release "v20.0.0-alpha1" --organization giantswarm --cluster-name "hc27f" --description "np1" --nodes-min 3 --nodes-max 10 --output nodepool.yaml
```

Please notice that since this is an alpha release, the clusters you create won't be upgradable to next releases.
Also, please don't upgrade your existing clusters to this release.

With ClusterAPI you get the freedom to configure many aspects of your Kubernetes clusters directly from the Custom Resources.
We don't support all the Custom Resources just yet, but are working hard to provide you all that are available on upstream ClusterAPI.
Currently we only support the `MachinePool` and `AzureMachinePool` objects for creating nodepools.
