# :zap: Giant Swarm Release v14.1.2 for Azure :zap:

<< Add description here >>

## Change details


### cluster-operator [0.24.0](https://github.com/giantswarm/cluster-operator/releases/tag/v0.24.0)

#### Changed
- Remove `VersionBundle` version from `CertConfigs` and add the `cert-operator.giantswarm.io/version` label. **This change requires using `cert-operator` 1.0.0 or later**.



### cert-operator [1.0.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.0)

#### Changed
- Update Kubernetes dependencies to 1.18 versions.
- Reconcile `CertConfig`s based on their `cert-operator.giantswarm.io/version` label.
#### Removed
- Stop using the `VersionBundle` version.
#### Added
- Add network policy resource.
- Added lookup for nodepool clusters in other namespaces than `default`.
#### [0.1.0-2] - 2020-08-11
#### Fixed
- Skip validation of reference versions like `0.1.0-1`.
- Continue to export vault token expiration time as 0 when lookup fails.
#### Changed
- Update `apiextensions` to `0.4.1` version.
- Set version `0.1.0` in `project.go`.
- Use `architect` `2.1.2` in github actions.
#### [0.1.0-1] - 2020-08-07
#### Added
- Add `k8s-jwt-to-vault-token` init container to ensure *vault* token secret exists.
- Add Github automation workflows.



### azure-scheduled-events [0.2.2](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.2.2)

#### Fixed
- Disable helm hook for new installations of the chart.



