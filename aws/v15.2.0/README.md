# :zap: Giant Swarm Release v15.2.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.20.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.10)

#### Feature
- Kubernetes 1.20.x is now built using Go 1.15.15 ([#104215](https://github.com/kubernetes/kubernetes/pull/104215), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Disable aufs module for gce clusters ([#103831](https://github.com/kubernetes/kubernetes/pull/103831), [@lizhuqi](https://github.com/lizhuqi)) [SIG Cloud Provider]
- Fix kube-apiserver metric reporting for the deprecated watch path of /api/<version>/watch/... ([#104191](https://github.com/kubernetes/kubernetes/pull/104191), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- Fix: Provide IPv6 support for internal load balancer ([#103794](https://github.com/kubernetes/kubernetes/pull/103794), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: ignore not a VMSS error for VMAS nodes in reconcileBackendPools ([#103997](https://github.com/kubernetes/kubernetes/pull/103997), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: return empty VMAS name if using standalone VM ([#103470](https://github.com/kubernetes/kubernetes/pull/103470), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed a bug that scheduler extenders are not called on preemptions ([#103019](https://github.com/kubernetes/kubernetes/pull/103019), [@ordovicia](https://github.com/ordovicia)) [SIG Scheduling]
- Fixes an issue cleaning up CertificateSigningRequest objects with an unparseable `status.certificate` field ([#103949](https://github.com/kubernetes/kubernetes/pull/103949), [@liggitt](https://github.com/liggitt)) [SIG Apps and Auth]
- Fixes issue with websocket-based watches of Service objects not closing correctly on timeout ([#102542](https://github.com/kubernetes/kubernetes/pull/102542), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.19 → v0.0.22
#### Removed
_Nothing has changed._

### aws-operator [10.7.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.7.0)

### Added

- Add security settings to S3 bucket to comply with aws policies `s3-bucket-public-read-prohibited,s3-bucket-ssl-requests-only,s3-bucket-public-write-prohibited,s3-bucket-server-side-encryption-enabled,s3-bucket-logging-enabled`, `aws-operator` will need additonal permissions `s3:PutBucketPublicAccessBlock` and `s3:PutBucketPolicy`.
