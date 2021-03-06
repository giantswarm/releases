## :zap: Tenant Cluster Release 11.3.0 for AWS :zap:

This release includes Kubernetes v1.16.9 as well as reliability and user experience improvements.

In addition, this release [replaces CoreOS with Flatcar Container Linux](https://www.giantswarm.io/blog/time-to-catch-a-new-train-flatcar-linux).
CoreOS has gone [end-of-life](https://coreos.com/os/eol/) and is being rapidly phased out.
Flatcar is a compatible fork of CoreOS which receives ongoing support.
To continue receiving security updates and to minimize the effort needed to migrate in the future, we recommend upgrading to this release.

Below, you can find more details on components that were changed with this release.

### Kubernetes [v1.16.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v1168)
- Updated from v1.16.3.

### aws-operator [v8.5.0](https://github.com/giantswarm/aws-operator/releases/tag/v8.5.0)
- Use Flatcar linux instead of CoreOS.
- Support setting OIDC username and groups prefix.
- Enabled per-cluster configuration of kube-proxy's `conntrackMaxPerCore` parameter.

### Flatcar Container Linux [2345.3.1](https://www.flatcar-linux.org/releases/#release-2345.3.1)
- Updated from CoreOS Container Linux 2191.5.0.
- Updated Linux Kernel to 4.19.107.
