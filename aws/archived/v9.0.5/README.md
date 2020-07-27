# :zap: Giant Swarm Release v9.0.5 for AWS :zap:

This release fixes a rare bug that would prevent the NGINX IC from being installed on a new cluster.

This bug would only occur on cluster creation if you had a nginx-ingress-controller-user-values configmap in the kube-system namespace while the cluster was still initialising.

Solution Engineers have already done the manual fix for affected customers.

It also modifies release templates to support the coming upgrade to Helm 3.

**Note to SEs when upgrading from 8.5.0 or 9.0.0:** Existing customer automation or processes that manage the configuration of coredns, nginx-ingress-controller, or cluster-autoscaler must be modified in order to work with the changed location and format of the *-user-values configmaps. Please see our docs on [Giant Swarm Release Versions: Versions that use the App Platform](https://docs.giantswarm.io/reference/release-versions/#versions-that-use-the-app-platform) for more details.

**Note for future 9.0.x releases:** Please include this note and the one above in all future 9.0.x releases.

## cluster-operator [v0.23.9](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.9)

- Fix a bug in user values migration logic for apps.

## cert-exporter (GS [v1.2.2](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#v122-2020-04-01))

- Change daemonset to use release revision not time for Helm 3 support.

## net-exporter [v1.7.1](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#v171-2020-04-01)

- Change daemonset to use release revision not time for Helm 3 support.
- Only set hosts arg if a value is present.
- Remove label from role ref in cluster role binding.
