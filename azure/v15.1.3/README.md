# :zap: Giant Swarm Release v15.1.3 for Azure :zap:

This is a patch release that brings the newest 1.20 Kubernetes version as well as updated version of
all the Giant Swarm's softwares.

It also improves the upgrade process when Azure Disk PVCs are used.
With this release, the node draining process should be much more reliable and the general impact
on the workloads should be as good as it can be.

## Change details


### kubernetes [1.20.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.12)

#### API Change
- Kube-apiserver: Fixes handling of CRD schemas containing literal null values in enums (#104990, @liggitt) [SIG API Machinery, Apps and Network]
#### Bug or Regression
- Detach volumes from vSphere nodes not tracked by attach-detach controller (#104910, @gnufied) [SIG Cloud Provider and Storage]
- Fix: consolidate logs for instance not found error (#105364, @nilo19) [SIG Cloud Provider]
- Fix: ignore not a VMSS error for VMAS nodes in EnsureBackendPoolDeleted. (#105404, @ialidzhikov) [SIG Cloud Provider]
- Fix: ignore the case when updating Azure tags (#104687, @nilo19) [SIG Cloud Provider]
- Revert PR #102925 which introduced unexpected scheduling behavior based on balanced resource allocation (#105239, @damemi) [SIG Scheduling]
- Updates golang.org/x/text to v0.3.6 to fix CVE-2020-28852 (#102602, @jonesbr17) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Node]
#### Other (Cleanup or Flake)
- Allow CSI drivers to just run offline expansion tests (#102665, @gnufied) [SIG Storage and Testing]


### azure-operator [5.8.2](https://github.com/giantswarm/azure-operator/releases/tag/v5.8.2)

### Changed

- Delegate Storage account type selection for master VM's disks to Azure API.
- Separate the drain and node deletion phases during node pool upgrades to avoid stuck disks.

### Fixed

- During an upgrade, fixed the detection of a master node being upgraded to wait before upgrading node pools.


### cert-operator [1.2.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.2.0)

#### Changed
- Introducing `v1alpha3` CR's.
#### Added
- Add check to ensure that the `Cluster` resource is in the same namespace as the `certConfig` before creating the secret there.



### containerlinux [2905.2.5](https://www.flatcar-linux.org/releases/#release-2905.2.5)

This release includes the fix for the following Linux Kernel CVE advisories: CVE-2021-41073, CVE-2020-16119, CVE-2021-3753, CVE-2021-3739, CVE-2021-40490.



### cert-exporter [2.0.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.0.0)

#### Changed
- Export presence of `giantswarm.io/service-type: managed` label in cert-manager `Issuer` and `ClusterIssuer` CR referenced by `Certificate` CR `issuerRef` spec field to `cert_exporter_certificate_cr_not_after` metric as `managed_issuer` label.
- Add `--monitor-files` and `--monitor-secrets` flags.
- Add Deployment to helm chart to avoid exporting secrets and certificate metrics from DaemonSets.
- Build container image using retagged giantswarm alpine.
- Run as non-root inside container.



### chart-operator [2.19.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.19.1)

#### Changed
- Deployment `hostNetwork` is enabled or not depending on `chartOperator.cni.install` value.



### net-exporter [1.10.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.3)

#### Changed
- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.



### node-exporter [1.8.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.8.0)

#### Changed
- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.



### cluster-autoscaler [1.23.0](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.23.0)

#### Changed
- Use new node selector `node-role.kubernetes.io/master` in place of deprecated one `kubernetes.io/role`.


### external-dns [2.6.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.6.0)

Updated from 2.4.0.

#### Added

- Add support for CAPZ clusters by detecting the Azure configuration file location.

#### Changed

- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.

