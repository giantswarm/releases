# :zap: Giant Swarm Release v18.2.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.23.15](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.15)

#### Feature
- Kubeadm: use the image registry registry.k8s.io instead of k8s.gcr.io for new clusters. During upgrade, migrate users to registry.k8s.io if they were using the default of k8s.gcr.io. ([#113393](https://github.com/kubernetes/kubernetes/pull/113393), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
#### Bug or Regression
- Fix endpoint reconciler not being able to delete the apiserver lease on shutdown ([#114153](https://github.com/kubernetes/kubernetes/pull/114153), [@aojea](https://github.com/aojea)) [SIG API Machinery]
- Fix for volume reconstruction of CSI ephemeral volumes ([#113347](https://github.com/kubernetes/kubernetes/pull/113347), [@dobsonj](https://github.com/dobsonj)) [SIG Node, Storage and Testing]
- Fix performance issue when creating large objects using SSA with fully unspecified schemas (preserveUnknownFields). ([#111914](https://github.com/kubernetes/kubernetes/pull/111914), [@aojea](https://github.com/aojea)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
#### Other (Cleanup or Flake)
- Kubelet now defaults to pulling the pause image from registry.k8s.io ([#114340](https://github.com/kubernetes/kubernetes/pull/114340), [@liggitt](https://github.com/liggitt)) [SIG Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/structured-merge-diff/v4: v4.2.1 → v4.2.3
#### Removed
_Nothing has changed._



### etcd [3.5.6](https://github.com/etcd-io/etcd/releases/tag/v3.5.6)

Not found


### app-operator [6.5.0](https://github.com/giantswarm/app-operator/releases/tag/v6.5.0)

#### Fixed
- Fix building URLs for OCI Repositories assigned to non-internal `Catalogs`.



### aws-operator [14.4.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.4.0)

#### Changed
- Bump k8scc to 15.4.0.



### node-exporter [1.15.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.15.0)

#### Fixed
- Fix collector `systemd`
- Fix duplicate scrapping by GiantSwarm Prometheus
#### Added
- Add values schema



### vertical-pod-autoscaler [2.5.3](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.5.3)

#### Changed
- Increased memory limits for updater, recommender and admissionController



