# :zap: Giant Swarm Release v29.6.4 for CAPA :zap:

This release backports a fix for reducing IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

## Changes compared to v29.6.3

### Components

- cluster-aws from v2.6.3 to v2.6.4

### cluster-aws [v2.6.3...v2.6.4](https://github.com/giantswarm/cluster-aws/compare/v2.6.3...v2.6.4)

#### Changed

- Reduce IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.

### Apps
