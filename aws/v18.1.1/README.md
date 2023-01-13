# :zap: Giant Swarm Release v18.1.1 for AWS :zap:

This release contains revised allocation calculation on the master nodes. Kubernetes api-server's cpu request are changed to be the half of the **available** CPUs in the VM. This fix will open up more space for other workloads scheduled to run on the master nodes.

## Change details


### aws-operator [14.1.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.1.1)

#### Changed
- Bump k8scc to [15.4.0](https://github.com/giantswarm/k8scloudconfig/blob/master/CHANGELOG.md#1540---2023-01-11). (Lower apiserver's cpu request to be 1/2 of the available CPUs in the VM.)



