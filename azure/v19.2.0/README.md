# :zap: Giant Swarm Release v19.2.0 for Azure :zap:

<< Add description here >>

## Change details


### app-operator [6.8.0](https://github.com/giantswarm/app-operator/releases/tag/v6.8.0)

#### Added
- Add Service Monitor by default to make it complain with the latest monitoring improvements 



### cert-operator [3.2.0](https://github.com/giantswarm/cert-operator/releases/tag/v3.2.0)

#### Fixed
- Expand policy expception to cover old deployments.



### cluster-operator [5.7.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.7.0)

#### Added
- Add necessary values for PSS policy warnings.
#### Removed
- Values: Remove release override for `nginx-ingress-controller`. ([#1625](https://github.com/giantswarm/cluster-operator/pull/1625))



### kubernetes [1.24.16](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.16)

#### Feature
- Kubernetes 1.24.x is now built with Go 1.20.5 ([#119203](https://github.com/kubernetes/kubernetes/pull/119203), [@cpanato](https://github.com/cpanato)) [SIG Release and Testing]
- Kubernetes is now built with Go 1.20.6 ([#119369](https://github.com/kubernetes/kubernetes/pull/119369), [@xmudrii](https://github.com/xmudrii)) [SIG Release and Testing]
#### Bug or Regression
- Only declare Job as finished after removing all Pod finalizers to avoid orphan Pods ([#119165](https://github.com/kubernetes/kubernetes/pull/119165), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and Testing]
- Updated cAdvisor to v0.44.2 - Fix metrics in cri-o when a container restarts ([#118800](https://github.com/kubernetes/kubernetes/pull/118800), [@harche](https://github.com/harche)) [SIG Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.44.1 → v0.44.2](https://github.com/google/cadvisor/compare/v0.44.1...v0.44.2)
#### Removed
_Nothing has changed._



### containerlinux [3510.2.5](https://www.flatcar-linux.org/releases/#release-3510.2.5)

 _Changes since **Stable 3510.2.4**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-3338](https://nvd.nist.gov/vuln/detail/CVE-2023-3338), [CVE-2023-3390](https://nvd.nist.gov/vuln/detail/CVE-2023-3390))
 
 #### Bug fixes:
 
 - Resolved the conflicting FD usage of libselinux and systemd which caused, e.g., a systemd crash on certain watchdog interaction during shutdown (patch in systemd 252.11)
 
 #### Updates:
 
 - Linux ([5.15.119](https://lwn.net/Articles/936675) (includes [5.15.118](https://lwn.net/Articles/935584)))
 - systemd ([252.11](https://github.com/systemd/systemd-stable/releases/tag/v252.11) (from 252.5))


### azuredisk-csi-driver [1.27.0](https://github.com/giantswarm/azuredisk-csi-driver-app/releases/tag/v1.27.0)

#### Fixed
- **Breaking:** added policy exceptions. this makes app dependent on kyverno.
- Added required values for pss policies.



### etcd-kubernetes-resources-count-exporter [1.4.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.4.0)

#### Changed
- Add Max memory (default 500Mi) for VPA.



### external-dns [2.38.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.38.0)

#### Changed
- Move CRD jobs into a separated subchart ([#275](https://github.com/giantswarm/external-dns-app/pull/275)).
- Prepare new values for alignment ([]()).
  - Add domainFilter and extraArgs values.
  - Add interval, namepsaceFilter and minEventSyncInterval values.
  - Add txtPrefix value with higher priority.
  - Add txtOwnerId value with higher priority.
  - Add annotationFilter value with higher priority.
- Allow projected volumes across all providers ([#282](https://github.com/giantswarm/external-dns-app/pull/282)).
#### Removed
- Hardcoded references to `provider==vmware` ([#277](https://github.com/giantswarm/external-dns-app/pull/277)).



### net-exporter [1.17.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.17.0)

#### Changed
- Add security context values to make chart comply to PodSecurityStandard restricted profile.



### node-exporter [1.17.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.17.1)

#### Changed
- fix apparmor annotation



### vertical-pod-autoscaler [3.5.4](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.4)

#### Changed
- Specified `failureThreshold` and `periodSeconds` for recommender's liveness probe.
- Upgrade dependency chart to 7.1.0.
- Upgrade VPA components to 0.14.0



### observability-bundle [0.7.2](https://github.com/giantswarm/observability-bundle/releases/tag/v0.7.2)

#### Changed
- Upgrade `prometheus-operator-app` to 5.0.7.



### k8s-dns-node-cache [1.2.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v1.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### security-bundle [0.16.2](https://github.com/giantswarm/security-bundle/releases/tag/v0.16.2)

#### Added
- Add depends-on annotation to set App depenencies.
#### Changed
- Update to `kyverno` (app) version 0.14.10, introducing Cilium Network Policy support.
- Update to `kyverno-policies` (app) version 0.20.0.



