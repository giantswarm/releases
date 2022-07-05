# :zap: Giant Swarm Release v16.5.0 for AWS :zap:

This is a security release featuring the latest version of Kubernetes 1.21 (1.21.14) and all of Giant Swarm applications.

## Change details


### cluster-operator [3.14.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.14.1)

#### Changed
- Update `aws-pod-identity-webhook` app version.



### app-operator [6.0.1](https://github.com/giantswarm/app-operator/releases/tag/v6.0.1)

#### Added
- Add support for Catalogs that define multiple repository mirrors to be used in case some of them are unreachable.
#### Changed
- Only run `PodMonitor` outside of bootstrap mode.



### aws-cni [1.11.2-iptables](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.2-iptables)

Upgraded from version 1.10.1. Please check [upstream changelog](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.2) for details.


### kubernetes [1.21.14](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.14)

Upgraded from 1.21.9.

This release brings bug and security fixes.

Please refer to the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#v1219) for all the details.


### calico [3.21.5](https://github.com/projectcalico/calico/releases/tag/v3.21.5)

Upgraded from 3.12.3.

This upgrade bring security and bug fixes.

Please refer to the [official changelog](https://projectcalico.docs.tigera.io/archive/v3.21/release-notes/#v3215) for all details.


### external-dns [2.15.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.0)

#### Changed
- Update test dependencies and py-helm-charts version to [0.7.0](https://github.com/giantswarm/pytest-helm-charts/blob/master/CHANGELOG.md) ([#173](https://github.com/giantswarm/external-dns-app/pull/173))
- Ignore IRSA annotation for service account when using AWS `external` access.



### chart-operator [2.24.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.1)

#### Changed
- Update `helmclient` to v4.10.1.



### node-exporter [1.13.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.13.0)

#### Changed
- Disable boot partition from the `filesystem` exporter.



### metrics-server [1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set `kubelet-preferred-address-types` to `Hostname` on `AWS`.



### cert-manager [2.15.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.0)

#### Changed
- Upgrade to upstream image [`v1.7.3`](https://github.com/jetstack/cert-manager/releases/tag/v1.7.3) which increases some hard-coded timeouts for certain ACME issuers (ZeroSSL and Sectigo) ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))
- Update kubectl container version to `1.24.2` ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))



### aws-ebs-csi-driver [2.14.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.14.0)

#### Changed
- Remove `imagePullSecrets` from values.yaml
- Bump aws-ebs-csi-driver version to `v1.6.2`.



### coredns [1.10.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.10.0)

#### Added
- Add `app.kubernetes.io/component` on deployments so that management-cluster-admission controller does not complain.



### kube-state-metrics [1.11.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.11.0)

#### Add
- Allow `application.giantswarm.io/team` label.



