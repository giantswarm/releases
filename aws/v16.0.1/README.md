# :zap: Giant Swarm Release v16.0.1 for AWS :zap:

This release provides a fix for `cert-operator` to ensure `certConfig` is in the same org namespace as the `Cluster` resource.

## Change details


### cert-operator [1.2.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.2.0)

#### Changed
- Introducing `v1alpha3` CR's.
#### Added
- Add check to ensure that the `Cluster` resource is in the same namespace as the `certConfig` before creating the secret there.



