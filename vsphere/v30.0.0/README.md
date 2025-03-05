# :zap: Giant Swarm Release v30.0.0 for vSphere :zap:

## Changes compared to v29.3.0

### Components

- cluster-vsphere from v0.68.1 to v0.69.0
- Flatcar from v4081.2.1 to [v4152.2.1](https://www.flatcar.org/releases#release-4152.2.1)
- Kubernetes from v1.29.13 to [v1.30.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md)

### cluster-vsphere [v0.68.1...v0.69.0](https://github.com/giantswarm/cluster-vsphere/compare/v0.68.1...v0.69.0)

#### Changed

- Remove requirement for `pods` and `services` as they are defaulted by the values schema.
- Chart: Update `cluster` to v2.1.1.
- Chart: Enable `coredns-extensions` and `etcd-defrag`.

### Apps

- capi-node-labeler from v0.5.0 to v1.0.1
- cert-exporter from v2.9.3 to v2.9.4
- cert-manager from v3.8.2 to v3.9.0
- cilium from v0.25.2 to v0.31.0
- coredns from v1.23.0 to v1.24.0
- coredns-extensions v0.1.2
- etcd-defrag v1.0.1
- etcd-k8s-res-count-exporter from v1.10.0 to v1.10.1
- external-dns from v3.1.0 to v3.2.0
- k8s-audit-metrics from v0.10.0 to v0.10.1
- metrics-server from v2.4.2 to v2.6.0
- node-exporter from v1.20.0 to v1.20.1
- vertical-pod-autoscaler from v5.3.1 to v5.4.0
- vertical-pod-autoscaler-crd from v3.1.2 to v3.2.0

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

### etcd-defrag [v1.0.1](https://github.com/giantswarm/etcd-defrag-app/releases/v1.0.1)

#### Added

- Chart: Add `moveLeader`. ([#11](https://github.com/giantswarm/etcd-defrag-app/pull/11))

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.24.0. ([#16](https://github.com/giantswarm/etcd-defrag-app/pull/16))
- Values: Rename `cluster` into `useClusterEndpoints`. ([#8](https://github.com/giantswarm/etcd-defrag-app/pull/8))

### etcd-k8s-res-count-exporter [v1.10.0...v1.10.1](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.0...v1.10.1)

#### Changed

- Set `readOnlyRootFilesystem` to true in the container security context.
- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod` and `.nancy-ignore`. ([#242](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/pull/242))

### external-dns [v3.1.0...v3.2.0](https://github.com/giantswarm/external-dns-app/compare/v3.1.0...v3.2.0)

#### Changed

- Update architect-orb and ATS.
- Add DNSEndpoints as a source for DNS records.

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
