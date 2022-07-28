# :zap: Giant Swarm Release v17.4.2 for AWS :zap:

This is a patch release bringing improvements to cert-manager, fixing an issue which can occur during upgrades to `v17.4.1`. Fresh installations of `v17.4.1` still work, existing workload clusters should be upgraded to this release.

## Change details

### cert-manager [2.15.2](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.2)

#### Fixed

- RBAC for `cmctl upgrade migrate-api-version` ([#249](https://github.com/giantswarm/cert-manager-app/pull/249)).
