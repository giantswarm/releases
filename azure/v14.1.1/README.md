# :zap: Giant Swarm Release v14.1.1 for Azure :zap:

This is a bugfix release to resolve a few bugs related to the cluster autoscaler.
We strongly suggest upgrading any 14.x workload cluster to this release to ensure the cluster autoscaler feature works properly.

Warning: to avoid downtimes in the ingress-based workloads, before upgrading to this release it is important to ensure your cluster has a recent version (1.14.0 or newer) 
of the `Nginx Ingress Controller APP` running. Please get in touch with your Solution Engineer before upgrading if you have any concern.

## Change details
### azure-operator [5.5.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.0)

### Added

- Add new handler that creates `AzureClusterIdentity` CRs and the related `Secrets` out of Giant Swarm's credential secrets.
- Ensure `AzureCluster` CR has the `SubscriptionID` field set.
- Reference `Spark` CR as bootstrap reference from the `MachinePool` CR.
- Ensure node pools min size is applied immediately when changed.

### Fixed

- Avoid blocking the whole `AzureConfig` handler on cluster creation because we can't update the `StorageClasses`.
- Avoid overriding the NP size when the scaling is changed by autoscaler.

### kubernetes [1.19.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.8)

#### API Change
- Kubernetes is now built using go1.15.8 ([#99093](https://github.com/kubernetes/kubernetes/pull/99093), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Aggregate errors when putting vmss ([#98350](https://github.com/kubernetes/kubernetes/pull/98350), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Avoid marking node as Ready until node has synced with API servers at least once ([#97996](https://github.com/kubernetes/kubernetes/pull/97996), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Cleanup subnet in frontend IP configs to prevent huge subnet request bodies in some scenarios. ([#98288](https://github.com/kubernetes/kubernetes/pull/98288), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix CSI-migrated inline EBS volumes failing to mount if their volumeID is prefixed by aws:// ([#96821](https://github.com/kubernetes/kubernetes/pull/96821), [@wongma7](https://github.com/wongma7)) [SIG Storage]
- Fix azure file migration issue ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx)) [SIG Auth, Cloud Provider and Storage]
- Fix the description of command line flags that can override --config ([#98873](https://github.com/kubernetes/kubernetes/pull/98873), [@changshuchao](https://github.com/changshuchao)) [SIG Scheduling]
- Fix to recover CSI volumes from certain dangling attachments ([#96617](https://github.com/kubernetes/kubernetes/pull/96617), [@yuga711](https://github.com/yuga711)) [SIG Apps and Storage]
- Fixed a bug that the kubelet cannot start on BtrfS. ([#98015](https://github.com/kubernetes/kubernetes/pull/98015), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Fixed a bug where aggregator_unavailable_apiservice metrics were reported for deleted apiservices. ([#96421](https://github.com/kubernetes/kubernetes/pull/96421), [@dgrisonnet](https://github.com/dgrisonnet)) [SIG API Machinery and Instrumentation]
- Fixed provisioning of Cinder volumes migrated to CSI when StorageClass with AllowedTopologies was used. ([#98311](https://github.com/kubernetes/kubernetes/pull/98311), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixes a panic in the disruption budget controller for PDB objects with invalid selectors ([#98776](https://github.com/kubernetes/kubernetes/pull/98776), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Apps]
- Kubeadm: get k8s CI version markers from k8s infra bucket ([#98836](https://github.com/kubernetes/kubernetes/pull/98836), [@hasheddan](https://github.com/hasheddan)) [SIG Cluster Lifecycle and Release]
- Kubelet should ignore cgroup driver check on Windows node. ([#98385](https://github.com/kubernetes/kubernetes/pull/98385), [@pacoxu](https://github.com/pacoxu)) [SIG Node]
- Performance regresssion #97685 has been fixed ([#98432](https://github.com/kubernetes/kubernetes/pull/98432), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Static pods will be deleted gracefully. ([#98103](https://github.com/kubernetes/kubernetes/pull/98103), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Truncates a message if it hits the NoteLengthLimit when the scheduler records an event for the pod that indicates the pod has failed to schedule. ([#98715](https://github.com/kubernetes/kubernetes/pull/98715), [@carlory](https://github.com/carlory)) [SIG Scheduling]
- Warning about using a deprecated volume plugin is logged only once. ([#96751](https://github.com/kubernetes/kubernetes/pull/96751), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
#### Other (Cleanup or Flake)
- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Resolves flakes in the Ingress conformance tests due to conflicts with controllers updating the Ingress object ([#98430](https://github.com/kubernetes/kubernetes/pull/98430), [@liggitt](https://github.com/liggitt)) [SIG Network and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.37.3 → v0.37.4](https://github.com/google/cadvisor/compare/v0.37.3...v0.37.4)
#### Removed
_Nothing has changed._



### app-operator [3.2.0](https://github.com/giantswarm/app-operator/releases/tag/v3.2.0)

#### Added
- Include `apiVersion`, `restrictions.compatibleProviders` in appcatalogentry CRs.
  
#### Changed
  
- Limit the number of AppCatalogEntry per app.
- Delete legacy finalizers on app CRs. 
- Reconciling appCatalog CRs only if pod is unique.
#### Fixed
- Updating status as cordoned if app CR has cordoned annotation.



### external-dns [2.1.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.1.0)

#### Added
- Allow the sync policy to be configured. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports customisation of the txt-owner-id (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports dry-run mode and warns the user if enabled. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
#### Changed
- Rework the way the txt prefix is generated (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Rework how the annotation filter value is generated (whilst still defaulting for default app). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))



### azure-scheduled-events [0.2.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.2.0)

### Added
- Remove the `Node` from Kubernetes API server right before approving the termination event.

### Fixed
- Keep looping on events if one loop errors out.
