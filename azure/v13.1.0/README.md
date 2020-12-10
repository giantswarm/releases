# :zap: Giant Swarm Release v13.1.0 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.1.0](https://github.com/giantswarm/aws-operator/releases/tag/v5.1.0)

Not found


### containerlinux [2605.9.0](https://www.flatcar-linux.org/releases/#release-2605.9.0)

Security fixes:

*   containerd ([CVE-2020-15257](https://nvd.nist.gov/vuln/detail/CVE-2020-15257))
*   glibc ([CVE-2019-9169](https://nvd.nist.gov/vuln/detail/CVE-2019-9169),[ CVE-2019-6488](https://nvd.nist.gov/vuln/detail/CVE-2019-6488),[ CVE-2019-7309](https://nvd.nist.gov/vuln/detail/CVE-2019-7309),[ CVE-2020-10029](https://nvd.nist.gov/vuln/detail/CVE-2020-10029),[ CVE-2020-1751](https://nvd.nist.gov/vuln/detail/CVE-2020-1751),[ CVE-2020-6096](https://nvd.nist.gov/vuln/detail/CVE-2020-6096),[ CVE-2018-20796](https://nvd.nist.gov/vuln/detail/CVE-2018-20796))
*   Linux ([CVE-2020-28941](https://nvd.nist.gov/vuln/detail/CVE-2020-28941), [CVE-2020-4788](https://nvd.nist.gov/vuln/detail/CVE-2020-4788), [CVE-2020-25669](https://nvd.nist.gov/vuln/detail/CVE-2020-25669), [CVE-2020-14351](https://nvd.nist.gov/vuln/detail/CVE-2020-14351))
*   glib ([CVE-2019-12450](https://nvd.nist.gov/vuln/detail/CVE-2019-12450))
*   open-iscsi ([CVE-2017-17840](https://nvd.nist.gov/vuln/detail/CVE-2017-17840))
*   samba ([CVE-2019-10197](https://nvd.nist.gov/vuln/detail/CVE-2019-10197),[ CVE-2020-10704](https://nvd.nist.gov/vuln/detail/CVE-2020-10704),[ CVE-2020-10745](https://nvd.nist.gov/vuln/detail/CVE-2020-10745),[ CVE-2019-3880](https://nvd.nist.gov/vuln/detail/CVE-2019-3880),[ CVE-2019-10218](https://nvd.nist.gov/vuln/detail/CVE-2019-10218))
*   shadow ([CVE-2019-19882](https://nvd.nist.gov/vuln/detail/CVE-2019-19882))
*   sssd ([CVE-2018-16883](https://nvd.nist.gov/vuln/detail/CVE-2018-16883),[ CVE-2019-3811](https://nvd.nist.gov/vuln/detail/CVE-2019-3811),[ CVE-2018-16838](https://nvd.nist.gov/vuln/detail/CVE-2018-16838))
*   trousers ([CVE-2020-24330](https://nvd.nist.gov/vuln/detail/CVE-2020-24330),[ CVE-2020-24331](https://nvd.nist.gov/vuln/detail/CVE-2020-24331))
*   cifs-utils ([CVE-2020-14342](https://nvd.nist.gov/vuln/detail/CVE-2020-14342))
*   ntp ([CVE-2020-11868](https://nvd.nist.gov/vuln/detail/CVE-2020-11868),[ CVE-2020-13817](https://nvd.nist.gov/vuln/detail/CVE-2020-13817),[ CVE-2018-8956](https://nvd.nist.gov/vuln/detail/CVE-2018-8956),[ CVE-2020-15025](https://nvd.nist.gov/vuln/detail/CVE-2020-15025))
*   bzip2 ([CVE-2019-12900](https://nvd.nist.gov/vuln/detail/CVE-2019-12900))
*   c-ares ([CVE-2017-1000381](https://nvd.nist.gov/vuln/detail/CVE-2017-1000381))
*   file ([CVE-2019-18218](https://nvd.nist.gov/vuln/detail/CVE-2019-18218))
*   json-c ([CVE-2020-12762](https://nvd.nist.gov/vuln/detail/CVE-2020-12762))
*   jq ([CVE-2015-8863](https://nvd.nist.gov/vuln/detail/CVE-2015-8863), [CVE-2016-4074](https://nvd.nist.gov/vuln/detail/CVE-2016-4074))
*   libuv ([CVE-2020-8252](https://nvd.nist.gov/vuln/detail/CVE-2020-8252))
*   libxml2 ([CVE-2019-20388](https://nvd.nist.gov/vuln/detail/CVE-2019-20388), [CVE-2020-7595](https://nvd.nist.gov/vuln/detail/CVE-2020-7595))
*   re2c ([CVE-2020-11958](https://nvd.nist.gov/vuln/detail/CVE-2020-11958))
*   tar ([CVE-2019-9923](https://nvd.nist.gov/vuln/detail/CVE-2019-9923))
*   sqlite ([CVE-2020-11656](https://nvd.nist.gov/vuln/detail/CVE-2020-11656), [CVE-2020-9327](https://nvd.nist.gov/vuln/detail/CVE-2020-9327), [CVE-2020-11655](https://nvd.nist.gov/vuln/detail/CVE-2020-11655), [CVE-2020-13630](https://nvd.nist.gov/vuln/detail/CVE-2020-13630), [CVE-2020-13435](https://nvd.nist.gov/vuln/detail/CVE-2020-13435), [CVE-2020-13434](https://nvd.nist.gov/vuln/detail/CVE-2020-13434), [CVE-2020-13631](https://nvd.nist.gov/vuln/detail/CVE-2020-13631), [CVE-2020-13632](https://nvd.nist.gov/vuln/detail/CVE-2020-13632), [CVE-2020-15358](https://nvd.nist.gov/vuln/detail/CVE-2020-15358))
*   tcpdump and pcap ([CVE-2018-10103](https://nvd.nist.gov/vuln/detail/CVE-2018-10103), [CVE-2018-10105](https://nvd.nist.gov/vuln/detail/CVE-2018-10105), [CVE-2019-15163](https://nvd.nist.gov/vuln/detail/CVE-2019-15163), [CVE-2018-14461](https://nvd.nist.gov/vuln/detail/CVE-2018-14461), [CVE-2018-14462](https://nvd.nist.gov/vuln/detail/CVE-2018-14462), [CVE-2018-14463](https://nvd.nist.gov/vuln/detail/CVE-2018-14463), [CVE-2018-14464](https://nvd.nist.gov/vuln/detail/CVE-2018-14464), [CVE-2018-14465](https://nvd.nist.gov/vuln/detail/CVE-2018-14465), [CVE-2018-14466](https://nvd.nist.gov/vuln/detail/CVE-2018-14466), [CVE-2018-14467](https://nvd.nist.gov/vuln/detail/CVE-2018-14467), [CVE-2018-14468](https://nvd.nist.gov/vuln/detail/CVE-2018-14468), [CVE-2018-14469](https://nvd.nist.gov/vuln/detail/CVE-2018-14469), [CVE-2018-14470](https://nvd.nist.gov/vuln/detail/CVE-2018-14470), [CVE-2018-14880](https://nvd.nist.gov/vuln/detail/CVE-2018-14880), [CVE-2018-14881](https://nvd.nist.gov/vuln/detail/CVE-2018-14881), [CVE-2018-14882](https://nvd.nist.gov/vuln/detail/CVE-2018-14882), [CVE-2018-16227](https://nvd.nist.gov/vuln/detail/CVE-2018-16227), [CVE-2018-16228](https://nvd.nist.gov/vuln/detail/CVE-2018-16228), [CVE-2018-16229](https://nvd.nist.gov/vuln/detail/CVE-2018-16229), [CVE-2018-16230](https://nvd.nist.gov/vuln/detail/CVE-2018-16230), [CVE-2018-16300](https://nvd.nist.gov/vuln/detail/CVE-2018-16300), [CVE-2018-16451](https://nvd.nist.gov/vuln/detail/CVE-2018-16451), [CVE-2018-16452](https://nvd.nist.gov/vuln/detail/CVE-2018-16452), [CVE-2019-15166](https://nvd.nist.gov/vuln/detail/CVE-2019-15166), [CVE-2018-14879](https://nvd.nist.gov/vuln/detail/CVE-2018-14879), [CVE-2017-16808](https://nvd.nist.gov/vuln/detail/CVE-2017-16808), [CVE-2018-19519](https://nvd.nist.gov/vuln/detail/CVE-2018-19519), [CVE-2019-15161](https://nvd.nist.gov/vuln/detail/CVE-2019-15161), [CVE-2019-15165](https://nvd.nist.gov/vuln/detail/CVE-2019-15165), [CVE-2019-15164](https://nvd.nist.gov/vuln/detail/CVE-2019-15164), [CVE-2019-1010220](https://nvd.nist.gov/vuln/detail/CVE-2019-1010220))
*   libbsd ([CVE-2019-20367](https://nvd.nist.gov/vuln/detail/CVE-2019-20367))
*   rsync and zlib ([CVE-2016-9840](https://nvd.nist.gov/vuln/detail/CVE-2016-9840), [CVE-2016-9841](https://nvd.nist.gov/vuln/detail/CVE-2016-9841), [CVE-2016-9842](https://nvd.nist.gov/vuln/detail/CVE-2016-9842), [CVE-2016-9843](https://nvd.nist.gov/vuln/detail/CVE-2016-9843))

Bug fixes

*   Added systemd-tmpfiles directives for /opt and /opt/bin to ensure that the folders have correct permissions even when /opt/ was once created by containerd ([Flatcar#279](https://github.com/kinvolk/Flatcar/issues/279))
*   Make the automatic filesystem resizing more robust against a race and add more logging ([kinvolk/init#31](https://github.com/kinvolk/init/pull/31))
*   Allow inactive network interfaces to be bound to a bonding interface, by encoding additional configuration for systemd-networkd-wait-online ([afterburn PR #10](https://github.com/flatcar-linux/afterburn/pull/10))
*   Do not configure ccache in Jenkins ([scripts PR #100](https://github.com/flatcar-linux/scripts/pull/100))
*   Azure: Exclude bonded SR-IOV network interfaces with newer drivers from networkd (in addition to the old drivers) to prevent them being configured instead of just the bond interface ([init PR#29](https://github.com/flatcar-linux/init/pull/29),[ bootengine PR#19](https://github.com/flatcar-linux/bootengine/pull/19))

Changes:

*   Update-engine now detects rollbacks and reports them as errors to the update server ([PR#6](https://github.com/flatcar-linux/update_engine/pull/6))
*   The zstd tools were added (version 1.4.4)
*   The kernel config CONFIG_PSI was set to support[ Pressure Stall Information](https://www.kernel.org/doc/html/latest/accounting/psi.html), more information also under[ https://facebookmicrosites.github.io/psi/docs/overview](https://facebookmicrosites.github.io/psi/docs/overview) ([Flatcar#162](https://github.com/flatcar-linux/Flatcar/issues/162))
*   The kernel config CONFIG_BPF_JIT_ALWAYS_ON was set to use the BPF just-in-time compiler by default for faster execution
*   The kernel config CONFIG_POWER_SUPPLY was set
*   The kernel configs CONFIG_OVERLAY_FS_METACOPY and CONFIG_OVERLAY_FS_REDIRECT_DIR were set. With the first overlayfs will only copy up metadata when a metadata-specific operation like chown/chmod is performed. The full file will be copied up later when the file is opened for write operations. With the second, which is equivalent to setting "redirect_dir=on" in the kernel command-line, overlayfs will copy up the directory first before the actual content ([Flatcar#170](https://github.com/kinvolk/Flatcar/issues/170)).
*   Remove unnecessary kernel module nf-conntrack-ipv4 ([overlay PR#649](https://github.com/flatcar-linux/coreos-overlay/pull/649))
*   Compress kernel modules with xz ([overlay PR#628](https://github.com/flatcar-linux/coreos-overlay/pull/628))
*   Add containerd-runc-shim-v* binaries required by kubelet custom CRI endpoints ([overlay PR#623](https://github.com/flatcar-linux/coreos-overlay/pull/623))
*   Equinix Metal (Packet): Exclude unused network interfaces from networkd, disregard the state of the bonded interfaces for the network-online.target and only require the bond interface itself to have at least one active link instead of routable which requires both links to be active ([afterburn PR#10](https://github.com/flatcar-linux/afterburn/pull/10))
*   QEMU: Use flatcar.autologin kernel command line parameter for auto login on the console ([Flatcar #71](https://github.com/flatcar-linux/Flatcar/issues/71))

Updates:

*   Linux ([5.4.81](https://lwn.net/Articles/838790/))
*   Linux firmware ([20200918](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20200918))
*   systemd ([246.6](https://github.com/systemd/systemd-stable/releases/tag/v246.6))
*   glibc ([2.32](https://lwn.net/Articles/828210/))
*   Docker ([19.03.14](https://github.com/docker/docker-ce/releases/tag/v19.03.14))
*   containerd ([1.4.3](https://github.com/containerd/containerd/releases/tag/v1.4.3))
*   tini[ (0.18](https://github.com/krallin/tini/releases/tag/v0.18.0))
*   libseccomp[ (2.5.0](https://github.com/seccomp/libseccomp/releases/tag/v2.5.0))
*   audit[ (2.8.5](https://github.com/linux-audit/audit-userspace/releases/tag/v2.8.5))
*   bzip2 ([1.0.8](https://sourceware.org/git/?p=bzip2.git;a=blob;f=CHANGES;h=30afead2586b6d64f50988a41d394a0131b38949;hb=HEAD#l342))
*   c-ares[ (1.61.1](https://github.com/c-ares/c-ares/releases/tag/cares-1_16_1))
*   cryptsetup[ (2.3.2](https://gitlab.com/cryptsetup/cryptsetup/-/tags/v2.3.2))
*   cifs-utils (6.11)
*   dbus-glib (0.110)
*   dracut[ (050](https://github.com/dracutdevs/dracut/releases/tag/050))
*   elfutils (0.178)
*   glib (2.64.5)
*   json-c[ (0.15](https://github.com/json-c/json-c/releases/tag/json-c-0.15-20200726))
*   jq ([1.6](https://github.com/stedolan/jq/releases/tag/jq-1.6))
*   libuv[ (1.39.0](https://github.com/libuv/libuv/releases/tag/v1.39.0))
*   libxml2[ (2.9.10](https://gitlab.gnome.org/GNOME/libxml2/-/tags/v2.9.10))
*   ntp (4.2.8_p15)
*   open-iscsi (2.1.2)
*   samba (4.11.13)
*   shadow (4.8)
*   sssd (2.3.1)
*   strace (5.9)
*   talloc (2.3.1)
*   tar[ (1.32](https://git.savannah.gnu.org/cgit/tar.git/tag/?h=release_1_32))
*   tdb (1.4.3)
*   tevent (0.10.2)
*   SDK/developer container: GCC (9.3.0), binutils (2.35), gdb (9.2)
*   Go ([1.15.5](https://go.googlesource.com/go/+/refs/tags/go1.15.5), [1.12.17](https://go.googlesource.com/go/+/refs/tags/go1.12.17)) (only in SDK)
*   Rust ([1.46.0](https://blog.rust-lang.org/2020/08/27/Rust-1.46.0.html)) (only in SDK)
*   file ([5.39](https://github.com/file/file/tree/FILE5_39)) (only in SDK)
*   gdbus-codegen ([2.64.5](https://gitlab.gnome.org/GNOME/glib/-/tags/2.64.5)) (only in SDK)
*   meson ([0.55.3](https://github.com/mesonbuild/meson/releases/tag/0.55.3)) (only in SDK)
*   re2c ([2.0.3](https://re2c.org/releases/release_notes.html#release-2-0-3)) (only in SDK)
*   VMware: open-vm-tools (11.2.0)


### kubernetes [1.18.13](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.13)

#### Feature
- Add a new flag to set priority for the kubelet on Windows nodes so that workloads cannot overwhelm the node there by disrupting kubelet process. ([#96158](https://github.com/kubernetes/kubernetes/pull/96158), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node]
#### Bug or Regression
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
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### cert-exporter [1.4.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.4.0)

#### Added
- Add new metric (`cert_exporter_secret_not_after`) which tracks expiry of TLS certificates stored in Kubernetes secrets. ([#92](https://github.com/giantswarm/cert-exporter/pull/92))



### chart-operator [2.5.2](https://github.com/giantswarm/chart-operator/releases/tag/v2.5.2)

#### Added
- Add Vertical Pod Autoscaler support.



### metrics-server [1.1.1](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.1.1)

#### Changed
- Updated metrics-server version to 0.3.6.
- Updated architect-orb to 0.10.0.



