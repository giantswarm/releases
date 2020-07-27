## :zap:  Giant Swarm Release 12.1.0 for KVM :zap:

In this release:

- NGINX Ingress Controller LoadBalancer Service external traffic policy has been made configurable.
- Management of NGINX IC LoadBalancer Service is moved from azure-operator to NGINX IC App itself to:
  - Enable external traffic policy configurability, in a way consistent with other NGINX IC configuration options.
  - Lay the foundation for making NGINX IC App optional and not pre-installed in a future Azure platform release.

Along with azure-operator and NGINX IC, minor improvements were also made to CoreDNS, cluster-operator, Flatcar, and Kubernetes.

Below, you can find more details on components that were changed with this release.

### cluster-operator [0.23.13](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.13)

- Enabled NGINX App managed NodePort Service on KVM.

### kvm-operator [v3.12.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.12.0)

- Stop provisioning NGINX ingress controller NodePort Service.

### nginx-ingress-controller v0.34.1 ([Giant Swarm app v1.8.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v181-2020-07-28))

- Made NGINX NodePort Service external traffic policy configurable.
- Made NGINX NodePort Service node ports configurable.
