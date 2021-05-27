# :zap: Giant Swarm Release v14.2.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [10.3.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.3.0)

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.
- Cancel update loop if source or target release is not found.
- Updated IPAM library to avoid IP conflicts.
#### Added
- Clean up VPC peerings from a cluster VPC when is cluster deleted.
- Clean up Application and Network loadbalancers created by Kubernetes when cluster is deleted.
- Add new flatcar AMIs.
#### Changed
- Fix issues with etcd initial cluster resolving into ELB and causing errors.
- Update `k8scloudconfig` to version `v10.5.0` to support kubernetes `v1.20`.
- Use `networkctl reload` for managing networking to avoid bug in `systemd`.



### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.

#### Changed
- Update Kubernetes dependencies to 1.18 versions.
- Reconcile `CertConfig`s based on their `cert-operator.giantswarm.io/version` label.

#### Removed
- Stop using the `VersionBundle` version.

#### Added
- Add network policy resource.
- Added lookup for nodepool clusters in other namespaces than `default`.



### cluster-operator [3.7.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.7.1)

#### Fixed
- Add `AllowedLabels` to clusterconfigmap resource to prevent unnecessary updates.

#### Added
- Create app CR for per cluster app-operator instance.

#### Removed
- Do not add `VersionBundle` to new `CertConfig` specs (`CertConfig`s are now versioned using a label). **This change requires using `cert-operator`



### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases.
- Emit events when config maps and secrets referenced in App CRs are updated.
- Cache k8sclient, helmclient for later use.
- Apply the namespaceConfig to the desired chart.
- Install apps in CAPI Workload Clusters.
- Apply `compatibleProvider`,`namespace` metadata validation based on the relevant `AppCatalogEntry` CR.
- Add annotations from Helm charts to AppCatalogEntry CRs.
- Enable Vertical Pod Autoscaler.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.
- Restore chart-operator when it had been deleted.
- Use backoff in chart CR watcher to wait until kubeconfig secret exists.

#### Changed
- Updated Helm to v3.5.3.
- Replace status webhook with chart CR status watcher.
- Sort AppCatalogEntry CRs by version and created timestamp.
- Watch cluster namespace for per workload cluster instances of app-operator.



### containerlinux [2765.2.3](https://www.flatcar-linux.org/releases/#release-2765.2.3)


#### Security fixes
* Linux
    - [CVE-2021-28964](https://nvd.nist.gov/vuln/detail/CVE-2021-28964)
    - [CVE-2021-28972](https://nvd.nist.gov/vuln/detail/CVE-2021-28972)
    - [CVE-2021-28971](https://nvd.nist.gov/vuln/detail/CVE-2021-28971)
    - [CVE-2021-28951](https://nvd.nist.gov/vuln/detail/CVE-2021-28951)
    - [CVE-2021-28952](https://nvd.nist.gov/vuln/detail/CVE-2021-28952)
    - [CVE-2021-29266](https://nvd.nist.gov/vuln/detail/CVE-2021-29266)
    - [CVE-2021-28688](https://nvd.nist.gov/vuln/detail/CVE-2021-28688)
    - [CVE-2021-29264](https://nvd.nist.gov/vuln/detail/CVE-2021-29264)
    - [CVE-2021-29649](https://nvd.nist.gov/vuln/detail/CVE-2021-29649)
    - [CVE-2021-29650](https://nvd.nist.gov/vuln/detail/CVE-2021-29650)
    - [CVE-2021-29646](https://nvd.nist.gov/vuln/detail/CVE-2021-29646)
    - [CVE-2021-29647](https://nvd.nist.gov/vuln/detail/CVE-2021-29647)
    - [CVE-2021-29154](https://nvd.nist.gov/vuln/detail/CVE-2021-29154)
    - [CVE-2021-29155](https://nvd.nist.gov/vuln/detail/CVE-2021-29155)
    - [CVE-2021-23133](https://nvd.nist.gov/vuln/detail/CVE-2021-23133)
    - [CVE-2021-27365](https://nvd.nist.gov/vuln/detail/CVE-2021-27365)
    - [CVE-2021-27364](https://nvd.nist.gov/vuln/detail/CVE-2021-27364)
    - [CVE-2021-27363](https://nvd.nist.gov/vuln/detail/CVE-2021-27363)
    - [CVE-2021-28038](https://nvd.nist.gov/vuln/detail/CVE-2021-28038)
    - [CVE-2021-28039](https://nvd.nist.gov/vuln/detail/CVE-2021-28039)
    - [CVE-2021-28375](https://nvd.nist.gov/vuln/detail/CVE-2021-28375)
    - [CVE-2021-28660](https://nvd.nist.gov/vuln/detail/CVE-2021-28660)
    - [CVE-2021-27218](https://nvd.nist.gov/vuln/detail/CVE-2021-27218)
    - [CVE-2021-27219](https://nvd.nist.gov/vuln/detail/CVE-2021-27219)
    - [CVE-2020-25639](https://nvd.nist.gov/vuln/detail/CVE-2020-25639)
    - [CVE-2021-27365](https://nvd.nist.gov/vuln/detail/CVE-2021-27365)
    - [CVE-2021-27364](https://nvd.nist.gov/vuln/detail/CVE-2021-27364)
    - [CVE-2021-27363](https://nvd.nist.gov/vuln/detail/CVE-2021-27363)
    - [CVE-2021-28038](https://nvd.nist.gov/vuln/detail/CVE-2021-28038)
    - [CVE-2021-28039](https://nvd.nist.gov/vuln/detail/CVE-2021-28039)
    - [CVE-2021-26931](https://nvd.nist.gov/vuln/detail/CVE-2021-26931)
    - [CVE-2021-26930](https://nvd.nist.gov/vuln/detail/CVE-2021-26930)
    - [CVE-2021-26932](https://nvd.nist.gov/vuln/detail/CVE-2021-26932)
* Openssl
    - [CVE-2021-23840](https://nvd.nist.gov/vuln/detail/CVE-2021-23840)
    - [CVE-2021-23841](https://nvd.nist.gov/vuln/detail/CVE-2021-23841)
    - [CVE-2020-1971](https://nvd.nist.gov/vuln/detail/CVE-2020-1971)
    - [CVE-2021-23840](https://nvd.nist.gov/vuln/detail/CVE-2021-23840)
    - [CVE-2021-23841](https://nvd.nist.gov/vuln/detail/CVE-2021-23841)
    - [CVE-2021-3449](https://nvd.nist.gov/vuln/detail/CVE-2021-3449)
    - [CVE-2021-3450](https://nvd.nist.gov/vuln/detail/CVE-2021-3450)
* containerd
    - [GHSA-6g2q-w5j3-fwh4](https://github.com/containerd/containerd/security/advisories/GHSA-6g2q-w5j3-fwh4)



#### Bug fixes
- Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))
- GCE: The old interface name ens4v1 which was replaced by eth0 due to a broken udev rule was restored, but now as alternative interface name, and eth0 will stay the primary name for consistency across cloud environments. ([init#38](https://github.com/kinvolk/init/pull/38))
- Include firmware files for all modules shipped in our image ([Issue #359](https://github.com/kinvolk/Flatcar/issues/359), [PR #887](https://github.com/kinvolk/coreos-overlay/pull/887))
- Add explicit path to the binary call in the coreos-metadata unit file ([Issue #360](https://github.com/kinvolk/Flatcar/issues/360))


#### Changes
- The virtio network interfaces got predictable interface names as alternative interface names, and thus these names can also be used to match for a specific interface in case there is more than one and the eth0 and eth1 name assignment is not stable. ([init#38](https://github.com/kinvolk/init/pull/38))

#### Updates
- Linux ([5.10.32](https://lwn.net/Articles/853762/))
- openssl ([1.1.1k](https://mta.openssl.org/pipermail/openssl-announce/2021-March/000197.html))
- open-iscsi ([2.1.4](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.4))
- Containerd ([1.4.4](https://github.com/containerd/containerd/releases/tag/v1.4.4))


### cert-manager [2.7.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.1)

#### Changed
- Set authoritative nameserver to `coredns` when using `dns01` ACME solver. ([#162](https://github.com/giantswarm/cert-manager-app/pull/162))



### cluster-autoscaler [1.19.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.3)

#### Added
- Allow users to set container resources



### chart-operator [2.14.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.14.0)

#### Changed
- Cancel the release resource when the manifest object already exists.
- Cancel the release resource when helm returns an unknown error.



### external-dns [2.3.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.3.1)

#### Changed
- Increase memory limit to 100Mi since we ran into out of memory problems. This will make the app more stable.



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

#### Added
- Added new configuration value `extraArgs`.



### net-exporter [1.10.1](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.1)




