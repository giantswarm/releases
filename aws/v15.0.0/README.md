# :zap: Giant Swarm Release v15.0.0 for AWS :zap:

**Warning:** From AWS workload cluster release v15.0.0, the automatic termination of unhealthy nodes is enabled by default. For more information about the feature and information how to disable it, please follow the [official documentation](https://docs.giantswarm.io/advanced/automatic-node-termination/). 

## Change details


### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases.
- Emit events when config maps and secrets referenced in App CRs are updated.



### cluster-operator [3.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.8.0)

### Changed
- Adjust helm chart to be used with config-controller.

### Fixed

- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.
- Fix clusterIPRange value in configmap.
- Fix kubeconfig resource to search secrets in all namespaces.



### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### external-dns [2.3.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.3.1)

#### Changed

- Increase memory limit to 100Mi since we ran into out of memory problems. This will make the app more stable.



### aws-operator [10.5.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.5.0)

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.
- Cancel update loop if source or target release is not found.
- Updated IPAM library to avoid IP conflicts.
#### Added
- Enabled EBS CSI migration.
- Clean up VPC peerings from a cluster VPC when is cluster deleted.
- Clean up Application and Network loadbalancers created by Kubernetes when cluster is deleted.
- Add new flatcar AMIs.
#### Changed
- Fix issues with etcd initial cluster resolving into ELB and causing errors.
- Update `k8scloudconfig` to version `v10.5.0` to support kubernetes `v1.20`.
- Use `networkctl reload` for managing networking to avoid bug in `systemd`.
- Avoid TCCPN stack failure by checking if a control-plane tag exists before adding it.
- Look up cloud tags in all namespaces
- Find certs in all namespaces
- Enable terminate unhealthy node feature by default.
- Add node termination counter per cluster metric.
#### Removed
- Removed default storage-class annotation, EBS CSI driver is taking over.


### containerlinux [2765.2.3](https://www.flatcar-linux.org/releases/#release-2765.2.3)


**Security fixes**



*   Linux ([CVE-2021-28964](https://nvd.nist.gov/vuln/detail/CVE-2021-28964), [CVE-2021-28972](https://nvd.nist.gov/vuln/detail/CVE-2021-28972), [CVE-2021-28971](https://nvd.nist.gov/vuln/detail/CVE-2021-28971), [CVE-2021-28951](https://nvd.nist.gov/vuln/detail/CVE-2021-28951), [CVE-2021-28952](https://nvd.nist.gov/vuln/detail/CVE-2021-28952), [CVE-2021-29266](https://nvd.nist.gov/vuln/detail/CVE-2021-29266), [CVE-2021-28688](https://nvd.nist.gov/vuln/detail/CVE-2021-28688), [CVE-2021-29264](https://nvd.nist.gov/vuln/detail/CVE-2021-29264), [CVE-2021-29649](https://nvd.nist.gov/vuln/detail/CVE-2021-29649), [CVE-2021-29650](https://nvd.nist.gov/vuln/detail/CVE-2021-29650), [CVE-2021-29646](https://nvd.nist.gov/vuln/detail/CVE-2021-29646), [CVE-2021-29647](https://nvd.nist.gov/vuln/detail/CVE-2021-29647), [CVE-2021-29154](https://nvd.nist.gov/vuln/detail/CVE-2021-29154), [CVE-2021-29155](https://nvd.nist.gov/vuln/detail/CVE-2021-29155), [CVE-2021-23133](https://nvd.nist.gov/vuln/detail/CVE-2021-23133))

**Bug fixes**



*   Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))

**Updates**



*   Linux ([5.10.32](https://lwn.net/Articles/853762/))


### kubernetes [1.20.7](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.7)

### API Change

- We have added a new Priority & Fairness rule that exempts all probes (/readyz, /healthz, /livez) to prevent 
  restarting of "healthy" kube-apiserver instance(s) by kubelet. ([#101112](https://github.com/kubernetes/kubernetes/pull/101112), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]

### Feature

- Kubernetes is now built using go1.15.11 ([#101192](https://github.com/kubernetes/kubernetes/pull/101192), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built using go1.15.12 ([#101845](https://github.com/kubernetes/kubernetes/pull/101845), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]

### Bug or Regression

- Azurefile: Normalize share name to not include capital letters ([#100731](https://github.com/kubernetes/kubernetes/pull/100731), [@kassarl](https://github.com/kassarl)) [SIG Cloud Provider and Storage]
- EndpointSlice IP validation now matches Endpoints IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSlice controllers are less likely to create duplicate EndpointSlices. ([#101763](https://github.com/kubernetes/kubernetes/pull/101763), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Ensure service deleted when the Azure resource group has been deleted ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix panic in JSON logging format caused by missing Duration encoder ([#101158](https://github.com/kubernetes/kubernetes/pull/101158), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fix smb mount PermissionDenied issue on Windows ([#99550](https://github.com/kubernetes/kubernetes/pull/99550), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider, Storage and Windows]
- Fix: azure file inline volume namespace issue in csi migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx)) [SIG Apps, Cloud Provider, Node and Storage]
- Fix: not tagging static public IP ([#101752](https://github.com/kubernetes/kubernetes/pull/101752), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: set "host is down" as corrupted mount ([#101398](https://github.com/kubernetes/kubernetes/pull/101398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug where startupProbe stopped working after a container's first restart ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed port-forward memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Kubectl create service now respects namespace flag ([#101005](https://github.com/kubernetes/kubernetes/pull/101005), [@zxh326](https://github.com/zxh326)) [SIG CLI]
- Kubelet: improve the performance when waiting for a synchronization of the node list with the kube-apiserver ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123)) [SIG Node]
- No support endpointslice in linux userpace mode ([#101503](https://github.com/kubernetes/kubernetes/pull/101503), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Renames the timeout field for the DelegatingAuthenticationOptions to TokenRequestTimeout and set the timeout only for the token review client. Previously the timeout was also applied to watches making them reconnecting every 10 seconds. ([#101103](https://github.com/kubernetes/kubernetes/pull/101103), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth and Cloud Provider]
- Respect ExecProbeTimeout=false for dockershim ([#101126](https://github.com/kubernetes/kubernetes/pull/101126), [@jackfrancis](https://github.com/jackfrancis)) [SIG Node and Testing]

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._


### aws-ebs-csi-driver [2.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.1.0)

#### Changed

- Update aws-ebs-csi-driver to v1.1.0.



### aws-cni [1.8.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.8.0)

Changes since v1.7.10:

* Bug - Use symmetric return path for non-VPC traffic - alternate solution (#1475, @kishorj)
* Bug - Gracefully handle failed ENI SG update (#1341, @jayanthvn)
* Bug - Fix CNI crashing when there is no available IP addresses (#1499, @M00nF1sh)
* Bug - Use primary ENI SGs if SG is null for Custom networking (#1259, @jayanthvn)
* Bug - Don't cache dynamic VPC IPv4 CIDR info (#1113, @anguslees)
* Improvement - Address Excessive API Server calls from CNI Pods (#1419, @achevuru)
* Improvement - refine ENI tagging logic (#1482, @M00nF1sh)
* Improvement - Change tryAssignIPs to assign up to configured WARM_IP_TARGET (#1279, @jacksontj)
* Improvement - Use regional STS endpoint (#1332, @nithu0115)
* Improvement - Update containernetworking dependencies (#1200, @mogren)
* Improvement - Split Calico manifest into two (#1410, @caseydavenport)
* Improvement - Update Calico manifest to support ARM & AMD (#1282, @jayanthvn)
* Improvement - Auto gen of AWS CNI, metrics helper and calico artifacts through helm (#1271, @jayanthvn)
* Improvement - Refactor EC2 Metadata IMDS code (#1225, @anguslees)
* Improvement - Unnecessary logging for each CNI invocation (#1469, @jayanthvn)
* Improvement - New instance types (#1463, @jayanthvn)
* Improvement - Use 'exec' ENTRYPOINTs (#1432, @anguslees)
* Improvement - Fix logging texts for ENI cleanup (#1209, @mogren)
* Improvement - Remove Duplicated vlan IPTable rules (#1208, @mogren)
* Improvement - Minor code cleanup (#1198, @mogren)
* HelmChart - Adding flags to support overriding container runtime endpoint. (#1443, @haouc)
* HelmChart - Add podLabels to amazon-vpc-cni chart (#1440, @haouc)
* HelmChart - Add workflow to sync aws-vpc-cni helm chart to eks-charts (#1430, @fawadkhaliq)
* Testing - Remove validation of VPC CIDRs from ip rules (#1476, @kishorj)
* Testing - Updated agent version (#1474, @cgchinmay)
* Testing - Fix for CI failure (#1470, @achevuru)
* Testing - Binary for mtu and veth prefix check (#1458, @cgchinmay)
* Testing - add test to verify cni-metrics-helper puts metrics to CW (#1461, @abhipth)
* Testing - add e2e test for security group for pods (#1459, @abhipth)
* Testing - Added Test cases for EnvVars check on CNI daemonset (#1431, @cgchinmay)
* Testing - add test to verify host networking setup & cleanup (#1457, @abhipth)
* Testing - Runners failing because of docker permissions (#1456, @jayanthvn)
* Testing - decouple test helper input struct from netlink library (#1455, @abhipth)
* Testing - add custom networking e2e test suite (#1445, @abhipth)
* Testing - add integration test for ipamd env variables (#1453, @abhipth)
* Testing - add agent for testing pod networking (#1448, @abhipth)
* Testing - fix format of commited code to fix unit test step (#1449, @abhipth)
* Testing - Unblocks Github Action Integration Tests (#1435, @couralex6)
* Testing - add warm ENI/IP target integration tests (#1438, @abhipth)
* Testing - add service connectivity test (#1436, @abhipth)
* Testing - add network connectivity test (#1424, @abhipth)
* Testing - add ginkgo automation framework (#1416, @abhipth)
* Testing - Add some test coverage to allocating ENIs (#1234, @mogren)
* Testing - Add some minimal tests to metrics (#1228, @mogren)

Changes since v1.7.9:

* Improvement - Multi card support - Prevent route override for primary ENI across multi-cards ENAs (#1396 , [@jayanthvn](https://github.com/Jayanthvn))



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

#### Added
- Added new configuration value `extraArgs`.



### cluster-autoscaler [1.20.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.3)

#### Changed
- Set docker.io as the default registry
- Allow users to set container resources


### net-exporter [1.10.1](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.1)




### cert-manager [2.7.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.0)

- Update to upstream `v1.3.1` ([#155](https://github.com/giantswarm/cert-manager-app/pull/155)). This mitigates failed cert-manager-app installations due to CRD conversion issues.



### chart-operator [2.14.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.14.0)

#### Changed
- Cancel the release resource when the manifest object already exists.
- Cancel the release resource when helm returns an unknown error.



