# :zap: This Week in Provider-Independent Announcements (October 16, 2020):zap:

## Giant Swarm Catalog (Managed Apps)

1. [NGINX Ingress Controller app v1.10.0](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#1100---2020-10-07) upgrades ingress-nginx-controller from `v0.35.0` to `v0.40.2` improving performance and reliability.
2. [Kong v1.0.0](https://github.com/giantswarm/kong-app/blob/master/CHANGELOG.md#100---2020-10-13) upgrades to upstream chart v1.11.0. (Note: Contains breaking changes from upstream. Solution Engineers are working with customers to prepare.)
3. [Prometheus Operator v0.4.0](https://github.com/giantswarm/prometheus-operator-app/blob/master/CHANGELOG.md#040---2020-10-15) upgrades to upstream chart v10.1.0. It also makes required changes related to the upstream chart name change from `prometheus-operator` (EOL in Nov 2020) to `kube-prometheus`
4. [Cert Manager v2.3.0](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#230---2020-10-02) brings the first stable release from upstream, v1.0.2.

## User Interface Releases

1. [Happa](https://github.com/giantswarm/happa/releases) displays `EOL` (end of life) for Kubernetes releases that no longer receive updates upstream, makes it more clear when a cluster is being created or upgrading, and fixes a few minor glitches.
2. [gsctl](https://github.com/giantswarm/gsctl/releases) also displays Kubernetes release end of life information, as mentioned above, in various commands.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
