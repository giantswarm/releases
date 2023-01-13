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



### aws-ebs-csi-driver [2.19.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.19.0)

#### Fixed
- Fix scraping monitoring port.
#### Added
- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in `deployment/ebs-plugin` if defined in `values.yaml`.
- Support for using `cluster-apps-operator` specific `cluster.proxy` values.



### cert-exporter [2.3.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.3.1)

#### Fixed
- Allow eviction for cert-exporter-deployment.



### chart-operator [2.33.2](https://github.com/giantswarm/chart-operator/releases/tag/v2.33.2)




### coredns [1.13.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.13.0)

#### Added
- `values.schema.json` file
#### Changed
- Move nodeselector `label:value` to values.yaml to allow customizing it for CAPZ
- Add toleration for `node-role.kubernetes.io/control-plane` to masters instance



### external-dns [2.23.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.23.1)

#### Fixed
- Restore missing pod annotations in deployment ([#232](https://github.com/giantswarm/external-dns-app/pull/232)).



### kiam [2.6.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.6.0)

- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in deployment if defined in `values.yaml`.
- Support for using `cluster-apps-operator` specific `cluster.proxy` values.



### kube-state-metrics [1.14.2](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.14.2)

#### Changed
- Re add to Azure collections.



### metrics-server [2.0.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.0.0)

#### Changed
- Switch to HA setup.



### net-exporter [1.13.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.13.0)

#### Added
- Add helm chart values schema
#### Changed
- Update to Go 1.18
- Update github.com/giantswarm/k8sclient to v7.0.1
- Update github.com/giantswarm/micrologger to v1.0.0
- Update github.com/miekg/dns to v1.1.50
- Update k8s.io deps to v0.26.0
- Update docker-kubectl to 1.25.4



