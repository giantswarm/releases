# :zap: Giant Swarm Release v15.0.0-beta1 for Azure :zap:

This is the first Giant Swarm Azure release featuring Kubernetes 1.20.

This release includes many improvements to make the upgrade process even more reliable, with many improvements
specifically related to Spot Instances.

This release also uses Vnet peering in place of VPN Gateway to connect the Workload Clusters with the Management Cluster.
There should be no impact in terms of connectivity from the workload point of view, but if you are using vnet peering
in your tenant clusters please get in touch with your solution engineer.

## Change details


### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

Updated from version 3.2.1.

Highlight of changes:

#### Fixed

- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.

#### Changed

- Updated Helm to v3.5.3.

#### Added

- Enable Vertical Pod Autoscaler.

### azure-operator [5.6.1-k8s-1-20](https://github.com/giantswarm/azure-operator/releases/tag/v5.6.1-k8s-1-20)

Updated from version 5.5.3.

Highlight of changes:

#### Changed

- Replace VPN Gateway with VNet Peering.
- Update OperatorKit to v4.3.1 to drop usage of self-link which is not supported in k8s 1.20 anymore. 
- Avoid creating too many worker nodes at the same time when upgrading node pools.
- Don't wait for new workers to be up during spot instances node pools upgrades.

#### Fixed

- When deleting a node pool, also delete the VMSS role assignment.

#### Removed

- Support for single tenant BYOC credentials (warning: the operator will error at startup if any organization credentials is not multi tenant).

### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

Updated from 0.1.0.

Highlight of changes:

#### Changed

- Update Kubernetes dependencies to 1.18 versions.

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.

#### Added

- Add network policy resource.
- Added lookup for nodepool clusters in other namespaces than default.

#### Removed

- Stop using the VersionBundle version.

### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

Updated from 0.23.22.

Highlight of changes:

#### Changed
- Dropped ensuring cluster CRDs from controllers.
- Adjust helm chart to be used with config-controller.
- Migrate to Go modules.
- Update certs package to v2.0.0.
- Refactor to use slightly newer dependency versions.
- Align version bundle version and project version.
- Remove VersionBundle version from CertConfigs and add the cert-operator.giantswarm.io/version label. This change requires using cert-operator 1.0.0 or later.

#### Added

- Assign app catalog name from the component in release CR.
- Create app CR for per cluster app-operator instance.
- Add appfinalizer resource to remove finalizers from workload cluster app CRs.

#### Fixed

- Add AllowedLabels to clusterconfigmap resource to prevent unnecessary updates.

### kubernetes [1.20.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.6)

Updated from 1.19.10.

Please refer to the [official release announcement](https://kubernetes.io/blog/2020/12/08/kubernetes-1-20-release-announcement/) for details about the 1.20 Kubernetes version
and to the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md) for details about
changes in the patch releases up to 1.20.6.

### containerlinux [2765.2.3](https://www.flatcar-linux.org/releases/#release-2765.2.3)

Updated from 2765.2.2

**Security fixes**

*   Linux ([CVE-2021-28964](https://nvd.nist.gov/vuln/detail/CVE-2021-28964), [CVE-2021-28972](https://nvd.nist.gov/vuln/detail/CVE-2021-28972), [CVE-2021-28971](https://nvd.nist.gov/vuln/detail/CVE-2021-28971), [CVE-2021-28951](https://nvd.nist.gov/vuln/detail/CVE-2021-28951), [CVE-2021-28952](https://nvd.nist.gov/vuln/detail/CVE-2021-28952), [CVE-2021-29266](https://nvd.nist.gov/vuln/detail/CVE-2021-29266), [CVE-2021-28688](https://nvd.nist.gov/vuln/detail/CVE-2021-28688), [CVE-2021-29264](https://nvd.nist.gov/vuln/detail/CVE-2021-29264), [CVE-2021-29649](https://nvd.nist.gov/vuln/detail/CVE-2021-29649), [CVE-2021-29650](https://nvd.nist.gov/vuln/detail/CVE-2021-29650), [CVE-2021-29646](https://nvd.nist.gov/vuln/detail/CVE-2021-29646), [CVE-2021-29647](https://nvd.nist.gov/vuln/detail/CVE-2021-29647), [CVE-2021-29154](https://nvd.nist.gov/vuln/detail/CVE-2021-29154), [CVE-2021-29155](https://nvd.nist.gov/vuln/detail/CVE-2021-29155), [CVE-2021-23133](https://nvd.nist.gov/vuln/detail/CVE-2021-23133))

**Bug fixes**

*   Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))

**Updates**

*   Linux ([5.10.32](https://lwn.net/Articles/853762/))


### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

Updated from 1.6.0.

#### Changed
- Set docker.io as the default registry

### chart-operator [2.14.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.14.0)

Updated from 2.12.0.

Highlights of changes:

#### Fixed

- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.

#### Changed
- giantswarm-critical PriorityClass only managed when E2E.
- Cancel the release resource when the manifest object already exists.
- Cancel the release resource when helm returns an unknown error.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

Updated from 1.3.0.

#### Changed
- Set docker.io as the default registry



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

Updated from 1.2.1.

#### Added

- Added new configuration value `extraArgs`.

#### Changed

- Set docker.io as the default registry

### net-exporter [1.10.1](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.1)

Updated from 1.9.2.

Highlights of changes:

#### Changed

- Set docker.io as the default registry.
- Update kubectl image to v1.18.8.
- Add label selector for pods to help lower memory usage.
- Allow to customize dns service.

### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

Updated from 1.7.1.

#### Changed
- Set docker.io as the default registry



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

Updated from 1.2.0.

#### Changed
- Set docker.io as the default registry
- Update coredns to upstream version 1.8.0.


### cluster-autoscaler [1.20.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.2)

Updated from version 1.19.10.

#### Changed
- Update cluster-autoscaler to version 1.20.0.
- Set docker.io as the default registry



