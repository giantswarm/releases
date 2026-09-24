# :zap: Giant Swarm Release v35.1.1 for Azure :zap:

## Changes compared to v35.1.0

### Components

- cluster-azure from v9.3.0 to v9.3.1
- cluster from v8.3.0 to v8.3.1

### cluster [v8.3.0...v8.3.1](https://github.com/giantswarm/cluster/compare/v8.3.0...v8.3.1)

#### Added

- Add the `app.kubernetes.io/component` label with the app name to the resources rendered per app, so they can be listed together.

#### Changed

- Cilium: Replace the catch-all `- operator: Exists` toleration on the `hubble-relay`, `hubble-ui` and `certgen` components with an explicit list.
