# :zap: Giant Swarm Release v15.1.0 for KVM :zap:

This release includes changes in calico and etcd that should lead to better performance.

## Change details


### calico [3.20.1](https://github.com/projectcalico/calico/releases/tag/v3.20.1)

#### Bug fixes
 - Updated ubi base images and CentOS repos to stop CVE false positives from being reported. [node #1214](https://github.com/projectcalico/node/pull/1214) (@mgleung)
 - calico/node logs write to /var/log/calico within the container by default, in addition to stdout [node #1134](https://github.com/projectcalico/node/pull/1134) (@song-jiang)
 - Improve error message for conflicting routes in CNI plugin [cni-plugin #1164](https://github.com/projectcalico/cni-plugin/pull/1164) (@mgleung)
 - Bugfix: blackhole routing table with No-OIF / InterfaceNone-only is clobbering all other routes in the same routing table because if netlink.RT_FILTER_OIF is specified with a netlink.Route{LinkIndex: 0}, it will return all routes using the remaining applicable filter (netlink.RT_FILTER_TABLE / Table 254) link routes. [felix #2995](https://github.com/projectcalico/felix/pull/2995) (@electricjesus)
 - Fix slow performance when updating a Kubernetes namespace when there are many Pods (and in turn, slow startup performance when there are many namespaces). [felix #2967](https://github.com/projectcalico/felix/pull/2967) (@fasaxc)
 - Fixes a benign error caused by attempting to delete the same IPAMBlock twice. [kube-controllers #822](https://github.com/projectcalico/kube-controllers/pull/822) (@caseydavenport)



### etcd [3.5.0](https://github.com/etcd-io/etcd/releases/tag/v3.5.0)

Not found


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



### kvm-operator [3.18.2](https://github.com/giantswarm/kvm-operator/releases/tag/v3.18.2)

#### Changed
- Disable apiserver flow control to mitigate etcd memory usage issues temporarily.



