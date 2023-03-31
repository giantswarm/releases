# :zap: Giant Swarm Release v18.3.0 for AWS :zap:

This release contains changes that improves the stability of cluster upgrades

## Change details


### aws-operator [14.8.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.8.1)

#### Changed
- Set ENV for nftables in `aws-cni`.
- Bump k8s-api-healthz image to 0.2.0.

#### Fixed
- Don't mark master instance as unhealthy if local etcd instance is unresponsive but the whole etcd cluster is also down.
- Don't mark master instance as unhealthy if local API server instance is unresponsive but the whole API server is also down.


