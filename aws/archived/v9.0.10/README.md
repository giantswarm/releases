# :zap: Tenant Cluster Release v9.0.10 for AWS :zap:

This release provides a new aws-operator which is fixing an issue with overlapping networks when creating legacy clusters together with node pools clusters.

## Change details

### aws-operator [5.5.6](https://github.com/giantswarm/aws-operator/releases/tag/v5.5.6)

- Check node pools clusters for reserved subnets.
