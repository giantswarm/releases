# :zap: Giant Swarm Release v29.6.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v29.5.0

### Components



* cluster-aws from 2.5.0 to [2.5.0](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...v2.5.0)






### cluster-aws [2.5.0](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...v2.5.0)

#### Added

- Add aws-node-termination-handler bundle
- Values: Add `global.providerSpecific.controlPlaneAmi` & `global.providerSpecific.nodePoolAmi`.
- Make ASG lifecycle hook heartbeat timeout configurable
- Add `global.providerSpecific.additionalNodeTags`. Field used to specify tags applied to nodes only.
- Expose the `maxHealthyPercentage` property to allow setting the maximum percentage of healthy machines in the Auto Scaling Group during upgrades.
- Allow to enable `auditd` through `global.components.auditd.enabled` helm value.
- Chart: Support multiple service account issuers.\
- Add `global.metadata.labels` to values schema. This field is used to add labels to the cluster resources.
- Enable `observability-policies` app.
- Add the Management Cluster name as a tag to the AWS resources created by CAPA.
- Add the node pool name as a tag to the AWS resources associated with the node pool.
- Add `irsa-servicemonitors` and `aws-ebs-csi-driver-servicemonitors` apps.
- Add configuration for `aws-pod-web-identity-webook` app to include region into the IRSA enabled pods.
- Chart: Add `aws-pod-identity-webhook` app. ([#581](https://github.com/giantswarm/cluster-aws/pull/581)).
- Worker nodes - Add `nonRootVolumes` fields to mount `/var/lib` and `/var/log` on separate disk volumes.
- BREAKING CHANGE: values `global.controlplane.containerdVolumeSizeGB` and `global.controlplane.kubeletVolumeSizeGB` merged into single value `.global.controlPlane.libVolumeSizeGB` which define size of disk volume used for `/var/lib` mount point.
- Make Cilium ENI-based IP allocation configurable with high-level `global.connectivity.cilium.ipamMode` value. This feature was previously introduced as prototype and is now fully working.
- Allow to configure SELinux mode through `global.components.selinux.mode` helm value.
- Add `log` volume to control-plane nodes.
- Add option to configure instance metadata http tokens for EC2 instances to enable or disable IMDSv2 enforcement.
- Allow customizing instance refresh parameters.
- Smart defaulting for AWS availability zones using actual AZs in the region of choice rather than hardcoded values.
- Make Cilium ENI-based IP allocation configurable with new high-level `global.connectivity.cilium.ipamMode` value (prototype)
- Add automatic support for deploying to AWS China.
- Add network-policies helm release.
- Allow customers to specify optional extraConfigs in HelmRelease apps.
- Include cluster-test-catalog in the CI, so we can more easily test dev builds of subcharts.
- Add propagating tags from `cluster-aws` to resources managed my `ebs-csi-driver`.
- CI: trigger automated e2e tests on Renovate PRs.
- Add new annotation for vintage irsa domain which is only used for migrating vintage clusters.
- Use 443 as the default api-server Load Balancer port.
- Make Helm values configurable for aws-cloud-controller-manager, aws-ebs-csi-driver, cilium, coredns and vertical-pod-autoscaler-crd
- Expose value to configure launch template overrides, used to override the instance type specified by the launch template with multiple instance types that can be used to launch instances.
- Add `global.metadata.preventDeletion` to add the [deletion prevention label](https://docs.giantswarm.io/advanced/deletion-prevention/) to cluster resources
- Allow cluster-autoscaler handling MachinePools.
- Add additional tag for cluster autoscaler to MachinePool ASGs.
- Add option to force CGroups v1.
- Allow configuration of `AWSCluster.spec.AdditionalTags` value and add a default giantswarm tag.
- Add teleport.service: Secure SSH access via Teleport.
- Add `controlPlane.loadBalancerIngressAllowCidrBlocks` to configure control plane load balancer ingress rules.
- Add values neccessery for migration from vintage.
- Add support for Spot instances.
- Support creating `CiliumNetworkPolicy` manifests that allow egress requests to DNS and conditionally the proxy host (via [`cilium-app`](https://github.com/giantswarm/cilium-app))
- Set value for `controller-manager` `terminated-pod-gc-threshold` to `125` ( consistent with vintage )
- Fix defaulting of node pool for AWSCluster CR.
- Add CNI/CSI/coredns apps as HelmReleases.
- Migrating from Ubuntu AMI to Flatcar AMI is a **breaking change** that requires manual steps.
- Add JSON schema related makefile.
- generate `values.yaml` from `values.schema.json` with `make generate-values`
- normalize `values.schema.json` with `make normalize-schema`
- validate that `values.schema.json` is according to requirements with `make validate-schema`
- Add full configuration values documentation.
- Add `"helm.sh/resource-policy": keep` annotation to AWSCluster,
- Values schema:
- Add /managementCluster and /provider to account for values injected by controllers.
- Add `MachineHealthCheck` for control plane nodes.
- Made registry configurations `connectivity.containerRegistries` dynamic to accept as many container registries and mirrors as needed.
- Expose helm value for customers to decide whether VPC endpoint should be created by Giantswarm.
- Customize tags per individual subnet.
- Add value to specify which AWS account ID to use when associating Route53 Resolver Rules with workload cluster VPC.
- More configuration options when defining subnets to be created
- `controlPlane.subnetTags`, `bastion.subnetTags` and `machinePools[].subnetTags` to target specific subnets
- Add icon to Chart.yaml
- Add and propagate `no_proxy` value to the underlying components.
- Add cluster base domain to no proxy config.
- Add schema for items of the arrays `.machinePools[*].availabilityZones` and `.machinePools[*].customNodeTaints`.
- Add IRSA domain placeholder replacer as postKubeadm script.
- Add `containerd` registry auth workaround to bug https://github.com/giantswarm/roadmap/issues/1737.
- Add option to specify oidc CA PEM in order to autheticate againts OIDC with custom CA.
- Add option to configure containerd registry authentication for `docker.io`.
- Add external resource gc annotation.
- Add full proxy configuration for private clusters.
- Support setting node taints using `customNodeTaints`
- IRSA for CAPA.
- Make subnets configurable.
- Set `aws.giantswarm.io/vpc-mode` annotation on AWSCluster.
- Set cluster to paused when vpcMode is set to private.
- Support for specifying private VPC configuration (not yet used)
- Support for specifying private DNS zone configuration.
- Validation of vpcMode and apiMode combination being valid
- Add support for configuring outgoing proxy for the cluster.
- Allow configuration of loadbalancer for Control Plane API (`internet-facing` will be default).
- Network topology mode annotations
- Add role label to bastion machine.
- `hash` function to ensure immutable resources change be changed via recreate/replacement
- Add OIDC support for k8s api.
- Add `localhost` and `api` domain to the certSANs of apiserver certificates.
- Support for specifying the `clusterName` (defaults to chart release name)
- Lookup AWS region if not set in values
- Lookup AWS Availability Zones if not set in values
- Allow app platform to take over managing coredns
- Rename `networkSpec` to `network` in AWSCluster CR due renaming in `v1beta1`.
- Prefix machine pool with cluster id.
- Set etcd max db size to 8 GB.
- Add encryption provider config for k8s secrets.
- Add `audit-policy` to kubernetes api.
- Fix AWSMachinePool min and max values.
- Add labels to machine metadata to `AWSMachineTemplate` CRs.
- Add `sourceIdenityRef` to AWSClusterRoleIdentity CR.
- Fix aws cluster role identity value reference.

#### Changed

- Chart: Update `cluster` to [v1.7.0](https://github.com/giantswarm/cluster/releases/tag/v1.7.0).
- Add `teleport-init` systemd unit to handle initial token setup before `teleport` service starts
- Improve `teleport` service reliability by adding proper file and service dependencies and pre-start checks
- Chart: Update `cluster` to v1.4.1.
- Set provider specific configuration for cilium CNI ENI values.
- Validate that machine pool availability zones belong to the selected region.
- CI: Bump release version.
- Apps: Use `catalog` from Release CR.
- Update cluster chart to v1.1.0.
- This sets cilium `kubeProxyReplacement` config to `"true"` instead to `"strict"` (`"strict"` has been deprecated since cilium v1.14, see [this upstream cilium](https://github.com/cilium/cilium/issues/32711) issue for more details).
- Update `ami` named template to correctly render OS image name with the new format `flatcar-stable-<flatcar version>-kube-<kubernetes version>-tooling-<capi-image-builder version>-gs`.
- Update cluster chart version to v1.0.0. This update adds MC Zot deployment as a registry mirror for `gsoci.azurecr.io` registry. This is the new default behavior.
- Apps: Fix service monitor dependencies.
- Chart: Update `cluster` chart to v0.36.0. ([#703](https://github.com/giantswarm/cluster-aws/pull/703))
- Update cluster chart to 0.35.0
- Update cluster chart to 0.33.1
- First major release, breaking changes not allowed in minor releases anymore.
- Set maximum number of pods (kubelet `--max-pods`) for Cilium ENI mode due to restrictions by number of ENIs and IPs per ENI. This change will roll nodes even if not using Cilium ENI mode, since a new script is introduced.
- Update cluster chart to 0.31.4
- Update cluster chart to 0.31.2
- Set `prometheus-blackbox-exporter` and `k8s-audit-metrics` as enabled.
- Set environment variable `COREOS_EC2_IPV4_LOCAL` to inject value to kubeadm configuration.
- Set environmane variable `COREOS_EC2_HOSTNAME` to inject value to kubeadm configuration.
- Update aws-cloud-controller-manager-app to v1.25.14-gs3.
- Update cluster chart to v0.27.0. More details in [cluster chart v0.27.0 release notes](https://github.com/giantswarm/cluster/releases/tag/v0.27.0).
- Always render `userConfig` values reference to configmap for `aws-pod-identity-webhook-app`.
- Store EC2 user data (Ignition config) for machine pools in S3 bucket to overcome the size limit (requires this new CAPA feature and `AWSMachinePool.spec.ignition` field)
- Update CAPA CR references to API version `v1beta2`
- Set token audience for `aws-pod-identity-webhook` based on AWS region.
- Update cluster chart to v0.25.0 and enable all default apps. More details in [cluster chart v0.25.0 release notes](https://github.com/giantswarm/cluster/releases/tag/v0.25.0).
- Upgrade default-apps-aws App to v0.52.0 or newer.
- Update default-apps-aws Helm value `.Values.deleteOptions.moveAppsHelmOwnershipToClusterAws` to true.
- All apps, except observability-bundle and security-bundle will get `app-operator.giantswarm.io/paused: true` annotation, so wait few minutes for the change to get applied by the Helm post-upgrade hook.
- Delete default-apps-aws.
- App resources for all default apps will get deleted. Wait few minutes for this to happen.
- Chart resources on the workload cluster will stay, so all apps will continue running.
- Upgrade cluster-aws App to v0.76.0.
- cluster-aws will deploy all default apps, wait a few minutes for all Apps to be successfully deployed.
- Chart resources on the workload cluster will get updated, as newly deployed App resources will take over the reconciliation of the existing Chart resources.
- Control-plane nodes - combine kubelet disk `/var/lib/kubelet` and containerd disk `/var/lib/containerd` into single disk `/var/lib` to share the volume space and save cost.
- Update cluster chart to v0.22.0.
- Update cluster chart to v0.18.0. This updates teleport node labels and will roll nodes.
- Update instanceWarmup to 10' to be on pair with Vintage
- Enable extraPolicies from network-policies-app.
- Disable and remove extraPolicies from cilium-app.
- Values: Separate `app` and `helmRelease` definition. ([#581](https://github.com/giantswarm/cluster-aws/pull/581))
- Update cluster chart to v0.17.0. This updates cilium app from v0.21.0 to v0.22.0.
- Update Availability Zones in helm/cluster-aws/files/azs-in-region.yaml
- AMI: Use new AMI which includes latest teleport binary v15.1.7
- Chart: Bump `cluster` to v0.16.0.
- Use cleanup hook job HelmRelease from cluster chart.
- Chart: Bump `cluster` to v0.13.0.
- Change image lookup format for base OS image. osImageVariant is set to "2" for this kubernetes version. This is a **breaking change** that requires manual steps. For the next kubernetes versions, osImageVariant should not be set.
- Fix allow list API port 443.
- Chart: Bump `cluster` to v0.11.1.
- Chart: Bump `cluster` to v0.11.0.
- Use cilium and network-policies from cluster chart, and remove them from cluster-aws.
- Use default HelmRepositories from cluster chart.
- Use vertical-pod-autoscaler-crd HelmRelease from cluster chart.
- Use coredns HelmRelease from cluster chart.
- Remove default HelmRepositories from cluster-aws.
- Remove vertical-pod-autoscaler-crd HelmRelease from cluster-aws.
- Remove coredns HelmRelease from cluster-aws.
- Render MachineHealthCheck resource from the cluster chart.
- Remove MachineHealthCheck resource.
- Render MachinePool and KubeadmConfig resources from the cluster chart.
- Remove MachinePool and KubeadmConfig resources.
- Update cluster chart version to the latest v0.7.1 release.
- Render control plane resources from the cluster chart.
- Remove KubeadmControlPlane resource.
- Use `cluster.connectivity.proxy.noProxy` Helm template from cluster chart to render NO_PROXY in cluster-aws.
- Rename CI files, so they are used in GitHub action that checks Helm rendering.
- Remove ingress and egress rules from the security group that AWS creates by default when creating a new VPC.
- Remove unnecessary architect brackets cleanup.
- Use CI values to render templates locally.
- Bumped kubernetes version to 1.25.16. This change also enforces PSS.
- Use `gsoci.azurecr.io` for `kubeadm` container images.
- Use `gsoci.azurecr.io` for sandbox container image (pause container).
- Update `coredns` to `1.21.0` to use `gsoci.azurecr.io`.
- Update `aws-cloud-controller-manager` to `1.25.14-gs2` to use `gsoci.azurecr.io`.
- Bump cilium-app to v0.19.2 (upgrades Cilium to v1.14.5 and fixes a `CiliumNetworkPolicy` definition for reaching coredns)
- Remove allow-all Cilium network policies.
- Add topology annotations to AWSCluster
- Add `cluster` chart as subchart.
- Render Cluster resource from the `cluster` chart.
- Delete Cluster resource template.
- Add missing kubelet configuration to align it with vintage config.
- Remove bastion and ssh configuration on nodes.
- Fill `AWSCluster.spec.network.subnets[*].id` field for managed subnets for compatibility with CAPA v2.3.0
- Login to the management cluster and run the script (e.g: `./migrate.sh organization my-cluster`)
- Verify the `app.yaml` file and apply it to the management cluster (e.g: `kubectl apply -f app.yaml`)
- Move Helm values property `.Values.metadata` to `.Values.global.metadata`.
- Move Helm values property `.Values.connectivity` to `.Values.global.connectivity`.
- Move Helm values property `.Values.controlPlane` to `.Values.global.controlPlane`.
- Move Helm values property `.Values.nodePools` to `.Values.global.nodePools`.
- Move Helm values property `.Values.managementCluster` to `.Values.global.managementCluster`.
- Move Helm values property `.Values.baseDomain` to `.Values.global.connectivity.baseDomain`.
- Move Helm values property `.Values.providerSpecific` to `.Values.global.providerSpecific`.
- Move Helm values property `.Values.global.connectivity.containerRegistries` to `.Values.global.components.containerd.containerRegistries`.
- Bump the Kubernetes version to `v1.24.16`.
- Bump Teleport version to `v14.1.3`.
- Enable Teleport by default.
- Change schema validation allowing to add additional properties in `global`.
- Support longer node pool names and allow dashes.
- Bump cilium-app to v0.18.0 (upgrades Cilium to v1.14.3)
- Bump `coredns` version to `1.19.0` and fix values.
- Align job that cleans `HelmReleases` and `HelmCharts` with other providers.
- Remove dependency between `cilium` and CPI so that `cilium` is installed as soon as possible.
- Add `https://` scheme prefix to service-account-issuer URI
- Update kubernetes version to `1.24.14`.
- Use fixed alias CloudFront domain for IRSA
- Tolerate CAPI taints on uninitialized nodes when scheduling cilium relay and ui.
- Decrease `interval` on `HelmReleases` to make things more reactive.
- Remove finalizers from `HelmCharts` when removing this app to avoid leaving leftovers in the management cluster.
- Use CAPBK to provision bastion node with Flatcar AMI.
- Use CAPBK to provision control plane nodes with Flatcar AMI.
- Use CAPBK to provision worker nodes with Flatcar AMI.
- Migrating from Ubuntu AMI to Flatcar AMI is a **breaking change** that requires manual steps.
- Apply default OS setting for flatcar and os hardening.
- Update CAPA CRs API version from `v1beta1` to `v1beta2`.
- Values schema: disallow additional properties on the `.nodePools` object. This is a **breaking change** where node pool names are in use that do not match the pattern `^[a-z0-9]{5,10}$`.
- Removed `connectivity.network.podCidr` and `connectivity.network.serviceCidr`. Replaced by `connectivity.network.pods.cidrBlocks` and `connectivity.network.services.cidrBlocks`.
- Remove `app.kubernetes.io/version` from common labels. They are part of hashes, but we don't want to always roll nodes just because we are deploying a new version.
- Remove `architect` templating from `Chart.yaml` file.
- Remove control plane replicas value `controlPlane.replicas`. Now it's hardcoded to 3 nodes.
- Set `r6i.xlarge` as the new default AWS instance type for the control plane and node pools.
- Added value `.metadata.servicePriority` to the schema to set the cluster's relative priority.
- Updated `cluster-shared` chart dependency to `0.6.5`
- Moved the core components feature flags to their configuration, as the `featureGates` field is for `kubeadm` feature flags.
- Enable `CronJobTimeZone` feature gate in the kubelet.
- Set kubernetes `1.24.10` as the default version.
- Switch from the in-tree cloud-controller-manager to the external one. This requires version `v0.26.0` of `default-apps-aws`.
- Rename `defaultMachinePools` to `internal.nodePools` to fit new schema requirements and make clear that it should not be changed by customers.
- Default to using `giantswarm.azurecr.io` as Docker Hub mirror.
- Values schema:
- Added annotations
- Applied normalization using `schemalint normalize`
- Added property schema for /connectivity/containerRegistries
- Added property schema for subnetTags objects
- Added default values
- Move /ami to /providerSpecific/ami
- Move /awsClusterRoleIdentityName to /providerSpecific/awsClusterRoleIdentityName
- Move /region to /providerSpecific/region
- Move /flatcarAWSAccount to /providerSpecific/flatcarAwsAccount
- Move /clusterName to /metadata/name
- Move /clusterDescription to /medatada/description
- Move /organization to /metadata/organization
- Move /oidc to /controlPlane/oidc
- Move /bastion to /connectivity/bastion
- Move /network/serviceCIDR to /connectivity/network/serviceCidr
- Move /network/podCIDR to /connectivity/network/podCidr
- Move /proxy to /connectivity/proxy
- Rename /proxy/no_proxy to /connectivity/proxy/noProxy
- Rename /proxy/http_proxy to /connectivity/proxy/httpProxy
- Rename /proxy/https_proxy to /connectivity/proxy/httpsProxy
- Move /sshSSOPublicKey to /connectivity/sshSsoPublicKey
- Remove unused /includeClusterResourceSet
- Remove /aws/awsClusterRole (previously deprecated)
- Move /hashSalt to /internal/hashSalt
- Move /kubernetesVersion to /internal/kubernetesVersion
- Move /network/dnsMode to /connectivity/dns/mode
- Move /network/dnsAssignAdditionalVPCs to /connectivity/dns/additionalVpc and change to type array
- Move /network/vpcCIDR to /connectivity/network/vpcCidr
- Move /network/apiMode to /controlPlane/apiMode
- Move /network/resolverRulesOwnerAccount to /connectivity/dns/resolverRulesOwnerAccount
- Move /network/prefixListID to /connectivity/topology/prefixListId
- Move /network/topologyMode to /connectivity/topology/mode
- Move /network/transitGatewayID to /connectivity/topology/transitGatewayId
- Move /network/vpcEndpointMode to /connectivity/vpcEndpointMode
- Move /network/vpcMode to /connectivity/vpcMode
- Move /network/availabilityZoneUsageLimit to /connectivity/availabilityZoneUsageLimit
- Move /network/subnets to /connectivity/subnets
- Rename /machinePools to /nodePools
- Disallow additional properties on the root level
- Fail in Helm template if `dnsMode=public` is combined with a `baseDomain` ending with `.internal`.
- Set `/var/lib/kubelet` permissions to `0750` to fix `node-exporter` issue.
- Bump kubernetes version to `1.23.16`
- Subnets are now specified on the `AWSCluster` resource by default rather than relying on CAPA code to default them. The same sizing as the CAPA default have been used.
- Use Giant Swarm image repository for official Kubernetes images
- Override image repository to `registry.k8s.io` because kubeadm of Kubernetes v1.23.15 tries to pull the official image incorrectly, resulting in failing cluster upgrades, and `k8s.gcr.io` is outdated
- Change default NTP server as AWS NTP server.
- Deprecate confusingly named `aws.awsClusterRole` in favor of `aws.awsClusterRoleIdentityName`. The value refers to an `AWSClusterRoleIdentity` object, not directly to an IAM role name/ARN.
- Bump Kubernetes to 1.23.15
- Dowgrade to using Ubuntu 20.04 as base OS.
- Run bastion on private IP if vpc mode is set to private.
- Remove registry authetication workaround.
- Make `baseDomain` a required value.
- Allow scraping of k8s core components.
- Bump external-dns to latest release
- Make `kubeadm` skip the phase where it installs `coredns` as it will be installed by as a default app.
- Bumped Kubernetes to v1.23
- Update [cluster-shared](https://github.com/giantswarm/cluster-shared) from v0.3.0 to v0.6.3.
- Make `kubeadm` skip the phase where it installs `kube-proxy` as we will use `cilium` as a replacement.
- Enable tcp forwarding for sshd on bastion.
- Updated to Kubernetes 1.22.15
- Updated to using Ubuntu 22.04 as base OS
- `.Values.controlPlane.apiLoadbalancerScheme` has been removed in favour of `.Values.network.apiMode`
- Default network topology mode changed to 'None'
- Use our Giant Swarm built AMIs
- Bump default Kubernetes version to 1.22.12
- Set pod CIDR to 100.64.0.0/12 to match what we set in Cilium (and to not clash with AWS CIDR)
- Fix values schema.
- Make bastion optional.
- Add team label to helm resources.
- Add `values.schema.json` file.
- Remove helm lookup function for SSH CA cert and use value fro central vault instead.
- Updated to latest `cluster-shared` library chart
- Switched to using `cluster-shared` for PSPs and coredns-adopter
- Upgrade to `vbeta1` version for all CRs.

#### Fixed

- Fix aws-nth-bundle to use the MC's kubeconfig context if it's in a different organization namespace.
- Only try to render subnet tags if they are defined by the user.
- Fixed China IRSA suffix
- Added an annotation to Kubernetes resources to resolve an issue where deletion was stuck due to hanging load balancers.
- Update aws-ebs-csi-driver-app from v2.30.0 to v2.30.1. This fixes accidental installation of PSPs which could break the upgrade to previous `cluster-aws` versions which didn't have this fix yet.
- Fix selecting right AWS region for Machine Pools when cluster is in different AWS region than MC.
- Update Chart.lock with current version of dependencies.
- Pass `clusterID` to `aws-ebs-csi-driver` app's values for volume tagging purposes.
- Update network-policies to avoid installing deny-all policy.
- Remove duplicate containerd config as it's already deployed by the cluster chart.
- Do not hardcode cilium k8s service port. Use `global.controlPlane.apiServerPort`.
- Fix removing allow-all Cilium network policies
- Change `KubeadmConfig` bootstrap config reference to ensure nodes get rolled when making changes to node specification (requires newer versions of CAPI/CAPA as shown in the original [issue](https://github.com/kubernetes-sigs/cluster-api/issues/8858)). Add machine pool instance warmup setting (5 minutes) to ensure nodes do not get replaced too quickly.
- Run kubeadm after containerd to avoid node startup problems
- Added option to customize app via configmap or secret with `global.apps.{app_name}.extraConfigs`.
- In-line custom values for app moved from `global.apps.{app_name}` to `global.apps.{app_name}.values`.
- Set node pool subnet filters to include avaiability zone.
- Fix error messages if `global.connectivity.baseDomain` is missing
- Fixed issue when deleting node pools that would prevent the deletion, caused by the fact that `MachinePool` and `AWSMachinePool` CRs were annotated with `"helm.sh/resource-policy": keep`.
- Fix containerd config that was breaking in newer flatcar versions.
- The value to configure the control plane load balancer ingress rules is filtered to avoid duplicates and to always contain GiantSwarm VPN IPs.
- Move labels to AWSMachineTemplate manifest to avoid unnecessary rolling/no rolling.
- Make AWS instances names independent of helm label to prevent unnecessary rolling.
- Revert to install default Cilium policies again. Some operators' "allow access to API nodes" `NetworkPolicy`s are not effective and Cilium first needs to be upgraded, including a recent upstream fix to the [known issue](https://github.com/cilium/cilium/issues/20550).
- Accept old service account issuer URI without `https://` prefix as well. This fixes the breaking change introduced in v0.38.4. Existing service account tokens, and the operators/applications using them, will keep working even before the tokens get rotated with the new service account issuer URI. When upgrading, it is recommended to skip earlier releases and immediately jump from v0.38.3 (or older) to _this one_.
- Fix job that removes `HelmReleases` and `HelmCharts`.
- Delete `HelmReleases` and `HelmCharts` clean-up jobs when they are successful.
- Delete all `HelmCharts` on the organization namespace that contain the cluster name on its name.
- Add always-required values to `noProxy` list for aws-cloud-controller-manager-app and aws-ebs-csi-driver-app (only relevant for private clusters with proxy)
- Forbid additional properties under `connectivity.proxy` to avoid typos
- Fix RBAC for `HelmReleases` clean up job.
- Specify `HelmChart` type when patching `HelmCharts` in job that removes finalizers.
- Fix job that removes finalizers by dropping namespace from the `HelmChart` name when using it for patching.
- Remove duplicate label `cluster.x-k8s.io/cluster-name` in bastion MachineDeployment.
- Fix rendering `oidc.pem` by mistake when not specified
- Run machine pools and control plane nodes on private subnets.
- Use region defaulting wherever possible, removing `region` from schema.
- Quote bastion subnet tag filters in order to avoid type conversion errors.
- Quote subnet tag filters in order to avoid type conversion errors.
- Change sed to fix replacement for Cloudfront placeholder.
- Added missing prefixListID for UserManaged network topology
- Add `https://` for IRSA service account issuer.
- Bumped cluster-shared to latest with coredns-adopter apiserver polling
- Handle default values in worker machine pool values
- Immutable AWSMachineTemplate
- Ensure the `KubeadmControlPlane` `.spec.version` value is always prefixed with `v`
- Add the missing `api-audiences` attribute to the `KubeadmControlPlane` template, to fix the use of IRSA service account tokens.
- Re-added Ubuntu 22.04 with correct lookup
- Rolled back to Ubuntu 20.04
- Fix helm context for proxy helper function.
- Improved hash function to hash based on whole `.Spec` rather than just provided values
- AZ list rendering
- Ensure availability zone restrictions are added to the subnet filters
- Fix subnet filter to relevant with `tag:` prefix.
- Limit subnet filter to relevant, cluster owned, subnets
- Ensure worker nodes are only launched in private subnets
- Added the OS version to the imageLookupBaseOS
- Fixed app version label.
- Quoted boolean to a string
- Pod CIDR as array rather than string
- Fix bastion secret.





### Apps



* aws-nth-bundle from 1.2.0 to [1.2.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.0)





### aws-nth-bundle [1.2.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.0)

#### Added

- Send spot instance interruption and instance state change events to SQS queue so that aws-node-termination-handler can react to them

#### Changed

- Add dependency for servicemonitors
- Move to default catalog
- First release




