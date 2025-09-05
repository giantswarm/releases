# :zap: Giant Swarm Release v31.1.2 for CAPA :zap:

This release updates the `cluster-aws` chart to fix an issue around Helm values schema validation when using certain node pool fields.

## Changes compared to v31.1.1

### Components

- cluster-aws from v3.6.1 to v3.6.2

### cluster-aws [v3.6.1...v3.6.2](https://github.com/giantswarm/cluster-aws/compare/v3.6.1...v3.6.2)

#### Fixed

- Add `cluster` chart nodepool fields to the schema.
