# :zap: Giant Swarm Release v19.3.3 for AWS :zap:

This is a patch release that fixes an issue with clusters running in IMDSv2 mode only. If you are not using IMDSv2 only mode, you can skip this release.

As a reminder, enabling `IMDS v2` only support for the EC2 instances can be set via an annotation `alpha.aws.giantswarm.io/metadata-v2: required` on `AWSMachineDeployment` (for nodepools) and to `AWSControlPlane` (for control-plane nodes). The setting can be either configured before the upgrade or can be triggered to be effective with nodes rollout.

## Change details


### aws-operator [15.1.0](https://github.com/giantswarm/aws-operator/releases/tag/v15.1.0)

#### Fixed
- [Backport] Bump k8scc to 16.8.1 fix issues with IMDS v2.



### containerlinux [3815.2.1](https://www.flatcar-linux.org/releases/#release-3815.2.1)

_Changes since **Stable 3815.2.0**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-52429](https://nvd.nist.gov/vuln/detail/CVE-2023-52429), [CVE-2023-52434](https://nvd.nist.gov/vuln/detail/CVE-2023-52434), [CVE-2023-52435](https://nvd.nist.gov/vuln/detail/CVE-2023-52435), [CVE-2024-0340](https://nvd.nist.gov/vuln/detail/CVE-2024-0340), [CVE-2024-1151](https://nvd.nist.gov/vuln/detail/CVE-2024-1151), [CVE-2024-23850](https://nvd.nist.gov/vuln/detail/CVE-2024-23850), [CVE-2024-23851](https://nvd.nist.gov/vuln/detail/CVE-2024-23851), [CVE-2024-26582](https://nvd.nist.gov/vuln/detail/CVE-2024-26582), [CVE-2024-26583](https://nvd.nist.gov/vuln/detail/CVE-2024-26583), [CVE-2024-26586](https://nvd.nist.gov/vuln/detail/CVE-2024-26586), [CVE-2024-26593](https://nvd.nist.gov/vuln/detail/CVE-2024-26593))
 
 #### Bug fixes:
 
 - Fixed that systemd-sysext images can extend directories where Flatcar extensions are also shipping files, e.g., that the sysext-bakery Kubernetes extension works when OEM extensions are present ([sysext-bakery#50](https://github.com/flatcar/sysext-bakery/issues/50))
 - Fixed the handling of OEM update payloads in a Nebraska response with self-hosted packages in an airgapped environment ([update_engine#39](https://github.com/flatcar/update_engine/pull/39))
 - Restored support for custom OEMs supplied in the PXE boot where `/usr/share/oem` brings the OEM partition contents ([Flatcar#1376](https://github.com/flatcar/Flatcar/issues/1376))
 
 #### Changes:
 
 
 #### Updates:
 
 - Linux ([6.1.81](https://lwn.net/Articles/964562) (includes [6.1.80](https://lwn.net/Articles/964174), [6.1.79](https://lwn.net/Articles/963358), [6.1.78](https://lwn.net/Articles/962559)))
 - ca-certificates ([3.98](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_98.html))
 - keyutils ([1.6.3](https://git.kernel.org/pub/scm/linux/kernel/git/dhowells/keyutils.git/commit/?id=cb3bb194cca88211cbfcdde2f10c0f43c3fb8ec3) (includes [1.6.2](https://git.kernel.org/pub/scm/linux/kernel/git/dhowells/keyutils.git/commit/?id=454f80f537e5d1aad506599b6776e4cc1cf5f0f2)))


