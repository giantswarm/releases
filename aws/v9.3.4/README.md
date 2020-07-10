# :zap: Giant Swarm Release v9.3.4 for AWS :zap:

This release updates managed apps to the latest releases and includes security fixes for 3 CVEs:
- IPv6 rogue router advertisement vulnerability [CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-13597)
- Intel Microcode vulnerabilities [CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543)
- [CVE-2020-8558](https://github.com/kubernetes/kubernetes/issues/92315) which allows for neighboring hosts to bypass localhost boundary

Below, you can find more details on components that were changed with this release.

### calico [v3.10.4](https://docs.projectcalico.org/archive/v3.10/release-notes/)

- Fixes IPv6 rogue router advertisement vulnerability [CVE-2020-13597](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-13597).

For complete details for changes since v3.10.1 please check the upstream release notes at https://docs.projectcalico.org/archive/v3.10/release-notes/.

### cert-exporter [v1.2.3](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#v123-2020-05-15)

- Updated prometheus/client_golang dependency.
- Migrated from dep to go modules.
- Moved to App deployment.

### cluster-autoscaler v1.16.5 ([Giant Swarm app v1.16.0](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md#v1160-2020-05-26))

- Upgraded upstream cluster-autoscaler from v1.16.2 to v1.16.5 - [changelog](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.16.5).
- Set `scan-interval` to 30 seconds (from 10 seconds) to save resources.
- Set `scale-down-unneeded-time` to 5 minutes (from the default of 10 minutes) to release unneeded nodes earlier.
- Lower `scaleDownUtilizationThreshold` to 0.5.

### cluster-operator [v0.23.10](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.10)

- Align with NGINX IC App 1.7.0, move of LB Service management from azure-operator to the app itself.

### containerlinux [2512.2.1](https://www.flatcar-linux.org/releases/#release-2512.2.1)

Security fixes:
- Mitigations for Intel Microcode vulnerabilities ([CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543))

Changes:
- A source code and licensing overview is available under `/usr/share/licenses/INFO`

Updates:
- Linux [4.19.128](https://lwn.net/Articles/822841/)
- intel-microcode [20200609](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20200609)

### coredns v1.6.5 ([Giant Swarm app v1.1.10](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v1110-2020-06-29))

- Make resource requests/limits configurable.
- Make forward options optional.

### kube-state-metrics v1.9.5 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
- Fixed invalid cluster role binding for Helm 3 compatibility.

### Kubernetes 1.16.11 
- Updated from Kubernetes 1.16.9 - 
changelog since [v1.16.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v11610) and
since [v1.16.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v1169)
- Includes a fix for CVE-2020-8558, which allows for neighboring hosts to bypass localhost boundary

### metrics-server v0.3.3 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### net-exporter [v1.9.0](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#190---2020-06-29)

- Add `ntp` collector.
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

### nginx-ingress-controller v0.34.0 ([Giant Swarm app v1.7.2](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v172-2020-07-10))

- Upgraded upstream ingress-nginx from v0.30.0 to v0.34.0 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0340).
- Improved observability, enabled monitoring Service by default for monitoring targets discovery and removed support for disabling it.
