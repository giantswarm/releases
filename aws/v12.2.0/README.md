# :zap: Giant Swarm Release v12.2.0 for AWS :zap:

This release includes a number of fixes, improvements, and picks up various component and app upgrades.

**If you are upgrading from 12.1.0 or older platform release, this release uses Cert Manager 0.16.1 and will no longer reconcile existing resources** due to changes in its API. Manual intervention is required to update affected resources. While the negative impact to your workloads is low-to-none, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

## Change details

### aws-operator [9.1.1](https://github.com/giantswarm/aws-operator/blob/v9.1.1/CHANGELOG.md#v911)

* Updated Flatcar Releases
* Updated AWS CNI installation manifests
* Updated k8scloudconfig to disable Calico installing the CNI binaries

### aws-cni [v1.7.3](https://github.com/aws/amazon-vpc-cni-k8s/blob/v1.7.3/CHANGELOG.md#v173)

* Bug - Avoid deleting ENIs being created by older CNI versions ([#1109](https://github.com/aws/amazon-vpc-cni-k8s/pull/1109))
* Bug - Wait for ENI and secondary IPs ([#1174](https://github.com/aws/amazon-vpc-cni-k8s/pull/1174))
* Improvement - Refresh subnet/CIDR information periodically ([#903](https://github.com/aws/amazon-vpc-cni-k8s/pull/903))

### cluster-operator [v3.2.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.2.0)

* Introducing Kubernetes events
* Add monitoring labels.

### calico v3.15.3

* Fix import from libcalico-go
* Fix import from libcalico-go
* Update pins - pick up FelixConfiguration
* Update pins - pick up FelixConfiguration

### kiam v3.6.0 (Giant Swarm app [v1.5.0](https://github.com/giantswarm/kiam-app/blob/v1.5.0/CHANGELOG.md))

* Setting gRPC environment variables
* Use deep liveness probe for kiam agent.
* Align charts with upstream.

### Etcd [3.4.13](https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.4.md)

* A log warning is added when etcd use any existing directory that has a permission different than 700 on Linux and 777 on Windows.

### Kubernetes [1.17.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#v11712)

* Fix a concurrent map writes error in kubelet ([#93773](https://github.com/kubernetes/kubernetes/pull/93773))
* Fixes a bug evicting pods after a taint with a limited tolerationSeconds toleration is removed from a node ([#93722](https://github.com/kubernetes/kubernetes/pull/93722))
* Fixes an issue that can result in namespaced custom resources being orphaned when their namespace is deleted, if the CRD defining the custom resource is removed concurrently with namespaces being deleted, then recreated. ([#93790](https://github.com/kubernetes/kubernetes/pull/93790))
* Kube-apiserver: fixed a bug returning inconsistent results from list requests which set a field or label selector and set a paging limit ([#94002](https://github.com/kubernetes/kubernetes/pull/94002))
* Update CNI plugins to v0.8.7 ([#94367](https://github.com/kubernetes/kubernetes/pull/94367))

Check the [upstream changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#v11712) for details on all changes.

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
