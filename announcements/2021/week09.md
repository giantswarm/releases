# :zap: This week in provider-independent announcements (March 5, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[NGINX IC v1.15.0](https://docs.giantswarm.io/changes/managed-apps/nginx-ingress-controller-app/v1.15.0/) upgrades the container images (controller container to v0.44.0 and kube-webhook-certgen to v1.5.1). It also removes conflicting admission webhook api versions...

[Kong v1.2.0](https://docs.giantswarm.io/changes/managed-apps/kong-app/v1.2.0/) upgrades Kong to v2.3...

[Loki v0.1.1-alpha](https://docs.giantswarm.io/changes/managed-apps/loki-app/v0.1.1-alpha/) adds annotation to route alerts to the apps team, in preparation for offering Loki as a managed app in production.

## Apps supported with best-effort

[Jaeger Operator v0.2.3](https://docs.giantswarm.io/changes/playground-apps/jaeger-operator-app/v0.2.3/) upgrades to upstream v0.2.3.

[Fluent Logshipping v0.6.1](https://docs.giantswarm.io/changes/playground-apps/fluent-logshipping-app/v0.6.1/) fixes the output config format for AWS outputs, as well as mounts journald path and sets it correctly in fluent-bit config.

[Strimzi Kafka Operator v0.2.0](https://docs.giantswarm.io/changes/playground-apps/strimzi-kafka-operator-app/v0.2.0/) upgrades to upstream v0.21.1.

## gsctl

- The [gsctl create kubeconfig](https://docs.giantswarm.io/ui-api/gsctl/create-kubeconfig/) command now supports [Kubie](https://github.com/sbstp/kubie), which is a tool for safely setting the context for kubectl by spawning a new shell.
- The configuration path can now be customized using the GSCTL_CONFIG_DIR environment variable.

## Documentation

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
