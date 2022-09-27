# :zap: Giant Swarm Release v17.4.1 for AWS :zap:

This is a patch release bringing improvements to app-operator and chart-operator such as pod and container security contexts for PSS and more. It also contains a change in cert-manager to automatically upgrade stored custom resources stored in deprecated apiversions.

## Change details


### app-operator [6.1.0](https://github.com/giantswarm/app-operator/releases/tag/v6.1.0)

#### Changed
- Use downward API to set deployment env var `KUBERNETES_SERVICE_HOST` to `status.hostIP`.
- Change `initialBootstrapMode` configuration value to `bootstrapMode`.
- Tighten pod and container security contexts for PSS restricted policies.
#### Added
- Allow to set api server pod port when enabling `initialBootstrapMode`.



### chart-operator [2.25.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.25.0)

#### Changed
- Tighten pod and container security contexts for PSS restricted policies.
- Use downward API to set deployment env var `KUBERNETES_SERVICE_HOST` to `status.hostIP`.
- Change `initialBootstrapMode` configuration value to `bootstrapMode`.
- Use private Helm client for installing app-operators from control-plane-test-catalog
#### Added
- Allow to set api server pod port when enabling `initialBootstrapMode`.



### cert-manager [2.15.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.1)

#### Fixed
- Automatically try to execute `cmctl upgrade migrate-api-version` in crd install job to upgrade stored apiversions ([#245](https://github.com/giantswarm/cert-manager-app/pull/245))



### vertical-pod-autoscaler [2.4.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.1)

#### Fixed
- Correct selector in admission controller PDB
