# :zap: This week in provider-independent announcements (February 05, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[External DNS v2.0.0](https://docs.giantswarm.io/changes/managed-apps/external-dns-app/v2.0.0/) allows customers to install and use External DNS for their own purpose, separate from the preinstalled app in AWS and Azure clusters. [External DNS v2.1.0](waiting for link) adds guardrails (dry-run mode, etc.) to help prevent misconfigurations. :warning: [Requires NGINX IC v1.14.0+](https://github.com/giantswarm/external-dns-app#limitations).

[NGINX IC v1.14.0](https://docs.giantswarm.io/changes/managed-apps/nginx-ingress-controller-app/v1.14.0/) is released to work well with customer's separate External DNS (see above).

[Cert Manager v2.4.2](https://docs.giantswarm.io/changes/managed-apps/cert-manager-app/v2.4.2/) provides the option to automatically delete secrets when the parent Certificate is deleted, reducing unnecessary alerts.

[K8s Initiator v0.9.2.](https://docs.giantswarm.io/changes/playground-apps/k8s-initiator-app/v0.9.2/) adds flexibility by adding support for user-specific tolerations and customer RBAC rules. We also added usage examples to increase user-friendliness. 

## Documentation

We re-organized our documentation to make it easier for you to find the content you need. The new left-hand-side menu gives you a better overview and allows to jump quickly. Please visit https://docs.giantswarm.io/ and let us know what you think!

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
