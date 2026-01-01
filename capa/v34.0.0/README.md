# :zap: Giant Swarm Release v34.0.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v33.1.1

### Components

- cluster-aws from v6.4.1 to v7.2.2
- Flatcar from v4459.2.1 to [v4459.2.2](https://www.flatcar.org/releases/#release-4459.2.2)
- Kubernetes from v1.33.6 to [v1.34.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.3)
- os-tooling from v1.26.2 to v1.26.3

### cluster-aws [v6.4.1...v7.2.2](https://github.com/giantswarm/cluster-aws/compare/v6.4.1...v7.2.2)

#### :warning: Breaking Changes

- The following IAM permissions have been removed from the control plane nodes
- autoscaling:SetDesiredCapacity
- autoscaling:TerminateInstanceInAutoScalingGroup
- Removed `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers` value, as that's the default behavior now. It cannot be overridden anymore.

#### Added

- *This change will roll the control plane nodes* Add `preKubeadmCommand` to wait for the API server load balancer DNS to be resolvable before running kubeadm on control plane nodes. This prevents kubeadm from failing when the ELB DNS record hasn't propagated yet.
- *This change will roll the nodes* Add Crossplane IAM Roles, policies and instance profiles for worker and control plane nodes. Instead of having an IAM Role per node pool, now we'll use the same for all node pools.
- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- *This change will roll the nodes on Karpenter node pools* Attach the `lb` Security Group to Karpenter nodes.
- *This change will roll the nodes on Karpenter node pools* Name instance on AWS after the nodepool name.

#### Changed

- Chart: Update `cluster` to v5.1.1.
- Chart: Update `cluster` to v5.1.0.
- Chart: Update `cluster` to v5.0.0.
- Reduce redundant parts of JSON schema for Karpenter vs. MachinePool types of node pools
- Adjust node max pods based on the `nodeCidrMaskSize`

#### Fixed

- Ensure `AWSCluster.spec.network.subnets.tags` is not rendered as `null`
- Add missing documentation for node pools (health checks were not listed)
- Ensure defaulting `maxHealthyPercentage` since Helm does not use the default from the schema

### Apps

- cert-exporter from v2.9.14 to v2.9.15
- cluster-autoscaler from v1.33.1-2 to v1.34.1-1
- etcd-k8s-res-count-exporter from v1.10.11 to v1.10.12
- external-dns from v3.2.0 to v3.3.0
- k8s-audit-metrics from v0.10.10 to v0.10.11
- node-exporter from v1.20.9 to v1.20.10
- observability-bundle from v2.3.2 to v2.4.1
- Added priority-classes v0.3.0
- security-bundle from v1.15.0 to v1.16.1

### cert-exporter [v2.9.14...v2.9.15](https://github.com/giantswarm/cert-exporter/compare/v2.9.14...v2.9.15)

#### Changed

- Go: Update dependencies.

### cluster-autoscaler [v1.33.1-2...v1.34.1-1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.33.1-2...v1.34.1-1)

#### Changed

- Chart: Update to upstream v1.34.1.

### etcd-k8s-res-count-exporter [v1.10.11...v1.10.12](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.11...v1.10.12)

#### Changed

- Go: Update dependencies.

### external-dns [v3.2.0...v3.3.0](https://github.com/giantswarm/external-dns-app/compare/v3.2.0...v3.3.0)

#### Changed

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

### k8s-audit-metrics [v0.10.10...v0.10.11](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.10...v0.10.11)

#### Changed

- Go: Update dependencies.

### node-exporter [v1.20.9...v1.20.10](https://github.com/giantswarm/node-exporter-app/compare/v1.20.9...v1.20.10)

#### Removed

- Repository: Remove integration tests.

### observability-bundle [v2.3.2...v2.4.1](https://github.com/giantswarm/observability-bundle/compare/v2.3.2...v2.4.1)

#### Changed

- Upgrade `kube-prometheus-stack-app` to 19.0.0
- Update alloy-app to 0.16.0
  - Bumps alloy to 1.12.0

#### Fixed

- Fixed KSM metrics for endpoints

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).

### security-bundle [v1.15.0...v1.16.1](https://github.com/giantswarm/security-bundle/compare/v1.15.0...v1.16.1)

#### Changed

- Add missing dependency to all apps.
- Allow to set multiple dependencies on the depends-on annotation.
- Rename `edgedb` to `gel`.
- Update `cloudnative-pg` (app) to v0.0.12.
- Update `gel` (app) to v1.0.1.
