# :zap: Giant Swarm Release v35.2.3 for Azure :zap:

## Changes compared to v35.2.2

### Apps

- external-dns from v3.6.0 to v3.6.1

### external-dns [v3.6.0...v3.6.1](https://github.com/giantswarm/external-dns-app/compare/v3.6.0...v3.6.1)

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0, whose fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
