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


### calico [3.12.3](https://github.com/projectcalico/calico/releases/tag/v3.12.3)

#### Other changes
- Remove unnecessary packages from CNI plugin docker image [cni-plugin #902](https://github.com/projectcalico/cni-plugin/pull/902) (@gianlucam76)
- Remove unnecessary packages from pod2daemon docker image [pod2daemon #38](https://github.com/projectcalico/pod2daemon/pull/38) (@gianlucam76)
