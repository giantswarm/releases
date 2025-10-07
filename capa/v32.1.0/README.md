# :zap: Giant Swarm Release v32.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.0.0


### Apps

- aws-nth-bundle from v1.2.2 to v1.3.0

### aws-nth-bundle [v1.2.2...v1.3.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.2...v1.3.0)

#### Changed

- Upgrade aws-nth-crossplane-resources to v1.3.0, fixing support for multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters, and supporting heartbeat sending
- Upgrade aws-node-termination-handler-app to v1.23.0, enabling heartbeats by default and upgrading to upstream application version v1.25.2 which fixes a resource leak bug relevant to heartbeat sending
- Upgrade aws-nth-crossplane-resources to v1.1.1, supporting multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters
