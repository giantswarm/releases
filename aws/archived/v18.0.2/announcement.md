**Workload cluster release v18.0.2 for AWS is available**. This release fixes incompability with Cgroups V1 that was introduced in v18.0.0. The issue came from dropping dockershim and switching to containerd. Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v18.0.2/).

> **_Info:_** Cilium support that was introduced in 18.0.0 was rolled back to AWS CNI because of an unresolved upstream bug. Cilium support is postponed to version v19.0.0.
