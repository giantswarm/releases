# :zap: Giant Swarm Release v17.3.4 for AWS :zap:

This release provides quality of life improvements and bug fixes to operators and apps.

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


### metrics-server [v1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set kubelet-preferred-address-types to Hostname on AWS.


### vertical-pod-autoscaler [v2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

#### Changed
- Change log-level from default=4 to 3.
- Rename resources to include release name as prefix and avoid name collision.
- Upgrade vertical-pod-autoscaler to 0.10.0

    API changes:
    Added support for alternative recommenders.
    Added support for per-VPA Object MinReplicas.
    
    Other notable changes:
    Added support for running VPA out of cluster.
    Use v1 API for storage instead of v1beta2.
    Allow configuring default update threshold.
    Use v1 API to register admission webhook.
    
    Bug fixes:
    Use correct timestamp for checkpoints.
    Issues with setting limits.
    Deploying VPA in different namespaces.
    Loading history.

- Use patched docker image tagged 0.10.0-oomfix for recommender and updater (see giantswarm/roadmap#923).

#### Fixed
- Fix admission-controller webhook-service name.
- Fix webhook name in generated secret certificate.
- Prefix generated secret certificate with release-name.


### node-exporter [v1.12.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.12.0)

#### Added
- Add options to be able to disable nvme and conntrack collectors.

#### Changed
- Enabled diskstats collector.
- Disable the fibrechannel collector.
- Disable the tapestats collector.
