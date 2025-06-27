# :zap: Giant Swarm Release v28.5.5 for CAPA :zap:

Upgrade cluster-aws to handle IMDS Hop Limit and patch kubernetes version.

## Changes compared to v28.5.4

### Components

- cluster-aws from v1.3.10 to v1.3.11
- Kubernetes from v1.28.14 to [v1.28.15](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.28.md#v1.28.15)

### cluster-aws [v1.3.10...v1.3.11](https://github.com/giantswarm/cluster-aws/compare/v1.3.10...v1.3.11)

#### Changed

- Reduce IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

### Apps
