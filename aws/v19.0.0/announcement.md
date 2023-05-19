**Workload cluster release v19.0.0 for AWS is available**. This release introduces Kubernetes 1.24, replaces the use of AWS CNI with Cilium and upgrades most components. We are also deprecating and removing use of `kiam` by default in favor of `IAM Roles for Service Accounts (IRSA)`. Please check [IAM permissions requirements](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#330---2023-05-11). Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v19.0.0/).

> **Warning:** This is very critical upgrade, follow the release notes for all important details. Your Account Engineer will reach out to you and prepare a checklist to avoid downtime.

> **_CNI Warning:_** Upgrade includes changes from AWS CNI to Cilium, please check `Cilium` section for further details.

> **_IAM Warning:_** The `kiam` application will be removed from clusters in favor of IAM Roles for Service Accounts (IRSA), please check `IRSA` section for further details.
