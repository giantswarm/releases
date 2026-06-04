# :zap: Giant Swarm Release v31.2.0 for CAPA :zap:

Mitigates the "Dirty Frag" Linux kernel vulnerabilities (esp4/esp6, rxrpc) that could allow local privilege escalation and container escape on affected nodes.

## Changes compared to v31.1.2

### Components

- cluster-aws from v3.6.2 to v3.8.0
- cluster from v2.5.1 to v2.6.1
- Flatcar from v4152.2.3 to [v4593.2.1](https://www.flatcar.org/releases/#release-4593.2.1)
- os-tooling from v1.26.1 to v1.31.0

### cluster-aws [v3.6.2...v3.8.0](https://github.com/giantswarm/cluster-aws/compare/v3.6.2...v3.8.0)

#### Changed

- Support newer Flatcar versions which require a larger root volume size. For ease of migration, enforce at least 15 GB even if a smaller, explicit size is specified in chart values.

### cluster [v2.5.1...v2.6.1](https://github.com/giantswarm/cluster/compare/v2.5.1...v2.6.1)

#### Added

- Add `enabled` flag to `providerIntegration.workers.kubeadmConfig` to disable the rendering of the `KubeadmConfig` resource.
- containerd: Add flag to enable SELinux.

#### Changed

- Values: Fix schema. ([#580](https://github.com/giantswarm/cluster/pull/580))
