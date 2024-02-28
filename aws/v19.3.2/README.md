# :zap: Giant Swarm Release v19.3.2 for AWS :zap:

This patch release introduces latest stable Flatcar release that includes important CVEs, especially the Leaky Vessels for the containerd version [CVE-2024-21626](https://nvd.nist.gov/vuln/detail/CVE-2024-21626). We have also included our latest changes to the `security-bundle`.

## Change details


### containerlinux [3815.2.0](https://www.flatcar-linux.org/releases/#release-3815.2.0)

 _Changes since **Stable 3760.2.0**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-46838](https://nvd.nist.gov/vuln/detail/CVE-2023-46838), [CVE-2023-50431](https://nvd.nist.gov/vuln/detail/CVE-2023-50431), [CVE-2023-6610](https://nvd.nist.gov/vuln/detail/CVE-2023-6610), [CVE-2023-6915](https://nvd.nist.gov/vuln/detail/CVE-2023-6915), [CVE-2024-1085](https://nvd.nist.gov/vuln/detail/CVE-2024-1085), [CVE-2024-1086](https://nvd.nist.gov/vuln/detail/CVE-2024-1086), [CVE-2024-23849](https://nvd.nist.gov/vuln/detail/CVE-2024-23849))
 - Go ([CVE-2023-39326](https://nvd.nist.gov/vuln/detail/CVE-2023-39326), [CVE-2023-45285](https://nvd.nist.gov/vuln/detail/CVE-2023-45285))
 - VMWare: open-vm-tools ([CVE-2023-34058](https://nvd.nist.gov/vuln/detail/CVE-2023-34058), [CVE-2023-34059](https://nvd.nist.gov/vuln/detail/CVE-2023-34059))
 - docker ([CVE-2024-24557](https://nvd.nist.gov/vuln/detail/CVE-2024-24557))
 - nghttp2 ([CVE-2023-44487](https://nvd.nist.gov/vuln/detail/CVE-2023-44487))
 - runc ([CVE-2024-21626](https://nvd.nist.gov/vuln/detail/CVE-2024-21626))
 - samba ([CVE-2023-4091](https://nvd.nist.gov/vuln/detail/CVE-2023-4091))
 - zlib ([CVE-2023-45853](https://nvd.nist.gov/vuln/detail/CVE-2023-45853))
 
 #### Bug fixes:
 
 - Added a workaround for old airgapped/proxied update-engine clients to be able to update to this release ([Flatcar#1332](https://github.com/flatcar/Flatcar/issues/1332), [update_engine#38](https://github.com/flatcar/update_engine/pull/38))
 - Forwarded the proxy environment variables of `update-engine.service` to the postinstall script to support fetching OEM systemd-sysext payloads through a proxy ([Flatcar#1326](https://github.com/flatcar/Flatcar/issues/1326))
 - Set TTY used for fetching server_context to RAW mode before running cloudinit on cloudsigma ([scripts#1280](https://github.com/flatcar/scripts/pull/1280))
 
 #### Changes:
 
 - **torcx was replaced by systemd-sysext in the OS image**. Learn more about sysext and how to customise OS images [here](https://www.flatcar.org/docs/latest/provisioning/sysext/).
     (which is now also a legacy option because systemd-sysext offers a more robust and better structured way of customisation, including OS independent updates).
   - Torcx entered deprecation 2 years ago in favour of [deploying plain Docker binaries](https://www.flatcar.org/docs/latest/container-runtimes/use-a-custom-docker-or-containerd-version/)
   - Torcx has been removed entirely; if you use torcx to extend the Flatcar base OS image, please refer to our [conversion script](https://www.flatcar.org/docs/latest/provisioning/sysext/#torcx-deprecation) and to the sysext documentation mentioned above for migrating.
   - Consequently, `update_engine` will not perform torcx sanity checks post-update anymore.
   - Relevant changes: [scripts#1216](https://github.com/flatcar/scripts/pull/1216), [update_engine#30](https://github.com/flatcar/update_engine/pull/30), [Mantle#466](https://github.com/flatcar/mantle/pull/466), [Mantle#465](https://github.com/flatcar/mantle/pull/465).
 - **NOTE:** The docker btrfs storage driver has been de-prioritised; BTRFS backed storage will now default to the `overlay2` driver
     ([changelog](https://docs.docker.com/engine/release-notes/23.0/#bug-fixes-and-enhancements-6), [upstream pr](https://github.com/moby/moby/pull/42661)).
 - **NOTE:** If you are already using btrfs-backed Docker storage and are upgrading to this new version, Docker will automatically use the `btrfs` storage driver for backwards-compatibility with your deployment.
  - **Docker will remove the `btrfs` driver entirely in a future version. Please consider migrating your deployments to the `overlay2` driver.**
     Using the btrfs driver can still be enforced by creating a respective [docker config](https://docs.docker.com/storage/storagedriver/btrfs-driver/#configure-docker-to-use-the-btrfs-storage-driver) at `/etc/docker/daemon.json`.
 - cri-tools, runc, containerd, docker, and docker-cli are now built from Gentoo upstream ebuilds. Docker received a major version upgrade - it was updated to Docker 24 (from Docker 20; see "updates").
 - GCP OEM images now use a systemd-sysext image for layering additional platform-specific software on top of `/usr` and being part of the OEM A/B updates ([flatcar#1146](https://github.com/flatcar/Flatcar/issues/1146))
 - Added a `flatcar-update --oem-payloads <yes|no>` flag to skip providing OEM payloads, e.g., for downgrades ([init#114](https://github.com/flatcar/init/pull/114))
 
 #### Updates:
- Linux ([6.1.77](https://lwn.net/Articles/961012) (includes [6.1.76](https://lwn.net/Articles/960442), [6.1.75](https://lwn.net/Articles/959513), [6.1.74](https://lwn.net/Articles/958863)))
- Linux Firmware ([20231111](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20231111) (includes [20231030](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20231030)))
- Go ([1.20.12](https://go.dev/doc/devel/release#go1.20.12))
- Azure: WALinuxAgent ([v2.9.1.1](https://github.com/Azure/WALinuxAgent/releases/tag/v2.9.1.1))
- DEV: Azure ([3.11.6](https://docs.python.org/release/3.11.6/whatsnew/changelog.html#python-3-11-6))
- DEV: iperf ([3.15](https://github.com/esnet/iperf/releases/tag/3.15))
- DEV: smartmontools ([7.4](https://www.smartmontools.org/browser/tags/RELEASE_7_4/smartmontools/NEWS))
- SDK: Rust ([1.73.0](https://github.com/rust-lang/rust/releases/tag/1.73.0))
- SDK: Python ([3.11.0](https://github.com/platformdirs/platformdirs/releases/tag/3.11.0) (includes [23.2](https://github.com/pypa/packaging/releases/tag/23.2)))
- VMWare: open-vm-tools ([12.3.5](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.3.5))
- acpid ([2.0.34](https://sourceforge.net/p/acpid2/code/ci/2.0.34/tree/Changelog))
- ca-certificates ([3.97](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_97.html))
- containerd ([1.7.9](https://github.com/containerd/containerd/releases/tag/v1.7.9) (includes [1.7.8](https://github.com/containerd/containerd/releases/tag/v1.7.8), [1.7.13](https://github.com/containerd/containerd/releases/tag/v1.7.13), [1.7.10](https://github.com/containerd/containerd/releases/tag/v1.7.10)))
- cri-tools ([1.27.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.27.0))
- ding-libs ([0.6.2](https://github.com/SSSD/ding-libs/releases/tag/0.6.2))
- docker ([24.0.9](https://github.com/moby/moby/releases/tag/v24.0.9) (includes [24.0.6](https://docs.docker.com/engine/release-notes/24.0/), [23.0](https://docs.docker.com/engine/release-notes/23.0/)))
- efibootmgr ([18](https://github.com/rhboot/efibootmgr/releases/tag/18))
- efivar ([38](https://github.com/rhboot/efivar/releases/tag/38))
- ethtool ([6.5](https://git.kernel.org/pub/scm/network/ethtool/ethtool.git/tree/NEWS?h=v6.5))
- hwdata ([v0.375](https://github.com/vcrhonek/hwdata/releases/tag/v0.375) (includes [0.374](https://github.com/vcrhonek/hwdata/commits/v0.374)))
- iproute2 ([6.5.0](https://marc.info/?l=linux-netdev&m=169401822317373&w=2))
- ipvsadm ([1.31](https://git.kernel.org/pub/scm/utils/kernel/ipvsadm/ipvsadm.git/tag/?h=v1.31) (includes [1.30](https://git.kernel.org/pub/scm/utils/kernel/ipvsadm/ipvsadm.git/tag/?h=v1.30), [1.29](https://git.kernel.org/pub/scm/utils/kernel/ipvsadm/ipvsadm.git/tag/?h=v1.29), [1.28](https://git.kernel.org/pub/scm/utils/kernel/ipvsadm/ipvsadm.git/tag/?h=v1.28)))
- json-c ([0.17](https://github.com/json-c/json-c/blob/json-c-0.17-20230812/ChangeLog))
- libffi ([3.4.4](https://github.com/libffi/libffi/releases/tag/v3.4.4) (includes [3.4.3](https://github.com/libffi/libffi/releases/tag/v3.4.3), [3.4.2](https://github.com/libffi/libffi/releases/tag/v3.4.2)))
- liblinear ([246](https://github.com/cjlin1/liblinear/releases/tag/v246))
- libmnl ([1.0.5](https://git.netfilter.org/libmnl/log/?h=libmnl-1.0.5))
- libnetfilter_conntrack ([1.0.9](https://git.netfilter.org/libnetfilter_conntrack/log/?h=libnetfilter_conntrack-1.0.9))
- libnetfilter_cthelper ([1.0.1](https://git.netfilter.org/libnetfilter_cthelper/log/?id=8cee0347cc6969c39bb64000dfaa676a8f9e30f0))
- libnetfilter_cttimeout ([1.0.1](https://git.netfilter.org/libnetfilter_cttimeout/log/?id=068d36d6291f53a0a609ab1f695aa06e94ce3d30))
- libnfnetlink ([1.0.2](https://git.netfilter.org/libnfnetlink/log/?h=libnfnetlink-1.0.2))
- libsodium ([1.0.19](https://github.com/jedisct1/libsodium/releases/tag/1.0.19-RELEASE))
- libunistring ([1.1](https://git.savannah.gnu.org/gitweb/?p=libunistring.git;a=blob;f=NEWS;h=5a43ddd7011d62a952733f6c0b7ad52aa4f385c7;hb=8006860b710aae2e8442088c3ddc7d819dfa8ac7))
- libunwind ([1.7.2](https://github.com/libunwind/libunwind/releases/tag/v1.7.2) (includes [1.7.0](https://github.com/libunwind/libunwind/releases/tag/v1.7.0)))
- liburing ([2.3](https://github.com/axboe/liburing/blob/liburing-2.3/CHANGELOG))
- mpc ([1.3.1](https://sympa.inria.fr/sympa/arc/mpc-discuss/2022-12/msg00049.html) (includes [1.3.0](https://sympa.inria.fr/sympa/arc/mpc-discuss/2022-12/msg00028.html)))
- mpfr ([4.2.1](https://gitlab.inria.fr/mpfr/mpfr/-/blob/4.2.1/NEWS))
- nghttp2 ([1.57.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.57.0) (includes [1.56.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.56.0), [1.55.1](https://github.com/nghttp2/nghttp2/releases/tag/v1.55.1), [1.55.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.55.0), [1.54.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.54.0), [1.53.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.53.0), [1.52.0](https://github.com/nghttp2/nghttp2/releases/tag/v1.57.0)))
- nspr ([4.35](https://hg.mozilla.org/projects/nspr/log/b563bfc16c887c48b038b7b441fcc4e40a126d3b))
- ntp ([4.2.8p17](https://www.ntp.org/support/securitynotice/4_2_8p17-release-announcement/))
- nvme-cli ([v2.6](https://github.com/linux-nvme/nvme-cli/releases/tag/v2.6) (includes [v1.6](https://github.com/linux-nvme/libnvme/releases/tag/v1.6)))
- protobuf ([21.12](https://github.com/protocolbuffers/protobuf/releases/tag/v21.12) (includes [21.11](https://github.com/protocolbuffers/protobuf/releases/tag/v21.11), [21.10](https://github.com/protocolbuffers/protobuf/releases/tag/v21.10)))
- runc ([1.1.12](https://github.com/opencontainers/runc/releases/tag/v1.1.12))
- samba ([4.18.8](https://www.samba.org/samba/history/samba-4.18.8.html))
- sqlite ([3.43.2](https://www.sqlite.org/releaselog/3_43_2.html))
- squashfs-tools ([4.6.1](https://github.com/plougher/squashfs-tools/releases/tag/4.6.1) (includes [4.6](https://github.com/plougher/squashfs-tools/releases/tag/4.6)))
- thin-provisioning-tools ([1.0.6](https://github.com/jthornber/thin-provisioning-tools/blob/v1.0.6/CHANGES))

 _Changes since **Beta 3815.1.0**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-46838](https://nvd.nist.gov/vuln/detail/CVE-2023-46838), [CVE-2023-50431](https://nvd.nist.gov/vuln/detail/CVE-2023-50431), [CVE-2023-6610](https://nvd.nist.gov/vuln/detail/CVE-2023-6610), [CVE-2023-6915](https://nvd.nist.gov/vuln/detail/CVE-2023-6915), [CVE-2024-1085](https://nvd.nist.gov/vuln/detail/CVE-2024-1085), [CVE-2024-1086](https://nvd.nist.gov/vuln/detail/CVE-2024-1086), [CVE-2024-23849](https://nvd.nist.gov/vuln/detail/CVE-2024-23849))
 - docker ([CVE-2024-24557](https://nvd.nist.gov/vuln/detail/CVE-2024-24557))
 - runc ([CVE-2024-21626](https://nvd.nist.gov/vuln/detail/CVE-2024-21626))
 
 #### Bug fixes:
 
 - Added a workaround for old airgapped/proxied update-engine clients to be able to update to this release ([Flatcar#1332](https://github.com/flatcar/Flatcar/issues/1332), [update_engine#38](https://github.com/flatcar/update_engine/pull/38))
 - Forwarded the proxy environment variables of `update-engine.service` to the postinstall script to support fetching OEM systemd-sysext payloads through a proxy ([Flatcar#1326](https://github.com/flatcar/Flatcar/issues/1326))
 
 #### Changes:
 
 - Added a `flatcar-update --oem-payloads <yes|no>` flag to skip providing OEM payloads, e.g., for downgrades ([init#114](https://github.com/flatcar/init/pull/114))
 
 #### Updates:
 
 - Linux ([6.1.77](https://lwn.net/Articles/961012) (includes [6.1.76](https://lwn.net/Articles/960442), [6.1.75](https://lwn.net/Articles/959513), [6.1.74](https://lwn.net/Articles/958863)))
 - ca-certificates ([3.97](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_97.html))
 - containerd ([1.7.13](https://github.com/containerd/containerd/releases/tag/v1.7.13))
 - docker ([24.0.9](https://github.com/moby/moby/releases/tag/v24.0.9))
 - runc ([1.1.12](https://github.com/opencontainers/runc/releases/tag/v1.1.12))

### security-bundle [v1.6.2](https://github.com/giantswarm/security-bundle/releases/tag/v1.6.2)

#### Changes:

- Introduces new upstream Kyverno [version v1.11](https://kyverno.io/blog/2023/11/16/kyverno-1.11-released/).
- Introduces new `options` value to handle [App timeouts](https://docs.giantswarm.io/getting-started/app-platform/installation-configuration/).
- Introduces new `global.namespace` value to install all apps in a different namespace.

#### App updates

- kyverno [v0.17.6](https://github.com/giantswarm/kyverno-app/releases/tag/v0.17.6)
- kyverno-policy-operator [v0.0.7](https://github.com/giantswarm/kyverno-policy-operator/releases/tag/v0.0.7)
- exception-recommender [v0.1.1](https://github.com/giantswarm/exception-recommender/releases/tag/v0.1.1)
- trivy [v0.10.0](https://github.com/giantswarm/trivy-app/releases/tag/v0.10.0)
- trivy-operator [v0.7.2](https://github.com/giantswarm/trivy-operator-app/releases/tag/v0.7.2)
- starboard-exporter [v0.7.8](https://github.com/giantswarm/starboard-exporter/releases/tag/v0.7.8)
- falco [v0.8.0](https://github.com/giantswarm/falco-app/releases/tag/v0.8.0)


