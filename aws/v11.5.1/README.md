# :zap: Giant Swarm Release v11.5.1 for AWS :zap:

This patch release fixes two known issues found in our v11.5.0 release:

- The calculation of the maximum number of Pods to run per worker node was faulty, in some cases resulting in a too high limit. This would lead to Pods being scheduled to nodes where they could not get an IP address. The calculation is now improved to take into account all cases where network devices and IP addresses would be used in addition to Pod IPs.
- When deleting a cluster, the related custom resources would not always be cleaned up properly.

## Change details

### cluster-operator [v2.3.2](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.2)

- Handle error basedomain not found gracefully, so that G8sControlPlane CR and MachineDeployment CRs can be deleted

### aws-operator [v8.7.5](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.5)

### calico [v3.15.1](https://github.com/projectcalico/calico/releases/tag/v3.15.1)

- Fix an issue with service IP advertisement breaking host service connectivity

Complete release notes can be found at [docs.projectcalico.org/v3.15/release-notes](https://docs.projectcalico.org/v3.15/release-notes/)

#### Changed

- Adjust number of host network pods on worker node for aws-cni
