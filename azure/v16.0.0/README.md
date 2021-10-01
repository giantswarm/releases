# :zap: Giant Swarm Release v16.0.0 for Azure :zap:

This is the first Giant Swarm Azure stable release featuring Kubernetes 1.21.

In this release a few of the Azure resources related to the API server load balancer have been renamed to comply
to Cluster API naming. If you have a feature that relies on the naming of Load Balancers, Health Probes, Backend Pools,
or Load Balancing Rules please get in touch with your Solution Engineer. 

With this release the deprecated Azure MSI extensions has been removed from node pools' VMs because it was unused.
If you rely on the MSI extension being present for your workloads please talk with your Solution Engineer.

## Change details


### azure-operator [5.9.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.9.0)

#### Changed

- Use go embed in place of pkger.
- Rename API backend pool to comply with CAPZ.
- Rename API Load Balancing rule to comply with CAPZ.
- Rename API health probe to comply with CAPZ.
- Set `DisableOutputSnat` to true for API Load Balancer Load Balancing Rule to comply with CAPZ.
- Bumped `k8scloudconfig` to support Kubernetes 1.21

#### Fixed

- Ensure Spark CR release version label is updated when upgrading a cluster.

#### Removed

- Remove MSI extension from node pools.
- Remove VPN gateway cleanup code.


### kubernetes [1.21.4](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.4)

Updated from 1.20.8.

Please be aware this is a major release of Kubernetes that brings deprecations and dropped APIs.
Please read the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#whats-new-major-themes) 
before upgrading to ensure your workloads are not affected.

Detailed changes since previous release:

- Changes in Kubernetes [v1.20.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#changelog-since-v1208)
- Changes in Kubernetes [v1.20.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#changelog-since-v1209)
- Changes in Kubernetes [v1.21.0](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#changelog-since-v1200)
- Changes in Kubernetes [v1.21.1](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#changelog-since-v1210)
- Changes in Kubernetes [v1.21.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#changelog-since-v1211)
- Changes in Kubernetes [v1.21.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#changelog-since-v1212)
- Changes in Kubernetes [v1.21.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#changelog-since-v1213)


### containerlinux [2905.2.3](https://www.flatcar-linux.org/releases/#release-2905.2.3)

Updated from 2765.2.6.

This release brings Linux Kernel version 5.10.61, Docker 19.03.15 and openssl 1.1.1l as well as a number of security fixes.

Please refer to [Flatcar's release notes](https://www.flatcar-linux.org/releases/#release-2905.2.3) for all details.


### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

Updated from 1.7.1.

#### Added

- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.

#### Changed

- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.


### chart-operator [2.19.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.19.0)

Updated from 2.18.0.

#### Removed

- Remove `tillermigration` resource now Helm 3 migration is complete.

#### Changed

- Increase memory limit for deploying large charts in workload clusters.


### coredns [1.6.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.0)

Updated from 1.4.1.

#### Changed

- Make `targetCPUUtilizationPercentage` in HPA configurable.
- Update `coredns` to upstream version [1.8.3](https://coredns.io/2021/02/24/coredns-1.8.3-release/).
- Increase maximum replica count to 50 when using horizontal pod autoscaling.


### external-dns [2.6.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.6.0)

Updated from 2.4.0.

#### Added

- Add support for CAPZ clusters by detecting the Azure configuration file location.

#### Changed

- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.


### kube-state-metrics [1.4.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.4.0)

Updated from 1.3.1.

#### Changed

- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.


### metrics-server [1.5.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.5.0)

Updated from 1.3.0.

#### Changed

- Bumped API version for RoleBinding to v1 as it was using a deprecated version (removed in 1.22).


### cluster-autoscaler [1.21.0-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.0-gs1) 

Updated from 1.20.3.

#### Changed

- Updated cluster-autoscaler to version `1.21.0`.
- Use new node selector `node-role.kubernetes.io/master` in place of deprecated one `kubernetes.io/role`.
- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.
- Add `VerticalPodAutoscaler` resource to adjust limits automatically.


### app-operator [5.2.0](https://github.com/giantswarm/app-operator/releases/tag/v5.2.0)

Updated from 4.4.0.

#### Changed

- Reject App CRs with version labels with the legacy `1.0.0` value.
- Validate `.spec.catalog` using Catalog CRs instead of AppCatalog CRs.
- Create `AppCatalogEntry` CRs into the same namespace of Catalog CR.
- Include `chart.keywords`, `chart.description` and `chart.upstreamChartVersion` in `AppCatalogEntry` CRs.
- Create `AppCatalog` CRs from `Catalog` CRs for compatibility with existing app-operator releases.
- Prepare helm values to configuration management.
- Use `Catalog` CRs in `App` controller.
- Reconcile to `Catalog` CRs instead of `AppCatalog`.
- Get `Chart` CRD from the GitHub resources.
- Get metadata constants from k8smetadata library not apiextensions.

#### Fixed

- Fix creating `AppCatalog` CRs in appcatalogsync resource.
- For the chart CR watcher get the kubeconfig secret from the chart-operator app
  CR to avoid hardcoding it.
- Quote namespace in helm templates to handle numeric workload cluster IDs.


### cluster-operator [0.27.2](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.2)

Updated from 0.27.1.

#### Changed

- Use `app-operator-konfigure` configmap for the app-operator per workload cluster.

#### Fixed

- Fixed default value for calico subnet.


### etcd [3.4.16](https://github.com/etcd-io/etcd/releases/tag/v3.4.16)

Updated from 3.14.4.

- Changelog in etcd [v3.14.15](https://github.com/etcd-io/etcd/blob/main/CHANGELOG-3.4.md#v3415-2021-02-26).
- Changelog in etcd [v3.14.16](https://github.com/etcd-io/etcd/blob/main/CHANGELOG-3.4.md#v3416-2021-05-11).
