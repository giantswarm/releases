# :zap: Giant Swarm Release v13.0.0 for Azure :zap:

This is the first release to support Kubernetes 1.18 on Azure.

This release also includes support for Kubernetes node pools.
A node pool is a subset of the Kubernetes nodes. They enable having pools of nodes with different configurations (like a different instance size) within one cluster.
After cluster creation with 1 node pools, additional node pools can be freely added and removed from the cluster.

If you have access to the Control Plane API you can manage your clusters directly from there.
The clusters that you create are now represented by [Cluster API](https://cluster-api.sigs.k8s.io/) CRDs ([Custom Resource Definition](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)).
Using [our kubectl plugin](https://github.com/giantswarm/kubectl-gs/) you can easily create the Custom Resources required to create a cluster.

## Change details

### azure-operator [5.0.0](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0)

Bug fixes:

- Fixed firewall rules to allow prometheus to scrape node-level exporters from all node pools.
- Try to send only one request to VMSS Azure API from `nodepool` handler which reduces the risk of reaching API rate limit and getting 429 errors.

Changes:

- Decouple `Service` from controllers using an slice of controllers.
- Retry failed ARM deployments regardless of the checksum check.
- Master instances now use a dedicated NAT gateway for egress traffic.
- Removed the ETCD Private Load Balancer, reusing the API public one for ETCD traffic (needed by HA masters).
- Use `DataDisks` field to define VM disks when creating node pools.
- Don't error if certificates are not present yet. Cancel reconciliation and wait until next loop instead.
- Set cluster-autoscaler-enabled tag to false when min replicas and max replicas are the same for a node pool.
- Removed instance watchdog to save on VMSS API calls.
- Do not use public SSH keys coming from the CRs.
- Get the storage account type to use for node pools' VMSS from the AzureMachinePool CR.
- Add provider independent controllers to manage labeling and setting owner references in other provider dependent objects.
- Enable persistent volume `expansion` support in the default `Storage Classes`.
- Added to all VMSSes the tags needed by Cluster Autoscaler.

Updates:

- Updated backward incompatible Kubernetes dependencies to v1.18.5.
- Updated CAPI to `v0.3.9` and CAPZ to `v0.4.7`, using GiantSwarm forks that contain k8s 1.18 changes.


### containerlinux [2512.5.0](https://www.flatcar-linux.org/releases/#release-2512.5.0)

Changes:
- Update public key to include a [new subkey](https://www.flatcar-linux.org/security/image-signing-key/)

Updates:
- Linux [4.19.145](https://lwn.net/Articles/831367/)



### kubernetes [1.18.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.10)

#### Design
- Prevent logging of docker config contents if file is malformed ([#95347](https://github.com/kubernetes/kubernetes/pull/95347), [@sfowl](https://github.com/sfowl)) [SIG Auth and Node]
#### Bug or Regression
- Do not fail sorting empty elements. ([#94666](https://github.com/kubernetes/kubernetes/pull/94666), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Ensure getPrimaryInterfaceID not panic when network interfaces for Azure VMSS are null ([#94801](https://github.com/kubernetes/kubernetes/pull/94801), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix bug where loadbalancer deletion gets stuck because of missing resource group &#35;75198 ([#93962](https://github.com/kubernetes/kubernetes/pull/93962), [@phiphi282](https://github.com/phiphi282)) [SIG Cloud Provider]
- Fix detach azure disk issue when vm not exist ([#95177](https://github.com/kubernetes/kubernetes/pull/95177), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix etcd_object_counts metric reported by kube-apiserver ([#94818](https://github.com/kubernetes/kubernetes/pull/94818), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fix network_programming_latency metric reporting for Endpoints/EndpointSlice deletions, where we don't have correct timestamp ([#95363](https://github.com/kubernetes/kubernetes/pull/95363), [@wojtek-t](https://github.com/wojtek-t)) [SIG Network and Scalability]
- Fix scheduler cache snapshot when a Node is deleted before its Pods ([#95154](https://github.com/kubernetes/kubernetes/pull/95154), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Fix the `cloudprovider_azure_api_request_duration_seconds` metric buckets to correctly capture the latency metrics. Previously, the majority of the calls would fall in the "+Inf" bucket. ([#95375](https://github.com/kubernetes/kubernetes/pull/95375), [@marwanad](https://github.com/marwanad)) [SIG Cloud Provider and Instrumentation]
- Fix: azure disk resize error if source does not exist ([#93011](https://github.com/kubernetes/kubernetes/pull/93011), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: detach azure disk broken on Azure Stack ([#94885](https://github.com/kubernetes/kubernetes/pull/94885), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a bug where improper storage and comparison of endpoints led to excessive API traffic from the endpoints controller ([#94934](https://github.com/kubernetes/kubernetes/pull/94934), [@damemi](https://github.com/damemi)) [SIG Apps, Network and Testing]
- Gracefully delete nodes when their parent scale set went missing ([#95289](https://github.com/kubernetes/kubernetes/pull/95289), [@bpineau](https://github.com/bpineau)) [SIG Cloud Provider]
- Kubeadm: warn but do not error out on missing "ca.key" files for root CA, front-proxy CA and etcd CA, during "kubeadm join --control-plane" if the user has provided all certificates, keys and kubeconfig files which require signing with the given CA keys. ([#94988](https://github.com/kubernetes/kubernetes/pull/94988), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
#### Other (Cleanup or Flake)
- Masks ceph RBD adminSecrets in logs when logLevel >= 4 ([#95245](https://github.com/kubernetes/kubernetes/pull/95245), [@sfowl](https://github.com/sfowl)) [SIG Storage]



### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)

#### Other changes
 - Add FelixConfiguration parameters to explicitly allow encapsulated packets from workloads. [libcalico-go #1302](https://github.com/projectcalico/libcalico-go/pull/1302) (@doublek)
 - Respect explicit configuration for drop rules for encapsulated packets originating from workloads. [felix #2487](https://github.com/projectcalico/felix/pull/2487) (@doublek)



### chart-operator [2.3.5](https://github.com/giantswarm/chart-operator/releases/tag/v2.3.5)

#### Fixed
- Stop repeating helm upgrade for the failed helm release.



### app-operator [2.3.3](https://github.com/giantswarm/app-operator/releases/tag/v2.3.3)

#### Added
- Delete chart-operator helm release and chart CR so it can be re-installed.



### kube-state-metrics [1.2.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.2.0)

#### Added
- Added monitoring annotations and common labels.
#### Changed
- Deploy `kube-state-metrics-app` on installations as part of app collection.



### cert-exporter [1.3.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.3.0)

#### Added
- Add Network Policy.
#### Changed
- Remove `hostNetwork` and `hostPID` capabilities.



### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.2)

#### Changed
- Updated backward incompatible Kubernetes dependencies to v1.18.5.
#### Fixed
- Fixed indentation problem with the daemonset template.



### node-exporter [1.4.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.4.2)

#### Changed
- Disable `btrfs`,`softnet`,`rapl` and `thermal_zone` to reduce memory usage.
- Increase memory limit to `75Mi`.



### external-dns [1.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v1.5.0)

#### Changed
- Upgrade upstream external-dns from v0.7.3 to [v0.7.4](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.4).

