# :zap: Giant Swarm Release v34.1.1 for Azure :zap:

## Changes compared to v34.1.0

### Apps

- cert-manager from v3.11.0 to v3.13.0

### cert-manager [v3.11.0...v3.13.0](https://github.com/giantswarm/cert-manager-app/compare/v3.11.0...v3.13.0)

#### Added

- Add control plane node toleration to CA injector deployment.

#### Changed

- Upgrade cert-manager to v1.19.4.

#### Removed

- Remove PodSecurityPolicy (PSP) and related resources.
- Remove Giant Swarm PSP to PSS migration logic.
