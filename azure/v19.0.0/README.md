# :zap: Giant Swarm Release v19.0.0 for Azure :zap:

This is the first Azure release featuring Kubernetes 1.24. For increased performance the kubernetes api-server's cpu request are changed to be the half of the **available** CPUs in the VM. Release consists of upgrades to most of components that are listed below. 

Important to note this release adds two new components:
- [k8s-dns-node-cache](https://github.com/giantswarm/k8s-dns-node-cache-app) - improving cluster DNS performance by running a dns caching agent on cluster nodes as a DaemonSet.
- [observability-bundle](https://github.com/giantswarm/observability-bundle) - revised monitoring solution that provides necessary components to enable observability capabilities in a workload cluster. This upgrade can temporairly affect prometheus but the monitoring team is paged immediately to perform manual steps and cleanup. Hence there should be little to no impact.

## Change details


### app-operator [6.5.0](https://github.com/giantswarm/app-operator/releases/tag/v6.5.0)

#### Fixed
- Fix building URLs for OCI Repositories assigned to non-internal `Catalogs`.



### azure-operator [7.0.0](https://github.com/giantswarm/azure-operator/releases/tag/v7.0.0)

#### Changed
- Bump k8scc to [15.4.0](https://github.com/giantswarm/k8scloudconfig/blob/master/CHANGELOG.md#1540---2023-01-11) for k8s 1.24 support.
- Change apiserver's cpu request to be 1/2 of the available CPUs in the VM.



### cert-operator [3.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v3.0.1)

#### Fixed
- Allow running unique and non unique cert-operators in the same namespace.



### cluster-operator [5.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.3.0)

#### Changed
- Enable IRSA by default on v19+ clusters.



### kubernetes [1.24.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.9)

Bumped from version 1.23.9. Please read upstream [release announcement](https://kubernetes.io/blog/2022/05/03/kubernetes-1-24-release-announcement/) for all details.


### containerlinux [3374.2.1](https://www.flatcar-linux.org/releases/#release-3374.2.1)

Upgraded from `3227.2.1`.

Please refer to the [upstream changelog](https://www.flatcar.org/releases/#stable-release) for all details.

### azure-cloud-controller-manager [1.24.6-gs1](https://github.com/giantswarm/azure-cloud-controller-manager-app/releases/tag/v1.24.6-gs1)

#### Changed

- Bump to upstream version v1.24.5.

#### Added
- Add support to make `CPI` run on `CAPZ` based clusters.


### azure-cloud-node-manager [1.24.6-gs1](https://github.com/giantswarm/azure-cloud-node-manager-app/releases/tag/v1.24.6-gs1)

### Changed

- Bump to upstream version v1.24.5.

### Fixed

- Fix resources in yaml. it was defined twice



### azuredisk-csi-driver [1.25.2-gs1](https://github.com/giantswarm/azuredisk-csi-driver-app/releases/tag/v1.25.2-gs1)

### Changed

- Bumped azuredisk-csi to upstream version 1.25.0.


### cert-exporter [2.3.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.3.1)

#### Fixed
- Allow eviction for cert-exporter-deployment.



### chart-operator [2.33.2](https://github.com/giantswarm/chart-operator/releases/tag/v2.33.2)

#### Changed

- Add support for new control-plane label in k8s 1.24.

#### Added

- Support for running behind a proxy.
- Add support to run in private cloud clusters, which cannot provide any working externalDNSIP.

### coredns [1.13.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.13.0)

#### Added
- Possibility to set scale down stabilizationWindowSeconds behaviour

#### Changed
- Move nodeselector `label:value` to values.yaml to allow customizing it for CAPZ
- Add toleration for `node-role.kubernetes.io/control-plane` to masters instance



### etcd-kubernetes-resources-count-exporter [1.0.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.0.0)

#### Changed
- Push to azure and AWS app collection.
- Run in kube-system by default.
- Change default registry to docker.io.



### external-dns [2.22.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.22.0)

#### Added
- Support for running behind a proxy.
- Add nodeSelector, affinity, topologySpreadContraints and tolerations values to align to upstream ([223](https://github.com/giantswarm/external-dns-app/pull/223))

#### Changed 
- ServiceAccount: Align to upstream ([#222](https://github.com/giantswarm/external-dns-app/pull/222)).
- Labels: Add labels from values.
- Allow overrides of service account annotations ([#221](https://github.com/giantswarm/external-dns-app/pull/221)).



### kube-state-metrics [1.14.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.14.1)

#### Added

- Allow topology.kubernetes.io/region & topology.kubernetes.io/zone labels.

#### Changed

- Upgrade kube-state-metrics to 2.6.0


### metrics-server [2.0.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.0.0)

#### Changed
- Switch to HA setup.



### net-exporter [1.13.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.13.0)

#### Added
- Add helm chart values schema
#### Changed
- Update to Go 1.18
- Update github.com/giantswarm/k8sclient to v7.0.1
- Update github.com/giantswarm/micrologger to v1.0.0
- Update github.com/miekg/dns to v1.1.50
- Update k8s.io deps to v0.26.0
- Update docker-kubectl to 1.25.4



### node-exporter [1.15.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.15.0)

#### Added

- Enable ethtool collector.

#### Fixed

- Fix collector systemd
- Fix duplicate scrapping by GiantSwarm Prometheus


### cluster-autoscaler [1.24.0-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.0-gs1)

#### Changed
- Update cluster-autoscaler to version `1.24.0`.



### azure-scheduled-events [0.8.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.8.0)

#### Changed
- Bumped dependencies.
- Switched to go 1.18. 



### vertical-pod-autoscaler [2.5.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.5.2)

### Changed

- Increased memory limits for updater, recommender and admissionController
- Using custom docker image with openssl to fix vpa-certgen job



### observability-bundle [0.1.8](https://github.com/giantswarm/observability-bundle/releases/tag/v0.1.8)

#### Changed
- Upgrade `prometheus-agent` from 0.1.6 to 0.1.7.


### k8s-dns-node-cache [1.1.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v1.1.0)

#### Added
- First release featuring upstream version 1.21.4.

