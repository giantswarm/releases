# :zap: Giant Swarm Release v35.0.0 for EKS :zap:

## Changes compared to v34.0.1

### Components

- cluster-eks from v1.2.1 to v3.0.0
- cluster from v4.0.2 to v7.0.0
- Added containerd [v2.3.2](https://github.com/containerd/containerd/releases/tag/v2.3.2)
- Flatcar from v4459.2.3 to [v4593.2.5](https://www.flatcar.org/releases/#release-4593.2.5)
- Kubernetes from v1.34.4 to [v1.35.7](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.7)
- os-tooling from v1.26.4 to v1.34.0

### cluster-eks [v1.2.1...v3.0.0](https://github.com/giantswarm/cluster-eks/compare/v1.2.1...v3.0.0)

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

### cluster [v4.0.2...v7.0.0](https://github.com/giantswarm/cluster/compare/v4.0.2...v7.0.0)

#### Added

- Add `cert-manager-crossplane-resources` HelmRelease.
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

- Migrate `coredns` HelmRelease values to the new `coredns-app` zone-aware interface.
- Rename the internal `coredns` control plane helper to align with the `controlPlane` values key.
- Updated `cert-manager` to v4.0.0 and migrated the values to match the new chart's schema.
- Support templating on the `global.apps.<name>.extraConfigs.name` field.
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
- cilium from v1.3.4 to v1.5.1
- coredns from v1.30.0 to v1.32.0
- network-policies from v0.1.3 to v0.2.0
- observability-bundle from v2.5.0 to v3.3.1
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

### cilium [v1.3.4...v1.5.1](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.5.1)

#### Added

- `sync/verify-images.sh`, run as the last step of `sync/sync.sh` and therefore in CI: renders the chart across four scenarios covering every image-bearing component and fails the sync unless each image is served from `gsoci.azurecr.io/giantswarm/` or is explicitly listed in `sync/unmirrored-images.txt`. It also asserts that every configured image actually appears in a render, so a scenario that stops covering an image fails instead of silently narrowing the check. This replaces the removed patch's `fail` guards, which were the only thing that made image drift visible on an upstream bump.

#### Changed

- Upgrade Cilium to [v1.19.7](https://github.com/cilium/cilium/releases/tag/v1.19.7).
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Upgrade Cilium to [v1.19.6](https://github.com/cilium/cilium/releases/tag/v1.19.6).
- Switch Hubble TLS certificate provisioning from `hubble.tls.auto.method: helm` to `cronJob` ([giantswarm#37201](https://github.com/giantswarm/giantswarm/issues/37201)). The helm method minted certificates once and never renewed them, deterministically breaking hubble-relay when the leaf certs (1 year) or the CA (3 years) expired. With the cronJob method a `hubble-generate-certs` CronJob re-issues the leaf certificates every 4 months. **On upgrade, the Helm-owned `cilium-ca`, `hubble-server-certs` and `hubble-relay-client-certs` secrets are deleted and re-created by a one-shot certgen job with a fresh 3-year CA; agents and hubble-relay hot-reload the new certificates without restarts.**
- Wire certgen's `--ca-enforce-validity-throughout-leaves-duration` flag (new value `certgen.enforceCAValidityThroughoutLeavesDuration`, default `true`): the certgen job now fails roughly one year before the CA would no longer cover new leaf certificates, instead of silently issuing leafs that outlive the CA. certgen never rotates an existing CA on its own ([cilium/certgen#500](https://github.com/cilium/certgen/issues/500)).
- Relax `hubble-relay` `podAffinity` to `preferredDuringSchedulingIgnoredDuringExecution` so Karpenter can drain the last cilium-agent-bearing node during upgrades/consolidation instead of getting stuck on a required co-location with cilium.
- Upgrade Cilium to [v1.19.5](https://github.com/cilium/cilium/releases/tag/v1.19.5).
- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).
- Upgrade Cilium to [v1.19.3](https://github.com/cilium/cilium/releases/tag/v1.19.3).
- Upgrade Cilium to [v1.19.2](https://github.com/cilium/cilium/releases/tag/v1.19.2).
- Upgrade Cilium to [v1.19.1](https://github.com/cilium/cilium/releases/tag/v1.19.1).
- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.
- Upgrade Cilium to [v1.18.7](https://github.com/cilium/cilium/releases/tag/v1.18.7).

#### Removed

- **Removed the `image.registry` value and the sync patch behind it** ([roadmap#3264](https://github.com/giantswarm/roadmap/issues/3264)). The registry is now part of each image's `repository` value, exactly as upstream ships it, so the chart no longer patches upstream's `cilium.image`/`cilium.operator.image` helpers into a `(list $ <image>)` signature and no longer rewrites all 35 call sites with `sed`. `helm/cilium/templates/_helpers.tpl` and `helm/cilium/templates/cilium-operator/_helpers.tpl` are now byte-identical to upstream, and 15 template patches disappear from `diffs/`.

#### Fixed

- Fix four image references that were rendered as unpullable double-prefixed paths, because their `repository` already carried a registry while the removed patch prefixed `image.registry` on top: `gsoci.azurecr.io/ghcr.io/spiffe/spire-server`, `.../spire-agent`, `.../docker.io/library/busybox` (SPIRE mutual authentication) and `.../docker.io/istio/ztunnel` (`encryption.type=ztunnel`). All four now resolve. Latent until now because both features are disabled by default.
- Fix rendering of the certgen job specs (`hubble/tls-cronjob` and `clustermesh-apiserver/tls-cronjob`): the image reference was not converted to the Giant Swarm `cilium.image` helper signature because the image-registries sync patch only processed `*.yaml` templates, so enabling the cronJob method failed with `required list, but got "map"`.
- Add NetworkPolicies allowing the `hubble-generate-certs` and `clustermesh-apiserver-generate-certs` certgen pods egress to the Kubernetes API.

### coredns [v1.30.0...v1.32.0](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.32.0)

#### Added

- Wire up the full set of CoreDNS `forward`, `cache`, and `kubernetes` block parameters in the structured zone config:
  - `forward`: `maxIdleConns`, `maxConnectAttempts`, `dohMethod`, `tls`, `tlsServername`, `next`, `nextOnNodata`, `failfastAllUnhealthyUpstreams`, `failover`, `resolver`.
  - `cache`: `zones`, `serveStale.verifyTimeout`, `disable.successZones`, `disable.denialZones`.
  - `kubernetes`: `endpoint`, `tls`, `kubeconfig`, `apiserverQPS`, `apiserverBurst`, `apiserverMaxInflight`, `namespaceLabels`, `fallthroughZones`, `multicluster`, `startupTimeout`.

#### Changed

- Chart: Make tolerations configurable.
- Rebuild with `app-build-suite` 2.2.0 (via `architect` orb 9.6.0): the packaged chart now carries Artifact Hub metadata (`artifacthub.io/license` and a Support link). No functional chart changes.
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

### observability-bundle [v2.5.0...v3.3.1](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v3.3.1)

#### Added

- Add KSM metrics for Gateway API `ListenerSet` and `ReferenceGrant` resources.
- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle
- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.
- Add KSM metrics for Gateway API resources.

#### Changed

- Values: Generate schema for Alloy PodLogs CRDs.
- Values: Add Cilium as dependency for Alloy apps & Kube Prometheus Stack.
- Values: Update Alloy apps to v0.21.2.
- Update Gateway API KSM configs to `v1` for `Gateway`, `GatewayClass`, `HTTPRoute`, `GRPCRoute`, `TLSRoute` and `BackendTLSPolicy`.
- Update `kube-prometheus-stack` and `prometheus-operator-crd` to 22.0.0
- Update `alloy-app` to 0.21.0
- HelmReleases: honor the App platform `priority` field (1-150, default 25) on `extraConfigs` entries. `spec.valuesFrom` now reproduces the App platform merge order — all configMaps before all secrets (a secret always overrides a configMap), each kind ordered by priority around the user-config layer — preserving the App CR merge semantics after the migration. ([giantswarm#36096](https://github.com/giantswarm/giantswarm/issues/36096))
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- Remove 'cluster-values' ConfigMap reference from HelmReleases.
- Add new `alloy-podlogs-crds` chart.
- Update alloy-app to 0.20.0
- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0
- Update kube-prometheus-stack to 20.1.0
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).
- Update alloy-app to 0.17.1
- Update kube-prometheus-stack to 20.0.0
- Update prometheus-operator-crd to 20.0.0

#### Removed

- Values: Remove unused catalog.

### teleport-kube-agent [v0.10.8...v0.11.1](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.1)

#### Changed

- Values: Tolerate `node.cloudprovider.kubernetes.io/uninitialized`.
- Values: Ignore taints regardless of value.
- Values: Pass HTTP proxy settings to sub-chart.
- Updated `teleport-kube-agent` to upstream version `v18.7.6`.
