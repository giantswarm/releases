# :zap: Giant Swarm Release v15.0.0-beta1 for Azure :zap:

<< Add description here >>

## Change details


### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases.
- Emit events when config maps and secrets referenced in App CRs are updated.



### azure-operator [5.6.1-k8s-1-20](https://github.com/giantswarm/azure-operator/releases/tag/v5.6.1-k8s-1-20)

Not found


### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

#### Changed
- Dropped ensuring cluster CRDs from controllers.



### kubernetes [1.20.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.6)

#### API Change
- Fixes using server-side apply with APIService resources ([#100714](https://github.com/kubernetes/kubernetes/pull/100714), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Apps and Testing]
- Regenerate protobuf code to fix CVE-2021-3121 ([#100501](https://github.com/kubernetes/kubernetes/pull/100501), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery, Apps, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
#### Feature
- AWS cloudprovider supports auto-discovering subnets without any kubernetes.io/cluster/<clusterName> tags. It also supports additional service annotation service.beta.kubernetes.io/aws-load-balancer-subnets to manually configure the subnets. ([#97431](https://github.com/kubernetes/kubernetes/pull/97431), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- Kubernetes is now built using go1.15.10 ([#100375](https://github.com/kubernetes/kubernetes/pull/100375), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- ## Changelog
  
  ### General
  - Fix priority expander falling back to a random choice even though there is a higher priority option to choose
  - Clone `kubernetes/kubernetes` in `update-vendor.sh` shallowly, instead of fetching all revisions
  - Speed up binpacking by reducing the number of PreFilter calls (call once per pod instead of #pods*#nodes times)
  - Speed up finding unneeded nodes by 5x+ in very large clusters by reducing the number of PreFilter calls
  - Expose `--max-nodes-total` as a metric
  - Errors in `IncreaseSize` changed from type `apiError` to `cloudProviderError`
  - Make `build-in-docker` and `test-in-docker` work on Linux systems with SELinux enabled
  - Fix an error where existing nodes were not considered as destinations while finding place for pods in scale-down simulations
  - Remove redundant log lines and reduce severity around parsing kubeEnv
  - Don't treat nodes created by virtual kubelet as nodes from non-autoscaled node groups
  - Remove redundant logging around calculating node utilization
  - Add configurable `--network` and `--rm` flags for docker in `Makefile`
  - Subtract DaemonSet pods' requests from node allocatable in the denominator while computing node utilization
  - Include taints by condition when determining if a node is unready/still starting
  - Fix `update-vendor.sh` to work on OSX and zsh
  - Add best-effort eviction for DaemonSet pods while scaling down non-empty nodes
  - Add build support for ARM64
  
  ### AliCloud
  - Add missing daemonsets and replicasets to ALI example cluster role
  
  ### Apache CloudStack
  - Add support for Apache CloudStack
  
  ### AWS
  - Regenerate list of EC2 instances
  - Fix pricing endpoint in AWS China Region
  
  ### Azure
  - Add optional jitter on initial VMSS VM cache refresh, keep the refreshes spread over time
  - Serve from cache for the whole period of ongoing throttling
  - Fix unwanted VMSS VMs cache invalidation
  - Enforce setting the number of retries if cloud provider backoff is enabled
  - Don't update capacity if VMSS provisioning state is updating
  - Support allocatable resources overrides via VMSS tags
  - Add missing stable labels in template nodes
  - Proactively set instance status to deleting on node deletions
  
  ### Cluster API
  - Migrate interaction with the API from using internal types to using Unstructured
  - Improve tests to work better with constrained resources
  - Add support for node auto-discovery
  - Add support for `--cloud-config`
  - Update group identifier to use for Cluster API annotations
  
  ### Exoscale
  - Add support for Exoscale
  
  ### GCE
  - Decrease the number of GCE Read Requests made while deleting nodes
  - Base pricing of custom instances on their instance family type
  - Add pricing information for missing machine types
  - Add pricing information for different GPU types
  - Ignore the new `topology.gke.io/zone` label when comparing groups
  - Add missing stable labels to template nodes
  
  ### HuaweiCloud
  - Add auto scaling group support
  - Implement node group by AS
  - Implement getting desired instance number of node group
  - Implement increasing node group size
  - Implement TemplateNodeInfo
  - Implement caching instances
  
  ### IONOS
  - Add support for IONOS
  
  ### Kubemark
  - Skip non-kubemark nodes while computing node information for node groups.
  
  ### Magnum
  - Add Magnum support in the Cluster Autoscaler helm chart
  
  ### Packet
  - Allow empty nodepools
  - Add support for multiple nodepools
  - Add pricing support
  
  ## Image
  Image: `k8s.gcr.io/autoscaling/cluster-autoscaler:v1.20.0` ([#97012](https://github.com/kubernetes/kubernetes/pull/97012), [@towca](https://github.com/towca)) [SIG Cloud Provider]
- Fixed a bug where a high churn of events was causing master instability by reducing the maximum number of objects (events) attached to a single etcd lease. ([#100084](https://github.com/kubernetes/kubernetes/pull/100084), [@mborsz](https://github.com/mborsz)) [SIG API Machinery, Instrumentation and Scalability]
- Fixed a race condition on API server startup ensuring previously created webhook configurations are effective before the first write request is admitted. ([#95783](https://github.com/kubernetes/kubernetes/pull/95783), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixes a data race issue in the priority and fairness API server filter ([#100667](https://github.com/kubernetes/kubernetes/pull/100667), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Kubectl: Fixed panic when describing an ingress backend without an API Group ([#100541](https://github.com/kubernetes/kubernetes/pull/100541), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- Reverts breaking change to inline AzureFile volumes in v1.20.2-v1.20.5; referenced secrets are now correctly searched for in the same namespace as the pod as in previous releases. ([#100399](https://github.com/kubernetes/kubernetes/pull/100399), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- The endpointslice mirroring controller mirrors endpoints annotations and labels to the generated endpoint slices, it also ensures that updates on any of these fields on the endpoints are mirrored. 
  The well-known annotation endpoints.kubernetes.io/last-change-trigger-time is skipped and not mirrored. ([#100443](https://github.com/kubernetes/kubernetes/pull/100443), [@aojea](https://github.com/aojea)) [SIG Apps, Network and Testing]
- The maximum number of ports allowed in EndpointSlices has been increased from 100 to 20,000 ([#99795](https://github.com/kubernetes/kubernetes/pull/99795), [@robscott](https://github.com/robscott)) [SIG Network]
#### Uncategorized
- GCE L4 Loadbalancers now handle > 5 ports in service spec correctly. ([#99595](https://github.com/kubernetes/kubernetes/pull/99595), [@prameshj](https://github.com/prameshj)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/yuin/goldmark: [v1.1.27 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.1.27...v1.2.1)
- golang.org/x/sync: cd5d95a → 67f06af
- golang.org/x/tools: c1934b7 → 113979e
- sigs.k8s.io/structured-merge-diff/v4: v4.0.2 → v4.0.3
#### Removed
_Nothing has changed._



### containerlinux [2765.2.3](https://www.flatcar-linux.org/releases/#release-2765.2.3)


**Security fixes**



*   Linux ([CVE-2021-28964](https://nvd.nist.gov/vuln/detail/CVE-2021-28964), [CVE-2021-28972](https://nvd.nist.gov/vuln/detail/CVE-2021-28972), [CVE-2021-28971](https://nvd.nist.gov/vuln/detail/CVE-2021-28971), [CVE-2021-28951](https://nvd.nist.gov/vuln/detail/CVE-2021-28951), [CVE-2021-28952](https://nvd.nist.gov/vuln/detail/CVE-2021-28952), [CVE-2021-29266](https://nvd.nist.gov/vuln/detail/CVE-2021-29266), [CVE-2021-28688](https://nvd.nist.gov/vuln/detail/CVE-2021-28688), [CVE-2021-29264](https://nvd.nist.gov/vuln/detail/CVE-2021-29264), [CVE-2021-29649](https://nvd.nist.gov/vuln/detail/CVE-2021-29649), [CVE-2021-29650](https://nvd.nist.gov/vuln/detail/CVE-2021-29650), [CVE-2021-29646](https://nvd.nist.gov/vuln/detail/CVE-2021-29646), [CVE-2021-29647](https://nvd.nist.gov/vuln/detail/CVE-2021-29647), [CVE-2021-29154](https://nvd.nist.gov/vuln/detail/CVE-2021-29154), [CVE-2021-29155](https://nvd.nist.gov/vuln/detail/CVE-2021-29155), [CVE-2021-23133](https://nvd.nist.gov/vuln/detail/CVE-2021-23133))

**Bug fixes**



*   Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))

**Updates**



*   Linux ([5.10.32](https://lwn.net/Articles/853762/))


### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry



### chart-operator [2.14.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.14.0)

#### Changed
- Cancel the release resource when the manifest object already exists.
- Cancel the release resource when helm returns an unknown error.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

#### Added
- Added new configuration value `extraArgs`.



### net-exporter [1.10.1](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.1)




### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed
- Set docker.io as the default registry



### cluster-autoscaler [1.20.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.2)

#### Changed
- Set docker.io as the default registry



