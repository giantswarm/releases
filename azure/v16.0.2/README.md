# :zap: Giant Swarm Release v16.0.2 for Azure :zap:

This is a patch release that brings the newest 1.21 Kubernetes version as well as updated version of
all the Giant Swarm's softwares.

It also improves the upgrade process when Azure Disk PVCs are used.
With this release, the node draining process should be much more reliable and the general impact
on the workloads should be as good as it can be.

## Change details


### kubernetes [1.21.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.6)

#### API Change
- Kube-apiserver: Fixes handling of CRD schemas containing literal null values in enums (#104989, @liggitt) [SIG API Machinery, Apps and Network]
#### Feature
- Kubernetes is now built with Golang 1.16.9 (#105672, @cpanato) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Failing Test
- Backport: e2e.test now uses the secure port to retrieve metrics data to ensure compatibility with 1.21 and 1.22 (for upgrade tests) (#104328, @pohly) [SIG API Machinery, Instrumentation, Storage and Testing]
#### Bug or Regression
- Aggregate errors when putting vmss
  fix: consolidate logs for instance not found error (#105365, @nilo19) [SIG Cloud Provider]
- Allow CSI drivers to just run offline expansion tests (#102740, @gnufied) [SIG Storage and Testing]
- Bump klog to v2.9.0, fixing log lines that render as byte arrays (#105407, @ehashman) [SIG Architecture, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Fix overriding logs files on reboot. (#105614, @rphillips) [SIG Node]
- Fix winkernel kube-proxy to only use dual stack when host and networking supports it (#101047, @jsturtevant) [SIG Network and Windows]
- Fix: ignore not a VMSS error for VMAS nodes in EnsureBackendPoolDeleted. (#105402, @ialidzhikov) [SIG Cloud Provider]
- Fix: ignore the case when updating Azure tags (#104686, @nilo19) [SIG Cloud Provider]
- Revert PR #102925 which introduced unexpected scheduling behavior based on balanced resource allocation (#105238, @damemi) [SIG Scheduling]
- Updates golang.org/x/text to v0.3.6 to fix CVE-2020-28852 (#102601, @jonesbr17) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]


### azure-operator [5.10.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.10.0)

### Changed

- Delegate Storage account type selection for master VM's disks to Azure API.
- Separate the drain and node deletion phases during node pool upgrades to avoid stuck disks.

### Fixed

- During an upgrade, fixed the detection of a master node being upgraded to wait before upgrading node pools.


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



### cluster-autoscaler [1.21.0-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.0-gs2)

#### Fixed

- Fix RBAC for cluster autoscaler 1.21.


