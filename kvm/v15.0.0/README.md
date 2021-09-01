# :zap: Giant Swarm Release v15.0.0 for KVM :zap:

This is the first release for KVM with Kubernetes 1.20 and Calico 3.19. It also migrates the Calico datastore from etcd to Kubernetes.

## Change details

### calico [3.19.2](https://github.com/projectcalico/calico/releases/tag/v3.19.2)

View the changelogs for Calico 3.16 through 3.19:
- [3.16](https://docs.projectcalico.org/archive/v3.16/release-notes/#v3160)
- [3.17](https://docs.projectcalico.org/archive/v3.17/release-notes/#v3170)
- [3.18](https://docs.projectcalico.org/archive/v3.18/release-notes/#v3180)
- [3.19](https://docs.projectcalico.org/archive/v3.19/release-notes/#v3190)


### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

#### Added

- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.

#### Changed

- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.


### chart-operator [2.19.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.19.0)

#### Removed

- Remove `tillermigration` resource now Helm 3 migration is complete.

#### Added

- Add releasemaxhistory resource which ensures we retry at a reduced rate when
there are repeated failed upgrades.

#### Changed

- Increase memory limit for deploying large charts in workload clusters.
- Upgrade Helm release when failed even if version or values have not changed
to handle situations like failed webhooks where we should retry.
- Prepare helm values to configuration management.
- Update architect-orb to v3.0.0.
For CAPI clusters:
- Add tolerations to start on `NotReady` nodes for installing CNI.
- Create `giantswarm-critical` priority class.
- Use host network to allow installing CNI packaged as an app.

#### Fixed

- Improve status message when helm release has failed max number of attempts.


### coredns [1.6.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.0)

#### Changed

- Make `targetCPUUtilizationPercentage` in HPA configurable.
- Update `coredns` to upstream version [1.8.3](https://coredns.io/2021/02/24/coredns-1.8.3-release/).
- Increase maximum replica count to 50 when using horizontal pod autoscaling.


### kubernetes [1.20.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.10)

View the major changes since Kubernetes v1.19 [here](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#changelog-since-v1190).


### kube-state-metrics [1.4.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.4.0)

#### Changed

- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.


### kvm-operator [3.18.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.18.0)

#### Changed

- Upgrade `k8scloudconfig` to `v10.8.1` which includes a change to better determine if memory eviction thresholds are crossed.
- Update for compatibility with Calico v3.19.


### metrics-server [1.4.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.4.0)

#### Changed

- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.


### net-exporter [1.10.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.3)

#### Changed

- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.

### node-exporter [1.8.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.8.0)

#### Changed

- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.
