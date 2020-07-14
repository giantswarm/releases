# :zap: Giant Swarm Release v11.5.0 for AWS :zap:

This release makes cluster scaling more reliable with introduction of docker registry mirrors.

Specifically it enables tenant clusters, created using this release, to utilize images in additional image registries we set up. This in turn allows clusters to scale, even if one of the registries (e.g. `quay.io`) is unavailable or experiencing service disruption.

## aws-operator [v8.7.2](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.2)

### Added

- Add `--service.registry.mirrors` flag for setting registry mirror domains.
- Make registry domain & mirrors configurable based on region.

### Changed

- Replace `--service.registrydomain` with `--service.registry.domain`.
- Update `k8s-setup-network-env` image to `0.2.0`.

### Fixed

- Fix failing of ELB collector cache in case there is no ELB in AWS account


