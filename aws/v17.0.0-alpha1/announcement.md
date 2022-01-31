**Workload cluster release v17.0.0-alpha1 for AWS is available**. It provides support for Kubernetes 1.22 and has Control Groups v2 enabled by default. Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v17.0.0-alpha1/).

> **_Warning:_** This is an **`alpha preview release`** intended only for testing Kubernetes v1.22 changes and Control Groups v2 compatibility. It doesn't have any support for upgrades to and from.

> **_Warning:_** Kubernets v1.22 removed certain APIs and features. More details are available in the [upstream blog post](https://kubernetes.io/blog/2021/07/14/upcoming-changes-in-kubernetes-1-22/).

**Known Issues**
- Java applications are unable to identify memory limits when using a `JRE` prior to v15 in a Control Groups v2 environment. Support was added in `JRE` v15 and later. More details are available in the [upstream issue](https://bugs.openjdk.java.net/browse/JDK-8230305). We recommend you to use the latest LTS JRE available (currently v17).
