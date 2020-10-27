# :zap: Giant Swarm Release v12.5.2 for AWS :zap:

This release fixes an issue that causes app-operator to crash when handling cluster deletion.

## Change details



### app-operator [2.3.5](https://github.com/giantswarm/app-operator/releases/tag/v2.3.5)

#### Fixed
- Skip removing finalizer for chart-operator chart CR if its not present.
- Skip deleting chart-operator in case of cluster deletion.
