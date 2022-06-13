# :zap: Giant Swarm Release v17.1.0 for Azure :zap:

<< Add description here >>

## Change details


### app-operator [5.12.0](https://github.com/giantswarm/app-operator/releases/tag/v5.12.0)

#### Added
- Add `initialBootstrapMode` flag to allow deploying CNI as managed apps.



### azure-operator [5.20.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.20.0)

#### Changed
- Bumped k8scc to latest version to fix `localhost` node name problem.



### cert-operator [2.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v2.0.1)

#### Fixed
- Bump go module major version.



### containerlinux [3139.2.1](https://www.flatcar-linux.org/releases/#release-3139.2.1)

New **Stable** Release **3139.2.1**

_Changes since **Stable 3139.2.0**_

#### Security fixes:

- Linux ([CVE-2022-28390](https://nvd.nist.gov/vuln/detail/CVE-2022-28390), [CVE-2022-0168](https://nvd.nist.gov/vuln/detail/CVE-2022-0168), [CVE-2022-1158](https://nvd.nist.gov/vuln/detail/CVE-2022-1158), [CVE-2022-1353](https://nvd.nist.gov/vuln/detail/CVE-2022-1353), [CVE-2022-1198](https://nvd.nist.gov/vuln/detail/CVE-2022-1198), [CVE-2022-28389](https://nvd.nist.gov/vuln/detail/CVE-2022-28389), [CVE-2022-28388](https://nvd.nist.gov/vuln/detail/CVE-2022-28388), [CVE-2022-1516](https://nvd.nist.gov/vuln/detail/CVE-2022-1516), [CVE-2022-1263](https://nvd.nist.gov/vuln/detail/CVE-2022-1263), [CVE-2022-29582](https://nvd.nist.gov/vuln/detail/CVE-2022-29582), [CVE-2022-1204](https://nvd.nist.gov/vuln/detail/CVE-2022-1204), [CVE-2022-1205](https://nvd.nist.gov/vuln/detail/CVE-2022-1205), [CVE-2022-0500](https://nvd.nist.gov/vuln/detail/CVE-2022-0500), [CVE-2022-23222](https://nvd.nist.gov/vuln/detail/CVE-2022-23222))
- nvidia-drivers ([CVE-2022-21814](https://nvd.nist.gov/vuln/detail/CVE-2022-21814), [CVE-2022-21813](https://nvd.nist.gov/vuln/detail/CVE-2022-21813))
- Go ([CVE-2022-24675](https://nvd.nist.gov/vuln/detail/CVE-2022-24675))

#### Bug fixes:

- AWS: specify correct console (ttyS0) on kernel command line for ARM64 instances ([coreos-overlay#1628](https://github.com/flatcar-linux/coreos-overlay/pull/1628))
- GCE: Restored oem-gce.service functionality on GCP ([coreos-overlay#1813](https://github.com/flatcar-linux/coreos-overlay/pull/1813))
- Added pahole to developer container, without it kernel modules built against /usr/src/linux may fail to probe with an 'invalid relocation target' error ([coreos-overlay#1839](https://github.com/flatcar-linux/coreos-overlay/pull/1839))

#### Changes:

- Merge the Flatcar Pro features into the regular Flatcar images ([coreos-overlay#1679](https://github.com/flatcar-linux/coreos-overlay/pull/1679)) 
- GCE: Enabled GVE kernel driver, which adds support for Google Virtual NIC on GCP ([coreos-overlay#1802](https://github.com/flatcar-linux/coreos-overlay/pull/1802))
- SDK: Dropped the mantle binaries (kola, ore, etc.) from the SDK, they are now provided by the `ghcr.io/flatcar-linux/mantle` image ([coreos-overlay#1827](https://github.com/flatcar-linux/coreos-overlay/pull/1827), [scripts#275](https://github.com/flatcar-linux/scripts/pull/275))

#### Updates:

- Linux ([5.15.37](https://lwn.net/Articles/893264) (includes [5.15.36](https://lwn.net/Articles/892812), [5.15.35](https://lwn.net/Articles/892002), [5.15.34](https://lwn.net/Articles/891251), [5.15.33](https://lwn.net/Articles/890722)))
- Go ([1.17.9](https://go.googlesource.com/go/+/refs/tags/go1.17.9))
- ca-certificates ([3.78](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_78.html))
- nvidia-drivers ([510.47.03](https://docs.nvidia.com/datacenter/tesla/tesla-release-notes-510-47-03/index.html)) 
- GCE: google compute-image-packages ([20190124](https://github.com/GoogleCloudPlatform/compute-image-packages/releases/tag/20190124))


### calico [3.21.5](https://github.com/projectcalico/calico/releases/tag/v3.21.5)

Not found


### etcd [3.5.4](https://github.com/etcd-io/etcd/releases/tag/v3.5.4)

Not found


### kubernetes [1.22.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.10)

#### Bug or Regression
- Correct event registration for multiple scheduler plugins; this fixes a potential significant delay in re-queueing unschedulable pods. ([#109447](https://github.com/kubernetes/kubernetes/pull/109447), [@ahg-g](https://github.com/ahg-g)) [SIG Scheduling and Testing]
- Existing InTree AzureFile PVs which don't have a secret namespace defined will now work properly after enabling CSI migration - the namespace will be obtained from ClaimRef. ([#108000](https://github.com/kubernetes/kubernetes/pull/108000), [@RomanBednar](https://github.com/RomanBednar)) [SIG Cloud Provider and Storage]
- Failure to start a container cannot accidentally result in the pod being considered "Succeeded" in the presence of deletion. ([#108883](https://github.com/kubernetes/kubernetes/pull/108883), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Kubeadm: add the flag "--experimental-initial-corrupt-check" to etcd static Pod manifests to ensure etcd member data consistency ([#109076](https://github.com/kubernetes/kubernetes/pull/109076), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### cert-exporter [2.2.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.2.0)

#### Changed
- Change priorityClass to `system-node-critical` for the daemonset.



### chart-operator [2.24.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0)

#### Added
- Split Helm client into private Helm client for `giantswarm`-namespaced apps and public Helm client for rest of the apps.
- Add Helm release failure reason when it is known, and if there is a currently successfully released version

### Changed
- Add chart-pull-failed error to differentiate between issues when pulling chart tarball and other problems.
- Always create giantswarm-critical priority class if it does not exist.
- Update `helmclient` to v4.10.0.

### Fixed
- Fix missing PriorityClass issue.


### coredns [1.9.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.9.0)

#### Added
- Add toleration for `node.cloudprovider.kubernetes.io/uninitialized`.
#### Changed
- Update `coredns` to upstream version [1.8.7](https://coredns.io/2021/12/09/coredns-1.8.7-release/).
 



### external-dns [2.14.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.14.0)

#### Added
- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.



### metrics-server [1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set `kubelet-preferred-address-types` to `Hostname` on `AWS`.



### net-exporter [1.12.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.12.0)

#### Changed
- Use parameter for CoreDNS namespace (defaulted to kube-system)



### node-exporter [1.12.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.12.0)

#### Changed
- Enabled `diskstats` collector.



### azure-scheduled-events [0.7.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.7.0)

#### Added
- Add Vertical Pod Autoscaler CR.



### vertical-pod-autoscaler [2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

#### Changed
- Use patched docker image tagged `0.10.0-oomfix` for `recommender` and updater (see https://github.com/giantswarm/roadmap/issues/923).



### vertical-pod-autoscaler-crd [1.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v1.0.1)

#### Added
- Add cluster singleton restriction so app can only be installed once.



### kube-state-metrics [1.10.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.10.0)

#### Changed
- Add `Node Pool` labels to the default allowed labels in `--metric-labels-allowlist`.



