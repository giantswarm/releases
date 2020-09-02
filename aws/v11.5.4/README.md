# :zap: Giant Swarm Release v11.5.4 for AWS :zap:

This release upgrades external-dns app to v1.4.0 to improve observability.

**If you are upgrading from 11.5.3, upgrading to this release will not roll your nodes. It will only update the apps.**

## Change details

### external-dns v0.7.3 [Giant Swarm app 1.4.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#140---2020-08-21)

- Updated from v0.7.2. Check the [upstream changelog](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.3) for details on all changes.
- Added monitoring headless Service.
- Added more Giant Swarm custom monitoring annotations.
- Explicitly exposed metrics container port.
- Use default external-dns metrics port 7979.
