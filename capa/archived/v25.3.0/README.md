# :zap: Giant Swarm Release v25.3.0 for CAPA :zap:

This release allows to have different node pools with different configuration to support legacy cgroupsv1.

## Changes compared to v25.2.1

### Components

- cluster-aws from v1.3.2 to v1.3.4

### cluster-aws [v1.3.2...v1.3.4](https://github.com/giantswarm/cluster-aws/compare/v1.3.2...v1.3.4)

#### Changed

- Bump `cluster` chart to `1.0.3` so that we can configure node pools for cgroupsv1.
- Chart: Update `cluster` to v1.0.2.
  - Chart: Add OS tooling named template.
