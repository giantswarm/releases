# :zap: Giant Swarm Release v25.1.2 for CAPA :zap:

This release allows to have different node pools with different configuration to support legacy cgroupsv1.

## Changes compared to v25.1.1

### Components

- cluster-aws from v1.1.2 to v1.1.3

### cluster-aws [v1.1.2...v1.1.3](https://github.com/giantswarm/cluster-aws/compare/v1.1.2...v1.1.3)

#### Changed

- Bump `cluster` chart to `0.35.3` so that we can configure node pools for cgroupsv1.
