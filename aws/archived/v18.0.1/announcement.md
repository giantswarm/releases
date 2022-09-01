**Workload cluster release v18.0.1 for AWS is available**. It provides Kubernetes 1.23 and upgrades most components. CloudFront and a private S3 bucket are now used by IRSA in non-China regions. Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v18.0.1/).

> **_Info:_** Cilium support that was introduced in 18.0.0 was rolled back to AWS CNI because of an unresolved upstream bug. Cilium support is postponed to version v19.0.0.
