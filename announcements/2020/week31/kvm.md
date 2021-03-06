# :zap: This Week in KVM Releases (July 31, 2020):zap:

[KVM v12.1.0](https://github.com/giantswarm/releases/blob/master/kvm/v12.1.0) supports multiple NGINX Ingress Controllers per tenant cluster, moves the management of NGINX IC NodePort Service from kvm-operator to NGINX IC app itself, and fixes Quay being a single point of failure.

[KVM v12.2.0](https://github.com/giantswarm/releases/blob/master/kvm/v12.2.0) turns NGINX Ingress Controller App from a pre-installed to an optional component on KVM. (As detailed in release notes, please ensure cluster is on KVM 12.1.x platform release first before upgrading to 12.2.0+)

**Optional Apps Releases**  
(Can be installed from App Catalogs)

1. [NGINX IC v1.8.3](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#183---2020-07-31) fixes controller RBAC permissions, granting "get" and "update" of leader election ConfigMap lock. (Note: NGINX is Optional and not Pre-Installed in KVM 12.2.x+ clusters)
2. [NGINX IC v1.8.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#181---2020-07-28) made node ports configurable for NodePort Service type. (Note: NGINX is Optional and not Pre-Installed in KVM 12.2.x+ clusters)
3. [NGINX IC v1.8.0](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#180---2020-07-24) supported multiple NGINX IC App installations per tenant cluster and dropped deprecated configuration properties. (Note: NGINX is Optional and not Pre-Installed in KVM 12.2.x+ clusters)
4. [Prometheus Operator v0.3.4](https://github.com/giantswarm/prometheus-operator-app/blob/master/CHANGELOG.md#034---2020-07-22) is upgraded to the latest version with minor bug fixes.

**Web Interface Releases**
1. [Happa v0.10.35](https://github.com/giantswarm/happa/releases/tag/v0.10.35) enables support for optional ingress controllers on KVM v12.2.0+, always offers the latest version of the NGINX IC app, and adds multiple UI improvements and bug fixes.


---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
