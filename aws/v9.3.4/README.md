# :zap: Giant Swarm Release v9.3.4 for AWS :zap:

This release updates managed apps to the latest releases.

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

### cluster-operator [v0.23.10](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.10)

- Align with NGINX IC App 1.7.0, move of LB Service management from azure-operator to the app itself.

### coredns v1.6.5 ([Giant Swarm app v1.1.10](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v1110-2020-06-29))

- Make resource requests/limits configurable.
- Make forward options optional.

### kube-state-metrics v1.9.5 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
- Fixed invalid cluster role binding for Helm 3 compatibility.

### metrics-server v0.3.3 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### net-exporter [v1.9.0](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#190---2020-06-29)

- Add `ntp` collector.
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### nginx-ingress-controller v0.33.0 ([Giant Swarm app v1.7.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v171-2020-07-07))

- Upgraded upstream ingress-nginx from v0.30.0 to v0.33.0 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0330).
- Improved observability, enabled monitoring Service by default for monitoring targets discovery and removed support for disabling it.
