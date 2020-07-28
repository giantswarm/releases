## :zap:  Giant Swarm Release 12.1.0 for KVM :zap:

**This release includes two significant improvements to NGINX Ingress Controller. It also includes a fix for Quay being a single point of failure by using Docker mirroring feature. This ensures availability of all images needed for node bootstrap, thus the cluster creation/scaling doesn’t depend on Quay availability anymore.**

**The two NGINX Ingress Controller improvements:**

- First, multiple NGINX Ingress Controllers per tenant cluster are now supported, enabling separation of internal vs external traffic, dev vs prod, and so on.
- Second, management of NGINX IC NodePort Service is moved from kvm-operator to NGINX IC App itself. This enables configurability of external traffic policy and lays the foundation for making NGINX IC App optional and not pre-installed in a future KVM platform release.

Along with kvm-operator, cluster-operator and NGINX IC, release includes several upstream component upgrades.

**Note for cluster upgrades:**

Please manually delete `nginx-ingress-controller` NodePort Service in `kube-system` namespace. Upgrading the cluster then recreates the NodePort Service and moves its management from kvm-operator to NGINX IC. To minimize downtime, please delegate cluster upgrades to your SE.**

**Note for future 12.1.x releases:**

To prevent downtime, please persist this and the above note until all customers are on 12.1.0 and above.

Below, you can find more details on components that were changed with this release.

### cluster-operator [0.23.13](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.13)

- Enable NGINX App managed NodePort Service on KVM.

### kube-state-metrics v1.9.7 ([Giant Swarm app v1.1.1](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#111---2020-07-22))

- Updated kube-state-metrics version from 1.9.5 to 1.9.7. Check the [upstream changelog](https://github.com/kubernetes/kube-state-metrics/blob/master/CHANGELOG.md#v197--2020-05-24) for details on all changes.

### kvm-operator [v3.12.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.12.1)

- Add registry mirrors support.
- Stop provisioning NGINX IC NodePort Service.

### metrics-server v0.3.6 ([Giant Swarm app v1.1.1](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#111---2020-07-23))

- Updated metrics-server version from 0.3.3 to 0.3.6. Check the [upstream changelog](https://github.com/kubernetes-sigs/metrics-server/releases) for details on all changes.

### nginx-ingress-controller v0.34.1 ([Giant Swarm app v1.8.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#181---2020-07-28))

- Support multiple NGINX IC App installations per tenant cluster.
- Made NGINX NodePort Service external traffic policy configurable.
- Made NGINX NodePort Service node ports configurable.
- Drop support for deprecated configuration properties.

### node-exporter v1.0.1 ([Giant Swarm app v1.3.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#130---2020-07-23))

- Updated node-exporter version from 0.18.1 to 1.0.1. Check the [upstream changelog](https://github.com/prometheus/node_exporter/blob/master/CHANGELOG.md#101--2020-06-15) for details on all changes.
