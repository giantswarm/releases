# :zap: Giant Swarm Release v12.0.0 for AWS :zap:

This release provides Kubernetes 1.17. It is based on new aws-operator and cluster-operator versions and picks up upgrades tomany components.

## Change details

### kubernetes [1.17.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.9)

- CVE-2020-8557 (Medium): Node-local denial of service via container /etc/hosts file. See https://github.com/kubernetes/kubernetes/issues/93032 for more details.
- Extend kube-apiserver /readyz with new "informer-sync" check ensuring that internal informers are synced. ([#92644](https://github.com/kubernetes/kubernetes/pull/92644), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Testing]
- Kubeadm: add the deprecated flag --port=0 to kube-controller-manager and kube-scheduler manifests to disable insecure serving. Without this flag the components by default serve (e.g. /metrics) insecurely on the default node interface (controlled by --address). Users that wish to override this behavior and enable insecure serving can pass a custom --port=X via kubeadm's "extraArgs" mechanic for these components. ([#92720](https://github.com/kubernetes/kubernetes/pull/92720), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: during "join", don't re-add an etcd member if it already exists in the cluster. ([#92118](https://github.com/kubernetes/kubernetes/pull/92118), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- hyperkube: Use debian-hyperkube-base@v1.1.1 image
  
    Includes iproute2 to fix a regression in hyperkube images
    when using hyperkube as a kubelet ([#92625](https://github.com/kubernetes/kubernetes/pull/92625), [@justaugustus](https://github.com/justaugustus)) [SIG Cluster Lifecycle, Network and Release]

### aws-operator [v8.7.5](https://github.com/giantswarm/aws-operator/blob/master/CHANGELOG.md#875---2020-07-30)

- Adjust MAX_PODS setting for master and worker nodes to max IP's per ENI when using aws-cni

### calico [v3.15.1](https://github.com/projectcalico/calico/releases/tag/v3.15.1)

- Fix issue with service IP advertisement breaking host service connectivity

Complete release notes can be found at [docs.projectcalico.org/v3.15/release-notes](https://docs.projectcalico.org/v3.15/release-notes/)

### chart-operator [v0.13.2](https://github.com/giantswarm/chart-operator/releases/tag/v0.13.2)

- Add metrics for Helm releases with a mismatched namespace.
- Calculate md5sum from go struct.

### cluster-autoscaler [v1.17.3]() (Giant Swarm app [v1.17.X]()) - TODO

from 1.16.5 / Giant Swarm app v 1.16.0

- Switch leader election mechanism to use lease objects, includes RBAC rule update
- Nodes with small difference in available memory will now be considered similar for the purposes of balancing node pool sizes. This should increase the reliability of node pool balancing.
- Fixed a bug where Cluster Autoscaler incorrectly didn't take into account resource requests of initContainers.
- Allow marking pods with the `cluster-autoscaler.kubernetes.io/daemonset-pod` annotation. Marked pods will be treated as daemon set pods.
- Add TTL and batching to launch config cache mechanism.
- improved logging when a taint blocks scale-up
- Invalidate node instances cache after deleting failed and unregistered nodes
- Allow for greater node memory mismatch on comparison for similarity (bumped to 256K from 128K)
- Fix case when ASG size could be descresed twice
- Allow arbitrary placeholder AWS instance names

### kiam [v3.6]() (Giant Swarm app v1.3.1)

Check the [kiam changelog](https://github.com/uswitch/kiam/blob/master/CHANGELOG.md#v36) for details.

### app-operator v1.1.9

from 1.0.0

