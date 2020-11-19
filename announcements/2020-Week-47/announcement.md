# :zap: Announcements for this week (November 20, 2020) :zap:

First things first, we offer a new and powerful way to **follow releases, changes, and improvements**. Please check out our public [Releases and Changes](https://docs.giantswarm.io/changes/) section, which is a central place to look at new versions of our apps, new tenant cluster releases, UI improvements and even documentation changes.

Don't hesitate to tell us how this works for you and what you'd like to see improved.

## kubectl gs

- The [`template cluster`](https://docs.giantswarm.io/reference/kubectl-gs/template-cluster/) and [`template nodepool`](https://docs.giantswarm.io/reference/kubectl-gs/template-nodepool/) commands for creating cluster and node pool manifests no longer provide the `--region` flag, as this value is now filled in automatically by the Control Plane Kubernetes API.

## Documentation

- We updated the guide on [Preparing a Kubernetes cluster for the use of GPUs](https://docs.giantswarm.io/guides/kubernetes-gpu/) to make it easier to follow, especially on Azure.
- The [Releases and Changes](https://docs.giantswarm.io/changes/) mentioned above.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
