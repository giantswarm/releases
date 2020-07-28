## :zap:  Giant Swarm Release 12.1.0 for KVM :zap:

**This release includes two significant improvements to NGINX Ingress Controller. It also adds support for image registry mirrors, improving reliability.**

**The two NGINX Ingress Controller improvements:**
* First, multiple NGINX Ingress Controllers per tenant cluster are now supported, enabling separation of internal vs external traffic, dev vs prod, and so on.
* Second, management of NGINX IC NodePort Service is moved from kvm-operator to NGINX IC App itself. This enables configurability of external traffic policy and lays the foundation for making NGINX IC App optional and not pre-installed in a future KVM platform release.

**To change the ownership, `nginx-ingress-controller` NodePort Service in `kube-system` namespace needs to be manually deleted. It then gets to be recreated by triggering the cluster upgrade. To minimize the downtime, please delegate cluster upgrades to your SE.**

**Note for future 12.1.x releases:**

To prevent downtime, please persist this and the above note until all customers are on 12.1.0 and above.

Below, you can find more details on components that were changed with this release.

### cluster-operator [0.23.13](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.13)

- Enable NGINX App managed NodePort Service on KVM.

### kvm-operator [v3.12.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.12.1)

- Add registry mirrors support.
- Stop provisioning NGINX IC NodePort Service.

### nginx-ingress-controller v0.34.1 ([Giant Swarm app v1.8.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#181---2020-07-28))

- Support multiple NGINX IC App installations per tenant cluster.
- Made NGINX NodePort Service external traffic policy configurable.
- Made NGINX NodePort Service node ports configurable.
- Drop support for deprecated configuration properties.
