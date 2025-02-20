# :zap: Giant Swarm Release v30.0.0 for CAPA :zap:

## Changes compared to v29.6.1

### Components

- cluster-aws from v2.6.1 to v3.0.0
- Flatcar from v4081.2.1 to [v4152.2.1](https://www.flatcar.org/releases#release-4152.2.1)
- Kubernetes from v1.29.13 to [v1.30.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md)

### cluster-aws [v2.6.1...v3.0.0](https://github.com/giantswarm/cluster-aws/compare/v2.6.1...v3.0.0)

#### Added

- Values: Add schema for `cilium-crossplane-resources`.

#### Changed

- Scale down cilium-operator before deleting a cluster (only in eni mode)
- Chart: Update `cluster` to v2.0.1.
- Chart: Enable `coredns-extensions` and `etcd-defrag`.

### Apps

- aws-ebs-csi-driver from v2.30.1 to v3.0.1
- aws-pod-identity-webhook from v1.18.0 to v1.19.0
- capi-node-labeler from v0.5.0 to v1.0.1
- cert-exporter from v2.9.3 to v2.9.4
- cert-manager from v3.8.2 to v3.9.0
- cilium from v0.25.2 to v0.31.0
- cloud-provider-aws from v1.29.3-gs1 to v1.30.7-gs2
- cluster-autoscaler from v1.29.3-gs1 to v1.30.3-gs2
- coredns from v1.23.0 to v1.24.0
- coredns-extensions v0.1.2
- etcd-defrag v1.0.0
- etcd-k8s-res-count-exporter from v1.10.0 to v1.10.1
- k8s-audit-metrics from v0.10.0 to v0.10.1
- metrics-server from v2.4.2 to v2.6.0
- node-exporter from v1.20.0 to v1.20.1
- vertical-pod-autoscaler from v5.3.1 to v5.4.0
- vertical-pod-autoscaler-crd from v3.1.2 to v3.2.0

### aws-ebs-csi-driver [v2.30.1...v3.0.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v2.30.1...v3.0.1)

#### Added

- Repository: Some chores. ([#235](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/235))
  - Repository: Add `Makefile.custom.mk`.
- Chart: Add `snapshot-controller` NetworkPolicy. ([#246](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/246))
  - Kustomization: Add `snapshot-controller` NetworkPolicy.

#### Changed

- Harden security context for controller and node.
- Repository: Some chores. ([#235](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/235))
  - ABS: Rework `main.yaml`.
  - CircleCI: Rework `config.yml`.
  - Repository: Rework `README.md`.
  - Repository: Move `.gitignore` & `kustomization-snapshotter.yaml` to `vendor/external-snapshotter/`.
  - Chart: Rework `.kube-linter.yaml`.
  - Vendir: Rework `vendir.yml`.
  - Chart: Rework `Chart.yaml`.
  - Chart: Revert image to v1.37.0.
  - Renovate: Ignore `values.yaml`.
- Chart: Sync to upstream. ([#243](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/243))
  - Chart: Reorder labels.
  - Chart: Fix network policies.
- Chart: Add `snapshot-controller` NetworkPolicy. ([#246](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/246))
  - Vendir: Sync to `vendor/external-snapshotter/upstream`.
  - Kustomization: Set namespace.
  - Kustomization: Extend common labels.
  - Kustomization: Extract CRD patches.
  - Kustomization: Extract service account patches.
  - Kustomization: Extract deployment patches.
- Change to use ImagePullPolicy as specified via values.
- Upgrade to release v1.37.0
- Enable Volume Snapshotter by default
- Switch to Helm managed CRDs

#### Removed

- Repository: Some chores. ([#235](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/235))
  - Repository: Remove `.nancy-ignore`.
  - Chart: Remove pod `securityContext` from `external-snapshotter`.
  - Chart: Remove `.helmignore`.
  - Chart: Remove `CHANGELOG.md`.

### aws-pod-identity-webhook [v1.18.0...v1.19.0](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.18.0...v1.19.0)

#### Changed

- Add support for rolling `Deployments` owned by unknown CRs, like the case of Crossplane providers.

### capi-node-labeler [v0.5.0...v1.0.1](https://github.com/giantswarm/capi-node-labeler-app/compare/v0.5.0...v1.0.1)

#### Changed

- Main: Improve sleep. ([#125](https://github.com/giantswarm/capi-node-labeler-app/pull/125))
- Go: Update `go.mod` and `.nancy-ignore`. ([#123](https://github.com/giantswarm/capi-node-labeler-app/pull/123))

### cert-exporter [v2.9.3...v2.9.4](https://github.com/giantswarm/cert-exporter/compare/v2.9.3...v2.9.4)

#### Changed

- Repository: Some chores. ([#418](https://github.com/giantswarm/cert-exporter/pull/418))
- Go: Update `go.mod` and `.nancy-ignore`. ([#437](https://github.com/giantswarm/cert-exporter/pull/437))

### cert-manager [v3.8.2...v3.9.0](https://github.com/giantswarm/cert-manager-app/compare/v3.8.2...v3.9.0)

#### Added

- Adds new sync method based on Vendir to sync from upstream

#### Changed

- Updates Cert-manager Chart to Upstream 1.16.2

### cilium [v0.25.2...v0.31.0](https://github.com/giantswarm/cilium-app/compare/v0.25.2...v0.31.0)

#### Changed

- Upgrade Cilium to [v1.16.6](https://github.com/cilium/cilium/releases/tag/v1.16.6).
- Move provider specific custom CNI configuration to subchart.
- Improve security defaults for:
  - Hubble UI
  - Hubble Relay
  - Cilium Operator

#### Removed

- Delete defaultPolicies and extraPolicies templates.

### cloud-provider-aws [v1.29.3-gs1...v1.30.7-gs2](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.29.3-gs1...v1.30.7-gs2)

#### Added

- Add security context to the container for PSS.

#### Changed

- Chart: Remove duplicate service account. ([#87](https://github.com/giantswarm/aws-cloud-controller-manager-app/pull/87))
- Chart: Update to upstream v1.30.7.

#### Remove

- Remove PSP manifest.

### cluster-autoscaler [v1.29.3-gs1...v1.30.3-gs2](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.29.3-gs1...v1.30.3-gs2)

#### Changed

- Values: Enable Pod Security Standards. ([#296](https://github.com/giantswarm/cluster-autoscaler-app/pull/296))
- Chart: Update to upstream v1.30.3. ([#298](https://github.com/giantswarm/cluster-autoscaler-app/pull/298))
- Update `PolicyExceptions` apiVersion to `v2beta1`. ([#282](https://github.com/giantswarm/cluster-autoscaler-app/pull/282))

### coredns [v1.23.0...v1.24.0](https://github.com/giantswarm/coredns-app/compare/v1.23.0...v1.24.0)

#### Changed

- Update `coredns` image to [1.12.0](https://github.com/coredns/coredns/releases/tag/v1.12.0).
- Disable HPA Memory target.
- Increase threshold for HPA CPU target to 80%.

### coredns-extensions [v0.1.2](https://github.com/giantswarm/coredns-extensions-app/releases/v0.1.2)

#### Added

- Add VPA for CoreDNS deployments.
- Add value to enable or disable VPA resources.

#### Changed

- Push App to the default-catalog.
- Publish App in giantswarm-catalog.

### etcd-defrag [v1.0.0](https://github.com/giantswarm/etcd-defrag-app/releases/v1.0.0)

#### Added

- Chart: Add `moveLeader`. ([#11](https://github.com/giantswarm/etcd-defrag-app/pull/11))

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.23.0. ([#10](https://github.com/giantswarm/etcd-defrag-app/pull/10))
- Values: Rename `cluster` into `useClusterEndpoints`. ([#8](https://github.com/giantswarm/etcd-defrag-app/pull/8))

### etcd-k8s-res-count-exporter [v1.10.0...v1.10.1](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.0...v1.10.1)

#### Changed

- Set `readOnlyRootFilesystem` to true in the container security context.
- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod` and `.nancy-ignore`. ([#242](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/pull/242))

### k8s-audit-metrics [v0.10.0...v0.10.1](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.0...v0.10.1)

#### Changed

- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod` and `.nancy-ignore`. ([#248](https://github.com/giantswarm/k8s-audit-metrics/pull/248))

### metrics-server [v2.4.2...v2.6.0](https://github.com/giantswarm/metrics-server-app/compare/v2.4.2...v2.6.0)

#### Added

- Add VPA setting for `metrics-server`.

#### Changed

- Upgrade metrics-server to v0.7.2.
- Chart: Update PolicyExceptions to v2beta1. ([#226](https://github.com/giantswarm/metrics-server-app/pull/226))

### node-exporter [v1.20.0...v1.20.1](https://github.com/giantswarm/node-exporter-app/compare/v1.20.0...v1.20.1)

#### Changed

- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod`. ([#322](https://github.com/giantswarm/node-exporter-app/pull/322))

### vertical-pod-autoscaler [v5.3.1...v5.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.3.1...v5.4.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v10.0.0 ([#335](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/335))

### vertical-pod-autoscaler-crd [v3.1.2...v3.2.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.1.2...v3.2.0)

#### Changed

- Chart: Sync to upstream. ([#126](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/126))
