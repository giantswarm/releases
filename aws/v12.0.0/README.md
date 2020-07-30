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

### kube-state-metrics v1.9.7 (Giant Swarm app [v1.1.1](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md))

- Switch mutatingWebhookConfiguration to use v1 api

Check the [kube-state-metrics changelog](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.9.7) for more details.

### metrics-server v0.3.6 (Giant Swarm app [v1.1.1](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md))

- Fix: Don't break metric storage when duplicate pod metrics encountered

Check the [metrics-server changelog](https://github.com/kubernetes-sigs/metrics-server/releases) for more details.

### node-exporter v1.0.1 (Giant Swarm app [v1.3.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md))

From 0.18.1 (app v1.2.0)

#### Breaking changes

- The netdev collector CLI argument `--collector.netdev.ignored-devices` was renamed to `--collector.netdev.device-blacklist` in order to conform with the systemd collector.
- The label named `state` on `node_systemd_service_restart_total` metrics was changed to `name` to better describe the metric.
- Refactoring of the `mdadm` collector changes several metrics
  - node_md_disks_active is removed
  - node_md_disks now has a state label for "failed", "spare", "active" disks.
  - node_md_is_active is replaced by node_md_state with a state set of "active", "inactive", "recovering", "resync".
- Additional label `mountaddr` added to NFS device metrics to distinguish mounts from the same URL, but different IP addresses.
- Metrics `node_cpu_scaling_frequency_min_hrts` and node_cpu_scaling_frequency_max_hrts of the cpufreq collector were renamed to `node_cpu_scaling_frequency_min_hertz` and `node_cpu_scaling_frequency_max_hertz`.
- Collectors that are enabled, but are unable to find data to collect, now return 0 for `node_scrape_collector_success`.
