# :zap: Giant Swarm Release v34.0.0 for VMware Cloud Director :zap:

## Changes compared to v33.1.1

### Components

- cluster-cloud-director from v2.4.0 to v3.1.1
- Kubernetes from v1.33.6 to [v1.34.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.3)
- os-tooling from v1.26.2 to v1.26.3

### cluster-cloud-director [v2.4.0...v3.1.1](https://github.com/giantswarm/cluster-cloud-director/compare/v2.4.0...v3.1.1)

#### Added

- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- Add `"helm.sh/resource-policy": keep` annotation to `VCDCluster` CR so that it doesn't get removed by Helm when uninstalling this chart. The CAPI controllers will take care of removing it, following the expected deletion order.

#### Changed

- Chart: Update `cluster` to v5.1.1.
- Chart: Update `cluster` to v5.1.0.
- Chart: Update `cluster` to v5.0.0.

### Apps

- external-dns from v3.2.0 to v3.3.0
- Added priority-classes v0.3.0

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

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).
