# :zap: Giant Swarm Release v14.1.2 for Azure :zap:

This is a bugfix release to resolve a few bugs related to the cluster autoscaler.
We strongly suggest upgrading any 14.x workload cluster to this release to ensure the cluster autoscaler feature works properly.

Warning: to avoid downtimes in the ingress-based workloads, before upgrading to this release it is important to ensure your cluster has a recent version (1.14.0 or newer) of the `Nginx Ingress Controller APP` running. Please get in touch with your Solution Engineer before upgrading if you have any concern.

## Change details

### azure-operator [5.5.1](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.1)

### Fixed

- Fix a race condition when upgrading node pools with 0 replicas.
- Fix Upgrading condition for node pools with autoscaler enabled.

### Added

- Add new handler that creates `AzureClusterIdentity` CRs and the related `Secrets` out of Giant Swarm's credential secrets.
- Ensure `AzureCluster` CR has the `SubscriptionID` field set.
- Reference `Spark` CR as bootstrap reference from the `MachinePool` CR.
- Ensure node pools min size is applied immediately when changed.


### azure-scheduled-events [0.2.2](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.2.2)

#### Fixed
- Disable helm hook for new installations of the chart.



