# :zap: Giant Swarm Release v19.0.1 for Azure :zap:

This is a maintainance and security release, featuring latest Flatcar Container Linux and Kubernetes versions.
It is now possible to provide a custom user config for `bundle` default apps.
Lastly, we enabled the `CronJobTimeZone` feature, see https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones .

## Change details


### azure-operator [7.1.0](https://github.com/giantswarm/azure-operator/releases/tag/v7.1.0)

#### Changed
- Bump k8scc to support improve reliability of calico deployment process for new clusters.
- Enable `CronJobTimeZone` feature gate.



### kubernetes [1.24.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.10)

#### API Change
- Kubernetes 1.24 is now built with go1.19.4 ([#113956](https://github.com/kubernetes/kubernetes/pull/113956), [@liggitt](https://github.com/liggitt)) [SIG Apps, Architecture, Auth, Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Release, Scheduling, Storage and Testing]
#### Feature
- Kubernetes is now built with Go 1.19.5 ([#115012](https://github.com/kubernetes/kubernetes/pull/115012), [@cpanato](https://github.com/cpanato)) [SIG Release and Testing]
#### Bug or Regression
- Client-go: fixes potential data races retrying requests using a custom io.Reader body; with this fix, only requests with no body or with string / []byte / runtime.Object bodies can be retried ([#113933](https://github.com/kubernetes/kubernetes/pull/113933), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Do not include preemptor pod metadata in the event message ([#115024](https://github.com/kubernetes/kubernetes/pull/115024), [@mimowo](https://github.com/mimowo)) [SIG Scheduling]
- Failed pods associated with a job with `parallelism = 1` are recreated by the job controller honoring exponential backoff delay again. However, for jobs with `parallelism > 1`, pods might be created without exponential backoff delay. ([#115021](https://github.com/kubernetes/kubernetes/pull/115021), [@nikhita](https://github.com/nikhita)) [SIG Apps]
- Fix a regression that the scheduler always goes through all Filter plugins. ([#114526](https://github.com/kubernetes/kubernetes/pull/114526), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix bug in CRD Validation Rules (beta) and ValidatingAdmissionPolicy (alpha) where all admission requests could result in `internal error: runtime error: index out of range [3] with length 3 evaluating rule: <rule name>` under certain circumstances. ([#114865](https://github.com/kubernetes/kubernetes/pull/114865), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery]
- Fix performance issue when creating large objects using SSA with fully unspecified schemas (preserveUnknownFields). ([#111915](https://github.com/kubernetes/kubernetes/pull/111915), [@aojea](https://github.com/aojea)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
- Fixed StatefulSet to show the valid status even if the new replica creation fails. ([#112083](https://github.com/kubernetes/kubernetes/pull/112083), [@gjkim42](https://github.com/gjkim42)) [SIG Apps and Testing]
- Fixing issue in Winkernel Proxier - Unexpected active TCP connection drops while horizontally scaling the endpoints for a LoadBalancer Service with External Traffic Policy: Local ([#114040](https://github.com/kubernetes/kubernetes/pull/114040), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Fixing issue with Winkernel Proxier - No ingress load balancer rules with endpoints to support load balancing when all the endpoints are terminating. ([#114451](https://github.com/kubernetes/kubernetes/pull/114451), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Kube-apiserver: bugfix DeleteCollection API fails if request body is non-empty ([#113968](https://github.com/kubernetes/kubernetes/pull/113968), [@sxllwx](https://github.com/sxllwx)) [SIG API Machinery]
- Optimizing loadbalancer creation with the help of attribute Internal Traffic Policy: Local ([#114466](https://github.com/kubernetes/kubernetes/pull/114466), [@princepereira](https://github.com/princepereira)) [SIG Network]
- Update the system-validators library to v1.8.0 ([#114060](https://github.com/kubernetes/kubernetes/pull/114060), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- [aws] Fixed a bug which reduces the number of unnecessary calls to STS in the event of assume role failures in the legacy cloud provider ([#110706](https://github.com/kubernetes/kubernetes/pull/110706), [@prateekgogia](https://github.com/prateekgogia)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cel-go: [v0.10.2 → v0.10.4](https://github.com/google/cel-go/compare/v0.10.2...v0.10.4)
- k8s.io/system-validators: v1.7.0 → v1.8.0
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.33 → v0.0.35
- sigs.k8s.io/structured-merge-diff/v4: v4.2.1 → v4.2.3
#### Removed
_Nothing has changed._



### etcd [3.5.7](https://github.com/etcd-io/etcd/releases/tag/v3.5.7)

#### etcd server
- Fix [Remove memberID from data corrupt alarm](https://github.com/etcd-io/etcd/pull/14852).
- Fix [Allow non mutating requests pass through quotaKVServer when NOSPACE](https://github.com/etcd-io/etcd/pull/14884).
- Fix [nil pointer panic for readonly txn due to nil response](https://github.com/etcd-io/etcd/pull/14899).
- Fix [The last record which was partially synced to disk isn't automatically repaired](https://github.com/etcd-io/etcd/pull/15069).
- Fix [etcdserver might promote a non-started learner](https://github.com/etcd-io/etcd/pull/15096).
#### Package `clientv3`
- Reverted the fix to [auth invalid token and old revision errors in watch](https://github.com/etcd-io/etcd/pull/14995).
#### Security
- Use [distroless base image](https://github.com/etcd-io/etcd/pull/15016) to address critical Vulnerabilities.
- Updated [base image from base-debian11 to static-debian11 and removed dependency on busybox](https://github.com/etcd-io/etcd/pull/15037).
- Bumped [some dependencies](https://github.com/etcd-io/etcd/pull/15018) to address some HIGH Vulnerabilities.
#### Go
- Require [Go 1.17+](https://github.com/etcd-io/etcd/pull/15019).
- Compile with [Go 1.17+](https://go.dev/doc/devel/release#go1.17)



### app-operator [6.6.0](https://github.com/giantswarm/app-operator/releases/tag/v6.6.0)

#### Added
- Add support for dependencies between apps using `app-operator.giantswarm.io/depends-on` annotation.



### cluster-operator [5.5.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.5.0)

#### Fixed
- Fix user config CM mapping for bundle apps.
#### Added
- Read app dependencies from Release CR to avoid deadlock installing apps in new clusters.



### containerlinux [3374.2.3](https://www.flatcar-linux.org/releases/#release-3374.2.3)

_Changes since **Stable 3374.2.2**_

#### Security fixes:

- Linux ([CVE-2022-3169](https://nvd.nist.gov/vuln/detail/CVE-2022-3169), [CVE-2022-3344](https://nvd.nist.gov/vuln/detail/CVE-2022-3344), [CVE-2022-3424](https://nvd.nist.gov/vuln/detail/CVE-2022-3424), [CVE-2022-3521](https://nvd.nist.gov/vuln/detail/CVE-2022-3521), [CVE-2022-3534](https://nvd.nist.gov/vuln/detail/CVE-2022-3534), [CVE-2022-3545](https://nvd.nist.gov/vuln/detail/CVE-2022-3545), [CVE-2022-3643](https://nvd.nist.gov/vuln/detail/CVE-2022-3643), [CVE-2022-4378](https://nvd.nist.gov/vuln/detail/CVE-2022-4378), [CVE-2022-45869](https://nvd.nist.gov/vuln/detail/CVE-2022-45869), [CVE-2022-45934](https://nvd.nist.gov/vuln/detail/CVE-2022-45934), [CVE-2022-47518](https://nvd.nist.gov/vuln/detail/CVE-2022-47518), [CVE-2022-47519](https://nvd.nist.gov/vuln/detail/CVE-2022-47519), [CVE-2022-47520](https://nvd.nist.gov/vuln/detail/CVE-2022-47520), [CVE-2022-47521](https://nvd.nist.gov/vuln/detail/CVE-2022-47521))
- git ([CVE-2022-23521](https://nvd.nist.gov/vuln/detail/CVE-2022-23521), [CVE-2022-41903](https://nvd.nist.gov/vuln/detail/CVE-2022-41903))

#### Bug fixes:

- Fix "ext4 deadlock under heavy I/O load" kernel issue. The patch for this is included provisionally while we stay with Kernel 5.15.86. ([Flatcar#847](https://github.com/flatcar/Flatcar/issues/847), [coreos-overlay#2402](https://github.com/flatcar/coreos-overlay/pull/2402))

#### Changes:


#### Updates:

- Linux ([5.15.86](https://lwn.net/Articles/918808) (includes [5.15.85](https://lwn.net/Articles/918329), [5.15.84](https://lwn.net/Articles/918206), [5.15.83](https://lwn.net/Articles/917896), [5.15.82](https://lwn.net/Articles/917400), [5.15.81](https://lwn.net/Articles/916763), [5.15.80](https://lwn.net/Articles/916003)))
- git ([2.37.5](https://github.com/git/git/blob/v2.37.5/Documentation/RelNotes/2.37.5.txt))



### external-dns [2.23.2](https://github.com/giantswarm/external-dns-app/releases/tag/v2.23.2)

#### Fixed
- Hardcode `external-dns.name` default name dropping the `-app` suffix ([#235](https://github.com/giantswarm/external-dns-app/pull/235))



### observability-bundle [0.1.9](https://github.com/giantswarm/observability-bundle/releases/tag/v0.1.9)

#### Changed
- Upgrade `prometheus-operator-app` to 3.0.0.
- Upgrade `prometheus-operator-crd` to 3.0.0.



