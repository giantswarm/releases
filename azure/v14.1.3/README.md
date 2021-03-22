# :zap: Giant Swarm Release v14.1.3 for Azure :zap:

This releases increases the Azure Events Termination timeout from 5 to 15 minutes for better upgrade experience while workloads are moved to new nodes.
The draining process of the nodes has been improved as well.

## Change details


### azure-operator [5.5.2](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.2)

#### Changed
- Increase VMSS termination events timeout to 15 minutes.
#### Fixed
- Avoid logging errors when trying to create the workload cluster k8s client and cluster is not ready yet.



### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

#### Security
- Restrict ingress to only expose the status endpoint.



### azure-scheduled-events [0.3.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.3.0)

#### Fixed
- Ensure to wait long enough when draining a node before considering the node drained.
#### Changed
- Change drain timeout to 15 minutes.



