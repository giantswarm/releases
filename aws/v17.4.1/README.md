# :zap: Giant Swarm Release v17.4.1 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [6.1.0](https://github.com/giantswarm/app-operator/releases/tag/v6.1.0)

#### Changed
- Use downward API to set deployment env var `KUBERNETES_SERVICE_HOST` to `status.hostIP`.
- Change `initialBootstrapMode` configuration value to `bootstrapMode`.
- Tighten pod and container security contexts for PSS restricted policies.
#### Added
- Allow to set api server pod port when enabling `initialBootstrapMode`.



### cert-manager [2.15.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.1)

#### Fixed
- Automatically try to execute `cmctl upgrade migrate-api-version` in crd install job to upgrade stored apiversions ([#245](https://github.com/giantswarm/cert-manager-app/pull/245))



