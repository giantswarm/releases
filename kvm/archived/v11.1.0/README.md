## :zap: Tenant Cluster Release 11.1.0 for KVM :zap:

This release upgrades the NGINX Ingress Controller app to upstream v0.28.0. It contains several non-breaking changes that fix a few issues. We also enabled the metrics service. This allows prometheus-operator to scrape for metrics and improves monitoring of the app.

### net-exporter [v1.6.0](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#160-2020-01-29)

- Allow disabling TCP DNS check.

### nginx-ingress-controller v0.27.1 ([Giant Swarm App v1.2.0](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v120-2020-01-21))

- Add metrics Service for prometheus-operator support.
- Allow configuring which SSL/TLS protocols should be enabled.
