# :zap: Giant Swarm Release v15.1.0 for KVM :zap:

This release includes changes in calico and etcd that should lead to better performance.

## Change details


### calico [3.20.1](https://github.com/projectcalico/calico/releases/tag/v3.20.1)

#### Bug fixes
- Updated ubi base images and CentOS repos to stop CVE false positives from being reported.
- calico/node logs write to /var/log/calico within the container by default, in addition to stdout.
- Improve error message for conflicting routes in CNI plugin.
- Bugfix: blackhole routing table with No-OIF / InterfaceNone-only is clobbering all other routes in the same routing table because if netlink.RT_FILTER_OIF is specified with a netlink.Route{LinkIndex: 0}, it will return all routes using the remaining applicable filter (netlink.RT_FILTER_TABLE / Table 254) link routes.
- Fix slow performance when updating a Kubernetes namespace when there are many Pods (and in turn, slow startup performance when there are many namespaces).
- Fixes a benign error caused by attempting to delete the same IPAMBlock twice.
- Fix that calico/node would fail to set NetworkUnavailable to false for etcd clusters with mismatched node names.
- Stop ARP traffic being dropped due to RPF check.
- Disable VXLAN tunnel checksum offload on kernels < v5.7.
- Improve routing loop prevention to handle when advertising Service LoadBalancer IPs.
- Retry setting AWS EC2 source/destination check until successful.
- Install blackhole routes in VXLAN mode.
- Fix that podIP annotation could be incorrectly clobbered for stateful set pods.
- Reinstates logic that falls back to the status of the pod during termination if the pod IP annotation is not set by the Calico CNI plugin.
- Fix issue with serviceaccount names larger than 63 characters.
- Fix error parsing pod deletion updates in kube-controllers.

#### Other changes

- calico/node marks nodes with NetworkUnavailable=true on shutdown node.
- Typha now gives newly connected clients an extra grace period to catch up after sending the snapshot. Should reduce the possibility of cyclic disconnects.
- Added enhanced error logging for IPAM failures.
- Add IP address garbage collection to kube-controllers.
- Calico will now release empty IPAM blocks from nodes that no longer need them so they can be used elsewhere.
- Mount CNI plugin directory into calico/node to enable configuration updates.

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

### kubernetes [1.20.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.11)

#### Bug or Regression
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104448](https://github.com/kubernetes/kubernetes/pull/104448))
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104152](https://github.com/kubernetes/kubernetes/pull/104152))
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric ([#100720](https://github.com/kubernetes/kubernetes/pull/100720)) [SIG Apps, Instrumentation, Scheduling and Storage]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104348](https://github.com/kubernetes/kubernetes/pull/104348))
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104674](https://github.com/kubernetes/kubernetes/pull/104674))

#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958)) [SIG API Machinery and Node]

### kvm-operator [3.18.2](https://github.com/giantswarm/kvm-operator/releases/tag/v3.18.2)

#### Changed
- Disable apiserver flow control to mitigate etcd memory usage issues temporarily.
