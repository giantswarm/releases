# :zap: Giant Swarm Release v32.2.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.1.0

### Components

- cluster-aws from v5.3.0 to v6.4.1
- Flatcar from v4230.2.4 to [v4459.2.2](https://www.flatcar.org/releases/#release-4459.2.2)
- Kubernetes from v1.32.9 to [v1.32.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.11)
- os-tooling from v1.26.2 to v1.26.3

### cluster-aws [v5.3.0...v6.4.1](https://github.com/giantswarm/cluster-aws/compare/v5.3.0...v6.4.1)

#### Added

- *This change will roll the nodes on Karpenter node pools* Attach the `lb` Security Group to Karpenter nodes.
- *This change will roll the nodes on Karpenter node pools* Name instance on AWS after the nodepool name.
- Add node-problem-detector-app, disabled by default.
- Add `capa-karpenter-taint-remover` to handle CAPA - Karpenter taint race condition.
- Add standard tags to IRSA infrastructure.
- Expose value to configure `terminationGracePeriod` in the karpenter node pools.

#### Changed

- Tidy up dependencies on `azs-getter`.
- Make `global.baseDomain` and `global.managementCluster` required values. These values will be passed to the chart when deploying it from the `cluster-app-installation-values` ConfigMap in the default namespace.
- Extract required values to its own central file to avoid repeating the `required` keyword and error messages. This is normally done automatically by a Kyverno policy.
- Change the default root disk size for Karpenter node pools. Karpenter will choose the cheapest instances, and certain instances, like `g6f.xlarge` come with some drivers that require a larger disk.
- Chart: Update `cluster` to v4.3.0.
- Change default consolidation time to 6 hours to avoid constant node rolling.
- Rename `capa-karpenter-taint-remover` app.
- Set `terminationGracePeriod` default to 30m, to avoid having `karpenter` nodes stuck in `Deleting` state due to `Pods` blocking the deletion i.e. PDBs.
- Chart: Update `cluster` to v4.2.0.
- The container registry passed as value to default apps is set to `gsoci.azurecr.io`, regardless of the cluster region. The mirroring feature of `containerd` will make sure the right registry is used.
- Switch to HelmReleases to install `karpenter` and `karpenter-crossplane-resources` charts.
- Bump flux `HelmReleases` api version to v2.
- Reduce heartbeat timeout for ASG lifecycle hooks to from 30 minutes to 3 minutes since aws-node-termination-handler-app (NTH) can now send heartbeats
- Configure the following `startupTaints` to help `karpenter` ignore pending `Pods` due to these taints that will be removed after the node starts, avoiding unnecessary instance provisioning:
  - `node.cluster.x-k8s.io/uninitialized:NoSchedule`
  - `node.cilium.io/agent-not-ready:NoSchedule`
  - `ebs.csi.aws.com/agent-not-ready:NoExecute`
- Include `cilium` ENI mode pod CIDRs in the NodePort Services security group ingress rules.

#### Removed

- Removed `capi-node-labeler` app. From now on, the worker nodes won't have the `node-role.kubernetes.io/worker` or `node.kubernetes.io/worker` labels.

### Apps

- aws-pod-identity-webhook from v2.0.0 to v2.1.0
- capi-node-labeler from v1.1.5 to v1.1.6
- cert-exporter from v2.9.13 to v2.9.15
- cilium from v1.3.1 to v1.3.3
- coredns from v1.28.2 to v1.29.0
- etcd-defrag from v1.2.2 to v1.2.3
- etcd-k8s-res-count-exporter from v1.10.10 to v1.10.12
- external-dns from v3.2.0 to v3.4.0
- k8s-audit-metrics from v0.10.9 to v0.10.11
- network-policies from v0.1.1 to v0.1.3
- node-exporter from v1.20.8 to v1.20.10
- observability-bundle from v2.3.2 to v2.5.0
- observability-policies from v0.0.2 to v0.0.3
- security-bundle from v1.14.0 to v1.16.1
- teleport-kube-agent from v0.10.6 to v0.10.7

### aws-pod-identity-webhook [v2.0.0...v2.1.0](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v2.0.0...v2.1.0)

#### Changed

- Set VPA `minAllowed` CPU to 50m. Otherwise VPA will set the CPU to tiny values that will cause CPU throttling.

### capi-node-labeler [v1.1.5...v1.1.6](https://github.com/giantswarm/capi-node-labeler-app/compare/v1.1.5...v1.1.6)

#### Changed

- Go: Update dependencies.

### cert-exporter [v2.9.13...v2.9.15](https://github.com/giantswarm/cert-exporter/compare/v2.9.13...v2.9.15)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### cilium [v1.3.1...v1.3.3](https://github.com/giantswarm/cilium-app/compare/v1.3.1...v1.3.3)

#### Changed

- Upgrade Cilium to [v1.18.5](https://github.com/cilium/cilium/releases/tag/v1.18.5).
- Upgrade Cilium to [v1.18.4](https://github.com/cilium/cilium/releases/tag/v1.18.4).

### coredns [v1.28.2...v1.29.0](https://github.com/giantswarm/coredns-app/compare/v1.28.2...v1.29.0)

#### Changed

- Update `coredns` image to [1.14.0](https://github.com/coredns/coredns/releases/tag/v1.14.0).
- Update `coredns` image to [1.13.2](https://github.com/coredns/coredns/releases/tag/v1.13.2).

### etcd-defrag [v1.2.2...v1.2.3](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.2...v1.2.3)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.36.0. ([#69](https://github.com/giantswarm/etcd-defrag-app/pull/69))

### etcd-k8s-res-count-exporter [v1.10.10...v1.10.12](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.10...v1.10.12)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### external-dns [v3.2.0...v3.4.0](https://github.com/giantswarm/external-dns-app/compare/v3.2.0...v3.4.0)

#### Changed

- Sync to upstream helm chart [1.20.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.20.0).
  - Add option to set annotationPrefix.
  - Fixed the missing schema for .provider.webhook.serviceMonitor configs.
  - Fixed incorrect indentation of selector labels under spec.template.spec.topologySpreadConstraints when topologySpreadConstraints is set.
- Use kubectl-apply-job when installing CRDs.
- Upgrade external-dns to v0.20.0.
- Update DNSEndpoints CRD.
- Sync to upstream helm chart `1.19.0`.
  - Grant `discovery.k8s.io/endpointslices` permission only when using `service` source.
  - Update RBAC for `Service` source to support `EndpointSlices`.
  - Allow extraArgs to also be a map enabling overrides of individual values.
  - Set defaults for `automountServiceAccountToken` and `serviceAccount.automountServiceAccountToken` to `true` in Helm chart values.
  - Correctly handle `txtPrefix` and `txtSuffix` arguments when both are provided.
  - Add ability to generate schema with `helm plugin schema`.
  - Regenerate JSON schema with `helm-values-schema-json' plugin.
  - Added ability to configure `imagePullSecrets` via helm `global` value.
  - Added options to configure `labelFilter` and `managedRecordTypes` via dedicated helm values.
  - Allow templating `serviceaccount.annotations` keys and values, by rendering them using the `tpl` built-in function.
  - Added support for `extraContainers` argument.
  - Added support for setting `excludeDomains` argument.
  - Added support for setting `dnsConfig`.
  - Added support for webhook providers.
- Restrict managed record types to A and CNAME.

### k8s-audit-metrics [v0.10.9...v0.10.11](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.9...v0.10.11)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.

### network-policies [v0.1.1...v0.1.3](https://github.com/giantswarm/network-policies-app/compare/v0.1.1...v0.1.3)

#### Added

- Add support for Kamaji.

#### Fixed

- Fixed broken templating.

### node-exporter [v1.20.8...v1.20.10](https://github.com/giantswarm/node-exporter-app/compare/v1.20.8...v1.20.10)

#### Changed

- Go: Update dependencies.

#### Removed

- Repository: Remove integration tests.

### observability-bundle [v2.3.2...v2.5.0](https://github.com/giantswarm/observability-bundle/compare/v2.3.2...v2.5.0)

#### Added

- Add KSM metrics `kube_servicemonitor_info` and `kube_podmonitor_info` for ServiceMonitor and PodMonitor resources
- Add KSM metrics `kube_podlog_info` for PodLog resource

#### Changed

- Upgrade `kube-prometheus-stack-app` to 19.0.0
- Update alloy-app to 0.16.0
  - Bumps alloy to 1.12.0

#### Fixed

- Fixed KSM metrics for endpoints

### observability-policies [v0.0.2...v0.0.3](https://github.com/giantswarm/observability-policies-app/compare/v0.0.2...v0.0.3)

#### Fixed

- Missing RBAC for kyverno-report-controller

### security-bundle [v1.14.0...v1.16.1](https://github.com/giantswarm/security-bundle/compare/v1.14.0...v1.16.1)

#### Added

- Add `kubescape` (app) version v0.0.4.

#### Changed

- Add missing dependency to all apps.
- Allow to set multiple dependencies on the depends-on annotation.
- Rename `edgedb` to `gel`.
- Update `cloudnative-pg` (app) to v0.0.12.
- Update `gel` (app) to v1.0.1.
- Update `kyverno` (app) to v0.21.1.
- Update `kyverno-crds` (app) to v1.15.0.

### teleport-kube-agent [v0.10.6...v0.10.7](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.6...v0.10.7)

#### Added

- Add `ephemeral-storage` requests and limits to satisfy Kyverno policy `require-emptydir-requests-and-limits`.

#### Changed

- Enable upstream-provided Prometheus PodMonitor to scrape metrics from Teleport Kube Agent pods.
