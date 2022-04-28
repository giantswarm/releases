# :zap: Giant Swarm Release v15.2.2 for AWS :zap:

This release provides a security fix for [CVE-2021-25741](https://nvd.nist.gov/vuln/detail/CVE-2021-25741).

## Change details


### kubernetes [1.20.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.11)

#### Bug or Regression
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
