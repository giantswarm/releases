# :zap: Giant Swarm Release v31.0.0 for CAPA :zap:

This release along with k8s and application upgrades also brings several new features for the product. Node Pools have been extended with [new Karpenter type](https://github.com/giantswarm/cluster-aws/blob/main/helm/cluster-aws/values.schema.json#L151), integrating the solution fully with the Giant Swarm cluster lifecycle instead of a Managed Application. Karpenter application will now be deployed as a part of the Giant Swarm clusters out of the box if configured. For further configuration please check [our example](https://github.com/giantswarm/cluster-aws/blob/main/helm/cluster-aws/ci/test-karpenter-values.yaml) of the Karpenter Node Pool usage.
Additionally, we have extended the Cluster configuration to support multiple VPC CIDRs under `global.connectivity.network.vpcCidr`, please read the [schema documentation](https://github.com/giantswarm/cluster-aws/blob/main/helm/cluster-aws/README.md) for more details. 
Finally we are slowly introducing changes to `IAM roles for service accounts` (IRSA) management on GS side, where the infrastructure required will be fully managed by Crossplane instead of `irsa-operator` and `capa-iam-operator`. There is no impact for customers, but the change will allow Giant Swarm to pair the IAM permissions for required applications with their actual releases and deployments, moving away from single operators implementing all the roles. The Karpenter application will be the first one to use it.

For any questions regarding new features or their usage, please reach out to Giant Swarm. For customers running Karpenter as a Managed Application from Giant Swarm catalog, it is save to upgrade to this release without any changes. Application will work as expected until migrated to the new Node Pool type.

## Changes compared to v30.1.3

### Components

- cluster-aws from v3.2.2 to v3.4.0
- Flatcar from v4152.2.1 to [v4152.2.3](https://www.flatcar-linux.org/releases/#release-4152.2.3)
- Kubernetes from v1.30.11 to [v1.31.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.9)
- os-tooling from v1.24.0 to v1.26.1

### cluster-aws [v3.2.2...v3.4.0](https://github.com/giantswarm/cluster-aws/compare/v3.2.2...v3.4.0)

#### Added

- Adopt IRSA infrastructure with Crossplane. It can be disabled to use IRSA Operator.
- Support multiple VPC CIDRs
- Add `karpenter` support
  - Expose new values to configure karpenter node pools.
  - Deploy `karpenter` app when `karpenter` node pools are configured.
- Add `cert-manager-crossplane-resources` App in private clusters so `DNS01` `clusterIssuer`.
- Add configuration for `DNS01` `clusterIssuer` deployed by `cert-manager-app` in private clusters.
- Apply startup taint `ebs.csi.aws.com/agent-not-ready` for AWS EBS CSI driver on worker nodes.

#### Changed

- Reduce IMDS Response Hop Limit to 2 if pod networking is in ENI mode to increase security.
- Configure HelmReleases to retry indefinitely when installation or upgrade fails by setting retries: -1.
- Chart: Update `cluster` to v2.4.0.

### Apps

- Added cert-manager-crossplane-resources v0.1.0
- Added karpenter-bundle v2.0.0
- Added karpenter-nodepools v0.1.0
- capi-node-labeler from v1.0.2 to v1.1.1
- cert-exporter from v2.9.5 to v2.9.7
- cert-manager from v3.9.0 to v3.9.1
- cilium from v0.31.5 to v1.2.1
- cilium-crossplane-resources from v0.2.0 to v0.2.1
- cloud-provider-aws from v1.30.8-gs1 to v1.31.5-gs1
- cluster-autoscaler from v1.30.4-gs1 to v1.31.2-gs2
- coredns from v1.24.0 to v1.25.0
- etcd-defrag from v1.0.2 to v1.0.5
- etcd-k8s-res-count-exporter from v1.10.3 to v1.10.5
- k8s-audit-metrics from v0.10.2 to v0.10.4
- net-exporter from v1.22.0 to v1.23.0
- node-exporter from v1.20.2 to v1.20.3
- observability-bundle from v1.11.0 to v2.0.0
- observability-policies from v0.0.1 to v0.0.2
- security-bundle from v1.10.1 to v1.11.0
- teleport-kube-agent from v0.10.4 to v0.10.5
- vertical-pod-autoscaler from v5.4.0 to v5.5.1
- vertical-pod-autoscaler-crd from v3.2.0 to v3.3.1


### capi-node-labeler [v1.0.2...v1.1.1](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.0.2...v1.1.1)

#### Changed

- Go: Update dependencies.
- Improve Control Plane node detection.
- Taint Control Plane nodes if not already tainted.
- Go: Update dependencies.

### cert-exporter [v2.9.5...v2.9.7](https://github.com/giantswarm/cert-exporter/compare/v2.9.5...v2.9.7)

#### Changed

- Go: Update dependencies.
- Fix linting issues.
- Go: Update dependencies.

### cert-manager [v3.9.0...v3.9.1](https://github.com/giantswarm/cert-manager-app/compare/v3.9.0...v3.9.1)

#### Added

- Added Vertical Pod Autoscaler support for `controller` pods.
- Added renovate configutarion

#### Removed

- Removed dependabot configuration

### cert-manager-crossplane-resources [v0.1.0](https://github.com/giantswarm/cert-manager-crossplane-resources/releases/tag/v0.1.0)

#### Added

- Added support for `Azure`
- Included the `giantswarm.io/cluster` label

#### Changed

- Restructured Chart to support multiple cloud providers

### cilium [v0.31.5...v1.2.1](https://github.com/giantswarm/cilium-app/compare/v0.31.5...v1.2.1)

#### Changed

- Enable conntrack accounting in Cilium agent by default.
- Re-enable Cilium agent and operator metrics port.
- Add resource requests and limits to Hubble UI and Relay.
- Add resource requests and limits to Cilium Operator.
- Upgrade Cilium to [v1.17.4](https://github.com/cilium/cilium/releases/tag/v1.17.4).
- Cilium v1.17.4 disables kubernetes api connectivity check for liveness probes. (Upstream PR: https://github.com/cilium/cilium/pull/38703)
- Upgrade Cilium to [v1.17.3](https://github.com/cilium/cilium/releases/tag/v1.17.3).
- Upgrade Cilium to [v1.17.2](https://github.com/cilium/cilium/releases/tag/v1.17.2).
- Remove cleanup kube-proxy patch.
- Identity computation label exclusion list regular expressions. Remove `controller-uid`, since this is excluded by default now.
- Upgrade Cilium to [v1.17.0](https://github.com/cilium/cilium/releases/tag/v1.17.0).
- Use upstream default value for `prometheus.metrics`.
- Enable Envoy Proxy in standalone DaemonSet.

### cilium-crossplane-resources [v0.2.0...v0.2.1](https://github.com/giantswarm/cilium-crossplane-resources/compare/v0.2.0...v0.2.1)

#### Added

- Included the `giantswarm.io/cluster` label

### cloud-provider-aws [v1.30.8-gs1...v1.31.5-gs1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.30.8-gs1...v1.31.5-gs1)

#### Changed

- Chart: Update to upstream v1.31.5.

### cluster-autoscaler [v1.30.4-gs1...v1.31.2-gs2](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.30.4-gs1...v1.31.2-gs2)

#### Added

- Add additional labels to ignore during ASG balancing check
- Support adding additional labels to the `PodMonitor` resource via the `podMonitor.additionalLabels` value.

#### Changed

- Chart: Use v1.31.2.
- Chart: Update to upstream v1.31.2. ([#325](https://github.com/giantswarm/cluster-autoscaler-app/pull/325))

### coredns [v1.24.0...v1.25.0](https://github.com/giantswarm/coredns-app/compare/v1.24.0...v1.25.0)

#### Changed

- Update `coredns` image to [1.12.1](https://github.com/coredns/coredns/releases/tag/v1.12.1).

### etcd-defrag [v1.0.2...v1.0.5](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.2...v1.0.5)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.28.0. ([#34](https://github.com/giantswarm/etcd-defrag-app/pull/34))
- Chart: Update dependency ahrtr/etcd-defrag to v0.27.0. ([#29](https://github.com/giantswarm/etcd-defrag-app/pull/29))
- Chart: Update dependency ahrtr/etcd-defrag to v0.26.0. ([#22](https://github.com/giantswarm/etcd-defrag-app/pull/22))

### etcd-k8s-res-count-exporter [v1.10.3...v1.10.5](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.3...v1.10.5)

#### Changed

- Go: Update dependencies.

#### Fixed

- Fix linting issues.
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.2...v0.10.4](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.2...v0.10.4)

#### Changed

- Go: Update dependencies.

#### Fixed

- Fix linting issues.
- Go: Update dependencies.

### karpenter-bundle [v2.0.0](https://github.com/giantswarm/karpenter-bundle/releases/tag/v2.0.0)

#### Added

- First release

#### Changed

- Add karpenter-app dependency on karpenter-crossplane-resources app.
- Bump karpenter to `v1.5.0`.
- Bump karpenter-app to `v0.14.0`.
- Update `karpenter-capa-taint-remover` to allow scheduling on all taints.
- Update `karpenter-crossplane-resources` app version to add support for vintage OIDC issuer on migrated clusters
- Update karpenter to update flowschema API
- Update interruption queue settings
- Update SQS Policy URL

### karpenter-nodepools [v0.1.0](https://github.com/giantswarm/karpenter-nodepools/releases/tag/v0.1.0)

#### Changed

- changed: `app.giantswarm.io` label group was changed to `application.giantswarm.io`

### net-exporter [v1.22.0...v1.23.0](https://github.com/giantswarm/net-exporter/compare/v1.22.0...v1.23.0)

#### Changed

- Check for errors when closing connections.
- Switch from Endpoints to EndpointSlices for neighbors discovery.

### node-exporter [v1.20.2...v1.20.3](https://github.com/giantswarm/node-exporter-app/compare/v1.20.2...v1.20.3)

#### Changed

- Go: Update dependencies.

### observability-bundle [v1.11.0...v2.0.0](https://github.com/giantswarm/observability-bundle/compare/v1.11.0...v2.0.0)

#### Added

- Add support for enabling pre-configured custom resources in KSM
- Add metrics containing labels for Crossplane resources

#### Changed

- Upgrade `alloy-app` from 0.10.0 to 0.11.0
  - This bumps the version of `Alloy` from 1.8.3 to 1.9.0
- Upgrade `alloy-app` from 0.9.0 to 0.10.0
  - This bumps the version of `Alloy` from 1.7.1 to 1.8.3
- Reconfigure Flux-related part of the KSM to use wildcards instead of hardcoded versions.
- Rename Flux-related metrics produced by the KSM.
- Upgrade `kube-prometheus-stack` to 72.3.0
  - Bumps prometheus-operator to 0.82.0
  - Bumps prometheus-operator CRDs to 0.82.0
- Upgrade `kube-prometheus-stack` to 72.3.0
  - Bumps prometheus-operator to 0.82.0
- Upgrade `kube-prometheus-stack` from 69.5.1 to 70.1.1
  - Bumps prometheus-operator to 0.81.0
  - Bumps prometheus to 3.2.1

#### Fixed

- Fix catalog for alloy apps as it is now pushed to the default catalog.

#### Removed

- Clean up old and deprecated telemetry collectors:
  - `promtail`
  - `grafana-agent`
  - `promtheus-agent`
- Disable PodSecurityPolicies by default as PodSecurityPolicies are deprecated and removed in Kubernetes v1.25+ clusters

### observability-policies [v0.0.1...v0.0.2](https://github.com/giantswarm/observability-policies-app/compare/v0.0.1...v0.0.2)

#### Changed

- Add Cluster Role to allow latest Kyverno versions to work (https://github.com/giantswarm/giantswarm/issues/33416)
- Switch `.Values.disabled` to `.Values.enabled` to follow best practices.

### security-bundle [v1.10.1...v1.11.0](https://github.com/giantswarm/security-bundle/compare/v1.10.1...v1.11.0)

#### Added

- Add `policy-api-crds` app to manage Policy API CRDs.

#### Changed

- Update `trivy` (app) to v0.13.4.
- Update `cloudnative-pg` (app) to v0.0.7.
- Update `starboard-exporter` (app) to v0.8.1.
- Update `kyverno-policy-operator` (app) to v0.0.11.
- Update `cloudnative-pg` (app) to v0.0.9.

#### Notes

**Note:** Kyverno `PolicyExceptions` (API group `kyverno.io`) versions `v2alpha1` and `v2beta1` are deprecated and will be removed in the next Kyverno minor release (v1.14). Please update all Kyverno PolicyExceptions to `v2`. No action is required for Giant Swarm Policy API `PolicyExceptions` (API group `policy.giantswarm.io`), which are handled automatically.

### teleport-kube-agent [v0.10.4...v0.10.5](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.4...v0.10.5)

#### Added

- Set Home URL in chart metadata.

### vertical-pod-autoscaler [v5.4.0...v5.5.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.4.0...v5.5.1)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v10.2.1. ([#355](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/355))
- Chart: Update Helm release vertical-pod-autoscaler to v10.1.0. ([#350](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/350))
- Chart: Update Helm release vertical-pod-autoscaler to v10.2.0. ([#351](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/351))
- Chart: Update Helm release vertical-pod-autoscaler to v10.0.1. ([#346](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/346))

### vertical-pod-autoscaler-crd [v3.2.0...v3.3.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.2.0...v3.3.1)

#### Changed

- Chart: Sync to upstream. ([#146](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/146))
- Chart: Sync to upstream. ([#140](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/140))
- Chart: Sync to upstream. ([#136](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/136))
