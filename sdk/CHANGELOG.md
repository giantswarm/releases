# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [0.5.0] - 2024-07-23

### Added

- Add Azure provider to the list of supported providers.

## [0.4.0] - 2024-06-21

### Added

- `Client`
  - `GetReleasesForGitReference` returns all releases for the specified provider and from the specified git reference.
  - `GetReleaseForGitReference` returns a release with the specified release version for the specified provider and from
    the specified git reference.
  - `GetNewReleasesForGitReference` returns newly added releases for the specified provider and from the specified git reference.

## [0.3.0] - 2024-06-12

### Added

- Release CRD is generated in the SDK's `manifests` directory.

### Fixed

- When creating a new release, builder sets Release date to now (current moment).

## [0.2.0] - 2024-06-07

### Added

- `Client` gets Release resources from GitHub `giantswarm/releases` repo
  - `GetRelease` func to get a Release resource for the specified provider and release version
  - `GetLatestRelease` to get a Release resource for the latest release version for the specified provider
- `Builder` builds Release resources for the new releases based on the existing releases, with functions:
  - `WithClusterApp` to override `cluster-$provider` app,
  - `WithApp` to override a default app,
  - `WithPreReleasePrefix` to set a custom prefix that is prepended to the pre-release of the custom release version,
  - `WithRandomPreRelease` to specify that custom release version will have a random pre-release with specified length.

## [0.1.0] - 2024-06-05

### Added

- `Release` type that represents a Giant Swarm workload cluster release.

[Unreleased]: https://github.com/giantswarm/releases/compare/sdk/v0.5.0...HEAD
[0.5.0]: https://github.com/giantswarm/releases/releases/tag/sdk/v0.5.0
[0.4.0]: https://github.com/giantswarm/releases/releases/tag/sdk/v0.4.0
[0.3.0]: https://github.com/giantswarm/releases/releases/tag/sdk/v0.3.0
[0.2.0]: https://github.com/giantswarm/releases/releases/tag/sdk/v0.2.0
[0.1.0]: https://github.com/giantswarm/releases/releases/tag/sdk/v0.1.0
