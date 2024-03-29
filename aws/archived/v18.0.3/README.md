# :zap: Giant Swarm Release v18.0.3 for AWS :zap:

This release fixes a misconfiguration of IRSA that caused downtime in the API server during upgrades between v17 and v18 with IRSA feature enabled. 

***IRSA highlights***
- Please read the [updated documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)
- Prior to upgrades please reach out to your Account Engineer and GiantSwarm team will help you in seemless migration if you have already enabled IRSA.
- Please remember to adapt the IAM policies prior to upgrade as specified in the [documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)

***Features removed from v18.0.0***

- Cilium CNI was rolled back, and this release still uses AWS CNI. This was done because of an upstream bug that prevented some advanced networking features to work properly.
- Out-of-tree cloud controller manager was removed, and in-tree alternative was enabled again because of an upstream bug regarding AWS CNI. Bug is already fixed in 1.24 version of upstream repo. 

## Change details


### aws-operator [13.2.4](https://github.com/giantswarm/aws-operator/releases/tag/v13.2.4)

### Changed

- Avoid duplicate --service-account-signing-key-file flag being set for API server.
- Add old cloudfront domain name as service-account-issuer when domain alias is enabled in IRSA.

### Fixed

- Add cluster API endpoint as sts audience.

### metrics-server [2.0.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.0.0)

### Changed

- Switch to HA setup.
