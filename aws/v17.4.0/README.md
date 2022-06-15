# :zap: Giant Swarm Release v17.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [11.14.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.14.0)

#### Added 
- Added new flatcar 3139.2.2 image release.



### cluster-operator [4.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.3.0)

#### Changed
- Do not update "app-operator.giantswarm.io/version" label on app-operators when their value is 0.0.0 (aka they are reconciled by the management cluster app-operator). This is a use-case for App Bundles for example, because the App CRs they contain should be created in the MC so should be reconciled by the MC app-operator.



### app-operator [5.12.0](https://github.com/giantswarm/app-operator/releases/tag/v5.12.0)

#### Added
- Add `initialBootstrapMode` flag to allow deploying CNI as managed apps.



### aws-cni [1.11.2](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.2)

* Improvement -  [Updated golang to Go 1.18](https://github.com/aws/amazon-vpc-cni-k8s/pull/1991) (@orsenthil)
* Improvement -  [Updated containernetworking/cni version to 0.8.1 to address CVE-2021-20206](https://github.com/aws/amazon-vpc-cni-k8s/pull/1996) (@orsenthil)
* Improvement -  [Updated CNI Plugins to v1.1.1](https://github.com/aws/amazon-vpc-cni-k8s/pull/1997) (@orsenthil)



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



### external-dns [2.14.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.14.0)

#### Added
- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.



### chart-operator [2.24.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0)

#### Changed
- Add `chart-pull-failed` error to differentiate between issues when pulling chart tarball and other problems.
#### Fixed
- Fix missing `PriorityClass` issue.



### node-exporter [1.12.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.12.0)

#### Changed
- Enabled `diskstats` collector.



### vertical-pod-autoscaler [2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

#### Changed
- Use patched docker image tagged `0.10.0-oomfix` for `recommender` and updater (see https://github.com/giantswarm/roadmap/issues/923).



### metrics-server [1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set `kubelet-preferred-address-types` to `Hostname` on `AWS`.


### aws-ebs-csi-driver [1.6.2](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/blob/master/CHANGELOG.md#v162)

#### Notable changes
- Address CVE ALAS-2022-1792


### webhook-exporer [0.1.0](https://github.com/giantswarm/webhook-exporter/releases/tag/v0.1.0)

#### Added
- Initial webhook-exporter.



