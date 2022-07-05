# :zap: Giant Swarm Release v17.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [11.16.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.16.0)

#### Added
- Added new flatcar 3139.2.3 image release.
#### Changed
- Tighten pod and container security contexts for PSS restricted policies.
- Bump `k8scc` to enable `auditd` monitoring for `execve` syscalls.



### cluster-operator [4.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.3.0)

#### Changed
- Do not update "app-operator.giantswarm.io/version" label on app-operators when their value is 0.0.0 (aka they are reconciled by the management cluster app-operator). This is a use-case for App Bundles for example, because the App CRs they contain should be created in the MC so should be reconciled by the MC app-operator.



### app-operator [6.0.1](https://github.com/giantswarm/app-operator/releases/tag/v6.0.1)

#### Added
- Add support for Catalogs that define multiple repository mirrors to be used in case some of them are unreachable.
#### Changed
- Only run `PodMonitor` outside of bootstrap mode.



### aws-cni [1.11.2-nftables](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.2-nftables)

Not found


### containerlinux [3139.2.3](https://www.flatcar-linux.org/releases/#release-3139.2.3)

New **Stable** Release **3139.2.3**

Changes since **Stable 3139.2.2**

#### Security fixes:

- Linux ([CVE-2022-1789](https://nvd.nist.gov/vuln/detail/CVE-2022-1789), [CVE-2022-1852](https://nvd.nist.gov/vuln/detail/CVE-2022-1852), [CVE-2022-1972](https://nvd.nist.gov/vuln/detail/CVE-2022-1972), [CVE-2022-1973](https://nvd.nist.gov/vuln/detail/CVE-2022-1973), [CVE-2022-2078](https://nvd.nist.gov/vuln/detail/CVE-2022-2078), [CVE-2022-32250](https://nvd.nist.gov/vuln/detail/CVE-2022-32250), [CVE-2022-32981](https://nvd.nist.gov/vuln/detail/CVE-2022-32981))
- libpcre2 ([CVE-2022-1586](https://nvd.nist.gov/vuln/detail/CVE-2022-1586), [CVE-2022-1587](https://nvd.nist.gov/vuln/detail/CVE-2022-1587))

#### Updates:

- Linux ([5.15.48](https://lwn.net/Articles/898124) (includes [5.15.47](https://lwn.net/Articles/897904), [5.15.46](https://lwn.net/Articles/897377), [5.15.45](https://lwn.net/Articles/897167), [5.15.44](https://lwn.net/Articles/896647)))
- ca-certificates ([3.79](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_79.html))
- libpcre2 ([10.40](https://github.com/PCRE2Project/pcre2/blob/pcre2-10.40/NEWS))


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



### calico [3.21.5](https://github.com/projectcalico/calico/releases/tag/v3.21.5)

Not found


### external-dns [2.15.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.0)

#### Changed
- Update test dependencies and py-helm-charts version to [0.7.0](https://github.com/giantswarm/pytest-helm-charts/blob/master/CHANGELOG.md) ([#173](https://github.com/giantswarm/external-dns-app/pull/173))
- Ignore IRSA annotation for service account when using AWS `external` access.



### chart-operator [2.24.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.1)

#### Changed
- Update `helmclient` to v4.10.1.



### node-exporter [1.13.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.13.0)

#### Changed
- Disable boot partition from the `filesystem` exporter.



### vertical-pod-autoscaler [2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

#### Changed
- Use patched docker image tagged `0.10.0-oomfix` for `recommender` and updater (see https://github.com/giantswarm/roadmap/issues/923).



### metrics-server [1.7.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.7.0)

#### Changed
- Set `kubelet-preferred-address-types` to `Hostname` on `AWS`.



### cert-manager [2.15.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.0)

#### Changed
- Upgrade to upstream image [`v1.7.3`](https://github.com/jetstack/cert-manager/releases/tag/v1.7.3) which increases some hard-coded timeouts for certain ACME issuers (ZeroSSL and Sectigo) ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))
- Update kubectl container version to `1.24.2` ([#243](https://github.com/giantswarm/cert-manager-app/pull/243))



### aws-ebs-csi-driver [2.14.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.14.0)

#### Changed
- Remove `imagePullSecrets` from values.yaml
- Bump aws-ebs-csi-driver version to `v1.6.2`.



### cluster-autoscaler [1.22.2-gs7](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs7)

Not found


### coredns [1.10.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.10.0)

#### Added
- Add `app.kubernetes.io/component` on deployments so that management-cluster-admission controller does not complain.



### kube-state-metrics [1.11.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.11.0)

#### Add
- Allow `application.giantswarm.io/team` label.



### etcd-kubernetes-resources-count-exporter [0.5.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v0.5.0)

Added new default app to expose count of kubernetes objects in the cluster.



