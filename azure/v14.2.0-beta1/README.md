# :zap: Giant Swarm Release v14.2.0-beta1 for Azure :zap:

This is a beta release for Azure that replaces the previously used `VPN gateway` with `VNet peering`, cutting cluster deployment time
by several minutes. Most workload clusters should be up and running in about 10 minutes from this release on.

Being a `beta` release, you are free to deploy new clusters as well as update existing ones, but please keep in mind that it is discouraged
to use it for production workload clusters.

For further information get in touch with your solution engineer.

## Change details


### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

#### Changed
- Dropped ensuring cluster CRDs from controllers.



### azure-operator [5.6.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.6.0)

#### Changed
- Replace VPN Gateway with VNet Peering.
- Update OperatorKit to `v4.3.1` to drop usage of self-link which is not supported in k8s 1.20 anymore.
#### Removed
- Support for single tenant BYOC credentials (warning: the operator will error at startup if any organization credentials is not multi tenant).



### kubernetes [1.19.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.10)

#### API Change
- Fixes using server-side apply with APIService resources ([#100713](https://github.com/kubernetes/kubernetes/pull/100713), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Apps, Scheduling and Testing]
- Regenerate protobuf code to fix CVE-2021-3121 ([#100515](https://github.com/kubernetes/kubernetes/pull/100515), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
#### Feature
- AWS cloudprovider supports auto-discovering subnets without any kubernetes.io/cluster/<clusterName> tags. It also supports additional service annotation service.beta.kubernetes.io/aws-load-balancer-subnets to manually configure the subnets. ([#97431](https://github.com/kubernetes/kubernetes/pull/97431), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- AWS cloudprovider will ignore provisioning load balancers if the annotation service.beta.kubernetes.io/aws-load-balancer-type is external or nlb-ip ([#97975](https://github.com/kubernetes/kubernetes/pull/97975), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- Kubernetes is now built using go1.15.10 ([#100520](https://github.com/kubernetes/kubernetes/pull/100520), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Fixed a bug where a high churn of events was causing master instability by reducing the maximum number of objects (events) attached to a single etcd lease. ([#100450](https://github.com/kubernetes/kubernetes/pull/100450), [@mborsz](https://github.com/mborsz)) [SIG API Machinery and Instrumentation]
- Fixed a race condition on API server startup ensuring previously created webhook configurations are effective before the first write request is admitted. ([#95783](https://github.com/kubernetes/kubernetes/pull/95783), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixes a data race issue in the priority and fairness API server filter ([#100669](https://github.com/kubernetes/kubernetes/pull/100669), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Kubectl: Fixed panic when describing an ingress backend without an API Group ([#100542](https://github.com/kubernetes/kubernetes/pull/100542), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- Reverts breaking change to inline AzureFile volumes in v1.19.7-v1.19.9; referenced secrets are now correctly searched for in the same namespace as the pod as in previous releases. ([#100398](https://github.com/kubernetes/kubernetes/pull/100398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- The endpointslice mirroring controller mirrors endpoints annotations and labels to the generated endpoint slices, it also ensures that updates on any of these fields on the endpoints are mirrored. 
  The well-known annotation endpoints.kubernetes.io/last-change-trigger-time is skipped and not mirrored. ([#100443](https://github.com/kubernetes/kubernetes/pull/100443), [@aojea](https://github.com/aojea)) [SIG Apps, Network and Testing]
- The maximum number of ports allowed in EndpointSlices has been increased from 100 to 20,000 ([#99795](https://github.com/kubernetes/kubernetes/pull/99795), [@robscott](https://github.com/robscott)) [SIG Network]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/yuin/goldmark: [v1.1.27 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.1.27...v1.2.1)
- golang.org/x/sync: cd5d95a → 67f06af
- golang.org/x/tools: c1934b7 → 113979e
- golang.org/x/xerrors: 9bdfabe → 5ec99f8
- sigs.k8s.io/structured-merge-diff/v4: v4.0.1 → v4.0.3
#### Removed
_Nothing has changed._



### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



### net-exporter [1.9.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.3)

#### Changed
- Set docker.io as the default registry
- Update kubectl image to v1.18.8.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.2.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.2.2)

#### Changed
- Set docker.io as the default registry



### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed
- Set docker.io as the default registry



### chart-operator [2.13.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.13.1)

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.



### azure-scheduled-events [0.4.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.4.0)

#### Added
- React to `Preempt`, `Reboot` and `Redeploy` events to drain nodes properly.
#### Change
- Use the `NotBefore` field from the event to calculate drain timeout.



### cluster-autoscaler [1.19.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.2)

#### Changed
- Set docker.io as the default registry



