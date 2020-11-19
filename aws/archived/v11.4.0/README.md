# :zap: Tenant Cluster Release v11.4.0 for AWS :zap:

This release introduces [high availability Kubernetes masters](https://docs.giantswarm.io/basics/ha-masters/),
which means that clusters can have three master nodes in different availability zones
instead of one master node.

Please upgrade [`gsctl`](https://docs.giantswarm.io/reference/gsctl/#install) to the latest
version in order to work with the master node high availability functions.

## Important

- High availability master nodes is the default setting as of this release. If needed,
  clusters can be created with only a single master node instead.
- Single master clusters as of this release can be converted to multi master. However
  a conversion from multi master to single master is not possible.
- Selecting the master node availability zone is no longer supported as of this release.
- Some v5 functions in the [Giant Swarm Rest API](https://docs.giantswarm.io/api/) regarding
  [cluster creation](https://docs.giantswarm.io/api/#operation/addClusterV5) and [fetching
  cluster details](https://docs.giantswarm.io/api/#operation/getClusterV5) are changing.
  The old `master` attribute in the request/response will be deprecated by August 31.
  Please change implementations to use the `master_nodes` attribute instead.
- Specifying the availability zones of master nodes is only supported via the Giant Swarm
  Control Plane API, not the Rest API

Read our [dedicated documentation article](https://docs.giantswarm.io/basics/ha-masters/)
for more details and instructions.

**Note for Solution Engineers:** This release contains an external-dns fix introduced in
[11.3.2](https://github.com/giantswarm/releases/blob/master/aws/v11.3.2/release-notes.md).
It requires manual intervention for cluster upgrades in China to work. When upgrading a
cluster, existing ingress `A+TXT` record sets do not get replaced with `CNAME+TXT` record sets
even when external-dns is configured with CNAMEs as preferred. After upgrading, delete the
ingress `A+TXT` record sets. external-dns will then automatically create `CNAME+TXT` record
sets.

## Change details

### aws-operator [v8.7.0](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.0)

- Several changes to support running clusters with multiple master nodes as well as migrating
  from a single to multiple masters.
- Several improvements regarding the deletion of tenant clusters and related AWS resources.
- Kubernetes master nodes in tenant clusters will now receive the label
  `giantswarm.io/control-plane` in metadata.
- Private subnets for all used availability zones are created by default. This fixes a problem
  with private services of type `LoadBalancer`.
- Fixed Prometheus metrics gathering from calico endpoint.
- The number of calls to the AWS API for collection of Elastic Load Balancer (ELB) details
  has been reduced by adding a local cache, to avoid throttling errors.
- The Docker image now uses alpine v3.12 as a base.

### cluster-operator [v2.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.0)

- Several improvements regarding the deletion of tenant clusters and related AWS resources.
- The Docker image now uses alpine v3.12 as a base.

### cluster-autoscaler 1.16.5 (Giant Swarm app [v1.16.0](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md))

- Version v1.16.5 introduces a new method to read AWS EC2 instance type details from an AWS API.
  Since this API is not reachable from the AWS China regions, the autoscaler is started with the
  `--aws-use-static-instance-list=true` flag.
- Set `scan-interval` to 30 seconds (from 10 seconds) to save resources and reduce AWS API calls.
- Set `scale-down-unneeded-time` to 5 minutes (from the default of 10 minutes) to release unneeded
  nodes earlier.
- Lower `scaleDownUtilizationThreshold` to 0.5.

### chart-operator [v0.13.0](https://github.com/giantswarm/chart-operator/releases/tag/v0.13.0)

Please check the [changelog](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md)
for changes since v0.12.4.

### cert-exporter [v1.2.3](https://github.com/giantswarm/cert-exporter/releases/tag/v1.2.3)

Please check the [changelog](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md)
for changes since v1.2.1.

### net-exporter [v1.8.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.8.0)

Please check the [changelog](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md)
for changes since v1.7.0.

### calico v3.10.4

- Fixes IPv6 rogue router advertisement vulnerability
  [CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-13597).

Complete details for changes since v3.10.1 please check the upstream release notes at
https://docs.projectcalico.org/archive/v3.10/release-notes/

### etcd v3.4.9

Please check the upstream changelogs for details on changes since version 3.3.17.

- https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.4.md
- https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.3.md
