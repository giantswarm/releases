## :zap:  Giant Swarm Release 12.0.0 for Azure :zap:

This is the first release to support kubernetes 1.17 on Azure.

Apart from this important change, there have been a high number of changes under the hood, including a few important
security fixes coming from upstream projects we upgraded.

This release also brings some technical changes that are needed to prepare the upcoming Cluster API release with
Node Pools support.

### Kubernetes v1.17.8

- Updated from v1.16.12 - 
changelog since [v1.17.7](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1177),
since [v1.17.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1176),
since [v1.17.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1175),
since [v1.17.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1174),
since [v1.17.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1173),
since [v1.17.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1172),
since [v1.17.1](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1171) and
since [v1.17.0](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changes)

### azure-operator [v4.2.0](https://github.com/giantswarm/azure-operator/blob/v4.2.0/CHANGELOG.md#420---2020-07-08)

- Changed how the Azure authentication works when connecting to a different Subscription than the Control Plane's one.
- Restricted storage account access to the local VNET only.
- The Azure MSI extension for linux is not deployed anymore.

### Calico v3.14.1 ([Giant Swarm app v1.1.10](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v1110-2020-06-29))

- Updated from 3.10.1 - 
changelog since [v3.14.0](https://docs.projectcalico.org/archive/v3.14/release-notes/#v3140),
changelog since [v3.13.4](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3134),
changelog since [v3.13.3](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3133),
changelog since [v3.13.2](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3132),
changelog since [v3.13.1](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3131),
changelog since [v3.13.0](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3130),
changelog since [v3.13.0](https://docs.projectcalico.org/archive/v3.13/release-notes/#v3130),
changelog since [v3.12.2](https://docs.projectcalico.org/archive/v3.12/release-notes/#v3122),
changelog since [v3.12.1](https://docs.projectcalico.org/archive/v3.12/release-notes/#v3121),
changelog since [v3.12.0](https://docs.projectcalico.org/archive/v3.12/release-notes/#v3120),
changelog since [v3.11.3](https://docs.projectcalico.org/archive/v3.11/release-notes/#v3113),
changelog since [v3.11.2](https://docs.projectcalico.org/archive/v3.11/release-notes/#v3112),
changelog since [v3.11.1](https://docs.projectcalico.org/archive/v3.11/release-notes/#v3111),
changelog since [v3.11.0](https://docs.projectcalico.org/archive/v3.11/release-notes/#v3110),
changelog since [v3.10.4](https://docs.projectcalico.org/archive/v3.10/release-notes/#v3104),
changelog since [v3.10.3](https://docs.projectcalico.org/archive/v3.10/release-notes/#v3103),
changelog since [v3.10.2](https://docs.projectcalico.org/archive/v3.10/release-notes/#v3102) and
changelog since [v3.10.1](https://docs.projectcalico.org/archive/v3.10/release-notes/#v3101).

- Includes a fix for [CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=2020-13597) that allows a Pod to be
compromised and redirect full or partial network traffic from the Node to the compromised Pod.


