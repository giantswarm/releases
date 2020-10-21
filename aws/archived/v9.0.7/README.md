## :zap: Giant Swarm Release 9.0.7 for AWS :zap:

This release updates managed apps to the latest releases.

**Note to SEs when upgrading from 8.5.0 or 9.0.0:** Existing customer automation or processes that manage the configuration of coredns, nginx-ingress-controller, or cluster-autoscaler must be modified in order to work with the changed location and format of the *-user-values configmaps. Please see our docs on [Giant Swarm Release Versions: Versions that use the App Platform](https://docs.giantswarm.io/reference/release-versions/#versions-that-use-the-app-platform) for more details.

**Note for future 9.0.x releases:** Please include this note and the one above in all future 9.0.x releases.

Below, you can find more details on components that were changed with this release.

### cert-exporter [v1.2.3](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#v123-2020-05-15)

- Updated prometheus/client_golang dependency.
- Migrated from dep to go modules.
- Moved to App deployment.

### cluster-autoscaler v1.16.5 ([Giant Swarm app v1.16.0](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md#v1160-2020-05-26))

- Upgraded upstream cluster-autoscaler from v1.16.2 to v1.16.5 - [changelog](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.16.5).
- Set `scan-interval` to 30 seconds (from 10 seconds) to save resources.
- Set `scale-down-unneeded-time` to 5 minutes (from the default of 10 minutes) to release unneeded nodes earlier.
- Lower `scaleDownUtilizationThreshold` to 0.5.

### cluster-operator [v0.23.12](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.12)

- Align with NGINX IC App 1.7.0, move of LB Service management from azure-operator to the app itself.

### coredns v1.6.5 ([Giant Swarm app v1.2.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v120-2020-07-13))

- Make resource requests/limits configurable.
- Make forward options optional.
- Apply a readiness probe.
- Increase the liveness probe failure threshold from 5 failures to 7 failures.

### kube-state-metrics v1.9.5 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
- Fixed invalid cluster role binding for Helm 3 compatibility.

### metrics-server v0.3.3 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### net-exporter [v1.9.0](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#190---2020-06-29)

- Add `ntp` collector.
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### nginx-ingress-controller v0.34.0 ([Giant Swarm app v1.7.2](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v172-2020-07-10))

- Upgraded upstream ingress-nginx from v0.30.0 to v0.34.0 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0340).
- Improved observability, enabled monitoring Service by default for monitoring targets discovery and removed support for disabling it.
