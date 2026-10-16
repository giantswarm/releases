# :zap: Giant Swarm Release v35.4.0 for AKS :zap:

This release adds Giant Swarm's Cilium for the BYO CNI opt-in (`global.connectivity.cni.managed: true` in cluster-aks). Cilium is only installed on clusters that opt in. Clusters on AKS-managed networking (the default) get no Cilium and no new dependencies. The BYO CNI mode is not supported yet: use it on test clusters only.

## Changes compared to v35.3.0


### Apps

- Added cilium v1.6.1

### cilium [v1.6.1](https://github.com/giantswarm/cilium-app/releases/tag/v1.6.1)

#### Added

- Declare the `io.giantswarm.application.audience` (`all`) and `io.giantswarm.application.managed` (`true`) annotations, completing the [chart metadata standard](https://docs.giantswarm.io/reference/platform-api/chart-metadata/).

#### Changed

- Upgrade Cilium to [v1.20.2](https://github.com/cilium/cilium/releases/tag/v1.20.2).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to `io.giantswarm.application.team`, in both this chart and the bundled `cilium-giantswarm-cni-customconf` subchart. No template reads `.Chart.Annotations` and no template emits a team label, so rendered output is byte-identical.
- Point `home` at this repository instead of `https://cilium.io/`, as the standard requires. Upstream remains listed under `sources`.
