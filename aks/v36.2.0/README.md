# :zap: Giant Swarm Release v36.2.0 for AKS :zap:

## Changes compared to v36.1.0

### Components

- Kubernetes from v1.36.4 to [v1.36.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.5)

### Apps

- cert-manager-crossplane-resources from v0.2.0 to v0.3.0
- network-policies from v0.3.2 to v0.5.0
- observability-bundle from v3.5.0 to v3.6.0

### cert-manager-crossplane-resources [v0.2.0...v0.3.0](https://github.com/giantswarm/cert-manager-crossplane-resources/compare/v0.2.0...v0.3.0)

#### Added

- Read the AWS account ID, partition, and OIDC domains from the `<clusterName>-crossplane-config` ConfigMap if `providers.aws.accountID`, `providers.aws.partition`, or `providers.aws.oidc.domain` is empty. This supports EKS clusters, whose OIDC domain is only known once the cluster is provisioned. Explicitly set values take precedence over the ConfigMap.

#### Changed

- Restrict AWS provider permissions for `ChangeResourceRecordSets` to TXT records related to acme challenges.
- Create a dedicated role for azure to reduce permissions to TXT records only.

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
