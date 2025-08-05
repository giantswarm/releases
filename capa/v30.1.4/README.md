# :zap: Giant Swarm Release v30.1.4 for CAPA :zap:

This release backports cluster-aws fix for reducing IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

## Changes compared to v30.1.3

### Components

- cluster-aws from v3.2.2 to v3.2.3

### cluster-aws [v3.2.2...v3.2.3](https://github.com/giantswarm/cluster-aws/compare/v3.2.2...v3.2.3)

#### Changed

- Reduce IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

### Apps
