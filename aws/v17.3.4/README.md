# :zap: Giant Swarm Release v17.3.4 for AWS :zap:

This release provides quality of life improvements and bug fixes to operators.

## Change details


### app-operator [5.12.0](https://github.com/giantswarm/app-operator/releases/tag/v5.12.0)

#### Added
- Add initialBootstrapMode flag to allow deploying CNI as managed apps.

#### Changed
- Only set resource limits on the deployment when the VPA is not available or disabled
- Increase min / max resource limits on VPA


### chart-operator [5.24.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0)

#### Added
- Split Helm client into private Helm client for giantswarm-namespaced apps and public Helm client for rest of the apps.

#### Changed
- Always create giantswarm-critical priority class if it does not exist.
- Add chart-pull-failed error to differentiate between issues when pulling chart tarball and other problems.

#### Fixed
- Fix missing PriorityClass issue.


### metrics-server-app [v1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set kubelet-preferred-address-types to Hostname on AWS.
