# :zap: Giant Swarm Release v26.4.4 for CAPA :zap:

Upgrade cluster-aws to handle IMDS Hop Limit.

## Changes compared to v26.4.3

### Components

- cluster-aws from v1.3.10 to v1.3.11

### cluster-aws [v1.3.10...v1.3.11](https://github.com/giantswarm/cluster-aws/compare/v1.3.10...v1.3.11)

#### Changed

- Reduce IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

