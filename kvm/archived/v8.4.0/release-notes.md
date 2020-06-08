## :zap:  Giant Swarm Release 8.4.0 for KVM :zap:

This release contains Kubernetes 1.14.6 and therefor fixes for CVE-2019-9512 and CVE-2019-9514.

### Kubernetes v1.14.6
- Update kubernetes to 1.14.6 (CVE-2019-9512, CVE-2019-9514) - [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.14.md#v1146)

### kvm-operator v3.8.0
- Fix NTP server configuration on tenant nodes.

### coreDNS v1.6.2 ([GS 0.7.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v070))
- Updated from 1.5.1 - [changelog](https://coredns.io/2019/08/13/coredns-1.6.2-release/)

### kube-state-metrics v1.7.2 ([GS 0.4.0](https://github.com/giantswarm/kubernetes-kube-state-metrics/blob/master/CHANGELOG.md#v040))
- Updated from 1.6.0 - [changelog](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.7.2)

### ignition
- Add name label for default and kube-system namespaces.
- Use v1 stable for giantswarm-critical priority class.
- Introduce explicit resource reservation for OS resources and container runtime.

### Calico v3.8.2
- Update calico to 3.8.2 - [changelog](https://docs.projectcalico.org/v3.8/release-notes/)
