# :zap: Giant Swarm Release v33.1.2 for CAPA :zap:

<< Add description here >>

## Changes compared to v33.1.1

### Components

- cluster-aws from v6.4.1 to v6.4.2

### cluster-aws [v6.4.1...v6.4.2](https://github.com/giantswarm/cluster-aws/compare/v6.4.1...v6.4.2)

#### Added

- Add `kubernetes.io/cluster/$clusterName: "owned"` and `sigs.k8s.io/cluster-api-provider-aws/cluster/$clusterName: "owned"` tags to the `IRSAClaim` CR so that resources created by Crossplane contain the expected tags. This also allows to find the S3 buckets that need to be deleted when removing a cluster.

#### Fixed

- Fix Karpenter NodePool subnet filtering: when users define custom `subnetTags`, the default `giantswarm.io/role: "nodes"` filter is no longer applied, allowing full control over subnet selection. The cluster ownership tag (`sigs.k8s.io/cluster-api-provider-aws/cluster/<cluster-name>: owned`) is still enforced for security.
- Fix Karpenter HelmRelease: add missing `valuesFrom` parent field for `extraConfigs`, enabling customers to use custom ConfigMaps and Secrets for Karpenter configuration.
