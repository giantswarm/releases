# :zap: Giant Swarm Release v15.1.1 for Azure :zap:

This is a bugfix release that solves an issue preventing clusters to be successfully created on new organizations.

There is no need to upgrade workload clusters running the 15.1.0 release to this one because there are no changes in the
workload cluster components.

## Change details

### azure-operator [5.8.1](https://github.com/giantswarm/azure-operator/releases/tag/v5.8.1)

#### Fixed
- Fix namespace in secret reference of `AzureClusterIdentity`.
