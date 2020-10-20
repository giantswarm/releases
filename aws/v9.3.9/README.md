# :zap: Giant Swarm Release v9.3.9 for AWS :zap:

**If you are upgrading from 9.3.8, upgrading to this release will not roll your nodes.**


This is a patch release to update cluster-operator mitigating the unexpected deletion of app CRs in clusters. 

## Change details


### cluster-operator [0.23.17](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02317---2020-10-19)

### Changed

- Get app-operator version from releases CR. 
- Delete all chartconfig migration logic.

### chart-operator [0.13.2](https://github.com/giantswarm/chart-operator/blob/helm2/CHANGELOG.md#v0132-2020-06-23)

### Changed

- Calculating md5sum from Chart go struct.
- Add metrics for Helm releases with a mismatched namespace.
