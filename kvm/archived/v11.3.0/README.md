## :zap: Tenant Cluster Release 11.3.0 for KVM :zap:

This release includes Kubernetes v1.16.9 as well as reliability and user experience improvements.

In addition, this release replaces CoreOS with Flatcar Container Linux.
CoreOS has gone [end-of-life](https://coreos.com/os/eol/) and is being rapidly phased out.
Flatcar is a compatible fork of CoreOS which receives ongoing support.
To continue receiving security updates and to minimize the effort needed to migrate in the future, we recommend upgrading to this release.

Below, you can find more details on components that were changed with this release.

### Kubernetes [v1.16.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v1168)
- Updated from v1.16.3.

### kvm-operator [v3.11.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.11.0)
- Use Flatcar linux instead of CoreOS.
- Streamlined image templating for core components for quicker and easier releases in the future.
- Support setting OIDC username and groups prefix.
- Enabled per-cluster configuration of kube-proxy's `conntrackMaxPerCore` parameter.

### Flatcar Linux [2345.3.1](https://www.flatcar-linux.org/releases/#release-2345.3.1)
- Updated from CoreOS 2247.6.0.
- Updated Linux Kernel to 4.19.107.
