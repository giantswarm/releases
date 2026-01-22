# :zap: Giant Swarm Release v33.1.4 for CAPA :zap:

<< Add description here >>

## Changes compared to v33.1.3

### Components

- cluster-aws from v6.4.3 to v7.2.5

### cluster-aws [v6.4.3...v7.2.5](https://github.com/giantswarm/cluster-aws/compare/v6.4.3...v7.2.5)

#### :warning: Breaking Changes

- The following IAM permissions have been removed from the control plane nodes
- autoscaling:SetDesiredCapacity
- autoscaling:TerminateInstanceInAutoScalingGroup
- Removed `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers` value, as that's the default behavior now. It cannot be overridden anymore.

#### Added

- Add `kubernetes.io/cluster/$clusterName: "owned"` and `sigs.k8s.io/cluster-api-provider-aws/cluster/$clusterName: "owned"` tags to the `IRSAClaim` CR so that resources created by Crossplane contain the expected tags. This also allows to find the S3 buckets that need to be deleted when removing a cluster.
- *This change will roll the control plane nodes* Add `preKubeadmCommand` to wait for the API server load balancer DNS to be resolvable before running kubeadm on control plane nodes. This prevents kubeadm from failing when the ELB DNS record hasn't propagated yet.
- *This change will roll the nodes* Add Crossplane IAM Roles, policies and instance profiles for worker and control plane nodes. Instead of having an IAM Role per node pool, now we'll use the same for all node pools.
- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- *This change will roll the nodes on Karpenter node pools* Attach the `lb` Security Group to Karpenter nodes.
- *This change will roll the nodes on Karpenter node pools* Name instance on AWS after the nodepool name.

#### Changed

- Chart: Update `cluster` to v5.1.2.
- Chart: Update `cluster` to v5.1.1.
- Chart: Update `cluster` to v5.1.0.
- Chart: Update `cluster` to v5.0.0.
- Reduce redundant parts of JSON schema for Karpenter vs. MachinePool types of node pools
- Adjust node max pods based on the `nodeCidrMaskSize`

#### Fixed

- Fix Karpenter schema definition: changed from `app` schema to `helmRelease` schema to correctly reflect that Karpenter is deployed as a HelmRelease resource. This fixes incorrect field definitions in `extraConfigs` (capitalized enum values `ConfigMap`/`Secret` and `optional` field instead of `priority`).
- Fix Karpenter NodePool subnet filtering: when users define custom `subnetTags`, the default `giantswarm.io/role: "nodes"` filter is no longer applied, allowing full control over subnet selection. The cluster ownership tag (`sigs.k8s.io/cluster-api-provider-aws/cluster/<cluster-name>: owned`) is still enforced for security.
- Fix Karpenter HelmRelease: add missing `valuesFrom` parent field for `extraConfigs`, enabling customers to use custom ConfigMaps and Secrets for Karpenter configuration.
- Ensure `AWSCluster.spec.network.subnets.tags` is not rendered as `null`
- Add missing documentation for node pools (health checks were not listed)
- Ensure defaulting `maxHealthyPercentage` since Helm does not use the default from the schema

#### Removed

- Remove `RolePolicyAttachment` crossplane custom resources as they are not needed when using `Role` and `RolePolicy`.

