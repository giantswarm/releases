# :zap: Giant Swarm Release v17.0.0 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.21.5](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.5)

#### Feature
- Kubernetes is now built with Golang 1.16.8 ([#104906](https://github.com/kubernetes/kubernetes/pull/104906), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Fix NodeAuthenticator tests in dualstack ([#104840](https://github.com/kubernetes/kubernetes/pull/104840), [@ardaguclu](https://github.com/ardaguclu)) [SIG Auth and Testing]
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104447](https://github.com/kubernetes/kubernetes/pull/104447), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fixed occasional pod cgroup freeze when using cgroup v1 and systemd driver.
  Fixed "failed to create container ... unit already exists" when using cgroup v1 and systemd driver. ([#104530](https://github.com/kubernetes/kubernetes/pull/104530), [@kolyshkin](https://github.com/kolyshkin)) [SIG CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104151](https://github.com/kubernetes/kubernetes/pull/104151), [@aojea](https://github.com/aojea)) [SIG Network]
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric ([#100720](https://github.com/kubernetes/kubernetes/pull/100720), [@dntosas](https://github.com/dntosas)) [SIG Apps, Instrumentation, Scheduling and Storage]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104347](https://github.com/kubernetes/kubernetes/pull/104347), [@mauriciopoppe](https://github.com/mauriciopoppe)) [SIG Storage]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104673](https://github.com/kubernetes/kubernetes/pull/104673), [@thockin](https://github.com/thockin)) [SIG Network]
#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]
#### Dependencies
#### Added
- github.com/bits-and-blooms/bitset: [v1.2.0](https://github.com/bits-and-blooms/bitset/tree/v1.2.0)
#### Changed
- github.com/cilium/ebpf: [v0.5.0 → v0.6.2](https://github.com/cilium/ebpf/compare/v0.5.0...v0.6.2)
- github.com/coreos/go-systemd/v22: [v22.3.1 → v22.3.2](https://github.com/coreos/go-systemd/v22/compare/v22.3.1...v22.3.2)
- github.com/golang/protobuf: [v1.4.3 → v1.5.0](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.0)
- github.com/google/go-cmp: [v0.5.4 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.4...v0.5.5)
- github.com/opencontainers/runc: [v1.0.0-rc95 → v1.0.2](https://github.com/opencontainers/runc/compare/v1.0.0-rc95...v1.0.2)
- github.com/opencontainers/selinux: [v1.8.0 → v1.8.2](https://github.com/opencontainers/selinux/compare/v1.8.0...v1.8.2)
- github.com/sirupsen/logrus: [v1.7.0 → v1.8.1](https://github.com/sirupsen/logrus/compare/v1.7.0...v1.8.1)
- google.golang.org/protobuf: v1.25.0 → v1.26.0
#### Removed
- github.com/willf/bitset: [v1.1.11](https://github.com/willf/bitset/tree/v1.1.11)



