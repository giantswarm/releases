# :zap: Giant Swarm Release v35.0.0 for CAPA :zap:

## Changes compared to v34.3.0

### Components

- cluster-aws from v7.7.1 to v8.5.0
- cluster from v5.3.1 to v6.5.0
- Kubernetes from v1.34.7 to [v1.35.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.5)

### cluster-aws [v7.7.1...v8.5.0](https://github.com/giantswarm/cluster-aws/compare/v7.7.1...v8.5.0)

#### Added

- Add optional `architecture` field to node pools (enum `x86_64` / `arm64`, default `x86_64`). When set to `arm64`, `imageLookupFormat` is rendered with an `arm64-` infix so the pool resolves the matching CAPI Flatcar arm64 AMI. Operators are responsible for adding the `kubernetes.io/arch=arm64:NoSchedule` taint via `customNodeTaints` on arm64 pools so amd64-only workloads don't land there. Control plane and existing x86_64 pools are unaffected.
- Apply tags to subnets if they're defined by ID. Clusters that were migrated from Vintage to CAPA have the subnet IDs defined in the values, and we previously didn't render tags for them, so tags did not get reconciled.
- Add `external-dns-crossplane-resources` HelmRelease to manage Route53 records via Crossplane, injecting `clusterName`, `accountID`, `baseDomain`, `oidcDomain`, and `oidcDomains`.

#### Changed

- Chart: Fix validation errors.
- Support newer Flatcar versions which require a larger root volume size. For ease of migration, enforce at least 15 GB even if a smaller, explicit size is specified in chart values.
- Fix `external-dns` IAM role ARN annotation to use `{name}-external-dns` instead of `{name}-Route53Manager-Role`.
- Extract `aws-oidc-domain` Helm helper (China-aware) and use it in `cert-manager-crossplane-resources` and `external-dns-crossplane-resources` templates.

#### Removed

- Chart: Remove unused `cluster-shared` library chart dependency.

#### Fixed

- Allow numbers in Karpenter pool CPU limit (`global.nodePools.*.limits.cpu`).

### cluster [v5.3.1...v6.5.0](https://github.com/giantswarm/cluster/compare/v5.3.1...v6.5.0)

#### Added

- Apps: Add External DNS Crossplane Resources.
- Apps: Deploy `cluster-autoscaler` inCluster in Azure.
- MachineDeployment: Add CAPI autoscaler annotations (`cluster-api-autoscaler-node-group-min-size`/`max-size`) when `minSize`/`maxSize` are set on a node pool (only in Azure).
- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))

#### Changed

- Control Plane: Make etcd image tag configurable. ([#841](https://github.com/giantswarm/cluster/pull/841))
- Chart: Require `global.release.version` if using Releases to give a better rendering error message.
- Chart: Fix validation errors.
- Configure `observability-bundle` with the management cluster name.
- Apps: Skip `kyverno-crds` dependency for `cluster-autoscaler` when deployed inCluster.
- Apps: Add cluster-probes HelmRelease to deploy ServiceMonitors for probing workload cluster API server endpoint from the management cluster. Configurable via `global.apps.clusterProbes` with default module `http_2xx_insecure` for self-signed certificates.
- Helpers: Use `.Chart.AppVersion` in `app.kubernetes.io/version` label.
- Cluster API: Migrate to API `v1beta2`.

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.

### Apps

- aws-ebs-csi-driver from v4.1.2 to v4.3.0
- aws-nth-bundle from v1.4.0 to v2.0.1
- aws-pod-identity-webhook from v2.2.0 to v2.3.0
- cert-exporter from v2.10.1 to v2.11.0
- cilium from v1.4.3 to v1.4.4
- cluster-autoscaler from v1.34.3-2 to v1.35.0-1
- Added cluster-autoscaler-crossplane-resources v1.0.0
- coredns from v1.30.0 to v1.30.1
- etcd-defrag from v1.2.6 to v1.2.7
- external-dns from v3.4.0 to v3.5.0
- Added external-dns-crossplane-resources v0.2.0
- karpenter from v2.3.0 to v2.4.0
- karpenter-taint-remover from v1.0.2 to v1.1.0
- net-exporter from v1.23.1 to v1.24.0
- observability-bundle from v2.8.0 to v2.9.0
- prometheus-blackbox-exporter from v0.7.0 to v0.8.0
- Added rbac-bootstrap v0.3.0
- teleport-kube-agent from v0.10.8 to v0.11.0

### aws-ebs-csi-driver [v4.1.2...v4.3.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.2...v4.3.0)

#### :warning: Breaking Changes

- **Workload chart renamed** from `aws-ebs-csi-driver-app` to `aws-ebs-csi-driver`. The OCI catalog artifact name changes accordingly.
- **Bundle values restructured**: upstream chart values are now under the `upstream:` key in the bundle `values.yaml`. The `giantswarm.workloadValues` helper handles the transformation automatically, so bundle users only need to place overrides under `upstream:` in their App CR ConfigMap.
- **Direct workload chart install**: if installing the workload chart directly (without the bundle), all upstream values must be under the `upstream:` key, and extras (`verticalPodAutoscaler`, `networkPolicy`, `global.podSecurityStandards`) are at the top level.

#### Added

- Add IRSA environment variables (`AWS_ROLE_ARN`, `AWS_WEB_IDENTITY_TOKEN_FILE`), projected ServiceAccountToken volume, and `AWS_REGION` to the EBS CSI controller, enabling IRSA authentication in CAPA clusters.
- Propagate proxy values from the bundle (`proxy.http`, `proxy.noProxy`) to the upstream chart (`proxy.http_proxy`, `proxy.no_proxy`) when set.

#### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format for both the app and bundle charts.

#### Fixed

- Re-enable metrics, force use of `ServiceMonitor` to avoid rendering without them if CRDs are not installed yet
- Fix VPA `updateMode` for `ebs-csi-node` DaemonSet from `Auto` to `Initial`. VPA cannot evict DaemonSet pods, so `Auto` mode silently produces recommendations without ever applying them. `Initial` correctly sets resources at pod creation time.

### aws-nth-bundle [v1.4.0...v2.0.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.4.0...v2.0.1)

#### Added

- Make the `aws-node-termination-handler` HelmRelease `dependsOn` list configurable via `awsNodeTerminationHandler.dependsOn`.
- Add `cluster.x-k8s.io/cluster-name` label to the HelmReleases.
- Add Flux `dependsOn` from the `aws-node-termination-handler` HelmRelease to the `prometheus-operator-crd` HelmRelease.

#### Changed

- Deploy Crossplane resources directly from the bundle instead of standalone helmrelease.
- Clean up values and helm template functions.
- Remove `cluster-values` ConfigMap reference from `aws-node-termination-handler` HelmRelease. Pass `clusterID` explicitly via inline values instead.

### aws-pod-identity-webhook [v2.2.0...v2.3.0](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v2.2.0...v2.3.0)

#### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.
- Use Helm `appVersion` as the default version for the restarter container image.

### cert-exporter [v2.10.1...v2.11.0](https://github.com/giantswarm/cert-exporter/compare/v2.10.1...v2.11.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.

### cilium [v1.4.3...v1.4.4](https://github.com/giantswarm/cilium-app/compare/v1.4.3...v1.4.4)

#### Changed

- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).

### cluster-autoscaler [v1.34.3-2...v1.35.0-1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.34.3-2...v1.35.0-1)

#### Changed

- Chart: Update to upstream v1.35.0.

### cluster-autoscaler-crossplane-resources [v1.0.0](https://github.com/giantswarm/cluster-autoscaler-crossplane-resources/releases/tag/v1.0.0)

#### Added

- Chart: Add IAM role.

#### Changed

- CircleCI: Push to `cluster-catalog`.

### coredns [v1.30.0...v1.30.1](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.30.1)

#### Changed

- Update `coredns` image to [1.14.3](https://github.com/coredns/coredns/releases/tag/v1.14.3).

### etcd-defrag [v1.2.6...v1.2.7](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.6...v1.2.7)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.40.0. ([#94](https://github.com/giantswarm/etcd-defrag-app/pull/94))

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

### external-dns-crossplane-resources [v0.2.0](https://github.com/giantswarm/external-dns-crossplane-resources/releases/tag/v0.2.0)

#### Added

- Add `awsPartition` value to support non-standard AWS partitions (e.g. China).
- Add `oidcDomain` and `oidcDomains` values to support multiple OIDC providers; the IAM trust policy now iterates over all provided domains.
- Add `managed-by: external-dns-crossplane-resources` tag to the IAM role.

#### Changed

- Add `io.giantswarm.application.audience: giantswarm` annotation.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.
- Move `required` validation for `awsPartition` and `accountID` into named template helpers.
- Replace deprecated `app.giantswarm.io/branch`, `application.kubernetes.io/managed-by`, and `giantswarm.io/service-type` labels with `app.kubernetes.io/*` equivalents.
- Update chart description, icon, and owning team (`cabbage`).

#### Removed

- Remove `name`, `serviceType`, `project.branch`, and `project.commit` values that were leftover hackathon scaffolding.
- Remove hardcoded `baseDomain`-based OIDC provider URL; replaced by `oidcDomain`/`oidcDomains`.

### karpenter [v2.3.0...v2.4.0](https://github.com/giantswarm/karpenter-app/compare/v2.3.0...v2.4.0)

#### Added

- Add `cluster.x-k8s.io/cluster-name` label to the karpenter HelmRelease.
- Add `iam:GetInstanceProfile` permission to Karpenter IAM role.
- Add karpenter CRDs.
- Set `helm.sh/resource-policy: keep` on the karpenter CRDs so they survive HelmRelease uninstall and prevent cascade-deleting `NodePool`/`NodeClaim`/`EC2NodeClass` resources.

#### Changed

- Switch e2e scale test from App CR to Flux HelmRelease for deploying hello-world, avoiding `values-schema-violation` errors caused by app-platform injected properties.
- Improve Crossplane ConfigMap fetching logic

### karpenter-taint-remover [v1.0.2...v1.1.0](https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.2...v1.1.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image.
- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.

### net-exporter [v1.23.1...v1.24.0](https://github.com/giantswarm/net-exporter/compare/v1.23.1...v1.24.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the net-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Bump `docker-kubectl` init container from `1.25.4` to `1.36.0`.

### observability-bundle [v2.8.0...v2.9.0](https://github.com/giantswarm/observability-bundle/compare/v2.8.0...v2.9.0)

#### Added

- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle

#### Changed

- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0

### prometheus-blackbox-exporter [v0.7.0...v0.8.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.7.0...v0.8.0)

#### Added

- Add toleration for `kubernetes.io/arch=arm64:NoSchedule` so the DaemonSet schedules on ARM worker nodes.

### rbac-bootstrap [v0.3.0](https://github.com/giantswarm/rbac-bootstrap-app/releases/tag/v0.3.0)

#### Added

- Add `io.giantswarm.application.managed` chart annotation for Backstage visibility.
- Add optional `cluster-reader` ClusterRole (off by default, enabled via `clusterReader.enabled: true`) that aggregates into the built-in `view` ClusterRole and grants read access (`get`/`list`/`watch`) on cluster-scoped resources.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.

### teleport-kube-agent [v0.10.8...v0.11.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.7.6`.
