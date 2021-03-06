# :zap: Tenant Cluster Release 10.1.2 for AWS :zap:

This release fixes a problem with the cluster role for chart-operator, which is responsible for installing and updating Managed Apps. The changes affect only Managed Apps and upgrading to this release will not cause cluster nodes to be rolled. **To prevent encountering problems when installing or updating an app, please upgrade to this release.**

**Please Note:** We are still resolving the issue with network traffic between node pools in 10.x and 11.x clusters reported yesterday (5th Feb 2020).

### chart-operator [v0.11.3](https://github.com/giantswarm/chart-operator/releases/tag/v0.11.3)
- Adjusted RBAC permissions.

### cluster-autoscaler v1.16.2 ([Giant Swarm app v1.1.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.1.3))
- Updated manifests for Kubernetes 1.16 compatibility.

### kiam v3.4 ([Giant Swarm app v1.0.4](https://github.com/giantswarm/kiam-app/releases/tag/v1.0.3))
- Updated manifests for Kubernetes 1.16 compatibility.

### cert-manager v0.9.0 ([Giant Swarm app v1.0.4](https://github.com/giantswarm/cert-manager-app/releases/tag/v1.0.4))
- Improvements for clusters with restrictive network policies.

### kube-state-metrics v1.9.2 ([Giant Swarm app v1.0.2](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#v102))
- Updated to upstream version 1.9.2 - [changelog](https://github.com/kubernetes/kube-state-metrics/blob/master/CHANGELOG.md#v192--2020-01-13).
- Adjusted RBAC configuration.

### net-exporter [v1.5.1](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#151-2020-01-08)
- Changed priority class to `system-node-critical`.
