# :zap: Tenant Cluster Release v11.5.5 for AWS :zap:

This release upgrades:
* aws-cni to v1.7.3 to fix stability issues.
* containerlinux to v2512.5.0 to fix security issues.

## Change details

### aws-operator [8.7.8](https://github.com/giantswarm/aws-operator/blob/v8.7.8/CHANGELOG.md#v878)

* Updated Flatcar Releases
* Updated AWS CNI installation manifests
* Updated k8scloudconfig to disable Calico installing the CNI binaries

### aws-cni [v1.7.3](https://github.com/aws/amazon-vpc-cni-k8s/blob/v1.7.3/CHANGELOG.md#v173)

* Bug - Avoid deleting ENIs being created by older CNI versions ([#1109](https://github.com/aws/amazon-vpc-cni-k8s/pull/1109))
* Bug - Wait for ENI and secondary IPs ([#1174](https://github.com/aws/amazon-vpc-cni-k8s/pull/1174))
* Improvement - Refresh subnet/CIDR information periodically ([#903](https://github.com/aws/amazon-vpc-cni-k8s/pull/903))


### containerlinux [2512.5.0](https://www.flatcar-linux.org/releases/#release-2512.5.0)

Security fixes:

* Linux kernel: Fix AF_PACKET overflow in tpacket_rcv [CVE-2020-14386](https://seclists.org/oss-sec/2020/q3/146)
* Bind: fixes for [CVE-2020-8616](https://nvd.nist.gov/vuln/detail/CVE-2020-8616), [CVE-2020-8617](https://nvd.nist.gov/vuln/detail/CVE-2020-8617), [CVE-2020-8620](https://nvd.nist.gov/vuln/detail/CVE-2020-8620), [CVE-2020-8621](https://nvd.nist.gov/vuln/detail/CVE-2020-8621), [CVE-2020-8622](https://nvd.nist.gov/vuln/detail/CVE-2020-8622), [CVE-2020-8623](https://nvd.nist.gov/vuln/detail/CVE-2020-8623), [CVE-2020-8624](https://nvd.nist.gov/vuln/detail/CVE-2020-8624)

Bug fixes:

* The static IP address configuration in the initramfs works again in the format `ip=<ip>::<gateway>:<netmask>:<hostname>:<iface>:none[:<dns1>[:<dns2>]]` ([flatcar-linux/bootengine#15](https://github.com/flatcar-linux/bootengine/pull/15))
* app-admin/{kubelet, etcd, flannel}-wrapper: don't overwrite the user supplied –insecure-options argument ([flatcar-linux/coreos-overlay#426](https://github.com/flatcar-linux/coreos-overlay/pull/426))
* etcd-wrapper: Adjust data dir permissions ([flatcar-linux/coreos-overlay#536](https://github.com/flatcar-linux/coreos-overlay/pull/536))

Changes:

* Update public key to include a [new subkey](https://www.flatcar-linux.org/security/image-signing-key/)
* Vultr support in Ignition ([flatcar-linux/ignition#13](https://github.com/flatcar-linux/ignition/pull/13))
* VMware OVF settings default to ESXi 6.5 and Linux 3.x

Updates:

* Linux [4.19.145](https://lwn.net/Articles/831367/)
* bind-tools [9.11.22](https://ftp.isc.org/isc/bind9/cur/9.11/RELEASE-NOTES-bind-9.11.22.txt)
* etcd-wrapper [3.3.24](https://github.com/etcd-io/etcd/releases/tag/v3.3.24)
* Git [2.26.2](https://raw.githubusercontent.com/git/git/v2.26.2/Documentation/RelNotes/2.26.2.txt)
