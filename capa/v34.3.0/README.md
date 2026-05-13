# :zap: Giant Swarm Release v34.3.0 for CAPA :zap:

Mitigates the "Dirty Frag" Linux kernel vulnerabilities (esp4/esp6, rxrpc) that could allow local privilege escalation and container escape on affected nodes.

## Changes compared to v34.2.0

### Components

- cluster-aws from v7.6.1 to v7.7.1
- Flatcar from v4459.2.4 to [v4593.2.1](https://www.flatcar.org/releases/#release-4593.2.1)
- os-tooling from v1.28.0 to v1.31.0

### cluster-aws [v7.6.1...v7.7.1](https://github.com/giantswarm/cluster-aws/compare/v7.6.1...v7.7.1)

#### Changed

- Support newer Flatcar versions which require a larger root volume size. For ease of migration, enforce at least 15 GB even if a smaller, explicit size is specified in chart values.

