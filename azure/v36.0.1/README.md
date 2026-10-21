# :zap: Giant Swarm Release v36.0.1 for Azure :zap:

## Changes compared to v36.0.0


### Apps

- external-dns from v3.5.0 to v3.5.1

### external-dns [v3.5.0...v3.5.1](https://github.com/giantswarm/external-dns-app/compare/v3.5.0...v3.5.1)

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0 and pull it from `oci://gsoci.azurecr.io/charts/giantswarm`. Its fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
