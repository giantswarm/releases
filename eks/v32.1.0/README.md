# :zap: Giant Swarm Release v32.1.0 for  :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- cluster-eks from v1.0.0 to v1.1.0
- Flatcar from v4152.2.3 to [v4459.2.3](https://www.flatcar.org/releases/#release-4459.2.3)
- Kubernetes from v1.32.9 to [v1.32.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.11)
- os-tooling from v1.26.1 to v1.26.4

### cluster-eks [v1.0.0...v1.1.0](https://github.com/giantswarm/cluster-eks/compare/v1.0.0...v1.1.0)

#### Fixed

- Fix availability zone handling in managed machine pool template.

### Apps

- aws-ebs-csi-driver from v3.0.5 to v4.1.1
- cilium from v1.3.0 to v1.3.4
- coredns from v1.28.1 to v1.29.1
- network-policies from v0.1.1 to v0.1.3
- teleport-kube-agent from v0.10.6 to v0.10.7

### aws-ebs-csi-driver [v3.0.5...v4.1.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v4.1.1)

#### Added

- Introduce bundle chart architecture with Crossplane IAM resources.
  - Add `aws-ebs-csi-driver-app-bundle` chart that includes:
- Crossplane IAM Role with EBS CSI driver permissions
- Flux HelmRelease to deploy the workload cluster chart
- ConfigMap for values passthrough
  - Bundle chart is installed on the management cluster and deploys the app chart to the workload cluster
  - IAM role uses OIDC federation (IRSA) and reads configuration from `<clusterID>-crossplane-config` ConfigMap
  - Both charts share the same version and are released together

#### Changed

- Refactor crossplane config data retrieval. Fail installation if the ConfigMap can't be found, otherwise the chart was creating invalid IAM roles.
- Change IAM role name for the ebs-csi-driver-controller, to differentiate it from the old one managed by the iam-operator.
- Remove dependency for the cloud-provider-aws in the aws-ebs-csi-driver HelmRelease. That dependency should be set in the bundle HelmRelease by the provider cluster chart
- Update CircleCI configuration to push both app and bundle charts
- Update README with bundle architecture documentation
- Chart: Sync to upstream. ([#338](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/338))
  - Chart: Update AWS EBS CSI Driver from v1.41.0 to v1.51.0.
  - Chart: ⚠️ URGENT: XFS Compatibility Issue - Newly formatted XFS volumes may fail to mount on nodes with older kernels (Amazon Linux 2). Use `node.legacyXFS: true` as workaround.
  - Chart: ⚠️ URGENT: Controller Health Checks - Controller now performs AWS API dry-run checks. Ensure proper IAM permissions and network connectivity.
  - Chart: ⚠️ URGENT: StorageClass Parameter Deprecation - `blockExpress` parameter is deprecated for `io2` volumes (now always uses 256,000 IOPS cap).
  - Chart: Add support for creating instant, point-in-time copies of EBS volumes within the same Availability Zone.
  - Chart: Add `debugLogs` parameter for maximum verbosity logging and debugging.
  - Chart: Add `metadataSources` configuration option for node metadata handling.
  - Chart: Add `disableMutation` parameter for service account mutation control.
  - Chart: Add support for updating node's max attachable volume count via `MutableCSINodeAllocatableCount` feature gate (Kubernetes 1.33+).
  - Chart: Update dependencies including AWS SDK, Prometheus, and various Go modules.
  - Chart: Add missing `enablePrometheusAnnotations` values for controller and node components.
  - Chart: Update sidecar container versions:
- csi-provisioner: v5.2.0 → v5.3.0
- csi-attacher: v4.8.1 → v4.9.0
- csi-snapshotter: v8.2.1 → v8.3.0
- livenessprobe: v2.14.0 → v2.16.0
- csi-resizer: v1.13.2 → v1.14.0
- csi-node-driver-registrar: v2.13.0 → v2.14.0
- volume-modifier-for-k8s: v0.5.1 → v0.8.0
- Configure `gsoci.azurecr.io` as the default container image registry.
- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

#### Fixed

- Fix boolean type of the expansion
- Allow volume expansion by default on gp3

### cilium [v1.3.0...v1.3.4](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.4)

#### Changed

- Upgrade Cilium to [v1.18.6](https://github.com/cilium/cilium/releases/tag/v1.18.6).
- Upgrade Cilium to [v1.18.5](https://github.com/cilium/cilium/releases/tag/v1.18.5).
- Upgrade Cilium to [v1.18.4](https://github.com/cilium/cilium/releases/tag/v1.18.4).
- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### coredns [v1.28.1...v1.29.1](https://github.com/giantswarm/coredns-app/compare/v1.28.1...v1.29.1)

#### Changed

- Update `coredns` image to [1.14.1](https://github.com/coredns/coredns/releases/tag/v1.14.1).
- Update `coredns` image to [1.14.0](https://github.com/coredns/coredns/releases/tag/v1.14.0).
- Update `coredns` image to [1.13.2](https://github.com/coredns/coredns/releases/tag/v1.13.2).
- Update `coredns` image to [1.13.1](https://github.com/coredns/coredns/releases/tag/v1.13.1).

### network-policies [v0.1.1...v0.1.3](https://github.com/giantswarm/network-policies-app/compare/v0.1.1...v0.1.3)

#### Added

- Add support for Kamaji.

#### Fixed

- Fixed broken templating.

### teleport-kube-agent [v0.10.6...v0.10.7](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.6...v0.10.7)

#### Added

- Add `ephemeral-storage` requests and limits to satisfy Kyverno policy `require-emptydir-requests-and-limits`.

#### Changed

- Enable upstream-provided Prometheus PodMonitor to scrape metrics from Teleport Kube Agent pods.
