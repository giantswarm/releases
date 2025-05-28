# :zap: Giant Swarm Release v28.3.0 for CAPA :zap:

This release allows to have different node pools with different configuration to support legacy cgroupsv1.

## Changes compared to v28.2.0

### Components

- cluster-aws from v1.3.3 to v1.3.4

### cluster-aws [v1.3.3...v1.3.4](https://github.com/giantswarm/cluster-aws/compare/v1.3.3...v1.3.4)

#### Changed

- Bump `cluster` chart to `1.0.3` so that we can configure node pools for cgroupsv1.
