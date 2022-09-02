**Workload cluster release v18.0.0 for AWS is available**. It provides Kubernetes 1.23, replaces the use of AWS CNI with Cilium and upgrades most components. CloudFront and a private S3 bucket are now used by IRSA in non-China regions. Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v18.0.0/).

> **_Warning:_** The AWS CNI pod subnets are no longer used by Cilium. Please add custom routes with the node subnet(s) CIDR(s) instead of the AWS CNI pod subnets CIDR before upgrading to this release.
