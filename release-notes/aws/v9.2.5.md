## :zap: Giant Swarm Release 9.2.5 for AWS :zap:

**If you are upgrading from 9.2.4, upgrading to this release will not roll your nodes. It will only update the apps.**

This release improves the reliability of chart-operator.

Below, you can find more details on components that were changed with this release.

### chart-operator [v0.13.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#v0130-2020-04-21)

- Fix update state calculation and status resource for long running deployments.
- Handle 503 responses when GitHub Pages is unavailable.
- Make HTTP client timeout configurable for pulling chart tarballs in AWS China.
- Switch from dep to go modules.
- Fix problem pushing chart to default app catalog.
- Always set chart CR annotations so update state calculation is accurate.
- Only update failed Helm releases if the chart values or version has changed.
- Deploy as a unique app in app collection in control plane clusters.
