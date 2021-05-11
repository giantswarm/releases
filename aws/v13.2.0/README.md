# :zap: Giant Swarm Release v13.2.0 for AWS :zap:

This release fixes an issue that can cause an IP conflict to occur in certain situations when a node pool is created. It also contains a newer version of Cert Manager to mitigate failed cert-manager-app installations due to Custom Resource Definition conversion issues.

## Change details


### aws-operator [9.3.10](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.10)

- Fix IPAM conflicts when creating a node pool


### cert-manager [2.7.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.0)

- Update to upstream `v1.3.1` ([#155](https://github.com/giantswarm/cert-manager-app/pull/155)). This mitigates failed cert-manager-app installations due to CRD conversion issues.
