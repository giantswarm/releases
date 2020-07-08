## :zap: Giant Swarm Release 12.0.0 for KVM :zap:

This is the first Giant Swarm release which includes Kubernetes v1.17. Components have been updated for improved reliability and observability. Further details about changes to individual components can be found below.

### Kubernetes v1.17.6

- Updated from v1.16.9 - [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v1175).
- **Cloud Provider Labels reach General Availability**: Added as a beta feature way back in v1.2, v1.17 sees the general availability of cloud provider labels.
- **Volume Snapshot Moves to Beta**: The Kubernetes Volume Snapshot feature is now beta in Kubernetes v1.17. It was introduced as alpha in Kubernetes v1.12, with a second alpha with breaking changes in Kubernetes v1.13.
- **CSI Migration Beta**: The Kubernetes in-tree storage plugin to Container Storage Interface (CSI) migration infrastructure is now beta in Kubernetes v1.17. CSI migration was introduced as alpha in Kubernetes v1.14.

### kvm-operator v3.11.2

- Modified Calico deployment to use `-bird-live` as a liveness probe improving observability of failed mesh networking.
- Updated node setup for 1.17 compatibility.
- Removed limits from `calico-kube-controllers` to prevent it from being OOM killed.

### Calico v3.14.0

- Updated from v3.10.1 - [changelog](https://docs.projectcalico.org/v3.14/release-notes/).
- TODO

### etcd v3.4.8

- Updated from v3.3.17 - [changelog](https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.4.md#v348-2020-05-18).

### coredns-app v1.1.10

- Make resource requests/limits configurable.

### cluster-operator v0.23.10

- Align with NGINX IC App 1.7.0, move of LB Service management from azure-operator to the app itself.

### kube-state-metrics-app v1.1.0

- Added `100.64.0.0/10` to the allowed egress subnets in NetworkPolicy.
- Fix invalid cluster role binding for Helm 3 compatibility.

### metrics-server-app v1.1.0

- Added `100.64.0.0/10` to the allowed egress subnets in NetworkPolicy.

### net-exporter v1.9.0

- Added `ntp` collector.
- Added `100.64.0.0/10` to the allowed egress subnets in NetworkPolicy.

### nginx-ingress-controller-app v1.7.1

- Upgraded ingress-nginx from 0.30.0 to 0.33.0 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0330).
- Improved observability, enabled monitoring Service by default for monitoring targets discovery and removed support for disabling it.
