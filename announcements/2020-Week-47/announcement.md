# :zap: Announcements for this week (November 20, 2020) :zap:

First things first, we offer a new and powerful way to **follow releases, changes, and improvements**. Please check out our public [Changes and Releases](https://docs.giantswarm.io/changes/) section, which is a central place to look at new versions of our apps, new tenant cluster releases, UI improvements and even documentation changes.

Don't hesitate to tell us how this works for you and what you'd like to see improved.

## Web UI

- Node pools can now be configured to a minimum size of zero nodes, to allow scaling to zero and saving costly resources when there are no workloads scheduled.
- Potential security vulnerabilities in two dependencies got fixed.
- We fixed several glitches, including when upgrading on Azure to releases above v13.0.0, handling of errors in cluster or node pool creation, the upgrade dialog, and the app installation dialog.

## gsctl

- With v0.26.0, the `gsctl upgrade cluster` command supports to set a specific target release to upgrade to via the `--release` flag. This allows top skip patch or minor releases, and also allows to upgrade to a beta release.

## kubectl gs

- The [`template cluster`](https://docs.giantswarm.io/reference/kubectl-gs/template-cluster/) and [`template nodepool`](https://docs.giantswarm.io/reference/kubectl-gs/template-nodepool/) commands for creating cluster and node pool manifests no longer provide the `--region` flag, as this value is now filled in automatically by the Control Plane Kubernetes API.
- Additionally, the `--release` and `--release-branch` flags have been removed from the [`template nodepool`](https://docs.giantswarm.io/reference/kubectl-gs/template-nodepool/) command for improved simplicity, they are also now filled in automatically by the Control Plane Kubernetes API.

## Documentation

- We updated the guide on [Preparing a Kubernetes cluster for the use of GPUs](https://docs.giantswarm.io/guides/kubernetes-gpu/) to make it easier to follow, especially on Azure.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
