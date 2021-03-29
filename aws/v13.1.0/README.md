# :zap: Giant Swarm Release v13.1.0 for AWS :zap:

This release provides security and bug fixes for various components.

## Change details


### aws-operator [9.3.9](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.9)

#### Fixed
- Added CNI CIDR to internal ELB Security Group
- Added new Flatcar AMI identifiers
- Added China Flatcar AMI identifiers

### cluster-operator [3.6.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.6.0)

#### Fixed
- Fix cluster status computation to correctly display rollbacks, version changes and multiple updates.
#### Added
- Add vertical pod autoscaler support
- Add `appversionlabel` resource to update version labels for optional app CRs
- Check existence of chart tarball for `release` CR `apps` in catalog
- Add unit tests for cluster status computation



### app-operator [3.2.0](https://github.com/giantswarm/app-operator/releases/tag/v3.2.0)

#### Added
- Add printer columns for Version, Last Deployed and Status to chart CRD in
tenant clusters.
- Use validation logic from the app library.
- Include restrictions data from app metadata files in appcatalogentry CRs.
- Include `apiVersion`, `restrictions.compatibleProviders` in appcatalogentry CRs.
#### Changed
- Using values service from the app library.
- Updated Helm to v3.4.2.
- Enable mutating and validating webhooks in app-admission-controller for tenant app CRs.
- Limit the number of AppCatalogEntry per app.
- Delete legacy finalizers on app CRs.
- Reconciling appCatalog CRs only if pod is unique.
#### Fixed
- Reuse clients in clients resource when app CR uses inCluster.
- Updating status as cordoned if app CR has cordoned annotation.



### aws-cni [1.7.8](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.7.8)

* Bug - [Rearrange Pod deletion workflow](https://github.com/aws/amazon-vpc-cni-k8s/pull/1315) (#1315, @SaranBalaji90)
* Improvement - [Replace DescribeNetworkInterfaces with paginated version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1333) (#1333, [@haouc](https://github.com/haouc))



### containerlinux [2605.12.0](https://www.flatcar-linux.org/releases/#release-2605.12.0)

**Security fixes**

* linux
  - [CVE-2020-14390](https://nvd.nist.gov/vuln/detail/CVE-2020-14390)
  - [CVE-2020-25284](https://nvd.nist.gov/vuln/detail/CVE-2020-25284)
  - [CVE-2020-25645](https://nvd.nist.gov/vuln/detail/CVE-2020-25645)
  - [CVE-2020-25643](https://nvd.nist.gov/vuln/detail/CVE-2020-25643)
  - [CVE-2020-25211](https://nvd.nist.gov/vuln/detail/CVE-2020-25211)
  - [CVE-2020-27673](https://nvd.nist.gov/vuln/detail/CVE-2020-27673)
  - [CVE-2020-27675](https://nvd.nist.gov/vuln/detail/CVE-2020-27675)
  - [CVE-2020-28941](https://nvd.nist.gov/vuln/detail/CVE-2020-28941)
  - [CVE-2020-4788](https://nvd.nist.gov/vuln/detail/CVE-2020-4788)
  - [CVE-2020-25669](https://nvd.nist.gov/vuln/detail/CVE-2020-25669)
  - [CVE-2020-14351](https://nvd.nist.gov/vuln/detail/CVE-2020-14351)
  - [CVE-2020-29661](https://nvd.nist.gov/vuln/detail/CVE-2020-29661)
  - [CVE-2020-29660](https://nvd.nist.gov/vuln/detail/CVE-2020-29660)
  - [CVE-2020-27830](https://nvd.nist.gov/vuln/detail/CVE-2020-27830)
  - [CVE-2020-28588](https://nvd.nist.gov/vuln/detail/CVE-2020-28588)
  - [CVE-2020-27815](https://www.openwall.com/lists/oss-security/2020/11/30/5)
  - [CVE-2020-29568](https://nvd.nist.gov/vuln/detail/CVE-2020-29568)
  - [CVE-2020-29569](https://nvd.nist.gov/vuln/detail/CVE-2020-29569)
  - [CVE-2020-28374](https://nvd.nist.gov/vuln/detail/CVE-2020-28374)
  - [CVE-2020-36158](https://nvd.nist.gov/vuln/detail/CVE-2020-36158)
* go
  - [CVE-2021-3114](https://github.com/golang/go/issues/43786)
* sudo
  - [CVE-2021-3156](https://nvd.nist.gov/vuln/detail/CVE-2021-3156)
  - [CVE-2021-23239](https://nvd.nist.gov/vuln/detail/CVE-2021-23239)
* containerd
  - [CVE-2020-15257](https://nvd.nist.gov/vuln/detail/CVE-2020-15257)
* glibc
  - [CVE-2019-9169](https://nvd.nist.gov/vuln/detail/CVE-2019-9169)
  - [CVE-2019-6488](https://nvd.nist.gov/vuln/detail/CVE-2019-6488)
  - [CVE-2019-7309](https://nvd.nist.gov/vuln/detail/CVE-2019-7309)
  - [CVE-2020-10029](https://nvd.nist.gov/vuln/detail/CVE-2020-10029)
  - [CVE-2020-1751](https://nvd.nist.gov/vuln/detail/CVE-2020-1751)
  - [CVE-2020-6096](https://nvd.nist.gov/vuln/detail/CVE-2020-6096)
  - [CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796)
* glib
  - [CVE-2019-12450](https://nvd.nist.gov/vuln/detail/CVE-2019-12450)
* open-iscsi
  - [CVE-2017-17840](https://nvd.nist.gov/vuln/detail/CVE-2017-17840)
* samba
  - [CVE-2019-10197](https://nvd.nist.gov/vuln/detail/CVE-2019-10197)
  - [CVE-2020-10704](https://nvd.nist.gov/vuln/detail/CVE-2020-10704)
  - [CVE-2020-10745](https://nvd.nist.gov/vuln/detail/CVE-2020-10745)
  - [CVE-2019-3880](https://nvd.nist.gov/vuln/detail/CVE-2019-3880)
  - [CVE-2019-10218](https://nvd.nist.gov/vuln/detail/CVE-2019-10218)
* shadow
  - [CVE-2019-19882](https://nvd.nist.gov/vuln/detail/CVE-2019-19882)
* sssd
  - [CVE-2018-16883](https://nvd.nist.gov/vuln/detail/CVE-2018-16883)
  - [CVE-2019-3811](https://nvd.nist.gov/vuln/detail/CVE-2019-3811)
  - [CVE-2018-16838](https://nvd.nist.gov/vuln/detail/CVE-2018-16838)
* trousers
  - [CVE-2020-24330](https://nvd.nist.gov/vuln/detail/CVE-2020-24330)
  - [CVE-2020-24331](https://nvd.nist.gov/vuln/detail/CVE-2020-24331)
* cifs-utils
  - [CVE-2020-14342](https://nvd.nist.gov/vuln/detail/CVE-2020-14342)
* ntp
  - [CVE-2020-11868](https://nvd.nist.gov/vuln/detail/CVE-2020-11868)
  - [CVE-2020-13817](https://nvd.nist.gov/vuln/detail/CVE-2020-13817)
  - [CVE-2018-8956](https://nvd.nist.gov/vuln/detail/CVE-2018-8956)
  - [CVE-2020-15025](https://nvd.nist.gov/vuln/detail/CVE-2020-15025)
* bzip2
  - [CVE-2019-12900](https://nvd.nist.gov/vuln/detail/CVE-2019-12900)
* c-ares
  - [CVE-2017-1000381](https://nvd.nist.gov/vuln/detail/CVE-2017-1000381)
* file
  - [CVE-2019-18218](https://nvd.nist.gov/vuln/detail/CVE-2019-18218)
* json-c
  - [CVE-2020-12762](https://nvd.nist.gov/vuln/detail/CVE-2020-12762)
* jq
  - [CVE-2015-8863](https://nvd.nist.gov/vuln/detail/CVE-2015-8863)
  - [CVE-2016-4074](https://nvd.nist.gov/vuln/detail/CVE-2016-4074)
* libuv
  - [CVE-2020-8252](https://nvd.nist.gov/vuln/detail/CVE-2020-8252)
* libxml2
  - [CVE-2019-20388](https://nvd.nist.gov/vuln/detail/CVE-2019-20388)
  - [CVE-2020-7595](https://nvd.nist.gov/vuln/detail/CVE-2020-7595)
* re2c
  - [CVE-2020-11958](https://nvd.nist.gov/vuln/detail/CVE-2020-11958)
* tar
  - [CVE-2019-9923](https://nvd.nist.gov/vuln/detail/CVE-2019-9923)
* sqlite
  - [CVE-2020-11656](https://nvd.nist.gov/vuln/detail/CVE-2020-11656)
  - [CVE-2020-9327](https://nvd.nist.gov/vuln/detail/CVE-2020-9327)
  - [CVE-2020-11655](https://nvd.nist.gov/vuln/detail/CVE-2020-11655)
  - [CVE-2020-13630](https://nvd.nist.gov/vuln/detail/CVE-2020-13630)
  - [CVE-2020-13435](https://nvd.nist.gov/vuln/detail/CVE-2020-13435)
  - [CVE-2020-13434](https://nvd.nist.gov/vuln/detail/CVE-2020-13434)
  - [CVE-2020-13631](https://nvd.nist.gov/vuln/detail/CVE-2020-13631)
  - [CVE-2020-13632](https://nvd.nist.gov/vuln/detail/CVE-2020-13632)
  - [CVE-2020-15358](https://nvd.nist.gov/vuln/detail/CVE-2020-15358)
* tcpdump and pcap
  - [CVE-2018-10103](https://nvd.nist.gov/vuln/detail/CVE-2018-10103)
  - [CVE-2018-10105](https://nvd.nist.gov/vuln/detail/CVE-2018-10105)
  - [CVE-2019-15163](https://nvd.nist.gov/vuln/detail/CVE-2019-15163)
  - [CVE-2018-14461](https://nvd.nist.gov/vuln/detail/CVE-2018-14461)
  - [CVE-2018-14462](https://nvd.nist.gov/vuln/detail/CVE-2018-14462)
  - [CVE-2018-14463](https://nvd.nist.gov/vuln/detail/CVE-2018-14463)
  - [CVE-2018-14464](https://nvd.nist.gov/vuln/detail/CVE-2018-14464)
  - [CVE-2018-14465](https://nvd.nist.gov/vuln/detail/CVE-2018-14465)
  - [CVE-2018-14466](https://nvd.nist.gov/vuln/detail/CVE-2018-14466)
  - [CVE-2018-14467](https://nvd.nist.gov/vuln/detail/CVE-2018-14467)
  - [CVE-2018-14468](https://nvd.nist.gov/vuln/detail/CVE-2018-14468)
  - [CVE-2018-14469](https://nvd.nist.gov/vuln/detail/CVE-2018-14469)
  - [CVE-2018-14470](https://nvd.nist.gov/vuln/detail/CVE-2018-14470)
  - [CVE-2018-14880](https://nvd.nist.gov/vuln/detail/CVE-2018-14880)
  - [CVE-2018-14881](https://nvd.nist.gov/vuln/detail/CVE-2018-14881)
  - [CVE-2018-14882](https://nvd.nist.gov/vuln/detail/CVE-2018-14882)
  - [CVE-2018-16227](https://nvd.nist.gov/vuln/detail/CVE-2018-16227)
  - [CVE-2018-16228](https://nvd.nist.gov/vuln/detail/CVE-2018-16228)
  - [CVE-2018-16229](https://nvd.nist.gov/vuln/detail/CVE-2018-16229)
  - [CVE-2018-16230](https://nvd.nist.gov/vuln/detail/CVE-2018-16230)
  - [CVE-2018-16300](https://nvd.nist.gov/vuln/detail/CVE-2018-16300)
  - [CVE-2018-16451](https://nvd.nist.gov/vuln/detail/CVE-2018-16451)
  - [CVE-2018-16452](https://nvd.nist.gov/vuln/detail/CVE-2018-16452)
  - [CVE-2019-15166](https://nvd.nist.gov/vuln/detail/CVE-2019-15166)
  - [CVE-2018-14879](https://nvd.nist.gov/vuln/detail/CVE-2018-14879)
  - [CVE-2017-16808](https://nvd.nist.gov/vuln/detail/CVE-2017-16808)
  - [CVE-2018-19519](https://nvd.nist.gov/vuln/detail/CVE-2018-19519)
  - [CVE-2019-15161](https://nvd.nist.gov/vuln/detail/CVE-2019-15161)
  - [CVE-2019-15165](https://nvd.nist.gov/vuln/detail/CVE-2019-15165)
  - [CVE-2019-15164](https://nvd.nist.gov/vuln/detail/CVE-2019-15164)
  - [CVE-2019-1010220](https://nvd.nist.gov/vuln/detail/CVE-2019-1010220)
* libbsd
  - [CVE-2019-20367](https://nvd.nist.gov/vuln/detail/CVE-2019-20367)
* rsync and zlib
  - [CVE-2016-9840](https://nvd.nist.gov/vuln/detail/CVE-2016-9840)
  - [CVE-2016-9841](https://nvd.nist.gov/vuln/detail/CVE-2016-9841)
  - [CVE-2016-9842](https://nvd.nist.gov/vuln/detail/CVE-2016-9842)
  - [CVE-2016-9843](https://nvd.nist.gov/vuln/detail/CVE-2016-9843)



**Bug fixes**


* Enabled missing systemd services ([#191](https://github.com/flatcar-linux/Flatcar/issues/191), [PR #612](https://github.com/flatcar-linux/coreos-overlay/pull/612))
* Fixed Docker torcx image unpacking error on machines with less than ~600 MB total RAM ([#32](https://github.com/flatcar-linux/Flatcar/issues/32))
* Solved adcli Kerberos Active Directory incompatibility ([#194](https://github.com/flatcar-linux/Flatcar/issues/194))
* Fixed the makefile path when building kernel modules with the developer container ([#195](https://github.com/flatcar-linux/Flatcar/issues/195))
* Removed the /etc/portage/savedconfig/ folder that contained a dump of the firmware config [flatcar-linux/coreos-overlay#613](https://github.com/flatcar-linux/coreos-overlay/pull/613)
* Ensured that the /etc/coreos to /etc/flatcar symlink always exists, relevant for the Container Linux Config transpiler (ct) when specifying directives for update: or locksmith: while also reformatting the rootfs ([baselayout PR#7](https://github.com/flatcar-linux/baselayout/pull/7))
* network: Restore KeepConfiguration=dhcp-on-stop ([kinvolk/init#30](https://github.com/kinvolk/init/pull/30))
* Added systemd-tmpfiles directives for /opt and /opt/bin to ensure that the folders have correct permissions even when /opt/ was once created by containerd ([Flatcar#279](https://github.com/kinvolk/Flatcar/issues/279))
* Make the automatic filesystem resizing more robust against a race and add more logging ([kinvolk/init#31](https://github.com/kinvolk/init/pull/31))
* Allow inactive network interfaces to be bound to a bonding interface, by encoding additional configuration for systemd-networkd-wait-online ([afterburn PR #10](https://github.com/flatcar-linux/afterburn/pull/10))
* Do not configure ccache in Jenkins ([scripts PR #100](https://github.com/flatcar-linux/scripts/pull/100))
* Azure: Exclude bonded SR-IOV network interfaces with newer drivers from networkd (in addition to the old drivers) to prevent them being configured instead of just the bond interface ([init PR#29](https://github.com/flatcar-linux/init/pull/29), [bootengine PR#19](https://github.com/flatcar-linux/bootengine/pull/19))
* The sysctl net.ipv4.conf.*.rp_filter is set to 0 for the Cilium CNI plugin to work ([Flatcar#181](https://github.com/kinvolk/Flatcar/issues/181))
* Package downloads in the developer container now use the correct URL again ([Flatcar#298](https://github.com/kinvolk/Flatcar/issues/298))
* networkd: avoid managing MAC addresses for veth devices ([kinvolk/init#33](https://github.com/kinvolk/init/pull/33))
* `/etc/iscsi/initiatorname.iscsi` is generated by the iscsi-init service ([#321](https://github.com/kinvolk/Flatcar/issues/321))
* Prevent iscsiadm buffer overflow ([#318](https://github.com/kinvolk/Flatcar/issues/318))

**Changes**

* GCE: Improved oslogin support and added shell aliases to run a Python Docker image ([PR #592](https://github.com/flatcar-linux/coreos-overlay/pull/592))
* Update-engine now detects rollbacks and reports them as errors to the update server ([PR#6](https://github.com/flatcar-linux/update_engine/pull/6))
* The zstd tools were added (version 1.4.4)
* The kernel config CONFIG_PSI was set to support [Pressure Stall Information](https://www.kernel.org/doc/html/latest/accounting/psi.html), more information also under https://facebookmicrosites.github.io/psi/docs/overview ([Flatcar#162](https://github.com/flatcar-linux/Flatcar/issues/162))
* The kernel config CONFIG_BPF_JIT_ALWAYS_ON was set to use the BPF just-in-time compiler by default for faster execution
* The kernel config CONFIG_POWER_SUPPLY was set
* The kernel configs CONFIG_OVERLAY_FS_METACOPY and CONFIG_OVERLAY_FS_REDIRECT_DIR were set. With the first overlayfs will only copy up metadata when a metadata-specific operation like chown/chmod is performed. The full file will be copied up later when the file is opened for write operations. With the second, which is equivalent to setting “redirect_dir=on” in the kernel command-line, overlayfs will copy up the directory first before the actual content ([Flatcar#170](https://github.com/kinvolk/Flatcar/issues/170)).
* Remove unnecessary kernel module nf-conntrack-ipv4 ([overlay PR#649](https://github.com/flatcar-linux/coreos-overlay/pull/649))
* Compress kernel modules with xz ([overlay PR#628](https://github.com/flatcar-linux/coreos-overlay/pull/628))
* Add containerd-runc-shim-v* binaries required by kubelet custom CRI endpoints ([overlay PR#623](https://github.com/flatcar-linux/coreos-overlay/pull/623))
* Equinix Metal (Packet): Exclude unused network interfaces from networkd, disregard the state of the bonded interfaces for the network-online.target and only require the bond interface itself to have at least one active link instead of routable which requires both links to be active ([afterburn PR#10](https://github.com/flatcar-linux/afterburn/pull/10))
* QEMU: Use flatcar.autologin kernel command line parameter for auto login on the console ([Flatcar #71](https://github.com/flatcar-linux/Flatcar/issues/71))
* The sysctl default config file is now applied under the prefix 60 which allows for custom sysctl config files to take effect when they start with a prefix of 70, 80, or 90 ([baselayout#13](https://github.com/kinvolk/baselayout/pull/13))
* Containerd CRI plugin got enabled by default, only the containerd socket path needs to be specified as kubelet parameter for Kubernetes 1.20 to use containerd instead of Docker ([Flatcar#283](https://github.com/kinvolk/Flatcar/issues/283))
* For users with a custom update server a machine alias setting in update-engine allows to give human-friendly names to client instances ([update-engine#8](https://github.com/kinvolk/update_engine/pull/8))
* Revert to building docker and containerd with go1.13 instead of go1.15. This reduces the SIGURG log spam ([Issue #315](https://github.com/kinvolk/Flatcar/issues/315) [PR #774](https://github.com/kinvolk/coreos-overlay/pull/774))
* The containerd socket is now available in the default location (`/run/containerd/containerd.sock`) and also as a symlink in the previous location (`/run/docker/libcontainerd/docker-containerd.sock`) ([#771](https://github.com/kinvolk/coreos-overlay/pull/771))
* With the iscsi update, the service unit has changed from iscsid to iscsi ([#791](https://github.com/kinvolk/coreos-overlay/pull/791))
* AWS Pro: include scripts to facilitate setup of EKS workers ([#794](https://github.com/kinvolk/coreos-overlay/pull/794)).
* Missed from earlier notes: with the previous open-iscsi update to 2.1.2, the service unit name changed from iscsid to iscsi ([#682](https://github.com/kinvolk/coreos-overlay/pull/682))

**Updates**

* linux ([5.4.92](https://lwn.net/Articles/843687/))
* linux firmware ([20200918](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20200918v))
* systemd ([246.6](https://github.com/systemd/systemd-stable/releases/tag/v246.6))
* open-iscsi ([2.1.3](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.3))
* go ([1.15.7](https://go.googlesource.com/go/+/refs/tags/go1.15.7))
* sudo ([1.9.5p2](https://github.com/sudo-project/sudo/releases/tag/SUDO_1_9_5p2))
* adcli ([0.9.0](https://cgit.freedesktop.org/realmd/adcli/tree/NEWS?h=0.9.0))
* GCE: oslogin ([20200910.00](https://github.com/GoogleCloudPlatform/guest-oslogin/releases/tag/20200910.00))
* glibc ([2.32](https://lwn.net/Articles/828210/))
* Docker ([19.03.14](https://github.com/docker/docker-ce/releases/tag/v19.03.14))
* containerd ([1.4.3](https://github.com/containerd/containerd/releases/tag/v1.4.3))
* tini ([0.18](https://github.com/krallin/tini/releases/tag/v0.18.0))
* libseccomp ([2.5.0](https://github.com/seccomp/libseccomp/releases/tag/v2.5.0))
* audit ([2.8.5](https://github.com/linux-audit/audit-userspace/releases/tag/v2.8.5))
* bzip2 ([1.0.8](https://sourceware.org/git/?p=bzip2.git;a=blob;f=CHANGES;h=30afead2586b6d64f50988a41d394a0131b38949;hb=HEAD#l342))
* c-ares ([1.61.1](https://github.com/c-ares/c-ares/releases/tag/cares-1_16_1))
* cryptsetup ([2.3.2](https://gitlab.com/cryptsetup/cryptsetup/-/tags/v2.3.2))
* cifs-utils (6.11)
* dbus-glib (0.110)
* dracut ([050](https://github.com/dracutdevs/dracut/releases/tag/050))
* elfutils (0.178)
* glib (2.64.5)
* json-c ([0.15](https://github.com/json-c/json-c/releases/tag/json-c-0.15-20200726))
* jq ([1.6](https://github.com/stedolan/jq/releases/tag/jq-1.6))
* libuv ([1.39.0](https://github.com/libuv/libuv/releases/tag/v1.39.0))
* libxml2 ([2.9.10](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.9.10))
* ntp (4.2.8_p15)
* samba (4.11.13)
* shadow (4.8)
* sssd (2.3.1)
* strace (5.9)
* talloc (2.3.1)
* tar ([1.32](https://git.savannah.gnu.org/cgit/tar.git/tag/?h=release_1_32))
* tdb (1.4.3)
* tevent (0.10.2)
* SDK/developer container: GCC (9.3.0), binutils (2.35), gdb (9.2)
* Rust ([1.46.0](https://blog.rust-lang.org/2020/08/27/Rust-1.46.0.html)) (only in SDK)
* file ([5.39](https://github.com/file/file/tree/FILE5_39)) (only in SDK)
* gdbus-codegen ([2.64.5](https://gitlab.gnome.org/GNOME/glib/-/tags/2.64.5)) (only in SDK)
* meson ([0.55.3](https://github.com/mesonbuild/meson/releases/tag/0.55.3)) (only in SDK)
* re2c ([2.0.3](https://re2c.org/releases/release_notes.html#release-2-0-3)) (only in SDK)
* VMware: open-vm-tools (11.2.0)


### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.4.13...v3.4.14) and [v3.4 upgrade guide](https://etcd.io/docs/latest/upgrades/upgrade_3_4/) for any breaking changes.
#### Package `clientv3`
- Fix [auth token invalid after watch reconnects](https://github.com/etcd-io/etcd/pull/12264). Get AuthToken automatically when clientConn is ready.
#### etcd server
- [Fix server panic](https://github.com/etcd-io/etcd/pull/12288) when force-new-cluster flag is enabled in a cluster which had learner node.
#### Package `netutil`
- Remove [`netutil.DropPort/RecoverPort/SetLatency/RemoveLatency`](https://github.com/etcd-io/etcd/pull/12491).
  - These are not used anymore. They were only used for older versions of functional testing.
  - Removed to adhere to best security practices, minimize arbitrary shell invocation.
#### `tools/etcd-dump-metrics`
- Implement [input validation to prevent arbitrary shell invocation](https://github.com/etcd-io/etcd/pull/12491).
#### Go
- Compile with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).



### kubernetes [1.18.17](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.17)

### Feature

- Add a new flag to set priority for the kubelet on Windows nodes so that workloads cannot overwhelm the node there by disrupting kubelet process. ([#96158](https://github.com/kubernetes/kubernetes/pull/96158), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node]


#### Failing Test
- Fix handing special characters in the volume path on Windows ([#99138](https://github.com/kubernetes/kubernetes/pull/99138), [@yujuhong](https://github.com/yujuhong)) [SIG Storage]

#### Bug or Regression
- Count pod overhead against an entity's ResourceQuota ([#99600](https://github.com/kubernetes/kubernetes/pull/99600), [@gjkim42](https://github.com/gjkim42)) [SIG API Machinery and Node]
- EndpointSlice controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100146](https://github.com/kubernetes/kubernetes/pull/100146), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#98943](https://github.com/kubernetes/kubernetes/pull/98943), [@CKchen0726](https://github.com/CKchen0726)) [SIG Apps and Node]
- Kubelet now cleans up orphaned volume directories automatically ([#95301](https://github.com/kubernetes/kubernetes/pull/95301), [@lorenz](https://github.com/lorenz)) [SIG Node and Storage]
- Resolves spurious `Failed to list *v1.Secret` or `Failed to list *v1.ConfigMap` messages in kubelet logs. ([#99538](https://github.com/kubernetes/kubernetes/pull/99538), [@liggitt](https://github.com/liggitt)) [SIG Auth and Node]
- We will no longer automatically delete all data when a failure is detected during creation of the volume data file on a CSI volume. Now we will only remove the data file and volume path. ([#96021](https://github.com/kubernetes/kubernetes/pull/96021), [@huffmanca](https://github.com/huffmanca)) [SIG Storage]
- Avoid marking node as Ready until node has synced with API servers at least once ([#99034](https://github.com/kubernetes/kubernetes/pull/99034), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Cleanup subnet in frontend IP configs to prevent huge subnet request bodies in some scenarios. ([#98290](https://github.com/kubernetes/kubernetes/pull/98290), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix CSI-migrated inline EBS volumes failing to mount if their volumeID is prefixed by aws:// ([#96821](https://github.com/kubernetes/kubernetes/pull/96821), [@wongma7](https://github.com/wongma7)) [SIG Storage]
- Fix azure file migration issue ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx)) [SIG Auth, Cloud Provider and Storage]
- Fix to recover CSI volumes from certain dangling attachments ([#96617](https://github.com/kubernetes/kubernetes/pull/96617), [@yuga711](https://github.com/yuga711)) [SIG Apps and Storage]
- Fixed a bug where aggregator_unavailable_apiservice metrics were reported for deleted apiservices. ([#96421](https://github.com/kubernetes/kubernetes/pull/96421), [@dgrisonnet](https://github.com/dgrisonnet)) [SIG API Machinery and Instrumentation]
- Fixed provisioning of Cinder volumes migrated to CSI when StorageClass with AllowedTopologies was used. ([#98311](https://github.com/kubernetes/kubernetes/pull/98311), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixes a panic in the disruption budget controller for PDB objects with invalid selectors ([#98777](https://github.com/kubernetes/kubernetes/pull/98777), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Apps]
- Kubeadm: get k8s CI version markers from k8s infra bucket ([#98836](https://github.com/kubernetes/kubernetes/pull/98836), [@hasheddan](https://github.com/hasheddan)) [SIG Cluster Lifecycle and Release]
- Kubelet should ignore cgroup driver check on Windows node. ([#98384](https://github.com/kubernetes/kubernetes/pull/98384), [@pacoxu](https://github.com/pacoxu)) [SIG Node]
- TerminationGracePeriodSeconds from pod spec is respected for the mirror pod
  Static pods will be deleted gracefully ([#99035](https://github.com/kubernetes/kubernetes/pull/99035), [@ehashman](https://github.com/ehashman)) [SIG Node and Testing]
- Truncates a message if it hits the NoteLengthLimit when the scheduler records an event for the pod that indicates the pod has failed to schedule. ([#98715](https://github.com/kubernetes/kubernetes/pull/98715), [@carlory](https://github.com/carlory)) [SIG Scheduling]
- Warning about using a deprecated volume plugin is logged only once. ([#96751](https://github.com/kubernetes/kubernetes/pull/96751), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fix Azure file share not deleted issue when the namespace is deleted ([#97417](https://github.com/kubernetes/kubernetes/pull/97417), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix counting error in service/nodeport/loadbalancer quota check ([#97829](https://github.com/kubernetes/kubernetes/pull/97829), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery and Network]
- Fix: azure file latency issue for metadata-heavy workloads ([#97082](https://github.com/kubernetes/kubernetes/pull/97082), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed bug in CPUManager with race on container map access ([#97427](https://github.com/kubernetes/kubernetes/pull/97427), [@klueska](https://github.com/klueska)) [SIG Node]
- GCE Internal LoadBalancer sync loop will now release the ILB IP address upon sync failure. An error in ILB forwarding rule creation will no longer leak IP addresses. ([#97740](https://github.com/kubernetes/kubernetes/pull/97740), [@prameshj](https://github.com/prameshj)) [SIG Cloud Provider and Network]
- Kubeadm: avoid detection of the container runtime for commands that do not need it ([#97849](https://github.com/kubernetes/kubernetes/pull/97849), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Cordoned nodes are now deregistered from AWS target groups. ([#85920](https://github.com/kubernetes/kubernetes/pull/85920), [@hoelzro](https://github.com/hoelzro)) [SIG Cloud Provider]
- Fixed FibreChannel volume plugin corrupting filesystems on detach of multipath volumes. ([#97013](https://github.com/kubernetes/kubernetes/pull/97013), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Remove ready file and its directory (which is created during volume SetUp) during emptyDir volume TearDown. ([#95770](https://github.com/kubernetes/kubernetes/pull/95770), [@jingxu97](https://github.com/jingxu97)) [SIG Storage]
- Avoid GCE API calls when initializing GCE CloudProvider for Kubelets.
  Avoid unnecessary GCE API calls when adding IP alises or reflecting them in Node object in GCE cloud provider. ([#96863](https://github.com/kubernetes/kubernetes/pull/96863), [@tosi3k](https://github.com/tosi3k)) [SIG Apps, Cloud Provider and Network]
- Bump node-problem-detector version to v0.8.5 to fix OOM detection in with Linux kernels 5.1+ ([#96716](https://github.com/kubernetes/kubernetes/pull/96716), [@tosi3k](https://github.com/tosi3k)) [SIG Cloud Provider, Scalability and Testing]
- Exposes and sets a default timeout for the SubjectAccessReview client for DelegatingAuthorizationOptions ([#96152](https://github.com/kubernetes/kubernetes/pull/96152), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Fix memory leak in kube-apiserver when underlying time goes forth and back. ([#96266](https://github.com/kubernetes/kubernetes/pull/96266), [@chenyw1990](https://github.com/chenyw1990)) [SIG API Machinery]
- Fix pull image error from multiple ACRs using azure managed identity ([#96355](https://github.com/kubernetes/kubernetes/pull/96355), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: resize Azure disk issue when it's in attached state ([#96705](https://github.com/kubernetes/kubernetes/pull/96705), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a bug that prevents kubectl to validate CRDs with schema using x-kubernetes-preserve-unknown-fields on object fields.
  Fix kubectl SchemaError on CRDs with schema using x-kubernetes-preserve-unknown-fields on array types. ([#96563](https://github.com/kubernetes/kubernetes/pull/96563), [@gautierdelorme](https://github.com/gautierdelorme)) [SIG API Machinery and Testing]
- Fixed kubelet creating extra sandbox for pods with RestartPolicyOnFailure after all containers succeeded ([#92614](https://github.com/kubernetes/kubernetes/pull/92614), [@tnqn](https://github.com/tnqn)) [SIG Node and Testing]
- Metric names for CSI and flexvolume drivers will include the driver name as well as the CSI plugin name. ([#96474](https://github.com/kubernetes/kubernetes/pull/96474), [@mattcary](https://github.com/mattcary)) [SIG Instrumentation and Storage]
- New Azure instance types do now have correct max data disk count information. ([#94340](https://github.com/kubernetes/kubernetes/pull/94340), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Cloud Provider and Storage]


### Other (Cleanup or Flake)

- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Client-go header logging (at verbosity levels >= 9) now masks `Authorization` header contents ([#95316](https://github.com/kubernetes/kubernetes/pull/95316), [@sfowl](https://github.com/sfowl)) [SIG API Machinery]

#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### cert-manager [2.4.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.4.3)

#### Changed
- Set docker.io as the default registry
- Made CRD install Job backoffLimit configurable (and increased the default value). ([#129](https://github.com/giantswarm/cert-manager-app/pull/129))
- Made backoffLimit for clusterissuer job configurable. ([#125](https://github.com/giantswarm/cert-manager-app/pull/125))
- Updated clusterissuer subchart API groups to `cert-manager.io/v1`. ([#124](https://github.com/giantswarm/cert-manager-app/pull/124))
- Update to upstream `v1.1.0`. ([#119](https://github.com/giantswarm/cert-manager-app/pull/119))
### Added
- Enabled configuration of certificate Secret deletion when the parent Certificate is deleted. ([#127](https://github.com/giantswarm/cert-manager-app/pull/127))



### external-dns [2.2.2](https://github.com/giantswarm/external-dns-app/releases/tag/v2.2.2)

#### Changed
- Set docker.io as the default registry
- Rework the way the txt prefix is generated (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Rework how the annotation filter value is generated (whilst still defaulting for default app). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Only template Secret if both required values are present in `values.yaml`. ([#53](https://github.com/giantswarm/external-dns-app/pull/53))
- Reworked the App to prepare it for customer use. ([#49](https://github.com/giantswarm/external-dns-app/pull/49))
  - General:
    - Pushes the app to the giantswarm app catalog.
    - Uses Helm release namespace.
    - Uses the release name for resource naming to avoid conflicts.
    - Added a values schema to catch incorrect values.
    - Generally makes the chart easier to use (fully documented values file).
  - external-dns options:
    - Allows customisation of the txt registry prefix.
    - Allows configuration of synchronisation interval.
    - Filter resources to reconcile via annotations.
  - AWS-specifc:
    - Allows the user to provide an IAM role to use.
    - Allows the user to provide the list of domains for external-dns to manage.
    - Allows configuration of batch size.
    - Allows configuration of CNAME instead of ALIAS records.
    - Allows configuration of the AWS zone type to update.
- Upgrade upstream external-dns from v0.7.4 to [v0.7.6](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.6).
### Fixed
- Adds additional options required for vmware installations. ([#74](https://github.com/giantswarm/external-dns-app/pull/74))
- Ensure CNAMEs are always used when AWS access is external. ([#62](https://github.com/giantswarm/external-dns-app/pull/62))
- Revert location of AWS API credentials in `values.yaml`. ([#57](https://github.com/giantswarm/external-dns-app/pull/57))
### Added
- Add crd source if the provider is vmware. ([#72](https://github.com/giantswarm/external-dns-app/pull/72))
- Allow the sync policy to be configured. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports customisation of the txt-owner-id (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports dry-run mode and warns the user if enabled. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))



### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry
- Check ca.crt expiries in TLS secrets. ([#109](https://github.com/giantswarm/cert-exporter/pull/109))
### Added
- Add exceptions in NetworkPolicies to allow DNS to work correctly through port 53.
- Add new metric (`cert_exporter_secret_not_after`) which tracks expiry of TLS certificates stored in Kubernetes secrets. ([#92](https://github.com/giantswarm/cert-exporter/pull/92))



### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Changed
- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.
- Deploy `giantswarm-critical` PriorityClass when it's not found.
### Added
- Updating namespace metadata using namespaceConfig in `Chart` CRs.
- Pause Chart CR reconciliation when it has chart-operator.giantswarm.io/paused=true annotation.
- Use diff key when logging differences between the current and desired release.
- Add support for skip CRD flag when installing Helm releases.
- Added last reconciled timestamp as metrics.
- Print difference between current release and desired release.
- Add Vertical Pod Autoscaler support.
### Fixed
- Stop updating Helm release if it has failed the previous 5 attempts.
- Only create VPA if autoscaling API group is present.



### kiam [1.7.1](https://github.com/giantswarm/kiam-app/releases/tag/v1.7.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.2.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.2.2)

#### Changed
- Set docker.io as the default registry



### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



