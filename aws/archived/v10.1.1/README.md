# :zap: Tenant Cluster Release 10.1.1 for AWS :zap:

This Node Pools release fixes a problem with CoreDNS in clusters using custom
cluster IP ranges. Additionally, CPU limits have been removed from multiple
default apps to increase their reliability.

## Component changes:

## cluster-operator 2.0.1

- Added additional settings for CoreDNS to cluster configmap.
- Fix cluster status conditions to be reconciled upon cluster creation.

## cert-exporter (GS [v1.2.1](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#121-2019-12-24))

- Removed CPU limits.

## cert-manager (GS [v1.0.3](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#v103-2020-01-03))

- Removed CPU limits.

## chart-operator 0.11.2

- Removed CPU limits.

## cluster-autoscaler (GS [v1.1.2](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md#v112-2020-01-03))

- Removed CPU limits.

## external-dns (GS [v1.1.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#v110))

- Added support AWS SDK configuration with explicit credentials.
- Removed CPU limits.

## kiam (GS [v1.0.2](https://github.com/giantswarm/kiam-app/blob/master/CHANGELOG.md#v102-2020-01-04))

- Removed CPU limits.

## node-exporter (GS [v1.2.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#120-2020-01-08))

- Removed CPU limits.
