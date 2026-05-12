# :zap: Giant Swarm Release v33.3.0 for CAPA :zap:

Mitigates the "Dirty Frag" Linux kernel vulnerabilities (esp4/esp6, rxrpc) that could allow local privilege escalation and container escape on affected nodes.

## Changes compared to v33.2.0

### Components

- cluster-aws from v6.4.4 to v6.5.0
- Flatcar from v4459.2.1 to [v4593.2.1](https://www.flatcar.org/releases/#release-4593.2.1)

### cluster-aws [v6.4.4...v6.5.0](https://github.com/giantswarm/cluster-aws/compare/v6.4.4...v6.5.0)

#### Changed

- Support newer Flatcar versions which require a larger root volume size. For ease of migration, enforce at least 15 GB even if a smaller, explicit size is specified in chart values.

