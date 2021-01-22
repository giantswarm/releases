# :zap: Giant Swarm Release v13.1.1 for Azure :zap:

This is a bug fix release aimed at solving a bug that was affecting CIDR selection during workload clusters creation.

There is no need to upgrade existing workload clusters to this release.

## Change details

### azure-operator [5.2.1](https://github.com/giantswarm/azure-operator/releases/tag/v5.2.1)

#### Fixed
- Ensure the management cluster's network space is never used for workload clusters.



