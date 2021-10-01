# :zap: Giant Swarm Release v15.3.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.20.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.11)

#### Bug or Regression
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104448](https://github.com/kubernetes/kubernetes/pull/104448), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104152](https://github.com/kubernetes/kubernetes/pull/104152), [@aojea](https://github.com/aojea)) [SIG Network]
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric ([#100720](https://github.com/kubernetes/kubernetes/pull/100720), [@dntosas](https://github.com/dntosas)) [SIG Apps, Instrumentation, Scheduling and Storage]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104348](https://github.com/kubernetes/kubernetes/pull/104348), [@mauriciopoppe](https://github.com/mauriciopoppe)) [SIG Storage]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104674](https://github.com/kubernetes/kubernetes/pull/104674), [@thockin](https://github.com/thockin)) [SIG Network]
#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### external-dns [2.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.5.0)

#### Changed
- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.



### cert-manager [2.10.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.10.0)

#### Changed
- Upgrade to upstream `v1.5.3` ([#184](https://github.com/giantswarm/cert-manager-app/pull/184)). This is the first version compatible with Kubernetes 1.22.
- Add metadata to enable metrics scraping ([#181](https://github.com/giantswarm/cert-manager-app/pull/181)).



### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

#### Added
- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.
#### Changed
- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.



