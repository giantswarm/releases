# :zap: Giant Swarm Release v36.1.0 for AKS :zap:

## Changes compared to v35.3.0

### Components

- cluster-aks from v0.7.0 to v0.9.0
- cluster from v8.3.0 to v9.0.2
- Kubernetes from v1.35.7 to [v1.36.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.4)

### cluster-aks [v0.7.0...v0.9.0](https://github.com/giantswarm/cluster-aks/compare/v0.7.0...v0.9.0)

#### Added

- Add `global.providerSpecific.nodeOSUpgrade.channel` to choose the AKS node OS auto-upgrade channel (`NodeImage`, `SecurityPatch`, `Unmanaged` or `None`) independently of the Kubernetes version. It defaults to `NodeImage`, which is what AKS already applies to clusters that don't set it, so existing clusters keep their weekly node image upgrades.
- Add `global.providerSpecific.nodeOSUpgrade.maintenanceWindow` to restrict when node OS upgrades may start. When set, the chart adds an `aksManagedNodeOSUpgradeSchedule` `MaintenanceConfiguration` to the `AzureASOManagedControlPlane` resources. Exactly one of the `daily`, `weekly`, `absoluteMonthly` or `relativeMonthly` schedules must be set, which the values schema validates.
- Add the `giantswarm.io/machine-pool` label to the nodes of each node pool.
- Add an optional per-node-pool `subnetArmId` override for bring-your-own VNet clusters, so node pools can use different subnets in the same VNet.
- Add a pull request template and a release PR body that trigger the `cluster-test-suites` E2E tests.
- Add the `giantswarm.io/prevent-deletion` label to `AzureASOManagedCluster` and `AzureASOManagedControlPlane` when `global.metadata.preventDeletion` is enabled.
- Enable Workload Identity.
- Project `{{$clusterName}}-cluster-aso-exports` ConfigMap with OIDC issuer profile and principal ID.
- Add `global.controlPlane.logging.audit` to ship API server audit logs from the AKS control plane to an Event Hub via an Azure diagnostic setting. Disabled by default.

#### Changed

- Always render `autoUpgradeProfile.upgradeChannel: none` on the ManagedCluster, so AKS never upgrades the Kubernetes version on its own; it is controlled only by the Giant Swarm `Release`.
- Disallow additional properties in `global.controlPlane`, `global.providerSpecific` and their nested objects, so mistyped or removed values fail schema validation instead of being silently ignored.
- Bump `OCIRepository` resources from the `source.toolkit.fluxcd.io/v1beta2` API to `source.toolkit.fluxcd.io/v1`.
- Bump ASO `containerservice.azure.com` resources to `v1api20250801` and `network.azure.com` resources to `v20250301`.
- Publish the chart to `cluster-catalog` / `cluster-test-catalog` instead of `giantswarm-catalog` / `giantswarm-test-catalog`.

#### Fixed

- Scrape AKS-managed CoreDNS by selecting `k8s-app: kube-dns` in the kube-prometheus-stack CoreDNS Service.
- Configure k8s-dns-node-cache for AKS so it redirects DNS traffic to the node-local cache.
- Increase default value of node pool setting `maxPods` to 110 regardless of network mode.
- Add `helm.sh/resource-policy: keep` to `AzureASOManagedMachinePool` resources so they are deleted by CAPI with the cluster rather than by Helm.

### cluster [v8.3.0...v9.0.2](https://github.com/giantswarm/cluster/compare/v8.3.0...v9.0.2)

#### Added

- Add `filesTemplateName` hook under `providerIntegration.workers.kubeadmConfig`. It names a provider template that renders a YAML list of files, once per node pool, so that a provider can deploy files to selected node pools only. A node pool for which the template renders nothing keeps its `KubeadmConfig` spec, and therefore its spec hash, unchanged.
- Enable `mergeDefaultEvictionSettings` to keep defaults for eviction like `nodefs.available` and `nodefs.inodesFree` which would otherwise be set to 0. This rolls all nodes.
- Add the `app.kubernetes.io/component` label with the app name to the resources rendered per app, so they can be listed together.

#### Changed

- `kubeadm`: Exclude `/etc/.systemd-confext` from `restorecon`.
- Cilium: Replace the catch-all `- operator: Exists` toleration on `cilium-operator` with an explicit list.
- Stop deleting and recreating `/etc/ssl/certs` on nodes for fixing SELinux labeling, so any custom certificates placed there directly are preserved.
- Cilium: Replace the catch-all `- operator: Exists` toleration on the `hubble-relay`, `hubble-ui` and `certgen` components with an explicit list.

#### Removed

- Chart: Remove App to HelmRelease migration.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### Apps

- cert-exporter from v2.12.0 to v2.12.1
- cert-manager from v4.1.1 to v4.1.2
- external-dns from v3.5.0 to v3.7.1
- k8s-dns-node-cache from v2.11.0 to v2.11.1
- net-exporter from v1.24.0 to v1.24.1
- network-policies from v0.3.0 to v0.3.2
- node-exporter from v1.20.13 to v1.21.0
- observability-bundle from v3.4.0 to v3.5.0
- prometheus-blackbox-exporter from v0.9.0 to v0.10.0
- Added rbac-bootstrap v0.3.0

### cert-exporter [v2.12.0...v2.12.1](https://github.com/giantswarm/cert-exporter/compare/v2.12.0...v2.12.1)

#### Fixed

- A cert file that cannot be read no longer aborts the scan of its whole cert path. Previously one unreadable file (such as a root-only `0600` `ca.crt` on an AKS node, where the exporter runs as an unprivileged user) stopped the walk, silently dropping every file sorting after it from the metrics. Unreadable files are now logged and skipped individually.

### cert-manager [v4.1.1...v4.1.2](https://github.com/giantswarm/cert-manager-app/compare/v4.1.1...v4.1.2)

#### Added

- cert-manager-giantswarm-clusterissuer: Allow setting `cnameStrategy` per DNS01 solver with `acme.dns01.cloudflare.cnameStrategy`, `acme.dns01.route53.cnameStrategy` and `acme.dns01.azureDNS.cnameStrategy`, so a `_acme-challenge` CNAME can delegate the challenge to a zone the solver can write to.

#### Changed

- Use `imagePullPolicy: IfNotPresent` for the `cert-manager-giantswarm-clusterissuer` hook Job, so a retried Helm hook reuses the image already present on the node instead of downloading it again.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Increase clusterIssuer job memory limit from 250Mi to 512Mi.
- The ACME HTTP-01 solver of the `letsencrypt-giantswarm` ClusterIssuer now uses Gateway API (`gatewayHTTPRoute`) instead of an Ingress, and targets Envoy Gateway.

#### Removed

- `giantSwarmClusterIssuer.acme.http01.enabled` and `giantSwarmClusterIssuer.acme.http01.ingressClassName` (which defaulted to `nginx`). Setting either now fails the Helm render with a message pointing at the replacement values.

#### Fixed

- Trim the `helm.sh/chart` label of the Giant Swarm resources (VPAs, PodLogs, proxy ConfigMap) to end on an alphanumeric character. A long dev version cut at 63 characters could end on a `.`, and the API server rejected the resource.

### external-dns [v3.5.0...v3.7.1](https://github.com/giantswarm/external-dns-app/compare/v3.5.0...v3.7.1)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Changed

- Upgrade external-dns to [v0.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.23.0).
- Sync to upstream helm chart [1.23.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.23.0).
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Upgrade external-dns to [v0.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.22.0).
- Sync to upstream helm chart [1.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.22.0).
  - Add `replicaCount` value to scale the deployment down to `0` or back to `1`.
  - Add `service.enabled` value to skip creating the `Service`.
  - Add `hostAliases` value to inject entries into the pod's `/etc/hosts`.
  - Add `crd` as a valid `registry` value, with the matching RBAC on `dnsrecords`.
  - Install the new `DNSRecord` CRD.
  - Narrow the RBAC verbs on `dnsendpoints/status` from `*` to `update`.
  - `policy` is now a required value. Our default of `sync` is unchanged.
  - Pin `annotationPrefix` to `external-dns.alpha.kubernetes.io/`, keeping the previous annotation prefix after upstream changed the default.

#### Fixed

- CRD install job: bump the `kubectl-apply-job` dependency to 0.14.0, whose fallback image is `gsoci.azurecr.io/giantswarm/docker-kubectl` instead of `quay.io/giantswarm/docker-kubectl:1.26.0`, which quay.io now refuses with 401.
- Values schema: accept arbitrary keys again in free-form maps and pass-through Kubernetes objects (annotations, labels, `nodeSelector`, `affinity`, security contexts, probes, resources, `extraArgs` as a map, `tlsConfig`, `secretConfiguration.data`), and `name` in `imagePullSecrets` items. This broke installs, for example with the IRSA `eks.amazonaws.com/role-arn` service account annotation on CAPA.
- Values schema: accept the `aws` values that cluster-aws still sets. They have no effect.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### k8s-dns-node-cache [v2.11.0...v2.11.1](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.11.0...v2.11.1)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Fixed

- Make the DNS Service redirected by the `CiliumLocalRedirectPolicy` configurable.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### net-exporter [v1.24.0...v1.24.1](https://github.com/giantswarm/net-exporter/compare/v1.24.0...v1.24.1)

#### Added

- Add `keywords` to `Chart.yaml`.
- Add the `io.giantswarm.application.audience` (`all`) and `io.giantswarm.application.managed`

#### Changed

- Replace `interface{}` with `any` and use for-range over integers (Go modernization).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Set `Chart.yaml` `apiVersion` to `v2`. The chart declares no dependencies and had no

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### network-policies [v0.3.0...v0.3.2](https://github.com/giantswarm/network-policies-app/compare/v0.3.0...v0.3.2)

#### Changed

- Convert the `allow-ingress-from-konnectivity` `CiliumNetworkPolicy` into a `CiliumClusterwideNetworkPolicy` named
- Always exclude the `karpenter` and `aws-load-balancer-controller` namespaces from `denyEgressToIMDS`.

### node-exporter [v1.20.13...v1.21.0](https://github.com/giantswarm/node-exporter-app/compare/v1.20.13...v1.21.0)

#### Added

- `disableSystemdCollector` value to turn the systemd collector off, mirroring the existing `disableConntrackCollector` and `disableNvmeCollector` toggles. The collector needs a D-Bus connection to the host, which is refused on nodes where AppArmor mediates D-Bus (such as AKS Ubuntu nodes running under the default containerd profile), making it fail on every scrape. Defaults to `false`, so behaviour is unchanged.

### observability-bundle [v3.4.0...v3.5.0](https://github.com/giantswarm/observability-bundle/compare/v3.4.0...v3.5.0)

#### Added

- Point kube-prometheus-stack's control-plane ServiceMonitors at the alloy-metrics token Secret.

#### Changed

- Update `alloy` apps to 0.23.1 (Alloy v1.19.2).
- Update `prometheus-operator-crd` to 24.0.0 (Prometheus Operator CRDs v0.94.0).
- Update `kube-prometheus-stack` to 24.0.0 (chart 91.2.3, Prometheus Operator v0.94.0).

#### Fixed

- KSM custom resource state: Set the Gateway API `TCPRoute` and `UDPRoute` collectors to `v1`, which is the version the API server serves.

### prometheus-blackbox-exporter [v0.9.0...v0.10.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.9.0...v0.10.0)

#### Added

- Probe `github.com`, `gsoci.azurecr.io` and `grafana.com` from every node (targets `egress-github`, `egress-registry`, `egress-grafana`), so an allowlist-style firewall change blocking a single domain becomes visible. The new `serviceMonitor.externalTargets` value is a map so per-installation, per-region or per-customer overrides can disable, change or add individual entries without copying the whole list. A separate `serviceMonitor.additionalExternalTargets` key takes regional/customer additions, structurally separated from the Giant Swarm defaults. Adds the `http_2xx_or_401` module for registry endpoints that answer unauthenticated requests with 401. See giantswarm/giantswarm#33409.
- Add the `http_2xx_egress` module, used by the internet egress targets. It carries a 15s timeout so `probe_success` reports whether an endpoint is reachable rather than whether it is fast. It is a separate module rather than a longer timeout on `http_2xx` because several installations pin `http_2xx` in their own custom values, which would silently revert the change there.

#### Removed

- Remove the inert `instance` metric relabeling from the ServiceMonitor template. It interpolated a `url` field that no target defines, so it rendered empty and Prometheus fell back to its `$1` default, leaving `instance` unchanged.

#### Fixed

- Give the internet egress targets a probe deadline above the cross-border baseline: `scrapeTimeout: 20s` on `http-giantswarm`, `egress-github`, `egress-registry` and `egress-grafana`, and a 15s timeout on `http_2xx_or_401`. The exporter applies `min(module timeout, scrapeTimeout - 0.5s offset)`, so the 5s `serviceMonitor.defaults.scrapeTimeout` capped every probe at a 4.5s deadline and raising a module timeout alone had no effect. Installations whose baseline latency is a large fraction of that deadline crossed it on endpoints that were still returning HTTP 200, making `probe_success` report latency rather than reachability.
- Point the `dns-tcp-internal` and `dns-udp-internal` ServiceMonitors at the `dns_*_internal` modules. They referenced the `_external` modules, so both probed `www.prometheus.io` and in-cluster DNS resolution was never monitored.

### rbac-bootstrap [v0.3.0](https://github.com/giantswarm/rbac-bootstrap-app/releases/tag/v0.3.0)

#### Added

- Add `io.giantswarm.application.managed` chart annotation for Backstage visibility.
- Add optional `cluster-reader` ClusterRole (off by default, enabled via `clusterReader.enabled: true`) that aggregates into the built-in `view` ClusterRole and grants read access (`get`/`list`/`watch`) on cluster-scoped resources.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.
