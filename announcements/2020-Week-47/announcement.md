# :zap: Announcements for this week (November 20, 2020) :zap:

First things first, we offer a new and powerful way to **follow releases, changes, and improvements**. Please check out our public [Releases and Changes](https://docs.giantswarm.io/changes/) section, which is a central place to look at new versions of our apps, new tenant cluster releases, UI improvements and even documentation changes.

Don't hesitate to tell us how this works for you and what you'd like to see improved.

## Web UI

- Node pools can now be configured to a minimum size of zero nodes, to allow scaling to zero and saving costly resources when there are no workloads scheduled.
- When cluster creation or node pool creation fails, errors are now handled more gracefully and a retry is possible in some cases.
- The "About upgrading" dialog you can open when a new tenant cluster release is available for your cluster has been improved in several details.
- The dialog for installing apps has been fixed so that long version numbers don't break the layout.
- We fixed several glitches related to upgrades from Azure tenant cluster releases below v13.x.x to v13.x.x.
- Potential security vulnerabilities in two dependencies got fixed.

## gsctl

- With v0.26.0, the `gsctl upgrade cluster` command supports to set a specific target release to upgrade to via the `--release` flag. This allows top skip patch or minor releases, and also allows to upgrade to a beta release.

## kubectl gs

- The [`template cluster`](https://docs.giantswarm.io/reference/kubectl-gs/template-cluster/) and [`template nodepool`](https://docs.giantswarm.io/reference/kubectl-gs/template-nodepool/) commands for creating cluster and node pool manifests no longer provide the `--region` flag, as this value is now filled in automatically by the Control Plane Kubernetes API.

## Documentation

- We updated the guide on [Preparing a Kubernetes cluster for the use of GPUs](https://docs.giantswarm.io/guides/kubernetes-gpu/) to make it easier to follow, especially on Azure.
- The [Releases and Changes](https://docs.giantswarm.io/changes/) mentioned above.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
