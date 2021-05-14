# :zap: Giant Swarm Release v13.2.0 for AWS :zap:

This release fixes an issue that can cause an IP conflict to occur in certain situations when a node pool is created. It also contains a newer version of Cert Manager to mitigate failed cert-manager-app installations due to Custom Resource Definition conversion issues.

## Change details


### aws-operator [9.3.10](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.10)

- Fix IPAM conflicts when creating a node pool


### cert-manager [2.7.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.0)

- Update to upstream `v1.3.1` ([#155](https://github.com/giantswarm/cert-manager-app/pull/155)). This mitigates failed cert-manager-app installations due to CRD conversion issues.
- cert-manager-app now requires kubernetes version >=1.16.0. ([#151](https://github.com/giantswarm/cert-manager-app/pull/151))
- Switch rbac rules from `extensions` to `networking.k8s.io` for ingresses. ([#151](https://github.com/giantswarm/cert-manager-app/pull/151))
- Allow strings and integers in values schema for resources requests and limits. ([#150](https://github.com/giantswarm/cert-manager-app/pull/150))
- Rename clusterissuer subchart to match it's name in its Chart.yaml. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140))
- Make pods of deployments use read-only file systems. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140))
- Make pre-install/pre-upgrade hooks use server side apply. Possibly fixes upgrade timeouts. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140))
