## :zap:  Giant Swarm Release 12.0.0 for Azure :zap:

This is the first release to support kubernetes 1.17 on Azure.

**Important Warning** During master upgrade from 11.4.0 to 12.0.0, within the time frame of 30 second we had noticed a spike in requests failures. This most likely is caused by Azure CNI upgrade and despite a lot of effort we had not found a solution to maintain upgrade path and avoid this disturbance. Please keep this in mind when scheduling an upgrade window, and contact your SE if you have further questions.

As of this release **NGINX Ingress Controller App** is **optional and not pre-installed** component on Azure.
This allows NGINX App installations to be managed independently from the base platform lifecycle. It is both benefit but also new responsibility to keep NGINX App installations up-to-date separately from rest of the cluster.
Making NGINX optional enables use of other ingress controller alternatives without wasting resources where NGINX is not the preferred option.
Upgrading existing tenant clusters with pre-installed NGINX will leave NGINX unchanged. Existing NGINX App custom resources will still have `giantswarm.io/managed-by: cluster-operator` label, but it will be ignored. The label will be cleaned up at a later point after all tenant clusters have been upgraded and Azure platform releases older than v12.0.0 archived.

Apart from that, this release contains many changes in other components, including important security fixes in Kubernetes and Calico.

This release also brings technical changes that are needed to prepare the upcoming Cluster API release with Node Pools support.

### Kubernetes v1.17.9

- **Cloud Provider Labels reach General Availability**: Added as a beta feature way back in v1.2, v1.17 sees the general availability of cloud provider labels.
- **Volume Snapshot Moves to Beta**: The Kubernetes Volume Snapshot feature is now beta in Kubernetes v1.17. It was introduced as alpha in Kubernetes v1.12, with a second alpha with breaking changes in Kubernetes v1.13.
- **CSI Migration Beta**: The Kubernetes in-tree storage plugin to Container Storage Interface (CSI) migration infrastructure is now beta in Kubernetes v1.17. CSI migration was introduced as alpha in Kubernetes v1.14.
- Includes a fix for **CVE-2020-8559** that allowed a privilege escalation issue from a compromised node to the cluster.
- Updated from **v1.16.12** - to review all changes please read changelog since [v1.17.0](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changes).

### azure-operator [v4.2.0](https://github.com/giantswarm/azure-operator/blob/v4.2.0/CHANGELOG.md#420---2020-07-08)

- Changed how the Azure authentication works when connecting to a different Subscription than the Control Plane's one so we can move forward with replacing VPN gateway with VNET peering for connecting customer clusters with the Giant Swarm Control Plane to increase performance and lower the costs.
- Restricted storage account access to the local VNET only to increase security.
- The Azure MSI extension for linux was not used so it will not be deployed anymore.

### Calico [v3.15.1](https://docs.projectcalico.org/archive/v3.15/release-notes/#v3151)

- Includes a fix for [**CVE-2020-13597**](https://cve.mitre.org/cgi-bin/cvename.cgi?name=2020-13597) that allows a Pod to be
compromised and redirect full or partial network traffic from the Node to the compromised Pod.
- Updated from **3.10.1** - to review all changes please read changelogs since
 [v3.15.0](https://docs.projectcalico.org/archive/v3.15/release-notes/#v3150),
 [v3.14](https://docs.projectcalico.org/archive/v3.14/release-notes),
 [v3.13](https://docs.projectcalico.org/archive/v3.13/release-notes),
 [v3.12](https://docs.projectcalico.org/archive/v3.12/release-notes),
 [v3.11](https://docs.projectcalico.org/archive/v3.11/release-notes), and
 [v3.10](https://docs.projectcalico.org/archive/v3.10/release-notes).

### CoreDNS v1.6.5 ([Giant Swarm app v1.2.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v120-2020-07-13))

- Apply a readiness probe to improve reliability of the clusters.
- Increase the liveness probe failure threshold from 5 failures to 7 failures.

### cluster-operator 0.23.12

- Support for making NGINX IC App optional and not pre-installed.

### external-dns v0.7.2 [Giant Swarm app 1.2.2](https://github.com/giantswarm/external-dns-app/releases/tag/v1.2.2)

- Updated from v0.5.18 - to review all changes please read changelog since [v0.5.18](https://github.com/kubernetes-sigs/external-dns/blob/master/CHANGELOG.md#v060---2020-02-11).
