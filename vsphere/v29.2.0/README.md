# :zap: Giant Swarm Release v29.2.0 for vSphere :zap:

## Changes compared to v29.1.0

### Components

- cluster-vsphere from v0.66.0 to v0.68.0
- Kubernetes from v1.29.10 to [v1.29.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.29.md#changelog-since-v12910)

### cluster-vsphere [v0.66.0...v0.68.0](https://github.com/giantswarm/cluster-vsphere/compare/v0.66.0...v0.68.0)

#### **Breaking change**

> [!CAUTION]
> The interface for setting annotations on the `vsphereCluster` CR has changed.
<details>
<summary>Upgrade guide: how to migrate values (from v0.66.0)</summary>

The `additionalVsphereClusterAnnotations` field is now an array of strings to accomodate for escape characters. Convert your dictionary to an array of strings as per the following example.

Old interface for annotations:

```yaml
global:
  providerSpecific:
    additionalVsphereClusterAnnotations:
      my-key: "my-value"
```
New interface for annotations:
```yaml
global:
  providerSpecific:
    additionalVsphereClusterAnnotations:
      - "my-key: value"
```
</details>

#### Changed

- Chart: Update `cluster` to [v1.7.0](https://github.com/giantswarm/cluster/releases/tag/v1.7.0).
  - Add `teleport-init` systemd unit to handle initial token setup before `teleport` service starts
  - Improve `teleport` service reliability by adding proper file and service dependencies and pre-start checks
- Change `global.providerSpecific.additionalVsphereClusterAnnotations` attribute into an array.

### Apps

- cert-manager from v3.8.1 to v3.8.2
- coredns from v1.22.0 to v1.23.0
- observability-bundle from v1.8.0 to v1.9.0

### cert-manager [v3.8.1...v3.8.2](https://github.com/giantswarm/cert-manager-app/compare/v3.8.1...v3.8.2)

#### Changed

- Changed ownership to team Shield

#### Removed

- Get rid of label `giantswarm.io/monitoring_basic_sli` as this slo generation label is not used anymore.

### coredns [v1.22.0...v1.23.0](https://github.com/giantswarm/coredns-app/compare/v1.22.0...v1.23.0)

#### Changed

- Update `coredns` image to [1.11.4](https://github.com/coredns/coredns/releases/tag/v1.11.4).
- Explicitly expose liveness and readiness probe ports in deployments.

#### Removed

- Remove PodSecurityPolicy and associated Resources and values.

### observability-bundle [v1.8.0...v1.9.0](https://github.com/giantswarm/observability-bundle/compare/v1.8.0...v1.9.0)

#### Added

- Add `alloy` v0.7.0 as `alloyEvents`.

#### Changed

- Upgrade `alloy-logs` and `alloy-metrics` to chart 0.7.0.
  - Bumps `alloy` from 1.4.2 to 1.5.0
- upgrade `kube-prometheus-stack` from 65.1.1 to 66.2.1
  - prometheus-operator CRDs from 0.75.0 to 0.78.1
  - prometheus-operator from 0.77.1 to 0.78.1
  - prometheus from 2.54.1 to 2.55.1
  - kube-state-metrics from 2.13.0 to 2.14.0
  - grafana from 8.5.0 to 8.6.0