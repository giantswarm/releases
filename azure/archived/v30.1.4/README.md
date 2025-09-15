# :zap: Giant Swarm Release v30.1.4 for Azure :zap:

## Changes compared to v29.5.1

### Components

- cluster-azure from v1.6.1 to v2.1.2
- Flatcar from v4081.2.1 to [v4152.2.1](https://www.flatcar-linux.org/releases/#release-4152.2.1)
- Kubernetes from v1.29.13 to [v1.30.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.30.md#v1.30.11)
- os-tooling from v1.22.1 to v1.24.0

### cluster-azure [v1.6.1...v2.1.2](https://github.com/giantswarm/cluster-azure/compare/v1.6.1...v2.1.2)

#### Added

- Add `global.metadata.preventDeletion` to add the [deletion prevention label](https://docs.giantswarm.io/advanced/deletion-prevention/) to cluster resources
- Enable network-policies-app from `cluster` chart with DNS policies enabled. This makes `kube-system` and `giantswarm` namespaces to be `deny` by default.
- Add default tag `giantswarm-cluster` to all resources.
- Allow adding custom tags to resources using `providerSpecific.additionalResourceTags` value.
- Add `azurefile-csi-driver` app as helmrelease.
- Add validation of machineDeployment name using Schema Regex
- Add flags to disable PSPs.
- Render `containerd` configuration at cluster creation time
  - add support for containerd registry mirrors
  - add support for containerd registry credentials
- Set value for `controller-manager` `terminated-pod-gc-threshold` to `125` ( consistent with vintage )
- Deploy private links for private clusters.
- Add full configuration values documentation.
- Add support for private clusters.
- Add support for failuredomains field to MachineDeployments
- Generate SAN entries for `api.<clusterName>.<baseDomain>` (e.g. `api.glippy.azuretest.gigantic.io`)
- Add option to specify the `giantswarm.io/service-priority` cluster label.
- Add icon property to Chart metadata.
- Pre-Create /var/lib/kubelet with `0750` if it does not exist already to address issue with node-exporter
- Add example manifests to create cluster
- Add support for Bastion host as a MachineDeployment
- Add support for MachineDeployments
- Add MachineDeployments to Values.yaml
- Add MachineHealthChecks for Worker Nodes in MachineDeployments. Enabled by default
- Enable PodSecurityPolicy admission plugin when version is `lt` 1.25.0
- Add helm chart dependency for `cluster-shared` , required by the PSP admission controller
- Default to 3 replicas for control plane
- add giantswam user to the KCP and Machinepool configuration
- Add support for custom taints and labels on machinepools
  - also add hardcoded `role=worker` and `giantswarm.io/machine-pool` labels
- Add support for custom taints on control plane nodes
- Set EvictionThresholds soft and hard on all nodes
- Add a script to calculate the `kube-reserved` settings for nodes based on the available CPU and Memory using the formulas defined by [GKE](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#memory_cpu)
  - The memory reservation is slighly less aggressive than what GKE suggests
- Initial support to create a workload cluster via CAPI/CAPZ.
- Add support for creating cluster with `UserAssigned Identity` for `VM Identity`
- Add `cluster.x-k8s.io/watch-filter: capi` to common labels.
- Added github automation

#### Changed

- Chart: Update `cluster` to v2.2.3.
- Chart: Update Cilium configuration.
- Chart: Update `cluster` to v2.2.1.
- Chart: Update `cluster` to v2.2.0.
- Chart: Update `cluster` to v2.1.1.
- Chart: Reduce default etcd volume size to 50 GB.
- Chart: Update `cluster` to [v1.7.0](https://github.com/giantswarm/cluster/releases/tag/v1.7.0).
  - Add `teleport-init` systemd unit to handle initial token setup before `teleport` service starts
  - Improve `teleport` service reliability by adding proper file and service dependencies and pre-start checks
- Make `external-dns-private` app depend on the `prometheus-operator-crd` app, because it uses `ServiceMonitors`.
- Chart: Update `cluster` to [v1.4.1](https://github.com/giantswarm/cluster/releases/tag/v1.4.1)
  - Allow to enable auditd service through `global.components.auditd.enabled`.
  - Allow configuring `kube-controller-manager` `--node-cidr-mask-size` flag.
- Chart: Update `cluster` to [v1.2.2](https://github.com/giantswarm/cluster/releases/tag/v1.2.2)
  - Set `MachineDeployment` Kubernetes version from release
- Apps: Use `catalog` from Release CR.
- Chart: Update `cluster` to v1.1.0. ([#325](https://github.com/giantswarm/cluster-azure/pull/325))
  - Machine Template: Adapt new image format.
  - Apps: Enable `observability-policies`.
- Update cluster chart version to v1.0.0. This update adds MC Zot deployment as a registry mirror for `gsoci.azurecr.io` registry. This is the new default behavior.
- Respect `global.apps.externalDnsPrivate` to overwrite configuration of `external-dns-private` app.
- Add `allowedSubscriptions` parameter for multi-subscription use case.
- Use `.Values.global.managementCluster` for teleport node labels.
- Update `azurefile-csi-driver-app` to `1.30.2-gs1`
- Update `cluster` chart to v0.32.0. More details in [cluster chart v0.32.0 release notes](https://github.com/giantswarm/cluster/releases/tag/v0.32.0).
- Use MachineHealth resource from `cluster` chart.
- Use MachineDeployment resource from `cluster` chart.
- Update cluster chart to v0.27.0. More details in [cluster chart v0.27.0 release notes](https://github.com/giantswarm/cluster/releases/tag/v0.27.0).
- Disable and remove permissive policies from cilium-app.
- Import HelmRepositories from `cluster` chart and delete the HelmRepositories from this chart. This adds the `cluster-catalog`.
- Bump `cluster` chart from `0.21.0` to `0.26.0`.
- Use KubeadmControlPlane resource from `cluster` chart.
- Bump flatcar to `3815.2.0`.
- Add `cluster` chart as subchart.
- Render Cluster resource from the `cluster` chart.
- Delete Cluster resource template.
- Bump `azurefile-csi-driver-app` to `1.26.0-gs5`.
- Update teleport node labels - add `ins=` label and remove `cluster=` label condition check, such that MC nodes have this label.
- Allow additional fields for `privateEndpoints`.
- Allow adding `privateEndpoints` to subnets.
- Use `Standard_D4s_v5` for control plane and worker nodes.
- Use 2 replicas for workers by default.
- Upgrade K8S version to `1.25.16`.
- Upgrade kubectl version to `1.25.15`.
- Disable PSPs by default.
- Enable Host Encryption for workers and control plane virtual machines.
- Use `gsoci.azurecr.io` for `kubeadm` container images.
- Use `gsoci.azurecr.io` for sandbox container image (pause container).
- Update `coredns` to `1.21.0` to use `gsoci.azurecr.io`.
- Update `cillium` to `0.19.2` to use `gsoci.azurecr.io`.
- Update `azure-cloud-controller-manager-app` to `1.24.18-gs6` to use `gsoci.azurecr.io`.
- Update `azure-cloud-node-manager-app` to `1.24.18-gs6` to use `gsoci.azurecr.io`.
- Update `azuredisk-csi-driver-app` to `1.26.2-gs6` to use `gsoci.azurecr.io`.
- Update `azurefile-csi-driver-app` to `1.26.0-gs4` to use `gsoci.azurecr.io`.
- Enable teleport by default.
- Upgrade Flatcar image to [3510.2.5](https://www.flatcar.org/releases#release-3510.2.5)
- Upgrade K8S version to `1.24.17`
- Fix left-over azurefile-csi-driver helmreleases during cleanup.
- Adapt cleanup hook for cluster policies.
- :boom: Migrate CNI / CPI / CSI and VPA CRD apps to helmreleases in cluster-azure - requires `default-apps-azure` 0.0.24
- Use multiple volumes for `containerd`,`kubelet`,`root` and `etcd` mounts for **control plane** nodes
- Restricted `.providerSpecific.location` value to a set of defined region names.
- Disallow additional properties on the `.metadata.labels` object.
- Value `.providerSpecific.subscriptionId` marked as required, constrained to UUID format.
- `Enabled Admission Plugins` is now handled dynamically based on the kubernetes version of the cluster that is being installed
- `Feature Gates` is now handled dynamically based on the kubernetes version of the cluster that is being installed
- Add support for configurable labels to Cluster CR
- Upgrade `cluster-shared` dependency to `0.6.5`
- Add value schema constraints to all numeric types, using `exclusiveMinimum` or `minimum` of zero.
- Upgrade Flatcar image to [3510.2.1](https://www.flatcar.org/releases#release-3510.2.1)
- Upgrade K8S version to `1.24.13`
- :boom: Breaking - Skip `kube-proxy` during kubeadm init/join to replace with cilium-proxy
  - This change requies default-apps >= 0.0.17
- Add `identity spec` to hash calculation for bastion node
- Add `connectivity.allowedCIDRs` to define a list of network addresses to connect to the API server.
- Support defining custom vnet settings ( in the /internal section  of the schema )
  - VNET name and ResourceGroup
  - precreated subnet names
- Allow defining the scope of the SystemAssigned Identity on WC nodes
- Rename JSON schema makefile commands to `normalize-schema`, `validate-schema`, `generate-values`.
- Add replacement of pause image for kubelet and containerd to use `quay.io/giantswarm/pause`
- Revert `cilium kube-proxy` replacement - do not skip kube-proxy
  - Requires default-apps => 0.0.15
- :boom: Breaking - Skip `kube-proxy` during kubeadm init/join to replace with cilium-proxy
  - This change requies default-apps >= 0.0.14
- Remove machinepool code , this code is currently not used and it will confused the team that picks up this APP
- Add support for creating WC with SystemAssigned Identities and make it the default - `Contributor` Role in the `resourceGroup` where the cluster Lives
- Switch Cluster Images from Ubuntu to Flatcar
- Port hardening and tuning settings from Vintage to CAPZ Flatcar
- Fix `schema-normalize` Make target to actually do the normalize
- **Breaking change** to values schema - make sure to update your values before updating to this releaseValues schema:
  - Rename /machineDeployments to /nodePools
  - Remove /machinePools from schema
- Values schema: Use draft 2020-12 and update default value encoding based on latest `schemalint normalize` output.
- Cluster Example: Update to match release 0.0.12 changes
- Add `managementCluster`, `baseDomain` and `provider` properties to the schema because they are added by the AppOperator and the schema has `additionalProperties: false`
- Re-Add selector to Bastion machineDeployment , this is a required field and the webhook validation fail without it ( only in our kind mc-bootstrap)
- Update example manifests to create cluster
- Re-Add selector to machineDeployment , this is a required field and the webhook validation fail without it ( only in our kind mc-bootstrap)
- Disallow additional properties on the values scherma root level.
- Reduce default network range from 10.0.0.0/8 (default CAPZ) to 10.0.0.0/16.
- **Breaking change** to values schema - make sure to update your values before updating to this releaseValues schema:
  - Renamed /azure to /providerSpecific
  - Moved /bastion to /connectivity/bastion
  - Moved /oidc to /controlPlane/oidc
  - Moved /defaults to /internal/defaults
  - Moved /attachCapzControllerIdentity into /internal/identy
  - Moved /enablePerClusterIdentity into /internal/identy
  - Moved /sshSSOPublicKey to /connectivity/sshSSOPublicKey
  - Moved /kubernetesVersion to /internal/kubernetesVersion
- Move common templates between MachineDeployments and MachinePools into an helper file ( \_machine_helpers.tpl )
- replace version with `0.0.0-dev` in Chart.yaml since we use App Build Suite
- Allow customizing the `identityRef` in the `AzureCluster`
- Fix MachinePool naming by removing the hashed name from all resources. This is not needed for MachinePools , like it is for MachineDeployments
- Skip `coredns` installation phase in `kubeadmbootstrapconfiguration` , we install it as an App
- Do not consider the `labels` in the ControlPlane AzureMachineTemplate when calculating name hash to avoid rolling control plane nodes unecessarily
- Change default values ssh key to RSA one ( since azure does not support ed25519 )
- Update schema json

#### Fixed

- Use correct context at `MachineDeployment` helper.
- Render external-dns for Azure private clusters correctly.
- Render cert-manager configuration for Azure private clusters correctly.
- Add missing hack for manipulating /etc/hosts for private clusters.
- Fix containerd config that was breaking in newer flatcar versions.

#### Removed

- Remove Cilium deprecated values.
- Remove unused `internal` values from `values.schema.json`.
- Drop duplicated workflow "compare_rendering.yaml" file.
- SSH inbound SG rule from VPN
- Bastion and ssh configuration on nodes.
- Remove CSIMigration feature flag (enabled by default with k8s 1.23).
- Removed `baseDomain` from CI values.
- Values schema
  - Removed redundant and unused /clusterName and /clusterDescription properties.
  - Removed unused /includeClusterResourceSet

### Apps

- Added coredns-extensions v0.1.2
- Added etcd-defrag v1.0.2
- azure-cloud-controller-manager from v1.29.8-gs1 to v1.30.14-gs1
- azure-cloud-node-manager from v1.29.8-gs1 to v1.30.14-gs1
- azuredisk-csi-driver from v1.30.2-gs2 to v1.30.12-gs1
- azurefile-csi-driver from v1.30.2-gs1 to v1.30.10-gs1
- capi-node-labeler from v0.5.0 to v1.0.2
- cert-exporter from v2.9.3 to v2.9.5
- cert-manager from v3.8.2 to v3.9.0
- cilium from v0.25.2 to v0.31.5
- cilium-servicemonitors from v0.1.2 to v0.1.3
- coredns from v1.23.0 to v1.24.0
- etcd-k8s-res-count-exporter from v1.10.0 to v1.10.3
- external-dns from v3.1.0 to v3.2.0
- k8s-audit-metrics from v0.10.0 to v0.10.2
- metrics-server from v2.4.2 to v2.6.0
- net-exporter from v1.21.0 to v1.22.0
- node-exporter from v1.20.0 to v1.20.2
- observability-bundle from v1.9.0 to v1.11.0
- security-bundle from v1.9.1 to v1.10.1
- teleport-kube-agent from v0.10.3 to v0.10.4
- vertical-pod-autoscaler from v5.3.1 to v5.4.0
- vertical-pod-autoscaler-crd from v3.1.2 to v3.2.0


### azure-cloud-controller-manager [v1.29.8-gs1...v1.30.14-gs1](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v1.29.8-gs1...v1.30.14-gs1)

#### Added

- Add Annotations and labels for use of azure workload identity.

#### Changed

- Chart: Update to upstream v1.30.14. ([#110](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/110))
- Chart: Update to upstream v1.30.13. ([#103](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/103))
- Chart: Update to upstream v1.30.6. ([#87](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/87))

### azure-cloud-node-manager [v1.29.8-gs1...v1.30.14-gs1](https://github.com/giantswarm/azure-cloud-node-manager-app/compare/v1.29.8-gs1...v1.30.14-gs1)

#### Added

- Add Annotations and labels for use of azure workload identity.

#### Changed

- Chart: Update to upstream v1.30.14. ([#100](https://github.com/giantswarm/azure-cloud-node-manager-app/pull/100))
- Chart: Update to upstream v1.30.13. ([#93](https://github.com/giantswarm/azure-cloud-node-manager-app/pull/93))
- Chart: Update to upstream v1.30.6. ([#77](https://github.com/giantswarm/azure-cloud-node-manager-app/pull/77))

### azuredisk-csi-driver [v1.30.2-gs2...v1.30.12-gs1](https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.30.2-gs2...v1.30.12-gs1)

#### Changed

- Chart: Update to upstream v1.30.12.

### azurefile-csi-driver [v1.30.2-gs1...v1.30.10-gs1](https://github.com/giantswarm/azurefile-csi-driver-app/compare/v1.30.2-gs1...v1.30.10-gs1)

#### Changed

- Chart: Update to upstream v1.30.10.

### capi-node-labeler [v0.5.0...v1.0.2](https://github.com/giantswarm/capi-node-labeler-app/compare/v0.5.0...v1.0.2)

#### Changed

- Go: Update dependencies.
- Main: Improve sleep. ([#125](https://github.com/giantswarm/capi-node-labeler-app/pull/125))
- Go: Update `go.mod` and `.nancy-ignore`. ([#123](https://github.com/giantswarm/capi-node-labeler-app/pull/123))

### cert-exporter [v2.9.3...v2.9.5](https://github.com/giantswarm/cert-exporter/compare/v2.9.3...v2.9.5)

#### Changed

- Go: Update dependencies.
- Repository: Some chores. ([#418](https://github.com/giantswarm/cert-exporter/pull/418))
- Go: Update `go.mod` and `.nancy-ignore`. ([#437](https://github.com/giantswarm/cert-exporter/pull/437))

### cert-manager [v3.8.2...v3.9.0](https://github.com/giantswarm/cert-manager-app/compare/v3.8.2...v3.9.0)

#### Added

- Adds new sync method based on Vendir to sync from upstream

#### Changed

- Updates Cert-manager Chart to Upstream 1.16.2

### cilium [v0.25.2...v0.31.5](https://github.com/giantswarm/cilium-app/compare/v0.25.2...v0.31.5)

#### Changed

- Reenable Cilium agent metrics.
- Upgrade Cilium to [v1.16.10](https://github.com/cilium/cilium/releases/tag/v1.16.10).
- Upgrade Cilium to [v1.16.9](https://github.com/cilium/cilium/releases/tag/v1.16.9).
- Upgrade Cilium to [v1.16.8](https://github.com/cilium/cilium/releases/tag/v1.16.8).
- Upgrade Cilium to [v1.16.7](https://github.com/cilium/cilium/releases/tag/v1.16.7).
- Upgrade Cilium to [v1.16.6](https://github.com/cilium/cilium/releases/tag/v1.16.6).
- Upgrade Cilium to [v1.16.5](https://github.com/cilium/cilium/releases/tag/v1.16.5).
- Revert back to upstream default of using image digests for container images.
- Upgrade Cilium to [v1.16.3](https://github.com/cilium/cilium/releases/tag/v1.16.3).
- Move provider specific custom CNI configuration to subchart.
- Upgrade Cilium to [v1.16.1](https://github.com/cilium/cilium/releases/tag/v1.16.1).
- Upgrade Cilium to [v1.16.0](https://github.com/cilium/cilium/releases/tag/v1.16.0).
- Disable digest in all images.
- Improve security defaults for:
  - Hubble UI
  - Hubble Relay
  - Cilium Operator

#### Removed

- Delete defaultPolicies and extraPolicies templates.

### cilium-servicemonitors [v0.1.2...v0.1.3](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.2...v0.1.3)

#### Changed

- Change ownership from `phoenix` to `cabbage`.
- Use the app-build-suite.

### coredns [v1.23.0...v1.24.0](https://github.com/giantswarm/coredns-app/compare/v1.23.0...v1.24.0)

#### Changed

- Update `coredns` image to [1.12.0](https://github.com/coredns/coredns/releases/tag/v1.12.0).
- Disable HPA Memory target.
- Increase threshold for HPA CPU target to 80%.

### coredns-extensions [v0.1.2](https://github.com/giantswarm/coredns-extensions/releases/tag/v0.1.2)

#### Added

- Add VPA for CoreDNS deployments.
- Add value to enable or disable VPA resources.

#### Changed

- Push App to the default-catalog.
- Publish App in giantswarm-catalog.

### etcd-defrag [v1.0.2](https://github.com/giantswarm/etcd-defrag/releases/tag/v1.0.2)

#### Added

- Chart: Add `moveLeader`. ([#11](https://github.com/giantswarm/etcd-defrag-app/pull/11))

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.25.0. ([#17](https://github.com/giantswarm/etcd-defrag-app/pull/17))
- Chart: Update dependency ahrtr/etcd-defrag to v0.24.0. ([#16](https://github.com/giantswarm/etcd-defrag-app/pull/16))
- Chart: Update dependency ahrtr/etcd-defrag to v0.23.0. ([#10](https://github.com/giantswarm/etcd-defrag-app/pull/10))
- Values: Rename `cluster` into `useClusterEndpoints`. ([#8](https://github.com/giantswarm/etcd-defrag-app/pull/8))

### etcd-k8s-res-count-exporter [v1.10.0...v1.10.3](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.0...v1.10.3)

#### Changed

- Go: Update dependencies.
- Go: Update dependencies.
- Set `readOnlyRootFilesystem` to true in the container security context.
- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod` and `.nancy-ignore`. ([#242](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/pull/242))

### external-dns [v3.1.0...v3.2.0](https://github.com/giantswarm/external-dns-app/compare/v3.1.0...v3.2.0)

#### Changed

- Update architect-orb and ATS.
- Add DNSEndpoints as a source for DNS records.

### k8s-audit-metrics [v0.10.0...v0.10.2](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.0...v0.10.2)

#### Changed

- Go: Update dependencies.
- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod` and `.nancy-ignore`. ([#248](https://github.com/giantswarm/k8s-audit-metrics/pull/248))

### metrics-server [v2.4.2...v2.6.0](https://github.com/giantswarm/metrics-server-app/compare/v2.4.2...v2.6.0)

#### Added

- Add VPA setting for `metrics-server`.
- Chart: Update PolicyExceptions to v2beta1. ([#226](https://github.com/giantswarm/metrics-server-app/pull/226))

#### Changed

- Upgrade metrics-server to v0.7.2.

### net-exporter [v1.21.0...v1.22.0](https://github.com/giantswarm/net-exporter/compare/v1.21.0...v1.22.0)

#### Changed

- Narrow down CiliumNetworkPolicy to allow desired traffic only.

#### Removed

- Remove NetworkPolicy resource and rely on CiliumNetworkPolicy only.

### node-exporter [v1.20.0...v1.20.2](https://github.com/giantswarm/node-exporter-app/compare/v1.20.0...v1.20.2)

#### Changed

- Go: Update dependencies.
- Update Kyverno `PolicyExceptions` to `v2beta1`.
- Go: Update `go.mod`. ([#322](https://github.com/giantswarm/node-exporter-app/pull/322))

### observability-bundle [v1.9.0...v1.11.0](https://github.com/giantswarm/observability-bundle/compare/v1.9.0...v1.11.0)

#### Changed

- prometheus-operator will not check promql syntax for prometheusRules that are labelled `observability.giantswarm.io/rule-type: logs`
- Upgrade `alloy` to chart 0.9.0.
  - Bumps `alloy` from to 1.5.1 to 1.7.1
- Upgrade `alloy` to chart 0.8.0.
  - Bumps `alloy` from to 1.5.0 to 1.6.1
- Upgrade `kube-prometheus-stack` from 66.2.1 to 69.5.1
  - Bumps prometheus-operator to 0.80.1
  - Bumps prometheus to 3.0.1

### security-bundle [v1.9.1...v1.10.1](https://github.com/giantswarm/security-bundle/compare/v1.9.1...v1.10.1)

#### Added

- Add e2e tests for the `security-bundle` and all is components

#### Changed

- Update `kyverno-crds` (app) to v1.13.1.
- Update `kyverno` (app) to v0.19.0.
- Update `kyverno-crds` (app) to v1.13.0.
- Update `kyverno-policies` (app) to v0.23.0.
- Update `edgedb` (app) to v0.1.0.
- Update `falco` (app) to v0.10.0.
- Update `trivy` (app) to v0.13.2.

### teleport-kube-agent [v0.10.3...v0.10.4](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.3...v0.10.4)

#### Added

- Add headless service on `diag` port 3000.

#### Changed

- Migrated to ABS

### vertical-pod-autoscaler [v5.3.1...v5.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.3.1...v5.4.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v10.0.0 ([#335](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/335))

### vertical-pod-autoscaler-crd [v3.1.2...v3.2.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v3.1.2...v3.2.0)

#### Changed

- Chart: Sync to upstream. ([#126](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/126))
