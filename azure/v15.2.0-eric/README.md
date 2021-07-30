# :zap: Giant Swarm Release v15.2.0-eric for Azure :zap:

<< Add description here >>

## Change details


### aws-operator [10.6.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.6.1)

#### Changed
- Upgrade `k8scloudconfig` to v10.8.1 which includes a change to better determine if memory eviction thresholds are crossed.



### kubernetes [1.21.3](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.3)

#### Feature
- Kubernetes is now built with Golang 1.16.6 ([#103670](https://github.com/kubernetes/kubernetes/pull/103670), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
#### Bug or Regression
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102969](https://github.com/kubernetes/kubernetes/pull/102969), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/structured-merge-diff/v4: v4.1.0 → v4.1.2
#### Removed
_Nothing has changed._



### containerlinux [2905.2.0](https://www.flatcar-linux.org/releases/#release-2905.2.0)

_Changes since **Beta 2905.1.0**_

**Security Fixes**



* containerd ([CVE-2021-32760](https://nvd.nist.gov/vuln/detail/CVE-2021-32760))
* curl (CVE-2021-22922, CVE-2021-22923, CVE-2021-22924, CVE-2021-22925, CVE-2021-22926)
* linux ([CVE-2020-26541](https://nvd.nist.gov/vuln/detail/CVE-2020-26541), [CVE-2021-35039](https://nvd.nist.gov/vuln/detail/CVE-2021-35039), [CVE-2021-22543](https://nvd.nist.gov/vuln/detail/CVE-2021-22543), CVE-2021-3609, CVE-2021-3655, [CVE-2021-33909](https://nvd.nist.gov/vuln/detail/CVE-2021-33909))

**Updates**



* Linux ([5.10.52](https://lwn.net/Articles/863648/))
* curl ([7.78](https://curl.se/changes.html#7_78_0))
* containerd ([1.5.4](https://github.com/containerd/containerd/releases/tag/v1.5.4))

_Changes since **Stable 2765.2.6**_

**Security Fixes:**



* Linux ([CVE-2020-26541](https://nvd.nist.gov/vuln/detail/CVE-2020-26541), [CVE-2021-35039](https://nvd.nist.gov/vuln/detail/CVE-2021-35039), [CVE-2021-22543](https://nvd.nist.gov/vuln/detail/CVE-2021-22543), CVE-2021-3609, CVE-2021-3655, [CVE-2021-33909](https://nvd.nist.gov/vuln/detail/CVE-2021-33909), [CVE-2021-34693](https://nvd.nist.gov/vuln/detail/CVE-2021-34693), [CVE-2021-33624](https://nvd.nist.gov/vuln/detail/CVE-2021-33624))
* containerd ([CVE-2021-32760](https://nvd.nist.gov/vuln/detail/CVE-2021-32760))
* curl (CVE-2021-22922, CVE-2021-22923, CVE-2021-22924, CVE-2021-22925, CVE-2021-22926)
* boost ([CVE-2012-2677](https://nvd.nist.gov/vuln/detail/CVE-2012-2677))
* Docker ([CVE-2021-21285](https://nvd.nist.gov/vuln/detail/CVE-2021-21285),[ CVE-2021-21284](https://nvd.nist.gov/vuln/detail/CVE-2021-21284))
* c-ares ([CVE-2020-8277](https://nvd.nist.gov/vuln/detail/CVE-2020-8277))
* coreutils ([CVE-2017-7476](https://nvd.nist.gov/vuln/detail/CVE-2017-7476))
* dbus ([CVE-2020-35512](https://nvd.nist.gov/vuln/detail/CVE-2020-35512))
* dnsmasq ([CVE-2020-25681](https://nvd.nist.gov/vuln/detail/CVE-2020-25681),[ CVE-2020-25682](https://nvd.nist.gov/vuln/detail/CVE-2020-25682),[ CVE-2020-25683](https://nvd.nist.gov/vuln/detail/CVE-2020-25683),[ CVE-2020-25684](https://nvd.nist.gov/vuln/detail/CVE-2020-25683),[ CVE-2020-25685](https://nvd.nist.gov/vuln/detail/CVE-2020-25685),[ CVE-2020-25686](https://nvd.nist.gov/vuln/detail/CVE-2020-25686),[ CVE-2020-25687](https://nvd.nist.gov/vuln/detail/CVE-2020-25687))
* git ([CVE-2021-21300](https://nvd.nist.gov/vuln/detail/CVE-2021-21300))
* glib ([CVE-2021-28153](https://nvd.nist.gov/vuln/detail/CVE-2021-28153),[ CVE-2021-27218](https://nvd.nist.gov/vuln/detail/CVE-2021-27218),[ CVE-2021-27219](https://nvd.nist.gov/vuln/detail/CVE-2021-27219))
* gnutls ([CVE-2021-20231](https://nvd.nist.gov/vuln/detail/CVE-2021-20231),[ CVE-2021-20232](https://nvd.nist.gov/vuln/detail/CVE-2021-20232))
* intel-microcode ([CVE-2020-8696](https://nvd.nist.gov/vuln/detail/CVE-2020-8696),[ CVE-2020-8698](https://nvd.nist.gov/vuln/detail/CVE-2020-8698))
* libxml2 ([CVE-2021-3516](https://nvd.nist.gov/vuln/detail/CVE-2021-3516),[CVE-2021-3517](https://nvd.nist.gov/vuln/detail/CVE-2021-3517),[CVE-2021-3518](https://nvd.nist.gov/vuln/detail/CVE-2021-3518),CVE-2021-3541)
* ncurses ([CVE-2019-17594](https://nvd.nist.gov/vuln/detail/CVE-2019-17594),[ CVE-2019-17595](https://nvd.nist.gov/vuln/detail/CVE-2019-17595))
* openldap ([CVE-2020-36221](https://nvd.nist.gov/vuln/detail/CVE-2020-36221),[ CVE-2020-36222](https://nvd.nist.gov/vuln/detail/CVE-2020-36222),[ CVE-2020-36223](https://nvd.nist.gov/vuln/detail/CVE-2020-36223),[ CVE-2020-36224](https://nvd.nist.gov/vuln/detail/-2020-36224),[ CVE-2020-36225](https://nvd.nist.gov/vuln/detail/CVE-2020-36225),[ CVE-2020-36226](https://nvd.nist.gov/vuln/detail/CVE-2020-36226),[ CVE-2020-36227](https://nvd.nist.gov/vuln/detail/CVE-2020-36227),[ CVE-2020-36228](https://nvd.nist.gov/vuln/detail/CVE-2020-36228),[ CVE-2020-36229](https://nvd.nist.gov/vuln/detail/CVE-2020-36229),[ CVE-2020-36230](https://nvd.nist.gov/vuln/detail/CVE-2020-36230))
* samba ([CVE-2020-14318](https://nvd.nist.gov/vuln/detail/CVE-2020-14318),[ CVE-2020-14323](https://nvd.nist.gov/vuln/detail/CVE-2020-14323),[ CVE-2020-14383](https://nvd.nist.gov/vuln/detail/CVE-2020-14383))
* sqlite ([CVE-2021-20227](https://nvd.nist.gov/vuln/detail/CVE-2021-20227))
* binutils ([CVE-2021-20197](https://nvd.nist.gov/vuln/detail/CVE-2021-20197),[CVE-2021-3487](https://nvd.nist.gov/vuln/detail/CVE-2021-3487))

**Bug Fixes:**



* passwd: use correct GID for tss ([baselayout#15](https://github.com/kinvolk/baselayout/pull/15))
* flatcar-eks: add missing mkdir and update to latest versions ([coreos-overlay#817](https://github.com/kinvolk/coreos-overlay/pull/817))
* gmerge: Stop installing gmerge script ([coreos-overlay#828](https://github.com/kinvolk/coreos-overlay/pull/828))
* Add explicit path to the binary call in the coreos-metadata unit file ([Issue #360](https://github.com/kinvolk/Flatcar/issues/360))
* Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))

**Changes**



* Docker: disabled SELinux support in the Docker daemon
* The pam_faillock PAM module was enabled as replacement for the removed pam_tally2 module and will temporarily lock an account if there were login attempts with a wrong password. The faillock command can be used to show the current state. With pam_tally2 there was no limit for wrong password login attempts but with faillock the default is already restricting the attempts. The default behavior was relaxed to allow 5 wrong passwords per two minutes, and a one minute account lock time. This does not apply to logins with an SSH key. ([baselayout#17](https://github.com/kinvolk/baselayout/pull/17))
* The etcd and flannel services are now run with Docker and any rkt-based customizations of the etcd-member and flanneld services not supported anymore. Also, because the flanneld service relies on Docker and will restart Docker after applying the new configuration, it is not possible anymore to set Requires=flanneld.service for docker.service and instead it’s enough to have flanneld.service enabled. ([coreos-overlay#857](https://github.com/kinvolk/coreos-overlay/pull/857))
* toolbox: replace rkt with docker ([coreos-overlay#881](https://github.com/kinvolk/coreos-overlay/pull/881))
* flatcar-install: add parameters to make wget more resilient ([init#35](https://github.com/kinvolk/init/pull/35))
* flatcar-install: Add -D flag to only download the image file ([Flatcar#248](https://github.com/kinvolk/Flatcar/issues/248))
* flatcar-install: Detect device mapper (e.g., LVM/LUKS) usage when searching for free drives with the -s flag ([Flatcar#332](https://github.com/kinvolk/Flatcar/issues/332))
* motd: Add OEM information to motd output ([init#34](https://github.com/kinvolk/init/pull/34))
* open-iscsi: Command substitution in iscsi-init system service ([coreos-overlay#801](https://github.com/kinvolk/coreos-overlay/pull/801))
* sshd: use secure crypto algos only ([kinvolk/coreos-overlay#852](https://github.com/kinvolk/coreos-overlay/pull/852))
* kernel: enable kernel config CONFIG_BPF_LSM ([kinvolk/coreos-overlay#846](https://github.com/kinvolk/coreos-overlay/pull/846))
* bootengine: set hostname for EC2 and OpenStack from metadata ([kinvolk/coreos-overlay#848](https://github.com/kinvolk/coreos-overlay/pull/848))
* Make the hostname setting units optional. Having the hostname units as required by the initrd.target meant that if the unit failed the machine wouldn’t start, disrupting the whole boot. ([bootengine#23](https://github.com/kinvolk/bootengine/pull/23))
* Enable using iSCSI netroot devices on Flatcar ([bootengine#22](https://github.com/kinvolk/bootengine/pull/22))
* systemd-networkd: Do not manage loopback network interface ([bootengine#24 init#40](https://github.com/kinvolk/bootengine/pull/24))
* containerd: Removed the containerd-stress binary ([coreos-overlay#858](https://github.com/kinvolk/coreos-overlay/pull/858))
* dhcpcd: Removed the dhcpcd binary from the image, systemd-networkd is the only DHCP client ([coreos-overlay#858](https://github.com/kinvolk/coreos-overlay/pull/858))
* samba: Update to EAPI=7, add new USE flags and remove deps on icu ([kinvolk/coreos-overlay#864](https://github.com/kinvolk/coreos-overlay/pull/864))
* GCE: The oem-gce.service was ported to use systemd-nspawn instead of rkt. A one-time action is required to fetch the new service file because the OEM partition is not updated: sudo curl -s -S -f -L -o /etc/systemd/system/oem-gce.service https://raw.githubusercontent.com/kinvolk/coreos-overlay/fe7b0047ef5b634ebe04c9627bbf1ce3008ee5fa/coreos-base/oem-gce/files/units/oem-gce.service && sudo systemctl daemon-reload && sudo systemctl restart oem-gce.service
* SDK: update portage and related packages to newer versions ([coreos-overlay#840](https://github.com/kinvolk/coreos-overlay/pull/840))
* SDK: Drop jobs parameter in flatcar-scripts ([flatcar-scripts#121](https://github.com/kinvolk/flatcar-scripts/pull/121))
* SDK: delete Go 1.6 ([coreos-overlay#827](https://github.com/kinvolk/coreos-overlay/pull/827))
* Update sys-apps/coreutils and make sure they have split-usr disabled for generic images ([coreos-overlay#829](https://github.com/kinvolk/coreos-overlay/pull/829))
* systemd: Fix unit installation ([coreos-overlay#810](https://github.com/kinvolk/coreos-overlay/pull/810))

**Updates**



* Linux ([5.10.52](https://lwn.net/Articles/863648/))
* Linux firmware ([20210511](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20210511))
* boost ([1.75.0](https://www.boost.org/users/history/version_1_75_0.html))
* docker ([19.03.15](https://docs.docker.com/engine/release-notes/19.03/#190315))
* c-ares ([1.17.1](https://c-ares.haxx.se/changelog.html#1_17_1))
* curl ([7.78](https://curl.se/changes.html#7_78_0))
* containerd ([1.5.4](https://github.com/containerd/containerd/releases/tag/v1.5.4))
* coreutils ([8.32](https://git.savannah.gnu.org/cgit/coreutils.git/tree/NEWS?h=v8.32))
* cri-tools ([1.19.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.19.0))
* dbus ([1.10.32](https://lists.freedesktop.org/archives/ftp-release/2020-July/000759.html))
* dnsmasq ([2.83](https://thekelleys.org.uk/dnsmasq/CHANGELOG))
* go ([1.16.5](https://go.googlesource.com/go/+/refs/tags/go1.16.5))
* git ([2.26.3](https://raw.githubusercontent.com/git/git/v2.26.3/Documentation/RelNotes/2.26.3.txt))
* glib ([2.66.8](https://gitlab.gnome.org/GNOME/glib/-/releases/2.66.8))
* gnutls ([3.7.1](https://gitlab.com/gnutls/gnutls/-/tags/3.7.1))
* intel-microcode ([20210216](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20210216))
* libxml2 ([2.9.12](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.9.12))
* multipath-tools ([0.8.5](https://github.com/opensvc/multipath-tools/releases/tag/0.8.5))
* ncurses ([6.2](https://invisible-island.net/ncurses/announce-6.2.html))
* open-iscsi ([2.1.4](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.4))
* openldap ([2.4.58](https://www.openldap.org/software/release/announce.html))
* openssh ([8.6_p1](https://www.openssh.com/txt/release-8.6))
* runc ([1.0.0_rc95](https://github.com/opencontainers/runc/releases/tag/v1.0.0-rc95))
* samba ([4.12.9](https://www.samba.org/samba/history/samba-4.12.9.html))
* sqlite ([3.34.1](https://www.sqlite.org/releaselog/3_34_1.html))
* systemd ([247.6](https://github.com/systemd/systemd-stable/releases/tag/v247.6))
* zstd ([1.4.9](https://github.com/facebook/zstd/releases/tag/v1.4.9))
* SDK: Rust ([1.52.1](https://blog.rust-lang.org/2021/05/10/Rust-1.52.1.html))
* SDK: QEMU ([5.2.0](https://wiki.qemu.org/ChangeLog/5.2))
* SDK: cmake ([3.18.5](https://cmake.org/cmake/help/latest/release/3.18.html#id1))
* SDK: binutils ([2.36.1](https://sourceware.org/pipermail/binutils/2021-February/115240.html))

**Deprecation**



* docker-1.12, rkt and kubelet-wrapper are deprecated and removed from Stable, also from subsequent channels in the future. Please read the[ removal announcement](https://groups.google.com/g/flatcar-linux-user/c/MeinndLqJO4) to know more


### cert-exporter [1.7.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.7.1)

#### Fixed
- Fix configuration version in `Chart.yaml`.



### chart-operator [2.18.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.18.0)

#### Added
- Add releasemaxhistory resource which ensures we retry at a reduced rate when
there are repeated failed upgrades.
#### Changed
- Upgrade Helm release when failed even if version or values have not changed
to handle situations like failed webhooks where we should retry.



### external-dns [2.4.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.4.0)

#### Changed
- Upgrade upstream external-dns from v0.7.6 to [v0.8.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.8.0).
- Allow to configure the minimum interval between two consecutive synchronizations triggered from kubernetes events through `externalDNS.minEventSyncInterval`.



### net-exporter [1.10.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.2)

#### Changed
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.



### cluster-autoscaler [1.21.0](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.0)

#### Added
- Add `VerticalPodAutoscaler` resource to adjust limits automatically.



