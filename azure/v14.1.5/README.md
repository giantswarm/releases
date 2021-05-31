# :zap: Giant Swarm Release v14.1.5 for Azure :zap:

This workload cluster release contains fixes for azure-operator handling multi-tenant service principal secrets as a prerequisite for migration from VPN Gateway to VNET peering. Moreover, we have extended azure-scheduled-events app with additional Azure VMSS events handling. 

Please contact your Solution Architect in order to validate if this release is necessary for your use case.

## Change details


### azure-operator [5.5.3](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.3)

#### Fixed
- Fix wrong setup of multi-account service principals.



### kubernetes [1.19.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.10)

#### API Change
- Fixes using server-side apply with APIService resources ([#100713](https://github.com/kubernetes/kubernetes/pull/100713), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Apps, Scheduling and Testing]
- Regenerate protobuf code to fix CVE-2021-3121 ([#100515](https://github.com/kubernetes/kubernetes/pull/100515), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
#### Feature
- Kubernetes is now built using go1.15.10 ([#100520](https://github.com/kubernetes/kubernetes/pull/100520), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Fixed a bug where a high churn of events was causing master instability by reducing the maximum number of objects (events) attached to a single etcd lease. ([#100450](https://github.com/kubernetes/kubernetes/pull/100450), [@mborsz](https://github.com/mborsz)) [SIG API Machinery and Instrumentation]
- Fixed a race condition on API server startup ensuring previously created webhook configurations are effective before the first write request is admitted. ([#95783](https://github.com/kubernetes/kubernetes/pull/95783), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixes a data race issue in the priority and fairness API server filter ([#100669](https://github.com/kubernetes/kubernetes/pull/100669), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Kubectl: Fixed panic when describing an ingress backend without an API Group ([#100542](https://github.com/kubernetes/kubernetes/pull/100542), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- Reverts breaking change to inline AzureFile volumes in v1.19.7-v1.19.9; referenced secrets are now correctly searched for in the same namespace as the pod as in previous releases. ([#100398](https://github.com/kubernetes/kubernetes/pull/100398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- The endpointslice mirroring controller mirrors endpoints annotations and labels to the generated endpoint slices, it also ensures that updates on any of these fields on the endpoints are mirrored. 
  The well-known annotation endpoints.kubernetes.io/last-change-trigger-time is skipped and not mirrored. ([#100443](https://github.com/kubernetes/kubernetes/pull/100443), [@aojea](https://github.com/aojea)) [SIG Apps, Network and Testing]
- The maximum number of ports allowed in EndpointSlices has been increased from 100 to 20,000 ([#99795](https://github.com/kubernetes/kubernetes/pull/99795), [@robscott](https://github.com/robscott)) [SIG Network]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/yuin/goldmark: [v1.1.27 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.1.27...v1.2.1)
- golang.org/x/sync: cd5d95a → 67f06af
- golang.org/x/tools: c1934b7 → 113979e
- golang.org/x/xerrors: 9bdfabe → 5ec99f8
- sigs.k8s.io/structured-merge-diff/v4: v4.0.1 → v4.0.3
#### Removed
_Nothing has changed._



### containerlinux [2765.2.2](https://www.flatcar-linux.org/releases/#release-2765.2.2)

**Security fixes**
*   Linux ([CVE-2021-27365](https://nvd.nist.gov/vuln/detail/CVE-2021-27365), [CVE-2021-27364](https://nvd.nist.gov/vuln/detail/CVE-2021-27364), [CVE-2021-27363](https://nvd.nist.gov/vuln/detail/CVE-2021-27363), [CVE-2021-28038](https://nvd.nist.gov/vuln/detail/CVE-2021-28038),[CVE-2021-28039](https://nvd.nist.gov/vuln/detail/CVE-2021-28039), [CVE-2021-28375](https://nvd.nist.gov/vuln/detail/CVE-2021-28375), [CVE-2021-28660](https://nvd.nist.gov/vuln/detail/CVE-2021-28660), [CVE-2021-27218](https://nvd.nist.gov/vuln/detail/CVE-2021-27218), [CVE-2021-27219](https://nvd.nist.gov/vuln/detail/CVE-2021-27219))
*   openssl ([CVE-2021-23840](https://nvd.nist.gov/vuln/detail/CVE-2021-23840),[ CVE-2021-23841](https://nvd.nist.gov/vuln/detail/CVE-2021-23841), [CVE-2020-1971](https://nvd.nist.gov/vuln/detail/CVE-2020-1971),[ CVE-2021-23840](https://nvd.nist.gov/vuln/detail/CVE-2021-23840),[ CVE-2021-23841](https://nvd.nist.gov/vuln/detail/CVE-2021-23841), [CVE-2021-3449](https://nvd.nist.gov/vuln/detail/CVE-2021-3449),[ CVE-2021-3450](https://nvd.nist.gov/vuln/detail/CVE-2021-3450))

**Bug Fixes**
*   GCE: The old interface name ens4v1 which was replaced by eth0 due to a broken udev rule was restored, but now as alternative interface name, and eth0 will stay the primary name for consistency across cloud environments. ([init#38](https://github.com/kinvolk/init/pull/38))

**Changes**
*   The virtio network interfaces got predictable interface names as alternative interface names, and thus these names can also be used to match for a specific interface in case there is more than one and the eth0 and eth1 name assignment is not stable. ([init#38](https://github.com/kinvolk/init/pull/38))

**Updates**
*   Linux ([5.10.25](https://lwn.net/Articles/849951/))
*   openssl ([1.1.1k](https://mta.openssl.org/pipermail/openssl-announce/2021-March/000197.html))
*   open-iscsi ([2.1.4](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.4))


### azure-scheduled-events [0.4.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.4.0)

#### Added
- React to `Preempt`, `Reboot` and `Redeploy` events to drain nodes properly.
#### Change
- Use the `NotBefore` field from the event to calculate drain timeout.



