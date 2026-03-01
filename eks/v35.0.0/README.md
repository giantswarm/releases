# :zap: Giant Swarm Release v35.0.0 for EKS :zap:

<< Add description here >>

## Changes compared to v34.0.0

### Components

- cluster-eks from v1.2.1 to v1.3.0
- cluster from v4.0.2 to v5.3.0
- Flatcar from v4459.2.3 to [v4459.2.4](https://www.flatcar.org/releases/#release-4459.2.4)
- Kubernetes from v1.34.4 to [v1.35.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.2)

### cluster-eks [v1.2.1...v1.3.0](https://github.com/giantswarm/cluster-eks/compare/v1.2.1...v1.3.0)

#### Changed

- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

### cluster [v4.0.2...v5.3.0](https://github.com/giantswarm/cluster/compare/v4.0.2...v5.3.0)

#### Added

- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.
- Add `insecure` flag to containerd mirrors to configure them as http instead of hard coding https.
- Add `debug.level` flag to containerd mirrors to configure containerd logging verbosity.
- Add `overridePath` flag to containerd mirrors to configure containerd `override_path` flag.
- Add support for Kubernetes Structured Authentication Configuration for OIDC providers. This feature allows configuring multiple OIDC issuers and is supported on **Kubernetes 1.34+**. The feature is disabled by default and can be enabled via `global.controlPlane.oidc.structuredAuthentication.enabled`. When enabled, the API server uses `AuthenticationConfiguration` instead of legacy `--oidc-*` flags. **Note:** Existing OIDC configuration (legacy flags) is automatically migrated and configured as the first issuer in the new structure, ensuring a seamless transition. This feature also adds support for custom CEL expressions for claim mappings and claim validation rules.
- Add `priority-classes` app to deploy curated Giant Swarm priority classes.
- Add `MachineHealthCheck` for machine pool worker nodes (requires "machine pool machines" feature in the CAPI infrastructure provider). This can be turned on per node pool (default off). Use `maxUnhealthy=20%` as default (for control plane nodes, the default remains at `40%`).
- Conditionally add node-problem-detector-app if specific conditions are enabled in a node pool's machine health check properties.
- Enable `MutableCSINodeAllocatableCount` feature gate on all v1.33+ clusters by default.
- Add `minKubernetesVersion` field to feature gate definitions for better version control.
- Chart: Add CDI support for containerd.
- `NodeRestriction` admission plugin is now enabled by default on all clusters.
- `OwnerReferencesPermissionEnforcement` admission plugin is now enabled by default on all clusters.
- Expose `rollingUpdate` values in the node pools to configure the `MachineDeployments`.
- Chart: Add feature gate support for scheduler and kubelet.

#### Changed

- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.
- Control Plane: Move `node-cidr-mask-size` patch out of `enablePriorityAndFairness` conditional block. ([#741](https://github.com/giantswarm/cluster/pull/741))
- Apps: Make Teleport Kube Agent depend on Prometheus Operator CRD. ([#733](https://github.com/giantswarm/cluster/pull/733))
- Chart: Update sandbox image to v3.10.1. ([#734](https://github.com/giantswarm/cluster/pull/734))
- Chart: Render `cloud-config` flag for Kubernetes < v1.34.0 only. ([#736](https://github.com/giantswarm/cluster/pull/736))
- Chart: Always render `cloud-provider` flag. ([#738](https://github.com/giantswarm/cluster/pull/738))
- Chart: Update sandbox image to v3.10. ([#731](https://github.com/giantswarm/cluster/pull/731))
- Refactor containerd configuration to use `config_path` (`/etc/containerd/certs.d`) and `hosts.toml` for registry mirrors, ensuring proper fallback order (local cache -> mirrors -> upstream).
- Move containerd registry authentication to `hosts.toml` headers, as `registry.configs` is ignored when `config_path` is enabled.
- Helpers: Fix `cluster.app.in-release`.
- Add required install values to a `required.yaml` file and update values schema.
- **Breaking**: The `MutableCSINodeAllocatableCount` feature gate is not recommended for use with Kubernetes v1.33+ clusters as it may cause compatibility issues.
- Update HelmRelease ApiVersion to from v2beta1 to v2.
- Chart: Render `cloud-config` flag for Kubernetes < v1.33.0 only.

#### Fixed

- Cleanup job now also deletes HelmChart CRs to prevent leftover resources when suspending HelmReleases during cluster deletion.

#### Removed

- Remove helm `Job` that cleans up `HelmReleases`. This was needed because we were letting Helm delete the infra cluster and control plane Custom Resources, instead of letting CAPI controllers handle the deletion. This has been fixed, so the `Job` is no longer required.

### Apps

- aws-ebs-csi-driver from v4.1.1 to v4.1.2
- aws-ebs-csi-driver-servicemonitors from v0.1.0 to v0.1.2
- cert-exporter from v2.9.15 to v2.10.0
- cert-manager from v3.9.4 to v3.11.0
- chart-operator-extensions from v1.1.2 to v1.1.3
- cilium from v1.3.4 to v1.4.1
- cilium-servicemonitors from v0.1.3 to v0.1.4
- cluster-autoscaler from v1.34.1-1 to v1.24.3
- coredns-extensions from v0.1.2 to v0.1.3
- k8s-dns-node-cache from v2.9.1 to v2.9.2
- karpenter from v1.4.0 to v2.2.0
- karpenter-taint-remover from v1.0.1 to v1.0.2
- metrics-server from v2.7.0 to v2.8.0
- net-exporter from v1.23.0 to v1.23.1
- node-exporter from v1.20.10 to v1.20.11
- observability-bundle from v2.5.0 to v2.8.0
- observability-policies from v0.0.3 to v0.0.4
- priority-classes from v0.3.0 to v0.3.1
- prometheus-blackbox-exporter from v0.5.0 to v0.7.0
- security-bundle from v1.16.1 to v1.17.0
- vertical-pod-autoscaler from v6.1.1 to v6.1.2
- vertical-pod-autoscaler-crd from v4.1.1 to v4.1.2

### aws-ebs-csi-driver [v4.1.1...v4.1.2](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.1...v4.1.2)

#### Changed

- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.

### aws-ebs-csi-driver-servicemonitors [v0.1.0...v0.1.2](https://github.com/giantswarm/aws-ebs-csi-driver-servicemonitors-app/compare/v0.1.0...v0.1.2)

#### Changed

- Migrate to App Build Suite (ABS).

#### Fixed

- Remove duplicate `application.giantswarm.io/team` label in PodMonitor that caused install failure. The label is already included via the common labels helper.

### cert-exporter [v2.9.15...v2.10.0](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.10.0)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Values: Tune resources.
- Go: Update dependencies.

### cert-manager [v3.9.4...v3.11.0](https://github.com/giantswarm/cert-manager-app/compare/v3.9.4...v3.11.0)

#### Added

- Add Vertical Pod Autoscaler (VPA) support for webhook pods.
- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.
- Add PodLogs for log collection.

#### Fixed

- Fix `controller` Vertical Pod Autoscaler (VPA) resource syntax.

### chart-operator-extensions [v1.1.2...v1.1.3](https://github.com/giantswarm/chart-operator-extensions/compare/v1.1.2...v1.1.3)

#### Changed

- Migrate Chart.yaml annotations to new format as per https://docs.giantswarm.io/reference/platform-api/chart-metadata/

### cilium [v1.3.4...v1.4.1](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.4.1)

#### Changed

- Upgrade Cilium to [v1.19.1](https://github.com/cilium/cilium/releases/tag/v1.19.1).
- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.
- Upgrade Cilium to [v1.18.7](https://github.com/cilium/cilium/releases/tag/v1.18.7).

### cilium-servicemonitors [v0.1.3...v0.1.4](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.3...v0.1.4)

#### Changed

- Migrate chart metadata annotations

### cluster-autoscaler [v1.34.1-1...v1.24.3](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.34.1-1...v1.24.3)

#### Changed

- Change ScaleDownUtilizationThreshold default from 0.5 to 0.7
- Update cluster-autoscaler to version `1.24.3`.

### k8s-dns-node-cache [v2.9.1...v2.9.2](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.1...v2.9.2)

#### Changed

- Upgrade application to version 1.26.7 (includes coredns 1.13.1)

### karpenter [v1.4.0...v2.2.0](https://github.com/giantswarm/karpenter-app/compare/v1.4.0...v2.2.0)

#### Added

- Add `PodLogs` and `PodMonitor` custom resources for observability data ingestion.
- Deployment: Add HTTP proxy support.
- Add e2e tests for this app.
- Add `karpenter-bundle` chart that consolidates `karpenter-app` and `karpenter-crossplane-resources` into a single deployable bundle. The bundle includes:
  - HelmRelease and OCIRepository for deploying karpenter to workload clusters
  - IAM roles for karpenter and nodeclassgenerator via Crossplane
  - SQS queue and CloudWatch event rules for interruption handling

#### Changed

- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.
- Use only `clustertest` v3 instead of v2 and v3. We also upgraded to `apptest-framework` v3 due to this.

### karpenter-taint-remover [v1.0.1...v1.0.2](https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.1...v1.0.2)

#### Changed

- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

### metrics-server [v2.7.0...v2.8.0](https://github.com/giantswarm/metrics-server-app/compare/v2.7.0...v2.8.0)

#### Changed

- Upgrade metrics-server to v0.8.1.
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).

### net-exporter [v1.23.0...v1.23.1](https://github.com/giantswarm/net-exporter/compare/v1.23.0...v1.23.1)

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### node-exporter [v1.20.10...v1.20.11](https://github.com/giantswarm/node-exporter-app/compare/v1.20.10...v1.20.11)

#### Changed

- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

#### Fixed

- Removed duplicated `app` label which is already added by the selector helper.

### observability-bundle [v2.5.0...v2.8.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.8.0)

#### Added

- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.
- Add KSM metrics for Gateway API resources.

#### Changed

- Update kube-prometheus-stack to 20.1.0
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).
- Update alloy-app to 0.17.1
- Update kube-prometheus-stack to 20.0.0
- Update prometheus-operator-crd to 20.0.0

### observability-policies [v0.0.3...v0.0.4](https://github.com/giantswarm/observability-policies-app/compare/v0.0.3...v0.0.4)

#### Changed

- Rename app to `observability-policies`
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).

### priority-classes [v0.3.0...v0.3.1](https://github.com/giantswarm/priority-classes/compare/v0.3.0...v0.3.1)

#### Fixed

- Sanitize `Chart.Version` used in labels. This is needed because flux apapends the digest to the version using the `+` character which is not allowed in labels.

### prometheus-blackbox-exporter [v0.5.0...v0.7.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.5.0...v0.7.0)

#### Added

- Add `http_2xx_insecure` module with `insecure_skip_verify: true` to support probing workload cluster API servers from the management cluster. The MC's service account CA (`http_2xx_k8sca`) only covers the MC itself; workload clusters have their own CA which is not available to the blackbox exporter, making TLS verification impossible without this module.

#### Changed

- Set `priorityClassName` to `system-node-critical` to ensure DaemonSet pods are scheduled even on full nodes.
- Migrate to App Build Suite (ABS) for Helm chart building.

### security-bundle [v1.16.1...v1.17.0](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v1.17.0)

#### Changed

- Update `kyverno` (app) to v0.23.0.
- Update `kyverno-crds` (app) to v1.16.0.
- Update `reports-server` (app) to v0.1.0.
- Update `cloudnative-pg` (app) to v0.0.13.
- Update `kubescape` (app) to v0.0.5.
- Update `starboard-exporter` (app) to v1.0.2.

### vertical-pod-autoscaler [v6.1.1...v6.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.1.1...v6.1.2)

#### Fixed

- Pushed helm chart to OCI repository.

### vertical-pod-autoscaler-crd [v4.1.1...v4.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.1.1...v4.1.2)

#### Fixed

- Pushed helm chart to OCI repository.
