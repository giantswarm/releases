# :zap: Giant Swarm Release v35.0.0 for EKS :zap:

## Changes compared to v34.0.1

### Components

- cluster-eks from v1.2.1 to v2.0.0
- cluster from v4.0.2 to v6.4.0
- Flatcar from v4459.2.3 to [v4593.2.1](https://www.flatcar.org/releases/#release-4593.2.1)
- Kubernetes from v1.34.4 to [v1.35.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.5)
- os-tooling from v1.26.4 to v1.31.0

### cluster-eks [v1.2.1...v2.0.0](https://github.com/giantswarm/cluster-eks/compare/v1.2.1...v2.0.0)

#### Added

- Add `appVersion` field to `Chart.yaml`.

#### Changed

- Disable kube-proxy and vpc addon by default since cilium replaces it.
- Override cluster-autoscaler `nodeSelector` for EKS — remove control-plane selector since EKS has no control-plane nodes.
- Disable coredns `mastersInstance` and null its control-plane nodeSelector for EKS.
- Enable coredns adopter job for EKS.
- Enable CoreDNS, Cilium and network-policies Apps.
- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

#### Removed

- Remove `cluster-shared` dependency — coredns-adopter job is now handled by `coredns-app` directly.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for `app.kubernetes.io/version` labels.
- Rename `coreDnsExtensions` config template from `EKSCorednsHelmValues` to `EKSCoreDNSExtensionsHelmValues` to avoid confusion with the `coreDns` template name.
- Move `coreDns` Helm values config from `apps/` to the helmrelease config file, since coreDns is deployed as a HelmRelease.
- Use `cluster.providerIntegration.workers.defaultNodePools` of parent chart and define a working default for `replicas` (must be between `minSize` and `maxSize`).
- Only use `ServiceMonitor` and `VerticalPodAutoscaler` in aws-ebs-csi-driver if needed dependencies are enabled

### cluster [v4.0.2...v6.4.0](https://github.com/giantswarm/cluster/compare/v4.0.2...v6.4.0)

#### Added

- Apps: Add External DNS Crossplane Resources.
- Apps: Deploy `cluster-autoscaler` inCluster in Azure.
- MachineDeployment: Add CAPI autoscaler annotations (`cluster-api-autoscaler-node-group-min-size`/`max-size`) when `minSize`/`maxSize` are set on a node pool (only in Azure).
- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))
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

- Configure `observability-bundle` with the management cluster name.
- Apps: Skip `kyverno-crds` dependency for `cluster-autoscaler` when deployed inCluster.
- Apps: Add cluster-probes HelmRelease to deploy ServiceMonitors for probing workload cluster API server endpoint from the management cluster. Configurable via `global.apps.clusterProbes` with default module `http_2xx_insecure` for self-signed certificates.
- Helpers: Use `.Chart.AppVersion` in `app.kubernetes.io/version` label.
- Cluster API: Migrate to API `v1beta2`.
- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.
- Control Plane: Move `node-cidr-mask-size` patch out of `enablePriorityAndFairness` conditional block. ([#741](https://github.com/giantswarm/cluster/pull/741))
- Refactored control plane resource configuration: replaced `providerIntegration.resourcesApi.controlPlaneResourceEnabled` boolean flag= with a unified `controlPlaneResource` object containing `enabled` (boolean) and `provider` (enum: `kubeadm`|`kamaji`) fields. This provides a cleaner, more extensible API for supporting multiple control plane providers.
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

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.
- Remove helm `Job` that cleans up `HelmReleases`. This was needed because we were letting Helm delete the infra cluster and control plane Custom Resources, instead of letting CAPI controllers handle the deletion. This has been fixed, so the `Job` is no longer required.

#### Fixed

- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.
- Cleanup job now also deletes HelmChart CRs to prevent leftover resources when suspending HelmReleases during cluster deletion.

### Apps

- aws-ebs-csi-driver from v4.1.1 to v4.3.0
- aws-ebs-csi-driver-servicemonitors from v0.1.0 to v0.1.2
- aws-nth-bundle from v1.3.0 to v2.0.1
- cert-exporter from v2.9.15 to v2.11.0
- cert-manager from v3.9.4 to v3.13.0
- cert-manager-crossplane-resources from v0.1.0 to v0.1.1
- chart-operator-extensions from v1.1.2 to v1.1.3
- cilium from v1.3.4 to v1.4.4
- cilium-servicemonitors from v0.1.3 to v0.1.4
- cluster-autoscaler from v1.34.1-1 to v1.35.0-1
- coredns from v1.30.0 to v1.30.1
- coredns-extensions from v0.1.2 to v0.1.3
- external-dns from v3.4.0 to v3.5.0
- k8s-dns-node-cache from v2.9.1 to v2.11.0
- karpenter from v1.4.0 to v2.4.0
- karpenter-taint-remover from v1.0.1 to v1.1.0
- metrics-server from v2.7.0 to v2.8.0
- net-exporter from v1.23.0 to v1.24.0
- node-exporter from v1.20.10 to v1.20.11
- observability-bundle from v2.5.0 to v2.9.0
- observability-policies from v0.0.3 to v0.0.4
- priority-classes from v0.3.0 to v0.3.1
- prometheus-blackbox-exporter from v0.5.0 to v0.8.0
- Added rbac-bootstrap v0.3.0
- security-bundle from v1.16.1 to v1.17.1
- teleport-kube-agent from v0.10.8 to v0.11.0
- vertical-pod-autoscaler from v6.1.1 to v6.1.2
- vertical-pod-autoscaler-crd from v4.1.1 to v4.1.2

### aws-ebs-csi-driver [v4.1.1...v4.3.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.1...v4.3.0)

#### :warning: Breaking Changes

- **Workload chart renamed** from `aws-ebs-csi-driver-app` to `aws-ebs-csi-driver`. The OCI catalog artifact name changes accordingly.
- **Bundle values restructured**: upstream chart values are now under the `upstream:` key in the bundle `values.yaml`. The `giantswarm.workloadValues` helper handles the transformation automatically, so bundle users only need to place overrides under `upstream:` in their App CR ConfigMap.
- **Direct workload chart install**: if installing the workload chart directly (without the bundle), all upstream values must be under the `upstream:` key, and extras (`verticalPodAutoscaler`, `networkPolicy`, `global.podSecurityStandards`) are at the top level.

#### Added

- Add IRSA environment variables (`AWS_ROLE_ARN`, `AWS_WEB_IDENTITY_TOKEN_FILE`), projected ServiceAccountToken volume, and `AWS_REGION` to the EBS CSI controller, enabling IRSA authentication in CAPA clusters.
- Propagate proxy values from the bundle (`proxy.http`, `proxy.noProxy`) to the upstream chart (`proxy.http_proxy`, `proxy.no_proxy`) when set.
- Add VPA templates for controller (Deployment) and node (DaemonSet).
- Add Kyverno PolicyException template for Pod Security Standards.
- Add `ignorePaths` to `.kube-linter.yaml` for upstream subchart.
- Forward `verticalPodAutoscaler`, `networkPolicy`, and `global.podSecurityStandards` as extras.

#### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format for both the app and bundle charts.
- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.
- Migrate from forked upstream chart to unmodified upstream as Helm dependency (alias `upstream`).
- Restructure bundle values into explicit BUNDLE-ONLY / UPSTREAM / EXTRAS sections.
- Extract `giantswarm.combineImage` and `giantswarm.setValues` into separate reusable helpers.
- Add `clusterID` derivation from release name as fallback.
- Use `clusterID` helper consistently across all bundle templates.
- Gate NetworkPolicy templates with `networkPolicy.enabled`.
- Rewrite README with architecture diagram, terminology table, value flow, and upgrade notes.

#### Fixed

- Re-enable metrics, force use of `ServiceMonitor` to avoid rendering without them if CRDs are not installed yet
- Fix VPA `updateMode` for `ebs-csi-node` DaemonSet from `Auto` to `Initial`. VPA cannot evict DaemonSet pods, so `Auto` mode silently produces recommendations without ever applying them. `Initial` correctly sets resources at pod creation time.
- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.

### aws-ebs-csi-driver-servicemonitors [v0.1.0...v0.1.2](https://github.com/giantswarm/aws-ebs-csi-driver-servicemonitors-app/compare/v0.1.0...v0.1.2)

#### Changed

- Migrate to App Build Suite (ABS).

#### Fixed

- Remove duplicate `application.giantswarm.io/team` label in PodMonitor that caused install failure. The label is already included via the common labels helper.

### aws-nth-bundle [v1.3.0...v2.0.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.3.0...v2.0.1)

#### Added

- Make the `aws-node-termination-handler` HelmRelease `dependsOn` list configurable via `awsNodeTerminationHandler.dependsOn`.
- Add `cluster.x-k8s.io/cluster-name` label to the HelmReleases.
- Add Flux `dependsOn` from the `aws-node-termination-handler` HelmRelease to the `prometheus-operator-crd` HelmRelease.

#### Changed

- Deploy Crossplane resources directly from the bundle instead of standalone helmrelease.
- Clean up values and helm template functions.
- Remove `cluster-values` ConfigMap reference from `aws-node-termination-handler` HelmRelease. Pass `clusterID` explicitly via inline values instead.
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.

### cert-exporter [v2.9.15...v2.11.0](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.11.0)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Values: Tune resources.
- Go: Update dependencies.

#### Fixed

- Parse all PEM blocks in secrets and certificate files, not just the first one. This fixes false alerts when multiple certificates are concatenated (e.g. Kyverno webhook cert rotation).

### cert-manager [v3.9.4...v3.13.0](https://github.com/giantswarm/cert-manager-app/compare/v3.9.4...v3.13.0)

#### Added

- Add control plane node toleration to CA injector deployment.
- Add Vertical Pod Autoscaler (VPA) support for webhook pods.
- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.
- Add PodLogs for log collection.

#### Changed

- Upgrade cert-manager to v1.19.4.

#### Removed

- Remove PodSecurityPolicy (PSP) and related resources.
- Remove Giant Swarm PSP to PSS migration logic.

#### Fixed

- Fix `controller` Vertical Pod Autoscaler (VPA) resource syntax.

### cert-manager-crossplane-resources [v0.1.0...v0.1.1](https://github.com/giantswarm/cert-manager-crossplane-resources/compare/v0.1.0...v0.1.1)

#### Changed

- Update `architect-orb` to v6.15.0.

### chart-operator-extensions [v1.1.2...v1.1.3](https://github.com/giantswarm/chart-operator-extensions/compare/v1.1.2...v1.1.3)

#### Changed

- Migrate Chart.yaml annotations to new format as per https://docs.giantswarm.io/reference/platform-api/chart-metadata/

### cilium [v1.3.4...v1.4.4](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.4.4)

#### Changed

- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).
- Upgrade Cilium to [v1.19.3](https://github.com/cilium/cilium/releases/tag/v1.19.3).
- Upgrade Cilium to [v1.19.2](https://github.com/cilium/cilium/releases/tag/v1.19.2).
- Upgrade Cilium to [v1.19.1](https://github.com/cilium/cilium/releases/tag/v1.19.1).
- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.
- Upgrade Cilium to [v1.18.7](https://github.com/cilium/cilium/releases/tag/v1.18.7).

### cilium-servicemonitors [v0.1.3...v0.1.4](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.3...v0.1.4)

#### Changed

- Migrate chart metadata annotations

### cluster-autoscaler [v1.34.1-1...v1.35.0-1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.34.1-1...v1.35.0-1)

#### Added

- Validate that `managementCluster` (when `isManagementCluster=true`) or `clusterID` (otherwise) are set, failing early with a clear error message.
- Add support for CAPI mode (`kubeconfig-incluster`): run cluster-autoscaler on the management cluster using a pre-existing kubeconfig to connect to the workload cluster.
- Add `clusterAPI` values section for configuring CAPI mode (autodiscovery, kubeconfig secret, configmaps namespace).
- Add `rbac.clusterScoped` toggle to support namespace-scoped RBAC (no ClusterRole/ClusterRoleBinding) for CAPI deployments.

#### Changed

- Chart: Update to upstream v1.35.0.
- Migrate test infrastructure from pipenv to uv.
- Deploy the Kyverno policy exception in the `policy-exceptions` Namespace.
- Deploy the Kyverno PolicyException as a Helm `pre-install,pre-upgrade` hook so it takes effect before chart resources are created.
- Chart: Update to upstream v1.34.3.
- Chart: Update to upstream v1.34.2.

### coredns [v1.30.0...v1.30.1](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.30.1)

#### Changed

- Update `coredns` image to [1.14.3](https://github.com/coredns/coredns/releases/tag/v1.14.3).

### external-dns [v3.4.0...v3.5.0](https://github.com/giantswarm/external-dns-app/compare/v3.4.0...v3.5.0)

#### Changed

- Update VPA `updatePolicy.updateMode` from deprecated `Auto` to `Recreate`.
- Upgrade external-dns to [v0.21.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.21.0).
- Sync to upstream helm chart [1.21.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.21.1).
  - Add `namespaceOverride` value to support deploying external-dns into a different namespace than the Helm release (useful for subchart usage).
  - Add `enableGatewayListenerSets` value to opt into Gateway API ListenerSet resource support.
  - Add `sourceNamespace` value (used with `namespaced=true`) to watch resources in a namespace different from the deployment namespace.
  - Avoid creating cluster-scoped namespace RBAC when `gatewayNamespace` is set, reducing required permissions.
  - Fix `extraArgs` map handling: boolean values now render as `--flag` / `--no-flag` and string values are properly quoted.
- Use external-dns.namespace in VPA and NetworkPolicy resources.

### k8s-dns-node-cache [v2.9.1...v2.11.0](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.1...v2.11.0)

#### Added

- Add `configmap.log.enabled` helm value to toggle CoreDNS query logging (default: `false`).
- Make `AAAA NOERROR` configurable for IPv6.

#### Changed

- Upgrade application to version 1.26.7 (includes coredns 1.13.1)

### karpenter [v1.4.0...v2.4.0](https://github.com/giantswarm/karpenter-app/compare/v1.4.0...v2.4.0)

#### Added

- Add `cluster.x-k8s.io/cluster-name` label to the karpenter HelmRelease.
- Add `iam:GetInstanceProfile` permission to Karpenter IAM role.
- Add karpenter CRDs.
- Set `helm.sh/resource-policy: keep` on the karpenter CRDs so they survive HelmRelease uninstall and prevent cascade-deleting `NodePool`/`NodeClaim`/`EC2NodeClass` resources.
- Add `PodLogs` and `PodMonitor` custom resources for observability data ingestion.
- Deployment: Add HTTP proxy support.
- Add e2e tests for this app.
- Add `karpenter-bundle` chart that consolidates `karpenter-app` and `karpenter-crossplane-resources` into a single deployable bundle. The bundle includes:
  - HelmRelease and OCIRepository for deploying karpenter to workload clusters
  - IAM roles for karpenter and nodeclassgenerator via Crossplane
  - SQS queue and CloudWatch event rules for interruption handling

#### Changed

- Switch e2e scale test from App CR to Flux HelmRelease for deploying hello-world, avoiding `values-schema-violation` errors caused by app-platform injected properties.
- Improve Crossplane ConfigMap fetching logic
- Migrate workload chart to use unmodified upstream Karpenter v1.8.1 chart as a Helm dependency (`alias: upstream`), eliminating fork maintenance.
- Bundle chart: add `giantswarm.workloadValues` transformer to route values under `upstream:` key with extras (`podLogs`, `global`) at top level.
- Bundle chart: convert proxy settings to `controller.env` entries for upstream compatibility.
- Bundle chart: add `giantswarm.combineImage` helper to merge split `registry`+`repository` into single `repository` path.
- Restructure bundle `values.yaml` into annotated BUNDLE-ONLY / UPSTREAM / EXTRAS sections.
- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format for both the karpenter and karpenter-bundle charts.
- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.

#### Removed

- Remove all forked upstream templates from workload chart (replaced by upstream dependency).
- Remove `vendir.yml`, `vendir.lock.yml`, `vendor/` directory, and `Makefile.custom.mk`.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.
- Use only `clustertest` v3 instead of v2 and v3. We also upgraded to `apptest-framework` v3 due to this.

### karpenter-taint-remover [v1.0.1...v1.1.0](https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.1...v1.1.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image.
- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.
- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

### metrics-server [v2.7.0...v2.8.0](https://github.com/giantswarm/metrics-server-app/compare/v2.7.0...v2.8.0)

#### Changed

- Upgrade metrics-server to v0.8.1.
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).

### net-exporter [v1.23.0...v1.24.0](https://github.com/giantswarm/net-exporter/compare/v1.23.0...v1.24.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the net-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Bump `docker-kubectl` init container from `1.25.4` to `1.36.0`.

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### node-exporter [v1.20.10...v1.20.11](https://github.com/giantswarm/node-exporter-app/compare/v1.20.10...v1.20.11)

#### Changed

- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

#### Fixed

- Removed duplicated `app` label which is already added by the selector helper.

### observability-bundle [v2.5.0...v2.9.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.9.0)

#### Added

- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle
- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.
- Add KSM metrics for Gateway API resources.

#### Changed

- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0
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

### prometheus-blackbox-exporter [v0.5.0...v0.8.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.5.0...v0.8.0)

#### Added

- Add toleration for `kubernetes.io/arch=arm64:NoSchedule` so the DaemonSet schedules on ARM worker nodes.
- Add `http_2xx_insecure` module with `insecure_skip_verify: true` to support probing workload cluster API servers from the management cluster. The MC's service account CA (`http_2xx_k8sca`) only covers the MC itself; workload clusters have their own CA which is not available to the blackbox exporter, making TLS verification impossible without this module.

#### Changed

- Set `priorityClassName` to `system-node-critical` to ensure DaemonSet pods are scheduled even on full nodes.
- Migrate to App Build Suite (ABS) for Helm chart building.

### rbac-bootstrap [v0.3.0](https://github.com/giantswarm/rbac-bootstrap-app/releases/tag/v0.3.0)

#### Added

- Add `io.giantswarm.application.managed` chart annotation for Backstage visibility.
- Add optional `cluster-reader` ClusterRole (off by default, enabled via `clusterReader.enabled: true`) that aggregates into the built-in `view` ClusterRole and grants read access (`get`/`list`/`watch`) on cluster-scoped resources.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.

### security-bundle [v1.16.1...v1.17.1](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v1.17.1)

#### Added

- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.

#### Changed

- Update `falco` (app) to v0.11.2.
- Update `gel` (app) to v1.0.2.
- Update `kubescape` (app) to v0.0.6.
- Update `reports-server` (app) to v0.1.3.
- Update `starboard-exporter` (app) to v1.0.3.
- Update `trivy` (app) to v0.14.2.
- Update `trivy-operator` (app) to v0.12.2.
- Migrate chart annotations to OCI-compatible format.
- Update `kyverno` (app) to v0.23.0.
- Update `kyverno-crds` (app) to v1.16.0.
- Update `reports-server` (app) to v0.1.0.
- Update `cloudnative-pg` (app) to v0.0.13.
- Update `kubescape` (app) to v0.0.5.
- Update `starboard-exporter` (app) to v1.0.2.

### teleport-kube-agent [v0.10.8...v0.11.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.7.6`.

### vertical-pod-autoscaler [v6.1.1...v6.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.1.1...v6.1.2)

#### Fixed

- Pushed helm chart to OCI repository.

### vertical-pod-autoscaler-crd [v4.1.1...v4.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.1.1...v4.1.2)

#### Fixed

- Pushed helm chart to OCI repository.
