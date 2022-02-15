# :zap: Giant Swarm Release v17.0.0-alpha1 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.22.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.6)

#### Feature
- Kube-apiserver: when merging lists, Server Side Apply now prefers the order of the submitted request instead of the existing persisted object ([#107568](https://github.com/kubernetes/kubernetes/pull/107568), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
#### Bug or Regression
- An inefficient lock in EndpointSlice controller metrics cache has been reworked. Network programming latency may be significantly reduced in certain scenarios, especially in clusters with a large number of Services. ([#107168](https://github.com/kubernetes/kubernetes/pull/107168), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Client-go: fix that paged list calls with ResourceVersionMatch set would fail once paging kicked in. ([#107335](https://github.com/kubernetes/kubernetes/pull/107335), [@fasaxc](https://github.com/fasaxc)) [SIG API Machinery]
- Fix a panic when using invalid output format in kubectl create secret command ([#107346](https://github.com/kubernetes/kubernetes/pull/107346), [@rikatz](https://github.com/rikatz)) [SIG CLI]
- Fix: azuredisk parameter lowercase translation issue ([#107429](https://github.com/kubernetes/kubernetes/pull/107429), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixes a rare race condition handling requests that timeout ([#107459](https://github.com/kubernetes/kubernetes/pull/107459), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Mount-utils: Detect potential stale file handle ([#107039](https://github.com/kubernetes/kubernetes/pull/107039), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage]
#### Other (Cleanup or Flake)
- Updates konnectivity-network-proxy to v0.0.27. This includes a memory leak fix for the network proxy ([#107187](https://github.com/kubernetes/kubernetes/pull/107187), [@rata](https://github.com/rata)) [SIG API Machinery, Auth and Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.39.2 → v0.39.3](https://github.com/google/cadvisor/compare/v0.39.2...v0.39.3)
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.22 → v0.0.27
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.1
#### Removed
_Nothing has changed._



### cert-operator [1.3.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.3.0)

#### Changed
- Use `RenewSelf` instead of `LookupSelf` to prevent expiration of `Vault token`.



### cluster-operator [3.13.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.13.0)

#### Changed
- Removed encryption key creation. Encryption keys will be managed by `encryption-provider-operator`.



### app-operator [5.6.0](https://github.com/giantswarm/app-operator/releases/tag/v5.6.0)

#### Changed
- Get tarball URL for chart CRs from index.yaml for better community app catalog support.
#### Fixed
- Fix error handling in chart CR watcher when chart CRD not installed.



### azure-operator [5.14.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.14.0)

#### Added
- Add support for feature that enables forcing cgroups v1 for Flatcar version `3033.2.0` and above.
#### Changed
- Upgraded to giantswarm/exporterkit v1.0.0
- Upgraded to giantswarm/microendpoint v1.0.0
- Upgraded to giantswarm/microkit v1.0.0
- Upgraded to giantswarm/micrologger v0.6.0
- Upgraded to giantswarm/versionbundle v1.0.0
- Upgraded to spf13/viper v1.10.0



### chart-operator [2.20.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.20.0)

#### Changed
- Update Helm to v3.6.3.
- Use controller-runtime client to remove CAPI dependency.
#### Removed
- Remove unused helm 2 release collector.



### external-dns [2.9.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.9.0)

This release contains some changes to mitigate rate limiting on AWS clusters. Please take note of the defaults
for values `aws.batchChangeInterval`, `aws.zonesCacheDuration`, `externalDNS.interval`
and `externalDNS.minEventSyncInterval`.
If you already specify `--aws-batch-change-interval` or `--aws-zones-cache-duration`, please migrate to the new values `aws.batchChangeInterval` and `aws.zonesCacheDuration`.
#### Added
- Allow to set `--aws-batch-change-interval` through `aws.batchChangeInterval` value. Default `10s`.
- Allow to set `--aws-zones-cache-duration` through `aws.zonesCacheDuration` value. Default `3h`.
#### Changed
- Set default `externalDNS.interval` to `5m`.
- Set default `externalDNS.minEventSyncInterval` to `30s`.
- Allow setting Route53 credentials (`externalDNS.aws_access_key_id` and `externalDNS.aws_secret_access_key`) indepentent from `aws.access` value.
- Allow setting the AWS default region (`aws.region`) indepentent from `aws.access` value.
- Allow to omit the `--domain-filter` flag completely by setting `externalDNS.domainFilterList` to `null`.



### cluster-autoscaler [1.22.2-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs2)

Not found


### azure-scheduled-events [0.6.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.6.0)

#### Added
- Add `priorityClassName: "system-node-critical"` to Daemonset to give higher priority during scheduling.



### containerlinux [3033.2.2](https://www.flatcar-linux.org/releases/#release-3033.2.2)

**Breaking changes**

* CGroupsV2 are enabled by default. Applications might need to be updated if they don't have support. There are
several known issues:
  - Java applications must use JRE >= 15; Please see OpenJDK
[upstream issue](https://bugs.openjdk.java.net/browse/JDK-8230305) for more details.

**Security fixes**

* Linux
    - [CVE-2021-43976](https://nvd.nist.gov/vuln/detail/CVE-2021-43976)
    - [CVE-2022-0330](https://nvd.nist.gov/vuln/detail/CVE-2022-0330)
    - [CVE-2022-22942](https://nvd.nist.gov/vuln/detail/CVE-2022-22942)
    - [CVE-2021-4135](https://nvd.nist.gov/vuln/detail/CVE-2021-4135)
    - [CVE-2021-4155](https://nvd.nist.gov/vuln/detail/CVE-2021-4155)
    - [CVE-2021-28711](https://nvd.nist.gov/vuln/detail/CVE-2021-28711)
    - [CVE-2021-28712](https://nvd.nist.gov/vuln/detail/CVE-2021-28712)
    - [CVE-2021-28713](https://nvd.nist.gov/vuln/detail/CVE-2021-28713)
    - [CVE-2021-28714](https://nvd.nist.gov/vuln/detail/CVE-2021-28714)
    - [CVE-2021-28715](https://nvd.nist.gov/vuln/detail/CVE-2021-28715)
    - [CVE-2021-39685](https://nvd.nist.gov/vuln/detail/CVE-2021-39685)
    - [CVE-2021-44733](https://nvd.nist.gov/vuln/detail/CVE-2021-44733)
    - [CVE-2021-45095](https://nvd.nist.gov/vuln/detail/CVE-2021-45095)
    - [CVE-2022-0185](https://nvd.nist.gov/vuln/detail/CVE-2022-0185)
    - [CVE-2021-4002](https://nvd.nist.gov/vuln/detail/CVE-2021-4002)
    - [CVE-2020-27820](https://nvd.nist.gov/vuln/detail/CVE-2020-27820)
    - [CVE-2021-4001](https://nvd.nist.gov/vuln/detail/CVE-2021-4001)
    - [CVE-2021-43975](https://nvd.nist.gov/vuln/detail/CVE-2021-43975)
    - [CVE-2021-42739](https://nvd.nist.gov/vuln/detail/CVE-2021-42739)
    - [CVE-2021-3760](https://nvd.nist.gov/vuln/detail/CVE-2021-3760)
    - [CVE-2021-3772](https://nvd.nist.gov/vuln/detail/CVE-2021-3772)
    - [CVE-2021-42327](https://nvd.nist.gov/vuln/detail/CVE-2021-42327)
    - [CVE-2021-43056](https://nvd.nist.gov/vuln/detail/CVE-2021-43056)
    - [CVE-2021-43267](https://nvd.nist.gov/vuln/detail/CVE-2021-43267)
    - [CVE-2021-43389](https://nvd.nist.gov/vuln/detail/CVE-2021-43389)
    - CVE-2021-3609
    - [CVE-2021-3653](https://nvd.nist.gov/vuln/detail/CVE-2021-3653)
    - CVE-2021-3655
    - [CVE-2021-3656](https://nvd.nist.gov/vuln/detail/CVE-2021-3656)
    - [CVE-2021-3760](https://nvd.nist.gov/vuln/detail/CVE-2021-3760)
    - [CVE-2021-3772](https://nvd.nist.gov/vuln/detail/CVE-2021-3772)
    - [CVE-2020-26541](https://nvd.nist.gov/vuln/detail/CVE-2020-26541)
    - [CVE-2021-35039](https://nvd.nist.gov/vuln/detail/CVE-2021-35039)
    - [CVE-2021-37576](https://nvd.nist.gov/vuln/detail/CVE-2021-37576)
    - [CVE-2021-22543](https://nvd.nist.gov/vuln/detail/CVE-2021-22543)
    - [CVE-2021-33909](https://nvd.nist.gov/vuln/detail/CVE-2021-33909)
    - [CVE-2021-34556](https://nvd.nist.gov/vuln/detail/CVE-2021-34556)
    - [CVE-2021-35477](https://nvd.nist.gov/vuln/detail/CVE-2021-35477)
    - [CVE-2021-38166](https://nvd.nist.gov/vuln/detail/CVE-2021-38166)
    - [CVE-2021-38205](https://nvd.nist.gov/vuln/detail/CVE-2021-38205)
    - [CVE-2021-42327](https://nvd.nist.gov/vuln/detail/CVE-2021-42327)
    - [CVE-2021-43056](https://nvd.nist.gov/vuln/detail/CVE-2021-43056)
    - [CVE-2021-43267](https://nvd.nist.gov/vuln/detail/CVE-2021-43267)
    - [CVE-2021-43389](https://nvd.nist.gov/vuln/detail/CVE-2021-43389)
* Go
    - [CVE-2021-29923](https://nvd.nist.gov/vuln/detail/CVE-2021-29923)
    - [CVE-2021-39293](https://nvd.nist.gov/vuln/detail/CVE-2021-39293)
    - [CVE-2021-38297](https://nvd.nist.gov/vuln/detail/CVE-2021-38297)
    - [CVE-2021-39293](https://nvd.nist.gov/vuln/detail/CVE-2021-39293)
    - [CVE-2021-44717](https://nvd.nist.gov/vuln/detail/CVE-2021-44717)
    - [CVE-2021-44716](https://nvd.nist.gov/vuln/detail/CVE-2021-44716)
    - [CVE-2021-41771](https://nvd.nist.gov/vuln/detail/CVE-2021-41771)
    - [CVE-2021-41772](https://nvd.nist.gov/vuln/detail/CVE-2021-41772)
* bash
    - [CVE-2019-9924](https://nvd.nist.gov/vuln/detail/CVE-2019-9924)
    - [CVE-2019-18276](https://nvd.nist.gov/vuln/detail/CVE-2019-18276)
* binutils
    - [CVE-2021-3530](https://nvd.nist.gov/vuln/detail/CVE-2021-3530)
    - [CVE-2021-3549](https://nvd.nist.gov/vuln/detail/CVE-2021-3549)
* ca-certificates
    - [CVE-2021-43527](https://nvd.nist.gov/vuln/detail/CVE-2021-43527)
* containerd
    - [CVE-2021-43816](https://nvd.nist.gov/vuln/detail/CVE-2021-43816)
    - [CVE-2021-41103](https://nvd.nist.gov/vuln/detail/CVE-2021-41103)
    - [CVE-2021-41190](https://nvd.nist.gov/vuln/detail/CVE-2021-41190)
* curl
    - [CVE-2021-22945](https://nvd.nist.gov/vuln/detail/CVE-2021-22945)
    - [CVE-2021-22946](https://nvd.nist.gov/vuln/detail/CVE-2021-22946)
    - [CVE-2021-22947](https://nvd.nist.gov/vuln/detail/CVE-2021-22947)
* Docker
    - [CVE-2021-41092](https://nvd.nist.gov/vuln/detail/CVE-2021-41092)
    - [CVE-2021-41089](https://nvd.nist.gov/vuln/detail/CVE-2021-41089)
    - [CVE-2021-41091](https://nvd.nist.gov/vuln/detail/CVE-2021-41091)
    - [CVE-2021-41190](https://nvd.nist.gov/vuln/detail/CVE-2021-41190)
* expat
    - [CVE-2022-23852](hrrp://nvd.nist.gov/vuln/detail/CVE-2022-23852)
    - [CVE-2022-23990](https://nvd.nist.gov/vuln/detail/CVE-2022-23990)
    - [CVE-2021-45960](https://nvd.nist.gov/vuln/detail/CVE-2021-45960)
    - [CVE-2021-46143](https://nvd.nist.gov/vuln/detail/CVE-2021-46143)
    - [CVE-2022-22822](https://nvd.nist.gov/vuln/detail/CVE-2022-22822)
    - [CVE-2022-22823](https://nvd.nist.gov/vuln/detail/CVE-2022-22823)
    - [CVE-2022-22824](https://nvd.nist.gov/vuln/detail/CVE-2022-22824)
    - [CVE-2022-22825](https://nvd.nist.gov/vuln/detail/CVE-2022-22825)
    - [CVE-2022-22826](https://nvd.nist.gov/vuln/detail/CVE-2022-22826)
    - [CVE-2022-22827](https://nvd.nist.gov/vuln/detail/CVE-2022-22827)
* git
    - [CVE-2021-40330](https://nvd.nist.gov/vuln/detail/CVE-2021-40330)
* glibc
    - [CVE-2021-3998](https://nvd.nist.gov/vuln/detail/CVE-2021-3998)
    - [CVE-2021-3999](https://nvd.nist.gov/vuln/detail/CVE-2021-3999)
    - [CVE-2022-23218](https://nvd.nist.gov/vuln/detail/CVE-2022-23218)
    - [CVE-2022-23219](https://nvd.nist.gov/vuln/detail/CVE-2022-23219)
    - [CVE-2021-38604](https://nvd.nist.gov/vuln/detail/CVE-2021-38604)
* gnupg
    - [CVE-2020-25125](https://nvd.nist.gov/vuln/detail/CVE-2020-25125)
* libgcrypt
    - [CVE-2021-40528](https://nvd.nist.gov/vuln/detail/CVE-2021-40528)
* nettle
    - [CVE-2021-20305](https://nvd.nist.gov/vuln/detail/CVE-2021-20305)
    - [CVE-2021-3580](https://nvd.nist.gov/vuln/detail/CVE-2021-3580)
* polkit
    - [CVE-2021-40340](https://nvd.nist.gov/vuln/detail/CVE-2021-40340)
    - [CVE-2021-3560](https://nvd.nist.gov/vuln/detail/CVE-2021-3560)
* sssd
    - [CVE-2021-3621](https://nvd.nist.gov/vuln/detail/CVE-2021-3621)
* util-linux
    - [CVE-2021-37600](https://nvd.nist.gov/vuln/detail/CVE-2021-37600)
* vim
    - [CVE-2021-3770](https://nvd.nist.gov/vuln/detail/CVE-2021-3770)
    - [CVE-2021-3778](https://nvd.nist.gov/vuln/detail/CVE-2021-3778)
    - [CVE-2021-3796](https://nvd.nist.gov/vuln/detail/CVE-2021-3796)
* SDK: bison
    - [CVE-2020-14150](https://nvd.nist.gov/vuln/detail/CVE-2020-14150)
    - [CVE-2020-24240](https://nvd.nist.gov/vuln/detail/CVE-2020-24240)
* SDK: perl
    - [CVE-2020-10878](https://nvd.nist.gov/vuln/detail/CVE-2020-10878)

**Bug fixes**
* SDK: Fixed build error popping up in the new SDK Container because policycoreutils used the wrong ROOT to update the SELinux store ([flatcar-linux/coreos-overlay#1502](https://github.com/flatcar-linux/coreos-overlay/pull/1502))
* Fixed leak of SELinux policy store to the root filesystem top directory due to wrong store path in policycoreutils instead of /var/lib/selinux ([flatcar-linux/Flatcar#596](https://github.com/flatcar-linux/Flatcar/issues/596))
* Ensured that the /run/xtables.lock coordination file exists for modifications of the xtables backend from containers (must be bind-mounted) or the iptables-legacy binaries on the host ([flatcar-linux/init#57](https://github.com/flatcar-linux/init/pull/57))
* dev container: Fix github URL for coreos-overlay and portage-stable to use repos from flatcar-linux org directly instead of relying on redirects from the kinvolk org. This fixes checkouts with emerge-gitclone inside dev-container. ([flatcar-linux/scripts#194](https://github.com/flatcar-linux/scripts/pull/194))
* SDK: Fixed build error popping up in the new SDK Container because policycoreutils used the wrong ROOT to update the SELinux store ([flatcar-linux/coreos-overlay#1502](https://github.com/flatcar-linux/coreos-overlay/pull/1502))
* arm64: the Polkit service does not crash anymore. ([flatcar-linux/Flatcar#156](https://github.com/flatcar-linux/Flatcar/issues/156))
* toolbox: fixed support for multi-layered docker images ([toolbox#5](https://github.com/flatcar-linux/toolbox/pull/5))
* Run emergency.target on ignition/torcx service unit failure in dracut ([bootengine#28](https://github.com/flatcar-linux/bootengine/pull/28))
* Fix vim warnings on missing file, when built with USE=”minimal” ([portage-stable#260](https://github.com/flatcar-linux/portage-stable/pull/260))
* The Torcx profile docker-1.12-no got fixed to reference the current Docker version instead of 19.03 which wasn’t found on the image, causing Torcx to fail to provide Docker ([PR#1456](https://github.com/flatcar-linux/coreos-overlay/pull/1456))
* Use https protocol instead of git for Github URLs ([flatcar-linux/coreos-overlay#1394](https://github.com/flatcar-linux/coreos-overlay/pull/1394))

**Changes**
* Backported elf support for iproute2 ([flatcar-linux/coreos-overlay#1256](https://github.com/flatcar-linux/coreos-overlay/pull/1526))
* Added GPIO support ([coreos-overlay#1236](https://github.com/flatcar-linux/coreos-overlay/pull/1236))
* Enabled SELinux in permissive mode on ARM64 ([coreos-overlay#1245](https://github.com/flatcar-linux/coreos-overlay/pull/1245))
* The iptables command uses the nftables kernel backend instead of the iptables backend, you can also migrate to using the nft tool instead of iptables. Containers with iptables binaries that use the iptables backend will result in mixing both kernel backends which is supported but you have to look up the rules separately (on the host you can use the iptables-legacy and friends).
* Added missing SELinux rule as initial step to resolve Torcx unpacking issue ([coreos-overlay#1426](https://github.com/flatcar-linux/coreos-overlay/pull/1426))

**Updates**
* Linux ([5.10.96](https://lwn.net/Articles/883442))
* Linux Firmware ([20211216](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20211216))
* expat ([2.4.4](https://github.com/libexpat/libexpat/blob/R_2_4_4/expat/Changes))
* ca-certificates ([3.74](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_74.html))
* Docker ([20.10.12](https://docs.docker.com/engine/release-notes/#201012))
* containerd ([1.5.9](https://github.com/containerd/containerd/releases/tag/v1.5.9))
* Go ([1.17.5](https://go.googlesource.com/go/+/refs/tags/go1.17.5))
* systemd ([249.4](https://github.com/systemd/systemd-stable/blob/v249.4/NEWS))
* bash ([5.1_p8](https://lists.gnu.org/archive/html/info-gnu/2020-12/msg00003.html))
* binutils ([2.37](https://sourceware.org/pipermail/binutils/2021-July/117384.html))
* curl ([7.79.1](https://curl.se/changes.html#7_79_1))
* duktape ([2.6.0](https://github.com/svaarala/duktape/blob/master/doc/release-notes-v2-6.rst))
* ebtables ([2.0.11](https://lwn.net/Articles/806179/))
* git ([2.32.0](https://github.com/git/git/blob/master/Documentation/RelNotes/2.32.0.txt))
* gnupg ([2.2.29](https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000461.html))
* iptables ([1.8.7](https://lwn.net/Articles/843069/))
* keyutils ([1.6.1](https://git.kernel.org/pub/scm/linux/kernel/git/dhowells/keyutils.git/tag/?h=v1.6.1))
* ldb ([2.3.0](https://gitlab.com/samba-team/samba/-/tags/ldb-2.3.0))
* libgcrypt ([1.9.4](https://dev.gnupg.org/T5402))
* libmnl ([1.0.4](https://marc.info/?l=netfilter-devel&m=146745072727070&w=2))
* libnftnl ([1.2.0](https://marc.info/?l=netfilter&m=162194376520385&w=2))
* libtirpc ([1.3.2](https://www.spinics.net/lists/linux-nfs/msg84129.html))
* lvm2 ([2.02.188](https://github.com/lvmteam/lvm2/releases/tag/v2_02_188))
* nettle ([3.7.3](https://lists.gnu.org/archive/html/info-gnu/2021-06/msg00002.html))
* nftables ([0.9.9](https://lwn.net/Articles/857369/))
* net-tools ([2.10](https://sourceforge.net/p/net-tools/code/ci/v2.10/tree/))
* openssh ([8.7_p1-r1](https://www.openssh.com/txt/release-8.7))
* open-vm-tools ([11.3.5](https://github.com/vmware/open-vm-tools/releases/tag/stable-11.3.5))
* polkit ([0.119](https://gitlab.freedesktop.org/polkit/polkit/-/blob/0.119/NEWS))
* realmd ([0.17.0](https://gitlab.freedesktop.org/realmd/realmd/-/tags/0.17.0))
* runc ([1.0.3](https://github.com/opencontainers/runc/releases/tag/v1.0.3))
* talloc ([2.3.2](https://gitlab.com/samba-team/samba/-/tags/talloc-2.3.2))
* util-linux ([2.37.2](https://github.com/karelzak/util-linux/blob/v2.37.2/NEWS))
* vim ([8.2.3428](https://github.com/vim/vim/releases/tag/v8.2.3428))
* xenstore ([4.14.2](https://xenproject.org/downloads/xen-project-archives/xen-project-4-14-series/xen-project-4-14-2/))
* SDK: gnuconfig (20210107)
* SDK: google-cloud-sdk ([355.0.0](https://groups.google.com/g/google-cloud-sdk-announce/c/HoJuttxnzNQ))
* SDK: meson (0.57.2)
* SDK: mtools (4.0.35)
* SDK: perl ([5.34.0](https://perldoc.perl.org/perl5340delta))
* SDK: Rust ([1.55.0](https://blog.rust-lang.org/2021/09/09/Rust-1.55.0.html))
* SDK: texinfo ([6.8](https://github.com/debian-tex/texinfo/releases/tag/upstream%2F6.8))

