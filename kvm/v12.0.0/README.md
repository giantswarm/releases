# :zap: Giant Swarm Release 12.0.0 for KVM :zap:

This is the first Giant Swarm release which includes Kubernetes v1.17. Many core components and default apps have been updated for improved reliability and observability. Further details about changes to individual components can be found below.

---

### Kubernetes [1.17.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.8)

- **Cloud Provider Labels reach General Availability**: Added as a beta feature way back in v1.2, v1.17 sees the general availability of cloud provider labels.
- **Volume Snapshot Moves to Beta**: The Kubernetes Volume Snapshot feature is now beta in Kubernetes v1.17. It was introduced as alpha in Kubernetes v1.12, with a second alpha
with breaking changes in Kubernetes v1.13.
- **CSI Migration Beta**: The Kubernetes in-tree storage plugin to Container Storage Interface (CSI) migration infrastructure is now beta in Kubernetes v1.17. CSI migration was
introduced as alpha in Kubernetes v1.14.


### Calico [3.14.1](https://github.com/projectcalico/calico/releases/tag/v3.14.1)

- Added port 6443 (Kubernetes API server) to failsafe ports for felix.
- Fixed IPv6 rogue router advertisement vulnerability ([CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-13597)).


### etcd [3.4.9](https://github.com/etcd-io/etcd/releases/tag/v3.4.9)

- Added [missing CRC checksum check in WAL validate method otherwise causes panic](https://github.com/etcd-io/etcd/pull/11924).
  - See https://github.com/etcd-io/etcd/issues/11918.
- Compiled with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).


### kvm-operator [3.12.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.12.0)

- Improved upgrades from KVM v11 releases.
- Fixed CR validation errors during cluster creation.
- Upgraded node image to QEMU 4.2.0 and Fedora 32.
- Modified Calico deployment to use `-bird-live` as a liveness probe improving observability of failed mesh networking.
- Removed limits from `calico-kube-controllers` to prevent it from being OOM killed.


### cluster-operator [0.23.12](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.12)

- Aligned with NGINX IC App 1.7.0 moving LB Service management from the operator to the app itself.


### Container Linux [2512.2.1](https://www.flatcar-linux.org/releases/#release-2512.2.1)

- Fixed Intel Microcode vulnerabilities ([CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543)).
- Updated to Linux [4.19.128](https://lwn.net/Articles/822841/).
- Updated to intel-microcode [20200609](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20200609).


### cert-exporter [1.2.3](https://github.com/giantswarm/cert-exporter/releases/tag/v1.2.3)

- Updated prometheus/client_golang dependency.
- Changed to App-based deployment.


### CoreDNS [1.2.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.2.0)

- Made resource requests/limits configurable.
- Make forward options optional.
- Apply a readiness probe.
- Increase the liveness probe failure threshold from 5 failures to 7 failures.


### kube-state-metrics [1.1.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.1.0)

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
- Fixed invalid cluster role binding for Helm 3 compatibility.


### metrics-server [1.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.1.0)

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.


### net-exporter [1.9.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.0)

- Added `ntp` collector.


### nginx-ingress-controller [1.7.2](https://github.com/giantswarm/nginx-ingress-controller-app/releases/tag/v1.7.2)

- Upgraded upstream `ingress-nginx` from 0.30.0 to 0.34.0.
- Improved observability, enabled monitoring Service by default for monitoring targets discovery and removed support for disabling it.
- Added support for additional service for internal traffic. Existing service can still be configured to be either for public (default) or internal traffic.
- Made monitoring headless Service non-optional.
- Enabled managed app monitoring via monitoring service.


### node-exporter [1.2.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.2.0)

- Changed Priority Class to `system-node-critical`.
