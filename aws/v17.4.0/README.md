# :zap: Giant Swarm Release v17.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [11.15.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.14.1)

#### Fixed
- Fix principal ARN for Route53 trusted entity.

#### Changed
- Set default upgrade batch to 10% from 33%
- Set default pause time to 10 minutes
- Remove `imagePullSecrets`


### cluster-autoscaler [1.22.2-gs7](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs7)

#### Changed
Enable balance similar nodepools by default

#### Fixed
Ignore labels to consider nodepools similar groups

#### Added
Support to add extra arguments

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


### kubernetes [1.22.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.11)

#### Bug or Regression
- Bug Fix: Kube-proxy dropped endpointSlice's local endpoints when upgrading from 1.20 to 1.22 ([#110245](https://github.com/kubernetes/kubernetes/pull/110245), [@xh4n3](https://github.com/xh4n3)) [SIG Network]
- EndpointSlices marked for deletion are now ignored during reconciliation. ([#110482](https://github.com/kubernetes/kubernetes/pull/110482), [@aryan9600](https://github.com/aryan9600)) [SIG Apps and Network]
- Fixed a kubelet issue that could result in invalid pod status updates to be sent to the api-server where pods would be reported in a terminal phase but also report a ready condition of true in some cases. ([#110481](https://github.com/kubernetes/kubernetes/pull/110481), [@bobbypage](https://github.com/bobbypage)) [SIG Node and Testing]
- Pods will now post their readiness during termination. ([#110418](https://github.com/kubernetes/kubernetes/pull/110418), [@aojea](https://github.com/aojea)) [SIG Network, Node and Testing]

#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._

### cert-manager [2.15.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.0)

#### Changed
- Upgrade to upstream image [`v1.7.3`](https://github.com/jetstack/cert-manager/releases/tag/v1.7.3) which increases some hard-coded timeouts for certain ACME issuers (ZeroSSL and Sectigo) ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))
- Update kubectl container version to `1.24.2` ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))

### external-dns [2.15.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.0)

#### Added
- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.

#### Changed
- Update test dependencies and py-helm-charts version to [0.7.0](https://github.com/giantswarm/pytest-helm-charts/blob/master/CHANGELOG.md) ([#173](https://github.com/giantswarm/external-dns-app/pull/173))
- Ignore IRSA annotation for service account when using AWS `external` access.


### chart-operator [2.24.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0)

#### Added
- Add Helm release failure reason when it is known, and if there is a currently successfully released version (this includes values schema validation errors as well)
- Split Helm client into private Helm client for giantswarm-namespaced apps and public Helm client for rest of the apps.


#### Changed
- Add `chart-pull-failed` error to differentiate between issues when pulling chart tarball and other problems.
- Always create giantswarm-critical priority class if it does not exist.
- Add initialBootstrapMode flag to allow deploying CNI as managed apps.

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



