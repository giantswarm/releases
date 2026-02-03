# :zap: Giant Swarm Release v34.0.0 for vSphere :zap:

## OIDC Structured Authentication (optional)

This release introduces optional support for [Kubernetes Structured Authentication Configuration](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#using-authentication-configuration) for OIDC providers.

### Minimal example

```yaml
global:
  controlPlane:
    oidc:
      structuredAuthentication:
        enabled: true
        issuers:
          - issuerUrl: https://your-idp.example.com
            clientId: kubernetes
```

### Example with customization

```yaml
global:
  controlPlane:
    oidc:
      structuredAuthentication:
        enabled: true
        issuers:
          - issuerUrl: https://your-idp.example.com
            clientId: kubernetes
            usernameClaim: email          # Optional: use 'email' instead of 'sub'
            groupsClaim: roles            # Optional: use 'roles' instead of 'groups'
            usernamePrefix: "oidc:"       # Optional: prefix usernames
            groupsPrefix: "oidc:"         # Optional: prefix groups
```

### Migration from legacy OIDC configuration

If you already use OIDC with the legacy configuration, add `structuredAuthentication.enabled: true` to migrate:

```yaml
global:
  controlPlane:
    oidc:
      issuerUrl: https://your-idp.example.com
      clientId: kubernetes
      structuredAuthentication:
        enabled: true
```

This will automatically convert your legacy configuration to the new structured format.

## Changes compared to v33.1.1

### Components

- cluster-vsphere from v3.4.0 to v4.1.2
- Flatcar from v4459.2.1 to [v4459.2.2](https://www.flatcar.org/releases/#release-4459.2.2)
- Kubernetes from v1.33.6 to [v1.34.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.3)
- os-tooling from v1.26.2 to v1.26.3

### cluster-vsphere [v3.4.0...v4.1.2](https://github.com/giantswarm/cluster-vsphere/compare/v3.4.0...v4.1.2)

#### Added

- Add the `priority-classes` default app, enabled by default. This app provides standardised `PriorityClass` resources like `giantswarm-critical` and `giantswarm-high`, which should replace the previous inconsistent per-app priority classes.
- Add `"helm.sh/resource-policy": keep` annotation to `VSphereCluster` CR so that it doesn't get removed by Helm when uninstalling this chart. The CAPI controllers will take care of removing it, following the expected deletion order.
- Add `"helm.sh/resource-policy": keep` annotation to the provider secret. This is to ensure that it isn't removed by Helm, thus leading to a race condition when deleting the cluster as the vSphere cleaner needs it to clean up resources in vSphere.

#### Changed

- Chart: Update `cluster` to v5.1.2.
- Chart: Update `cluster` to v5.1.1.
- Chart: Update `cluster` to v5.1.0.
- Chart: Update `cluster` to v5.0.0.
- Chart: Update `cluster` to v4.6.0.
- Chart: Update `cluster` to v4.5.1.
- Chart: Update `cluster` to v4.5.0.

### Apps

- cert-exporter from v2.9.14 to v2.9.15
- cilium from v1.3.2 to v1.3.4
- cloud-provider-vsphere from v2.0.1 to v2.2.0
- coredns from v1.28.3 to v1.29.1
- etcd-k8s-res-count-exporter from v1.10.11 to v1.10.12
- network-policies from v0.1.1 to v0.1.3
- node-exporter from v1.20.9 to v1.20.10
- observability-bundle from v2.3.2 to v2.5.0
- Added priority-classes v0.3.0
- security-bundle from v1.15.0 to v1.16.1
- vsphere-csi-driver from v3.4.2 to v3.4.3

### cert-exporter [v2.9.14...v2.9.15](https://github.com/giantswarm/cert-exporter/compare/v2.9.14...v2.9.15)

#### Changed

- Go: Update dependencies.

### cilium [v1.3.2...v1.3.4](https://github.com/giantswarm/cilium-app/compare/v1.3.2...v1.3.4)

#### Changed

- Upgrade Cilium to [v1.18.6](https://github.com/cilium/cilium/releases/tag/v1.18.6).
- Upgrade Cilium to [v1.18.5](https://github.com/cilium/cilium/releases/tag/v1.18.5).

### cloud-provider-vsphere [v2.0.1...v2.2.0](https://github.com/giantswarm/cloud-provider-vsphere-app/compare/v2.0.1...v2.2.0)

#### Added

- Add kamaji.enabled value. If set to true, a deployment instead of the dameonset will be used for CPI controller components.

#### Changed

- Update to upstream `1.34.0`.

### coredns [v1.28.3...v1.29.1](https://github.com/giantswarm/coredns-app/compare/v1.28.3...v1.29.1)

#### Changed

- Update `coredns` image to [1.14.1](https://github.com/coredns/coredns/releases/tag/v1.14.1).
- Update `coredns` image to [1.14.0](https://github.com/coredns/coredns/releases/tag/v1.14.0).

### etcd-k8s-res-count-exporter [v1.10.11...v1.10.12](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.11...v1.10.12)

#### Changed

- Go: Update dependencies.

### network-policies [v0.1.1...v0.1.3](https://github.com/giantswarm/network-policies-app/compare/v0.1.1...v0.1.3)

#### Added

- Add support for Kamaji.

#### Fixed

- Fixed broken templating.

### node-exporter [v1.20.9...v1.20.10](https://github.com/giantswarm/node-exporter-app/compare/v1.20.9...v1.20.10)

#### Removed

- Repository: Remove integration tests.

### observability-bundle [v2.3.2...v2.5.0](https://github.com/giantswarm/observability-bundle/compare/v2.3.2...v2.5.0)

#### Added

- Add KSM metrics `kube_servicemonitor_info` and `kube_podmonitor_info` for ServiceMonitor and PodMonitor resources
- Add KSM metrics `kube_podlog_info` for PodLog resource

#### Changed

- Upgrade `kube-prometheus-stack-app` to 19.0.0
- Update alloy-app to 0.16.0
  - Bumps alloy to 1.12.0

#### Fixed

- Fixed KSM metrics for endpoints

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).

### security-bundle [v1.15.0...v1.16.1](https://github.com/giantswarm/security-bundle/compare/v1.15.0...v1.16.1)

#### Changed

- Add missing dependency to all apps.
- Allow to set multiple dependencies on the depends-on annotation.
- Rename `edgedb` to `gel`.
- Update `cloudnative-pg` (app) to v0.0.12.
- Update `gel` (app) to v1.0.1.

### vsphere-csi-driver [v3.4.2...v3.4.3](https://github.com/giantswarm/vsphere-csi-driver-app/compare/v3.4.2...v3.4.3)

#### Changed

- Update upstream chart to `v3.3.1`
- Make deployment `affinity` and `tolerations` configurable in `values.yaml`.
