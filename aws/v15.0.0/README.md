# :zap: Giant Swarm Release v15.0.0 for AWS :zap:

**Warning:** From AWS workload cluster release v15.0.0, the automatic termination of unhealthy nodes is enabled by default. For more information about the feature and information how to disable it, please follow the [official documentation](https://docs.giantswarm.io/advanced/automatic-node-termination/). 

## Change details


### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases;
- Emit events when config maps and secrets referenced in App CRs are updated;
- Cache k8sclient, helmclient for later use;
- Apply the namespaceConfig to the desired chart;
- Install apps in CAPI Workload Clusters;
- Apply `compatibleProvider`, `namespace` metadata validation based on the relevant `AppCatalogEntry` CR;
- Add annotations from Helm charts to AppCatalogEntry CRs;
- Enable Vertical Pod Autoscaler.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support;
- Restore chart-operator when it had been deleted;
- Use backoff in chart CR watcher to wait until kubeconfig secret exists;

#### Changed
- Updated Helm to v3.5.3;
- Replace status webhook with chart CR status watcher;
- Sort AppCatalogEntry CRs by version and created timestamp;
- Watch cluster namespace for per workload cluster instances of app-operator.



### cluster-operator [3.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.8.0)

#### Changed
- Adjust helm chart to be used with config-controller.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support;
- Fix clusterIPRange value in configmap;
- Fix kubeconfig resource to search secrets in all namespaces;
- Add `AllowedLabels` to clusterconfigmap resource to prevent unnecessary updates.

#### Added
- Create app CR for per cluster app-operator instance;
- Add `appfinalizer` resource to remove finalizers from workload cluster app CRs.

#### Removed
- Do not add `VersionBundle` to new `CertConfig` specs (`CertConfig`s are now versioned using a label). **This change requires using `cert-operator` 1.0.0 or later.**



### aws-operator [10.6.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.6.1)

#### Added
- S3 vpc endpoint to AWS CNI subnet;
- Clean up VPC peerings from a cluster VPC when is cluster deleted;
- Clean up Application and Network loadbalancers created by Kubernetes when cluster is deleted;
- Add new flatcar AMIs.

#### Changed
- Upgrade `k8scloudconfig` to v10.8.1 which includes a change to better determine if memory eviction thresholds are crossed;
- Update Flatcar AMI's to the latest stable releases;
- Enabled EBS CSI migration;
- Avoid TCCPN stack failure by checking if a control-plane tag exists before adding it;
- Look up cloud tags in all namespaces;
- Find certs in all namespaces;
- Enable `terminate unhealthy node` feature by default;
- Add node termination counter per cluster metric;
- Fix issues with etcd initial cluster resolving into ELB and causing errors;
- Update `k8scloudconfig` to version `v10.5.0` to support kubernetes `v1.20`;
- Use `networkctl reload` for managing networking to avoid bug in `systemd`.

#### Removed
- Removed default storage-class annotation, EBS CSI driver is taking over.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support;
- Cancel update loop if source or target release is not found;
- Updated IPAM library to avoid IP conflicts.



### containerlinux [2765.2.6](https://www.flatcar-linux.org/releases/#release-2765.2.6)


**Security fixes**

* Linux ([CVE-2020-26558](https://nvd.nist.gov/vuln/detail/CVE-2020-26558), [CVE-2021-0129](https://nvd.nist.gov/vuln/detail/CVE-2021-0129), [CVE-2020-24587](https://nvd.nist.gov/vuln/detail/CVE-2020-24587), [CVE-2020-24586](https://nvd.nist.gov/vuln/detail/CVE-2020-24586), [CVE-2020-24588](https://nvd.nist.gov/vuln/detail/CVE-2020-24588), [CVE-2020-26139](https://nvd.nist.gov/vuln/detail/CVE-2020-26139), [CVE-2020-26145](https://nvd.nist.gov/vuln/detail/CVE-2020-26145), [CVE-2020-26147](https://nvd.nist.gov/vuln/detail/CVE-2020-26147), [CVE-2020-26141](https://nvd.nist.gov/vuln/detail/CVE-2020-26141), [CVE-2021-3564](https://nvd.nist.gov/vuln/detail/CVE-2021-3564), [CVE-2021-28691](https://nvd.nist.gov/vuln/detail/CVE-2021-28691), [CVE-2021-3587](https://nvd.nist.gov/vuln/detail/CVE-2021-3587), [CVE-2021-3573](https://nvd.nist.gov/vuln/detail/CVE-2021-3573))

**Bug fixes**

* Update-engine sent empty requests when restarted before a pending reboot ([Flatcar#388](https://github.com/kinvolk/Flatcar/issues/388))
motd login prompt list of failed services: The output of “systemctl list-units –state=failed –no-legend” contains a bullet point which is not expected and ended up being taken as the unit name of failed units which was previously on the start of the line. Filtered the bullet point out to stay compatible with the old behavior in case upstream would remove the bullet point again. ([coreos-overlay#1042](https://github.com/kinvolk/coreos-overlay/pull/1042))

**Updates**

* Linux ([5.10.43](https://lwn.net/Articles/859022/))


### kubernetes [1.20.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.8)

### Feature

- Kubernetes is now built using Go 1.15.13 ([#102786](https://github.com/kubernetes/kubernetes/pull/102786), [@thejoycekung](https://github.com/thejoycekung)) [SIG Cloud Provider, Instrumentation, Release and Testing]

### Failing Test

- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]

### Bug or Regression

- Added jitter factor to lease controller that better smears load on kube-apiserver over time. ([#101652](https://github.com/kubernetes/kubernetes/pull/101652), [@marseel](https://github.com/marseel)) [SIG API Machinery and Scalability]
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: avoid setting cached Sku when updating VMSS and VMSS instances ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix a bug on the endpoint slices mirroring controller where endpoint NotReadyAddresses were mirrored as Ready to the corresponding EndpointSlice ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix errors when accessing Windows container stats for Dockershim ([#98510](https://github.com/kubernetes/kubernetes/pull/98510), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node and Windows]
- Fix removing pods from podTopologyHints mapping ([#101896](https://github.com/kubernetes/kubernetes/pull/101896), [@aheng-ch](https://github.com/aheng-ch)) [SIG Node]
- Fix: avoid nil-pointer panic when checking the frontend IP configuration ([#101739](https://github.com/kubernetes/kubernetes/pull/101739), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed garbage collection of dangling VolumeAttachments for PersistentVolumes migrated to CSI on startup of kube-controller-manager. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt)) [SIG Apps and Storage]
- Improve speed of vSphere PV provisioning and reduce number of API calls ([#102350](https://github.com/kubernetes/kubernetes/pull/102350), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- Kubeadm: remove the "ephemeral_storage" request from the etcd static pod that kubeadm deploys on stacked etcd control plane nodes. This request has caused sporadic failures on some setups due to a problem in the kubelet with cadvisor and the LocalStorageCapacityIsolation feature gate. See this issue for more details: https://github.com/kubernetes/kubernetes/issues/99305 ([#102673](https://github.com/kubernetes/kubernetes/pull/102673), [@jackfrancis](https://github.com/jackfrancis)) [SIG Cluster Lifecycle]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- Reverted the previous fix for portforward cleanup because it introduced a kubelet regression which can lead into segmentation faults. ([#102586](https://github.com/kubernetes/kubernetes/pull/102586), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]

### Other (Cleanup or Flake)

- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102341](https://github.com/kubernetes/kubernetes/pull/102341), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]

## Dependencies

### Added
_Nothing has changed._

### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.15 → v0.0.19

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



### cluster-autoscaler [1.20.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.3)

#### Changed
- Set docker.io as the default registry;
- Allow users to set container resources;
- Update cluster-autoscaler to version `1.20.0`;
- Update templates to run app in Control-Plane cluster.

#### Added
- Push app to the control plane app catalog;
- Push app to the aws/azure app-collections.

#### Fixed
- Remove priority class used in integration tests that is no longer required.



### cert-manager [2.7.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.0)

#### Added
- Add support for *dns01* ACME solver.

#### Changed
- Update to upstream `v1.3.1` ([#155](https://github.com/giantswarm/cert-manager-app/pull/155)). This mitigates failed cert-manager-app installations due to CRD conversion issues;
- cert-manager-app now requires kubernetes version >=1.16.0. ([#151](https://github.com/giantswarm/cert-manager-app/pull/151));
- Switch rbac rules from `extensions` to `networking.k8s.io` for ingresses. ([#151](https://github.com/giantswarm/cert-manager-app/pull/151));
- Rename clusterissuer subchart to match it's name in its Chart.yaml. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140));
- Make pods of deployments use read-only file systems. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140));
- Make pre-install/pre-upgrade hooks use server side apply. Possibly fixes upgrade timeouts. ([#140](https://github.com/giantswarm/cert-manager-app/pull/140)).

#### Fixed
- Allow strings and integers in values schema for resources requests and limits. ([#150](https://github.com/giantswarm/cert-manager-app/pull/150)).



### chart-operator [2.14.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.14.0)

#### Changed
- Cancel the release resource when the manifest object already exists.
- Cancel the release resource when helm returns an unknown error.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.

