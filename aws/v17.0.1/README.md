# :zap: Giant Swarm Release v17.0.1 for AWS :zap:

This release allows one replica of `coredns` to run on the control plane nodes for clusters without any node pools.

> **_Warning:_** Kubernetes v1.22 removed certain APIs and features. More details are available in the [upstream blog post](https://kubernetes.io/blog/2021/07/14/upcoming-changes-in-kubernetes-1-22/).

**Known Issues**
- Java applications are unable to identify memory limits when using a `JRE` prior to v15 in a Control Groups v2 environment. Support was added in `JRE` v15 and later. More details are available in the [upstream issue](https://bugs.openjdk.java.net/browse/JDK-8230305). We recommend using the latest LTS JRE available (currently v17) to ensure continued compatibility with future releases.

**Control Groups v1**
To ensure a smooth transition, in case you need time to modify applications to make them compatible with Control Groups v2, we provide a mechanism that will allow using Control Groups v1 on specific node pools. More details are available in our [documentation](https://docs.giantswarm.io/advanced/forcing-cgroupsv1/).

## Change details


### coredns [1.8.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.8.0)

#### Changed
- Add deployment to run one replica of coredns in master nodes (for clusters with no node pools).



