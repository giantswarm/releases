# :zap: Giant Swarm Release v18.4.0 for AWS :zap:

This release contains changes that address several vulnerabilities and overall improvements. Most important change is extending the IRSA trust policy for external-dns IAM role so it can be used by multiple external-dns in your workload clusters with IRSA enabled.

This release also adds a new component `cilium-prerequisites` that installs `CiliumNetworkPolicy` CRDs towards the easier and downtime-free Cilium migration. This application can also be installed from the catalog.

## Change details


### aws-operator [14.13.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.13.1)

#### Fixed
- Allow to enable ACLs for a S3 buckets.



### containerlinux [3374.2.5](https://www.flatcar-linux.org/releases/#release-3374.2.5)

 _Changes since **Stable 3374.2.4**_
 
 #### Security fixes:
 
 - Linux ([CVE-2022-4129](https://nvd.nist.gov/vuln/detail/CVE-2022-4129), [CVE-2022-4382](https://nvd.nist.gov/vuln/detail/CVE-2022-4382), [CVE-2022-4842](https://nvd.nist.gov/vuln/detail/CVE-2022-4842), [CVE-2023-1073](https://nvd.nist.gov/vuln/detail/CVE-2023-1073), [CVE-2023-1074](https://nvd.nist.gov/vuln/detail/CVE-2023-1074), [CVE-2023-23559](https://nvd.nist.gov/vuln/detail/CVE-2023-23559))
 
 #### Bug fixes:
 
 - Excluded the special Kubernetes network interfaces `nodelocaldns` and `kube-ipvs0` from being managed with systemd-networkd which interfered with the setup ([init#89](https://github.com/flatcar/init/pull/89)).
 
 #### Updates:
 
 - Linux ([5.15.92](https://lwn.net/Articles/922340) (includes [5.15.91](https://lwn.net/Articles/921851), [5.15.90](https://lwn.net/Articles/921029)))



### cilium-prerequisites [0.1.1](https://github.com/giantswarm/cilium-prerequisites/releases/tag/v0.1.1)

#### Fixed 
- Fixed kube-linter.



### observability-bundle [0.4.2](https://github.com/giantswarm/observability-bundle/releases/tag/v0.4.2)

#### Changed
- Upgrade `prometheus-agent-app` to 0.4.1.



### vertical-pod-autoscaler [3.4.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.2)

#### Changed
- Remove circleci job for pushing to shared app collection



### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap



