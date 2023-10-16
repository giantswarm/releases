# :zap: Giant Swarm Release v19.2.0 for AWS :zap:

This is a security release bringing newest stable Flatcar where most of recent CVEs have been adressed. Additionally functionality of Cilium ENI mode has been improved by fixing maximum pod calculations per node as well as VPC peerings.

## Change details


### app-operator [6.8.1](https://github.com/giantswarm/app-operator/releases/tag/v6.8.1)

#### Fixed
- Use the right name for Chart CR to be deleted.



### aws-operator [14.23.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.23.0)

#### Fixed
- Cleanup `kube-proxy` VPA after switching to Cilium.
- Bump k8scc to enable max pod calculations when cilium is in ENI IPAM mode.



### cluster-operator [5.9.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.9.0)

#### Changed
- Disable masquerading when cilium is in ENI mode.



### containerlinux [3602.2.0](https://www.flatcar-linux.org/releases/#release-3602.2.0)

_Changes since **Beta 3602.1.6**_
 
#### Security fixes:
 
 - Linux ([CVE-2023-42755](https://nvd.nist.gov/vuln/detail/CVE-2023-42755))
 
 #### Bug fixes:
 
 - Triggered re-reading of partition table to fix adding partitions to the boot disk ([scripts#1202](https://github.com/flatcar/scripts/pull/1202))
 
 #### Changes:
 
 - Use qcow2 compressed format instead of additional compression layer in Qemu images ([Flatcar#1135](https://github.com/flatcar/Flatcar/issues/1135), [scripts#1132](https://github.com/flatcar/scripts/pull/1132))
 
 #### Updates:
 
 - Linux ([5.15.133](https://lwn.net/Articles/945380))

_Changes compared to **Stable 3510.2.8**_

#### Security fixes:
 
 - Linux ([CVE-2023-42752](https://nvd.nist.gov/vuln/detail/CVE-2023-42752), [CVE-2023-42753](https://nvd.nist.gov/vuln/detail/CVE-2023-42753), [CVE-2023-42755](https://nvd.nist.gov/vuln/detail/CVE-2023-42755), [CVE-2023-4623](https://nvd.nist.gov/vuln/detail/CVE-2023-4623), [CVE-2023-4921](https://nvd.nist.gov/vuln/detail/CVE-2023-4921))
 - Go ([CVE-2023-24532](https://nvd.nist.gov/vuln/detail/CVE-2023-24532), [CVE-2023-24534](https://nvd.nist.gov/vuln/detail/CVE-2023-24534), [CVE-2023-24536](https://nvd.nist.gov/vuln/detail/CVE-2023-24536), [CVE-2023-24537](https://nvd.nist.gov/vuln/detail/CVE-2023-24537), [CVE-2023-24538](https://nvd.nist.gov/vuln/detail/CVE-2023-24538), [CVE-2023-24539](https://nvd.nist.gov/vuln/detail/CVE-2023-24539), [CVE-2023-24540](https://nvd.nist.gov/vuln/detail/CVE-2023-24540), [CVE-2023-29400](https://nvd.nist.gov/vuln/detail/CVE-2023-29400), [CVE-2022-41723](https://nvd.nist.gov/vuln/detail/CVE-2022-41723), [CVE-2022-41724](https://nvd.nist.gov/vuln/detail/CVE-2022-41724), [CVE-2022-41725](https://nvd.nist.gov/vuln/detail/CVE-2022-41725))
 - bash ([CVE-2022-3715](https://nvd.nist.gov/vuln/detail/CVE-2022-3715))
 - c-ares ([CVE-2022-4904](https://nvd.nist.gov/vuln/detail/CVE-2022-4904))
 - containerd ([CVE-2023-25153](https://nvd.nist.gov/vuln/detail/CVE-2023-25153), [CVE-2023-25173](https://nvd.nist.gov/vuln/detail/CVE-2023-25173))
 - curl ([CVE-2023-23914](https://nvd.nist.gov/vuln/detail/CVE-2023-23914), [CVE-2023-23915](https://nvd.nist.gov/vuln/detail/CVE-2023-23915) and [CVE-2023-23916](https://nvd.nist.gov/vuln/detail/CVE-2023-23916), [CVE-2023-27533](https://nvd.nist.gov/vuln/detail/CVE-2023-27533), [CVE-2023-27534](https://nvd.nist.gov/vuln/detail/CVE-2023-27534), [CVE-2023-27535](https://nvd.nist.gov/vuln/detail/CVE-2023-27535), [CVE-2023-27536](https://nvd.nist.gov/vuln/detail/CVE-2023-27536), [CVE-2023-27537](https://nvd.nist.gov/vuln/detail/CVE-2023-27537), [CVE-2023-27538](https://nvd.nist.gov/vuln/detail/CVE-2023-27538))
 - Docker ([CVE-2023-28840](https://nvd.nist.gov/vuln/detail/CVE-2023-28840), [CVE-2023-28841](https://nvd.nist.gov/vuln/detail/CVE-2023-28841), [CVE-2023-28842](https://nvd.nist.gov/vuln/detail/CVE-2023-28842))
 - e2fsprogs ([CVE-2022-1304](https://nvd.nist.gov/vuln/detail/CVE-2022-1304))
 - git ([CVE-2023-22490](https://nvd.nist.gov/vuln/detail/CVE-2023-22490), [CVE-2023-23946](https://nvd.nist.gov/vuln/detail/CVE-2023-23946))
 - GnuTLS ([CVE-2023-0361](https://nvd.nist.gov/vuln/detail/CVE-2023-0361))
 - intel-microcode ([CVE-2022-21216](https://nvd.nist.gov/vuln/detail/CVE-2022-21216), [CVE-2022-33196](https://nvd.nist.gov/vuln/detail/CVE-2022-33196), [CVE-2022-38090](https://nvd.nist.gov/vuln/detail/CVE-2022-38090))
 - less ([CVE-2022-46663](https://nvd.nist.gov/vuln/detail/CVE-2022-46663))
 - libxml2 ([CVE-2023-28484](https://nvd.nist.gov/vuln/detail/CVE-2023-28484), [CVE-2023-29469](https://nvd.nist.gov/vuln/detail/CVE-2023-29469))
 - OpenSSH ([CVE-2023-25136](https://nvd.nist.gov/vuln/detail/CVE-2023-25136), [CVE-2023-28531](https://nvd.nist.gov/vuln/detail/CVE-2023-28531), [CVE-2023-38408](https://nvd.nist.gov/vuln/detail/CVE-2023-38408))
 - OpenSSL ([CVE-2022-4203](https://nvd.nist.gov/vuln/detail/CVE-2022-4203), [CVE-2022-4304](https://nvd.nist.gov/vuln/detail/CVE-2022-4304), [CVE-2022-4450](https://nvd.nist.gov/vuln/detail/CVE-2022-4450), [CVE-2023-0215](https://nvd.nist.gov/vuln/detail/CVE-2023-0215), [CVE-2023-0216](https://nvd.nist.gov/vuln/detail/CVE-2023-0216), [CVE-2023-0217](https://nvd.nist.gov/vuln/detail/CVE-2023-0217), [CVE-2023-0286](https://nvd.nist.gov/vuln/detail/CVE-2023-0286), [CVE-2023-0401](https://nvd.nist.gov/vuln/detail/CVE-2023-0401), [CVE-2023-0464](https://nvd.nist.gov/vuln/detail/CVE-2023-0464), [CVE-2023-0465](https://nvd.nist.gov/vuln/detail/CVE-2023-0465), [CVE-2023-0466](https://nvd.nist.gov/vuln/detail/CVE-2023-0466), [CVE-2023-1255](https://nvd.nist.gov/vuln/detail/CVE-2023-1255))
 - runc ([CVE-2023-25809](https://nvd.nist.gov/vuln/detail/CVE-2023-25809), [CVE-2023-27561](https://nvd.nist.gov/vuln/detail/CVE-2023-27561), [CVE-2023-28642](https://nvd.nist.gov/vuln/detail/CVE-2023-28642))
 - tar ([CVE-2022-48303](https://nvd.nist.gov/vuln/detail/CVE-2022-48303))
 - torcx ([CVE-2022-32149](https://nvd.nist.gov/vuln/detail/CVE-2022-32149))
 - vim ([CVE-2023-0288](https://nvd.nist.gov/vuln/detail/CVE-2023-0288), [CVE-2023-0433](https://nvd.nist.gov/vuln/detail/CVE-2023-0433), [CVE-2023-1127](https://nvd.nist.gov/vuln/detail/CVE-2023-1127), [CVE-2023-1175](https://nvd.nist.gov/vuln/detail/CVE-2023-1175), [CVE-2023-1170](https://nvd.nist.gov/vuln/detail/CVE-2023-1170))
 - SDK: dnsmasq ([CVE-2022-0934](https://nvd.nist.gov/vuln/detail/CVE-2022-0934))
 - SDK: pkgconf ([CVE-2023-24056](https://nvd.nist.gov/vuln/detail/CVE-2023-24056))
 - SDK: python ([CVE-2023-24329](https://nvd.nist.gov/vuln/detail/CVE-2023-24329))
 
 #### Bug fixes:
 
 - Ensured that `/var/log/journal/` is created early enough for systemd-journald to persist the logs on first boot ([bootengine#60](https://github.com/flatcar/bootengine/pull/60), [baselayout#29](https://github.com/flatcar/baselayout/pull/29))
 - Fixed `journalctl --user` permission issue ([Flatcar#989](https://github.com/flatcar/Flatcar/issues/989))
 - Ensured that the folder `/var/log/sssd` is created if it doesn't exist, required for `sssd.service` ([Flatcar#1096](https://github.com/flatcar/Flatcar/issues/1096))
 - Fixed a miscompilation of getfacl causing it to dump core when executed ([scripts#809](https://github.com/flatcar/scripts/pull/809))
 - Restored the reboot warning and delay for non-SSH console sessions ([locksmith#21](https://github.com/flatcar/locksmith/pull/21))
 - Triggered re-reading of partition table to fix adding partitions to the boot disk ([scripts#1202](https://github.com/flatcar/scripts/pull/1202))
 - Worked around a bash regression in `flatcar-install` and added error reporting for disk write failures ([Flatcar#1059](https://github.com/flatcar/Flatcar/issues/1059))
 
 #### Changes:
 
 - Added `pigz` to the image, a parallel gzip implementation, which is useful to speed up the (de)compression for large container image imports/exports ([coreos-overlay#2504](https://github.com/flatcar/coreos-overlay/pull/2504))
 - Added a new `flatcar-reset` tool and boot logic for selective OS resets to reconfigure the system with Ignition while avoiding config drift ([bootengine#55](https://github.com/flatcar/bootengine/pull/55), [init#91](https://github.com/flatcar/init/pull/91))
 - Enabled elfutils support in systemd-coredump. A backtrace will now appear in the journal for any program that dumps core ([coreos-overlay#2489](https://github.com/flatcar/coreos-overlay/pull/2489))
 - Improved the OS reset tool to offer preview, backup and restore ([init#94](https://github.com/flatcar/init/pull/94))
 - On boot any files in `/etc` that are the same as provided by the booted `/usr/share/flatcar/etc` default for the overlay mount on `/etc` are deleted to ensure that future updates of `/usr/share/flatcar/etc` are propagated - to opt out create `/etc/.no-dup-update` in case you want to keep an unmodified config file as is or because you fear that a future Flatcar version may use the same file as you at which point your copy is cleaned up and any other future Flatcar changes would be applied ([bootengine#54](https://github.com/flatcar/bootengine/pull/54))
 - Switched systemd log reporting to the combined format of both unit description, as before, and now the unit name to easily find the unit ([coreos-overlay#2436](https://github.com/flatcar/coreos-overlay/pull/2436))
 - `/etc` is now set up as overlayfs with the original `/etc` folder being the store for changed files/directories and `/usr/share/flatcar/etc` providing the lower default directory tree ([bootengine#53](https://github.com/flatcar/bootengine/pull/53), [scripts#666](https://github.com/flatcar/scripts/pull/666))
 - Changed coreos-cloudinit to now set the short hostname instead of the FQDN when fetched from the metadata service ([coreos-cloudinit#19](https://github.com/flatcar/coreos-cloudinit/pull/19))
 - Use qcow2 compressed format instead of additional compression layer in Qemu images ([Flatcar#1135](https://github.com/flatcar/Flatcar/issues/1135), [scripts#1132](https://github.com/flatcar/scripts/pull/1132))
 
 #### Updates:
 
 - Linux ([5.15.133](https://lwn.net/Articles/945380) (includes [5.15.132](https://lwn.net/Articles/944877), [5.15.131](https://lwn.net/Articles/943755), [5.15.130](https://lwn.net/Articles/943404), [5.15.129](https://lwn.net/Articles/943113), [5.15.128](https://lwn.net/Articles/942866), [5.15.127](https://lwn.net/Articles/941775), [5.15.126](https://lwn.net/Articles/941273), [5.15.125](https://lwn.net/Articles/940798), [5.15.124](https://lwn.net/Articles/940339), [5.15.123](https://lwn.net/Articles/939424), [5.15.122](https://lwn.net/Articles/939104), [5.15.121](https://lwn.net/Articles/939016), [5.15.120](https://lwn.net/Articles/937404), [5.15.119](https://lwn.net/Articles/936675), [5.15.118](https://lwn.net/Articles/935584), [5.15.117](https://lwn.net/Articles/934622), [5.15.116](https://lwn.net/Articles/934320), [5.15.115](https://lwn.net/Articles/933909), [5.15.114](https://lwn.net/Articles/933280), [5.15.113](https://lwn.net/Articles/932883), [5.15.112](https://lwn.net/Articles/932134), [5.15.111](https://lwn.net/Articles/931652), [5.15.110](https://lwn.net/Articles/930600), [5.15.109](https://lwn.net/Articles/930263), [5.15.108](https://lwn.net/Articles/929679), [5.15.107](https://lwn.net/Articles/929015/), [5.15.106](https://lwn.net/Articles/928343), [5.15.105](https://lwn.net/Articles/927860), [5.15.104](https://lwn.net/Articles/926873), [5.15.103](https://lwn.net/Articles/926415), [5.15.102](https://lwn.net/Articles/925991), [5.15.101](https://lwn.net/Articles/925939), [5.15.100](https://lwn.net/Articles/925913), [5.15.99](https://lwn.net/Articles/925844)))
 - Linux Firmware ([20230404](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230404) (includes [20230310](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230310), [20230210](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230210)))
 - Go ([1.19.9](https://go.dev/doc/devel/release#go1.19.9) (includes [1.19.8](https://go.dev/doc/devel/release#go1.19.8), [1.19.7](https://go.dev/doc/devel/release#go1.19.7), [1.19.6](https://go.dev/doc/devel/release#go1.19.6)))
 - bash ([5.2](https://lists.gnu.org/archive/html/bash-announce/2022-09/msg00000.html))
 - bind tools ([9.16.37](https://bind9.readthedocs.io/en/v9_16_37/notes.html#notes-for-bind-9-16-37))
 - bpftool ([6.2.1](https://kernelnewbies.org/LinuxChanges#Linux_6.2.Tracing.2C_perf_and_BPF))
 - btrfs-progs ([6.0.2](https://btrfs.readthedocs.io/en/latest/CHANGES.html#btrfs-progs-6-0-2-2022-11-24), includes [6.0](https://btrfs.readthedocs.io/en/latest/CHANGES.html#btrfs-progs-6-0-2022-10-11))
 - c-ares ([1.19.0](https://c-ares.org/changelog.html#1_19_0))
 - containerd ([1.6.21](https://github.com/containerd/containerd/releases/tag/v1.6.21) (includes [1.6.20](https://github.com/containerd/containerd/releases/tag/v1.6.20), [1.6.19](https://github.com/containerd/containerd/releases/tag/v1.6.19) [1.6.18](https://github.com/containerd/containerd/releases/tag/v1.6.18))
 - curl ([8.0.1](https://curl.se/changes.html#8_0_1) (includes [7.88.1](https://curl.se/changes.html#7_88_1), [7.88.0](https://curl.se/changes.html#7_88_0)))
 - diffutils ([3.9](https://savannah.gnu.org/forum/forum.php?forum_id=10282))
 - Docker ([20.10.24](https://docs.docker.com/engine/release-notes/20.10/#201024))
 - e2fsprogs ([1.47.0](https://e2fsprogs.sourceforge.net/e2fsprogs-release.html##1.47.0) (includes [1.46.6](https://e2fsprogs.sourceforge.net/e2fsprogs-release.html#1.46.6)))
 - findutils ([4.9.0](https://lists.gnu.org/archive/html/info-gnu/2022-02/msg00003.html))
 - gcc ([12.2.1](https://gcc.gnu.org/gcc-12/changes.html))
 - gdb ([13.1.90](https://lwn.net/Articles/923819/))
 - git ([2.39.2](https://github.com/git/git/blob/v2.39.2/Documentation/RelNotes/2.39.2.txt))
 - GLib ([2.74.6](https://gitlab.gnome.org/GNOME/glib/-/tags/2.74.6) (includes [2.74.5](https://gitlab.gnome.org/GNOME/glib/-/tags/2.74.5)))
 - GnuTLS ([3.8.0](https://gitlab.com/gnutls/gnutls/-/blob/3.8.0/NEWS))
 - ignition ([2.15.0](https://coreos.github.io/ignition/release-notes/#ignition-2150-2023-02-21))
 - intel-microcode ([20230214](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20230214))
 - iperf ([3.13](https://github.com/esnet/iperf/blob/3.13/RELNOTES.md))
 - iputils ([20221126](https://github.com/iputils/iputils/releases/tag/20221126))
 - less ([608](http://www.greenwoodsoftware.com/less/news.608.html))
 - libarchive ([3.6.2](https://github.com/libarchive/libarchive/releases/tag/v3.6.2))
 - libpcap ([1.10.3](https://git.tcpdump.org/libpcap/blob/refs/tags/libpcap-1.10.3:/CHANGES) (includes [1.10.2](https://git.tcpdump.org/libpcap/blob/refs/tags/libpcap-1.10.2:/CHANGES)))
 - libpcre2 ([10.42](https://github.com/PCRE2Project/pcre2/blob/pcre2-10.42/NEWS))
 - libxml2 ([2.10.4](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.10.4))
 - multipath-tools ([0.9.4](https://github.com/opensvc/multipath-tools/commits/0.9.4))
 - OpenSSH ([9.3](http://www.openssh.com/releasenotes.html#9.3) (includes [9.2](http://www.openssh.com/releasenotes.html#9.2)))
 - OpenSSL ([3.0.8](https://github.com/openssl/openssl/blob/openssl-3.0.8/NEWS.md#major-changes-between-openssl-307-and-openssl-308-7-feb-2023))
 - pinentry ([1.2.1](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=pinentry.git;a=blob;f=NEWS;h=c080b34e57d01a6ccca9d2996d7096c42b1a3f84;hb=8ab1682e80a2b4185ee9ef66cbb44340245966fc))
 - qemu guest agent ([7.1.0](https://wiki.qemu.org/ChangeLog/7.1#Guest_agent))
 - readline ([8.2](https://lists.gnu.org/archive/html/info-gnu/2022-09/msg00013.html))
 - runc ([1.1.7](https://github.com/opencontainers/runc/releases/tag/v1.1.7) (includes [1.1.6](https://github.com/opencontainers/runc/releases/tag/v1.1.6), [1.1.5](https://github.com/opencontainers/runc/releases/tag/v1.1.5)))
 - socat ([1.7.4.4](https://repo.or.cz/socat.git/blob/refs/tags/tag-1.7.4.4:/CHANGES))
 - sqlite ([3.41.2](https://sqlite.org/releaselog/3_41_2.html))
 - strace ([6.1](https://github.com/strace/strace/releases/tag/v6.1))
 - traceroute (2.1.1)
 - vim ([9.0.1403](https://github.com/vim/vim/releases/tag/v9.0.1403) (includes [9.0.1363](https://github.com/vim/vim/releases/tag/v9.0.1363)))
 - XZ utils ([5.4.2](https://github.com/tukaani-project/xz/releases/tag/v5.4.2))
 - Zstandard ([1.5.4](https://github.com/facebook/zstd/releases/tag/v1.5.4) (includes [1.5.2](https://github.com/facebook/zstd/releases/tag/v1.5.2), [1.5.1](https://github.com/facebook/zstd/releases/tag/v1.5.1) and [1.5.0](https://github.com/facebook/zstd/releases/tag/v1.5.0)))
 - SDK: cmake ([3.25.2](https://cmake.org/cmake/help/v3.25/release/3.25.html))
 - SDK: dnsmasq ([2.89](https://lists.thekelleys.org.uk/pipermail/dnsmasq-discuss/2023q1/016859.html))
 - SDK: pahole ([1.24](https://github.com/acmel/dwarves/releases/tag/v1.24))
 - SDK: portage ([3.0.44](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.44))
 - SDK: python ([3.10.10](https://docs.python.org/3.10/whatsnew/changelog.html#python-3-10-10-final) (includes [3.10.9](https://docs.python.org/3.10/whatsnew/changelog.html#python-3-10-9-final), [3.10](https://www.python.org/downloads/release/python-3100/)))
 - SDK: Rust ([1.68.2](https://github.com/rust-lang/rust/releases/tag/1.68.2) (includes [1.68.0](https://github.com/rust-lang/rust/releases/tag/1.68.0), [1.67.1](https://github.com/rust-lang/rust/releases/tag/1.67.1)))
 - SDK: nano ([7.2](https://git.savannah.gnu.org/cgit/nano.git/tree/NEWS?h=v7.2))
 - VMware: open-vm-tools ([12.2.0](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.2.0))


### cert-exporter [2.7.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.7.0)

#### Changed
- Add Service Monitor.



### chart-operator [3.0.0](https://github.com/giantswarm/chart-operator/releases/tag/v3.0.0)

#### Removed
- Removed `giantswarm.io/monitoring: "true"` label from the `Service` resource. To get metrics `chart-operator` should
  be from now on used in conjunction with `chart-operator-extensions` version `v1.1.1` or later to deploy
  `ServiceMonitor` resource for it. It was split up as `chart-operator` is one of the first component to get into
  a cluster that will deploy most other things, for example Prometheus that will eventually actually deploy the
  CRD for `ServiceMonitor`.



### cilium [0.13.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.13.0)

#### Added
- Support removal of previously-deployed default policies by setting `defaultPolicies.enabled=false` and `defaultPolicies.remove=false`



### external-dns [2.42.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.42.0)

#### Changed
- Make CRD install job compliant with PSS ([#309](https://github.com/giantswarm/external-dns-app/pull/309)).



### etcd-kubernetes-resources-count-exporter [1.7.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.7.0)

#### Changed
- Replace condition for PSP CR installation.
- Disable debugging.



### coredns [1.19.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.19.0)

#### Changed
- Make App compliant with PSS policies ([#234](https://github.com/giantswarm/coredns-app/pull/234)):
  - Set seccompProfile to `RuntimeDefault`.
  - Fix capabilities typo.
  - Remove `NET_BIND_SERVICE` capabilities.
  - Set `runAsNonRoot` as true.



### metrics-server [2.4.1](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.4.1)

#### Added
- Add possibility to configure `hostNetwork`.
#### Changed
- Upgrade metrics-server to v0.6.4.



### net-exporter [1.18.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.18.0)

#### Changed
- Enable PSP resource deployment based on global value.



### observability-bundle [0.8.7](https://github.com/giantswarm/observability-bundle/releases/tag/v0.8.7)

#### Changed
- Upgrade `prometheus-agent` to 0.6.4.



