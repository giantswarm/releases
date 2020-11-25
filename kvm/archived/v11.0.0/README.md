## :zap: Tenant Cluster Release 11.0.0 for KVM :zap:

This is the first Giant Swarm release which includes Kubernetes v1.16. In addition, CPU limits have been removed from several supporting components and priority classes have been adjusted to ensure system reliability under heavy load. Further details about changes to individual components can be found below.

### Important upgrade notes
- This release includes a new [network policy](https://docs.giantswarm.io/guides/limiting-pod-communication-with-network-policies/#default-policies) which blocks network traffic to and from pods in the `giantswarm` and `kube-system` namespaces by default to improve security. Giant Swarm components have been modified to work under this environment, but any other pods in the cluster communicating with these system pods may cease to function without a network policy allowing traffic. Thus, if you are running pods that need to talk to system pods, you will need to allow those explicitly by adding a [network policy](https://docs.giantswarm.io/guides/limiting-pod-communication-with-network-policies/) for them.
- [As previously communicated](https://github.com/giantswarm/giantswarm/blob/master/news/2019/10/product/k8s_1.16_breaking_changes.md), Kubernetes v1.16 has removed several API groups which were already deprecated. Please ensure that you have migrated your deployments and pipelines from the deprecated to the new API groups. You can check the above-linked breaking change communication and consult with your Solution Engineer for help with that.
- Previous Giant Swarm releases included Tiller v2.14 in the `giantswarm` namespace which is not compatible with Kubernetes v1.16. This Tiller - used by Giant Swarm to manage the cluster - has therefore been updated to the Kubernetes v1.16-compatible v2.16. Any Tiller instances deployed in other namespaces will likewise need to be updated to v2.16 or greater to continue to function after an upgrade.

---

### Kubernetes v1.16.3
- Updated from v1.15.5 - [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md).
- **Custom resources**: CRDs are in widespread use as a way to extend Kubernetes to persist and serve new resource types, and have been available in beta since the 1.7 release. The 1.16 release marks the graduation of CRDs to general availability (GA).
- **Admission webhooks**: Admission webhooks are in widespread use as a Kubernetes extensibility mechanism and have been available in beta since the 1.9 release. The 1.16 release marks the graduation of admission webhooks to general availability (GA).
- **Overhauled metrics**: Kubernetes has previously made extensive use of a global metrics registry to register metrics to be exposed. By implementing a metrics registry, metrics are registered in more transparent means. Previously, Kubernetes metrics have been excluded from any kind of stability requirements.
- **Volume Extension**: There are quite a few enhancements in this release that pertain to volumes and volume modifications. Volume resizing support in CSI specs is moving to beta which allows for any CSI spec volume plugin to be resizable.
- **Node labels** `beta.kubernetes.io/metadata-proxy-ready`, `beta.kubernetes.io/metadata-proxy-ready` and `beta.kubernetes.io/kube-proxy-ds-ready` are no longer added on new nodes.
- **Network Policies**: Giantswarm added a `deny-all` network policy by default. This policy is applied to all pods in sensitive namespaces such as `giantswarm` or `kube-system`. To communicate with any pods in these namespaces you need to explicitly create a Network Policy that allows it.
- As previously communicated, resources under `apps/v1beta1` and `apps/v1beta2` groups have been moved to `apps/v1` instead. Similarly, `daemonsets`, `deployments`, `replicasets` resources under `extensions/v1beta1` have been moved to `apps/v1`, `networkpolicies` under `extensions/v1beta1` to `networking.k8s.io/v1`, and `podsecuritypolicies` under `extensions/v1beta1` to `policy/v1beta1`.

### kvm-operator v3.10.0
- Added a check to ensure resource versions of control plane endpoints are correct before updating IPs to prevent IPs from being left over.

### Calico v3.10.1
- Updated from v3.9.1 - [changelog](https://docs.projectcalico.org/v3.10/release-notes/).
- Calico now supports two new top-level selectors to make writing Calico network policy easier. The `namespaceSelector` allows you to select the namespace(s) to apply a global network policy to. This enables you to write a single network policy applicable to one or more namespaces.
- Calico now supports BGP advertisement of Kubernetes service ExternalIPs in addition to advertising ClusterIPs. Advertising service external IPs allows for more flexible routing architectures.
- Configurable default IPv4, IPv6 block sizes and pool node selectors.
- Typha is now run as a non-root user for improved security.

### cert-exporter [v1.2.1](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#121-2019-12-24)
- Removed CPU limits to improve reliability.

### chart-operator [v0.11.3](https://github.com/giantswarm/chart-operator/releases/tag/v0.11.3)
- Removed CPU limits to improve reliability.
- Adjusted RBAC permissions.

### cluster-operator [v0.21.4](https://github.com/giantswarm/cluster-operator/releases/tag/v0.21.4)
- Added additional settings for coredns to cluster configmap.

### coreDNS v1.6.5 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v110))
- Updated from upstream `coredns` v1.6.4 - [changelog](https://coredns.io/2019/11/05/coredns-1.6.5-release/).
- Removed CPU limits to improve reliability.
- Migrated to be deployed via an app CR not a chartconfig CR.

### CoreOS Container Linux v2247.6.0
- Updated from v2191.5.0 - [changelog](https://coreos.com/releases/#2247.6.0).
- Updated Linux kernel to 4.19.78.

### etcd v3.3.17
- Updated from v3.3.15 - [changelog](https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.3.md#v3317-2019-10-11).

### Helm v2.16.1 (primarily for Giant Swarm internal use)
- Updated from v2.14.3 - [changelog](https://github.com/helm/helm/releases/tag/v2.16.1).
- Helm v2.15 was the last feature release for Helm v2 as new feature development now happens in Helm v3. The v2.16 release includes fixes to issues that are too large of a change for a patch release.

### kube-state-metrics v1.9.2 ([Giant Swarm app v1.0.2](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#v102))
- Updated to upstream version 1.9.2 - [changelog](https://github.com/kubernetes/kube-state-metrics/blob/master/CHANGELOG.md#v192--2020-01-13).
- Adjusted RBAC configuration.

### net-exporter [v1.5.1](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#151-2020-01-08)
- Changed priority class to `system-node-critical`.

### nginx-ingress-controller v0.26.1 ([Giant Swarm app v1.1.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v111-2020-01-04))
- Updated manifests for Kubernetes v1.16 compatibility.

### node-exporter v0.18.1 ([Giant Swarm app v1.2.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#120-2020-01-08))
- Updated to upstream version 0.18.1 - [changelog](https://github.com/prometheus/node_exporter/blob/master/CHANGELOG.md#0181--2019-06-04).
- Changed priority class to `system-node-critical`.
