**Workload cluster release v19.0.0-beta1 for AWS is available**. This beta release introduces Kubernetes 1.24, replaces the use of AWS CNI with Cilium and upgrades most components. This release also deprecates and removes use of `kiam` by default in favor of `IAM Roles for Service Accounts (IRSA)`.

> **_CNI Warning:_** Upgrade includes changes from AWS CNI to Cilium, please check `Cilium breaking changes` section for further details.

> **_IAM Warning:_** The `kiam` application will be removed from clusters in favor of IAM Roles for Service Accounts (IRSA). please check `IRSA breaking changes` section for further details.

> **_Network Policies Warning:_** Cilium implementation of Kubernetes `NetworkPolicies` has known limitations regarding CIDR selectors. If you have NetworkPolicies containing `ipBlock` fields referencing IPs of the cluster nodes and/or Pod IPs and/or Service IPs and/or 0.0.0.0/0 your applications might stop working when upgrading to this release. Please get in touch with your SA for further details.
