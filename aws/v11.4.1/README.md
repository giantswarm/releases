# :zap: Giant Swarm Release v11.4.1 for AWS :zap:

This release re-activates the recent AWS release of high-availability (HA) masters while fixing issues with certain OIDC configurations.

## Change details


### aws-operator [8.7.1](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.1)

#### Added
- Add mapping between similar instance types `m4.16xlarge` and `m5.16xlarge`.
- Add `lifecycle` label to the `aws_operator_ec2_instance_status` metric to distinguish on-demand and spot.
#### Changed
- Use `k8s-apiserver` image which includes CAs to enable OIDC.
- Fix failing go template rendering of KMS encryption content.
- Use `0.1.1` tag for `k8s-api-heahtz` image.
- Use `0.1.0` tag for `k8s-setup-network-env` image.
- Use `0.1.0` tag for `aws-attach-etcd-dep` image.



### containerlinux [2512.2.1](https://www.flatcar-linux.org/releases/#release-2512.2.1)

Security fixes:
- Fix the Intel Microcode vulnerabilities ([CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543))

Changes:
- A source code and licensing overview is available under `/usr/share/licenses/INFO`

Updates:
- Linux [4.19.128](https://lwn.net/Articles/822841/)
- intel-microcode [20200609](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20200609)


### coredns [1.1.10](https://github.com/giantswarm/coredns-app/releases/tag/v1.1.10)

#### Changed
- Make resource requests/limits configurable.
- Applying Go modules.



### external-dns [1.2.2](https://github.com/giantswarm/external-dns-app/releases/tag/v1.2.2)

#### Changed
- Upgrade upstream external-dns from v0.5.18 to v0.7.2.



### kiam [1.3.0](https://github.com/giantswarm/kiam-app/releases/tag/v1.3.0)

#### Added
- Set kiam region flag for STS endpoint
#### Changed
- Upgrade architect-orb to 0.10.0



### kube-state-metrics [1.1.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.1.0)

#### Changed
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
#### Fixed
- Fix invalid cluster role binding for Helm 3 compatibility.



### metrics-server [1.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.1.0)

#### Changed
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.



### net-exporter [1.9.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.0)

#### Added
- Add `ntp` collector.


