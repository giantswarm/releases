# :zap: Giant Swarm Release v16.1.0 for KVM :zap:

This release includes changes in calico and etcd that should lead to better performance.

## Change details


### calico [3.20.1](https://github.com/projectcalico/calico/releases/tag/v3.20.1)

#### Bug fixes
 - Updated ubi base images and CentOS repos to stop CVE false positives from being reported. [node #1214](https://github.com/projectcalico/node/pull/1214)
 - calico/node logs write to /var/log/calico within the container by default, in addition to stdout [node #1134](https://github.com/projectcalico/node/pull/1134)
 - Improve error message for conflicting routes in CNI plugin [cni-plugin #1164](https://github.com/projectcalico/cni-plugin/pull/1164)
 - Bugfix: blackhole routing table with No-OIF / InterfaceNone-only is clobbering all other routes in the same routing table because if netlink.RT_FILTER_OIF is specified with a netlink.Route{LinkIndex: 0}, it will return all routes using the remaining applicable filter (netlink.RT_FILTER_TABLE / Table 254) link routes. [felix #2995](https://github.com/projectcalico/felix/pull/2995)
 - Fix slow performance when updating a Kubernetes namespace when there are many Pods (and in turn, slow startup performance when there are many namespaces). [felix #2967](https://github.com/projectcalico/felix/pull/2967)
 - Fixes a benign error caused by attempting to delete the same IPAMBlock twice. [kube-controllers #822](https://github.com/projectcalico/kube-controllers/pull/822)

### etcd [3.5.0](https://github.com/etcd-io/etcd/releases/tag/v3.5.0)

#### etcd server

- Fix corruption bug in defrag.
- Fix quorum protection logic when promoting a learner.
- Improve peer corruption checker to work when peer mTLS is enabled.
- Log `[CLIENT-PORT]/health` check in server side.
- Log successful etcd server-side health check in debug level.
- Improve compaction performance when latest index is greater than 1-million.
- Add log when etcdserver failed to apply command.
- Improve count-only range performance.
- Remove redundant storage restore operation to shorten the startup time.
- Fix deadlock bug in mvcc.
- Fix inconsistency between WAL and server snapshot.
- Improve logging around snapshot send and receive.
- Push down RangeOptions.limit argv into index tree to reduce memory overhead.
- Improve `runtime.FDUsage` call pattern to reduce objects malloc of Memory Usage and CPU Usage.
- Improve mvcc.watchResponse channel Memory Usage.
- Log expensive request info in UnaryInterceptor.
- Improve healthcheck by using v3 range request and its corresponding timeout.
- Fix server panic in slow writes warnings.
- Reduce around 30% memory allocation by logging range response size without marshal.


### kvm-operator [3.18.3](https://github.com/giantswarm/kvm-operator/releases/tag/v3.18.3)

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



