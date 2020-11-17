# :zap: Giant Swarm Release v13.0.0-alpha4 for Azure :zap:

## Change details

### azure-operator [5.0.0-alpha4](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0-alpha4)

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




