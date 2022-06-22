# :zap: Giant Swarm Release v17.1.0 for Azure :zap:

This is a maintenance azure release that provides the latest Kubernetes 1.22 version as well as the latest version of all giantswarm releases.
It also uses etcd 3.5 for improved performance and reliability.

Starting from this version it is possible to enable the automated rotation of the encryption key for Kubernetes secrets. This feature is disabled by default.
If you are interested in enabling it please ask your Solution engineer. 

## Change details


### app-operator [6.0.1](https://github.com/giantswarm/app-operator/releases/tag/v6.0.1)

Upgraded from 5.10.2.

This upgrade brings a bunch of bug fixes, including one to better handle the "app bundle" use case (used for example by the Giant Swarm security stack).

For detailed changelog please refer to the [changelog](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md).



### azure-operator [5.21.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.21.0)

Upgraded from 5.17.0.

This upgrade brings a lot of improvement in the bootstrap and upgrade processes and it fixes a few minor bugs.

It also includes a migration process for `MachinePool` and `AzureMachinPool` CRs from the old `experimental` api group to the new stable one. 

For detailed changelog please refer to the [changelog](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md).



### cluster-operator [4.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.3.0)

Upgraded from 3.12.0.

This upgrade brings a few bug fixes as well as support for automated rotation of Secret encryption keys.

For detailed changelog please refer to the [changelog](https://github.com/giantswarm/cluster-operator/blob/master/CHANGELOG.md).



### cert-operator [2.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v2.0.1)

Upgraded from 1.3.0.

This upgrade brings updated dependencies in the golang binary to address security issues in earlier version.

For detailed changelog please refer to the [changelog](https://github.com/giantswarm/cert-operator/blob/master/CHANGELOG.md).


### containerlinux [3139.2.2](https://www.flatcar-linux.org/releases/#release-3139.2.2)

Upgraded from 3033.2.2.

This upgrade brings fixes for a lot of security issues in all main operating system components, including Linux, golang, containerd and openssl. 

Please refer to the [official changelog](https://www.flatcar.org/releases/#release-3139.2.2) for all details.



### calico [3.21.5](https://github.com/projectcalico/calico/releases/tag/v3.21.5)

Upgraded from 3.12.3.

This upgrade bring security and bug fixes.

Please refer to the [official changelog](https://projectcalico.docs.tigera.io/archive/v3.21/release-notes/#v3215) for all details.


### etcd [3.5.4](https://github.com/etcd-io/etcd/releases/tag/v3.5.4)

Upgraded from 3.4.18.

This is a minor release bump, bringing several security and bug fixes and important performance improvements.

Please refer to the [official changelog](https://github.com/etcd-io/etcd/blob/main/CHANGELOG/CHANGELOG-3.5.md#v354-2022-04-24) and [the announcement blog post](https://etcd.io/blog/2021/announcing-etcd-3.5/) for all details.


### kubernetes [1.22.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.10)

Upgraded from 1.22.6.

This release brings bug and security fixes.

Please refer to the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.22.md#v12210) for all the details.



### cert-exporter [2.2.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.2.0)

Upgraded from 2.0.1.

This release brings improved reliability to the exporter.

Please refer to the [changelog](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md) for all the details.



### chart-operator [2.24.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0)

Upgraded from 2.20.1.

This release brings improved reliability to the operator.

Please refer to the [changelog](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md) for all the details.


### coredns [1.10.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.10.1)

Upgraded from 1.8.0.

This upgrade provides coredns version 1.8.7 as well as improvements in the helm chart.

Please refer to the [changelog](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md) for all the details.



### external-dns [2.14.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.14.0)

Upgraded from 2.9.0.

This upgrade provides external-dns version 0.11.0 as well as improvements in the helm chart.

Please refer to the [changelog](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md) for all the details.


### cluster-autoscaler [1.22.2-gs6](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs6)

Upgraded from 1.22.2-gs4.

This upgrade provides an improved helm chart.

Please refer to the [changelog](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md) for all the details.



### metrics-server [1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

Upgraded from 1.5.0.

This upgrade provides metrics-server version 0.5.2 as well as improvements in the helm chart.

Please refer to the [changelog](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md) for all the details.


### net-exporter [1.12.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.12.0)

Upgraded from 1.11.0.

This upgrade provides an improved helm chart.

Please refer to the [changelog](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md) for all the details.



### node-exporter [1.12.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.12.0)

Upgraded from 1.8.0.

This upgrade provides node-exporter 1.3.1 and enables the `diskstats` exporter to expose node IO metrics.

Please refer to the [changelog](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md) for all the details.


### azure-scheduled-events [0.7.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.7.0)

Upgraded from 0.6.1.

This upgrade provides an improved helm chart.

Please refer to the [changelog](https://github.com/giantswarm/azure-scheduled-events/blob/master/CHANGELOG.md) for all the details.


### vertical-pod-autoscaler [2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

Upgraded from 2.1.1.

This upgrade provides vertical-pod-autoscaler 0.10.0 including an internal fix to circumvent a bug in containerd that prevented VPA to work properly.

Please refer to the [changelog](https://github.com/giantswarm/vertical-pod-autoscaler-app/blob/master/CHANGELOG.md) for all the details.


### vertical-pod-autoscaler-crd [1.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v1.0.1)

Upgraded from 1.0.0.

#### Added
- Add cluster singleton restriction so app can only be installed once.



### kube-state-metrics [1.10.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.10.0)

Upgraded from 1.7.0.

This upgrade provides an improved helm chart.

Please refer to the [changelog](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md) for all the details.



