# :zap: Giant Swarm Release v20.1.1 for AWS :zap:

Fixes issue with cert-manager removing CRDs on rollback.

## Change details


### cert-manager [3.7.5](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.5)

#### Added
- Added annotation `helm.sh/resource-policy: keep` on CRDs to prevent them from being pruned in an unexpected rollback event.



