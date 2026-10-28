# :zap: Giant Swarm Release v35.4.0 for CAPA :zap:

## Changes compared to v35.3.0

### Components

- cluster-aws from v10.5.0 to v10.6.0

### cluster-aws [v10.5.0...v10.6.0](https://github.com/giantswarm/cluster-aws/compare/v10.5.0...v10.6.0)

#### Added

- Add `global.nodePools.<pool>.replicas` for Karpenter node pools to keep a fixed number of nodes (static node pool). This is a tech preview of the alpha Karpenter `StaticCapacity` feature and requires a release with Karpenter v1.14.1 and the matching `aws-resolver-rules-operator`.
- Add `global.nodePools.<pool>.limits.nodes` to cap the number of nodes of a Karpenter node pool.
- Add the `Balanced` consolidation policy for Karpenter node pools.

#### Fixed

- Default `limits.cpu` and `limits.memory` of Karpenter node pools separately, so setting only one of them no longer renders the other one empty.

### Apps

- cilium from v1.5.2 to v1.5.4
- cloud-provider-aws from v2.1.0 to v2.2.0
- cluster-autoscaler from v2.0.5 to v2.1.1
- coredns from v1.34.0 to v1.35.0
- etcd-defrag from v1.2.12 to v1.2.13
- external-dns from v3.6.1 to v3.7.1
- k8s-dns-node-cache from v2.11.0 to v2.11.1
- karpenter from v2.5.0 to v2.6.0
- network-policies from v0.3.2 to v0.5.0
- observability-bundle from v3.5.0 to v3.6.0

### cilium [v1.5.2...v1.5.4](https://github.com/giantswarm/cilium-app/compare/v1.5.2...v1.5.4)

#### Changed

- Upgrade Cilium to [v1.19.10](https://github.com/cilium/cilium/releases/tag/v1.19.10).
- Upgrade Cilium to [v1.19.9](https://github.com/cilium/cilium/releases/tag/v1.19.9).
- Exclude the Lens namespace label `k8slens-edit-resource-version` from identity computation.

### cloud-provider-aws [v2.1.0...v2.2.0](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v2.1.0...v2.2.0)

#### Added

- Add E2E tests

#### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.
- Chart: Update to upstream v1.36.1.

### cluster-autoscaler [v2.0.5...v2.1.1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v2.0.5...v2.1.1)

#### Changed

- Chart: Update to upstream v1.36.1.

#### Fixed

- RBAC: Allow reading `infrastructure.cluster.x-k8s.io` resources when `clusterAPI.enabled` is set.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### coredns [v1.34.0...v1.35.0](https://github.com/giantswarm/coredns-app/compare/v1.34.0...v1.35.0)

#### Added

- Support the forward `max_age` option (#528)

#### Changed

- Update `coredns` image to [1.15.0](https://github.com/coredns/coredns/releases/tag/v1.15.0).
- Update gsoci.azurecr.io/giantswarm/docker-kubectl to v1.37.1 (#520)

#### Fixed

- Render forward integer options when set to `0` (#528)
- Validate `additionalZones` cache, forward, kubernetes, log and loadbalance blocks against the public and cluster zone schemas (#528)

### etcd-defrag [v1.2.12...v1.2.13](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.12...v1.2.13)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.46.0. ([#156](https://github.com/giantswarm/etcd-defrag-app/pull/156))

### external-dns [v3.6.1...v3.7.1](https://github.com/giantswarm/external-dns-app/compare/v3.6.1...v3.7.1)

#### Changed

- Upgrade external-dns to [v0.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.23.0).
- Sync to upstream helm chart [1.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.23.0).

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0, whose fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
- Values schema: accept arbitrary keys again in free-form maps and pass-through Kubernetes objects (annotations, labels, `nodeSelector`, `affinity`, security contexts, probes, resources, `extraArgs` as a map, `tlsConfig`, `secretConfiguration.data`), and `name` in `imagePullSecrets` items. This broke installs, for example with the IRSA `eks.amazonaws.com/role-arn` service account annotation on CAPA.
- Values schema: accept the `aws` values that cluster-aws still sets. They have no effect.

### k8s-dns-node-cache [v2.11.0...v2.11.1](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.11.0...v2.11.1)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Fixed

- Make the DNS Service redirected by the `CiliumLocalRedirectPolicy` configurable.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### karpenter [v2.5.0...v2.6.0](https://github.com/giantswarm/karpenter-app/compare/v2.5.0...v2.6.0)

#### Added

- Add `ec2:DescribeInstanceStatus` permission to the Karpenter IAM role, which Karpenter requires since v1.12 for its interruption health checks.

#### Changed

- Update upstream Karpenter and its CRDs to v1.14.1. This might roll nodes because Karpenter v1.12 bumps the `EC2NodeClass` hash version from v4 to v5, which can mark existing nodes as drifted.
- Enable the `StaticCapacity` feature gate (`NodePool.spec.replicas`) in the `karpenter` chart defaults, as already done in `karpenter-bundle`. The feature is still alpha upstream, and we offer it as tech preview.

### network-policies [v0.3.2...v0.5.0](https://github.com/giantswarm/network-policies-app/compare/v0.3.2...v0.5.0)

#### Added

- Add `hubbleRelay` to allow the AKS ACNS `hubble-relay` through the default deny.
- Add `allowEgressToUpstreamDNS` to allow pods with `dnsPolicy: Default` to reach upstream DNS resolvers.
- Add `azureWorkloadIdentityWebhook` to allow the AKS workload identity webhook through the default deny.

#### Changed

- Move the chart to `helm/network-policies` to match its name, in preparation for devctl-generated CI.

### observability-bundle [v3.5.0...v3.6.0](https://github.com/giantswarm/observability-bundle/compare/v3.5.0...v3.6.0)

#### Added

- KSM custom resource state: Add `status_parent_accepted` and `status_parent_resolved_refs` gauges for Gateway API `HTTPRoute`, one series per parent with the condition `reason` as label.
- KSM custom resource state: Add `target_info` and `status_ancestor_accepted` metrics for Envoy Gateway `SecurityPolicy`, `BackendTrafficPolicy`, `ClientTrafficPolicy` and `EnvoyExtensionPolicy`.

#### Changed

- *!Breaking change*: Turn `apps.<appName>.dependsOn` from a list into a map from dependency name to `true`/`false`, so that a provider can turn off a single dependency (for example `kyverno-crds` on clusters without the security bundle) instead of overriding the whole list.
