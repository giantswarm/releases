# :zap: Giant Swarm Release v29.0.0 for CAPA :zap:

## Changes compared to v28.0.0

- cluster-aws from 1.0.1 to 1.2.1
- Flatcar from 3815.2.2 to 3815.2.5
- Kubernetes from 1.28.11 to 1.29.7
- cloud-provider-aws (formerly aws-cloud-controller-manager-app) from 1.28.6-gs1 to 1.29.3-gs1
- cluster-autoscaler from 1.28.5-gs1 to 1.29.3-gs1
- security-bundle from 1.7.0 to 1.8.0

### cilium [v0.25.1](https://github.com/giantswarm/cilium-app/releases/tag/v0.25.1)

#### Changed

- Fix regression setting Policy BPF Max map `policyMapMax` back to 65536 from 16384.
- Upgrade cilium to v1.15.6.

### k8s-dns-node-cache [v2.8.1](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.8.1)

#### Changed

- Reduce security exceptions.
  - Enable readOnly FS moving config to emptyDir volume.
  - Remove `NET_ADMIN` and drop `ALL` capabilities.
  - Add `NET_BIND_SERVICE` capability.
  - Add policy exception for `require-non-root-groups/autogen-check-runasgroup`.
  - Remove disallow-capabilities-* policy exceptions.
- Update PolicyException CR version to v2beta1.

### net-exporter [v1.21.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.21.0)

#### Changed

- Enable readOnlyRootFilesystem in securityContext.
- Update module google.golang.org/grpc to v1.65.0. 
- Update k8s modules to v0.30.2. 
- Update quay.io/giantswarm/alpine Docker tag to v3.20.1.
- Add node and app labels in ServiceMonitor.

### security-bundle [v1.7.0...v1.8.0](https://github.com/giantswarm/security-bundle/compare/v1.7.0...v1.8.0)

#### Added

- Add `kyverno-crds` app to handle Kyverno CRD install.

#### Changed

- Update `kyverno` (app) to v0.17.15. This version disables the CRD install job in favor of `kyverno-crds` App.
- Update `starboard-exporter` (app) to v0.7.11.
