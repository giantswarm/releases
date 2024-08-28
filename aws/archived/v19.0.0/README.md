# :zap: Giant Swarm Release v19.0.0 for AWS :zap:

> **WARNING:** Please talk to your Account Engineer prior to upgrading. You will be provided with a checklist to follow and validate your clusters.

This release includes upgrades of components and Kubernetes version to 1.24. The upgrade to `v19.0.0` involve two major changes for customers, namely the migration from the AWS VPC CNI to Cilium and the replacement of Kiam with IAM Roles for Service Accounts(IRSA) for authenticating pods against the AWS API.
Next sections are describing important changes we will introduce with the new release, the key benefits, what customers can do to prepare and how to avoid downtime during this crucial upgrade. 

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.3.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#330---2023-05-11) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Cilium

Say goodbye to slow network initialization times and hello to lightning-fast performance with [Cilium](https://github.com/cilium/cilium), our new Kubernetes CNI solution!

### Key Highlights

- Service Mesh integration: `Cilium` is designed to work seamlessly with popular service meshes like `Istio`, `Linkerd`, and `Envoy`. This allows for more advanced networking and security features, such as mTLS encryption and observability.
- `Cilium` uses a virtual network which provides more flexibility, faster network initialization, and more advanced networking features compared to `AWS CNI` affecting the IP addresses.
- Advanced networking features: `Cilium` supports advanced networking features such as load balancing, network segmentation, and eBPF-based packet filtering. These features allow for more granular control over network traffic and improve security. 
- Scalability: `Cilium's` eBPF-based data plane is highly scalable and performs well even at scale. It is also highly efficient, reducing overhead and maximizing performance.
- Improvements in pods time to come up due to cilium endpoint refresh substituting the kubeproxy refresh of iptables.
- More efficient usage of IP space - fully described in `Cilium and IP space` section below.

### What changes with Cilium?

With `Cilium`, you'll no longer be using the `AWS CNI` Pod subnets, so be sure to add custom routes with the `Node subnet(s) CIDR(s)` instead. 

Additionally, while `Cilium's Network Policy` provides powerful security features, support for setting `ipBlock` with `Pod IPs` is not implemented in Cilium, so be sure to inspect your workloads and configure `Network Policies` carefully. The Account Engineers will reach out to you and help to provide the `CiliumNetworkPolicies` before the upgrade in order to have no downtime during the switch. [Cilium-prerequisites](https://github.com/giantswarm/cilium-prerequisites) app that installs the `CiliumNetworkPolicy` CRD is available in the catalog as well as will be installed with GS version `v18.4.0` to provide seemless upgrade experience.

It's important to note that due to changes to `Cluster CR's` during the upgrade process, `GitOps` automation will have to be suspended and any applied changes backported to the repos before resuming. Keep this in mind as you prepare for the upgrade. This needs to be evaluated on a case-by-case basis, since different GitOps implementations might only keep _some_ parts of `Cluster` CRs in Git. Feel free to reach out to your Account Engineer to understand more about these changes.

To ensure a smooth transition to `Cilium`, we've prepared a [comprehensive upgrade process](https://handbook.giantswarm.io/docs/support-and-ops/ops-recipes/upgrade-to-cilium/) that explains every migration step in detail, so you can feel confident in following the process and avoid any potential issues. We have also extended our [documentation](https://docs.giantswarm.io/platform-overview/cluster-management/vintage/aws/#pod-networking) which describes differences between `AWS CNI` and `Cilium`

#### Cilium and AWS Load Balancer Controller

If you are running `aws-load-balancer-controller` inside your clusters for managing [Network Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html) and you did set the annotation `service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip`, you need to change it to `service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: instance`.

For further information, please checkout the [documentation](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.2/guide/service/annotations/#traffic-routing)

#### Cilium and pod CIDR

While switching to Cilium we are forced to change the CIDR used to assign IPs to Pods (192.168.0.0/16 by default).
The process is automated for the vast majority of the clusters, but if you had set up custom networking settings in your cluster the upgrade might be blocked by admission controllers. If that is the case, reach out to your SA and you'll receive guidance how to move on with the upgrade. Same thing applies if you don't want to stick with the default value and prefer to change it.

#### Cilium and improved IP space usage

Migration from AWS CNI to Cilium allowed us to improve the IP space delegation per WC, meaning that starting with AWS v19 release customers will be able to use full range of the CIDR instead of 25%.

#### AWS-CNI in releases prior to v19

By default at Giant Swarm the pod CIDR is set to 10.2.0.0/16 and can be customized in the AWSCluster CR.
The CIDR is added in the VPC CIDR and thus cannot overlap with other CIDRs in the cluster and any other CIDR in peered VPCs.

The pod CIDR space needs to be split into separate, contiguous CIDR space, one for each AZ. Since Giant Swarm supports 3 AZs, we need to divide the range into 4 subranges. In the default scenario, we end up with 4 /18 blocks (16k addresses each). One of the /18 blocks is unused meaning we have 16k (or 1/4th) IP addresses "lost".

#### Cilium in releases v19 and further

In v19 and further releases the default pod CIDR is 192.168.0.0/16. The default can be changed by setting a field in the AWSCluster CR as it was the case so far.

*WARNING*: If pods need to talk directly with private IP addresses reachable through the VPC (for example peered resources) then those target private IP addresses must be on a separate IP space.

This CIDR is used across the whole cluster, regardless of AZs and all addresses from the range can be used. In the default setting, each node gets assigned a /25 subnet for pods running in it. Meaning we can have as many as 65k pods in a cluster (~ 595 nodes in theory) but that can be increased by providing a larger pod cidr space to begin with.

#### Maximum number of pods in a node

With AWS-CNI, IP addresses assigned to pods are actually IP addresses assigned to the node itself.
Depending on the instance type, there is a limit on the number of IP addresses assignable to each instance. This means in practice that clusters using AWS-CNI will have less pods per node in principle. With Cilium, we can use the max number of pods per node as [suggested by k8s](https://kubernetes.io/docs/setup/best-practices/cluster-large/) which is 110.

#### Cilium Troubleshooting
We have included a small [ops-recipe](https://handbook.giantswarm.io/docs/support-and-ops/ops-recipes/cilium-troubleshooting/) for details how you can start troubleshoot Cilium issues.

## IAM roles for service accounts (IRSA)

By switching from `KIAM` to `IAM Roles for Service Accounts (IRSA)`, we're making it easier and more secure for your Kubernetes workloads to interact with AWS services. 

### Key Highlights

- Official AWS way to authenticate pods to AWS API.
- Reduced complexity: IRSA eliminates the need for a separate service like KIAM, streamlining your Kubernetes clusters.
- Regional STS (Security Token Service) rather than using global STS

### What changes with IRSA?

During the upgrade, we are removing `KIAM` as a default app in your workload clusters but it is possible to install it optionally. If you need to keep using KIAM in v19 clusters, please reach out to your SA.

Additionally, we are creating a `Cloudfront Domain Alias` (except China) for each cluster which is used as the [OpenID Connect (OIDC) identity provider](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc.html) to improve predictability and simplify IAM role creation. 

To ensure that your applications can assume the appropriate IAM roles, you need to add the `Cloudfront Domain Alias` to those roles as a [trust entity](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/#aws-release-v19).

We have also adjusted the `external-dns` IRSA trust policy to facilitate externalDNS role being assumed by any Service Account containing `external-dns` phrase to allow multiple app deployments.

To help make your transition to `IRSA` as easy as possible, we've added more context on our [official docs](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/).

## Other release highlights

### 🎣 DNS Node Cache

To improve the DNS performance of your cluster [k8s-dns-node-cache-app](https://github.com/giantswarm/k8s-dns-node-cache-app) will be deployed by default.

#### Key Highlights

- Faster DNS lookups: The app caches DNS lookups on each node, reducing the time it takes to resolve domain names.
- Lower latency: By caching DNS requests locally on each node, the app reduces the need to query external DNS servers, which can improve latency.
- Reducing network traffic: By caching DNS responses locally on each node, the app reduces the need for repeated queries to external DNS servers, which can reduce network traffic.

If you previously deployed `k8s-dns-node-cache-app` through the managed catalog, you can delete the application after the upgrade, as it will be automatically re-installed.

### Prometheus Blackbox Exporter

The [prometheus-blackbox-exporter](https://github.com/giantswarm/prometheus-blackbox-exporter-app) is a new monitoring component installed by default with release `v19`. 

#### Key Highlights

- Flexible monitoring: The blackbox exporter allows users to monitor endpoints from various protocols like HTTP, HTTPS, DNS, TCP, ICMP, and more.
- Real-time monitoring: The exporter provides real-time monitoring of the endpoints and helps detect issues before they turn into major problems.
- Customizable checks: The blackbox exporter can be customized to perform specific checks on the endpoints, which helps in identifying problems quickly.
- Integration with Prometheus: The exporter integrates seamlessly with Prometheus, allowing users to visualize and analyze data collected from the endpoints.

We're aiming to provide a comprehensive blackbox monitoring tool that can validate internal, DNS and external connectivity.

### 🔭 Cilium Hubble

`Cilium` will have [Hubble](https://github.com/cilium/hubble) enabled by default for troubleshooting and observability.

#### Key Highlights

- Provides real-time visibility into network traffic with advanced filtering and aggregation capabilities.
- Helps troubleshoot connectivity issues with its network flow and DNS query analysis features.

#### Caveats and know limitations

- Hubble's UI is not exposed by default, but can be reached using port forwarding. More information regarding the access in the [ops-recipe](https://handbook.giantswarm.io/docs/support-and-ops/ops-recipes/cilium-troubleshooting/#hubble-ui)

## Change details


### app-operator [6.7.0](https://github.com/giantswarm/app-operator/releases/tag/v6.7.0)

#### Changed
- Only include PodSecurityPolicy on clusters with policy/v1beta1 api available.
- Only include PodMonitor on clusters with monitoring.coreos.com/v1 api available.
#### Removed
- Stop pushing to `openstack-app-collection`.



### aws-operator [14.17.1-patch3](https://github.com/giantswarm/aws-operator/releases/tag/v14.17.1-patch3)

#### Added
- Add toleration for new control-plane taint.
#### Fixed
- Ensure `net.ipv4.conf.eth0.rp_filter` is set to `2` if aws-CNI is used.
- Make `routes-fixer` script compatible with alpine.
- Change AWS LB Controller Trust Policy for the new S3 bucket in China clusters.
### Changed
- Change Route53 Trust Policy to allow multiple applications to use the role.
- Update IAM policy for AWS LoadBalancer Controller.

### cluster-operator [5.6.1-patch1](https://github.com/giantswarm/cluster-operator/releases/tag/v5.6.1-patch1)

#### Fixed
- Don't enable Cilium network policies on Azure.
#### Changed
- Patch app operator version on all apps instead of just optional ones.

### k8s-dns-node-cache-app [v2.3.1](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.3.1)

### Changed
- Disable IPV6 queries.
- Remove VPA.
- Remove resource limits.


### aws-cloud-controller-manager [1.24.1-gs9](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.24.1-gs9)

#### Changed
- Adjusted VerticalPodAutoscaler minimum allowed CPU and memory

#### Fixed
- Quote environment variables that contain numeric values, because it's required by kubernetes.



### aws-ebs-csi-driver [2.21.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.21.1)

#### Fixed
- Use `string` type for the proxy parameters on the `values.schema.json` file.



### cert-exporter [2.5.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.5.1)

#### Changed
- Allow requests from the api-server.
- Update icon
- Disable PSPs for k8s 1.25 and newer.



### chart-operator [2.35.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.35.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### cilium [0.10.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.10.0)

#### Changed
- Enable PDB for `cilium-operator`.



### cluster-autoscaler [1.24.0-gs3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.0-gs3)

#### Changed
- Adjusted VerticalPodAutoscaler minimum allowed CPU and memory
- Add 'projected' volumes to the PSP.
- Add new-pod-scale-up-delay variable.
- Disable PSPs for k8s 1.25 and newer.



### coredns [1.17.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.17.1)

#### Added
- Add scaling based on custom metrics ([#209](https://github.com/giantswarm/coredns-app/pull/209)).
#### Changed
- Decouple PDB configuration from deployment updateStrategy ([#208](https://github.com/giantswarm/coredns-app/pull/208)).
- Disable IPV6.



### external-dns [2.37.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.37.1)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### metrics-server [2.2.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.
- Switch to `apiVersion: policy/v1` for PodDisruptionBudget.



### net-exporter [1.15.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.15.0)

#### Changed
- Allow requests from the api-server.
- Disable PSPs for k8s 1.25 and newer.



### node-exporter [1.16.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### vertical-pod-autoscaler [3.5.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.2)

#### Changed
- Remove circleci job for pushing to shared app collection
- Raised resources for updater and recommender.
- Drop all CAPabilities in container SecurityContext for Kyverno Policy compliance
- Set AllowPrivilegeEscalation=false in container SecurityContext for Kyverno Policy compliance


### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap





### observability-bundle [0.5.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.5.1)

#### Changed
- Remove cluster prefix to app name in _helpers.tpl



### prometheus-blackbox-exporter [0.3.2](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.3.2)

#### Added
- Add icon.



### cilium-servicemonitors [0.1.1](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.1)

#### Added
- Add overridability to the servicemonitors relabelings and metric_relabelings sections.



### cert-manager [2.25.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.25.0)

### Changed

- Remove control plane node toleration of CA injector deployment. This caused problems on single control plane node clusters. ([#362](https://github.com/giantswarm/cert-manager-app/pull/362))
- Update container image versions to use [v1.12.4](https://github.com/cert-manager/cert-manager/releases/tag/v1.12.4)



### kubernetes [1.24.13](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.13)

#### Changelog since v1.23.0

#### Major Themes

##### Dockershim Removed from kubelet

After its deprecation in v1.20, the dockershim component has been removed from the kubelet.
From v1.24 onwards, you will need to either use one of the other [supported runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/) (such as containerd or CRI-O)
or use cri-dockerd if you are relying on Docker Engine as your container runtime.
For more information about ensuring your cluster is ready for this removal, please
see [this guide](https://kubernetes.io/blog/2022/03/31/ready-for-dockershim-removal/).

##### Beta APIs Off by Default

[New beta APIs will not be enabled in clusters by default](https://github.com/kubernetes/enhancements/issues/3136).
Existing beta APIs and new versions of existing beta APIs, will continue to be enabled by default.

##### Signing Release Artifacts

Release artifacts are [signed](https://github.com/kubernetes/enhancements/issues/3031) using [cosign](https://github.com/sigstore/cosign)
signatures
and there is experimental support for [verifying image signatures](https://kubernetes.io/docs/tasks/administer-cluster/verify-signed-images/).
Signing and verification of release artifacts is part of [increasing software supply chain security for the Kubernetes release process](https://github.com/kubernetes/enhancements/issues/3027).

##### OpenAPI v3

Kubernetes 1.24 offers beta support for publishing its APIs in the [OpenAPI v3 format](https://github.com/kubernetes/enhancements/issues/2896).

##### Storage Capacity and Volume Expansion Are Generally Available

[Storage capacity tracking](https://github.com/kubernetes/enhancements/issues/1472)
supports exposing currently available storage capacity via [CSIStorageCapacity objects](https://kubernetes.io/docs/concepts/storage/storage-capacity/#api)
and enhances scheduling of pods that use CSI volumes with late binding.

[Volume expansion](https://github.com/kubernetes/enhancements/issues/284) adds support
for resizing existing persistent volumes.

##### NonPreemptingPriority to Stable

This feature adds [a new option to PriorityClasses](https://github.com/kubernetes/enhancements/issues/902),
which can enable or disable pod preemption.

##### Storage Plugin Migration

There is work under way to [migrate the internals of in-tree storage plugins](https://github.com/kubernetes/enhancements/issues/625) to call out to CSI Plugins,
while maintaining the original API.
The [Azure Disk](https://github.com/kubernetes/enhancements/issues/1490)
and [OpenStack Cinder](https://github.com/kubernetes/enhancements/issues/1489) plugins
have both been migrated.

##### gRPC Probes Graduate to Beta

With Kubernetes 1.24, the [gRPC probes functionality](https://github.com/kubernetes/enhancements/issues/2727)
has entered beta and is available by default. You can now [configure startup, liveness, and readiness probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#configure-probes) for your gRPC app
natively within Kubernetes, without exposing an HTTP endpoint or
using an extra executable.

##### Kubelet Credential Provider Graduates to Beta

Originally released as Alpha in Kubernetes 1.20, the kubelet's support for
[image credential providers](https://kubernetes.io/docs/tasks/kubelet-credential-provider/kubelet-credential-provider/)
has now graduated to Beta.
This allows the kubelet to dynamically retrieve credentials for a container image registry
using exec plugins, rather than storing credentials on the node's filesystem.

##### Contextual Logging in Alpha

Kubernetes 1.24 has introduced [contextual logging](https://github.com/kubernetes/enhancements/issues/3077)
that enables the caller of a function to control all aspects of logging (output formatting, verbosity, additional values and names).

##### Avoiding Collisions in IP allocation to Services

Kubernetes 1.24 introduced a new opt-in feature that allows you to
[soft-reserve a range for static IP address assignments](https://kubernetes.io/docs/concepts/services-networking/service/#service-ip-static-sub-range)
to Services.
With the manual enablement of this feature, the cluster will prefer automatic assignment from
the pool of Service IP addresses thereby reducing the risk of collision.

A Service `ClusterIP` can be assigned:

* dynamically, which means the cluster will automatically pick a free IP within the configured Service IP range.
* statically, which means the user will set one IP within the configured Service IP range.

Service `ClusterIP` are unique, hence, trying to create a Service with a `ClusterIP` that has already been allocated will return an error.

#### Urgent Upgrade Notes

##### (No, really, you MUST read this before you upgrade)

- Docker runtime support using dockershim in the kubelet is now completely removed in 1.24. The kubelet used to have a module called dockershim, which implements CRI support for Docker, and it has seen maintenance issues in the Kubernetes community. From 1.24 onwards, please move to a container runtime that is a full-fledged implementation of CRI (v1alpha1 or v1 compliant) as they become available. ([#97252](https://github.com/kubernetes/kubernetes/pull/97252), [@dims](https://github.com/dims))
- Fixed bug with leads to Node goes `Not-ready` state when credentials for vCenter stored in a secret and Zones feature is in use. Zone labels setup moved to KCM component, kubelet skips this step during startup in such case. If credentials stored in cloud-provider config file as plaintext current behaviour does not change and no action required. For proper functioning `kube-system:vsphere-legacy-cloud-provider` should be allowed to update node object if vCenter credentials stored in secret and Zone feature used. ([#101028](https://github.com/kubernetes/kubernetes/pull/101028), [@lobziik](https://github.com/lobziik))
- The `LegacyServiceAccountTokenNoAutoGeneration` feature gate is beta, and enabled by default. When enabled, Secret API objects containing service account tokens are no longer auto-generated for every ServiceAccount. Use the [TokenRequest](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/) API to acquire service account tokens, or if a non-expiring token is required, create a Secret API object for the token controller to populate with a service account token by following this [guide](https://kubernetes.io/docs/concepts/configuration/secret/#service-account-token-secrets). ([#108309](https://github.com/kubernetes/kubernetes/pull/108309), [@zshihang](https://github.com/zshihang))
- The calculations for Pod topology spread skew now exclude nodes that
  don't match the node affinity/selector. This may lead to unschedulable pods if you previously had pods
  matching the spreading selector on those excluded nodes (not matching the node affinity/selector),
  especially when the `topologyKey` is not node-level. Revisit the node affinity and/or pod selector in the
  topology spread constraints to avoid this scenario. ([#107009](https://github.com/kubernetes/kubernetes/pull/107009), [@kerthcet](https://github.com/kerthcet))
- Remove the deprecated flag `--experimental-check-node-capabilities-before-mount`. With CSI now GA, there is a better alternative. Remove any use of  `--experimental-check-node-capabilities-before-mount` from your kubelet scripts or manifests. ([#104732](https://github.com/kubernetes/kubernetes/pull/104732), [@mengjiao-liu](https://github.com/mengjiao-liu))
- `kubeadm.k8s.io/v1beta2` has been deprecated and will be removed in a future release, possibly in 3 releases (one year). You should start using `kubeadm.k8s.io/v1beta3` for new clusters. To migrate your old configuration files on disk you can use the `kubeadm config migrate` command. ([#107013](https://github.com/kubernetes/kubernetes/pull/107013), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: default the kubeadm configuration to the containerd socket (Unix: `unix:///var/run/containerd/containerd.sock`, Windows: `npipe:////./pipe/containerd-containerd`) instead of the one for Docker. If the `Init|JoinConfiguration.nodeRegistration.criSocket` field is empty during cluster creation and multiple sockets are found on the host always throw an error and ask the user to specify which one to use by setting the value in the field. Make sure you update any kubeadm configuration files on disk, to not include the dockershim socket unless you are still using kubelet version < 1.24 with kubeadm >= 1.24. Remove the DockerValidor and ServiceCheck for the `docker` service from kubeadm preflight. Docker is no longer special cased during host validation and ideally this task should be done in the now external cri-dockerd project where the importance of the compatibility matters. Use `crictl` for all communication with CRI sockets for actions like pulling images and obtaining a list of running containers instead of using the docker CLI in the case of Docker. ([#107317](https://github.com/kubernetes/kubernetes/pull/107317), [@neolit123](https://github.com/neolit123))
- The feature gate was mentioned as `csiMigrationRBD` where it should have been `CSIMigrationRBD` to be in parity with other migration plugins. This release correct the same and keep it as `CSIMigrationRBD`.
    users who have configured this feature gate as `csiMigrationRBD` has to reconfigure the same to `CSIMigrationRBD` from this release. ([#107554](https://github.com/kubernetes/kubernetes/pull/107554), [@humblec](https://github.com/humblec))
- The experimental dynamic log sanitization feature has been deprecated and removed in the 1.24 release. The feature is no longer available for use. ([#107207](https://github.com/kubernetes/kubernetes/pull/107207), [@ehashman](https://github.com/ehashman))
- Kubeadm: apply `second stage` of the plan to migrate kubeadm away from the usage of the word `master` in labels and taints. For new clusters, the label `node-role.kubernetes.io/master` will no longer be added to control plane nodes, only the label `node-role.kubernetes.io/control-plane` will be added. For clusters that are being upgraded to 1.24 with `kubeadm upgrade apply`, the command will remove the label `node-role.kubernetes.io/master` from existing control plane nodes. For new clusters, both the old taint `node-role.kubernetes.io/master:NoSchedule` and new taint `node-role.kubernetes.io/control-plane:NoSchedule` will be added to control plane nodes. In release 1.20 (`first stage`), a release note instructed to preemptively tolerate the new taint. For clusters that are being upgraded to 1.24 with `kubeadm upgrade apply`, the command will add the new taint `node-role.kubernetes.io/control-plane:NoSchedule` to existing control plane nodes. Please adapt your infrastructure to these changes. In 1.25 the old taint `node-role.kubernetes.io/master:NoSchedule` will be removed. ([#107533](https://github.com/kubernetes/kubernetes/pull/107533), [@neolit123](https://github.com/neolit123))

#### Changes by Kind

##### Deprecation

- Deprecated `Service.Spec.LoadBalancerIP`. This field was under-specified and its meaning varies across implementations.  As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available.  This field may be removed in a future API version. ([#107235](https://github.com/kubernetes/kubernetes/pull/107235), [@uablrek](https://github.com/uablrek))
- Kube-apiserver: the `--master-count` flag and `--endpoint-reconciler-type=master-count` reconciler are deprecated in favor of the lease reconciler ([#108062](https://github.com/kubernetes/kubernetes/pull/108062), [@aojea](https://github.com/aojea))
- Kube-apiserver: the insecure address flags `--address`, `--insecure-bind-address`, `--port` and `--insecure-port` (inert since 1.20) are removed ([#106859](https://github.com/kubernetes/kubernetes/pull/106859), [@knight42](https://github.com/knight42))
- Kubeadm: graduated the `UnversionedKubeletConfigMap` feature gate to Beta and enabled the feature by default. This implies that 1) for new clusters kubeadm will start using the `kube-system/kubelet-config` naming scheme for the kubelet ConfigMap and RBAC rules, instead of the legacy `kubelet-config-x.yy` naming. 2) during upgrade, kubeadm will only write the new scheme ConfigMap and RBAC objects. To disable the feature you can pass `UnversionedKubeletConfigMap: false` in the kubeadm config for new clusters. For upgrade on existing clusters you can also override the behavior by patching the ClusterConfiguration object in `kube-system/kubeadm-config`. More details in the associated KEP. ([#108027](https://github.com/kubernetes/kubernetes/pull/108027), [@neolit123](https://github.com/neolit123))
- Remove `tolerate-unready-endpoints` annotation in Service deprecated from 1.11, use `Service.spec.publishNotReadyAddresses` instead. ([#108020](https://github.com/kubernetes/kubernetes/pull/108020), [@tossmilestone](https://github.com/tossmilestone))
- Remove deprecated feature gates `ValidateProxyRedirects` and `StreamingProxyRedirects` ([#106830](https://github.com/kubernetes/kubernetes/pull/106830), [@pacoxu](https://github.com/pacoxu))
- Remove insecure serving configuration from cloud-provider package, which is consumed by cloud-controller-managers. ([#108953](https://github.com/kubernetes/kubernetes/pull/108953), [@nckturner](https://github.com/nckturner))
- The `--pod-infra-container-image` kubelet flag is deprecated and will be removed in future releases ([#108045](https://github.com/kubernetes/kubernetes/pull/108045), [@hakman](https://github.com/hakman))
- The `client.authentication.k8s.io/v1alpha1` ExecCredential has been removed. If you are using a client-go credential plugin that relies on the v1alpha1 API please contact the distributor of your plugin for instructions on how to migrate to the v1 API. ([#108616](https://github.com/kubernetes/kubernetes/pull/108616), [@margocrawf](https://github.com/margocrawf))
- The `node.k8s.io/v1alpha1` RuntimeClass API is no longer served. Use the `node.k8s.io/v1` API version, available since v1.20 ([#103061](https://github.com/kubernetes/kubernetes/pull/103061), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- The cluster addon for dashboard was removed. To install dashboard, see [here](https://github.com/kubernetes/dashboard/blob/master/docs/user/README.md). ([#107481](https://github.com/kubernetes/kubernetes/pull/107481), [@shu-mutou](https://github.com/shu-mutou))
- The in-tree Azure plugin has been deprecated. The Azure kubelogin plugin serves as an out-of-tree replacement via the kubectl/client-go credential plugin mechanism.  Users will now see a warning in the logs regarding this deprecation. ([#107904](https://github.com/kubernetes/kubernetes/pull/107904), [@sabbey37](https://github.com/sabbey37))
- The insecure address flags `--address` and `--port` in kube-controller-manager have had no effect since v1.20 and are removed in v1.24. ([#106860](https://github.com/kubernetes/kubernetes/pull/106860), [@knight42](https://github.com/knight42))
- The metadata.clusterName field is deprecated. This field has always been unwritable and always blank, but its presence is confusing, so we will remove it next release. Out of an abundance of caution, this release we have merely changed the name in the go struct to ensure any accidental client uses are found before complete removal. ([#108717](https://github.com/kubernetes/kubernetes/pull/108717), [@lavalamp](https://github.com/lavalamp))
- VSphere releases less than 7.0u2 are deprecated as of v1.24. Please consider upgrading vSphere to 7.0u2 or above. vSphere CSI Driver requires minimum vSphere 7.0u2.

  General Support for vSphere 6.7 will end on October 15, 2022. vSphere 6.7 Update 3 is deprecated in Kubernetes v1.24.  Customers are recommended to upgrade vSphere (both ESXi and vCenter) to 7.0u2 or above.  vSphere CSI Driver 2.2.3 and higher supports CSI Migration.

  Support for these deprecations will be available till October 15, 2022. ([#109089](https://github.com/kubernetes/kubernetes/pull/109089), [@deepakkinni](https://github.com/deepakkinni))

##### API Change

- Kubernetes 1.24 is now built with go1.19.4 ([#113956](https://github.com/kubernetes/kubernetes/pull/113956), [@liggitt](https://github.com/liggitt)) [SIG Apps, Architecture, Auth, Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Release, Scheduling, Storage and Testing]
- Protobuf serialization of metav1.MicroTime timestamps (used in `Lease` and `Event` API objects) has been corrected to truncate to microsecond precision, to match the documented behavior and JSON/YAML serialization. Any existing persisted data is truncated to microsecond when read from etcd. ([#111936](https://github.com/kubernetes/kubernetes/pull/111936), [@haoruan](https://github.com/haoruan)) [SIG API Machinery]
- Revert regression that prevented client-go latency metrics to be reported with a template URL to avoid label cardinality. ([#112056](https://github.com/kubernetes/kubernetes/pull/112056), [@aanm](https://github.com/aanm)) [SIG API Machinery]
- Add 2 new options for kube-proxy running in winkernel mode. `--forward-healthcheck-vip`, if specified as true, health check traffic whose destination is service VIP will be forwarded to kube-proxy's healthcheck service. `--root-hnsendpoint-name` specifies the name of the hns endpoint for the root network namespace. This option enables the pass-through load balancers like Google's GCLB to correctly health check the backend services. Without this change, the health check packets is dropped, and Windows node will be considered to be unhealthy by those load balancers. ([#99287](https://github.com/kubernetes/kubernetes/pull/99287), [@anfernee](https://github.com/anfernee))
- Added CEL runtime cost calculation into CustomerResource validation. CustomerResource validation will fail if runtime cost exceeds the budget. ([#108482](https://github.com/kubernetes/kubernetes/pull/108482), [@cici37](https://github.com/cici37))
- Added a new metric `webhook_fail_open_count` to monitor webhooks that fail to open. ([#107171](https://github.com/kubernetes/kubernetes/pull/107171), [@ltagliamonte-dd](https://github.com/ltagliamonte-dd))
- Adds a new Status subresource in Network Policy objects ([#107963](https://github.com/kubernetes/kubernetes/pull/107963), [@rikatz](https://github.com/rikatz))
- Adds support for `InterfaceNamePrefix` and `BridgeInterface` as arguments to `--detect-local-mode` option and also introduces a new optional `--pod-interface-name-prefix` and `--pod-bridge-interface` flags to kube-proxy. ([#95400](https://github.com/kubernetes/kubernetes/pull/95400), [@tssurya](https://github.com/tssurya))
- CEL CRD validation expressions may now reference existing object state using the identifier `oldSelf`. ([#108073](https://github.com/kubernetes/kubernetes/pull/108073), [@benluddy](https://github.com/benluddy))
- CRD deep copies should no longer contain shallow copies of `JSONSchemaProps.XValidations`. ([#107956](https://github.com/kubernetes/kubernetes/pull/107956), [@benluddy](https://github.com/benluddy))
- CRD writes will generate validation errors if a CEL validation rule references the identifier `oldSelf` on a part of the schema that does not support it. ([#108013](https://github.com/kubernetes/kubernetes/pull/108013), [@benluddy](https://github.com/benluddy))
- CSIStorageCapacity.storage.k8s.io: The v1beta1 version of this API is deprecated in favor of v1, and will be removed in v1.27. If a CSI driver supports storage capacity tracking, then it must get deployed with a release of external-provisioner that supports the v1 API. ([#108445](https://github.com/kubernetes/kubernetes/pull/108445), [@pohly](https://github.com/pohly))
- Custom resource requests with `fieldValidation=Strict` consistently require `apiVersion` and `kind`, matching non-strict requests ([#109019](https://github.com/kubernetes/kubernetes/pull/109019), [@liggitt](https://github.com/liggitt))
- Feature of `DefaultPodTopologySpread` is graduated to GA ([#108278](https://github.com/kubernetes/kubernetes/pull/108278), [@kerthcet](https://github.com/kerthcet))
- Feature of `NonPreemptingPriority` is graduated to GA ([#107432](https://github.com/kubernetes/kubernetes/pull/107432), [@denkensk](https://github.com/denkensk))
- Feature of `PodOverhead` is graduated to GA ([#108441](https://github.com/kubernetes/kubernetes/pull/108441), [@pacoxu](https://github.com/pacoxu))
- Fixed OpenAPI serialization of the x-kubernetes-validations field ([#107970](https://github.com/kubernetes/kubernetes/pull/107970), [@liggitt](https://github.com/liggitt))
- Fixed failed flushing logs in defer function when kubelet cmd exit 1. ([#104774](https://github.com/kubernetes/kubernetes/pull/104774), [@kerthcet](https://github.com/kerthcet))
- Fixes a regression in v1beta1 PodDisruptionBudget handling of `strategic merge patch`-type API requests for the `selector` field. Prior to 1.21, these requests would merge `matchLabels` content and replace `matchExpressions` content. In 1.21, patch requests touching the `selector` field started replacing the entire selector. This is consistent with server-side apply and the v1 PodDisruptionBudget behavior, but should not have been changed for v1beta1. ([#108138](https://github.com/kubernetes/kubernetes/pull/108138), [@liggitt](https://github.com/liggitt))
- Improve kubectl's user help commands readability ([#104736](https://github.com/kubernetes/kubernetes/pull/104736), [@lauchokyip](https://github.com/lauchokyip))
- Indexed Jobs graduated to stable. ([#107395](https://github.com/kubernetes/kubernetes/pull/107395), [@alculquicondor](https://github.com/alculquicondor))
- Introduce a v1alpha1 networking API for ClusterCIDRConfig ([#108290](https://github.com/kubernetes/kubernetes/pull/108290), [@sarveshr7](https://github.com/sarveshr7))
- Introduction of a new "sync_proxy_rules_no_local_endpoints_total" proxy metric. This metric represents the number of services with no internal endpoints. The "traffic_policy" label will contain both "internal" or "external". ([#108930](https://github.com/kubernetes/kubernetes/pull/108930), [@MaxRenaud](https://github.com/MaxRenaud))
- JobReadyPods graduates to Beta and it's enabled by default. ([#107476](https://github.com/kubernetes/kubernetes/pull/107476), [@alculquicondor](https://github.com/alculquicondor))
- Kube-apiserver: `--audit-log-version` and `--audit-webhook-version` now only support the default value of `audit.k8s.io/v1`. The v1alpha1 and v1beta1 audit log versions, deprecated since 1.13, have been removed. ([#108092](https://github.com/kubernetes/kubernetes/pull/108092), [@carlory](https://github.com/carlory))
- Kube-apiserver: the `metadata.selfLink` field can no longer be populated by kube-apiserver; it was deprecated in 1.16 and has not been populated by default since 1.20+. ([#107527](https://github.com/kubernetes/kubernetes/pull/107527), [@wojtek-t](https://github.com/wojtek-t))
- Kubelet external Credential Provider feature is moved to Beta. Credential Provider Plugin and Credential Provider Config API's updated from v1alpha1 to v1beta1 with no API changes. ([#108847](https://github.com/kubernetes/kubernetes/pull/108847), [@adisky](https://github.com/adisky))
- Make STS available replicas optional again. ([#109241](https://github.com/kubernetes/kubernetes/pull/109241), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- MaxUnavailable for StatefulSets, allows faster RollingUpdate by taking down more than 1 pod at a time. The number of pods you want to take down during a RollingUpdate is configurable using maxUnavailable parameter. ([#82162](https://github.com/kubernetes/kubernetes/pull/82162), [@krmayankk](https://github.com/krmayankk))
- Non-graceful node shutdown handling is enabled for stateful workload failovers ([#108486](https://github.com/kubernetes/kubernetes/pull/108486), [@sonasingh46](https://github.com/sonasingh46))
- Omit enum declarations from the static openapi file captured at https://git.k8s.io/kubernetes/api/openapi-spec. This file is used to generate API clients, and use of enums in those generated clients (rather than strings) can break forward compatibility with additional future values in those fields. See https://issue.k8s.io/109177 for details. ([#109178](https://github.com/kubernetes/kubernetes/pull/109178), [@liggitt](https://github.com/liggitt))
- OpenAPI V3 is turned on by default ([#109031](https://github.com/kubernetes/kubernetes/pull/109031), [@Jefftree](https://github.com/Jefftree))
- Pod affinity namespace selector and cross-namespace quota graduated to GA. The feature gate `PodAffinityNamespaceSelector` is locked and will be removed in 1.26. ([#108136](https://github.com/kubernetes/kubernetes/pull/108136), [@ahg-g](https://github.com/ahg-g))
- Promote IdentifyPodOS feature to beta. ([#107859](https://github.com/kubernetes/kubernetes/pull/107859), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Remove a v1alpha1 networking API for ClusterCIDRConfig ([#109436](https://github.com/kubernetes/kubernetes/pull/109436), [@JamesLaverack](https://github.com/JamesLaverack))
- Renamed metrics `evictions_number` to `evictions_total` and mark it as stable. The original `evictions_number` metrics name is marked as "Deprecated" and has been removed in kubernetes 1.23 . ([#106366](https://github.com/kubernetes/kubernetes/pull/106366), [@cyclinder](https://github.com/cyclinder))
- Skip x-kubernetes-validations rules if having fundamental error against the OpenAPIv3 schema. ([#108859](https://github.com/kubernetes/kubernetes/pull/108859), [@cici37](https://github.com/cici37))
- Support for gRPC probes is now in beta. GRPCContainerProbe feature gate is enabled by default. ([#108522](https://github.com/kubernetes/kubernetes/pull/108522), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- Suspend job to GA. The feature gate `SuspendJob` is locked and will be removed in 1.26. ([#108129](https://github.com/kubernetes/kubernetes/pull/108129), [@ahg-g](https://github.com/ahg-g))
- The AnyVolumeDataSource feature is now beta, and the feature gate is enabled by default. In order to provide user feedback on PVCs with data sources, deployers must install the VolumePopulators CRD and the data-source-validator controller. ([#108736](https://github.com/kubernetes/kubernetes/pull/108736), [@bswartz](https://github.com/bswartz))
- The CertificateSigningRequest `spec.expirationSeconds` API field has graduated to GA. The `CSRDuration` feature gate for the field is now unconditionally enabled and will be removed in 1.26. ([#108782](https://github.com/kubernetes/kubernetes/pull/108782), [@cfryanr](https://github.com/cfryanr))
- The `ServerSideFieldValidation` feature has graduated to beta and is now enabled by default. Kubectl 1.24 and newer will use server-side validation instead of client-side validation when writing to API servers with the feature enabled. ([#108889](https://github.com/kubernetes/kubernetes/pull/108889), [@kevindelgado](https://github.com/kevindelgado))
- The `ServiceLBNodePortControl` feature has graduated to GA. The feature gate will be removed in 1.26. ([#107027](https://github.com/kubernetes/kubernetes/pull/107027), [@uablrek](https://github.com/uablrek))
- The deprecated kube-controller-manager flag '--deployment-controller-sync-period' has been removed, it is not used by the deployment controller. ([#107178](https://github.com/kubernetes/kubernetes/pull/107178), [@SataQiu](https://github.com/SataQiu))
- The feature `DynamicKubeletConfig` has been removed from the kubelet. ([#106932](https://github.com/kubernetes/kubernetes/pull/106932), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- The infrastructure for contextual logging is complete (feature gate implemented, JSON backend ready). ([#108995](https://github.com/kubernetes/kubernetes/pull/108995), [@pohly](https://github.com/pohly))
- This adds an optional `timeZone` field as part of the CronJob spec to support running cron jobs in a specific time zone. ([#108032](https://github.com/kubernetes/kubernetes/pull/108032), [@deejross](https://github.com/deejross))
- Updated the default API priority-and-fairness config to avoid endpoint/configmaps operations from controller-manager to all match leader-election priority level. ([#106725](https://github.com/kubernetes/kubernetes/pull/106725), [@wojtek-t](https://github.com/wojtek-t))
- `topologySpreadConstraints` includes `minDomains` field to limit the minimum number of topology domains. ([#107674](https://github.com/kubernetes/kubernetes/pull/107674), [@sanposhiho](https://github.com/sanposhiho))

##### Feature

- Kubernetes is now built with Go 1.19.8 ([#117132](https://github.com/kubernetes/kubernetes/pull/117132), [@xmudrii](https://github.com/xmudrii)) [SIG Release and Testing]
- Kubelet TCP and HTTP probes are more effective using networking resources: conntrack entries, sockets, ... 
  This is achieved by reducing the TIME-WAIT state of the connection to 1 second, instead of the defaults 60 seconds. This allows kubelet to free the socket, and free conntrack entry and ephemeral port associated. ([#115143](https://github.com/kubernetes/kubernetes/pull/115143), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Kubeadm: use the image registry registry.k8s.io instead of k8s.gcr.io for new clusters. During upgrade, migrate users to registry.k8s.io if they were using the default of k8s.gcr.io. ([#113395](https://github.com/kubernetes/kubernetes/pull/113395), [@neolit123](https://github.com/neolit123)) [SIG Cloud Provider and Cluster Lifecycle]
- Kubernetes is now built with Go 1.19.5 ([#115012](https://github.com/kubernetes/kubernetes/pull/115012), [@cpanato](https://github.com/cpanato)) [SIG Release and Testing]- A new Priority and Fairness metric 'apiserver_flowcontrol_work_estimate_seats_samples' has been added that tracks the estimated seats associated with a request. ([#106628](https://github.com/kubernetes/kubernetes/pull/106628), [@tkashem](https://github.com/tkashem))
- Add a deprecated cmd flag for the time interval between flushing pods from unschedulable queue to active queue or backoff queue. ([#108017](https://github.com/kubernetes/kubernetes/pull/108017), [@denkensk](https://github.com/denkensk))
- Add one metrics(`kubelet_volume_stats_health_abnormal`) of volume health state to kubelet ([#105585](https://github.com/kubernetes/kubernetes/pull/105585), [@fengzixu](https://github.com/fengzixu))
- Add the metric `container_oom_events_total` to kubelet's cAdvisor metric endpoint. ([#108004](https://github.com/kubernetes/kubernetes/pull/108004), [@jonkerj](https://github.com/jonkerj))
- Added `SetTransform` to `SharedInformer` to allow users to transform objects before they are stored. ([#107507](https://github.com/kubernetes/kubernetes/pull/107507), [@alexzielenski](https://github.com/alexzielenski))
- Added a `proxy-url` flag into `kubectl config set-cluster`. ([#105566](https://github.com/kubernetes/kubernetes/pull/105566), [@ardaguclu](https://github.com/ardaguclu))
- Added a metric for measuring end-to-end volume mount timing. ([#107006](https://github.com/kubernetes/kubernetes/pull/107006), [@gnufied](https://github.com/gnufied))
- Added a new Priority and Fairness metric `apiserver_flowcontrol_request_dispatch_no_accommodation_total` to track the number of times a request dispatch attempt results in a no-accommodation status due to lack of available seats. ([#106629](https://github.com/kubernetes/kubernetes/pull/106629), [@tkashem](https://github.com/tkashem))
- Added a path `/header?key=` to `agnhost netexec` allowing one to view what the header value is of the incoming request.

  Ex:
  ```
  $ curl -H "X-Forwarded-For: something" 172.17.0.2:8080/header?key=X-Forwarded-For
  something
  ```
   ([#107796](https://github.com/kubernetes/kubernetes/pull/107796), [@alexanderConstantinescu](https://github.com/alexanderConstantinescu))
- Added completion for `kubectl config set-context`. ([#106739](https://github.com/kubernetes/kubernetes/pull/106739), [@kebe7jun](https://github.com/kebe7jun))
- Added field `add_ambient_capabilities` to the Capabilities message in the CRI-API. ([#104620](https://github.com/kubernetes/kubernetes/pull/104620), [@vinayakankugoyal](https://github.com/vinayakankugoyal))
- Added label selector flag to all `kubectl rollout` commands. ([#99758](https://github.com/kubernetes/kubernetes/pull/99758), [@aramperes](https://github.com/aramperes))
- Added more message for no PodSandbox container. ([#107116](https://github.com/kubernetes/kubernetes/pull/107116), [@yxxhero](https://github.com/yxxhero))
- Added prune flag into `diff` command to simulate `apply --prune`. ([#105164](https://github.com/kubernetes/kubernetes/pull/105164), [@ardaguclu](https://github.com/ardaguclu))
- Added support for `btrfs` resizing ([#108561](https://github.com/kubernetes/kubernetes/pull/108561), [@RomanBednar](https://github.com/RomanBednar))
- Added support for kubectl commands (`kubectl exec` and `kubectl port-forward`) via a SOCKS5 proxy. ([#105632](https://github.com/kubernetes/kubernetes/pull/105632), [@xens](https://github.com/xens))
- Adds `OpenAPIV3SchemaInterface` to `DiscoveryClient` and its variants for fetching OpenAPI v3 schema documents. ([#108992](https://github.com/kubernetes/kubernetes/pull/108992), [@alexzielenski](https://github.com/alexzielenski))
- Allow kubectl to manage resources by filename patterns without the shell expanding it first ([#102265](https://github.com/kubernetes/kubernetes/pull/102265), [@danielrodriguez](https://github.com/danielrodriguez))
- An alpha flag `--subresource` is added to get, patch, edit replace kubectl commands to fetch and update status and scale subresources. ([#99556](https://github.com/kubernetes/kubernetes/pull/99556), [@nikhita](https://github.com/nikhita))
- Apiextensions_openapi_v3_regeneration_count metric (alpha) will be emitted for OpenAPI V3. ([#109128](https://github.com/kubernetes/kubernetes/pull/109128), [@Jefftree](https://github.com/Jefftree))
- Apply ProxyTerminatingEndpoints to all traffic policies (external, internal, cluster, local). ([#108691](https://github.com/kubernetes/kubernetes/pull/108691), [@andrewsykim](https://github.com/andrewsykim))
- CEL regex patterns in x-kubernetes-valiation rules are compiled when CRDs are created/updated if the pattern is provided as a string constant in the expression. Any regex compile errors are reported as a CRD create/update validation error. ([#108617](https://github.com/kubernetes/kubernetes/pull/108617), [@jpbetz](https://github.com/jpbetz))
- CRD `x-kubernetes-validations` rules now support the CEL functions: `isSorted`, `sum`, `min`, `max`, `indexOf`, `lastIndexOf`, `find` and `findAll`. ([#108312](https://github.com/kubernetes/kubernetes/pull/108312), [@jpbetz](https://github.com/jpbetz))
- Changes the kubectl `--validate` flag from a bool to a string that accepts the values {true, strict, warn, false, ignore}
  - true/strict - perform validation and error the request on any invalid fields in the ojbect. It will attempt to perform server-side validation if it is enabled on the apiserver, otherwise it will fall back to client-side validation.
  - warn - perform server-side validation and warn on any invalid fields (but ultimately let the request succeed by dropping any invalid fields from the object). If validation is not available on the server, perform no validation.
  - false/ignore - perform no validation, silently dropping invalid fields from the object. ([#108350](https://github.com/kubernetes/kubernetes/pull/108350), [@kevindelgado](https://github.com/kevindelgado))
- Client-go metrics: change bucket distribution for `rest_client_request_duration_seconds` and `rest_client_rate_limiter_duration_seconds` from [0.001, 0.002, 0.004, 0.008, 0.016, 0.032, 0.064, 0.128, 0.256, 0.512] to [0.005, 0.025, 0.1, 0.25, 0.5, 1.0, 2.0, 4.0, 8.0, 15.0, 30.0, 60.0}] ([#106911](https://github.com/kubernetes/kubernetes/pull/106911), [@aojea](https://github.com/aojea))
- Client-go: add new histogram metric to record the size of the requests and responses. ([#108296](https://github.com/kubernetes/kubernetes/pull/108296), [@aojea](https://github.com/aojea))
- CycleState is now optimized for "write once and read many times". ([#108724](https://github.com/kubernetes/kubernetes/pull/108724), [@sanposhiho](https://github.com/sanposhiho))
- Enabled beta feature HonorPVReclaimPolicy by default. ([#109035](https://github.com/kubernetes/kubernetes/pull/109035), [@deepakkinni](https://github.com/deepakkinni))
- Env var for additional cli flags used in the csi-proxy binary when a Windows nodepool is created with `kube-up.sh` ([#107806](https://github.com/kubernetes/kubernetes/pull/107806), [@mauriciopoppe](https://github.com/mauriciopoppe))
- Feature of `PreferNominatedNode` is graduated to GA. ([#106619](https://github.com/kubernetes/kubernetes/pull/106619), [@chendave](https://github.com/chendave))
- In text format, log messages that previously used quoting to prevent multi-line output (for example, text="some \"quotation\", a\nline break") will now be printed with more readable multi-line output without the escape sequences. ([#107103](https://github.com/kubernetes/kubernetes/pull/107103), [@pohly](https://github.com/pohly))
- Increase default value of discovery cache TTL for kubectl to 6 hours. ([#107141](https://github.com/kubernetes/kubernetes/pull/107141), [@mk46](https://github.com/mk46))
- Introduce policy to allow the HPA to consume the `external.metrics.k8s.io` API group. ([#104244](https://github.com/kubernetes/kubernetes/pull/104244), [@dgrisonnet](https://github.com/dgrisonnet))
- Kube-apiserver: Subresources such as `status` and `scale` now support tabular output content types. ([#103516](https://github.com/kubernetes/kubernetes/pull/103516), [@ykakarap](https://github.com/ykakarap))
- Kube-apiserver: when merging lists, Server Side Apply now prefers the order of the submitted request instead of the existing persisted object. ([#107565](https://github.com/kubernetes/kubernetes/pull/107565), [@jiahuif](https://github.com/jiahuif))
- Kubeadm: added support for dry running `kubeadm reset`. The new flag `kubeadm reset --dry-run` is similar to the existing flag for `kubeadm init/join/upgrade` and allows you to see what changes would be applied. ([#107512](https://github.com/kubernetes/kubernetes/pull/107512), [@SataQiu](https://github.com/SataQiu))
- Kubeadm: added the flag `--experimental-initial-corrupt-check` to etcd static Pod manifests to ensure etcd member data consistency ([#109074](https://github.com/kubernetes/kubernetes/pull/109074), [@neolit123](https://github.com/neolit123))
- Kubeadm: better surface errors during `kubeadm upgrade` when waiting for the kubelet to restart static pods on control plane nodes ([#108315](https://github.com/kubernetes/kubernetes/pull/108315), [@Monokaix](https://github.com/Monokaix))
- Kubeadm: improve the strict parsing of user YAML/JSON configuration files. Next to printing warnings for unknown and duplicate fields (current state), also print warnings for fields with incorrect case sensitivity - e.g. `controlPlaneEndpoint` (valid), `ControlPlaneEndpoint` (invalid). Instead of only printing warnings during `init` and `join` also print warnings when downloading the ClusterConfiguration, KubeletConfiguration or KubeProxyConfiguration objects from the cluster. This can be useful if the user has patched these objects in their respective ConfigMaps with mistakes. ([#107725](https://github.com/kubernetes/kubernetes/pull/107725), [@neolit123](https://github.com/neolit123))
- Kubectl now supports shell completion for the <type>/<name> format for specifying resources.
  kubectl now provides shell completion for container names following the `--container/-c` flag of the `exec` command.
  kubectl's shell completion now suggests resource types for commands that only apply to pods. ([#108493](https://github.com/kubernetes/kubernetes/pull/108493), [@marckhouzam](https://github.com/marckhouzam))
- Kubelet: add `kubelet_volume_metric_collection_duration_seconds` metrics for volume disk usage calculation duration ([#107201](https://github.com/kubernetes/kubernetes/pull/107201), [@pacoxu](https://github.com/pacoxu))
- Kubelet: the following dockershim related flags are also removed along with dockershim `--experimental-dockershim-root-directory`, `--docker-endpoint`, `--image-pull-progress-deadline`, `--network-plugin`, `--cni-conf-dir`, `--cni-bin-dir`, `--cni-cache-dir`, `--network-plugin-mtu`. ([#106907](https://github.com/kubernetes/kubernetes/pull/106907), [@cyclinder](https://github.com/cyclinder))
- Kubernetes 1.24 bumped version of golang it is compiled with to go1.18, which introduced significant changes to its garbage collection algorithm. As a result, we observed an increase in memory usage for kube-apiserver in larger an heavily loaded clusters up to ~25% (with the benefit of API call latencies drop by up to 10x on 99th percentiles). If the memory increase is not acceptable for you you can mitigate by setting GOGC env variable (for our tests using GOGC=63 brings memory usage back to original value, although the exact value may depend on usage patterns on your cluster). ([#108870](https://github.com/kubernetes/kubernetes/pull/108870), [@dims](https://github.com/dims))
- Kubernetes 1.24 is built with go1.18, which will no longer validate certificates signed with a SHA-1 hash algorithm by default. See https://golang.org/doc/go1.18#sha1 for more details. If you are using certificates like this in admission or conversion ([#109024](https://github.com/kubernetes/kubernetes/pull/109024), [@stlaz](https://github.com/stlaz))
- Leader Migration is now GA. All new configuration files onwards should use version v1. ([#109072](https://github.com/kubernetes/kubernetes/pull/109072), [@jiahuif](https://github.com/jiahuif))
- Mark AzureDisk CSI migration as GA ([#107681](https://github.com/kubernetes/kubernetes/pull/107681), [@andyzhangx](https://github.com/andyzhangx))
- Move volume expansion feature to GA ([#108929](https://github.com/kubernetes/kubernetes/pull/108929), [@gnufied](https://github.com/gnufied))
- Moving MixedProtocolLBService from alpha to beta ([#109213](https://github.com/kubernetes/kubernetes/pull/109213), [@bridgetkromhout](https://github.com/bridgetkromhout))
- New "field_validation_request_duration_seconds" metric, measures how long requests take, indicating the value of the fieldValidation query parameter and whether or not server-side field validation is enabled on the apiserver ([#109120](https://github.com/kubernetes/kubernetes/pull/109120), [@kevindelgado](https://github.com/kevindelgado))
- New feature gate, ServiceIPStaticSubrange, to enable the new strategy in the Service IP allocators, so the IP range is subdivided and dynamic allocated ClusterIP addresses for Services are allocated preferently from the upper range. ([#106792](https://github.com/kubernetes/kubernetes/pull/106792), [@aojea](https://github.com/aojea))
- OpenAPI definitions served by kube-apiserver now include enum types by default. ([#108898](https://github.com/kubernetes/kubernetes/pull/108898), [@jiahuif](https://github.com/jiahuif))
- OpenStack Cinder CSI migration is now GA and switched on by default, Cinder CSI driver must be installed on clusters on OpenStack for Cinder volumes to work (has been since v1.21). ([#107462](https://github.com/kubernetes/kubernetes/pull/107462), [@dims](https://github.com/dims))
- PreFilter extension in the scheduler framework now returns not only status but also PreFilterResult ([#108648](https://github.com/kubernetes/kubernetes/pull/108648), [@ahg-g](https://github.com/ahg-g))
- Promoted graceful shutdown based on pod priority to beta ([#107986](https://github.com/kubernetes/kubernetes/pull/107986), [@wzshiming](https://github.com/wzshiming))
- Removed feature gate `SetHostnameAsFQDN`. ([#108038](https://github.com/kubernetes/kubernetes/pull/108038), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Removed kube-scheduler insecure flags. You can use `--bind-address` and `--secure-port` instead. ([#106865](https://github.com/kubernetes/kubernetes/pull/106865), [@jonyhy96](https://github.com/jonyhy96))
- Removed the `ImmutableEphemeralVolumes` feature gate. ([#107152](https://github.com/kubernetes/kubernetes/pull/107152), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Set `PodMaxUnschedulableQDuration` as 5 min. ([#108761](https://github.com/kubernetes/kubernetes/pull/108761), [@denkensk](https://github.com/denkensk))
- Support in-tree PV deletion protection finalizer. ([#108400](https://github.com/kubernetes/kubernetes/pull/108400), [@deepakkinni](https://github.com/deepakkinni))
- The `.spec.loadBalancerClass` field for Services is now generally available. ([#107979](https://github.com/kubernetes/kubernetes/pull/107979), [@XudongLiuHarold](https://github.com/XudongLiuHarold))
- The `NamespaceDefaultLabelName` feature gate, GA since v1.22, is now removed. ([#106838](https://github.com/kubernetes/kubernetes/pull/106838), [@mengjiao-liu](https://github.com/mengjiao-liu))
- The `kubectl logs` will now warn and default to the first container in a pod. This new behavior brings it in line with `kubectl exec`. ([#105964](https://github.com/kubernetes/kubernetes/pull/105964), [@kidlj](https://github.com/kidlj))
- The `v1` version of `LeaderMigrationConfiguration` supports only `leases` API for leader election. To use formerly supported mechanisms, please continue using `v1beta1`. ([#108016](https://github.com/kubernetes/kubernetes/pull/108016), [@jiahuif](https://github.com/jiahuif))
- The kubelet now creates an iptables chain named `KUBE-IPTABLES-HINT` in
  the `mangle` table. Containerized components that need to modify iptables
  rules in the host network namespace can use the existence of this chain
  to more-reliably determine whether the system is using iptables-legacy or
  iptables-nft. ([#109059](https://github.com/kubernetes/kubernetes/pull/109059), [@danwinship](https://github.com/danwinship))
- The output of `kubectl describe ingress` now includes an IngressClass name if available. ([#107921](https://github.com/kubernetes/kubernetes/pull/107921), [@mpuckett159](https://github.com/mpuckett159))
- The scheduler prints info logs when the extender returned an error. (`--v>5`) ([#107974](https://github.com/kubernetes/kubernetes/pull/107974), [@sanposhiho](https://github.com/sanposhiho))
- The script `cluster/gce/gci/configure.sh` now supports downloading `crictl` on ARM64 nodes ([#108034](https://github.com/kubernetes/kubernetes/pull/108034), [@tstapler](https://github.com/tstapler))
- Turn on `CSIMigrationAzureFile` by default on 1.24 ([#105070](https://github.com/kubernetes/kubernetes/pull/105070), [@andyzhangx](https://github.com/andyzhangx))
- Update the k8s.io/system-validators library to v1.7.0 ([#108988](https://github.com/kubernetes/kubernetes/pull/108988), [@neolit123](https://github.com/neolit123))
- Updated golang.org/x/net to v0.0.0-20211209124913-491a49abca63. ([#106949](https://github.com/kubernetes/kubernetes/pull/106949), [@cpanato](https://github.com/cpanato))
- Updates `kubectl kustomize` and `kubectl apply -k` to Kustomize v4.5.4 ([#108994](https://github.com/kubernetes/kubernetes/pull/108994), [@KnVerey](https://github.com/KnVerey))
- When invoked with `-list-images`, the `e2e.test` binary now also lists the images that might be needed for storage tests. ([#108458](https://github.com/kubernetes/kubernetes/pull/108458), [@pohly](https://github.com/pohly))
- `kubectl config delete-user` now supports completion ([#107142](https://github.com/kubernetes/kubernetes/pull/107142), [@dimbleby](https://github.com/dimbleby))
- `kubectl create token` can now be used to request a service account token, and permission to request service account tokens is added to the `edit` and `admin` RBAC roles ([#107880](https://github.com/kubernetes/kubernetes/pull/107880), [@liggitt](https://github.com/liggitt))
- `kubectl version` now includes information on the embedded version of Kustomize ([#108817](https://github.com/kubernetes/kubernetes/pull/108817), [@KnVerey](https://github.com/KnVerey))

##### Bug or Regression

- Fix missing delete events on informer re-lists to ensure all delete events are correctly emitted and using the latest known object state, so that all event handlers and stores always reflect the actual apiserver state as best as possible ([#115901](https://github.com/kubernetes/kubernetes/pull/115901), [@odinuge](https://github.com/odinuge)) [SIG API Machinery]
- Fix: Route controller should update routes with NodeIP changed ([#116360](https://github.com/kubernetes/kubernetes/pull/116360), [@lzhecheng](https://github.com/lzhecheng)) [SIG Cloud Provider]
- Kubelet: Fix fs quota monitoring on volumes ([#116795](https://github.com/kubernetes/kubernetes/pull/116795), [@pacoxu](https://github.com/pacoxu)) [SIG Storage]
- Fix the regression that introduced 34s timeout for DELETECOLLECTION calls ([#115482](https://github.com/kubernetes/kubernetes/pull/115482), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fixed bug which caused the status of Indexed Jobs to only be updated when there are newly completed indexes. The completed indexes are now updated if the .status.completedIndexes has values outside of the [0, .spec.completions> range ([#115457](https://github.com/kubernetes/kubernetes/pull/115457), [@danielvegamyhre](https://github.com/danielvegamyhre)) [SIG Apps]
- Golang.org/x/net updates to v0.7.0 to fix CVE-2022-41723 ([#115789](https://github.com/kubernetes/kubernetes/pull/115789), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Security and Storage]
- The Kubernetes API server now correctly detects and closes existing TLS connections when its client certificate file for kubelet authentication has been rotated. ([#115580](https://github.com/kubernetes/kubernetes/pull/115580), [@enj](https://github.com/enj)) [SIG API Machinery, Node and Testing]
- Client-go: fixes potential data races retrying requests using a custom io.Reader body; with this fix, only requests with no body or with string / []byte / runtime.Object bodies can be retried ([#113933](https://github.com/kubernetes/kubernetes/pull/113933), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Do not include preemptor pod metadata in the event message ([#115024](https://github.com/kubernetes/kubernetes/pull/115024), [@mimowo](https://github.com/mimowo)) [SIG Scheduling]
- Failed pods associated with a job with `parallelism = 1` are recreated by the job controller honoring exponential backoff delay again. However, for jobs with `parallelism > 1`, pods might be created without exponential backoff delay. ([#115021](https://github.com/kubernetes/kubernetes/pull/115021), [@nikhita](https://github.com/nikhita)) [SIG Apps]
- Fix a regression that the scheduler always goes through all Filter plugins. ([#114526](https://github.com/kubernetes/kubernetes/pull/114526), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix bug in CRD Validation Rules (beta) and ValidatingAdmissionPolicy (alpha) where all admission requests could result in `internal error: runtime error: index out of range [3] with length 3 evaluating rule: <rule name>` under certain circumstances. ([#114865](https://github.com/kubernetes/kubernetes/pull/114865), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery]
- Fix performance issue when creating large objects using SSA with fully unspecified schemas (preserveUnknownFields). ([#111915](https://github.com/kubernetes/kubernetes/pull/111915), [@aojea](https://github.com/aojea)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
- Fixed StatefulSet to show the valid status even if the new replica creation fails. ([#112083](https://github.com/kubernetes/kubernetes/pull/112083), [@gjkim42](https://github.com/gjkim42)) [SIG Apps and Testing]
- Fixing issue in Winkernel Proxier - Unexpected active TCP connection drops while horizontally scaling the endpoints for a LoadBalancer Service with External Traffic Policy: Local ([#114040](https://github.com/kubernetes/kubernetes/pull/114040), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Fixing issue with Winkernel Proxier - No ingress load balancer rules with endpoints to support load balancing when all the endpoints are terminating. ([#114451](https://github.com/kubernetes/kubernetes/pull/114451), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Kube-apiserver: bugfix DeleteCollection API fails if request body is non-empty ([#113968](https://github.com/kubernetes/kubernetes/pull/113968), [@sxllwx](https://github.com/sxllwx)) [SIG API Machinery]
- Optimizing loadbalancer creation with the help of attribute Internal Traffic Policy: Local ([#114466](https://github.com/kubernetes/kubernetes/pull/114466), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Update the system-validators library to v1.8.0 ([#114060](https://github.com/kubernetes/kubernetes/pull/114060), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- [aws] Fixed a bug which reduces the number of unnecessary calls to STS in the event of assume role failures in the legacy cloud provider ([#110706](https://github.com/kubernetes/kubernetes/pull/110706), [@prateekgogia](https://github.com/prateekgogia)) [SIG Cloud Provider]
- Fix endpoint reconciler not being able to delete the apiserver lease on shutdown ([#114138](https://github.com/kubernetes/kubernetes/pull/114138), [@aojea](https://github.com/aojea)) [SIG API Machinery]
- Fix for volume reconstruction of CSI ephemeral volumes ([#113346](https://github.com/kubernetes/kubernetes/pull/113346), [@dobsonj](https://github.com/dobsonj)) [SIG Node, Storage and Testing]
- Kube-apiserver: resolves possible hung connections using konnectivity network proxy with TCP or UDS HTTP connect configurations ([#113862](https://github.com/kubernetes/kubernetes/pull/113862), [@jkh52](https://github.com/jkh52)) [SIG API Machinery]
- Resolves an issue that causes winkernel proxier to treat stale VIPs as valid ([#113567](https://github.com/kubernetes/kubernetes/pull/113567), [@daschott](https://github.com/daschott)) [SIG Network and Windows]
- Updates golang.org/x/net to fix CVE-2022-41717 ([#114322](https://github.com/kubernetes/kubernetes/pull/114322), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
- Updates golang.org/x/net to v0.1.1-0.20221027164007-c63010009c80 to resolve CVE-2022-27664 ([#113459](https://github.com/kubernetes/kubernetes/pull/113459), [@aimuz](https://github.com/aimuz)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Release and Storage]
- Volumes are no longer detached from healthy nodes after 6 minutes timeout. 6 minute force-detach timeout is used only for unhealthy nodes (`node.status.conditions["Ready"] != true`). ([#110721](https://github.com/kubernetes/kubernetes/pull/110721), [@jsafrane](https://github.com/jsafrane)) [SIG Apps]
- Consider only plugin directory and not entire kubelet root when cleaning up mounts ([#112920](https://github.com/kubernetes/kubernetes/pull/112920), [@mattcary](https://github.com/mattcary)) [SIG Storage]
- Etcd: Update to v3.5.5 ([#113099](https://github.com/kubernetes/kubernetes/pull/113099), [@mk46](https://github.com/mk46)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle and Testing]
- Fixed a bug where a change in the `appProtocol` for a Service did not trigger a load balancer update. ([#113032](https://github.com/kubernetes/kubernetes/pull/113032), [@MartinForReal](https://github.com/MartinForReal)) [SIG Cloud Provider and Network]
- Kube-proxy, will restart in case it detects that the Node assigned pod.Spec.PodCIDRs have changed ([#113252](https://github.com/kubernetes/kubernetes/pull/113252), [@code-elinka](https://github.com/code-elinka)) [SIG Cloud Provider, Network, Node and Storage]
- Kubelet no longer reports terminated container metrics from cAdvisor ([#112963](https://github.com/kubernetes/kubernetes/pull/112963), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
- Kubelet: fix GetAllocatableCPUs method in cpumanager ([#113421](https://github.com/kubernetes/kubernetes/pull/113421), [@Garrybest](https://github.com/Garrybest)) [SIG Node]
- Pod logs using --timestamps are not broken up with timestamps anymore. ([#113516](https://github.com/kubernetes/kubernetes/pull/113516), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Allow Label section in vsphere e2e cloudprovider configuration ([#112479](https://github.com/kubernetes/kubernetes/pull/112479), [@gnufied](https://github.com/gnufied)) [SIG Storage and Testing]
- Kube-apiserver: gzip compression switched from level 4 to level 1 to improve large list call latencies in exchange for higher network bandwidth usage (10-50% higher). This increases the headroom before very large unpaged list calls exceed request timeout limits. ([#112399](https://github.com/kubernetes/kubernetes/pull/112399), [@shyamjvs](https://github.com/shyamjvs)) [SIG API Machinery]
- Kube-apiserver: resolved a regression that treated `304 Not Modified` responses from aggregated API servers as internal errors ([#112528](https://github.com/kubernetes/kubernetes/pull/112528), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Kubeadm: allow RSA and ECDSA format keys in preflight check ([#112535](https://github.com/kubernetes/kubernetes/pull/112535), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Fix an ephemeral port exhaustion bug caused by improper connection management that occurred when a large number of objects were handled by kubectl while exec auth was in use. ([#112337](https://github.com/kubernetes/kubernetes/pull/112337), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]
- Fix problem in updating VolumeAttached in node status ([#112304](https://github.com/kubernetes/kubernetes/pull/112304), [@xing-yang](https://github.com/xing-yang)) [SIG Apps]
- Kube-apiserver: redirect responses are no longer returned from backends by default. Set `--aggregator-reject-forwarding-redirect=false` to continue forwarding redirect responses. ([#112331](https://github.com/kubernetes/kubernetes/pull/112331), [@enj](https://github.com/enj)) [SIG API Machinery]
- UserName check for 'ContainerAdministrator' is now case-insensitive if runAsNonRoot is set to true on Windows. ([#112211](https://github.com/kubernetes/kubernetes/pull/112211), [@PushkarJ](https://github.com/PushkarJ)) [SIG Node, Testing and Windows]
- Fix JobTrackingWithFinalizers when a pod succeeds after the job is considered failed, which led to API conflicts that blocked finishing the job. ([#111664](https://github.com/kubernetes/kubernetes/pull/111664), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and Testing]
- Fix memory leak in the job controller related to JobTrackingWithFinalizers ([#111722](https://github.com/kubernetes/kubernetes/pull/111722), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps]
- Fix memory leak on kube-scheduler preemption ([#111803](https://github.com/kubernetes/kubernetes/pull/111803), [@amewayne](https://github.com/amewayne)) [SIG Scheduling]
- Fixed potential scheduler crash when scheduling with unsatisfied nodes in PodTopologySpread. ([#111511](https://github.com/kubernetes/kubernetes/pull/111511), [@kerthcet](https://github.com/kerthcet)) [SIG Scheduling]
- Fixing issue on Windows nodes where HostProcess containers may not be created as expected. ([#110966](https://github.com/kubernetes/kubernetes/pull/110966), [@marosset](https://github.com/marosset)) [SIG Node and Windows]
- If the parent directory of the file specified in the `--audit-log-path` argument does not exist, Kubernetes now creates it. ([#111225](https://github.com/kubernetes/kubernetes/pull/111225), [@vpnachev](https://github.com/vpnachev)) [SIG Auth]
- Namespace editors and admins can now create leases.coordination.k8s.io and should use this type for leaderelection instead of configmaps. ([#111515](https://github.com/kubernetes/kubernetes/pull/111515), [@deads2k](https://github.com/deads2k)) [SIG API Machinery and Auth]
- Reduce API server memory when many CRDs are loaded by sharing a single etcd3 client logger across all clients ([#111648](https://github.com/kubernetes/kubernetes/pull/111648), [@negz](https://github.com/negz)) [SIG API Machinery]
- Run kubelet, when there is an error exit, print the error log ([#110917](https://github.com/kubernetes/kubernetes/pull/110917), [@yangjunmyfm192085](https://github.com/yangjunmyfm192085)) [SIG Node]
- Fix a bug on endpointslices tests comparing the wrong metrics ([#110920](https://github.com/kubernetes/kubernetes/pull/110920), [@jluhrsen](https://github.com/jluhrsen)) [SIG Apps and Network]
- Fix a bug that caused the wrong result length when using --chunk-size and --selector together ([#110735](https://github.com/kubernetes/kubernetes/pull/110735), [@Abirdcfly](https://github.com/Abirdcfly)) [SIG API Machinery and Testing]
- Fix bug that prevented the job controller from enforcing activeDeadlineSeconds when set ([#110544](https://github.com/kubernetes/kubernetes/pull/110544), [@harshanarayana](https://github.com/harshanarayana)) [SIG Apps]
- Fix image pulling failure when IMDS is unavailable in kubelet startup ([#110523](https://github.com/kubernetes/kubernetes/pull/110523), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix printing resources with int64 fields ([#110572](https://github.com/kubernetes/kubernetes/pull/110572), [@sanchezl](https://github.com/sanchezl)) [SIG API Machinery]
- Fix unnecessary recreation of placeholder EndpointSlice ([#110732](https://github.com/kubernetes/kubernetes/pull/110732), [@jluhrsen](https://github.com/jluhrsen)) [SIG Apps and Network]
- Fixed a regression introduced in 1.24.0 where Azure load balancers were not kept up to date with the state of cluster nodes. In particular, nodes that are not in the ready state and are not newly created (i.e. not having the `node.cloudprovider.kubernetes.io/uninitialized` taint) now get removed from Azure load balancers. ([#109931](https://github.com/kubernetes/kubernetes/pull/109931), [@ricky-rav](https://github.com/ricky-rav)) [SIG Cloud Provider]
- Kubeadm: fix error adding extra prefix unix:// to CRI endpoints that were missing URL scheme ([#110634](https://github.com/kubernetes/kubernetes/pull/110634), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubeadm: fix the bug that configurable KubernetesVersion not respected during kubeadm join ([#111021](https://github.com/kubernetes/kubernetes/pull/111021), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- EndpointSlices marked for deletion are now ignored during reconciliation. ([#110484](https://github.com/kubernetes/kubernetes/pull/110484), [@aryan9600](https://github.com/aryan9600)) [SIG Apps and Network]
- Fixed a kubelet issue that could result in invalid pod status updates to be sent to the api-server where pods would be reported in a terminal phase but also report a ready condition of true in some cases. ([#110479](https://github.com/kubernetes/kubernetes/pull/110479), [@bobbypage](https://github.com/bobbypage)) [SIG Node and Testing]
- Pods will now post their readiness during termination. ([#110416](https://github.com/kubernetes/kubernetes/pull/110416), [@aojea](https://github.com/aojea)) [SIG Network, Node and Testing]
- The pod phase lifecycle guarantees that terminal Pods, those whose states are Unready or Succeeded, can not regress and will have all container stopped. Hence, terminal Pods will never be reachable and should not publish their IP addresses on the Endpoints or EndpointSlices, independently of the Service TolerateUnready option. ([#110258](https://github.com/kubernetes/kubernetes/pull/110258), [@robscott](https://github.com/robscott)) [SIG Apps, Network, Node and Testing]
- Fix JobTrackingWithFinalizers that:
  - was declaring a job finished before counting all the created pods in the status
  - was leaving pods with finalizers, blocking pod and job deletions
  
  JobTrackingWithFinalizers is still disabled by default. ([#109486](https://github.com/kubernetes/kubernetes/pull/109486), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and Testing]
- Kubeadm: only taint control plane nodes when the legacy "master" taint is present. This avoids a bug where "kubeadm upgrade" will re-taint a control plane node with the new "control plane" taint even if the user explicitly untainted the node. ([#109841](https://github.com/kubernetes/kubernetes/pull/109841), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- A node IP provided to kublet via `--node-ip` will now be preferred for when determining the node's primary IP and using the external cloud provider (CCM). ([#107750](https://github.com/kubernetes/kubernetes/pull/107750), [@stephenfin](https://github.com/stephenfin))
- A static pod that is rapidly updated was failing to start until the Kubelet was restarted. ([#107900](https://github.com/kubernetes/kubernetes/pull/107900), [@smarterclayton](https://github.com/smarterclayton))
- Add one metrics(`kubelet_volume_stats_health_abnormal`) of volume health state to kubelet ([#108758](https://github.com/kubernetes/kubernetes/pull/108758), [@fengzixu](https://github.com/fengzixu))
- Added a new label `type` to `apiserver_flowcontrol_request_execution_seconds` metric - it has the following values: - 'regular': indicates that it is a non long running request - 'watch': indicates that it is a watch request. ([#105517](https://github.com/kubernetes/kubernetes/pull/105517), [@tkashem](https://github.com/tkashem))
- Added a test to guarantee that conformance clusters require at least 2 untainted nodes. ([#106313](https://github.com/kubernetes/kubernetes/pull/106313), [@aojea](https://github.com/aojea))
- Adds PV deletion protection finalizer only when PV reclaimPolicy is Delete for dynamically provisioned volumes. ([#109205](https://github.com/kubernetes/kubernetes/pull/109205), [@deepakkinni](https://github.com/deepakkinni))
- Allowed attached volumes to be mounted quicker by skipping exponential backoff when checking for reported-in-use volumes. ([#106853](https://github.com/kubernetes/kubernetes/pull/106853), [@gnufied](https://github.com/gnufied))
- Alowed useful inclusion of `-args $prog_args` in KUBE_TEST_ARGS, when doing `make test-integration`. ([#107516](https://github.com/kubernetes/kubernetes/pull/107516), [@MikeSpreitzer](https://github.com/MikeSpreitzer))
- An inefficient lock in EndpointSlice controller metrics cache has been reworked. Network programming latency may be significantly reduced in certain scenarios, especially in clusters with a large number of Services. ([#107091](https://github.com/kubernetes/kubernetes/pull/107091), [@robscott](https://github.com/robscott))
- Apiserver will now reject connection attempts to `0.0.0.0/::` when handling a proxy subresource request. ([#107402](https://github.com/kubernetes/kubernetes/pull/107402), [@anguslees](https://github.com/anguslees))
- Bug: client-go clientset was not defaulting to the user agent, and was using the default golang agent for all the requests. ([#108772](https://github.com/kubernetes/kubernetes/pull/108772), [@aojea](https://github.com/aojea))
- Bump `sigs.k8s.io/apiserver-network-proxy/konnectivity-client@v0.0.30` to fix a goroutine leak in kube-apiserver when using egress selctor with the gRPC mode. ([#108437](https://github.com/kubernetes/kubernetes/pull/108437), [@andrewsykim](https://github.com/andrewsykim))
- CEL validation failure returns object type instead of object. ([#107090](https://github.com/kubernetes/kubernetes/pull/107090), [@cici37](https://github.com/cici37))
- CRI-API: IPs returned by `PodSandboxNetworkStatus`` are ignored by the kubelet for host-network pods. ([#106715](https://github.com/kubernetes/kubernetes/pull/106715), [@aojea](https://github.com/aojea))
- Call `NodeExpand` on all nodes in case of RWX volumes ([#108693](https://github.com/kubernetes/kubernetes/pull/108693), [@gnufied](https://github.com/gnufied))
- Changed node staging path for CSI driver to use a PV agnostic path. Nodes must be drained before updating the kubelet with this change. ([#107065](https://github.com/kubernetes/kubernetes/pull/107065), [@saikat-royc](https://github.com/saikat-royc))
- Client-go: fixed the paged list calls with `ResourceVersionMatch` set would fail once paging is kicked in. ([#107311](https://github.com/kubernetes/kubernetes/pull/107311), [@fasaxc](https://github.com/fasaxc))
- Correct event registration for multiple scheduler plugins; this fixes a potential significant delay in re-queueing unschedulable pods. ([#109442](https://github.com/kubernetes/kubernetes/pull/109442), [@ahg-g](https://github.com/ahg-g))
- Etcd: Update to v3.5.3 ([#109471](https://github.com/kubernetes/kubernetes/pull/109471), [@justaugustus](https://github.com/justaugustus))
- Existing InTree AzureFile PVs which don't have a secret namespace defined will now work properly after enabling CSI migration - the namespace will be obtained from ClaimRef. ([#108000](https://github.com/kubernetes/kubernetes/pull/108000), [@RomanBednar](https://github.com/RomanBednar))
- Failure to start a container cannot accidentally result in the pod being considered "Succeeded" in the presence of deletion. ([#107845](https://github.com/kubernetes/kubernetes/pull/107845), [@smarterclayton](https://github.com/smarterclayton))
- Fix a race in the timeout handler that could lead to kube-apiserver crashes ([#108455](https://github.com/kubernetes/kubernetes/pull/108455), [@Argh4k](https://github.com/Argh4k))
- Fix container creation errors for pods with cpu requests bigger than 256 cpus ([#106570](https://github.com/kubernetes/kubernetes/pull/106570), [@odinuge](https://github.com/odinuge))
- Fix issue where the job controller might not remove the job tracking finalizer from pods when deleting a job, or when the pod is orphan ([#108752](https://github.com/kubernetes/kubernetes/pull/108752), [@alculquicondor](https://github.com/alculquicondor))
- Fix libct/cg/fs2: fixed GetStats for unsupported hugetlb error on Raspbian Bullseye ([#106912](https://github.com/kubernetes/kubernetes/pull/106912), [@Letme](https://github.com/Letme))
- Fix the bug that the outdated services may be sent to the cloud provider ([#107631](https://github.com/kubernetes/kubernetes/pull/107631), [@lzhecheng](https://github.com/lzhecheng))
- Fix the overestimated cost of delegated API requests in kube-apiserver API priority & fairness ([#109188](https://github.com/kubernetes/kubernetes/pull/109188), [@wojtek-t](https://github.com/wojtek-t))
- Fix to allow `fsGroup` to be applied for CSI Inline Volumes ([#108662](https://github.com/kubernetes/kubernetes/pull/108662), [@dobsonj](https://github.com/dobsonj))
- Fixed CSI migration of Azure Disk in-tree StorageClasses with topology requirements in Azure regions that do not have availability zones. ([#109154](https://github.com/kubernetes/kubernetes/pull/109154), [@jsafrane](https://github.com/jsafrane))
- Fixed `--retries` functionality for negative values in `kubectl cp` ([#108748](https://github.com/kubernetes/kubernetes/pull/108748), [@atiratree](https://github.com/atiratree))
- Fixed `azureDisk` parameter lowercase translation issue. ([#107429](https://github.com/kubernetes/kubernetes/pull/107429), [@andyzhangx](https://github.com/andyzhangx))
- Fixed `azureFile` `volumeID` collision issue in CSI migration. ([#107575](https://github.com/kubernetes/kubernetes/pull/107575), [@andyzhangx](https://github.com/andyzhangx))
- Fixed a bug in attachdetach controller that didn't properly handle kube-apiserver errors leading to stuck attachments/detachments. ([#108167](https://github.com/kubernetes/kubernetes/pull/108167), [@jfremy](https://github.com/jfremy))
- Fixed a bug that a pod's `.status.nominatedNodeName` is not cleared properly, and thus over-occupied system resources. ([#106816](https://github.com/kubernetes/kubernetes/pull/106816), [@Huang-Wei](https://github.com/Huang-Wei))
- Fixed a bug that caused credentials in an exec plugin to override the static certificates set in a kubeconfig. ([#107410](https://github.com/kubernetes/kubernetes/pull/107410), [@margocrawf](https://github.com/margocrawf))
- Fixed a bug that could cause panic when a `/healthz` request times out. ([#107034](https://github.com/kubernetes/kubernetes/pull/107034), [@benluddy](https://github.com/benluddy))
- Fixed a bug that out-of-tree plugin is misplaced when using scheduler v1beta3 config ([#108613](https://github.com/kubernetes/kubernetes/pull/108613), [@Huang-Wei](https://github.com/Huang-Wei))
- Fixed a bug where a partial `EndpointSlice` update could cause node name information to be dropped from endpoints that were not updated. ([#108198](https://github.com/kubernetes/kubernetes/pull/108198), [@liggitt](https://github.com/liggitt))
- Fixed a bug where unwanted fields were being returned from a `create --dry-run`: uid and, if generateName was used, name. ([#107088](https://github.com/kubernetes/kubernetes/pull/107088), [@joejulian](https://github.com/joejulian))
- Fixed a bug where vSphere client connections where not being closed during testing. Leaked vSphere client sessions were causing resource exhaustion during automated testing. ([#107337](https://github.com/kubernetes/kubernetes/pull/107337), [@derek-pryor](https://github.com/derek-pryor))
- Fixed a panic when using invalid output format in `kubectl create secret` command. ([#107221](https://github.com/kubernetes/kubernetes/pull/107221), [@rikatz](https://github.com/rikatz))
- Fixed a rare race condition handling requests that timeout. ([#107452](https://github.com/kubernetes/kubernetes/pull/107452), [@liggitt](https://github.com/liggitt))
- Fixed a regression in 1.23 that incorrectly pruned data from array items of a custom resource that set `x-kubernetes-preserve-unknown-fields: true`. ([#107688](https://github.com/kubernetes/kubernetes/pull/107688), [@liggitt](https://github.com/liggitt))
- Fixed a regression in 1.23 where update requests to previously persisted `Service` objects that have not been modified since 1.19 can be rejected with an incorrect `spec.clusterIPs: Required value` error. ([#107847](https://github.com/kubernetes/kubernetes/pull/107847), [@thockin](https://github.com/thockin))
- Fixed a regression that could incorrectly reject pods with `OutOfCpu` errors if they were rapidly scheduled after other pods were reported as complete in the API. The Kubelet now waits to report the phase of a pod as terminal in the API until all running containers are guaranteed to have stopped and no new containers can be started.  Short-lived pods may take slightly longer (~1s) to report Succeeded or Failed after this change. ([#108366](https://github.com/kubernetes/kubernetes/pull/108366), [@smarterclayton](https://github.com/smarterclayton))
- Fixed bug in `TopologyManager` for ensuring aligned allocations on machines with more than 2 NUMA nodes ([#108052](https://github.com/kubernetes/kubernetes/pull/108052), [@klueska](https://github.com/klueska))
- Fixed bug in error messaging for basic-auth and ssh secret validations. ([#106179](https://github.com/kubernetes/kubernetes/pull/106179), [@vivek-koppuru](https://github.com/vivek-koppuru))
- Fixed detaching CSI volumes from nodes when a CSI driver name has prefix "csi-". ([#107025](https://github.com/kubernetes/kubernetes/pull/107025), [@jsafrane](https://github.com/jsafrane))
- Fixed duplicate port opening in kube-proxy when `--nodeport-addresses` is empty. ([#107413](https://github.com/kubernetes/kubernetes/pull/107413), [@tnqn](https://github.com/tnqn))
- Fixed handling of objects with invalid selectors. ([#107559](https://github.com/kubernetes/kubernetes/pull/107559), [@liggitt](https://github.com/liggitt))
- Fixed indexer bug that resulted in incorrect index updates if number of index values for a given object was changing during update ([#109137](https://github.com/kubernetes/kubernetes/pull/109137), [@wojtek-t](https://github.com/wojtek-t))
- Fixed kubectl bug where bash completions don't work if `--context` flag is specified with a value that contains a colon. ([#107439](https://github.com/kubernetes/kubernetes/pull/107439), [@brianpursley](https://github.com/brianpursley))
- Fixed performance regression in JSON logging caused by syncing stdout every time error was logged. ([#107035](https://github.com/kubernetes/kubernetes/pull/107035), [@serathius](https://github.com/serathius))
- Fixed regression in CPUManager that it will release exclusive CPUs in app containers inherited from init containers when the init containers were removed. ([#104837](https://github.com/kubernetes/kubernetes/pull/104837), [@eggiter](https://github.com/eggiter))
- Fixed static pod add and removes restarts in certain cases. ([#107695](https://github.com/kubernetes/kubernetes/pull/107695), [@rphillips](https://github.com/rphillips))
- Fixed: deleted a non-existent Azure disk issue. ([#107406](https://github.com/kubernetes/kubernetes/pull/107406), [@andyzhangx](https://github.com/andyzhangx))
- Fixed: do not return early in the node informer when there is no change of the topology label. ([#108149](https://github.com/kubernetes/kubernetes/pull/108149), [@nilo19](https://github.com/nilo19))
- Fixed: removed outdated ipv4 route when the corresponding node is deleted. ([#106164](https://github.com/kubernetes/kubernetes/pull/106164), [@nilo19](https://github.com/nilo19))
- Fixes bug in CronJob Controller V2 where it would lose track of jobs upon job template labels change. ([#107997](https://github.com/kubernetes/kubernetes/pull/107997), [@d-honeybadger](https://github.com/d-honeybadger))
- If drainer has nil for Ctx or Client it will error with `RunCordonOrUncordon`. ([#105297](https://github.com/kubernetes/kubernetes/pull/105297), [@jackfrancis](https://github.com/jackfrancis))
- Improved handling of unmount failures when device may be in-use by another container/process. ([#107789](https://github.com/kubernetes/kubernetes/pull/107789), [@gnufied](https://github.com/gnufied))
- Improved logging when volume times out waiting for attach/detach. ([#108628](https://github.com/kubernetes/kubernetes/pull/108628), [@RomanBednar](https://github.com/RomanBednar))
- Improved the rounding of `PodTopologySpread` scores to offer better scoring when spreading a low number of pods. ([#107384](https://github.com/kubernetes/kubernetes/pull/107384), [@sanposhiho](https://github.com/sanposhiho))
- Increase Azure ACR credential provider timeout ([#108209](https://github.com/kubernetes/kubernetes/pull/108209), [@andyzhangx](https://github.com/andyzhangx))
- Kube-apiserver: Server Side Apply merge order is reverted to match v1.22 behavior until http://issue.k8s.io/104641 is resolved. ([#106660](https://github.com/kubernetes/kubernetes/pull/106660), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: ensures the namespace of objects sent to admission webhooks matches the request namespace. Previously, objects without a namespace set would have the request namespace populated after mutating admission, and objects with a namespace that did not match the request namespace would be rejected after admission. ([#94637](https://github.com/kubernetes/kubernetes/pull/94637), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: removed `apf_fd` from server logs which could contain data identifying the requesting user ([#108631](https://github.com/kubernetes/kubernetes/pull/108631), [@jupblb](https://github.com/jupblb))
- Kube-proxy in iptables mode now only logs the full iptables input at `-v=9` rather than `-v=5`. ([#108224](https://github.com/kubernetes/kubernetes/pull/108224), [@danwinship](https://github.com/danwinship))
- Kube-proxy will no longer hold service node ports open on the node. Users are still advised not to run any listener on node ports range used by kube-proxy. ([#108496](https://github.com/kubernetes/kubernetes/pull/108496), [@khenidak](https://github.com/khenidak))
- Kubeadm: allow the `certs check-expiration` command to not require the existence of the cluster CA key (ca.key file) when checking the expiration of managed certificates in kubeconfig files. ([#106854](https://github.com/kubernetes/kubernetes/pull/106854), [@neolit123](https://github.com/neolit123))
- Kubeadm: during execution of the `certs check-expiration` command, treat the etcd CA as external if there is a missing etcd CA key file (etcd/ca.key) and perform the proper validation on certificates signed by the etcd CA. Additionally, make sure that the CA for all entries in the output table is included - for both certificates on disk and in kubeconfig files. ([#106891](https://github.com/kubernetes/kubernetes/pull/106891), [@neolit123](https://github.com/neolit123))
- Kubeadm: fixed a bug related to a warning printed if the `KubeletConfiguration` `resolvConf` field value does not match `/run/systemd/resolve/resolv.conf` ([#107785](https://github.com/kubernetes/kubernetes/pull/107785), [@chendave](https://github.com/chendave))
- Kubeadm: fixed a bug when using `kubeadm init --dry-run` with certificate authority files (`ca.key` / `ca.crt`) present in `/etc/kubernetes/pki`) ([#108410](https://github.com/kubernetes/kubernetes/pull/108410), [@Haleygo](https://github.com/Haleygo))
- Kubeadm: fixed a bug where Windows nodes fail to join an IPv6 cluster due to preflight errors ([#108769](https://github.com/kubernetes/kubernetes/pull/108769), [@SataQiu](https://github.com/SataQiu))
- Kubeadm: fixed the bug that `kubeadm certs generate-csr` command does not remove duplicated SANs ([#107982](https://github.com/kubernetes/kubernetes/pull/107982), [@SataQiu](https://github.com/SataQiu))
- Kubelet now checks "NoExecute" taint/toleration before accepting pods, except for static pods. ([#101218](https://github.com/kubernetes/kubernetes/pull/101218), [@gjkim42](https://github.com/gjkim42))
- Metrics Server image bumped to v0.5.2 ([#106492](https://github.com/kubernetes/kubernetes/pull/106492), [@serathius](https://github.com/serathius))
- Modified command line errors (for example, `kubectl list` -> `unknown command`) that were printed as log message with escaped line breaks instead of a multi-line plain text, making the error hard to read. ([#107044](https://github.com/kubernetes/kubernetes/pull/107044), [@pohly](https://github.com/pohly))
- Modified log messages that were logged with `"v":0` in JSON output although they were debug messages with a higher verbosity. ([#106978](https://github.com/kubernetes/kubernetes/pull/106978), [@pohly](https://github.com/pohly))
- No ([#107769](https://github.com/kubernetes/kubernetes/pull/107769), [@liurupeng](https://github.com/liurupeng)) [SIG Cloud Provider and Windows]
- NodeRestriction admission: nodes are now allowed to update PersistentVolumeClaim status fields `resizeStatus` and `allocatedResources` when the `RecoverVolumeExpansionFailure` feature is enabled. ([#107686](https://github.com/kubernetes/kubernetes/pull/107686), [@gnufied](https://github.com/gnufied))
- Only extend token lifetimes when `--service-account-extend-token-expiration` is true and the requested token audiences are empty or exactly match all values for `--api-audiences`. ([#105954](https://github.com/kubernetes/kubernetes/pull/105954), [@jyotimahapatra](https://github.com/jyotimahapatra))
- Prevent kube-scheduler from nominating a Pod that was already scheduled to a node ([#109245](https://github.com/kubernetes/kubernetes/pull/109245), [@alculquicondor](https://github.com/alculquicondor))
- Prevent unnecessary `Endpoints` and `EndpointSlice` updates caused by `Pod ResourceVersion` change ([#108078](https://github.com/kubernetes/kubernetes/pull/108078), [@tnqn](https://github.com/tnqn))
- Print `<default>` as the value in case kubectl describe ingress shows `default-backend:80` when no default backend is present ([#108506](https://github.com/kubernetes/kubernetes/pull/108506), [@jlsong01](https://github.com/jlsong01))
- Publishing kube-proxy metrics for Windows kernel-mode ([#106581](https://github.com/kubernetes/kubernetes/pull/106581), [@knabben](https://github.com/knabben))
- Re-adds response status and headers on verbose kubectl responses ([#108505](https://github.com/kubernetes/kubernetes/pull/108505), [@rikatz](https://github.com/rikatz))
- Record requests rejected with 429 in the apiserver_request_total metric ([#108927](https://github.com/kubernetes/kubernetes/pull/108927), [@wojtek-t](https://github.com/wojtek-t))
- Removed validation if AppArmor profiles are loaded on the local node. This should be handled by the container runtime. ([#97966](https://github.com/kubernetes/kubernetes/pull/97966), [@saschagrunert](https://github.com/saschagrunert))
- Replace the url label of `rest_client_request_duration_seconds` and `rest_client_rate_limiter_duration_seconds` metrics with a host label to prevent cardinality explosions and keep only the useful information. This is a breaking change required for security reasons. ([#106539](https://github.com/kubernetes/kubernetes/pull/106539), [@dgrisonnet](https://github.com/dgrisonnet))
- Restored `NumPDBViolations` info of nodes, when `HTTPExtender ProcessPreemption`. This info will be used in subsequent filtering steps - `pickOneNodeForPreemption` ([#105853](https://github.com/kubernetes/kubernetes/pull/105853), [@caden2016](https://github.com/caden2016))
- Reverted graceful node shutdown to match 1.21 behavior of setting pods that have not yet successfully completed to "Failed" phase if the GracefulNodeShutdown feature is enabled in kubelet. The GracefulNodeShutdown feature is beta and must be explicitly configured via kubelet config to be enabled in 1.21+. This changes 1.22 and 1.23 behavior on node shutdown to match 1.21. If you do not want pods to be marked terminated on node shutdown in 1.22 and 1.23, disable the GracefulNodeShutdown feature. ([#106901](https://github.com/kubernetes/kubernetes/pull/106901), [@bobbypage](https://github.com/bobbypage))
- Reverts the CRI API version surfaced by dockershim to v1alpha2 ([#106803](https://github.com/kubernetes/kubernetes/pull/106803), [@saschagrunert](https://github.com/saschagrunert))
- Services with "internalTrafficPolicy: Local" now behave more like
  "externalTrafficPolicy: Local". Also, "internalTrafficPolicy: Local,
  externalTrafficPolicy: Cluster" is now implemented correctly. ([#106497](https://github.com/kubernetes/kubernetes/pull/106497), [@danwinship](https://github.com/danwinship))
- Sets JobTrackingWithFinalizers, a beta feature, as disabled by default, due to unresolved bug https://github.com/kubernetes/kubernetes/issues/109485 ([#109487](https://github.com/kubernetes/kubernetes/pull/109487), [@alculquicondor](https://github.com/alculquicondor))
- Skip re-allocate logic if pod is already removed to avoid panic ([#108831](https://github.com/kubernetes/kubernetes/pull/108831), [@waynepeking348](https://github.com/waynepeking348))
- The Service field `spec.internalTrafficPolicy` is no longer defaulted for Services when the type is `ExternalName`. The field is also dropped on read when the Service type is `ExternalName`. ([#104846](https://github.com/kubernetes/kubernetes/pull/104846), [@andrewsykim](https://github.com/andrewsykim))
- The `ServerSideFieldValidation` feature has been reverted to alpha for 1.24. ([#109271](https://github.com/kubernetes/kubernetes/pull/109271), [@liggitt](https://github.com/liggitt))
- The `TopologyAwareHints` feature gate is now enabled by default. This will allow users to opt-in to Topology Aware Hints by setting the `service.kubernetes.io/topology-aware-hints` on a Service. This will not affect any Services without that annotation set. ([#108747](https://github.com/kubernetes/kubernetes/pull/108747), [@robscott](https://github.com/robscott))
- The deprecated flag `--really-crash-for-testing` was removed. ([#101719](https://github.com/kubernetes/kubernetes/pull/101719), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- The kubelet no longer forcefully closes active connections on heartbeat failures, using the HTTP2 health check mechanism to detect broken connections. Users can force the previous behavior of the kubelet by setting the environment variable DISABLE_HTTP2. ([#108107](https://github.com/kubernetes/kubernetes/pull/108107), [@aojea](https://github.com/aojea))
- This code change fixes the bug that UDP services would trigger unnecessary LoadBalancer updates. The root cause is that a field not working for non-TCP protocols is considered.
  ref: https://github.com/kubernetes-sigs/cloud-provider-azure/pull/1090 ([#107981](https://github.com/kubernetes/kubernetes/pull/107981), [@lzhecheng](https://github.com/lzhecheng))
- Topology translation of in-tree vSphere volume to vSphere CSI. ([#108611](https://github.com/kubernetes/kubernetes/pull/108611), [@divyenpatel](https://github.com/divyenpatel))
- Updating kubelet permissions check for Windows nodes to see if process is elevated instead of checking if process owner is in Administrators group ([#108146](https://github.com/kubernetes/kubernetes/pull/108146), [@marosset](https://github.com/marosset))
- `apiserver`, if configured to reconcile the `kubernetes.default` service endpoints, checks if the configured Service IP range matches the apiserver public address IP family, and fails to start if not. ([#106721](https://github.com/kubernetes/kubernetes/pull/106721), [@aojea](https://github.com/aojea))
- `kubectl version` now fails when given extra arguments. ([#107967](https://github.com/kubernetes/kubernetes/pull/107967), [@jlsong01](https://github.com/jlsong01))

##### Other (Cleanup or Flake)

- Service session affinity timeout tests are no longer required for Kubernetes network plugin conformance due to variations in existing implementations. New conformance tests will be developed to better express conformance in future releases. ([#112806](https://github.com/kubernetes/kubernetes/pull/112806), [@dcbw](https://github.com/dcbw)) [SIG Architecture, Network and Testing]
- Kubelet now defaults to pulling the pause image from registry.k8s.io ([#114341](https://github.com/kubernetes/kubernetes/pull/114341), [@liggitt](https://github.com/liggitt)) [SIG Node]
- '`build/dependencies.yaml`: remove the dependency on Docker. With the dockershim removal, core Kubernetes no longer
  has to track the latest validated version of Docker.' ([#107607](https://github.com/kubernetes/kubernetes/pull/107607), [@neolit123](https://github.com/neolit123))
- API server's deprecated `--experimental-encryption-provider-config` flag is now removed. Adapt your machinery to use the `--encryption-provider-config` flag that is available since v1.13. ([#108423](https://github.com/kubernetes/kubernetes/pull/108423), [@ialidzhikov](https://github.com/ialidzhikov))
- API server's deprecated `--target-ram-mb` flag is now removed. ([#108457](https://github.com/kubernetes/kubernetes/pull/108457), [@ialidzhikov](https://github.com/ialidzhikov))
- Added PreemptionPolicy in PriorityClass describe ([#108701](https://github.com/kubernetes/kubernetes/pull/108701), [@denkensk](https://github.com/denkensk))
- Added an e2e test to verify that the cluster is not vulnerable to CVE-2021-29923 when using Services with IPs with leading zeros, note that this test is a necessary but not sufficient condition, all the components in the clusters that consume IPs addresses from the APIs MUST interpret them as decimal or discard them. ([#107552](https://github.com/kubernetes/kubernetes/pull/107552), [@aojea](https://github.com/aojea))
- Added an example for the `kubectl plugin list` command. ([#106600](https://github.com/kubernetes/kubernetes/pull/106600), [@bergerhoffer](https://github.com/bergerhoffer))
- Added details about preemption in the event for scheduling failed. ([#107775](https://github.com/kubernetes/kubernetes/pull/107775), [@denkensk](https://github.com/denkensk))
- Allow KUBE_TEST_REPO_LIST to be a remote url ([#108429](https://github.com/kubernetes/kubernetes/pull/108429), [@dims](https://github.com/dims))
- Client-go: if resetting the body fails before a retry, an error is now surfaced to the user. ([#109050](https://github.com/kubernetes/kubernetes/pull/109050), [@MadhavJivrajani](https://github.com/MadhavJivrajani))
- Deprecate apiserver_dropped_requests_total metric. The same data can be read from apiserver_request_terminations_total metric. ([#109018](https://github.com/kubernetes/kubernetes/pull/109018), [@wojtek-t](https://github.com/wojtek-t))
- Deprecated types in `k8s.io/apimachinery/util/clock`. Please use `k8s.io/utils/clock` instead. ([#106850](https://github.com/kubernetes/kubernetes/pull/106850), [@MadhavJivrajani](https://github.com/MadhavJivrajani))
- E2e tests wait for `kube-root-ca.crt` to be populated in namespaces for use with projected service account tokens, reducing delays starting those test pods and errors in the logs. ([#107763](https://github.com/kubernetes/kubernetes/pull/107763), [@smarterclayton](https://github.com/smarterclayton))
- Endpoints and EndpointSlice controllers no longer populate [resourceVersion of targetRef](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/object-reference/#ObjectReference) in Endpoints and EndpointSlices ([#108450](https://github.com/kubernetes/kubernetes/pull/108450), [@tnqn](https://github.com/tnqn))
- Fixed default config flags for `NewDefaultKubectlCommand`. ([#107131](https://github.com/kubernetes/kubernetes/pull/107131), [@jonnylangefeld](https://github.com/jonnylangefeld))
- Fixed documentation typo in cloud-provider. ([#106445](https://github.com/kubernetes/kubernetes/pull/106445), [@majst01](https://github.com/majst01))
- Fixed spelling of implemented in pkg/proxy/apis/config/types.go line 206 ([#106453](https://github.com/kubernetes/kubernetes/pull/106453), [@davidleitw](https://github.com/davidleitw))
- Improve error message when applying CRDs before the CRD exists in a cluster ([#107363](https://github.com/kubernetes/kubernetes/pull/107363), [@eddiezane](https://github.com/eddiezane))
- Improved algorithm for selecting `best` non-preferred hint in the TopologyManager ([#108154](https://github.com/kubernetes/kubernetes/pull/108154), [@klueska](https://github.com/klueska))
- Kube-proxy doesn't set the sysctl `net.ipv4.conf.all.route_localnet=1` if no IPv4 loopback address is selected by the `nodePortAddresses` configuration parameter. ([#107684](https://github.com/kubernetes/kubernetes/pull/107684), [@aojea](https://github.com/aojea))
- Kubeadm: all warning messages are printed to stderr instead of stdout. ([#107467](https://github.com/kubernetes/kubernetes/pull/107467), [@SataQiu](https://github.com/SataQiu))
- Kubeadm: handled the removal of dockershim related flags for new kubeadm clusters. If kubelet <1.24 is on the host, kubeadm >=1.24 can continue using the built-in dockershim in the kubelet if the user passes the `{Init|Join}Configuration.nodeRegistration.criSocket` value in the kubeadm configuration to be equal to `unix:///var/run/dockershim.sock` on Unix or `npipe:////./pipe/dockershim` on Windows. If kubelet version >=1.24 is on the host, kubeadm >=1.24 will treat all container runtimes as "remote" using the kubelet flags `--container-runtime=remote --container-runtime-endpoint=scheme://some/path`. The special management for kubelet <1.24 will be removed in kubeadm 1.25. ([#106973](https://github.com/kubernetes/kubernetes/pull/106973), [@neolit123](https://github.com/neolit123))
- Kubeadm: make sure that `kubeadm init/join` always use a URL scheme (unix:// on Linux and npipe:// on Windows) when passing a value to the `--container-runtime-endpoint` kubelet flag. This flag's value is taken from the kubeadm configuration `criSocket` field or the `--cri-socket` CLI flag. Automatically add a missing URL scheme to the user configuration in memory, but warn them that they should also update their configuration on disk manually. During `kubeadm upgrade apply/node` mutate the `/var/lib/kubelet/kubeadm-flags.env` file on disk and the `kubeadm.alpha.kubernetes.io/cri-socket` annotation Node object if needed. These automatic actions are temporary and will be removed in a future release. In the future the kubelet may not support CRI endpoints without an URL scheme. ([#107295](https://github.com/kubernetes/kubernetes/pull/107295), [@neolit123](https://github.com/neolit123))
- Kubeadm: remove the `IPv6DualStack` feature gate. The feature has been GA and locked to enabled since 1.23. ([#106648](https://github.com/kubernetes/kubernetes/pull/106648), [@calvin0327](https://github.com/calvin0327))
- Kubeadm: removed the deprecated `output/v1alpha1` API used for machine readable output by some kubeadm commands. In 1.23 kubeadm started using the newer version `output/v1alpha2` for the same purpose. ([#107468](https://github.com/kubernetes/kubernetes/pull/107468), [@neolit123](https://github.com/neolit123))
- Kubeadm: removed the restriction that the `ca.crt` can only contain one certificate. If there is more than one certificate in the `ca.crt` file, kubeadm will pick the first one by default. ([#107327](https://github.com/kubernetes/kubernetes/pull/107327), [@SataQiu](https://github.com/SataQiu))
- Kubectl stack traces now only print at verbose `-v=99` and not `-v=6` ([#108053](https://github.com/kubernetes/kubernetes/pull/108053), [@eddiezane](https://github.com/eddiezane))
- Kubectl: restored `--dry-run`, `--dry-run=true`, and `--dry-run=false` for compatibility with pre-1.23 invocations. ([#107003](https://github.com/kubernetes/kubernetes/pull/107003), [@julianvmodesto](https://github.com/julianvmodesto))
- Kubelet config validation error messages are updated. ([#105360](https://github.com/kubernetes/kubernetes/pull/105360), [@shuheiktgw](https://github.com/shuheiktgw))
- Kubernetes e2e framework will use the url `invalid.registry.k8s.io/invalid` instead `invalid.com/invalid` for test that use an invalid registry. ([#107455](https://github.com/kubernetes/kubernetes/pull/107455), [@aojea](https://github.com/aojea))
- Marked kubelet `--container-runtime-endpoint` and `--image-service-endpoint` CLI flags as stable. ([#106954](https://github.com/kubernetes/kubernetes/pull/106954), [@saschagrunert](https://github.com/saschagrunert))
- Migrate `volume/csi/csi-client.go` logs to structured logging. ([#99441](https://github.com/kubernetes/kubernetes/pull/99441), [@CKchen0726](https://github.com/CKchen0726))
- Migrate statefulset files to structured logging ([#106109](https://github.com/kubernetes/kubernetes/pull/106109), [@h4ghhh](https://github.com/h4ghhh))
- Refactor kubelet command line for enabling features and "drop `RuntimeClass` feature gate" if present. Note that this feature has been on by default since 1.14 and was GA'ed in 1.20. ([#106882](https://github.com/kubernetes/kubernetes/pull/106882), [@cyclinder](https://github.com/cyclinder))
- Remove deprecated `--serviceaccount`, `--hostport`, `--requests` and `--limits` from kubectl run. ([#108820](https://github.com/kubernetes/kubernetes/pull/108820), [@mozillazg](https://github.com/mozillazg))
- Remove support for `node-expansion` between `node-stage` and `node-publish` ([#108614](https://github.com/kubernetes/kubernetes/pull/108614), [@gnufied](https://github.com/gnufied))
- Removed deprecated `generator` and `container-port` flags ([#106824](https://github.com/kubernetes/kubernetes/pull/106824), [@lauchokyip](https://github.com/lauchokyip))
- Removed kubelet `--non-masquerade-cidr` deprecated CLI flag ([#107096](https://github.com/kubernetes/kubernetes/pull/107096), [@hakman](https://github.com/hakman))
- Rename unschedulableQ to unschedulablePods ([#108919](https://github.com/kubernetes/kubernetes/pull/108919), [@denkensk](https://github.com/denkensk))
- SPDY transport in client-go will no longer follow redirects. ([#108531](https://github.com/kubernetes/kubernetes/pull/108531), [@tallclair](https://github.com/tallclair))
- ServerResources was deprecated in February 2019 (https://github.com/kubernetes/kubernetes/commit/618050e) and now it's being removed and ServerGroupsAndResources is suggested to be used instead ([#107180](https://github.com/kubernetes/kubernetes/pull/107180), [@ardaguclu](https://github.com/ardaguclu))
- The API server's deprecated `--deserialization-cache-size` flag is now removed. ([#108448](https://github.com/kubernetes/kubernetes/pull/108448), [@ialidzhikov](https://github.com/ialidzhikov))
- The `--container-runtime` kubelet flag is deprecated and will be removed in future releases. ([#107094](https://github.com/kubernetes/kubernetes/pull/107094), [@adisky](https://github.com/adisky))
- The `WarningHeaders` feature gate that is GA since v1.22 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#108394](https://github.com/kubernetes/kubernetes/pull/108394), [@ialidzhikov](https://github.com/ialidzhikov))
- The `e2e.test` binary supports a new `--kubelet-root` parameter to override the default `/var/lib/kubelet` path. CSI storage tests use this. ([#108253](https://github.com/kubernetes/kubernetes/pull/108253), [@pohly](https://github.com/pohly))
- The fluentd-elasticsearch addon is no longer included in the cluster directory. It is available from https://github.com/kubernetes-sigs/instrumentation-addons/tree/master/fluentd-elasticsearch. ([#107553](https://github.com/kubernetes/kubernetes/pull/107553), [@liggitt](https://github.com/liggitt))
- The scheduler framework option `runAllFilters` is removed. ([#108829](https://github.com/kubernetes/kubernetes/pull/108829), [@kerthcet](https://github.com/kerthcet))
- Updated cri-tools to [v1.23.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.23.0). ([#107604](https://github.com/kubernetes/kubernetes/pull/107604), [@saschagrunert](https://github.com/saschagrunert))
- Updated runc to 1.1.0 and updated cadvisor to 0.44.0 ([#109029](https://github.com/kubernetes/kubernetes/pull/109029), [@ehashman](https://github.com/ehashman))
- Updated runc to 1.1.1 ([#109104](https://github.com/kubernetes/kubernetes/pull/109104), [@kolyshkin](https://github.com/kolyshkin))
- Updated the error message to not use the `--max-resource-write-bytes` & `--json-patch-max-copy-bytes` string. ([#106875](https://github.com/kubernetes/kubernetes/pull/106875), [@warmchang](https://github.com/warmchang))
- Users who look at iptables dumps will see some changes in the naming and structure of rules. ([#109060](https://github.com/kubernetes/kubernetes/pull/109060), [@thockin](https://github.com/thockin))
- Windows Pause no longer has support for SAC releases 1903, 1909, 2004. Windows image support is now Ltcs 2019 (1809), 20H2, LTSC 2022 ([#107056](https://github.com/kubernetes/kubernetes/pull/107056), [@jsturtevant](https://github.com/jsturtevant))
- [k8s.io/utils/clock]: IntervalClock is now deprecated in favour of SimpleIntervalClock ([#108059](https://github.com/kubernetes/kubernetes/pull/108059), [@RaghavRoy145](https://github.com/RaghavRoy145))
- `kube-addon-manager` image version is bumped to 9.1.6 ([#108341](https://github.com/kubernetes/kubernetes/pull/108341), [@zshihang](https://github.com/zshihang))
- Add SourceVolumeMode field to VolumeSnapshotContents. Documentation for this alpha feature is pending. ([#665](https://github.com/kubernetes-csi/external-snapshotter/pull/665), [@RaunakShah](https://github.com/RaunakShah))
- Update snapshotter module to v6 and client module to v5. Documentation for this alpha feature is pending. ([#670],(https://github.com/kubernetes-csi/external-snapshotter/pull/670), [@RaunakShah](https://github.com/RaunakShah))

##### Uncategorized

- Deprecate kubectl version long output, will be replaced with kubectl version ` --short`. Users requiring full output should use `--output=yaml|json` instead. ([#108987](https://github.com/kubernetes/kubernetes/pull/108987), [@soltysh](https://github.com/soltysh))


#### Dependencies

##### Added
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/blang/semver/v4: [v4.0.0](https://github.com/blang/semver/v4/tree/v4.0.0)
- github.com/google/gnostic: [v0.5.7-v3refs](https://github.com/google/gnostic/tree/v0.5.7-v3refs)

##### Changed
- golang.org/x/net: 1e63c2f → v0.7.0
- golang.org/x/sys: v0.3.0 → v0.5.0
- golang.org/x/term: v0.3.0 → v0.5.0
- golang.org/x/text: v0.5.0 → v0.7.0
- github.com/google/cel-go: [v0.10.2 → v0.10.4](https://github.com/google/cel-go/compare/v0.10.2...v0.10.4)
- k8s.io/system-validators: v1.7.0 → v1.8.0
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.33 → v0.0.35
- sigs.k8s.io/structured-merge-diff/v4: v4.2.1 → v4.2.3
- github.com/yuin/goldmark: [v1.4.1 → v1.4.13](https://github.com/yuin/goldmark/compare/v1.4.1...v1.4.13)
- golang.org/x/mod: 9b9b3d8 → 86c51ed
- golang.org/x/sync: 036812b → 886fb93
- golang.org/x/tools: 897bd77 → v0.1.12
- github.com/stretchr/objx: [v0.2.0 → v0.4.0](https://github.com/stretchr/objx/compare/v0.2.0...v0.4.0)
- github.com/stretchr/testify: [v1.7.0 → v1.8.0](https://github.com/stretchr/testify/compare/v1.7.0...v1.8.0)
- go.uber.org/goleak: v1.1.10 → v1.2.0
- golang.org/x/lint: 6edffad → 83fdc39
- gopkg.in/yaml.v3: 496545a → v3.0.1
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.30 → v0.0.33
- github.com/cespare/xxhash/v2: [v2.1.1 → v2.1.2](https://github.com/cespare/xxhash/v2/compare/v2.1.1...v2.1.2)
- github.com/checkpoint-restore/go-criu/v5: [v5.0.0 → v5.3.0](https://github.com/checkpoint-restore/go-criu/v5/compare/v5.0.0...v5.3.0)
- github.com/cilium/ebpf: [v0.6.2 → v0.7.0](https://github.com/cilium/ebpf/compare/v0.6.2...v0.7.0)
- github.com/containerd/console: [v1.0.2 → v1.0.3](https://github.com/containerd/console/compare/v1.0.2...v1.0.3)
- github.com/containerd/containerd: [v1.4.11 → v1.4.12](https://github.com/containerd/containerd/compare/v1.4.11...v1.4.12)
- github.com/cpuguy83/go-md2man/v2: [v2.0.0 → v2.0.1](https://github.com/cpuguy83/go-md2man/v2/compare/v2.0.0...v2.0.1)
- github.com/cyphar/filepath-securejoin: [v0.2.2 → v0.2.3](https://github.com/cyphar/filepath-securejoin/compare/v0.2.2...v0.2.3)
- github.com/docker/distribution: [v2.7.1+incompatible → v2.8.1+incompatible](https://github.com/docker/distribution/compare/v2.7.1...v2.8.1)
- github.com/docker/docker: [v20.10.7+incompatible → v20.10.12+incompatible](https://github.com/docker/docker/compare/v20.10.7...v20.10.12)
- github.com/godbus/dbus/v5: [v5.0.4 → v5.0.6](https://github.com/godbus/dbus/v5/compare/v5.0.4...v5.0.6)
- github.com/golang/mock: [v1.5.0 → v1.6.0](https://github.com/golang/mock/compare/v1.5.0...v1.6.0)
- github.com/google/cadvisor: [v0.43.0 → v0.44.1](https://github.com/google/cadvisor/compare/v0.43.0...v0.44.1)
- github.com/moby/sys/mountinfo: [v0.4.1 → v0.6.0](https://github.com/moby/sys/mountinfo/compare/v0.4.1...v0.6.0)
- github.com/moby/term: [9d4ed18 → 3f7ff69](https://github.com/moby/term/compare/9d4ed18...3f7ff69)
- github.com/opencontainers/image-spec: [v1.0.1 → v1.0.2](https://github.com/opencontainers/image-spec/compare/v1.0.1...v1.0.2)
- github.com/opencontainers/runc: [v1.0.2 → v1.1.1](https://github.com/opencontainers/runc/compare/v1.0.2...v1.1.1)
- github.com/opencontainers/selinux: [v1.8.2 → v1.10.0](https://github.com/opencontainers/selinux/compare/v1.8.2...v1.10.0)
- github.com/prometheus/client_golang: [v1.11.0 → v1.12.1](https://github.com/prometheus/client_golang/compare/v1.11.0...v1.12.1)
- github.com/prometheus/common: [v0.28.0 → v0.32.1](https://github.com/prometheus/common/compare/v0.28.0...v0.32.1)
- github.com/prometheus/procfs: [v0.6.0 → v0.7.3](https://github.com/prometheus/procfs/compare/v0.6.0...v0.7.3)
- github.com/russross/blackfriday/v2: [v2.0.1 → v2.1.0](https://github.com/russross/blackfriday/v2/compare/v2.0.1...v2.1.0)
- github.com/seccomp/libseccomp-golang: [v0.9.1 → 3879420](https://github.com/seccomp/libseccomp-golang/compare/v0.9.1...3879420)
- github.com/spf13/cobra: [v1.2.1 → v1.4.0](https://github.com/spf13/cobra/compare/v1.2.1...v1.4.0)
- go.etcd.io/etcd/api/v3: v3.5.0 → v3.5.1
- go.etcd.io/etcd/client/pkg/v3: v3.5.0 → v3.5.1
- go.etcd.io/etcd/client/v3: v3.5.0 → v3.5.1
- golang.org/x/crypto: 32db794 → 8634188
- golang.org/x/oauth2: 2bc19b1 → d3ed0bb
- golang.org/x/time: 1f47c86 → 90d013b
- google.golang.org/genproto: fe13028 → 42d7afd
- k8s.io/gengo: 485abfe → c02415c
- k8s.io/klog/v2: v2.30.0 → v2.60.1
- k8s.io/kube-openapi: e816edb → 3ee0da9
- k8s.io/utils: cb0fa31 → 3a6ce19
- sigs.k8s.io/json: c049b76 → 9f7c6b3
- sigs.k8s.io/kustomize/api: v0.10.1 → v0.11.4
- sigs.k8s.io/kustomize/cmd/config: v0.10.2 → v0.10.6
- sigs.k8s.io/kustomize/kustomize/v4: v4.4.1 → v4.5.4
- sigs.k8s.io/kustomize/kyaml: v0.13.0 → v0.13.6

##### Removed
- cloud.google.com/go/firestore: v1.1.0
- github.com/armon/go-metrics: [f0300d1](https://github.com/armon/go-metrics/tree/f0300d1)
- github.com/armon/go-radix: [7fddfc3](https://github.com/armon/go-radix/tree/7fddfc3)
- github.com/bgentry/speakeasy: [v0.1.0](https://github.com/bgentry/speakeasy/tree/v0.1.0)
- github.com/bits-and-blooms/bitset: [v1.2.0](https://github.com/bits-and-blooms/bitset/tree/v1.2.0)
- github.com/bketelsen/crypt: [v0.0.4](https://github.com/bketelsen/crypt/tree/v0.0.4)
- github.com/containernetworking/cni: [v0.8.1](https://github.com/containernetworking/cni/tree/v0.8.1)
- github.com/fatih/color: [v1.7.0](https://github.com/fatih/color/tree/v1.7.0)
- github.com/googleapis/gnostic: [v0.5.5](https://github.com/googleapis/gnostic/tree/v0.5.5)
- github.com/hashicorp/consul/api: [v1.1.0](https://github.com/hashicorp/consul/api/tree/v1.1.0)
- github.com/hashicorp/consul/sdk: [v0.1.1](https://github.com/hashicorp/consul/sdk/tree/v0.1.1)
- github.com/hashicorp/errwrap: [v1.0.0](https://github.com/hashicorp/errwrap/tree/v1.0.0)
- github.com/hashicorp/go-cleanhttp: [v0.5.1](https://github.com/hashicorp/go-cleanhttp/tree/v0.5.1)
- github.com/hashicorp/go-immutable-radix: [v1.0.0](https://github.com/hashicorp/go-immutable-radix/tree/v1.0.0)
- github.com/hashicorp/go-msgpack: [v0.5.3](https://github.com/hashicorp/go-msgpack/tree/v0.5.3)
- github.com/hashicorp/go-multierror: [v1.0.0](https://github.com/hashicorp/go-multierror/tree/v1.0.0)
- github.com/hashicorp/go-rootcerts: [v1.0.0](https://github.com/hashicorp/go-rootcerts/tree/v1.0.0)
- github.com/hashicorp/go-sockaddr: [v1.0.0](https://github.com/hashicorp/go-sockaddr/tree/v1.0.0)
- github.com/hashicorp/go-syslog: [v1.0.0](https://github.com/hashicorp/go-syslog/tree/v1.0.0)
- github.com/hashicorp/go-uuid: [v1.0.1](https://github.com/hashicorp/go-uuid/tree/v1.0.1)
- github.com/hashicorp/go.net: [v0.0.1](https://github.com/hashicorp/go.net/tree/v0.0.1)
- github.com/hashicorp/golang-lru: [v0.5.0](https://github.com/hashicorp/golang-lru/tree/v0.5.0)
- github.com/hashicorp/hcl: [v1.0.0](https://github.com/hashicorp/hcl/tree/v1.0.0)
- github.com/hashicorp/logutils: [v1.0.0](https://github.com/hashicorp/logutils/tree/v1.0.0)
- github.com/hashicorp/mdns: [v1.0.0](https://github.com/hashicorp/mdns/tree/v1.0.0)
- github.com/hashicorp/memberlist: [v0.1.3](https://github.com/hashicorp/memberlist/tree/v0.1.3)
- github.com/hashicorp/serf: [v0.8.2](https://github.com/hashicorp/serf/tree/v0.8.2)
- github.com/magiconair/properties: [v1.8.5](https://github.com/magiconair/properties/tree/v1.8.5)
- github.com/mattn/go-colorable: [v0.0.9](https://github.com/mattn/go-colorable/tree/v0.0.9)
- github.com/mattn/go-isatty: [v0.0.3](https://github.com/mattn/go-isatty/tree/v0.0.3)
- github.com/miekg/dns: [v1.0.14](https://github.com/miekg/dns/tree/v1.0.14)
- github.com/mitchellh/cli: [v1.0.0](https://github.com/mitchellh/cli/tree/v1.0.0)
- github.com/mitchellh/go-homedir: [v1.0.0](https://github.com/mitchellh/go-homedir/tree/v1.0.0)
- github.com/mitchellh/go-testing-interface: [v1.0.0](https://github.com/mitchellh/go-testing-interface/tree/v1.0.0)
- github.com/mitchellh/gox: [v0.4.0](https://github.com/mitchellh/gox/tree/v0.4.0)
- github.com/mitchellh/iochan: [v1.0.0](https://github.com/mitchellh/iochan/tree/v1.0.0)
- github.com/pascaldekloe/goe: [57f6aae](https://github.com/pascaldekloe/goe/tree/57f6aae)
- github.com/pelletier/go-toml: [v1.9.3](https://github.com/pelletier/go-toml/tree/v1.9.3)
- github.com/posener/complete: [v1.1.1](https://github.com/posener/complete/tree/v1.1.1)
- github.com/ryanuber/columnize: [9b3edd6](https://github.com/ryanuber/columnize/tree/9b3edd6)
- github.com/sean-/seed: [e2103e2](https://github.com/sean-/seed/tree/e2103e2)
- github.com/shurcooL/sanitized_anchor_name: [v1.0.0](https://github.com/shurcooL/sanitized_anchor_name/tree/v1.0.0)
- github.com/spf13/cast: [v1.3.1](https://github.com/spf13/cast/tree/v1.3.1)
- github.com/spf13/jwalterweatherman: [v1.1.0](https://github.com/spf13/jwalterweatherman/tree/v1.1.0)
- github.com/spf13/viper: [v1.8.1](https://github.com/spf13/viper/tree/v1.8.1)
- github.com/subosito/gotenv: [v1.2.0](https://github.com/subosito/gotenv/tree/v1.2.0)
- gopkg.in/ini.v1: v1.62.0



### containerlinux [3510.2.0](https://www.flatcar-linux.org/releases/#release-3510.2.0)

_Changes since **Stable 3374.2.5**_

#### Security fixes:

- Linux ([CVE-2022-2196](https://nvd.nist.gov/vuln/detail/CVE-2022-2196), [CVE-2022-27672](https://nvd.nist.gov/vuln/detail/CVE-2022-27672), [CVE-2022-3707](https://nvd.nist.gov/vuln/detail/CVE-2022-3707), [CVE-2023-1078](https://nvd.nist.gov/vuln/detail/CVE-2023-1078), [CVE-2023-1281](https://nvd.nist.gov/vuln/detail/CVE-2023-1281), [CVE-2023-1513](https://nvd.nist.gov/vuln/detail/CVE-2023-1513), [CVE-2023-26545](https://nvd.nist.gov/vuln/detail/CVE-2023-26545))
- bind tools ([CVE-2022-2795](https://nvd.nist.gov/vuln/detail/CVE-2022-2795), [CVE-2022-2881](https://nvd.nist.gov/vuln/detail/CVE-2022-2881), [CVE-2022-2906](https://nvd.nist.gov/vuln/detail/CVE-2022-2906), [CVE-2022-3080](https://nvd.nist.gov/vuln/detail/CVE-2022-3080), [CVE-2022-38177](https://nvd.nist.gov/vuln/detail/CVE-2022-38177), [CVE-2022-38178](https://nvd.nist.gov/vuln/detail/CVE-2022-38178))
- binutils ([CVE-2022-38126](https://nvd.nist.gov/vuln/detail/CVE-2022-38126), [CVE-2022-38127](https://nvd.nist.gov/vuln/detail/CVE-2022-38127))
- containerd ([CVE-2022-23471](https://nvd.nist.gov/vuln/detail/CVE-2022-23471))
- cpio ([CVE-2021-38185](https://nvd.nist.gov/vuln/detail/CVE-2021-38185))
- curl ([CVE-2022-35252](https://nvd.nist.gov/vuln/detail/CVE-2022-35252), [CVE-2022-43551](https://nvd.nist.gov/vuln/detail/CVE-2022-43551), [CVE-2022-43552](https://nvd.nist.gov/vuln/detail/CVE-2022-43552),[CVE-2022-32221](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-35260](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-42915](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-42916](https://nvd.nist.gov/vuln/detail/CVE-2022-32221))
- dbus ([CVE-2022-42010](https://nvd.nist.gov/vuln/detail/CVE-2022-42010), [CVE-2022-42011](https://nvd.nist.gov/vuln/detail/CVE-2022-42011), [CVE-2022-42012](https://nvd.nist.gov/vuln/detail/CVE-2022-42012))
- git ([CVE-2022-39253](https://nvd.nist.gov/vuln/detail/CVE-2022-39253), [CVE-2022-39260](https://nvd.nist.gov/vuln/detail/CVE-2022-39260), [CVE-2022-23521](https://nvd.nist.gov/vuln/detail/CVE-2022-23521), [CVE-2022-41903](https://nvd.nist.gov/vuln/detail/CVE-2022-41903))
- glib ([fixes to normal form handling in GVariant](https://discourse.gnome.org/t/multiple-fixes-for-gvariant-normalisation-issues-in-glib/12835))
- Go ([CVE-2022-41717](https://nvd.nist.gov/vuln/detail/CVE-2022-41717))
- libarchive ([CVE-2022-36227](https://nvd.nist.gov/vuln/detail/CVE-2022-36227))
- libksba ([CVE-2022-47629](https://nvd.nist.gov/vuln/detail/CVE-2022-47629), [CVE-2022-3515](https://nvd.nist.gov/vuln/detail/CVE-2022-3515))
- libxml2 ([CVE-2022-40303](https://nvd.nist.gov/vuln/detail/CVE-2022-40303), [CVE-2022-40304](https://nvd.nist.gov/vuln/detail/CVE-2022-40304))
- logrotate ([CVE-2022-1348](https://nvd.nist.gov/vuln/detail/CVE-2022-1348))
- multipath-tools ([CVE-2022-41973](https://nvd.nist.gov/vuln/detail/CVE-2022-41973), [CVE-2022-41974](https://nvd.nist.gov/vuln/detail/CVE-2022-41974))
- sudo ([CVE-2023-22809](https://nvd.nist.gov/vuln/detail/CVE-2023-22809), [CVE-2022-43995](https://nvd.nist.gov/vuln/detail/CVE-2022-43995))
- systemd ([CVE-2022-3821](https://nvd.nist.gov/vuln/detail/CVE-2022-3821), [CVE-2022-4415](https://nvd.nist.gov/vuln/detail/CVE-2022-4415))
- vim ([CVE-2023-0049](https://nvd.nist.gov/vuln/detail/CVE-2023-0049), [CVE-2023-0051](https://nvd.nist.gov/vuln/detail/CVE-2023-0051), [CVE-2023-0054](https://nvd.nist.gov/vuln/detail/CVE-2023-0054), [CVE-2022-3705](https://nvd.nist.gov/vuln/detail/CVE-2022-3705), [CVE-2022-3491](https://nvd.nist.gov/vuln/detail/CVE-2022-3491), [CVE-2022-3520](https://nvd.nist.gov/vuln/detail/CVE-2022-3520), [CVE-2022-3591](https://nvd.nist.gov/vuln/detail/CVE-2022-3591), [CVE-2022-4141](https://nvd.nist.gov/vuln/detail/CVE-2022-4141), [CVE-2022-4292](https://nvd.nist.gov/vuln/detail/CVE-2022-4292), [CVE-2022-4293](https://nvd.nist.gov/vuln/detail/CVE-2022-4293),[CVE-2022-1725](https://nvd.nist.gov/vuln/detail/CVE-2022-1725), [CVE-2022-3234](https://nvd.nist.gov/vuln/detail/CVE-2022-3234), [CVE-2022-3235](https://nvd.nist.gov/vuln/detail/CVE-2022-3235), [CVE-2022-3278](https://nvd.nist.gov/vuln/detail/CVE-2022-3278), [CVE-2022-3256](https://nvd.nist.gov/vuln/detail/CVE-2022-3256), [CVE-2022-3296](https://nvd.nist.gov/vuln/detail/CVE-2022-3296), [CVE-2022-3297](https://nvd.nist.gov/vuln/detail/CVE-2022-3297), [CVE-2022-3324](https://nvd.nist.gov/vuln/detail/CVE-2022-3324), [CVE-2022-3352](https://nvd.nist.gov/vuln/detail/CVE-2022-3352), [CVE-2022-2042](https://nvd.nist.gov/vuln/detail/CVE-2022-2042), [CVE-2022-2124](https://nvd.nist.gov/vuln/detail/CVE-2022-2124), [CVE-2022-2125](https://nvd.nist.gov/vuln/detail/CVE-2022-2125), [CVE-2022-2126](https://nvd.nist.gov/vuln/detail/CVE-2022-2126), [CVE-2022-2129](https://nvd.nist.gov/vuln/detail/CVE-2022-2129), [CVE-2022-2175](https://nvd.nist.gov/vuln/detail/CVE-2022-2175), [CVE-2022-2182](https://nvd.nist.gov/vuln/detail/CVE-2022-2182), [CVE-2022-2183](https://nvd.nist.gov/vuln/detail/CVE-2022-2183), [CVE-2022-2206](https://nvd.nist.gov/vuln/detail/CVE-2022-2206), [CVE-2022-2207](https://nvd.nist.gov/vuln/detail/CVE-2022-2207), [CVE-2022-2208](https://nvd.nist.gov/vuln/detail/CVE-2022-2208), [CVE-2022-2210](https://nvd.nist.gov/vuln/detail/CVE-2022-2210), [CVE-2022-2231](https://nvd.nist.gov/vuln/detail/CVE-2022-2231), [CVE-2022-2257](https://nvd.nist.gov/vuln/detail/CVE-2022-2257), [CVE-2022-2264](https://nvd.nist.gov/vuln/detail/CVE-2022-2264), [CVE-2022-2284](https://nvd.nist.gov/vuln/detail/CVE-2022-2284), [CVE-2022-2285](https://nvd.nist.gov/vuln/detail/CVE-2022-2285), [CVE-2022-2286](https://nvd.nist.gov/vuln/detail/CVE-2022-2286), [CVE-2022-2287](https://nvd.nist.gov/vuln/detail/CVE-2022-2287), [CVE-2022-2288](https://nvd.nist.gov/vuln/detail/CVE-2022-2288), [CVE-2022-2289](https://nvd.nist.gov/vuln/detail/CVE-2022-2289), [CVE-2022-2304](https://nvd.nist.gov/vuln/detail/CVE-2022-2304), [CVE-2022-2343](https://nvd.nist.gov/vuln/detail/CVE-2022-2343), [CVE-2022-2344](https://nvd.nist.gov/vuln/detail/CVE-2022-2344), [CVE-2022-2345](https://nvd.nist.gov/vuln/detail/CVE-2022-2345), [CVE-2022-2522](https://nvd.nist.gov/vuln/detail/CVE-2022-2522), [CVE-2022-2816](https://nvd.nist.gov/vuln/detail/CVE-2022-2816), [CVE-2022-2817](https://nvd.nist.gov/vuln/detail/CVE-2022-2817), [CVE-2022-2819](https://nvd.nist.gov/vuln/detail/CVE-2022-2819), [CVE-2022-2845](https://nvd.nist.gov/vuln/detail/CVE-2022-2845), [CVE-2022-2849](https://nvd.nist.gov/vuln/detail/CVE-2022-2849), [CVE-2022-2862](https://nvd.nist.gov/vuln/detail/CVE-2022-2862), [CVE-2022-2874](https://nvd.nist.gov/vuln/detail/CVE-2022-2874), [CVE-2022-2889](https://nvd.nist.gov/vuln/detail/CVE-2022-2889), [CVE-2022-2923](https://nvd.nist.gov/vuln/detail/CVE-2022-2923), [CVE-2022-2946](https://nvd.nist.gov/vuln/detail/CVE-2022-2946), [CVE-2022-2980](https://nvd.nist.gov/vuln/detail/CVE-2022-2980), [CVE-2022-2982](https://nvd.nist.gov/vuln/detail/CVE-2022-2982), [CVE-2022-3016](https://nvd.nist.gov/vuln/detail/CVE-2022-3016), [CVE-2022-3099](https://nvd.nist.gov/vuln/detail/CVE-2022-3099), [CVE-2022-3134](https://nvd.nist.gov/vuln/detail/CVE-2022-3134), [CVE-2022-3153](https://nvd.nist.gov/vuln/detail/CVE-2022-3153))
- SDK: Python ([CVE-2015-20107](https://nvd.nist.gov/vuln/detail/CVE-2015-20107), [CVE-2020-10735](https://nvd.nist.gov/vuln/detail/CVE-2020-10735), [CVE-2021-3654](https://nvd.nist.gov/vuln/detail/CVE-2021-3654), [CVE-2022-37454](https://nvd.nist.gov/vuln/detail/CVE-2022-37454), [CVE-2022-42919](https://nvd.nist.gov/vuln/detail/CVE-2022-42919), [CVE-2022-45061](https://nvd.nist.gov/vuln/detail/CVE-2022-45061))
- SDK: qemu ([CVE-2022-4172](https://nvd.nist.gov/vuln/detail/CVE-2022-4172), [CVE-2020-14394](https://nvd.nist.gov/vuln/detail/CVE-2020-14394), [CVE-2022-0216](https://nvd.nist.gov/vuln/detail/CVE-2022-0216), [CVE-2022-35414](https://nvd.nist.gov/vuln/detail/CVE-2022-35414), [CVE-2022-3872](https://nvd.nist.gov/vuln/detail/CVE-2022-3872))
- SDK: rust ([CVE-2022-46176](https://nvd.nist.gov/vuln/detail/CVE-2022-46176), [CVE-2022-36113](https://nvd.nist.gov/vuln/detail/CVE-2022-36113), [CVE-2022-36114](https://nvd.nist.gov/vuln/detail/CVE-2022-36114))

#### Bug fixes:

- Added back Ignition support for Vagrant ([coreos-overlay#2351](https://github.com/flatcar/coreos-overlay/pull/2351))
- Added support for hardware security keys in update-ssh-keys ([update-ssh-keys#7](https://github.com/flatcar/update-ssh-keys/pull/7))
- Enabled IOMMU on arm64 kernels, the lack of which prevented some systems from booting ([coreos-overlay#2235](https://github.com/flatcar/coreos-overlay/pull/2235))
- Fixed a regression (in Alpha/Beta) where machines failed to boot if they didnt have the `core` user or group in `/etc/passwd` or `/etc/group` ([baselayout#26](https://github.com/flatcar/baselayout/pull/26))
- Fix "ext4 deadlock under heavy I/O load" kernel issue. The patch for this is included provisionally while we wait for it to be merged upstream ([Flatcar#847](https://github.com/flatcar/Flatcar/issues/847), [coreos-overlay#2315](https://github.com/flatcar/coreos-overlay/pull/2315))
- Restored the support to specify OEM partition files in Ignition when `/usr/share/oem` is given as initrd mount point ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))
- The rootfs setup in the initrd now runs systemd-tmpfiles on every boot, not only when Ignition runs, to fix a dbus failure due to missing files ([Flatcar#944](https://github.com/flatcar/Flatcar/issues/944))

#### Changes:

- Added `CONFIG_NF_CONNTRACK_BRIDGE` (for nf_conntrack_bridge) and `CONFIG_NFT_BRIDGE_META` (for nft_meta_bridge) to the kernel config to allow using conntrack rules for bridges in nftables and to match on bridge interface names ([coreos-overlay#2207](https://github.com/flatcar/coreos-overlay/pull/2207))
- Added new image signing pub key to `flatcar-install`, needed for download verification of releases built from July 2023 onwards, if you have copies of `flatcar-install` or the image signing pub key, you need to update them as well ([init#92](https://github.com/flatcar/init/pull/92))
- Change CONFIG_WIREGUARD kernel option to module to save space on boot partition ([coreos-overlay#2239](https://github.com/flatcar/coreos-overlay/pull/2239))
- Disable several arch specific arm64 kernel config options for unsupported platforms to save space on boot partition ([coreos-overlay#2239](https://github.com/flatcar/coreos-overlay/pull/2239))
- Specifying the OEM filesystem in Ignition to write files to `/usr/share/oem` is not needed anymore ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))
- Switched from `--strip-unneeded` to `--strip-debug` when installing kernel modules, which makes kernel stacktraces more accurate and makes debugging issues easier ([coreos-overlay#2196](https://github.com/flatcar/coreos-overlay/pull/2196))
- The flatcar-update tool got two new flags to customize ports used on the host while updating flatcar ([init#81](https://github.com/flatcar/init/pull/81))
- Toolbox now uses containerd to download and mount the image ([toolbox#7](https://github.com/flatcar/toolbox/pull/7))
- Add qemu-guest-agent to all amd64 images, it will be automatically enabled when qemu-ga virtio-port is detected ([coreos-overlay#2240](https://github.com/flatcar/coreos-overlay/pull/2240), [portage-stable#373](https://github.com/flatcar/portage-stable/pull/373))

#### Updates:

- Linux ([5.15.98](https://lwn.net/Articles/925080) (includes [5.15.97](https://lwn.net/Articles/925064), [5.15.96](https://lwn.net/Articles/924441), [5.15.95](https://lwn.net/Articles/924073), [5.15.94](https://lwn.net/Articles/923308), [5.15.93](https://lwn.net/Articles/922814)))
- Linux Firmware ([20230117](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230117))
- adcli ([0.9.2](https://gitlab.freedesktop.org/realmd/adcli/-/commits/8e88e3590a19006362ea8b8dfdc18bb88b3cb3b5/))
- bind tools ([9.16.36](https://bind9.readthedocs.io/en/v9_16_36/notes.html#notes-for-bind-9-16-36) (includes [9.16.34](https://bind9.readthedocs.io/en/v9_16_35/notes.html#notes-for-bind-9-16-34) and [9.16.35](https://bind9.readthedocs.io/en/v9_16_34/notes.html#notes-for-bind-9-16-35)))
- binutils ([2.39](https://sourceware.org/pipermail/binutils/2022-August/122246.html))
- bpftool ([5.19.12](https://lwn.net/Articles/909678/))
- ca-certificates ([3.89](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_89.html))
- containerd ([1.6.16](https://github.com/containerd/containerd/releases/tag/v1.6.16))
- cpio ([2.13](https://lists.gnu.org/archive/html/bug-cpio/2019-11/msg00000.html))
- curl ([7.87.0](https://curl.se/changes.html#7_87_0) (includes [7.85](https://curl.se/mail/archive-2022-08/0012.html)))
- dbus ([1.14.4](https://gitlab.freedesktop.org/dbus/dbus/-/raw/dbus-1.14.4/NEWS))
- Docker ([20.10.23](https://docs.docker.com/engine/release-notes/#201023))
- elfutils ([0.188](https://sourceware.org/pipermail/elfutils-devel/2022q4/005561.html) (includes [0.187](https://sourceware.org/pipermail/elfutils-devel/2022q2/004978.html)))
- Expat ([2.5.0](https://github.com/libexpat/libexpat/blob/R_2_5_0/expat/Changes))
- gawk ([5.2.1](https://lists.gnu.org/archive/html/help-gawk/2022-11/msg00008.html) (contains [5.2.0](https://lists.gnu.org/archive/html/help-gawk/2022-09/msg00000.html)))
- gettext ([0.21.1](https://git.savannah.gnu.org/gitweb/?p=gettext.git;a=blob;f=NEWS;h=cdbb16746c23555e70bb1e16917f5c349ce92d9e;hb=8b38ee827251cadbb90cb6cb576ae98702566288))
- git ([2.39.1](https://github.com/git/git/blob/v2.39.1/Documentation/RelNotes/2.39.1.txt) (includes [2.39.0](https://github.com/git/git/blob/v2.39.0/Documentation/RelNotes/2.39.0.txt)))
- glib ([2.74.4](https://gitlab.gnome.org/GNOME/glib/-/tags/2.74.4))
- Go ([1.19.5](https://go.dev/doc/devel/release#go1.19.5))
- glibc ([2.36](https://sourceware.org/pipermail/libc-alpha/2022-August/141193.html) (includes [2.35](https://savannah.gnu.org/forum/forum.php?forum_id=10111)))
- GnuTLS ([3.7.8](https://lists.gnupg.org/pipermail/gnutls-help/2022-September/004765.html))
- I2C tools ([4.3](https://git.kernel.org/pub/scm/utils/i2c-tools/i2c-tools.git/tree/CHANGES?id=d8bc1f1ff4b00a6bd988aa114100ae9b787f50d8))
- Intel Microcode ([20221108](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20221108))
- iptables ([1.8.8](https://www.netfilter.org/projects/iptables/files/changes-iptables-1.8.8.txt))
- iputils ([20211215](https://github.com/iputils/iputils/releases/tag/20211215))
- libcap ([2.66](https://sites.google.com/site/fullycapable/release-notes-for-libcap#h.d9ygdose5kw))
- libcap-ng ([0.8.3](https://people.redhat.com/sgrubb/libcap-ng/ChangeLog))
- libksba ([1.6.3](https://dev.gnupg.org/T6304))
- libseccomp ([2.5.4](https://github.com/seccomp/libseccomp/releases/tag/v2.5.4) (contains [2.5.2](https://github.com/seccomp/libseccomp/releases/tag/v2.5.2), [2.5.3](https://github.com/seccomp/libseccomp/releases/tag/v2.5.3)))
- libxml2 ([2.10.3](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.10.3))
- logrotate ([3.20.1](https://github.com/logrotate/logrotate/releases/tag/3.20.1))
- MIT Kerberos V ([1.20.1](https://web.mit.edu/kerberos/krb5-1.20/krb5-1.20.1.html))
- multipath-tools ([0.9.3](https://github.com/opensvc/multipath-tools/releases/tag/0.9.3))
- nettle ([3.8.1](https://git.lysator.liu.se/nettle/nettle/-/blob/990abad16ceacd070747dcc76ed16a39c129321e/ChangeLog))
- nmap ([7.93](https://nmap.org/changelog.html#7.93))
- OpenSSH ([9.1](http://www.openssh.com/releasenotes.html#9.1))
- rsync ([3.2.7](https://download.samba.org/pub/rsync/NEWS#3.2.7))
- shadow ([4.13](https://github.com/shadow-maint/shadow/releases/tag/4.13))
- sqlite ([3.40.1](https://www.sqlite.org/releaselog/3_40_1.html) (contains [3.40.0](https://www.sqlite.org/releaselog/3_40_0.html) and [3.39.4](https://sqlite.org/releaselog/3_39_4.html)))
- strace ([5.19](https://github.com/strace/strace/releases/tag/v5.19))
- sudo ([1.9.12_p2](https://github.com/sudo-project/sudo/releases/tag/SUDO_1_9_12p2))
- systemd ([252.5](https://github.com/systemd/systemd-stable/releases/tag/v252.5) (includes [252](https://github.com/systemd/systemd/releases/tag/v252)))
- vim ([9.0.1157](https://github.com/vim/vim/releases/tag/v9.0.1157) (includes [9.0.0469](https://github.com/vim/vim/releases/tag/v9.0.0469)))
- wget ([1.21.3](https://lists.gnu.org/archive/html/info-gnu/2022-02/msg00017.html))
- whois ([5.5.14](https://github.com/rfc1036/whois/commit/ab10466cf2e1ec4887f6a44375c3e29c1720157f))
- wireguard-tools ([1.0.20210914](https://github.com/WireGuard/wireguard-tools/releases/tag/v1.0.20210914))
- XZ utils ([5.4.1](https://github.com/tukaani-project/xz/releases/tag/v5.4.1) (includes [5.4.0](https://github.com/tukaani-project/xz/releases/tag/v5.4.0)))
- zlib ([1.2.13](https://github.com/madler/zlib/releases/tag/v1.2.13))
- OEM: python-oem ([3.9.16](https://www.python.org/downloads/release/python-3916/))
- SDK: boost ([1.81.0](https://www.boost.org/users/history/version_1_81_0.html))
- SDK: catalyst ([3.0.21](https://gitweb.gentoo.org/proj/catalyst.git/log/?h=3.0.21))
- SDK: cmake ([3.23.3](https://cmake.org/cmake/help/v3.23/release/3.23.html))
- SDK: file ([5.43](https://mailman.astron.com/pipermail/file/2022-September/000857.html) (includes [5.44](https://github.com/file/file/blob/FILE5_44/ChangeLog)))
- SDK: libpng ([1.6.39](http://www.libpng.org/pub/png/src/libpng-1.6.39-README.txt) (includes [1.6.38](http://www.libpng.org/pub/png/src/libpng-1.6.38-README.txt)))
- SDK: libxslt ([1.1.37](https://gitlab.gnome.org/GNOME/libxslt/-/tags/v1.1.37))
- SDK: meson ([0.62.2](https://mesonbuild.com/Release-notes-for-0-62-0.html))
- SDK: ninja ([1.11.0](https://groups.google.com/g/ninja-build/c/R2oCyDctDf8/m/-U94Y5I8AgAJ?pli=1))
- SDK: pahole ([1.23](https://git.kernel.org/pub/scm/devel/pahole/pahole.git/tag/?h=v1.23))
- SDK: perl ([5.36.0](https://perldoc.perl.org/5.36.0/perldelta))
- SDK: portage ([3.0.43](https://github.com/gentoo/portage/blob/portage-3.0.43/NEWS) (includes [3.0.42](https://github.com/gentoo/portage/blob/portage-3.0.42/NEWS), [3.0.41](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.41)))
- SDK: qemu ([7.2.0](https://wiki.qemu.org/ChangeLog/7.2) (includes [7.1.0](https://wiki.qemu.org/ChangeLog/7.1)))
- SDK: Rust ([1.67.0](https://github.com/rust-lang/rust/releases/tag/1.67.0))
- VMware: open-vm-tools ([12.1.5](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.1.5))

_Changes since **Beta 3510.1.0**_

#### Security fixes:


#### Bug fixes:

- Restored the support to specify OEM partition files in Ignition when `/usr/share/oem` is given as initrd mount point ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))

#### Changes:

- Added new image signing pub key to `flatcar-install`, needed for download verification of releases built from July 2023 onwards, if you have copies of `flatcar-install` or the image signing pub key, you need to update them as well ([init#92](https://github.com/flatcar/init/pull/92))
- Specifying the OEM filesystem in Ignition to write files to `/usr/share/oem` is not needed anymore ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))

#### Updates:

- ca-certificates ([3.89](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_89.html))

