**Workload cluster release v19.0.0-alpha1 for AWS is available**. This alpha release introduces Kubernetes 1.24, replaces the use of AWS CNI with Cilium and upgrades most components. **This release is solely to be used for testing in new playground clusters.** Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v19.0.0-alpha1/).

> **_CNI Warning:_** The AWS CNI pod subnets are no longer used by Cilium. Please add custom routes with the node subnet(s) CIDR(s) instead of the AWS CNI pod subnets CIDR before upgrading to this release.

> **_IAM Warning:_** The `kiam` application will be removed from clusters in favor of IAM Roles for Service Accounts (IRSA) solution. Please read release notes and prepare for migration.
