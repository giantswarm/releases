# :zap: This Week in Provider-Independent Announcements (October 16, 2020):zap:

## Giant Swarm Catalog (Managed Apps)

1. [Cert Manager v2.3.1](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#231---2020-10-29) removes deprecated APIs and upgrades to upstream v1.0.3.
2. [Kong v0.9.2](https://github.com/giantswarm/kong-app/blob/release-v0.9.x/CHANGELOG.md#092---2020-10-28) adds minReadySeconds (optional parameter) to the deployment template. This helps avoid downtime during rolling updates.
3. [EFK v0.3.4](https://github.com/giantswarm/efk-stack-app/blob/master/CHANGELOG.md#034---2020-10-30) improves reliability of upgrades. Builds now automatically do functional testing. Also, pod anti-affinity and PodDisruptionBudget are now enabled by default. Please [read before you deploy](https://github.com/giantswarm/efk-stack-app/tree/v0.3.4#important-notes---read-before-you-deploy).

## Documentation

1. We added [reference documentation for `kubectl gs`](https://docs.giantswarm.io/reference/kubectl-gs/), our kubectl plugin for the Control Plane Kubernetes API.
2. We updated our [security](https://docs.giantswarm.io/basics/security/) documentation and added a bunch of details.
---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
