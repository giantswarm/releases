# :zap: Giant Swarm Release v14.1.7 for Azure :zap:

This release upgrades Kubernetes to [v1.19.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#v11912).

## Change details


### azure-operator [5.5.4](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.4)

Meta release for guaranteed node rolling. No changes in functionality.


### kubernetes [1.19.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.12)

#### Feature
- Kubernetes is now built using Go 1.15.13 ([#102809](https://github.com/kubernetes/kubernetes/pull/102809), [@thejoycekung](https://github.com/thejoycekung)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Failing Test
- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
#### Bug or Regression
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: avoid setting cached Sku when updating VMSS and VMSS instances ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix errors when accessing Windows container stats for Dockershim ([#98510](https://github.com/kubernetes/kubernetes/pull/98510), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node and Windows]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed garbage collection of dangling VolumeAttachments for PersistentVolumes migrated to CSI on startup of kube-controller-manager. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt)) [SIG Apps and Storage]
- Fixes an issue where default RBAC policy could fail to reconcile on API server startup if an error was encountered ([#101954](https://github.com/kubernetes/kubernetes/pull/101954), [@voutcn](https://github.com/voutcn)) [SIG Auth]
- Improve speed of vSphere PV provisioning and reduce number of API calls ([#102351](https://github.com/kubernetes/kubernetes/pull/102351), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- Reverted the previous fix for portforward cleanup because it introduced a kubelet regression which can lead into segmentation faults. ([#102588](https://github.com/kubernetes/kubernetes/pull/102588), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
#### Other (Cleanup or Flake)
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102342](https://github.com/kubernetes/kubernetes/pull/102342), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._

