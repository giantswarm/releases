# :zap: Giant Swarm Release v20.0.0-alpha1 for Azure :zap:

We are happy to announce our first Giant Swarm ClusterAPI alpha release.
Using this release to create cluster means using all the power from ClusterAPI but inside the Giant Swarm product.

## Change details

This is an alpha release, but you can already create clusters on a similar fashion you create regular Giant Swarm clusters.
Please notice that since this is an alpha release, the clusters you create won't be upgradable to next releases.
Also, please don't upgrade your existing clusters to this release.

With ClusterAPI you get the freedom to configure many aspects of your Kubernetes clusters directly from the Custom Resources.
We don't support all the Custom Resources just yet, but are working hard to provide you all that it's available on upstream ClusterAPI.
Currently we only support the `MachinePool` and `AzureMachinePool` objects for creating nodepools.
