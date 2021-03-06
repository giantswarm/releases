# :zap: This Week in AWS Releases (August 7, 2020):zap:

[AWS v11.5.1](https://github.com/giantswarm/releases/tree/master/aws/v11.5.1) fixes a problem with the calculation of the maximum number of Pods per node and improves resource cleanup after cluster deletion.

**Optional Apps Releases**
(Can be installed from App Catalogs)

1. [NGINX IC v1.8.4](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#184---2020-08-06) fixes NetworkPolicy templating, to allow Pod ingress traffic (Prometheus scrape requests) on same port that the metrics/monitoring service advertises. (Note: NGINX is Optional and not Pre-Installed in AWS 10+ clusters)

**Web Interface Releases**

1. [Happa v0.10.36](https://github.com/giantswarm/happa/releases/tag/v0.10.36) makes cluster upgrades from within happa safer. It makes it no longer possible to trigger an upgrade for non-ready or already upgrading clusters. Additionally, the cluster release selection will no longer contain any empty rows.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
