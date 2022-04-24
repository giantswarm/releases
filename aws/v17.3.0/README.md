# :zap: Giant Swarm Release v17.3.0 for AWS :zap:

<< Add description here >>

## Change details

### aws-ebs-csi-driver [2.12.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.12.0)

#### Added
- Allow specifying `driverMode` for the `controller` component.
- Also push to control-plane app catalog.

#### Changed
- Allow specifying `nodeSelector` and `hostNetwork` for `controller` and `node`.
- Bump aws-ebs-csi-driver version to `v1.5.1`.


### aws-operator [11.9.2](https://github.com/giantswarm/aws-operator/releases/tag/v11.9.2)

#### Added
- Added separate service account flag for IRSA.
- Add `POD_SECURITY_GROUP_ENFORCING_MODE` to `aws-node` Daemonset.

#### Fixed
- Issuer S3 endpoint for IRSA.
- AWS Region Endpoint for IRSA.


### aws-cni [1.11.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.0)

Upgraded from version 1.10.2. Please check [upstream changelog](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.11.0) for details.

### cert-operator [2.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v2.0.1)

#### Fixed
- Bump go module major version.



### cluster-operator [4.0.2](https://github.com/giantswarm/cluster-operator/releases/tag/v4.0.2)

#### Fixed
- List apps by namespace.



### containerlinux [3139.2.0](https://www.flatcar-linux.org/releases/#release-3139.2.0)

New **Stable** Release **3139.2.0**

_Changes since **Stable 3033.2.4**_

#### Security fixes:

- Linux ([CVE-2022-1015](https://nvd.nist.gov/vuln/detail/CVE-2022-1015), [CVE-2022-1016](https://nvd.nist.gov/vuln/detail/CVE-2022-1016))
- Go ([CVE-2021-44716](https://nvd.nist.gov/vuln/detail/CVE-2021-44716), [CVE-2021-44717](https://nvd.nist.gov/vuln/detail/CVE-2021-44717))
- containerd ([CVE-2021-43816](https://nvd.nist.gov/vuln/detail/CVE-2021-43816), [CVE-2022-24769](https://nvd.nist.gov/vuln/detail/CVE-2022-24769))
- gcc ([CVE-2020-13844](https://nvd.nist.gov/vuln/detail/CVE-2020-13844))
- Ignition ([CVE-2020-14040](https://nvd.nist.gov/vuln/detail/CVE-2020-14040), [CVE-2021-38561](https://nvd.nist.gov/vuln/detail/CVE-2021-38561))
- krb5 ([CVE-2021-37750](https://nvd.nist.gov/vuln/detail/CVE-2021-37750))
- libarchive ([libarchive-1565](https://github.com/libarchive/libarchive/issues/1565), [libarchive-1566](https://github.com/libarchive/libarchive/issues/1566))
- OpenSSH ([CVE-2021-41617](https://nvd.nist.gov/vuln/detail/CVE-2021-41617))
- openssl ([CVE-2021-4044](https://nvd.nist.gov/vuln/detail/CVE-2021-4044))
- torcx ([CVE-2021-38561](https://nvd.nist.gov/vuln/detail/CVE-2021-38561), [CVE-2021-43565](https://nvd.nist.gov/vuln/detail/CVE-2021-43565))
- vim ([CVE-2021-3872](https://nvd.nist.gov/vuln/detail/CVE-2021-3872), [CVE-2021-3875](https://nvd.nist.gov/vuln/detail/CVE-2021-3875), [CVE-2021-3903](https://nvd.nist.gov/vuln/detail/CVE-2021-3903), [CVE-2021-3927](https://nvd.nist.gov/vuln/detail/CVE-2021-3927), [CVE-2021-3928](https://nvd.nist.gov/vuln/detail/CVE-2021-3928), [CVE-2021-3968](https://nvd.nist.gov/vuln/detail/CVE-2021-3968), [CVE-2021-3973](https://nvd.nist.gov/vuln/detail/CVE-2021-3973), [CVE-2021-3974](https://nvd.nist.gov/vuln/detail/CVE-2021-3974))
- SDK: edk2-ovmf ([CVE-2019-14584](https://nvd.nist.gov/vuln/detail/CVE-2019-14584), [CVE-2021-28210](https://nvd.nist.gov/vuln/detail/CVE-2021-28210), [CVE-2021-28211](https://nvd.nist.gov/vuln/detail/CVE-2021-28211), [CVE-2021-28213](https://nvd.nist.gov/vuln/detail/CVE-2021-28213))
- SDK: libxslt ([CVE-2021-30560](https://nvd.nist.gov/vuln/detail/CVE-2021-30560))
- SDK: mantle ([CVE-2021-3121](https://nvd.nist.gov/vuln/detail/CVE-2021-3121), [CVE-2021-38561](https://nvd.nist.gov/vuln/detail/CVE-2021-38561), [CVE-2021-43565](https://nvd.nist.gov/vuln/detail/CVE-2021-43565))
- SDK: QEMU ([CVE-2020-35504](https://nvd.nist.gov/vuln/detail/CVE-2020-35504), [CVE-2020-35505](https://nvd.nist.gov/vuln/detail/CVE-2020-35505), [CVE-2020-35506](https://nvd.nist.gov/vuln/detail/CVE-2020-35506), [CVE-2020-35517](https://nvd.nist.gov/vuln/detail/CVE-2020-35517), [CVE-2021-20203](https://nvd.nist.gov/vuln/detail/CVE-2021-20203), [CVE-2021-20255](https://nvd.nist.gov/vuln/detail/CVE-2021-20255), [CVE-2021-20257](https://nvd.nist.gov/vuln/detail/CVE-2021-20257), [CVE-2021-20263](https://nvd.nist.gov/vuln/detail/CVE-2021-20263), [CVE-2021-3409](https://nvd.nist.gov/vuln/detail/CVE-2021-3409), [CVE-2021-3416](https://nvd.nist.gov/vuln/detail/CVE-2021-3416), [CVE-2021-3527](https://nvd.nist.gov/vuln/detail/CVE-2021-3527), [CVE-2021-3544](https://nvd.nist.gov/vuln/detail/CVE-2021-3544), [CVE-2021-3545](https://nvd.nist.gov/vuln/detail/CVE-2021-3545), [CVE-2021-3546](https://nvd.nist.gov/vuln/detail/CVE-2021-3546), [CVE-2021-3582](https://nvd.nist.gov/vuln/detail/CVE-2021-3582), [CVE-2021-3607](https://nvd.nist.gov/vuln/detail/CVE-2021-3607), [CVE-2021-3608](https://nvd.nist.gov/vuln/detail/CVE-2021-3608), [CVE-2021-3682](https://nvd.nist.gov/vuln/detail/CVE-2021-3682))
- SDK: Rust ([CVE-2022-21658](https://nvd.nist.gov/vuln/detail/CVE-2022-21658))

#### Bug fixes:

- Excluded the Kubenet cbr0 interface from networkd's DHCP config and set it to Unmanaged to prevent interference and ensure that it is not part of the network online check ([init#55](https://github.com/flatcar-linux/init/pull/55))
- Fixed the dracut emergency Ignition log printing that had a scripting error causing the cat command to fail ([bootengine#33](https://github.com/flatcar-linux/bootengine/pull/33))
- network: Accept ICMPv6 Router Advertisements to fix IPv6 address assignment in the default DHCP setting ([init#51](https://github.com/flatcar-linux/init/pull/51), [coreos-cloudinit#12](https://github.com/flatcar-linux/coreos-cloudinit/pull/12), [bootengine#30](https://github.com/flatcar-linux/bootengine/pull/30))
- flatcar-update: Stopped checking for the `USER` environment variable which may not be set in all environments, causing the script to fail unless a workaround was used like prepending an additional `sudo` invocation ([init#58](https://github.com/flatcar-linux/init/pull/58))
- Reverted the Linux kernel commit which broke networking on AWS instances which use Intel 82559 NIC (c4/m4) ([Flatcar#665](https://github.com/flatcar-linux/Flatcar/issues/665), [coreos-overlay#1723](https://github.com/flatcar-linux/coreos-overlay/pull/1723))
- Re-added the `brd drbd nbd rbd xen-blkfront zram libarc4 lru_cache zsmalloc` kernel modules to the initramfs since they were missing compared to the Flatcar 3033.2.x releases where the 5.10 kernel is used ([bootengine#40](https://github.com/flatcar-linux/bootengine/pull/40))

#### Changes:

- Added a new flatcar-update tool to the image to ease manual updates, rollbacks, channel/release jumping, and airgapped updates ([init#53](https://github.com/flatcar-linux/init/pull/53))
- Update-engine now creates the `/run/reboot-required` flag file for [kured](https://github.com/weaveworks/kured) ([update_engine#15](https://github.com/flatcar-linux/update_engine/pull/15))
- Excluded special network interface devices like bridge, tunnel, vxlan, and veth devices from the default DHCP configuration to prevent networkd interference ([init#56](https://github.com/flatcar-linux/init/pull/56))
- Added CONFIG_NF_CT_NETLINK_HELPER (for libnetfilter_cthelper), CONFIG_NET_VRF (for virtual routing and forwarding) and CONFIG_KEY_DH_OPERATIONS (for keyutils) to the kernel config ([coreos-overlay#1524](https://github.com/flatcar-linux/coreos-overlay/pull/1524))
- Enabled the FIPS support for the Linux kernel, which users can now choose through a kernel parameter in `grub.cfg` (check it taking effect with `cat /proc/sys/crypto/fips_enabled`) ([coreos-overlay#1602](https://github.com/flatcar-linux/coreos-overlay/pull/1602))
- Enabled FIPS mode for cryptsetup ([portage-stable#312](https://github.com/flatcar-linux/portage-stable/pull/312))
- Rework the way we set up the default python intepreter in SDK - it is now without specifying a version. This should work fine as long as we keep having one version of python in SDK.
- Add a way to remove packages that are hard-blockers for update. A hard-blocker means that the package needs to be removed (for example with `emerge -C`) before an update can happen.
- Removed the pre-shipped `/etc/flatcar/update.conf` file, leaving it totally to the user to define the contents as it was unnecessarily overwriting the `/use/share/flatcar/update.conf` ([scripts#212](https://github.com/flatcar-linux/scripts/pull/212))

#### Updates:

- Linux ([5.15.32](https://lwn.net/Articles/889438)) (from 5.15.30)
- Linux headers ([5.15](https://lwn.net/Articles/876611/))
- GCC [9.4.0](https://lists.gnu.org/archive/html/info-gnu/2021-06/msg00000.html)
- acl ([2.3.1](https://git.savannah.nongnu.org/cgit/acl.git/log/?h=v2.3.1))
- attr ([2.5.1](https://git.savannah.nongnu.org/cgit/attr.git/log/?h=v2.5.1))
- audit ([3.0.6](https://listman.redhat.com/archives/linux-audit/2021-October/msg00000.html))
- boost ([1.76.0](https://www.boost.org/users/history/version_1_76_0.html))
- btrfs-progs ([5.15.1](https://btrfs.wiki.kernel.org/index.php/Changelog#btrfs-progs_v5.15_.28Nov_2021.29))
- ca-certificates ([3.77](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_77.html))
- containerd ([1.5.11](https://github.com/containerd/containerd/releases/tag/v1.5.11))
- coreutils ([8.32](https://lists.gnu.org/archive/html/coreutils-announce/2020-03/msg00000.html))
- diffutils ([3.8](https://lists.gnu.org/archive/html/info-gnu/2021-08/msg00000.html))
- ethtool ([5.10](https://git.kernel.org/pub/scm/network/ethtool/ethtool.git/tree/NEWS?h=v5.10))
- findutils ([4.8.0](https://savannah.gnu.org/forum/forum.php?forum_id=9914))
- glib ([2.68.4](https://gitlab.gnome.org/GNOME/glib/-/releases/2.68.4))
- i2c-tools ([4.2](https://git.kernel.org/pub/scm/utils/i2c-tools/i2c-tools.git/log/?h=v4.2))
- iproute2 ([5.15](https://lwn.net/ml/linux-kernel/20211101164705.6f4f2e41%40hermes.local/))
- ipset ([7.11](https://ipset.netfilter.org/changelog.html))
- iputils ([20210722](https://github.com/iputils/iputils/releases/tag/20210722))
- ipvsadm ([1.27](http://archive.linuxvirtualserver.org/html/lvs-devel/2013-09/msg00011.html))
- kmod ([29](https://git.kernel.org/pub/scm/utils/kernel/kmod/kmod.git/commit/?id=b6ecfc916a17eab8f93be5b09f4e4f845aabd3d1))
- libarchive [3.5.2](https://github.com/libarchive/libarchive/releases/tag/v3.5.2)
- libcap-ng ([0.8.2](https://github.com/stevegrubb/libcap-ng/releases/tag/v0.8.2))
- libseccomp ([2.5.1](https://github.com/seccomp/libseccomp/releases/tag/v2.5.1))
- lshw ([02.19.2b_p20210121](https://www.ezix.org/project/wiki/HardwareLiSter#Changes))
- lsof ([4.94.0](https://github.com/lsof-org/lsof/releases/tag/4.94.0))
- openssh ([8.8](http://www.openssh.com/txt/release-8.8))
- openssl ([3.0.2](https://www.openssl.org/news/changelog.html#openssl-30))
- parted ([3.4](https://savannah.gnu.org/forum/forum.php?forum_id=9924) (includes [3.3](https://savannah.gnu.org/forum/forum.php?forum_id=9569)))
- pciutils ([3.7.0](https://github.com/pciutils/pciutils/commit/864aecdea9c7db626856d8d452f6c784316a878c))
- polkit ([0.120](https://gitlab.freedesktop.org/polkit/polkit/-/blob/0.120/NEWS))
- runc ([1.1.0](https://github.com/opencontainers/runc/releases/tag/v1.1.0))
- sbsigntools ([0.9.4](https://git.kernel.org/pub/scm/linux/kernel/git/jejb/sbsigntools.git/tag/?h=v0.9.4))
- sed ([4.8](https://savannah.gnu.org/forum/forum.php?forum_id=9647))
- usbutils ([014](https://github.com/gregkh/usbutils/commit/57fb18e59cce31a50a1ca62d1e192512c905ba00))
- vim [8.2.3582](https://github.com/vim/vim/releases/tag/v8.2.3582)
- Azure: Python for OEM images ([3.9.8](https://www.python.org/downloads/release/python-398/))
- Azure: WALinuxAgent ([2.6.0.2](https://github.com/Azure/WALinuxAgent/releases/tag/v2.6.0.2))
- SDK: edk2-ovmf [202105](https://github.com/tianocore/edk2/releases/tag/edk2-stable202105)
- SDK: file ([5.40](https://mailman.astron.com/pipermail/file/2021-March/000478.html))
- SDK: ipxe [1.21.1](https://github.com/ipxe/ipxe/releases/tag/v1.21.1)
- SDK: mantle ([0.18.0](https://github.com/flatcar-linux/mantle/releases/tag/v0.18.0))
- SDK: perf ([5.15](https://kernelnewbies.org/LinuxChanges#Linux_5.15.Tracing.2C_perf_and_BPF))
- SDK: Python ([3.9.8](https://www.python.org/downloads/release/python-398/))
- SDK: qemu ([6.1.0](https://wiki.qemu.org/ChangeLog/6.1)
- SDK: Rust ([1.58.1](https://github.com/rust-lang/rust/releases/tag/1.58.1))
- SDK: seabios [1.14.0](https://seabios.org/Releases#SeaBIOS_1.14.0)
- SDK: sgabios [0.1_pre10](https://git.qemu.org/?p=sgabios.git;a=tree;h=a85446adb0e0)

_Changes since **Beta 3139.1.1**_

#### Security fixes:

- Linux ([CVE-2022-1015](https://nvd.nist.gov/vuln/detail/CVE-2022-1015), [CVE-2022-1016](https://nvd.nist.gov/vuln/detail/CVE-2022-1016))
- containerd ([CVE-2022-24769](https://nvd.nist.gov/vuln/detail/CVE-2022-24769))

#### Changes:

- Enabled FIPS mode for cryptsetup ([portage-stable#312](https://github.com/flatcar-linux/portage-stable/pull/312))

#### Updates:

- Linux ([5.15.32](https://lwn.net/Articles/889438)) (from 5.15.30)
- ca-certificates ([3.77](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_77.html))
- containerd ([1.5.11](https://github.com/containerd/containerd/releases/tag/v1.5.11))
- Azure: WALinuxAgent ([2.6.0.2](https://github.com/Azure/WALinuxAgent/releases/tag/v2.6.0.2))



### app-operator [5.9.0](https://github.com/giantswarm/app-operator/releases/tag/v5.9.0)

#### Changed
- Update `helmclient` to v4.10.0.
- Update giantswarm/appcatalog to `v0.7.0`, adding support for internal OCI chart catalogs.



### cert-exporter [2.2.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.2.0)

#### Changed
- Change priorityClass to `system-node-critical` for the daemonset.
- Make exporter's monitor flags configurable.

#### Fixed
- Allow egress to port 1053 to make in-cluster DNS queries work.
- Allow egress to port 443 to allow accessing vault.



### cert-manager [2.13.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.13.0)

#### Changed
- Use retagged container image for HTTP01 AcmeSolver ([#212](https://github.com/giantswarm/cert-manager-app/pull/212))
- Pin kubectl to 1.23.3 in crd-install and clusterissuer-install jobs ([#216](https://github.com/giantswarm/cert-manager-app/pull/216))
- Add `application.giantswarm.io/team` to default labels ([#224](https://github.com/giantswarm/cert-manager-app/pull/224)).



### chart-operator [2.21.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.21.0)

#### Changed
- Update `helmclient` to v4.10.0.



### cluster-autoscaler [1.22.2-gs6](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs6)

#### Added
- Support cloud provider alias names (GCP -> GCE).

#### Fixed
- Updated to correct cluster-autoscaler version.
- Use GS-built 1.22 image to deliver upstream [unreleased fix](https://github.com/kubernetes/autoscaler/pull/4600).



### coredns [1.9.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.9.0)

#### Added
- Add toleration for `node.cloudprovider.kubernetes.io/uninitialized`.
#### Changed
- Update `coredns` to upstream version [1.8.7](https://coredns.io/2021/12/09/coredns-1.8.7-release/).
 



### external-dns [2.9.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.9.1)

#### Changed
- Allow setting the AWS default region (`aws.region`) indepentent from any other value.



### kiam-watchdog [0.7.0](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.7.0)

#### Added
- Add PriorityClassName.



### kubernetes [1.22.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.9)

#### Bug or Regression

- Fixed a regression that could incorrectly reject pods with OutOfCpu errors if they were rapidly scheduled after other pods were reported as complete in the API. The Kubelet now waits to report the phase of a pod as terminal in the API until all running containers are guaranteed to have stopped and no new containers can be started.  Short-lived pods may take slightly longer (~1s) to report Succeeded or Failed after this change. ([#108749](https://github.com/kubernetes/kubernetes/pull/108749), [@bobbypage](https://github.com/bobbypage)) [SIG Apps, Node and Testing]
- Fixes error handling in a kubectl method used in downstream packages. ([#108520](https://github.com/kubernetes/kubernetes/pull/108520), [@heybronson](https://github.com/heybronson)) [SIG CLI]

#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### kube-state-metrics [1.10.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.10.0)

#### Changed
- Make `--metric-labels-allowlist` configurable through user values.
- Add Node Pool labels to the default allowed labels in `--metric-labels-allowlist`.



### net-exporter [1.12.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.12.0)

#### Changed
- Use parameter for CoreDNS namespace (defaulted to kube-system)



### vertical-pod-autoscaler [2.1.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.1.2)

#### Fixed
- Fixed default value for admission controller PDB.



### vertical-pod-autoscaler-crd [1.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v1.0.1)

#### Added
- Add cluster singleton restriction so app can only be installed once.



