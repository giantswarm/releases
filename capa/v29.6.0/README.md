# :zap: Giant Swarm Release v29.6.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v29.5.0

### Components



* cluster-aws from 2.5.0 to [2.6.0](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...v2.6.0)






### cluster-aws [2.6.0](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...v2.6.0)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools
- Use reduced IAM permissions on worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.

#### Fixed

- Explicitly set aws-node-termination-handler queue region so crash-loops are avoided, allowing faster startup





### Apps



* aws-nth-bundle from 1.2.0 to [1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)



* cilium from 0.25.1 to [0.25.2](https://github.com/giantswarm/cilium-app/compare/v0.25.1...v0.25.2)



* security-bundle from 1.8.2 to [1.9.1](https://github.com/giantswarm/security-bundle/compare/v1.8.2...v1.9.1)





### aws-nth-bundle [1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)

#### Added

- Forward proxy settings to `aws-node-termination-handler-app` as environment variables





### cilium [0.25.2](https://github.com/giantswarm/cilium-app/compare/v0.25.1...v0.25.2)

#### Changed

- Upgrade cilium to [v1.15.13](https://github.com/cilium/cilium/releases/tag/v1.15.13).





### security-bundle [1.9.1](https://github.com/giantswarm/security-bundle/compare/v1.8.2...v1.9.1)

#### Changed

- Update `trivy-operator` (app) to v0.10.3.
- Update `trivy` (app) to v0.13.1.
- Update `kyverno` (app) to v0.18.1.
- Update `kyverno-crds` (app) to v1.12.0.
- Update `kyverno-policies` (app) to v0.21.0.
- Update `starboard-exporter` (app) to v0.8.0.
- Update `trivy-operator` (app) to v0.10.2.
- Update `trivy` (app) to v0.13.0.
- Update `falco` (app) to v0.9.1.




