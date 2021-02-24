# :zap: Tenant Cluster Release v11.5.1 for AWS :zap:

This patch release fixes two known issues found in our v11.5.0 release:

- The calculation of the maximum number of Pods to run per worker node was faulty, in some cases resulting in a too high limit. This would lead to Pods being scheduled to nodes where they could not get an IP address. The calculation is now improved to take into account all cases where network devices and IP addresses would be used in addition to Pod IPs.
- When deleting a cluster, the related custom resources would not always be cleaned up properly.

## Change details

### cluster-operator [v2.3.2](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.2)

- Handle error `basedomain not found` gracefully, so that `G8sControlPlane` and `MachineDeployment` CRs are deleted reliably.

### aws-operator [v8.7.5](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.5)

- Adjust `MAX_PODS` for master and worker nodes to max IPs per Elastic Network Interface (ENI)

### calico [v3.15.1](https://github.com/projectcalico/calico/releases/tag/v3.15.1)

Complete release notes can be found at [docs.projectcalico.org/v3.15/release-notes](https://docs.projectcalico.org/v3.15/release-notes/)
