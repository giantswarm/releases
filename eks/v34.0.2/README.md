# :zap: Giant Swarm Release v34.0.2 for EKS :zap:

## Changes compared to v34.0.1

### Components

- Added containerd [v1.7.29](https://github.com/containerd/containerd/releases/tag/v1.7.29)

### Apps

- external-dns from v3.4.0 to v3.4.1

### external-dns [v3.4.0...v3.4.1](https://github.com/giantswarm/external-dns-app/compare/v3.4.0...v3.4.1)

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0 and pull it from `oci://gsoci.azurecr.io/charts/giantswarm`. Its fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
