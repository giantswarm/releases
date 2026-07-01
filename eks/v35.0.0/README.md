# :zap: Giant Swarm Release v35.0.0 for EKS :zap:

## Changes compared to v34.0.1

### Components

- cluster-eks from v1.2.1 to v2.2.0
- cluster from v4.0.2 to v6.7.0
- Flatcar from v4459.2.3 to [v4593.2.3](https://www.flatcar.org/releases/#release-4593.2.3)
- Kubernetes from v1.34.4 to [v1.35.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.6)
- os-tooling from v1.26.4 to v1.32.1

### cluster-eks [v1.2.1...v2.2.0](https://github.com/giantswarm/cluster-eks/compare/v1.2.1...v2.2.0)

#### Added

- Add `appVersion` field to `Chart.yaml`.

#### Changed

- Teleport Kube Agent: Conditionally disable PodMonitors and respect new values schema.
- Enable observabilityBundle app.
- Make: Align custom Makefile.
- Chart: Properly truncate & trim version labels.
- Chart: Fix validation errors.
- Chart: Only install `install-cilium-values-update-job` on initial installation.
- Disable kube-proxy and vpc addon by default since cilium replaces it.
- Override cluster-autoscaler `nodeSelector` for EKS — remove control-plane selector since EKS has no control-plane nodes.
- Disable coredns `mastersInstance` and null its control-plane nodeSelector for EKS.
- Enable coredns adopter job for EKS.
- Enable CoreDNS, Cilium and network-policies Apps.
- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

#### Removed

- Renovate: Do no longer trigger E2E Tests on new PRs.
- Remove `cluster-shared` dependency — coredns-adopter job is now handled by `coredns-app` directly.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for `app.kubernetes.io/version` labels.
- Rename `coreDnsExtensions` config template from `EKSCorednsHelmValues` to `EKSCoreDNSExtensionsHelmValues` to avoid confusion with the `coreDns` template name.
- Move `coreDns` Helm values config from `apps/` to the helmrelease config file, since coreDns is deployed as a HelmRelease.
- Use `cluster.providerIntegration.workers.defaultNodePools` of parent chart and define a working default for `replicas` (must be between `minSize` and `maxSize`).
- Only use `ServiceMonitor` and `VerticalPodAutoscaler` in aws-ebs-csi-driver if needed dependencies are enabled

### cluster [v4.0.2...v6.7.0](https://github.com/giantswarm/cluster/compare/v4.0.2...v6.7.0)

#### Added

- Feature Gates: Add support for defining maximum Kubernetes version.
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

- Add support for AKS clusters
- Add support for disabling the external autoscaler annotation (`cluster.x-k8s.io/replicas-managed-by: "external-autoscaler"`) on `MachinePools`.
- Refactor `providerIntegration.resourcesApi.machinePoolResourcesEnabled` into an object `providerIntegration.resourcesApi.machinePoolResources` and move `externalAutoscaler` under it (now `providerIntegration.resourcesApi.machinePoolResources.externalAutoscaler`).
- Bump Flux OCIRepository version to v1.
- Control Plane: Remove handling of clusterRole resources for kamaji Datastore CRs and create kamaji-etcd polex here.
- Control Plane: Make etcd image tag configurable. ([#841](https://github.com/giantswarm/cluster/pull/841))
- Chart: Require `global.release.version` if using Releases to give a better rendering error message.
- Chart: Fix validation errors.
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

- Control Plane: Ensure components start correctly when SELinux is set to enforcing.
- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.
- Cleanup job now also deletes HelmChart CRs to prevent leftover resources when suspending HelmReleases during cluster deletion.

### Apps

- aws-ebs-csi-driver from v4.1.1 to v4.3.0
- cilium from v1.3.4 to v1.4.5
- coredns from v1.30.0 to v1.31.0
- network-policies from v0.1.3 to v0.2.0
- observability-bundle from v2.5.0 to v2.9.1
- teleport-kube-agent from v0.10.8 to v0.11.1

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

### aws-nth-bundle [v1.3.0...v2.1.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.3.0...v2.1.0)

#### Added

- Make the `aws-node-termination-handler` HelmRelease `dependsOn` list configurable via `awsNodeTerminationHandler.dependsOn`.
- Add `cluster.x-k8s.io/cluster-name` label to the HelmReleases.
- Add Flux `dependsOn` from the `aws-node-termination-handler` HelmRelease to the `prometheus-operator-crd` HelmRelease.

#### Changed

- Use `gsoci.azurecr.io/giantswarm/aws-node-termination-handler` for the container image.
- Fix the Flux HelmRelease remediation policy (always remediate).
- Deploy Crossplane resources directly from the bundle instead of standalone helmrelease.
- Clean up values and helm template functions.
- Remove `cluster-values` ConfigMap reference from `aws-node-termination-handler` HelmRelease. Pass `clusterID` explicitly via inline values instead.
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.

### cert-exporter [v2.9.15...v2.11.1](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.11.1)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Values: Tune resources.
- Go: Update dependencies.

#### Fixed

- Add a `serialnumber` label to the `cert_exporter_not_after` and `cert_exporter_secret_not_after` metrics so concatenated certificates no longer collide into identical series. The collision made the registry fail `Gather()`, which blanked out the entire `/metrics` endpoint (regression from v2.10.1).
- Serve `/metrics` with `ContinueOnError` so a single problematic metric can no longer fail the whole scrape.
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

### cilium [v1.3.4...v1.4.5](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.4.5)

#### Changed

- Upgrade Cilium to [v1.19.5](https://github.com/cilium/cilium/releases/tag/v1.19.5).
- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).
- Upgrade Cilium to [v1.19.3](https://github.com/cilium/cilium/releases/tag/v1.19.3).
- Upgrade Cilium to [v1.19.2](https://github.com/cilium/cilium/releases/tag/v1.19.2).
- Upgrade Cilium to [v1.19.1](https://github.com/cilium/cilium/releases/tag/v1.19.1).
- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.
- Upgrade Cilium to [v1.18.7](https://github.com/cilium/cilium/releases/tag/v1.18.7).

### coredns [v1.30.0...v1.31.0](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.31.0)

#### Added

- Wire up the full set of CoreDNS `forward`, `cache`, and `kubernetes` block parameters in the structured zone config:
  - `forward`: `maxIdleConns`, `maxConnectAttempts`, `dohMethod`, `tls`, `tlsServername`, `next`, `nextOnNodata`, `failfastAllUnhealthyUpstreams`, `failover`, `resolver`.
  - `cache`: `zones`, `serveStale.verifyTimeout`, `disable.successZones`, `disable.denialZones`.
  - `kubernetes`: `endpoint`, `tls`, `kubeconfig`, `apiserverQPS`, `apiserverBurst`, `apiserverMaxInflight`, `namespaceLabels`, `fallthroughZones`, `multicluster`, `startupTimeout`.

#### Changed

- Standardize `values.yaml` comments to the `# @schema` / `# --` (helm-docs) convention and remove section-header dividers, so `values.schema.json` and the chart `README.md` are generated from the values file. Rendered manifests are unchanged.
- Update `coredns` image to [1.14.4](https://github.com/coredns/coredns/releases/tag/v1.14.4).
- update CoreDNS icon to light version
- Update `coredns` image to [1.14.3](https://github.com/coredns/coredns/releases/tag/v1.14.3).

#### Fixed

- Render the `health` directive in only the `.` server block. The health plugin is process-wide and can be enabled in just one Server Block, so emitting it in every zone block was invalid. `ready` is kept in every block (its readiness is aggregated across blocks).
- Correct the `coredns.*.cache.serveStale.refreshMode` schema enum to `immediate`/`verify` (was `immediate`/`background`), matching the CoreDNS cache plugin.

### network-policies [v0.1.3...v0.2.0](https://github.com/giantswarm/network-policies-app/compare/v0.1.3...v0.2.0)

#### Added

- Add support for AKS selector labels.

#### Changed

- Deprecated the .Values.kamaji in favour of the more generic .Values.konnectivityAgent to control the behaviour for the `konnectivity-agent`.

### observability-bundle [v2.5.0...v2.9.1](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.9.1)

#### Added

- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle
- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.
- Add KSM metrics for Gateway API resources.

#### Changed

- Update `alloy-app` to 0.20.1
- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0
- Update kube-prometheus-stack to 20.1.0
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).
- Update alloy-app to 0.17.1
- Update kube-prometheus-stack to 20.0.0
- Update prometheus-operator-crd to 20.0.0

### teleport-kube-agent [v0.10.8...v0.11.1](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.1)

#### Changed

- Values: Tolerate `node.cloudprovider.kubernetes.io/uninitialized`.
- Values: Ignore taints regardless of value.
- Values: Pass HTTP proxy settings to sub-chart.
- Updated `teleport-kube-agent` to upstream version `v18.7.6`.
