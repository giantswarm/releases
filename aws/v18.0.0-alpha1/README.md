# :zap: Giant Swarm Release v18.0.0-alpha1 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [13.0.0-alpha2](https://github.com/giantswarm/aws-operator/releases/tag/v13.0.0-alpha2)

Upgraded from 11.16.0.

Please refer to the [changelog](https://github.com/giantswarm/aws-operator/blob/master/CHANGELOG.md) for all details.


### kubernetes [1.23.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.9)

Upgraded from version 1.22.

This upgrade brings lots of changes, security and bug fixes and deprecations.

Please refer to the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.23.md) for all details.



### etcd [3.5.4](https://github.com/etcd-io/etcd/releases/tag/v3.5.4)

Upgraded from 3.4.18.

This is a minor release bump, bringing several security and bug fixes and important performance improvements.

Please refer to the [official changelog](https://github.com/etcd-io/etcd/blob/main/CHANGELOG/CHANGELOG-3.5.md#v354-2022-04-24) and [the announcement blog post](https://etcd.io/blog/2021/announcing-etcd-3.5/) for all details.


### app-operator [6.3.0](https://github.com/giantswarm/app-operator/releases/tag/v6.3.0)

#### Added
- If no userconfig configmap or secret reference is specified but one is found following the default naming convention (`*-user-values` / `*-user-secrets`) then the App resource is updated to reference the found configmap/secret.



### cluster-operator [4.4.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.4.0)

#### Changed
- Set `chartOperator.cni.install` to true to allow installing CNI as app.



### aws-ebs-csi-driver [2.16.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.16.1)

#### Fixed
- Changing controller `httpEndpoint` to `8610` because of overlapping ports.



### coredns [1.11.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.11.0)

#### Changed
- Update `coredns` to upstream version [1.9.3](https://coredns.io/2022/05/27/coredns-1.9.3-release/).



### cluster-autoscaler [1.23.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/1.23.1)

Upgraded from 1.22.2-gs7 to support newer kubernetes version.

Please refer to the [changelog](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md).



### metrics-server [1.8.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.8.0)

#### Changed
- Updated metrics-server version to 0.6.1.



### vertical-pod-autoscaler [2.4.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.2)

### Changed

- Change default webhook timeout to 5 seconds

### Fixed

- Correct selector in admission controller PDB


### aws-cloud-controller-manager [1.23.1-gs2](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.23.2-gs2)

This app was not present in previous release.


### cilium [0.2.6](https://github.com/giantswarm/cilium-app/releases/tag/v0.2.6)

Cilium is now the CNI in place of calico and AWS-cni.


### chart-operator [2.26.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.26.0)

#### Changed
- Use `127.0.0.1` as KUBERNETES_SERVICE_HOST when `bootstrapMode` is enabled.



