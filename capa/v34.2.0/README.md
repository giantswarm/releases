# :zap: Giant Swarm Release v34.2.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v34.1.0

### Components

- cluster-aws from v7.4.0 to v8.1.0
- cluster from v5.1.2 to v6.1.0
- Flatcar from v4459.2.3 to [v4459.2.4](https://www.flatcar.org/releases/#release-4459.2.4)
- Kubernetes from v1.34.5 to [v1.34.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.6)
- os-tooling from v1.26.4 to v1.27.0

### cluster-aws [v7.4.0...v8.1.0](https://github.com/giantswarm/cluster-aws/compare/v7.4.0...v8.1.0)

#### Added

- Add `appVersion` field to `Chart.yaml`.
- Enable scraping metrics and logs from the karpenter app.
- Allow to configure the name of the hosted zone to use for the workload cluster by setting `global.connectivity.dns.hostedZoneName`.
- Allow to configure the AWS IAM Role to use when managing the DNS delegation for the hosted zone by setting `global.connectivity.dns.delegationIdentityName`.
- Added new annotation `network.giantswarm.io/base-domain` with the base domain value used for the workload cluster.
- Add support for `network.giantswarm.io/wildcard-cname-target` annotation on `AWSCluster` via `global.connectivity.dns.wildcardCnameTarget`.

#### Changed

- Enable cert-manager DNS challenges by default.
- Reduced default karpenter consolidation from 6 hours to 1 hour.
- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

#### Fixed

- Set `appName` before `catalog` lookup in `aws-nth-app` template to ensure correct catalog resolution from Release CR.

### cluster [v5.1.2...v6.1.0](https://github.com/giantswarm/cluster/compare/v5.1.2...v6.1.0)

#### Added

- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))
- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.

#### Changed

- Helpers: Use `.Chart.AppVersion` in `app.kubernetes.io/version` label.
- Cluster API: Migrate to API `v1beta2`.
- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.

#### Fixed

- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.

### Apps

- aws-ebs-csi-driver from v4.1.1 to v4.1.2
- cert-exporter from v2.9.16 to v2.10.0
- cert-manager-crossplane-resources from v0.1.0 to v0.1.1
- cloud-provider-aws from v2.0.0 to v2.1.0
- coredns from v1.29.1 to v1.29.2
- etcd-defrag from v1.2.4 to v1.2.5
- karpenter from v2.1.0 to v2.2.0
- observability-bundle from v2.6.0 to v2.8.0
- prometheus-blackbox-exporter from v0.5.1 to v0.6.0

### aws-ebs-csi-driver [v4.1.1...v4.1.2](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.1...v4.1.2)

#### Changed

- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.

### cert-exporter [v2.9.16...v2.10.0](https://github.com/giantswarm/cert-exporter/compare/v2.9.16...v2.10.0)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Values: Tune resources.

### cert-manager-crossplane-resources [v0.1.0...v0.1.1](https://github.com/giantswarm/cert-manager-crossplane-resources/compare/v0.1.0...v0.1.1)

#### Changed

- Update `architect-orb` to v6.15.0.

### cloud-provider-aws [v2.0.0...v2.1.0](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v2.0.0...v2.1.0)

#### Changed

- Bump to upstream image v1.35.0

### coredns [v1.29.1...v1.29.2](https://github.com/giantswarm/coredns-app/compare/v1.29.1...v1.29.2)

#### Changed

- Update `coredns` image to [1.14.2](https://github.com/coredns/coredns/releases/tag/v1.14.2).

### etcd-defrag [v1.2.4...v1.2.5](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.4...v1.2.5)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.38.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/84))

### karpenter [v2.1.0...v2.2.0](https://github.com/giantswarm/karpenter-app/compare/v2.1.0...v2.2.0)

#### Changed

- Update ABS config to replace `.appVersion` in Chart.yaml with version detected by ABS.

#### Fixed

- Use `.Chart.AppVersion` instead of `.Chart.Version` for OCIRepository tag.

### observability-bundle [v2.6.0...v2.8.0](https://github.com/giantswarm/observability-bundle/compare/v2.6.0...v2.8.0)

#### Added

- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.

#### Changed

- Update kube-prometheus-stack to 20.1.0
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).
- Update alloy-app to 0.17.1
- Update kube-prometheus-stack to 20.0.0
- Update prometheus-operator-crd to 20.0.0

### prometheus-blackbox-exporter [v0.5.1...v0.6.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.5.1...v0.6.0)

#### Changed

- Set `priorityClassName` to `system-node-critical` to ensure DaemonSet pods are scheduled even on full nodes.
