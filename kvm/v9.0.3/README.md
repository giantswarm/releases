# :zap: Giant Swarm Release v9.0.3 for KVM :zap:

This release fixes a rare bug that would prevent the NGINX IC from being installed on a new cluster.

This bug would only occur on cluster creation if you had a nginx-ingress-controller-user-values configmap in the kube-system namespace while the cluster was still initialising.

Solution Engineers have already done the manual fix for affected customers.

## cluster-operator [v0.23.9](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.9)

- Fix a bug in user values migration logic for apps.
