# :zap: Giant Swarm Release v19.0.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [6.7.0](https://github.com/giantswarm/app-operator/releases/tag/v6.7.0)

#### Changed
- Only include PodSecurityPolicy on clusters with policy/v1beta1 api available.
- Only include PodMonitor on clusters with monitoring.coreos.com/v1 api available.
#### Removed
- Stop pushing to `openstack-app-collection`.



### aws-operator [14.17.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.17.1)

#### Added
- Add toleration for new control-plane taint.
#### Fixed
- Ensure `net.ipv4.conf.eth0.rp_filter` is set to `2` if aws-CNI is used.
- Make `routes-fixer` script compatible with alpine.


### cluster-operator [5.6.1](https://github.com/giantswarm/cluster-operator/releases/tag/v5.6.1)

#### Fixed
- Don't enable Cilium network policies on Azure.



### containerlinux [3510.2.0](https://www.flatcar-linux.org/releases/#release-3510.2.0)

_Changes since **Stable 3374.2.5**_

#### Security fixes:

- Linux ([CVE-2022-2196](https://nvd.nist.gov/vuln/detail/CVE-2022-2196), [CVE-2022-27672](https://nvd.nist.gov/vuln/detail/CVE-2022-27672), [CVE-2022-3707](https://nvd.nist.gov/vuln/detail/CVE-2022-3707), [CVE-2023-1078](https://nvd.nist.gov/vuln/detail/CVE-2023-1078), [CVE-2023-1281](https://nvd.nist.gov/vuln/detail/CVE-2023-1281), [CVE-2023-1513](https://nvd.nist.gov/vuln/detail/CVE-2023-1513), [CVE-2023-26545](https://nvd.nist.gov/vuln/detail/CVE-2023-26545))
- bind tools ([CVE-2022-2795](https://nvd.nist.gov/vuln/detail/CVE-2022-2795), [CVE-2022-2881](https://nvd.nist.gov/vuln/detail/CVE-2022-2881), [CVE-2022-2906](https://nvd.nist.gov/vuln/detail/CVE-2022-2906), [CVE-2022-3080](https://nvd.nist.gov/vuln/detail/CVE-2022-3080), [CVE-2022-38177](https://nvd.nist.gov/vuln/detail/CVE-2022-38177), [CVE-2022-38178](https://nvd.nist.gov/vuln/detail/CVE-2022-38178))
- binutils ([CVE-2022-38126](https://nvd.nist.gov/vuln/detail/CVE-2022-38126), [CVE-2022-38127](https://nvd.nist.gov/vuln/detail/CVE-2022-38127))
- containerd ([CVE-2022-23471](https://nvd.nist.gov/vuln/detail/CVE-2022-23471))
- cpio ([CVE-2021-38185](https://nvd.nist.gov/vuln/detail/CVE-2021-38185))
- curl ([CVE-2022-35252](https://nvd.nist.gov/vuln/detail/CVE-2022-35252), [CVE-2022-43551](https://nvd.nist.gov/vuln/detail/CVE-2022-43551), [CVE-2022-43552](https://nvd.nist.gov/vuln/detail/CVE-2022-43552),[CVE-2022-32221](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-35260](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-42915](https://nvd.nist.gov/vuln/detail/CVE-2022-32221), [CVE-2022-42916](https://nvd.nist.gov/vuln/detail/CVE-2022-32221))
- dbus ([CVE-2022-42010](https://nvd.nist.gov/vuln/detail/CVE-2022-42010), [CVE-2022-42011](https://nvd.nist.gov/vuln/detail/CVE-2022-42011), [CVE-2022-42012](https://nvd.nist.gov/vuln/detail/CVE-2022-42012))
- git ([CVE-2022-39253](https://nvd.nist.gov/vuln/detail/CVE-2022-39253), [CVE-2022-39260](https://nvd.nist.gov/vuln/detail/CVE-2022-39260), [CVE-2022-23521](https://nvd.nist.gov/vuln/detail/CVE-2022-23521), [CVE-2022-41903](https://nvd.nist.gov/vuln/detail/CVE-2022-41903))
- glib ([fixes to normal form handling in GVariant](https://discourse.gnome.org/t/multiple-fixes-for-gvariant-normalisation-issues-in-glib/12835))
- Go ([CVE-2022-41717](https://nvd.nist.gov/vuln/detail/CVE-2022-41717))
- libarchive ([CVE-2022-36227](https://nvd.nist.gov/vuln/detail/CVE-2022-36227))
- libksba ([CVE-2022-47629](https://nvd.nist.gov/vuln/detail/CVE-2022-47629), [CVE-2022-3515](https://nvd.nist.gov/vuln/detail/CVE-2022-3515))
- libxml2 ([CVE-2022-40303](https://nvd.nist.gov/vuln/detail/CVE-2022-40303), [CVE-2022-40304](https://nvd.nist.gov/vuln/detail/CVE-2022-40304))
- logrotate ([CVE-2022-1348](https://nvd.nist.gov/vuln/detail/CVE-2022-1348))
- multipath-tools ([CVE-2022-41973](https://nvd.nist.gov/vuln/detail/CVE-2022-41973), [CVE-2022-41974](https://nvd.nist.gov/vuln/detail/CVE-2022-41974))
- sudo ([CVE-2023-22809](https://nvd.nist.gov/vuln/detail/CVE-2023-22809), [CVE-2022-43995](https://nvd.nist.gov/vuln/detail/CVE-2022-43995))
- systemd ([CVE-2022-3821](https://nvd.nist.gov/vuln/detail/CVE-2022-3821), [CVE-2022-4415](https://nvd.nist.gov/vuln/detail/CVE-2022-4415))
- vim ([CVE-2023-0049](https://nvd.nist.gov/vuln/detail/CVE-2023-0049), [CVE-2023-0051](https://nvd.nist.gov/vuln/detail/CVE-2023-0051), [CVE-2023-0054](https://nvd.nist.gov/vuln/detail/CVE-2023-0054), [CVE-2022-3705](https://nvd.nist.gov/vuln/detail/CVE-2022-3705), [CVE-2022-3491](https://nvd.nist.gov/vuln/detail/CVE-2022-3491), [CVE-2022-3520](https://nvd.nist.gov/vuln/detail/CVE-2022-3520), [CVE-2022-3591](https://nvd.nist.gov/vuln/detail/CVE-2022-3591), [CVE-2022-4141](https://nvd.nist.gov/vuln/detail/CVE-2022-4141), [CVE-2022-4292](https://nvd.nist.gov/vuln/detail/CVE-2022-4292), [CVE-2022-4293](https://nvd.nist.gov/vuln/detail/CVE-2022-4293),[CVE-2022-1725](https://nvd.nist.gov/vuln/detail/CVE-2022-1725), [CVE-2022-3234](https://nvd.nist.gov/vuln/detail/CVE-2022-3234), [CVE-2022-3235](https://nvd.nist.gov/vuln/detail/CVE-2022-3235), [CVE-2022-3278](https://nvd.nist.gov/vuln/detail/CVE-2022-3278), [CVE-2022-3256](https://nvd.nist.gov/vuln/detail/CVE-2022-3256), [CVE-2022-3296](https://nvd.nist.gov/vuln/detail/CVE-2022-3296), [CVE-2022-3297](https://nvd.nist.gov/vuln/detail/CVE-2022-3297), [CVE-2022-3324](https://nvd.nist.gov/vuln/detail/CVE-2022-3324), [CVE-2022-3352](https://nvd.nist.gov/vuln/detail/CVE-2022-3352), [CVE-2022-2042](https://nvd.nist.gov/vuln/detail/CVE-2022-2042), [CVE-2022-2124](https://nvd.nist.gov/vuln/detail/CVE-2022-2124), [CVE-2022-2125](https://nvd.nist.gov/vuln/detail/CVE-2022-2125), [CVE-2022-2126](https://nvd.nist.gov/vuln/detail/CVE-2022-2126), [CVE-2022-2129](https://nvd.nist.gov/vuln/detail/CVE-2022-2129), [CVE-2022-2175](https://nvd.nist.gov/vuln/detail/CVE-2022-2175), [CVE-2022-2182](https://nvd.nist.gov/vuln/detail/CVE-2022-2182), [CVE-2022-2183](https://nvd.nist.gov/vuln/detail/CVE-2022-2183), [CVE-2022-2206](https://nvd.nist.gov/vuln/detail/CVE-2022-2206), [CVE-2022-2207](https://nvd.nist.gov/vuln/detail/CVE-2022-2207), [CVE-2022-2208](https://nvd.nist.gov/vuln/detail/CVE-2022-2208), [CVE-2022-2210](https://nvd.nist.gov/vuln/detail/CVE-2022-2210), [CVE-2022-2231](https://nvd.nist.gov/vuln/detail/CVE-2022-2231), [CVE-2022-2257](https://nvd.nist.gov/vuln/detail/CVE-2022-2257), [CVE-2022-2264](https://nvd.nist.gov/vuln/detail/CVE-2022-2264), [CVE-2022-2284](https://nvd.nist.gov/vuln/detail/CVE-2022-2284), [CVE-2022-2285](https://nvd.nist.gov/vuln/detail/CVE-2022-2285), [CVE-2022-2286](https://nvd.nist.gov/vuln/detail/CVE-2022-2286), [CVE-2022-2287](https://nvd.nist.gov/vuln/detail/CVE-2022-2287), [CVE-2022-2288](https://nvd.nist.gov/vuln/detail/CVE-2022-2288), [CVE-2022-2289](https://nvd.nist.gov/vuln/detail/CVE-2022-2289), [CVE-2022-2304](https://nvd.nist.gov/vuln/detail/CVE-2022-2304), [CVE-2022-2343](https://nvd.nist.gov/vuln/detail/CVE-2022-2343), [CVE-2022-2344](https://nvd.nist.gov/vuln/detail/CVE-2022-2344), [CVE-2022-2345](https://nvd.nist.gov/vuln/detail/CVE-2022-2345), [CVE-2022-2522](https://nvd.nist.gov/vuln/detail/CVE-2022-2522), [CVE-2022-2816](https://nvd.nist.gov/vuln/detail/CVE-2022-2816), [CVE-2022-2817](https://nvd.nist.gov/vuln/detail/CVE-2022-2817), [CVE-2022-2819](https://nvd.nist.gov/vuln/detail/CVE-2022-2819), [CVE-2022-2845](https://nvd.nist.gov/vuln/detail/CVE-2022-2845), [CVE-2022-2849](https://nvd.nist.gov/vuln/detail/CVE-2022-2849), [CVE-2022-2862](https://nvd.nist.gov/vuln/detail/CVE-2022-2862), [CVE-2022-2874](https://nvd.nist.gov/vuln/detail/CVE-2022-2874), [CVE-2022-2889](https://nvd.nist.gov/vuln/detail/CVE-2022-2889), [CVE-2022-2923](https://nvd.nist.gov/vuln/detail/CVE-2022-2923), [CVE-2022-2946](https://nvd.nist.gov/vuln/detail/CVE-2022-2946), [CVE-2022-2980](https://nvd.nist.gov/vuln/detail/CVE-2022-2980), [CVE-2022-2982](https://nvd.nist.gov/vuln/detail/CVE-2022-2982), [CVE-2022-3016](https://nvd.nist.gov/vuln/detail/CVE-2022-3016), [CVE-2022-3099](https://nvd.nist.gov/vuln/detail/CVE-2022-3099), [CVE-2022-3134](https://nvd.nist.gov/vuln/detail/CVE-2022-3134), [CVE-2022-3153](https://nvd.nist.gov/vuln/detail/CVE-2022-3153))
- SDK: Python ([CVE-2015-20107](https://nvd.nist.gov/vuln/detail/CVE-2015-20107), [CVE-2020-10735](https://nvd.nist.gov/vuln/detail/CVE-2020-10735), [CVE-2021-3654](https://nvd.nist.gov/vuln/detail/CVE-2021-3654), [CVE-2022-37454](https://nvd.nist.gov/vuln/detail/CVE-2022-37454), [CVE-2022-42919](https://nvd.nist.gov/vuln/detail/CVE-2022-42919), [CVE-2022-45061](https://nvd.nist.gov/vuln/detail/CVE-2022-45061))
- SDK: qemu ([CVE-2022-4172](https://nvd.nist.gov/vuln/detail/CVE-2022-4172), [CVE-2020-14394](https://nvd.nist.gov/vuln/detail/CVE-2020-14394), [CVE-2022-0216](https://nvd.nist.gov/vuln/detail/CVE-2022-0216), [CVE-2022-35414](https://nvd.nist.gov/vuln/detail/CVE-2022-35414), [CVE-2022-3872](https://nvd.nist.gov/vuln/detail/CVE-2022-3872))
- SDK: rust ([CVE-2022-46176](https://nvd.nist.gov/vuln/detail/CVE-2022-46176), [CVE-2022-36113](https://nvd.nist.gov/vuln/detail/CVE-2022-36113), [CVE-2022-36114](https://nvd.nist.gov/vuln/detail/CVE-2022-36114))

#### Bug fixes:

- Added back Ignition support for Vagrant ([coreos-overlay#2351](https://github.com/flatcar/coreos-overlay/pull/2351))
- Added support for hardware security keys in update-ssh-keys ([update-ssh-keys#7](https://github.com/flatcar/update-ssh-keys/pull/7))
- Enabled IOMMU on arm64 kernels, the lack of which prevented some systems from booting ([coreos-overlay#2235](https://github.com/flatcar/coreos-overlay/pull/2235))
- Fixed a regression (in Alpha/Beta) where machines failed to boot if they didn't have the `core` user or group in `/etc/passwd` or `/etc/group` ([baselayout#26](https://github.com/flatcar/baselayout/pull/26))
- Fix "ext4 deadlock under heavy I/O load" kernel issue. The patch for this is included provisionally while we wait for it to be merged upstream ([Flatcar#847](https://github.com/flatcar/Flatcar/issues/847), [coreos-overlay#2315](https://github.com/flatcar/coreos-overlay/pull/2315))
- Restored the support to specify OEM partition files in Ignition when `/usr/share/oem` is given as initrd mount point ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))
- The rootfs setup in the initrd now runs systemd-tmpfiles on every boot, not only when Ignition runs, to fix a dbus failure due to missing files ([Flatcar#944](https://github.com/flatcar/Flatcar/issues/944))

#### Changes:

- Added `CONFIG_NF_CONNTRACK_BRIDGE` (for nf_conntrack_bridge) and `CONFIG_NFT_BRIDGE_META` (for nft_meta_bridge) to the kernel config to allow using conntrack rules for bridges in nftables and to match on bridge interface names ([coreos-overlay#2207](https://github.com/flatcar/coreos-overlay/pull/2207))
- Added new image signing pub key to `flatcar-install`, needed for download verification of releases built from July 2023 onwards, if you have copies of `flatcar-install` or the image signing pub key, you need to update them as well ([init#92](https://github.com/flatcar/init/pull/92))
- Change CONFIG_WIREGUARD kernel option to module to save space on boot partition ([coreos-overlay#2239](https://github.com/flatcar/coreos-overlay/pull/2239))
- Disable several arch specific arm64 kernel config options for unsupported platforms to save space on boot partition ([coreos-overlay#2239](https://github.com/flatcar/coreos-overlay/pull/2239))
- Specifying the OEM filesystem in Ignition to write files to `/usr/share/oem` is not needed anymore ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))
- Switched from `--strip-unneeded` to `--strip-debug` when installing kernel modules, which makes kernel stacktraces more accurate and makes debugging issues easier ([coreos-overlay#2196](https://github.com/flatcar/coreos-overlay/pull/2196))
- The flatcar-update tool got two new flags to customize ports used on the host while updating flatcar ([init#81](https://github.com/flatcar/init/pull/81))
- Toolbox now uses containerd to download and mount the image ([toolbox#7](https://github.com/flatcar/toolbox/pull/7))
- Add qemu-guest-agent to all amd64 images, it will be automatically enabled when qemu-ga virtio-port is detected ([coreos-overlay#2240](https://github.com/flatcar/coreos-overlay/pull/2240), [portage-stable#373](https://github.com/flatcar/portage-stable/pull/373))

#### Updates:

- Linux ([5.15.98](https://lwn.net/Articles/925080) (includes [5.15.97](https://lwn.net/Articles/925064), [5.15.96](https://lwn.net/Articles/924441), [5.15.95](https://lwn.net/Articles/924073), [5.15.94](https://lwn.net/Articles/923308), [5.15.93](https://lwn.net/Articles/922814)))
- Linux Firmware ([20230117](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230117))
- adcli ([0.9.2](https://gitlab.freedesktop.org/realmd/adcli/-/commits/8e88e3590a19006362ea8b8dfdc18bb88b3cb3b5/))
- bind tools ([9.16.36](https://bind9.readthedocs.io/en/v9_16_36/notes.html#notes-for-bind-9-16-36) (includes [9.16.34](https://bind9.readthedocs.io/en/v9_16_35/notes.html#notes-for-bind-9-16-34) and [9.16.35](https://bind9.readthedocs.io/en/v9_16_34/notes.html#notes-for-bind-9-16-35)))
- binutils ([2.39](https://sourceware.org/pipermail/binutils/2022-August/122246.html))
- bpftool ([5.19.12](https://lwn.net/Articles/909678/))
- ca-certificates ([3.89](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_89.html))
- containerd ([1.6.16](https://github.com/containerd/containerd/releases/tag/v1.6.16))
- cpio ([2.13](https://lists.gnu.org/archive/html/bug-cpio/2019-11/msg00000.html))
- curl ([7.87.0](https://curl.se/changes.html#7_87_0) (includes [7.85](https://curl.se/mail/archive-2022-08/0012.html)))
- dbus ([1.14.4](https://gitlab.freedesktop.org/dbus/dbus/-/raw/dbus-1.14.4/NEWS))
- Docker ([20.10.23](https://docs.docker.com/engine/release-notes/#201023))
- elfutils ([0.188](https://sourceware.org/pipermail/elfutils-devel/2022q4/005561.html) (includes [0.187](https://sourceware.org/pipermail/elfutils-devel/2022q2/004978.html)))
- Expat ([2.5.0](https://github.com/libexpat/libexpat/blob/R_2_5_0/expat/Changes))
- gawk ([5.2.1](https://lists.gnu.org/archive/html/help-gawk/2022-11/msg00008.html) (contains [5.2.0](https://lists.gnu.org/archive/html/help-gawk/2022-09/msg00000.html)))
- gettext ([0.21.1](https://git.savannah.gnu.org/gitweb/?p=gettext.git;a=blob;f=NEWS;h=cdbb16746c23555e70bb1e16917f5c349ce92d9e;hb=8b38ee827251cadbb90cb6cb576ae98702566288))
- git ([2.39.1](https://github.com/git/git/blob/v2.39.1/Documentation/RelNotes/2.39.1.txt) (includes [2.39.0](https://github.com/git/git/blob/v2.39.0/Documentation/RelNotes/2.39.0.txt)))
- glib ([2.74.4](https://gitlab.gnome.org/GNOME/glib/-/tags/2.74.4))
- Go ([1.19.5](https://go.dev/doc/devel/release#go1.19.5))
- glibc ([2.36](https://sourceware.org/pipermail/libc-alpha/2022-August/141193.html) (includes [2.35](https://savannah.gnu.org/forum/forum.php?forum_id=10111)))
- GnuTLS ([3.7.8](https://lists.gnupg.org/pipermail/gnutls-help/2022-September/004765.html))
- I2C tools ([4.3](https://git.kernel.org/pub/scm/utils/i2c-tools/i2c-tools.git/tree/CHANGES?id=d8bc1f1ff4b00a6bd988aa114100ae9b787f50d8))
- Intel Microcode ([20221108](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20221108))
- iptables ([1.8.8](https://www.netfilter.org/projects/iptables/files/changes-iptables-1.8.8.txt))
- iputils ([20211215](https://github.com/iputils/iputils/releases/tag/20211215))
- libcap ([2.66](https://sites.google.com/site/fullycapable/release-notes-for-libcap#h.d9ygdose5kw))
- libcap-ng ([0.8.3](https://people.redhat.com/sgrubb/libcap-ng/ChangeLog))
- libksba ([1.6.3](https://dev.gnupg.org/T6304))
- libseccomp ([2.5.4](https://github.com/seccomp/libseccomp/releases/tag/v2.5.4) (contains [2.5.2](https://github.com/seccomp/libseccomp/releases/tag/v2.5.2), [2.5.3](https://github.com/seccomp/libseccomp/releases/tag/v2.5.3)))
- libxml2 ([2.10.3](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.10.3))
- logrotate ([3.20.1](https://github.com/logrotate/logrotate/releases/tag/3.20.1))
- MIT Kerberos V ([1.20.1](https://web.mit.edu/kerberos/krb5-1.20/krb5-1.20.1.html))
- multipath-tools ([0.9.3](https://github.com/opensvc/multipath-tools/releases/tag/0.9.3))
- nettle ([3.8.1](https://git.lysator.liu.se/nettle/nettle/-/blob/990abad16ceacd070747dcc76ed16a39c129321e/ChangeLog))
- nmap ([7.93](https://nmap.org/changelog.html#7.93))
- OpenSSH ([9.1](http://www.openssh.com/releasenotes.html#9.1))
- rsync ([3.2.7](https://download.samba.org/pub/rsync/NEWS#3.2.7))
- shadow ([4.13](https://github.com/shadow-maint/shadow/releases/tag/4.13))
- sqlite ([3.40.1](https://www.sqlite.org/releaselog/3_40_1.html) (contains [3.40.0](https://www.sqlite.org/releaselog/3_40_0.html) and [3.39.4](https://sqlite.org/releaselog/3_39_4.html)))
- strace ([5.19](https://github.com/strace/strace/releases/tag/v5.19))
- sudo ([1.9.12_p2](https://github.com/sudo-project/sudo/releases/tag/SUDO_1_9_12p2))
- systemd ([252.5](https://github.com/systemd/systemd-stable/releases/tag/v252.5) (includes [252](https://github.com/systemd/systemd/releases/tag/v252)))
- vim ([9.0.1157](https://github.com/vim/vim/releases/tag/v9.0.1157) (includes [9.0.0469](https://github.com/vim/vim/releases/tag/v9.0.0469)))
- wget ([1.21.3](https://lists.gnu.org/archive/html/info-gnu/2022-02/msg00017.html))
- whois ([5.5.14](https://github.com/rfc1036/whois/commit/ab10466cf2e1ec4887f6a44375c3e29c1720157f))
- wireguard-tools ([1.0.20210914](https://github.com/WireGuard/wireguard-tools/releases/tag/v1.0.20210914))
- XZ utils ([5.4.1](https://github.com/tukaani-project/xz/releases/tag/v5.4.1) (includes [5.4.0](https://github.com/tukaani-project/xz/releases/tag/v5.4.0)))
- zlib ([1.2.13](https://github.com/madler/zlib/releases/tag/v1.2.13))
- OEM: python-oem ([3.9.16](https://www.python.org/downloads/release/python-3916/))
- SDK: boost ([1.81.0](https://www.boost.org/users/history/version_1_81_0.html))
- SDK: catalyst ([3.0.21](https://gitweb.gentoo.org/proj/catalyst.git/log/?h=3.0.21))
- SDK: cmake ([3.23.3](https://cmake.org/cmake/help/v3.23/release/3.23.html))
- SDK: file ([5.43](https://mailman.astron.com/pipermail/file/2022-September/000857.html) (includes [5.44](https://github.com/file/file/blob/FILE5_44/ChangeLog)))
- SDK: libpng ([1.6.39](http://www.libpng.org/pub/png/src/libpng-1.6.39-README.txt) (includes [1.6.38](http://www.libpng.org/pub/png/src/libpng-1.6.38-README.txt)))
- SDK: libxslt ([1.1.37](https://gitlab.gnome.org/GNOME/libxslt/-/tags/v1.1.37))
- SDK: meson ([0.62.2](https://mesonbuild.com/Release-notes-for-0-62-0.html))
- SDK: ninja ([1.11.0](https://groups.google.com/g/ninja-build/c/R2oCyDctDf8/m/-U94Y5I8AgAJ?pli=1))
- SDK: pahole ([1.23](https://git.kernel.org/pub/scm/devel/pahole/pahole.git/tag/?h=v1.23))
- SDK: perl ([5.36.0](https://perldoc.perl.org/5.36.0/perldelta))
- SDK: portage ([3.0.43](https://github.com/gentoo/portage/blob/portage-3.0.43/NEWS) (includes [3.0.42](https://github.com/gentoo/portage/blob/portage-3.0.42/NEWS), [3.0.41](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.41)))
- SDK: qemu ([7.2.0](https://wiki.qemu.org/ChangeLog/7.2) (includes [7.1.0](https://wiki.qemu.org/ChangeLog/7.1)))
- SDK: Rust ([1.67.0](https://github.com/rust-lang/rust/releases/tag/1.67.0))
- VMware: open-vm-tools ([12.1.5](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.1.5))

_Changes since **Beta 3510.1.0**_

#### Security fixes:


#### Bug fixes:

- Restored the support to specify OEM partition files in Ignition when `/usr/share/oem` is given as initrd mount point ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))

#### Changes:

- Added new image signing pub key to `flatcar-install`, needed for download verification of releases built from July 2023 onwards, if you have copies of `flatcar-install` or the image signing pub key, you need to update them as well ([init#92](https://github.com/flatcar/init/pull/92))
- Specifying the OEM filesystem in Ignition to write files to `/usr/share/oem` is not needed anymore ([bootengine#58](https://github.com/flatcar/bootengine/pull/58))

#### Updates:

- ca-certificates ([3.89](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_89.html))


### kubernetes [1.24.13](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.13)

#### Feature
- Kubernetes is now built with Go 1.19.8 ([#117132](https://github.com/kubernetes/kubernetes/pull/117132), [@xmudrii](https://github.com/xmudrii)) [SIG Release and Testing]
#### Bug or Regression
- Fix missing delete events on informer re-lists to ensure all delete events are correctly emitted and using the latest known object state, so that all event handlers and stores always reflect the actual apiserver state as best as possible ([#115901](https://github.com/kubernetes/kubernetes/pull/115901), [@odinuge](https://github.com/odinuge)) [SIG API Machinery]
- Fix: Route controller should update routes with NodeIP changed ([#116360](https://github.com/kubernetes/kubernetes/pull/116360), [@lzhecheng](https://github.com/lzhecheng)) [SIG Cloud Provider]
- Kubelet: Fix fs quota monitoring on volumes ([#116795](https://github.com/kubernetes/kubernetes/pull/116795), [@pacoxu](https://github.com/pacoxu)) [SIG Storage]
#### Other (Cleanup or Flake)
- Service session affinity timeout tests are no longer required for Kubernetes network plugin conformance due to variations in existing implementations. New conformance tests will be developed to better express conformance in future releases. ([#112806](https://github.com/kubernetes/kubernetes/pull/112806), [@dcbw](https://github.com/dcbw)) [SIG Architecture, Network and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.35 → v0.0.36
#### Removed
_Nothing has changed._



### aws-cloud-controller-manager [1.24.1-gs7](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.24.1-gs7)

#### Fixed
- Quote environment variables that contain numeric values, because it's required by kubernetes.



### aws-ebs-csi-driver [2.21.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.21.1)

#### Fixed
- Use `string` type for the proxy parameters on the `values.schema.json` file.



### cert-exporter [2.5.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.5.1)

#### Changed
- Allow requests from the api-server.
- Update icon
- Disable PSPs for k8s 1.25 and newer.



### chart-operator [2.35.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.35.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### cilium [0.10.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.10.0)

#### Changed
- Enable PDB for `cilium-operator`.



### cluster-autoscaler [1.24.0-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.0-gs2)

#### Changed
- Add 'projected' volumes to the PSP.
- Add new-pod-scale-up-delay variable.
- Disable PSPs for k8s 1.25 and newer.



### coredns [1.17.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.17.0)

#### Added
- Add scaling based on custom metrics ([#209](https://github.com/giantswarm/coredns-app/pull/209)).
#### Changed
- Decouple PDB configuration from deployment updateStrategy ([#208](https://github.com/giantswarm/coredns-app/pull/208)).



### external-dns [2.37.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.37.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### metrics-server [2.2.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.
- Switch to `apiVersion: policy/v1` for PodDisruptionBudget.



### net-exporter [1.15.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.15.0)

#### Changed
- Allow requests from the api-server.
- Disable PSPs for k8s 1.25 and newer.



### node-exporter [1.16.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### vertical-pod-autoscaler [3.4.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.2)

#### Changed
- Remove circleci job for pushing to shared app collection



### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap



### etcd-kubernetes-resources-count-exporter [1.2.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### observability-bundle [0.5.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.5.1)

#### Changed
- Remove cluster prefix to app name in _helpers.tpl



### prometheus-blackbox-exporter [0.3.2](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.3.2)

#### Added
- Add icon.



### cilium-servicemonitors [0.1.1](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.1)

#### Added
- Add overridability to the servicemonitors relabelings and metric_relabelings sections.



