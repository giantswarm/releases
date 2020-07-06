# :zap: Giant Swarm Release v12.0.1 for KVM :zap:

<< Add description here >>

## Change details


### calico [3.14.1](https://github.com/projectcalico/calico/releases/tag/v3.14.1)

#### Other changes
- Add port 6443 (Kubernetes API server) to failsafe ports felix #2349 (@neiljerram)



### etcd [3.4.9](https://github.com/etcd-io/etcd/releases/tag/v3.4.9)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.4.8...v3.4.9) and [v3.4 upgrade guide](https://github.com/etcd-io/etcd/blob/master/Documentation/upgrades/upgrade_3_4.md) for any breaking changes.
#### Package `wal`
- Add [missing CRC checksum check in WAL validate method otherwise causes panic](https://github.com/etcd-io/etcd/pull/11924).
  - See https://github.com/etcd-io/etcd/issues/11918.
#### Go
- Compile with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).
<hr>



### kubernetes [1.17.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.17.8)

#### API Change
- Fixed: log timestamps now include trailing zeros to maintain a fixed width ([#91207](https://github.com/kubernetes/kubernetes/pull/91207), [@iamchuckss](https://github.com/iamchuckss)) [SIG Apps and Node]
#### Bug or Regression
- Fixes CSI volume attachment scaling issue by using informers. ([#91307](https://github.com/kubernetes/kubernetes/pull/91307), [@yuga711](https://github.com/yuga711)) [SIG API Machinery, Apps, Node, Storage and Testing]
- If firstTimestamp is not set use firstTimestamp or eventTime when printing event ([#91055](https://github.com/kubernetes/kubernetes/pull/91055), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Kubeadm increased to 5 minutes its timeout for the TLS bootstrapping process to complete upon join ([#89735](https://github.com/kubernetes/kubernetes/pull/89735), [@rosti](https://github.com/rosti)) [SIG Cluster Lifecycle]
- hyperkube: Use debian-hyperkube-base@v1.1.0 image
  
    A previous release built hyperkube using the debian-hyperkube-base@v1.0.0,
    which was updated to address a CVE in the CNI plugins.
    
    A side-effect of using this new image was that the networking packages
    (namely `iptables`) drifted from the versions used in the kube-proxy images.
  
    The following issues were filed on kube-proxy failures when using hyperkube:
    - https://github.com/kubernetes/kubernetes/issues/92275
    - https://github.com/kubernetes/kubernetes/issues/92272
    - https://github.com/kubernetes/kubernetes/issues/92250
  
    To address this, the new debian-hyperkube-base image (v1.1.0) uses the
    debian-iptables base image (v12.1.0), which includes iptables-wrapper, a
    script used to determine the correct iptables mode to run in. ([#92494](https://github.com/kubernetes/kubernetes/pull/92494), [@justaugustus](https://github.com/justaugustus)) [SIG Cluster Lifecycle and Release]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### kvm-operator [3.11.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.11.1)

#### Changed
- Use Release.Revision in Helm chart for Helm 3 support.
- Fix OIDC settings.



### cluster-operator [0.23.9](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.9)

Not found


### containerlinux [2512.2.1](https://www.flatcar-linux.org/releases/#release-2512.2.1)

## Flatcar updates

Security fixes:
- Fix the Intel Microcode vulnerabilities ([CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543))

Changes:
- A source code and licensing overview is available under `/usr/share/licenses/INFO`

Updates:
- Linux [4.19.128](https://lwn.net/Articles/822841/)
- intel-microcode [20200609](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20200609)


