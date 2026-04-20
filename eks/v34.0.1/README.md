# :zap: Giant Swarm Release v34.0.1 for EKS :zap:

This patch release fixes an issue with the adoption of the CoreDNS service.

## Changes compared to v34.0.0

### Apps

- coredns from v1.29.1 to v1.30.0

### coredns [v1.29.1...v1.30.0](https://github.com/giantswarm/coredns-app/compare/v1.29.1...v1.30.0)

#### Added

- Add `coredns-adopter` job to adopt default CoreDNS resources on EKS clusters (disabled by default).

#### Changed

- Update `coredns` image to [1.14.2](https://github.com/coredns/coredns/releases/tag/v1.14.2).
