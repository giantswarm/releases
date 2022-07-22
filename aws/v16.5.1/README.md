# :zap: Giant Swarm Release v16.5.1 for AWS :zap:

This is a patch release to downgrade Cert Manager to release 2.13.0 due to breaking changes fo API Versions for older k8s releases.

## Change details

### cert-manager [2.13.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.13.0)

#### Revert breaking changes from Cert Manager `2.15.0` update with API versions
