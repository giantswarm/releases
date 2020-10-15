## :zap:  Giant Swarm Release 12.2.0-2020092800  for KVM :zap:

**If you are upgrading from 12.2.0, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades all Helm releases managed by Giant Swarm to use [Helm v3.3.4](https://github.com/helm/helm/releases/tag/v3.3.4).

This lets us benefit from the improved security model and keep up to date with the community. We also remove the Tiller deployment from the giantswarm namespace, removing its gRPC endpoint, which reduces operational complexity.

If you are still using Helm 2 then these Helm releases will not be affected. However we encourage you to upgrade to Helm 3. As Helm 2 support ends on November 13th 2020. https://helm.sh/blog/helm-v2-deprecation-timeline/

Below, you can find more details on components that were changed with this release.

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.2.x and above.

## Change details

### app-operator [v2.3.1](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#231---2020-09-22)
### chart-operator [v2.3.5](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#233---2020-09-29)
- Updated Helm to [v3.3.4](https://github.com/helm/helm/releases/tag/v3.3.4).


### containerlinux [2512.5.0](https://www.flatcar-linux.org/releases/#release-2512.5.0)

Security fixes:

* Linux kernel: Fix AF_PACKET overflow in tpacket_rcv [CVE-2020-14386](https://seclists.org/oss-sec/2020/q3/146)
* Bind: fixes for [CVE-2020-8616](https://nvd.nist.gov/vuln/detail/CVE-2020-8616), [CVE-2020-8617](https://nvd.nist.gov/vuln/detail/CVE-2020-8617), [CVE-2020-8620](https://nvd.nist.gov/vuln/detail/CVE-2020-8620), [CVE-2020-8621](https://nvd.nist.gov/vuln/detail/CVE-2020-8621), [CVE-2020-8622](https://nvd.nist.gov/vuln/detail/CVE-2020-8622), [CVE-2020-8623](https://nvd.nist.gov/vuln/detail/CVE-2020-8623), [CVE-2020-8624](https://nvd.nist.gov/vuln/detail/CVE-2020-8624)

Changes:
* Update public key to include a new subkey
* Vultr support in Ignition ([flatcar-linux/ignition#13](https://github.com/flatcar-linux/ignition/pull/13))
* VMware OVF settings default to ESXi 6.5 and Linux 3.x

Updates:
* Linux [4.19.145](https://lwn.net/Articles/831367/)
