# :zap: Giant Swarm Release v19.0.0-beta1 for AWS :zap:

***THIS IS A BETA RELEASE - We highly advise to test your upgrades in lowest environment or new clusters***

This Giant Swarm release introduces Kubernetes 1.24, replaces the use of AWS CNI with Cilium and upgrades most components. This release also deprecates and removes use of `kiam` by default in favor of `IAM Roles for Service Accounts (IRSA)`.

> **_CNI Warning:_** Upgrade includes changes from AWS CNI to Cilium, please check `Cilium breaking changes` section for further details.

> **_IAM Warning:_** The `kiam` application will be removed from clusters in favor of IAM Roles for Service Accounts (IRSA). please check `IRSA breaking changes` section for further details.

> **_Network Policies Warning:_** Cilium implementation of Kubernetes `NetworkPolicies` has known limitations regarding CIDR selectors. If you have NetworkPolicies containing `ipBlock` fields referencing IPs of the cluster nodes and/or Pod IPs and/or Service IPs and/or 0.0.0.0/0 your applications might stop working when upgrading to this release. Please get in touch with your SA for further details.

:rotating_light: ***Cilium breaking changes***
- The AWS CNI pod subnets are no longer used by Cilium. Please add custom routes with the node subnet(s) CIDR(s) instead of the AWS CNI pod subnets CIDR before upgrading to this release.
- [Network Policy](https://docs.cilium.io/en/stable/concepts/kubernetes/policy/#networkpolicy-state) provided by Cilium does not cover support for setting [ipBlock with Pod IP](https://github.com/cilium/cilium/issues/9209). Components in need of this will have to use [CiliumNetworkPolicy](https://docs.cilium.io/en/stable/concepts/kubernetes/policy/#ciliumnetworkpolicy) which has wider scope. Please inspect your workloads and configured Network Policies carefully.
- The upgrade process [documentation](https://handbook.giantswarm.io/docs/support-and-ops/ops-recipes/upgrade-to-cilium/) is available. Please read it carefully and reach out in case of questions.
- Due to changes to CRs during upgrade the `gitops` automation will have to be suspended and any applied changes backported to the repos before resuming.

:rotating_light: ***IRSA breaking changes***
- Please read the [updated documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)
- Please remember that `kiam-app` will be removed by default during the upgrade.
- Please remember to adapt the IAM policies prior to upgrade as specified in the [documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)
- There have been breaking changes introduced to the `Cloudfront domain alias` to provide greater predictability and enable easier automation of AWS IAM role creation by customers. Please adjust AWS IAM roles accordingly to the [documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/#aws-release-v19) prior to the upgrade.

***General highlights***
- The [k8s-dns-node-cache-app](https://github.com/giantswarm/k8s-dns-node-cache-app/) is now deployed by default. After the upgrade please delete the application if you have deployed it through managed catalog. Afterwards the app will be installed automatically.
- The [prometheus-blackbox-exporter](https://github.com/giantswarm/prometheus-blackbox-exporter/) is a new monitoring component deployed by default. This is a blackbox monitoring tool that will validate internal, DNS and external connectivity.
- `Cilium` will have Hubble UI enabled by default for troubleshooting and observability.

## Change details


### app-operator [6.6.3](https://github.com/giantswarm/app-operator/releases/tag/v6.6.3)

### Changed

- Lowered resource requests and limits
- Changed VPA to consider unique and workload cluster operators as well and added support for min allowed fields of CPU and memory


### aws-operator [14.14.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.14.0)

### Fixed
- Use alpine as image for aws-cni's routes-fixer.

#### Added
- Added ami IDs for flatcar `3374.2.4` and `3374.2.5`.
- Added ami IDs for flatcar `3510.2.0`.

#### Changed
- Set ENV for nftables in `aws-cni`.
- Allow externalDNS role to be assumed by any SA containing "external-dns" to allow multiple app deployments.


### cluster-operator [5.6.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.6.0)

#### Added
- Add use of runtime/default seccomp profile.
#### Changed
- Disable `default policies` creation for Cilium App.
- Add `cert-manager` as dependency for `aws-pod-identity-webhook` app.
- Enable `ciliumNetworkPolicy.enabled` flag for all apps.
#### Fixed
- Don't remove finalizer for in-cluster apps when cluster is being deleted.



### aws-ebs-csi-driver [2.21.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.21.0)

#### Added
- Add cilium network policies.



### cert-exporter [2.4.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.4.0)

#### Added
- Add cilium network policies.



### cert-manager [2.21.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.21.0)

#### Added
- Chart: Add `CiliumNetworkPolicy`. ([#301](https://github.com/giantswarm/cert-manager-app/pull/301))



### cilium [0.9.3](https://github.com/giantswarm/cilium-app/releases/tag/v0.9.3)

#### Changed
- Use `image.registry` value as image registry for all containers in the chart.

#### Added
- Add network policy to allow exposing hubble UI through ingress.


### external-dns [2.35.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.35.0)

#### Changed
- Make CiliumNetworkPolicy CR creation be deployed or not with a flag in the Values.



### kube-state-metrics [1.15.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.15.0)

#### Added
- Add cilium network policies.



### metrics-server [2.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.1.0)

#### Added
- Add cilium network policies.



### net-exporter [1.14.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.14.0)

#### Added
- Add `Cilium Network Policy` to net-exporter.
#### Changed
- Don't push net-exporter to capa-app-collection because it's already a default app.
- Don't push net-exporter to cloud-director-app-collection and vsphere-app-collection because it's already in default app bundles.



### vertical-pod-autoscaler [3.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.0)

#### Added
- Add cilium network policies.



### observability-bundle [0.3.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.3.0)

#### Changed
- Add new app dependency mechanism (`app-operator.giantswarm.io/depends-on`) to the prometheus-operator-app and agent so they are not installed until the CRD app is deployed.
- prometheus-operator: drop `apiserver_request_slo_duration_seconds_bucket` metrics from apiserver
- upgrade `prometheus-operator-app` to 4.0.1 and `prometheus-operator-crd` to 4.0.0
- upgrade `prometheus-agent` to 0.3.0 to support chinese registry
#### Added
- Add `promtail-app` v1.0.1 disabled by default.



### k8s-dns-node-cache-app [2.1.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.1.0)

#### Changed
- Switch to ServiceMonitors for metrics scraping.



### prometheus-blackbox-exporter [0.3.1](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.3.1)

#### Fixed
- Change image registry for DaemonSet.



### cilium-servicemonitors [0.0.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.0.2)

#### Changed
- Add labels to servicemonitors



### irsa-servicemonitors [0.0.1](https://github.com/giantswarm/irsa-servicemonitors-app/releases/tag/v0.0.1)

#### Added

- First release.


