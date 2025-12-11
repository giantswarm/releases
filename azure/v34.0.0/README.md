# :zap: Giant Swarm Release v34.0.0 for Azure :zap:

## Changes compared to v33.1.0

### Components

- cluster-azure from v4.4.0 to v5.1.0
- Kubernetes from v1.33.6 to [v1.34.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.3)
- os-tooling from v1.26.2 to v1.26.3

### cluster-azure [v4.4.0...v5.1.0](https://github.com/giantswarm/cluster-azure/compare/v4.4.0...v5.1.0)

#### Added

- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- Add `"helm.sh/resource-policy": keep` annotation to `AzureCluster` CR so that it doesn't get removed by Helm when uninstalling this chart. The CAPI controllers will take care of removing it, following the expected deletion order.

#### Changed

- Chart: Update `cluster` to v5.1.0.
- Chart: Update `cluster` to v5.0.0.

### Apps

- coredns from v1.28.2 to v1.28.3
- external-dns from v3.2.0 to v3.3.0
- Added priority-classes v0.3.0

### coredns [v1.28.2...v1.28.3](https://github.com/giantswarm/coredns-app/compare/v1.28.2...v1.28.3)

#### Changed

- Update `coredns` image to [1.13.2](https://github.com/coredns/coredns/releases/tag/v1.13.2).

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
