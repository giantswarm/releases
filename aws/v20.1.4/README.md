# :zap: Giant Swarm Release v20.1.4 for AWS :zap:

This patch release fixes an issue with cert-manager acme-http01-solver-image argument and improves container security.

## Change details

### cert-manager [3.7.9](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.9)

#### Fix
- Remove quotes from acme-http01-solver-image argument. The quotes are used when looking up the image which causes an error.

#### Update
- Improves container security by setting `runAsGroup` and `runAsUser` greater than zero for all deployments.

### security-bundle [1.6.7](https://github.com/giantswarm/security-bundle/releases/tag/v1.6.7)

#### Changed
- Update `kyverno` (app) to v0.17.13.
- Update `trivy-operator` (app) to v0.8.1.
- Update `starboard-exporter` (app) to v0.7.10.
