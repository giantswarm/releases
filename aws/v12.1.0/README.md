# :zap: Giant Swarm Release v12.1.0 for AWS :zap:

**If you are upgrading from 12.0.0, upgrading to this release will not roll your nodes. It will only update the apps.**

## Change details

#### Significant changes

This release upgrades Cert Manager from upstream v0.9.0 to v0.15.2.

It is part of a larger effort to maximize reliability of upgrades, by keeping 3rd party components within 30 days of upstream releases. 
associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager. An
[upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh)
is provided to assist with the process.

**This is a breaking change.** Please ensure you have read the [upgrade notes](https://github.com/giantswarm/cert-manager-app#upgrading-from-v090-giant-swarm-app-v108-to-20x).

**Note for SEs:** when upgrading clusters to this release, ensure the App is also cordoned on the installation's
control plane.

### cert-manager [2.0.2](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.0.2)

#### Changed

- Upgrade cert-manager from 0.9.0 to 0.15.2. ([#31](https://github.com/giantswarm/cert-manager-app/pull/31))
