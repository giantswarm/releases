# :zap: Giant Swarm Release v20.1.4 for AWS :zap:

Fixes an issue with cert-manager acme-http01-solver-image argument and improves container security.

## Change details

### cert-manager [3.7.9](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.9)

#### Fix
- Remove quotes from acme-http01-solver-image argument. The quotes are used when looking up the image which causes an error.

### cert-manager [3.7.8](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.8)

#### Update
- Improves container security by setting `runAsGroup` and `runAsUser` greater than zero for all deployments.
