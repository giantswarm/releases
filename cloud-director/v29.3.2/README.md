# :zap: Giant Swarm Release v29.3.2 for VMware Cloud Director :zap:

## Changes compared to v29.3.1

### Apps

- cloud-provider-cloud-director from v0.4.0 to v0.5.0

### cloud-provider-cloud-director [v0.4.0...v0.5.0](https://github.com/giantswarm/cloud-provider-cloud-director-app/compare/v0.4.0...v0.5.0)

#### Changed

- Update to `1.6.1-gs` of CPI. (Upstream version `1.6.1` plus unreleased patch to address LB health monitor upgrade issue).
- Update to `1.6.0` of CSI. This also removes `--v=5` from the CSI plugin as this flag is no longer supported in this version.
