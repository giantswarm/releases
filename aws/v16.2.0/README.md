# :zap: Giant Swarm Release v16.2.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [5.3.0](https://github.com/giantswarm/app-operator/releases/tag/v5.3.0)

#### Changed
- Use dynamic client instead of generated client for watching chart CRs.
- Validate `.spec.kubeConfig.secret.name` in validation resource.



### containerlinux [2983.2.0](https://www.flatcar-linux.org/releases/#release-2983.2.0)

Upgraded from version 2905.2.6.

This upgrade provides the solution for a high number of security issues in the Linux Kernel, Containerd and Golang.
Please check details in the [upstream changelog page](https://www.flatcar-linux.org/releases/).



### etcd [3.4.18](https://github.com/etcd-io/etcd/releases/tag/v3.4.18)

Upgraded from version 3.4.16.

Please check details in the [upstream changelog page](https://github.com/etcd-io/etcd/blob/main/CHANGELOG-3.4.md).



### cert-exporter [2.0.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.0.0)

#### Changed
- Export presence of `giantswarm.io/service-type: managed` label in cert-manager `Issuer` and `ClusterIssuer` CR referenced by `Certificate` CR `issuerRef` spec field to `cert_exporter_certificate_cr_not_after` metric as `managed_issuer` label.
- Add `--monitor-files` and `--monitor-secrets` flags.
- Add Deployment to helm chart to avoid exporting secrets and certificate metrics from DaemonSets.
- Build container image using retagged giantswarm alpine.
- Run as non-root inside container.



