# :zap: Giant Swarm Release v17.1.0 for Azure :zap:

<< Add description here >>

## Change details


### app-operator [6.0.1](https://github.com/giantswarm/app-operator/releases/tag/v6.0.1)

#### Added
- Add support for Catalogs that define multiple repository mirrors to be used in case some of them are unreachable.
#### Changed
- Only run `PodMonitor` outside of bootstrap mode.



### azure-operator [5.21.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.21.0)

#### Changed
- Changes to EncryptionConfig in order to work with `encryption-provider-operator`.
#### Fixed
- Add pause annotation before deleting old machinepool and azuremachinepool CRs during migration to non-exp.
- Update ownerReference UIDs during migration to non-exp.
- Avoid updating `AzureCluster` at every reconciliation loop in the `subnet` resource.
- Avoid saving `AzureCluster` status if there are no changes to avoid useless reconciliation loops.



### cluster-operator [4.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.3.0)

#### Changed
- Do not update "app-operator.giantswarm.io/version" label on app-operators when their value is 0.0.0 (aka they are reconciled by the management cluster app-operator). This is a use-case for App Bundles for example, because the App CRs they contain should be created in the MC so should be reconciled by the MC app-operator.



### cert-operator [2.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v2.0.1)

#### Fixed
- Bump go module major version.



### containerlinux [3139.2.2](https://www.flatcar-linux.org/releases/#release-3139.2.2)

New **Stable** Release **3139.2.2**

_Changes since **Stable 3139.2.1**_

#### Security fixes:

- Linux ([CVE-2022-1734](https://nvd.nist.gov/vuln/detail/CVE-2022-1734), [CVE-2022-28893](https://nvd.nist.gov/vuln/detail/CVE-2022-28893), [CVE-2022-1012](https://nvd.nist.gov/vuln/detail/CVE-2022-1012), [CVE-2022-1729](https://nvd.nist.gov/vuln/detail/CVE-2022-1729))
- Go ([CVE-2022-29526](https://nvd.nist.gov/vuln/detail/CVE-2022-29526))

#### Bug fixes:

- Ensured `/etc/flatcar/update.conf` exists because it happens to be used as flag file for Ansible ([init#71](https://github.com/flatcar-linux/init/pull/71))
- GCP: Fixed shutdown script execution ([coreos-overlay#1912](https://github.com/flatcar-linux/coreos-overlay/pull/1912), [flatcar#743](https://github.com/flatcar-linux/Flatcar/issues/743))


#### Updates:

- Linux ([5.15.43](https://lwn.net/Articles/896231/) (includes [5.15.42](https://lwn.net/Articles/896226), [5.15.41](https://lwn.net/Articles/895645), [5.15.40](https://lwn.net/Articles/895318), [5.15.39](https://lwn.net/Articles/895070), [5.15.38](https://lwn.net/Articles/894357)))
- Go ([1.17.10](https://go.googlesource.com/go/+/refs/tags/go1.17.10))


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

#### Changed
- Add `chart-pull-failed` error to differentiate between issues when pulling chart tarball and other problems.
#### Fixed
- Fix missing `PriorityClass` issue.



### coredns [1.10.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.10.1)

#### Fixed
- Added component label to deployment labels as well



### external-dns [2.14.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.14.0)

#### Added
- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.



### cluster-autoscaler [1.22.2-gs6](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs6)

Not found


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



