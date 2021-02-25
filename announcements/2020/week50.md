# :zap: This Week in Provider-Independent Announcements (December 11, 2020):zap:

Please note that we are aligning some key terms in our product communication with the Kubernetes project's terminology:

- The term *management cluster* replaces the term *control plane*. The Kubernetes project uses the term control plane when dealing with the master nodes (as they were called formerly) of a cluster, so we are removing an ambiguity that confused people.
- The API formerly called *Control Plane Kubernetes API* is now called *Management API*.
- What we called *tenant cluster* is now called a *workload cluster*.

For a comprehensive and always-updated view of all new features and changes for apps, tenant clusters, UI, and documentation, please check out [Changes and Releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed Apps

- [NGINX IC v1.12.0](https://docs.giantswarm.io/changes/managed-apps/nginx-ingress-controller-app/v1.12.0/) adds options to disable updating of ingress loadbalancer status, as well as to set podAntiAffinity scheduling method via the values file.

## Web UI

We fixed a problem with the user account drop-down menu.

## Documentation

The search function will now make spelling suggestions, plus we updated and improved the content.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
