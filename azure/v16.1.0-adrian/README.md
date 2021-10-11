# :zap: Giant Swarm Release v16.1.0-adrian for Azure :zap:

<< Add description here >>

## Change details


### cert-operator [1.1.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.1.0)

#### Changed
- Adjust helm chart to be used with `config-controller`.
- Replace `jwt-go` with `golang-jwt/jwt`.
- Manage Secrets in the same namespace in which CertConfigs are found.
- Make `expirationThreshold` configurable.



