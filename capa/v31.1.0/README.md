# :zap: Giant Swarm Release v31.1.0 for CAPA :zap:

This release updates Kubernetes to the latest patch release v1.31.11.

During control plane upgrades, short-term warnings are now prevented by setting a fixed instead of dynamic AMI lookup string. This leads to nodes being rolled once when upgrading to this release.

We added an option to set the IMDSv2 request hop limit for EC2 instances ‒ this is usually not needed, except if security requirements such as AWS SCPs (service control policies) dictate a maximum.

Karpenter support keeps getting better: node-termination-handler is not installed anymore if only Karpenter node pools are used, as the same function is built into Karpenter (pod draining and EC2 instance termination handling). Nodes in such pools now also use the reduced IAM permission set (can be toggled in exceptional cases).

## Changes compared to v31.0.0

### Components

- cluster-aws from v3.4.0 to v3.6.0
- Kubernetes from v1.31.9 to [v1.31.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.11)

### cluster-aws [v3.4.0...v3.6.0](https://github.com/giantswarm/cluster-aws/compare/v3.4.0...v3.6.0)

#### Added

- Add `giantswarm.io/role: nodes` by default to private subnets used for nodes. Can be overwritten.
- Make IMDSv2 hop limit configurable

#### Changed

- Chart: Update `cluster` to v2.5.0.
- Only deploy `node-termination-handler` when there are non-karpenter node pools because karpenter takes care of node draining
- Change `imageLookupFormat` to use a static string rather than CAPI replacing the OS and Kubernetes versions. This rolls control plane nodes.

#### Fixed

- Use reduced IAM permissions on `karpenter` worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.

### Apps

- aws-nth-bundle from v1.2.1 to v1.2.2
- capi-node-labeler from v1.1.1 to v1.1.2
- cert-exporter from v2.9.7 to v2.9.8
- cilium from v1.2.1 to v1.2.2
- cluster-autoscaler from v1.31.2-gs2 to v1.31.3-gs1
- coredns from v1.25.0 to v1.26.0
- etcd-defrag from v1.0.5 to v1.0.6
- etcd-k8s-res-count-exporter from v1.10.5 to v1.10.6
- k8s-audit-metrics from v0.10.4 to v0.10.5
- k8s-dns-node-cache from v2.8.1 to v2.9.0
- karpenter-bundle from v2.0.0 to v2.1.0
- karpenter-nodepools from v0.1.0 to v0.2.0
- node-exporter from v1.20.3 to v1.20.4
- security-bundle from v1.11.0 to v1.12.0
- teleport-kube-agent from v0.10.5 to v0.10.6


### aws-nth-bundle [v1.2.1...v1.2.2](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.1...v1.2.2)

#### Changed

- Upgrade Node Termination Handler to 1.21.0.

### capi-node-labeler [v1.1.1...v1.1.2](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.1.1...v1.1.2)

#### Changed

- Go: Update dependencies.

### cert-exporter [v2.9.7...v2.9.8](https://github.com/giantswarm/cert-exporter/compare/v2.9.7...v2.9.8)

#### Changed

- Go: Update dependencies.

### cilium [v1.2.1...v1.2.2](https://github.com/giantswarm/cilium-app/compare/v1.2.1...v1.2.2)

#### Changed

- Upgrade Cilium to [v1.17.6](https://github.com/cilium/cilium/releases/tag/v1.17.6).
- Updated E2E tests to use apptest-framework v1.14.0
- Increase Cilium operator resource limits.

#### Removed

- Remove deprecated "partial" mode from Kube Proxy Replacement options.

### cluster-autoscaler [v1.31.2-gs2...v1.31.3-gs1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.31.2-gs2...v1.31.3-gs1)

#### Changed

- Chart: Update to upstream v1.31.3.

### coredns [v1.25.0...v1.26.0](https://github.com/giantswarm/coredns-app/compare/v1.25.0...v1.26.0)

#### Changed

- Update `coredns` image to [1.12.2](https://github.com/coredns/coredns/releases/tag/v1.12.2).

### etcd-defrag [v1.0.5...v1.0.6](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.5...v1.0.6)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.29.0. ([#43](https://github.com/giantswarm/etcd-defrag-app/pull/43))

### etcd-k8s-res-count-exporter [v1.10.5...v1.10.6](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.5...v1.10.6)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.4...v0.10.5](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.4...v0.10.5)

#### Changed

- Go: Update dependencies.

### k8s-dns-node-cache [v2.8.1...v2.9.0](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.8.1...v2.9.0)

#### Changed

- Upgrade application to version 1.26.4 (includes coredns 1.11.3)
- Increase ServiceMonitor's scrapping interval to 1m.
- Remove obsolete PSPs

### karpenter-bundle [v2.0.0...v2.1.0](https://github.com/giantswarm/karpenter-bundle/compare/v2.0.0...v2.1.0)

#### Removed

- Remove capa-karpenter-taint-remover because nodes are now in the `MachinePool` CR, so the taint will be removed by CAPI.

### karpenter-nodepools [v0.1.0...v0.2.0](https://github.com/giantswarm/karpenter-nodepools/compare/v0.1.0...v0.2.0)

#### Changed

- Improve json schema.
- Change subnet selector to avoid CNI subnets.

### node-exporter [v1.20.3...v1.20.4](https://github.com/giantswarm/node-exporter-app/compare/v1.20.3...v1.20.4)

#### Changed

- Go: Update to v1.24.5.

### security-bundle [v1.11.0...v1.12.0](https://github.com/giantswarm/security-bundle/compare/v1.11.0...v1.12.0)

#### Changed

- Update `trivy-operator` (app) to v0.11.1.
- Update `trivy` (app) to v0.14.0.
- Update `falco` (app) to v0.10.1.
- Update `cloudnative-pg` (app) to v0.0.10.
- Update `starboard-exporter` (app) to v0.8.2.
- Updated E2E tests to use apptest-framework v1.14.0

### teleport-kube-agent [v0.10.5...v0.10.6](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.5...v0.10.6)

#### Changed

- AppVersion upgrade to 17.5.4
