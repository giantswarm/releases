# :zap: Giant Swarm Release v36.1.0 for EKS :zap:

## Changes compared to v34.0.1

### Components

- cluster-eks from v1.2.1 to v4.0.0
- cluster from v4.0.2 to v9.4.0
- Added containerd [v2.3.2](https://github.com/containerd/containerd/releases/tag/v2.3.2)
- Flatcar from v4459.2.3 to [v4757.2.2](https://www.flatcar.org/releases/#release-4757.2.2)
- Kubernetes from v1.34.4 to [v1.36.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.6)
- os-tooling from v1.26.4 to v1.36.0

### cluster-eks [v1.2.1...v4.0.0](https://github.com/giantswarm/cluster-eks/compare/v1.2.1...v4.0.0)

#### Added

- Add the `giantswarm.io/machine-pool` label to the nodes of each EKS managed node group.
- Apply `global.nodePools.<name>.customNodeLabels` to the nodes of EKS managed node groups.
- Enable external-dns and external-dns-crossplane-resources, which creates the external-dns IAM role for the EKS OIDC provider. The OIDC domain is read from the `<cluster>-crossplane-config` ConfigMap.
- Enable cilium-servicemonitors and net-exporter.
- Enable the remaining default apps from `cluster-aws` that work on EKS: cert-exporter, cert-manager, chart-operator-extensions, cluster-autoscaler, coredns-extensions, k8s-dns-node-cache, metrics-server, node-exporter, observability-policies, priority-classes, prometheus-blackbox-exporter, security-bundle, vertical-pod-autoscaler, and vertical-pod-autoscaler-crd.
- Enable cluster-autoscaler-crossplane-resources, which creates the cluster-autoscaler IAM role for the EKS OIDC provider.
- Enable cert-manager-crossplane-resources, which creates the cert-manager IAM role for the EKS OIDC provider, and let cert-manager solve ACME challenges with Route53 DNS01 challenges.
- Add `global.nodePools.<pool>.replicas` for Karpenter node pools to keep a fixed number of nodes (static node pool). This is a tech preview of the alpha Karpenter `StaticCapacity` feature and requires a release with Karpenter v1.14.1 and the matching `aws-resolver-rules-operator`.
- Add `global.nodePools.<pool>.limits.nodes` to cap the number of nodes of a Karpenter node pool.
- Add the `Balanced` consolidation policy for Karpenter node pools.
- CI: Run the full set of E2E test suites automatically on release PRs, via `.github/release-pr-body.md`. Towards https://github.com/giantswarm/roadmap/issues/4334
- Add karpenter support. For a node pool, set `global.nodePools.<name>.type` to `karpenter` to have karpenter manage a pool instead of an EKS managed node group. For now, karpenter can only run on a managed node pool, not on karpenter-provisioned node pools, to ensure it keeps running and doesn't get stopped while doing any provisioning.
- Deploy the `karpenter-taint-remover` app
- Add `appVersion` field to `Chart.yaml`.

#### Changed

- Karpenter node pools: `customNodeLabels` can no longer override the `giantswarm.io/machine-pool` node label.
- Fail rendering if a `customNodeLabels` entry is not in `key=value` form.
- Bump `OCIRepository` resources from the `source.toolkit.fluxcd.io/v1beta2` API to `source.toolkit.fluxcd.io/v1`.
- Chart: Migrate Apps to HelmReleases. This skips the App to HelmRelease migration of `cluster` v8.x, since there are no live EKS clusters that would need it.
- Cilium: Tolerate the `karpenter.sh/unregistered` taint for `cilium-operator`, `hubble-relay`, `hubble-ui` and `certgen`, since `cluster` no longer tolerates all taints for them.
- CoreDNS: Drop the deprecated network values, the `cluster` chart defaults already set them via the new interface.
- Read the EKS node group update configuration from `global.nodePools.<name>.updateConfig`, where the schema has always documented it
- aws-ebs-csi-driver: Don't wait for a `cloud-provider-aws` HelmRelease, since the AWS cloud provider is part of the EKS control plane
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

- CoreDNS: Don't deploy the `coredns-controlplane` instance, since EKS has no control-plane nodes. `cluster` v7 sets `controlPlane.enabled`, which takes precedence over the deprecated `mastersInstance.enabled` that disabled it before. This requires coredns-app v1.31.0 or newer.
- Normalize lowercase `extraConfigs[*].kind` values of `global.apps.awsEbsCsiDriver` and `global.apps.karpenter` to the `ConfigMap` and `Secret` kinds that Flux accepts.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
- Make use of Kyverno and VPA CRDs conditional for Alloy components. This requires a newer version of observability-bundle, which takes `apps.<appName>.dependsOn` as a map.
- Use `.Chart.AppVersion` instead of `.Chart.Version` for `app.kubernetes.io/version` labels.
- Rename `coreDnsExtensions` config template from `EKSCorednsHelmValues` to `EKSCoreDNSExtensionsHelmValues` to avoid confusion with the `coreDns` template name.
- Move `coreDns` Helm values config from `apps/` to the helmrelease config file, since coreDns is deployed as a HelmRelease.
- Use `cluster.providerIntegration.workers.defaultNodePools` of parent chart and define a working default for `replicas` (must be between `minSize` and `maxSize`).
- Only use `ServiceMonitor` and `VerticalPodAutoscaler` in aws-ebs-csi-driver if needed dependencies are enabled

### cluster [v4.0.2...v9.4.0](https://github.com/giantswarm/cluster/compare/v4.0.2...v9.4.0)

#### Added

- Apps: Add `providerIntegration.apps.<app>.extraConfigs` so provider charts can add templated config maps or secrets to an app's values.
- Ignition: Expose `sysctl` and `sysfs` configuration per node pool via `kernel.sysctls` (`/etc/sysctl.d/90-node-pool.conf`) and `kernel.sysfs` (`/etc/tmpfiles.d/00-node-pool.conf`). The parameters required by kubelet's `protectKernelDefaults` cannot be set by name, and the Image Builder and Cilium sysctl configuration keeps precedence.
- Apps: Add `global.connectivity.cni.managed` (default `true`). When `false`, Cilium isn't rendered and is dropped from the dependencies of every app.
- Workers: Add `nodeLabelsTemplateName` and `taintsTemplateName` hooks under `providerIntegration.workers.kubeadmConfig`. They name provider templates, rendered once per node pool, that add kubelet node labels and taints.
- Add `filesTemplateName` hook under `providerIntegration.workers.kubeadmConfig`. It names a provider template that renders a YAML list of files, once per node pool, so that a provider can deploy files to selected node pools only. A node pool for which the template renders nothing keeps its `KubeadmConfig` spec, and therefore its spec hash, unchanged.
- Enable `mergeDefaultEvictionSettings` to keep defaults for eviction like `nodefs.available` and `nodefs.inodesFree` which would otherwise be set to 0. This rolls all nodes.
- Add the `app.kubernetes.io/component` label with the app name to the resources rendered per app, so they can be listed together.
- Add `preKubeadmCommandsTemplateName` and `postKubeadmCommandsTemplateName` hooks under `providerIntegration.controlPlane.kubeadmConfig` and `providerIntegration.workers.kubeadmConfig`. They name a provider template that renders a YAML list of additional kubeadm commands, once for the control plane and once per node pool for workers.
- Add `internal.advancedConfiguration.kubelet.evictionHard` values. Providers need them to tell autoscalers such as Karpenter how much of a node's resources is allocatable.
- SELinux: Add `global.components.selinux.writablePolicyStore` value (default `true`) to allow loading additional SELinux policies.
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

- Ignition: Rename `/etc/sysctl.d/hardening.conf` to `80-hardening.conf` so it sorts after Flatcar's `65-baselayout.conf` but before Image Builder's `99-sysctl.conf` and Cilium's `99-zzz-override_cilium.conf`. Previously the unprefixed file sorted last and overrode both. As a result, nodes now run with `fs.inotify.max_user_watches = 524288` (previously `16384`) and `net.ipv4.conf.all.rp_filter = 0` (previously `1`).
- Ignition: Clean up sysctl hardening. Drop keys already set by Image Builder or Cilium (`fs.inotify.max_user_watches`, `fs.inotify.max_user_instances`, `net.ipv4.conf.all.rp_filter`, `vm.overcommit_memory`, `kernel.panic`, `kernel.panic_on_oops`). Stop disabling `net.ipv4.tcp_timestamps`, since modern kernels randomize timestamp offsets.
- `kubeadm`: Exclude `/etc/.systemd-confext` from `restorecon`.
- Cilium: Replace the catch-all `- operator: Exists` toleration on `cilium-operator` with an explicit list.
- Stop deleting and recreating `/etc/ssl/certs` on nodes for fixing SELinux labeling, so any custom certificates placed there directly are preserved.
- Cilium: Replace the catch-all `- operator: Exists` toleration on the `hubble-relay`, `hubble-ui` and `certgen` components with an explicit list.
- Enable the `ClusterTrustBundle` and `ClusterTrustBundleProjection` feature gates (Kubernetes 1.33+) and the `PodCertificateRequest` feature gate (Kubernetes 1.35+) by default on kube-apiserver, kube-controller-manager and kubelet.
- SELinux: Keep AVC audit logs (required for SELinux policy generation).
- SELinux: Relabel the whole filesystem except read-only `/usr` (previously only `/etc/kubernetes`).
- SELinux: Correctly label CA certificates in `/etc/ssl/certs` for mounting into containers.
- App to HR Migration: Skip v35.0.0 pre-releases and update `docker-kubectl` to v1.36.4.
- Chart: Rework HelmRelease clean-up job.
- Chart: Migrate Apps to HelmReleases.
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

- Ignition: Remove `os-hardening.service`. Its `sysctl -p` duplicated what `systemd-sysctl` already applies at boot. Its `gpasswd` calls had no effect: glibc merges group memberships from `/usr/share/baselayout/group` (via `nss-usrfiles`) regardless of `/etc/group`.
- Apps: Remove the static `aks` exceptions for the Cilium dependency of network policies and CoreDNS.
- Workers: Remove the `karpenter.sh/unregistered:NoExecute` taint from Karpenter node pools. Provider charts must add it via `taintsTemplateName`.
- Chart: Remove App to HelmRelease migration.
- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.
- Remove helm `Job` that cleans up `HelmReleases`. This was needed because we were letting Helm delete the infra cluster and control plane Custom Resources, instead of letting CAPI controllers handle the deletion. This has been fixed, so the `Job` is no longer required.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
- Control Plane: Ensure components start correctly when SELinux is set to enforcing.
- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.
- Cleanup job now also deletes HelmChart CRs to prevent leftover resources when suspending HelmReleases during cluster deletion.

### Apps

- aws-ebs-csi-driver from v4.1.1 to v4.3.2
- cert-exporter from v2.9.15 to v2.12.1
- cert-manager from v3.9.4 to v4.1.2
- cert-manager-crossplane-resources from v0.1.0 to v0.3.0
- chart-operator-extensions from v1.1.2 to v1.1.3
- cilium from v1.3.4 to v1.6.2
- cilium-servicemonitors from v0.1.3 to v0.2.0
- cluster-autoscaler from v1.34.1-1 to v2.1.1
- Added cluster-autoscaler-crossplane-resources v1.0.0
- coredns from v1.30.0 to v1.35.0
- coredns-extensions from v0.1.2 to v0.1.3
- external-dns from v3.4.0 to v3.7.1
- Added external-dns-crossplane-resources v0.4.0
- k8s-dns-node-cache from v2.9.1 to v2.11.1
- karpenter from v1.4.0 to v2.6.0
- karpenter-taint-remover from v1.0.1 to v1.1.0
- metrics-server from v2.7.0 to v2.8.0
- net-exporter from v1.23.0 to v1.24.1
- network-policies from v0.1.3 to v0.5.0
- node-exporter from v1.20.10 to v1.21.0
- observability-bundle from v2.5.0 to v3.6.0
- observability-policies from v0.0.3 to v0.0.4
- priority-classes from v0.3.0 to v0.3.1
- prometheus-blackbox-exporter from v0.5.0 to v0.10.0
- security-bundle from v1.16.1 to v2.4.0
- teleport-kube-agent from v0.10.8 to v0.12.0
- vertical-pod-autoscaler from v6.1.1 to v6.1.2
- vertical-pod-autoscaler-crd from v4.1.1 to v4.1.2

### aws-ebs-csi-driver [v4.1.1...v4.3.2](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.1...v4.3.2)

#### :warning: Breaking Changes

- **Workload chart renamed** from `aws-ebs-csi-driver-app` to `aws-ebs-csi-driver`. The OCI catalog artifact name changes accordingly.
- **Bundle values restructured**: upstream chart values are now under the `upstream:` key in the bundle `values.yaml`. The `giantswarm.workloadValues` helper handles the transformation automatically, so bundle users only need to place overrides under `upstream:` in their App CR ConfigMap.
- **Direct workload chart install**: if installing the workload chart directly (without the bundle), all upstream values must be under the `upstream:` key, and extras (`verticalPodAutoscaler`, `networkPolicy`, `global.podSecurityStandards`) are at the top level.

#### Added

- HelmRelease: Add dependencies.
- Add IRSA environment variables (`AWS_ROLE_ARN`, `AWS_WEB_IDENTITY_TOKEN_FILE`), projected ServiceAccountToken volume, and `AWS_REGION` to the EBS CSI controller, enabling IRSA authentication in CAPA clusters.
- Propagate proxy values from the bundle (`proxy.http`, `proxy.noProxy`) to the upstream chart (`proxy.http_proxy`, `proxy.no_proxy`) when set.
- Add VPA templates for controller (Deployment) and node (DaemonSet).
- Add Kyverno PolicyException template for Pod Security Standards.
- Add `ignorePaths` to `.kube-linter.yaml` for upstream subchart.
- Forward `verticalPodAutoscaler`, `networkPolicy`, and `global.podSecurityStandards` as extras.

#### Changed

- Conditionally depend on Kyverno and VPA CRDs since security-bundle isn't yet implemented for cluster-eks and customers may disable that bundle for cluster adoption cases
- Run the e2e test suites automatically on release PRs by adding `.github/release-pr-body.md`.
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

### cert-exporter [v2.9.15...v2.12.1](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.12.1)

#### Added

- Regression tests covering the v2.11.1 metrics endpoint fix: concatenated and repeated certificates in secrets and files, and an endpoint level test asserting `/metrics` keeps returning 200 and serving the remaining metrics when a duplicate series is emitted.
- ATS: End to end test for a TLS secret whose `tls.crt` holds a concatenated certificate chain.
- DaemonSet: Add VPA.

#### Changed

- ATS: `cert_gen` now sets a certificate serial number, which defaulted to `0` for every generated certificate.
- Go: Update dependencies.
- Templates: Move PolicyException to `kube-system` namespace.
- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Values: Tune resources.

#### Fixed

- A cert file that cannot be read no longer aborts the scan of its whole cert path. Previously one unreadable file (such as a root-only `0600` `ca.crt` on an AKS node, where the exporter runs as an unprivileged user) stopped the walk, silently dropping every file sorting after it from the metrics. Unreadable files are now logged and skipped individually.
- Add a `serialnumber` label to the `cert_exporter_not_after` and `cert_exporter_secret_not_after` metrics so concatenated certificates no longer collide into identical series. The collision made the registry fail `Gather()`, which blanked out the entire `/metrics` endpoint (regression from v2.10.1).
- Serve `/metrics` with `ContinueOnError` so a single problematic metric can no longer fail the whole scrape.
- Parse all PEM blocks in secrets and certificate files, not just the first one. This fixes false alerts when multiple certificates are concatenated (e.g. Kyverno webhook cert rotation).

### cert-manager [v3.9.4...v4.1.2](https://github.com/giantswarm/cert-manager-app/compare/v3.9.4...v4.1.2)

#### Added

- cert-manager-giantswarm-clusterissuer: Allow setting `cnameStrategy` per DNS01 solver with `acme.dns01.cloudflare.cnameStrategy`, `acme.dns01.route53.cnameStrategy` and `acme.dns01.azureDNS.cnameStrategy`, so a `_acme-challenge` CNAME can delegate the challenge to a zone the solver can write to.
- Add control plane node toleration to CA injector deployment.
- Add Vertical Pod Autoscaler (VPA) support for webhook pods.
- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.
- Add PodLogs for log collection.

#### Changed

- Use `imagePullPolicy: IfNotPresent` for the `cert-manager-giantswarm-clusterissuer` hook Job, so a retried Helm hook reuses the image already present on the node instead of downloading it again.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Increase clusterIssuer job memory limit from 250Mi to 512Mi.
- The ACME HTTP-01 solver of the `letsencrypt-giantswarm` ClusterIssuer now uses Gateway API (`gatewayHTTPRoute`) instead of an Ingress, and targets Envoy Gateway.
- Upgrade `docker-kubectl` image to support arm64 architecture
- Updated `cert-manager` to upstream version `v1.20.3`.
- **Notes:** `cert-manager-edit` ClusterRole no longer grants `create` on `challenges.acme.cert-manager.io`, nor `create`/`patch`/`update` on `orders.acme.cert-manager.io`.
- Improved proxy settings by adding a proxy ConfigMap and setting upstream `envFrom` values for `controller`, `webhook` and `cainjector`.
- **Breaking:** Helm values to be passed to the upstream `cert-manager` chart will now need to use the `cert-manager` path instead of root. For example, the value `crds.enabled: true` must now be set with `cert-manager.crds.enabled: true`.
- Moved vendored chart to `helm/cert-manager/charts/` and adapted sync scripts to follow new structure.
- Upgrade cert-manager to v1.19.4.

#### Removed

- `giantSwarmClusterIssuer.acme.http01.enabled` and `giantSwarmClusterIssuer.acme.http01.ingressClassName` (which defaulted to `nginx`). Setting either now fails the Helm render with a message pointing at the replacement values.
- Remove PodSecurityPolicy (PSP) and related resources.
- Remove Giant Swarm PSP to PSS migration logic.

#### Fixed

- Trim the `helm.sh/chart` label of the Giant Swarm resources (VPAs, PodLogs, proxy ConfigMap) to end on an alphanumeric character. A long dev version cut at 63 characters could end on a `.`, and the API server rejected the resource.
- Fix `controller` Vertical Pod Autoscaler (VPA) resource syntax.

### cert-manager-crossplane-resources [v0.1.0...v0.3.0](https://github.com/giantswarm/cert-manager-crossplane-resources/compare/v0.1.0...v0.3.0)

#### Added

- Read the AWS account ID, partition, and OIDC domains from the `<clusterName>-crossplane-config` ConfigMap if `providers.aws.accountID`, `providers.aws.partition`, or `providers.aws.oidc.domain` is empty. This supports EKS clusters, whose OIDC domain is only known once the cluster is provisioned. Explicitly set values take precedence over the ConfigMap.
- Add `WorkloadIdentity` resource for AKS provider.

#### Changed

- Restrict AWS provider permissions for `ChangeResourceRecordSets` to TXT records related to acme challenges.
- Create a dedicated role for azure to reduce permissions to TXT records only.
- Update `architect-orb` to v6.15.0.

### chart-operator-extensions [v1.1.2...v1.1.3](https://github.com/giantswarm/chart-operator-extensions/compare/v1.1.2...v1.1.3)

#### Changed

- Migrate Chart.yaml annotations to new format as per https://docs.giantswarm.io/reference/platform-api/chart-metadata/

### cilium [v1.3.4...v1.6.2](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.6.2)

#### Added

- Declare the `io.giantswarm.application.audience` (`all`) and
- `sync/verify-images.sh`, run as the last step of `sync/sync.sh` and therefore in CI: renders the chart across four scenarios covering every image-bearing component and fails the sync unless each image is served from `gsoci.azurecr.io/giantswarm/` or is explicitly listed in `sync/unmirrored-images.txt`. It also asserts that every configured image actually appears in a render, so a scenario that stops covering an image fails instead of silently narrowing the check. This replaces the removed patch's `fail` guards, which were the only thing that made image drift visible on an upstream bump.

#### Changed

- Upgrade Cilium to [v1.20.4](https://github.com/cilium/cilium/releases/tag/v1.20.4).
- Envoy xDS mode defaults to `split` again, following upstream ([cilium/cilium#49383](https://github.com/cilium/cilium/pull/49383)).
- Restore upstream `cilium-operator` tolerations instead of tolerating all taints.
- Exclude the Lens namespace label `k8slens-edit-resource-version` from identity computation.
- Upgrade Cilium to [v1.19.9](https://github.com/cilium/cilium/releases/tag/v1.19.9).
- Upgrade Cilium to [v1.20.2](https://github.com/cilium/cilium/releases/tag/v1.20.2).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Point `home` at this repository instead of `https://cilium.io/`, as the standard requires.
- Upgrade Cilium to [v1.19.8](https://github.com/cilium/cilium/releases/tag/v1.19.8).
- Upgrade Cilium to [v1.20.1](https://github.com/cilium/cilium/releases/tag/v1.20.1) from v1.19.7. Please review the upstream [1.20 upgrade notes](https://docs.cilium.io/en/v1.20/operations/upgrade/) before rolling this out.
- Serve the ztunnel image from `gsoci.azurecr.io/giantswarm/cilium-ztunnel` instead of pulling it from upstream. Cilium v1.20 moved this image from `docker.io/istio/ztunnel` to `quay.io/cilium/ztunnel:v1.0.0`, and our mirror carries exactly that digest, so it no longer has to be allow-listed in `sync/unmirrored-images.txt`. Only used by `encryption.type=ztunnel`, which we do not support.
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

- Upstream removed these long-deprecated Helm values in v1.20. None of them are set by this chart's defaults, and because the chart's `values.schema.json` does not reject unknown keys, leftovers in existing values are **silently ignored** rather than rejected — check your values before upgrading:
  - `encryption.strictMode.{enabled,cidr,allowRemoteNodeIdentities}` → use `encryption.strictMode.egress.*`
  - `encryption.ipsec.encryptedOverlay`
  - `clustermesh.enableMCSAPISupport` → use `clustermesh.mcsapi.enabled` (MCS-API is now stable upstream)
  - `clustermesh.apiserver.tls.{server,admin,remote}.{cert,key}` and `clustermesh.apiserver.tls.enableSecrets` → enable auto-generation or pre-create the secrets
  - `hubble.redact.kafka.apiKey` → Kafka-aware L7 policy support and proxylib were removed upstream
  - `preflight.tofqdnsPreCache` → the preflight FQDN poller was removed upstream
  - `hubble.ui.backend.{livenessProbe,readinessProbe}.enabled`
- `sync/patches/certgen/`. Cilium v1.20 ships the `certgen.enforceCAValidityThroughoutLeavesDuration` value and wires `--ca-enforce-validity-throughout-leaves-duration` into both certgen job specs itself, so the Giant Swarm patch that added them became a no-op (it detected this and skipped). The default stays `true` and the rendered job specs are unchanged, so two more patches drop out of `diffs/`.
- **Removed the `image.registry` value and the sync patch behind it** ([roadmap#3264](https://github.com/giantswarm/roadmap/issues/3264)). The registry is now part of each image's `repository` value, exactly as upstream ships it, so the chart no longer patches upstream's `cilium.image`/`cilium.operator.image` helpers into a `(list $ <image>)` signature and no longer rewrites all 35 call sites with `sed`. `helm/cilium/templates/_helpers.tpl` and `helm/cilium/templates/cilium-operator/_helpers.tpl` are now byte-identical to upstream, and 15 template patches disappear from `diffs/`.

#### Fixed

- Fix four image references that were rendered as unpullable double-prefixed paths, because their `repository` already carried a registry while the removed patch prefixed `image.registry` on top: `gsoci.azurecr.io/ghcr.io/spiffe/spire-server`, `.../spire-agent`, `.../docker.io/library/busybox` (SPIRE mutual authentication) and `.../docker.io/istio/ztunnel` (`encryption.type=ztunnel`). All four now resolve. Latent until now because both features are disabled by default.
- Fix rendering of the certgen job specs (`hubble/tls-cronjob` and `clustermesh-apiserver/tls-cronjob`): the image reference was not converted to the Giant Swarm `cilium.image` helper signature because the image-registries sync patch only processed `*.yaml` templates, so enabling the cronJob method failed with `required list, but got "map"`.
- Add NetworkPolicies allowing the `hubble-generate-certs` and `clustermesh-apiserver-generate-certs` certgen pods egress to the Kubernetes API.

### cilium-servicemonitors [v0.1.3...v0.2.0](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.3...v0.2.0)

#### Added

- Add per-monitor `enabled` flag for the agent, hubble and operator.

#### Changed

- Switch all monitors from ServiceMonitor to PodMonitor.
- Migrate chart metadata annotations

### cluster-autoscaler [v1.34.1-1...v2.1.1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.34.1-1...v2.1.1)

#### Added

- Validate that `managementCluster` (when `isManagementCluster=true`) or `clusterID` (otherwise) are set, failing early with a clear error message.
- Add support for CAPI mode (`kubeconfig-incluster`): run cluster-autoscaler on the management cluster using a pre-existing kubeconfig to connect to the workload cluster.
- Add `clusterAPI` values section for configuring CAPI mode (autodiscovery, kubeconfig secret, configmaps namespace).
- Add `rbac.clusterScoped` toggle to support namespace-scoped RBAC (no ClusterRole/ClusterRoleBinding) for CAPI deployments.

#### Changed

- Chart: Update to upstream v1.36.1.
- Templates: Move PolicyException to `kube-system` namespace.
- Chart: Update to upstream v1.35.1.
- Templates: Move PolicyException to `giantswarm` namespace.
- RBAC: Grant access to DeviceClasses and ResourceClaims.
- RBAC: Grant access to `resourceslices.resource.k8s.io`.
- Helpers: Trim `.` suffix.
- Chart: Update to upstream v1.35.0.
- Migrate test infrastructure from pipenv to uv.
- Deploy the Kyverno policy exception in the `policy-exceptions` Namespace.
- Deploy the Kyverno PolicyException as a Helm `pre-install,pre-upgrade` hook so it takes effect before chart resources are created.
- Chart: Update to upstream v1.34.3.
- Chart: Update to upstream v1.34.2.

#### Fixed

- RBAC: Allow reading `infrastructure.cluster.x-k8s.io` resources when `clusterAPI.enabled` is set.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### cluster-autoscaler-crossplane-resources [v1.0.0](https://github.com/giantswarm/cluster-autoscaler-crossplane-resources/releases/tag/v1.0.0)

#### Added

- Chart: Add IAM role.

#### Changed

- CircleCI: Push to `cluster-catalog`.

### coredns [v1.30.0...v1.35.0](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.35.0)

#### Added

- Support the forward `max_age` option (#528)
- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.
- Wire up the full set of CoreDNS `forward`, `cache`, and `kubernetes` block parameters in the structured zone config:
  - `forward`: `maxIdleConns`, `maxConnectAttempts`, `dohMethod`, `tls`, `tlsServername`, `next`, `nextOnNodata`, `failfastAllUnhealthyUpstreams`, `failover`, `resolver`.
  - `cache`: `zones`, `serveStale.verifyTimeout`, `disable.successZones`, `disable.denialZones`.
  - `kubernetes`: `endpoint`, `tls`, `kubeconfig`, `apiserverQPS`, `apiserverBurst`, `apiserverMaxInflight`, `namespaceLabels`, `fallthroughZones`, `multicluster`, `startupTimeout`.

#### Changed

- Update `coredns` image to [1.15.0](https://github.com/coredns/coredns/releases/tag/v1.15.0).
- Update gsoci.azurecr.io/giantswarm/docker-kubectl to v1.37.1 (#520)
- Update `coredns` image to [1.14.7](https://github.com/coredns/coredns/releases/tag/v1.14.7).
- Update `coredns` image to [1.14.6](https://github.com/coredns/coredns/releases/tag/v1.14.6).
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Chart: Make tolerations configurable.
- Rebuild with `app-build-suite` 2.2.0 (via `architect` orb 9.6.0): the packaged chart now carries Artifact Hub metadata (`artifacthub.io/license` and a Support link). No functional chart changes.
- Standardize `values.yaml` comments to the `# @schema` / `# --` (helm-docs) convention and remove section-header dividers, so `values.schema.json` and the chart `README.md` are generated from the values file. Rendered manifests are unchanged.
- Update `coredns` image to [1.14.4](https://github.com/coredns/coredns/releases/tag/v1.14.4).
- update CoreDNS icon to light version
- Update `coredns` image to [1.14.3](https://github.com/coredns/coredns/releases/tag/v1.14.3).

#### Fixed

- Render forward integer options when set to `0` (#528)
- Validate `additionalZones` cache, forward, kubernetes, log and loadbalance blocks against the public and cluster zone schemas (#528)
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
- Honor the deprecated `configmap.log`, `loadbalancePolicy` and `configmap.cache` again. Since 1.31.0 `coredns.<zone>.log`, `coredns.<zone>.loadbalance` and `coredns.<zone>.cache.success.ttl` shipped defaults that shadowed them, so the old keys were silently ignored. They are now unset by default, restoring the documented fallback chain. Rendering with default values is unchanged.
- Render the `health` directive in only the `.` server block. The health plugin is process-wide and can be enabled in just one Server Block, so emitting it in every zone block was invalid. `ready` is kept in every block (its readiness is aggregated across blocks).
- Correct the `coredns.*.cache.serveStale.refreshMode` schema enum to `immediate`/`verify` (was `immediate`/`background`), matching the CoreDNS cache plugin.

### external-dns [v3.4.0...v3.7.1](https://github.com/giantswarm/external-dns-app/compare/v3.4.0...v3.7.1)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Changed

- Upgrade external-dns to [v0.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.23.0).
- Sync to upstream helm chart [1.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.23.0).
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Upgrade external-dns to [v0.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.22.0).
- Sync to upstream helm chart [1.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.22.0).
  - Add `replicaCount` value to scale the deployment down to `0` or back to `1`.
  - Add `service.enabled` value to skip creating the `Service`.
  - Add `hostAliases` value to inject entries into the pod's `/etc/hosts`.
  - Add `crd` as a valid `registry` value, with the matching RBAC on `dnsrecords`.
  - Install the new `DNSRecord` CRD.
  - Narrow the RBAC verbs on `dnsendpoints/status` from `*` to `update`.
  - `policy` is now a required value. Our default of `sync` is unchanged.
  - Pin `annotationPrefix` to `external-dns.alpha.kubernetes.io/`, keeping the previous annotation prefix after upstream changed the default.
- Update VPA `updatePolicy.updateMode` from deprecated `Auto` to `Recreate`.
- Upgrade external-dns to [v0.21.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.21.0).
- Sync to upstream helm chart [1.21.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.21.1).
  - Add `namespaceOverride` value to support deploying external-dns into a different namespace than the Helm release (useful for subchart usage).
  - Add `enableGatewayListenerSets` value to opt into Gateway API ListenerSet resource support.
  - Add `sourceNamespace` value (used with `namespaced=true`) to watch resources in a namespace different from the deployment namespace.
  - Avoid creating cluster-scoped namespace RBAC when `gatewayNamespace` is set, reducing required permissions.
  - Fix `extraArgs` map handling: boolean values now render as `--flag` / `--no-flag` and string values are properly quoted.
- Use external-dns.namespace in VPA and NetworkPolicy resources.

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0, whose fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
- Values schema: accept arbitrary keys again in free-form maps and pass-through Kubernetes objects (annotations, labels, `nodeSelector`, `affinity`, security contexts, probes, resources, `extraArgs` as a map, `tlsConfig`, `secretConfiguration.data`), and `name` in `imagePullSecrets` items. This broke installs, for example with the IRSA `eks.amazonaws.com/role-arn` service account annotation on CAPA.
- Values schema: accept the `aws` values that cluster-aws still sets. They have no effect.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### external-dns-crossplane-resources [v0.4.0](https://github.com/giantswarm/external-dns-crossplane-resources/releases/tag/v0.4.0)

#### Added

- Add configurable `serviceAccount.name`/`serviceAccount.namespace` and `azure.mountPath` values.

#### Changed

- Scope AWS values under `aws.*`. Root-level values remain supported as a fallback for backward compatibility.
- Templatize hardcoded resource, service account and tag values.
- Scope the IAM trust policy to the exact `serviceAccount` subject instead of a wildcard match.
- Enable Azure Workload Identity for the `azure` provider in addition to `aks`.
- Update chart metadata to reflect Azure support.

### k8s-dns-node-cache [v2.9.1...v2.11.1](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.1...v2.11.1)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.
- Add `configmap.log.enabled` helm value to toggle CoreDNS query logging (default: `false`).
- Make `AAAA NOERROR` configurable for IPv6.

#### Changed

- Upgrade application to version 1.26.7 (includes coredns 1.13.1)

#### Fixed

- Make the DNS Service redirected by the `CiliumLocalRedirectPolicy` configurable.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### karpenter [v1.4.0...v2.6.0](https://github.com/giantswarm/karpenter-app/compare/v1.4.0...v2.6.0)

#### Added

- Add `ec2:DescribeInstanceStatus` permission to the Karpenter IAM role, which Karpenter requires since v1.12 for its interruption health checks.
- Add `iam:CreateServiceLinkedRole` permission to the Karpenter IAM role, scoped to the `AWSServiceRoleForEC2Spot` service-linked role, so Karpenter can create such role when launching spot instances in accounts where it does not exist yet.
- Support karpenter on EKS clusters
  - Add `eks:DescribeCluster` permission to discover the cluster endpoint
  - Conditionally depend on Kyverno and pod logs CRDs since security-bundle isn't yet implemented for cluster-eks and customers may disable that bundle and also observability-bundle for cluster adoption cases
- HelmRelease: Add dependencies.
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

- Update upstream Karpenter and its CRDs to v1.14.1. This might roll nodes because Karpenter v1.12 bumps the `EC2NodeClass` hash version from v4 to v5, which can mark existing nodes as drifted.
- Enable the `StaticCapacity` feature gate (`NodePool.spec.replicas`) in the `karpenter` chart defaults, as already done in `karpenter-bundle`. The feature is still alpha upstream, and we offer it as tech preview.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Bump OCIRepository to v1.
- Chart: Move PolicyException to `kube-system` to avoid depending on creation of the `policy-exceptions` namespace.
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

- Removed `global.podSecurityStandards.enforced` helm value. The Kyverno `PolicyException` in `pss-exceptions.yaml` is now rendered unconditionally.
- Remove all forked upstream templates from workload chart (replaced by upstream dependency).
- Remove `vendir.yml`, `vendir.lock.yml`, `vendor/` directory, and `Makefile.custom.mk`.

#### Fixed

- Strip `-`, `_`, and `.` from the `helm.sh/chart` label to support dev versions
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

### net-exporter [v1.23.0...v1.24.1](https://github.com/giantswarm/net-exporter/compare/v1.23.0...v1.24.1)

#### Added

- Add `keywords` to `Chart.yaml`.
- Add the `io.giantswarm.application.audience` (`all`) and `io.giantswarm.application.managed`

#### Changed

- Replace `interface{}` with `any` and use for-range over integers (Go modernization).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Set `Chart.yaml` `apiVersion` to `v2`. The chart declares no dependencies and had no
- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the net-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Bump `docker-kubectl` init container from `1.25.4` to `1.36.0`.

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### network-policies [v0.1.3...v0.5.0](https://github.com/giantswarm/network-policies-app/compare/v0.1.3...v0.5.0)

#### Added

- Add `hubbleRelay` to allow the AKS ACNS `hubble-relay` through the default deny.
- Add `allowEgressToUpstreamDNS` to allow pods with `dnsPolicy: Default` to reach upstream DNS resolvers.
- Add `azureWorkloadIdentityWebhook` to allow the AKS workload identity webhook through the default deny.
- Add `keywords` to `Chart.yaml`.
- Declare the `io.giantswarm.application.audience` (`all`) and
- Add optional `denyEgressToIMDS` policy denying pod egress to the instance metadata service. Disabled by default.
- Add apptest-framework e2e test suite.
- Add support for AKS selector labels.

#### Changed

- Move the chart to `helm/network-policies` to match its name, in preparation for devctl-generated CI.
- Convert the `allow-ingress-from-konnectivity` `CiliumNetworkPolicy` into a `CiliumClusterwideNetworkPolicy` named
- Always exclude the `karpenter` and `aws-load-balancer-controller` namespaces from `denyEgressToIMDS`.
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Deprecated the .Values.kamaji in favour of the more generic .Values.konnectivityAgent to control the behaviour for the `konnectivity-agent`.

### node-exporter [v1.20.10...v1.21.0](https://github.com/giantswarm/node-exporter-app/compare/v1.20.10...v1.21.0)

#### Added

- `disableSystemdCollector` value to turn the systemd collector off, mirroring the existing `disableConntrackCollector` and `disableNvmeCollector` toggles. The collector needs a D-Bus connection to the host, which is refused on nodes where AppArmor mediates D-Bus (such as AKS Ubuntu nodes running under the default containerd profile), making it fail on every scrape. Defaults to `false`, so behaviour is unchanged.

#### Changed

- CircleCI: Do not override app version.
- Chart: Move PolicyException to `kube-system` namespace.
- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

#### Fixed

- Removed duplicated `app` label which is already added by the selector helper.

### observability-bundle [v2.5.0...v3.6.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v3.6.0)

#### Added

- KSM custom resource state: Add `status_parent_accepted` and `status_parent_resolved_refs` gauges for Gateway API `HTTPRoute`, one series per parent with the condition `reason` as label.
- KSM custom resource state: Add `target_info` and `status_ancestor_accepted` metrics for Envoy Gateway `SecurityPolicy`, `BackendTrafficPolicy`, `ClientTrafficPolicy` and `EnvoyExtensionPolicy`.
- Point kube-prometheus-stack's control-plane ServiceMonitors at the alloy-metrics token Secret.
- Add KSM metrics for Gateway API `ListenerSet` and `ReferenceGrant` resources.
- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle
- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.
- Add KSM metrics for Gateway API resources.

#### Changed

- *!Breaking change*: Turn `apps.<appName>.dependsOn` from a list into a map from dependency name to `true`/`false`, so that a provider can turn off a single dependency (for example `kyverno-crds` on clusters without the security bundle) instead of overriding the whole list.
- Update `alloy` apps to 0.23.1 (Alloy v1.19.2).
- Update `prometheus-operator-crd` to 24.0.0 (Prometheus Operator CRDs v0.94.0).
- Update `kube-prometheus-stack` to 24.0.0 (chart 91.2.3, Prometheus Operator v0.94.0).
- Values: Update Prometheus Operator CRD and Kube Prometheus Stack to v23.0.0.
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

#### Fixed

- KSM custom resource state: Set the Gateway API `TCPRoute` and `UDPRoute` collectors to `v1`, which is the version the API server serves.

### observability-policies [v0.0.3...v0.0.4](https://github.com/giantswarm/observability-policies-app/compare/v0.0.3...v0.0.4)

#### Changed

- Rename app to `observability-policies`
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).

### priority-classes [v0.3.0...v0.3.1](https://github.com/giantswarm/priority-classes/compare/v0.3.0...v0.3.1)

#### Fixed

- Sanitize `Chart.Version` used in labels. This is needed because flux apapends the digest to the version using the `+` character which is not allowed in labels.

### prometheus-blackbox-exporter [v0.5.0...v0.10.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.5.0...v0.10.0)

#### Added

- Probe `github.com`, `gsoci.azurecr.io` and `grafana.com` from every node (targets `egress-github`, `egress-registry`, `egress-grafana`), so an allowlist-style firewall change blocking a single domain becomes visible. The new `serviceMonitor.externalTargets` value is a map so per-installation, per-region or per-customer overrides can disable, change or add individual entries without copying the whole list. A separate `serviceMonitor.additionalExternalTargets` key takes regional/customer additions, structurally separated from the Giant Swarm defaults. Adds the `http_2xx_or_401` module for registry endpoints that answer unauthenticated requests with 401. See giantswarm/giantswarm#33409.
- Add the `http_2xx_egress` module, used by the internet egress targets. It carries a 15s timeout so `probe_success` reports whether an endpoint is reachable rather than whether it is fast. It is a separate module rather than a longer timeout on `http_2xx` because several installations pin `http_2xx` in their own custom values, which would silently revert the change there.
- Add VPA for `blackbox-exporter`. Uses `updateMode: Initial` for DaemonSet and `updateMode: Auto` for Deployment.
- Add toleration for `kubernetes.io/arch=arm64:NoSchedule` so the DaemonSet schedules on ARM worker nodes.
- Add `http_2xx_insecure` module with `insecure_skip_verify: true` to support probing workload cluster API servers from the management cluster. The MC's service account CA (`http_2xx_k8sca`) only covers the MC itself; workload clusters have their own CA which is not available to the blackbox exporter, making TLS verification impossible without this module.

#### Changed

- Set `priorityClassName` to `system-node-critical` to ensure DaemonSet pods are scheduled even on full nodes.
- Migrate to App Build Suite (ABS) for Helm chart building.

#### Removed

- Remove the inert `instance` metric relabeling from the ServiceMonitor template. It interpolated a `url` field that no target defines, so it rendered empty and Prometheus fell back to its `$1` default, leaving `instance` unchanged.

#### Fixed

- Give the internet egress targets a probe deadline above the cross-border baseline: `scrapeTimeout: 20s` on `http-giantswarm`, `egress-github`, `egress-registry` and `egress-grafana`, and a 15s timeout on `http_2xx_or_401`. The exporter applies `min(module timeout, scrapeTimeout - 0.5s offset)`, so the 5s `serviceMonitor.defaults.scrapeTimeout` capped every probe at a 4.5s deadline and raising a module timeout alone had no effect. Installations whose baseline latency is a large fraction of that deadline crossed it on endpoints that were still returning HTTP 200, making `probe_success` report latency rather than reachability.
- Point the `dns-tcp-internal` and `dns-udp-internal` ServiceMonitors at the `dns_*_internal` modules. They referenced the `_external` modules, so both probed `www.prometheus.io` and in-cluster DNS resolution was never monitored.
- Add `probe_target` label to ensure unique synthetic metrics

### security-bundle [v1.16.1...v2.4.0](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v2.4.0)

#### Added

- Add e2e scenarios covering trivy-operator `VulnerabilityReport` creation, starboard-exporter metrics for that report, kyverno restricted PSS enforcement, and kyverno-policy-operator `PolicyException` translation.
- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.

#### Changed

- Update `exception-recommender` (app) to v0.3.0.
- Update `falco` (app) to v0.13.0.
- Update `jiralert` (app) to v0.1.4.
- Update `kubescape` (app) to v0.1.1.
- Update `kyverno-policies` (app) to v0.27.1.
- Update `policy-api` (app) to v0.0.12.
- Update `starboard-exporter` (app) to v1.2.15.
- Update `trivy` (app) to v0.18.0.
- Update `trivy-operator` (app) to v0.15.0.
- Update `kyverno-policy-operator` (app) to v0.2.3.
- Update `policy-api` (app) to v0.0.9.
- Update `starboard-exporter` (app) to v1.2.3.
- Update `trivy` (app) to v0.17.0.
- Update `trivy-operator` (app) to v0.13.3.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Values: Make Kyverno and Kyverno Policy Operator depend on Cilium.
- Update `kyverno-policies` (app) to v0.26.1.
- Update `starboard-exporter` (app) to v1.1.4.
- Update `trivy` (app) to v0.16.0.
- Update `trivy-operator` (app) to v0.13.2.
- Update `cloudnative-pg` (app) to v0.1.0.
- Update `trivy` (app) to v0.15.0.
- Update `falco` (app) to v0.12.0.
- HelmReleases: honor the App platform `priority` field (1-150, default 25) on `extraConfigs` entries. `spec.valuesFrom` now reproduces the App platform merge order — all configMaps before all secrets (a secret always overrides a configMap), each kind ordered by priority around the user-config layer — preserving the App CR merge semantics after the migration. ([giantswarm#36096](https://github.com/giantswarm/giantswarm/issues/36096))
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- No longer pass the 'cluster-values' ConfigMap to the applications inside the bundle.
- Update `kyverno` (app) to v0.24.2.
  - This release includes a new Kyverno minor version. Please refer to the upstream release notes for the latest changes:
  - https://github.com/kyverno/kyverno/releases/tag/v1.17.0
- Update `kyverno-crds` (app) to v1.17.0.
- Update `kyverno-policies` (app) to v0.25.0.
- Update `kyverno-policy-operator` (app) to v0.2.2.
- Update `kubescape` (app) to v0.1.0.
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

#### Removed

- Remove `gel` (app).

#### Fixed

- Give `kubescape` a 15m install and upgrade timeout. It does not finish installing within Flux's 5m default, so it failed with `context deadline exceeded` and then retried indefinitely.
- Set `createNamespace` on every app in the bundle, so each one creates its target namespace instead of relying on another app to have created it first. Previously `kubescape` failed with `namespaces "kubescape" not found`, and the apps targeting `security-bundle` could only install after `kyverno-policy-operator` had created it.

### teleport-kube-agent [v0.10.8...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
- Values: Tolerate `node.cloudprovider.kubernetes.io/uninitialized`.
- Values: Ignore taints regardless of value.
- Values: Pass HTTP proxy settings to sub-chart.
- Updated `teleport-kube-agent` to upstream version `v18.7.6`.

### vertical-pod-autoscaler [v6.1.1...v6.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.1.1...v6.1.2)

#### Fixed

- Pushed helm chart to OCI repository.

### vertical-pod-autoscaler-crd [v4.1.1...v4.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.1.1...v4.1.2)

#### Fixed

- Pushed helm chart to OCI repository.
