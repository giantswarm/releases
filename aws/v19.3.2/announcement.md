**Workload cluster release v19.3.2 for AWS is available**. This patch release includes the latest stable Flatcar release containing important CVE patches, most notably the Leaky Vessels [CVE-2024-21626](https://nvd.nist.gov/vuln/detail/CVE-2024-21626) affecting `containerd`. It also includes an update to the `security-bundle` intended to improve upgrade stability. 

**Note:** This release does _not_ include patches for the other 3 Leaky Vessels BuildKit CVEs affecting `docker`. Normal cluster operations (e.g. pulling images) do not use `docker`, but if you run untrusted workloads which mount or use `docker` directly, please get in touch to discuss other mitigation options.

Further details can be found in the [release notes](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v19.3.2/).

