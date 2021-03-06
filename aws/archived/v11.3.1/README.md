## :zap: Tenant Cluster Release v11.3.1 for AWS :zap:

This release provides fixes for a race condition seen in some upgrades to v11.2.x when allocating the IP address for the master node.

In addition, we add support for disabling [external source network address translation (SNAT)](https://docs.aws.amazon.com/eks/latest/userguide/external-snat.html) for pod IP addresses via the Control Plane Kubernetes API, plus we pick up some minor component updates.

## aws-operator [v8.6.1](https://github.com/giantswarm/aws-operator/releases/tag/v8.6.1)

- Prevent an IP address collision when upgrading to a release managed by this aws-operator version
- Add support for an external SNAT setting on the `AWSCluster` custom resource [#2426](https://github.com/giantswarm/aws-operator/pull/2426)
- Use operatorkit v1.0.0 [#2449](https://github.com/giantswarm/aws-operator/pull/2449)

## cluster-operator [v2.2.0](https://github.com/giantswarm/cluster-operator/releases/tag/v2.2.0)

- Generate etcd certificates in preparation of master nodes high-availability [#1032](https://github.com/giantswarm/cluster-operator/pull/1032)
- Add pod CIDR service implementation using local caching, to ensure consistent state and reduce the number of Kubernetes API requests [#1054](https://github.com/giantswarm/cluster-operator/pull/1054)
