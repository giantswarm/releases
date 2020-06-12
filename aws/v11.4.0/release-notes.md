## :zap: Giant Swarm Release v11.4.0 for AWS :zap:

This release introduces master node high availability, which means that clusters will have three master nodes by default in three different availability zones (or ).

- Multiple master nodes are only supported in AWS regions having at least three availability zones. This excludes China/Beijing (`cn-north-1`).

**Note for SEs:** This release contains an external-dns fix introduced in [11.3.2](https://github.com/giantswarm/releases/blob/master/aws/v11.3.2/release-notes.md). It requires manual intervention for cluster upgrades in China to work. When upgrading a cluster, existing ingress A+TXT record sets do not get replaced with CNAME+TXT record sets even when external-dns is configured with CNAMEs as preferred. After upgrading, delete the ingress A+TXT record sets. external-dns will then automatically create CNAME+TXT record sets.

**Note for future 11.3.x releases:** Until upstream external-dns issue is fixed, please include this note and the one above in all future 11.3.x releases.

## Change details

### aws-operator vX.X.X

- Several changes to support running clusters with multiple master nodes as well as migrating from a single to multiple masters.
- Several improvements regarding the deletion of  tenant clusters and related AWS resources.
- Kubernetes master nodes in tenant clusters will now receive the label `giantswarm.io/control-plane` in metadata.
- Fixed Prometheus metrics gathering from calico endpoint.
- The Docker image now uses alpine v3.12 as a base.

### cluster-operator vX.X.X

- Several improvements regarding the deletion of  tenant clusters and related AWS resources.
- The Docker image now uses alpine v3.12 as a base.

### cluster-autoscaler 1.16.5 (Giant Swarm app [v1.16.0](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md))

- Version v1.16.5 introduces a new method to read AWS EC2 instance type details from an AWS API. Since this API is not reachable from the AWS China regions, the autoscaler is started with the `--aws-use-static-instance-list=true` flag.
- Set `scan-interval` to 30 seconds (from 10 seconds) to save resources and reduce AWS API calls.
- Set `scale-down-unneeded-time` to 5 minutes (from the default of 10 minutes) to release unneeded nodes earlier.
- Lower `scaleDownUtilizationThreshold` to 0.5.

### chart-operator 0.13.0

Please check the [changelog](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md) for changes since v0.12.4.

### cert-exporter 1.2.3

Please check the [changelog](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md) for changes since v1.2.1.

### net-exporter v1.8.0

Please check the [changelog](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md) for changes since v1.7.0.

### calico v3.10.4

- Fixes IPv6 rogue router advertisement vulnerability [CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-13597).

Complete details for changes since v3.10.1 please check the upstream release notes at https://docs.projectcalico.org/archive/v3.10/release-notes/

### etcd v3.4.9

Please check the upstream changelogs for details on changes since version 3.3.17.

- https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.4.md
- https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.3.md
