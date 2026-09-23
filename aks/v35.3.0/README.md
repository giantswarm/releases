# :zap: Giant Swarm Release v35.3.0 for AKS :zap:

## Changes compared to v35.2.0

### Components

- cluster-aks from v0.6.0 to v0.7.0

### cluster-aks [v0.6.0...v0.7.0](https://github.com/giantswarm/cluster-aks/compare/v0.6.0...v0.7.0)

#### Added

- Add `global.controlPlane.disableLocalAccounts` to turn off AKS local accounts, so the static cluster-admin credential can no longer be issued and all API server authentication goes through Entra ID. When set, the chart also points the ManagedCluster at `<cluster>-user-kubeconfig` via `operatorSpec.secrets.userCredentials`, because ASO cannot list admin credentials on such a cluster and CAPZ needs to own `<cluster>-kubeconfig` itself. Requires `global.controlPlane.aadProfile.managed: true`, which is validated at render time.

#### Fixed

- Turn the node-exporter systemd collector off on AKS. The collector opens a D-Bus connection to the host, which is refused on AKS Ubuntu nodes because the container runs under the default containerd AppArmor profile. It failed on every scrape, logging an error per node per minute and exporting no `node_systemd_*` metrics. Needs node-exporter-app with `disableSystemdCollector`.
