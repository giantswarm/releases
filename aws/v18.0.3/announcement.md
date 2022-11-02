This release fixes a misconfiguration of IRSA that caused downtime in the API server during upgrades between v17 and v18 with IRSA feature enabled. Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v18.0.3/).

> **_Info:_** Cilium support that was introduced in 18.0.0 was rolled back to AWS CNI because of an unresolved upstream bug. Cilium support is postponed to version v19.0.0.
