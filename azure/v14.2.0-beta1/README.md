# :zap: Giant Swarm Release v14.2.0-beta1 for Azure :zap:

This is a beta release for Azure that replaces the previously used `VPN gateway` with `VNet peering`, cutting cluster deployment time
by several minutes. Most workload clusters should be up and running in about 10 minutes from this release on.

Being a `beta` release, you are free to deploy new clusters as well as update existing ones, but please keep in mind that it is discouraged
to use it for production workload clusters.

For further information get in touch with your solution engineer.

## Change details


### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

#### Changed
- Dropped ensuring cluster CRDs from controllers.



### azure-operator [5.6.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.6.0)

#### Changed
- Replace VPN Gateway with VNet Peering.
- Update OperatorKit to `v4.3.1` to drop usage of self-link which is not supported in k8s 1.20 anymore.
#### Removed
- Support for single tenant BYOC credentials (warning: the operator will error at startup if any organization credentials is not multi tenant).



### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



### net-exporter [1.9.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.3)

#### Changed
- Set docker.io as the default registry
- Update kubectl image to v1.18.8.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.2.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.2.2)

#### Changed
- Set docker.io as the default registry



### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed
- Set docker.io as the default registry



### chart-operator [2.13.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.13.1)

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.



### cluster-autoscaler [1.19.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.2)

#### Changed
- Set docker.io as the default registry



