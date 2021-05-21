# :zap: This week in provider-independent announcements (May 21, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[Cloudflared v0.1.0](https://docs.giantswarm.io/changes/managed-apps/cloudflared-app/v0.1.0/) is now available. The app allows you to launch Cloudflare Argo Tunnels that route directly to services inside your cluster. Useful for clusters not running in a public cloud provider.

[Prometheus Operator v0.8.2](https://docs.giantswarm.io/changes/managed-apps/prometheus-operator-app/v0.8.2/) upgrades prometheus to [v2.26.1](https://github.com/prometheus/prometheus/releases/tag/v2.26.1) which contains a fix for [CVE-2021-29622](https://github.com/prometheus/prometheus/security/advisories/GHSA-vx57-7f4q-fpc7). The vulnerability allowed attackers to redirect to arbitray URLs through a specially crafted address (e.g.: http://127.0.0.1:9090/new/new<url>). Prior to that, [v0.8.1](https://docs.giantswarm.io/changes/managed-apps/prometheus-operator-app/v0.8.1/) eases installing on kubernetes >= 1.18 by enabling `ingressClass` and `pathType` by default.

[EFK v0.5.2](https://docs.giantswarm.io/changes/managed-apps/efk-stack-app/v0.5.2/) reverts upstream changes of deployment labels that cause upgrade conflicts.

[Cert Manager v2.7.1]() sets authoritative nameservers to `coredns` when using `dns01` ACME solver, in order to support DNS01 challenge introduced in [v2.6.0](https://docs.giantswarm.io/changes/managed-apps/cert-manager-app/v2.6.0/).

## User interfaces

[kubectl gs v1.28.0](https://docs.giantswarm.io/changes/kubectl-gs/kubectl-gs/v1.28.0/) adds `get appcatalogs` and `get apps` commands. It also now allows templating for China clusters.

We recently made many fixes and changes to [The Web UI](https://docs.giantswarm.io/changes/web-ui/) from v1.7.2 to 1.9.0 to improve the user experience, as well as prepare for new features. [CAPI? Management API?]

## Documentation

Many docs are updated and made more accurate including: [Advanced CoreDNS Configuration](https://docs.giantswarm.io/advanced/coredns/), [App CR's target namespace configuration](https://docs.giantswarm.io/app-platform/namespace-configuration/), and [Usage data recording in our web interface](https://docs.giantswarm.io/ui-api/web/usage-data/)

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
