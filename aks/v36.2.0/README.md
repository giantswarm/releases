# :zap: Giant Swarm Release v36.2.0 for AKS :zap:

## Changes compared to v36.1.0

### Components

- cluster-aks from v0.9.0 to v0.10.0
- cluster from v9.0.2 to v9.4.0

### cluster-aks [v0.9.0...v0.10.0](https://github.com/giantswarm/cluster-aks/compare/v0.9.0...v0.10.0)

#### Added

- Add render-time validation that `outboundType: userDefinedRouting` requires a BYO subnet, and document its route table prerequisite.
- Add Azure CNI Overlay (`global.controlPlane.networking.networkPluginMode`) and Advanced Container Networking Services observability and security (`global.controlPlane.networking.advancedNetworking`), with render-time validation of unsupported combinations.
- Add the BYO CNI opt-in `global.connectivity.cni.managed` (default `false`). When `true`, AKS installs no CNI and the cluster runs Giant Swarm's Cilium with AKS BYO CNI values. Not supported yet.

#### Changed

- Make the observability-bundle apps depend on the `cilium` HelmRelease only on BYO CNI clusters (`global.connectivity.cni.managed`), using the map format of `apps.<appName>.dependsOn` that a newer version of observability-bundle requires.
- Make the security-bundle apps depend on Cilium on BYO CNI clusters, as the bundle does by default.

#### Fixed

- Allow the ACNS `hubble-relay` through the network-policies default deny when ACNS observability is on.
- Fix the security-bundle values key `apps.cloudNativePG` to `apps.cloudnativePG`, so that the cloudnative-pg HelmRelease no longer depends on the `cilium` HelmRelease, which AKS doesn't have.

### cluster [v9.0.2...v9.4.0](https://github.com/giantswarm/cluster/compare/v9.0.2...v9.4.0)

#### Added

- Apps: Add `providerIntegration.apps.<app>.extraConfigs` so provider charts can add templated config maps or secrets to an app's values.
- Ignition: Expose `sysctl` and `sysfs` configuration per node pool via `kernel.sysctls` (`/etc/sysctl.d/90-node-pool.conf`) and `kernel.sysfs` (`/etc/tmpfiles.d/00-node-pool.conf`). The parameters required by kubelet's `protectKernelDefaults` cannot be set by name, and the Image Builder and Cilium sysctl configuration keeps precedence.
- Apps: Add `global.connectivity.cni.managed` (default `true`). When `false`, Cilium isn't rendered and is dropped from the dependencies of every app.
- Workers: Add `nodeLabelsTemplateName` and `taintsTemplateName` hooks under `providerIntegration.workers.kubeadmConfig`. They name provider templates, rendered once per node pool, that add kubelet node labels and taints.

#### Changed

- Ignition: Rename `/etc/sysctl.d/hardening.conf` to `80-hardening.conf` so it sorts after Flatcar's `65-baselayout.conf` but before Image Builder's `99-sysctl.conf` and Cilium's `99-zzz-override_cilium.conf`. Previously the unprefixed file sorted last and overrode both. As a result, nodes now run with `fs.inotify.max_user_watches = 524288` (previously `16384`) and `net.ipv4.conf.all.rp_filter = 0` (previously `1`).
- Ignition: Clean up sysctl hardening. Drop keys already set by Image Builder or Cilium (`fs.inotify.max_user_watches`, `fs.inotify.max_user_instances`, `net.ipv4.conf.all.rp_filter`, `vm.overcommit_memory`, `kernel.panic`, `kernel.panic_on_oops`). Stop disabling `net.ipv4.tcp_timestamps`, since modern kernels randomize timestamp offsets.

#### Removed

- Ignition: Remove `os-hardening.service`. Its `sysctl -p` duplicated what `systemd-sysctl` already applies at boot. Its `gpasswd` calls had no effect: glibc merges group memberships from `/usr/share/baselayout/group` (via `nss-usrfiles`) regardless of `/etc/group`.
- Apps: Remove the static `aks` exceptions for the Cilium dependency of network policies and CoreDNS.
- Workers: Remove the `karpenter.sh/unregistered:NoExecute` taint from Karpenter node pools. Provider charts must add it via `taintsTemplateName`.

### Apps

- cert-manager-crossplane-resources from v0.2.0 to v0.3.0
- Added cilium v1.6.1
- network-policies from v0.3.2 to v0.5.0
- observability-bundle from v3.5.0 to v3.6.0

### cert-manager-crossplane-resources [v0.2.0...v0.3.0](https://github.com/giantswarm/cert-manager-crossplane-resources/compare/v0.2.0...v0.3.0)

#### Added

- Read the AWS account ID, partition, and OIDC domains from the `<clusterName>-crossplane-config` ConfigMap if `providers.aws.accountID`, `providers.aws.partition`, or `providers.aws.oidc.domain` is empty. This supports EKS clusters, whose OIDC domain is only known once the cluster is provisioned. Explicitly set values take precedence over the ConfigMap.

#### Changed

- Restrict AWS provider permissions for `ChangeResourceRecordSets` to TXT records related to acme challenges.
- Create a dedicated role for azure to reduce permissions to TXT records only.

### cilium [v1.6.1](https://github.com/giantswarm/cilium-app/releases/tag/v1.6.1)

#### Added

- Declare the `io.giantswarm.application.audience` (`all`) and

#### Changed

- Upgrade Cilium to [v1.20.2](https://github.com/cilium/cilium/releases/tag/v1.20.2).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Point `home` at this repository instead of `https://cilium.io/`, as the standard requires.

### network-policies [v0.3.2...v0.5.0](https://github.com/giantswarm/network-policies-app/compare/v0.3.2...v0.5.0)

#### Added

- Add `hubbleRelay` to allow the AKS ACNS `hubble-relay` through the default deny.
- Add `allowEgressToUpstreamDNS` to allow pods with `dnsPolicy: Default` to reach upstream DNS resolvers.
- Add `azureWorkloadIdentityWebhook` to allow the AKS workload identity webhook through the default deny.

#### Changed

- Move the chart to `helm/network-policies` to match its name, in preparation for devctl-generated CI.

### observability-bundle [v3.5.0...v3.6.0](https://github.com/giantswarm/observability-bundle/compare/v3.5.0...v3.6.0)

#### Added

- KSM custom resource state: Add `status_parent_accepted` and `status_parent_resolved_refs` gauges for Gateway API `HTTPRoute`, one series per parent with the condition `reason` as label.
- KSM custom resource state: Add `target_info` and `status_ancestor_accepted` metrics for Envoy Gateway `SecurityPolicy`, `BackendTrafficPolicy`, `ClientTrafficPolicy` and `EnvoyExtensionPolicy`.

#### Changed

- *!Breaking change*: Turn `apps.<appName>.dependsOn` from a list into a map from dependency name to `true`/`false`, so that a provider can turn off a single dependency (for example `kyverno-crds` on clusters without the security bundle) instead of overriding the whole list.
