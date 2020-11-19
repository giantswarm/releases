# :zap:  Tenant Cluster Release 9.0.0 for AWS :zap:

### Kubernetes v1.15.5
- Updated from v1.14.6 - [changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.15.md#kubernetes-v115-release-notes)
- Includes a fix for CVE-2019-11253 related to json/yaml decoding where large or malformed documents could consume excessive server resources. Request bodies for normal API requests (create/delete/update/patch operations of regular resources) are now limited to 3MB.

### Calico v3.9.1
- Updated from v3.8.2 - [changelog](https://docs.projectcalico.org/v3.9/release-notes/)

### CoreOS Container Linux v2191.5.0
- Updated from v2135.4.0 - [changelog](https://coreos.com/releases/#2191.5.0)

### etcd v3.3.15
- Updated from v3.3.13 - [changelog](https://github.com/etcd-io/etcd/blob/master/CHANGELOG-3.3.md#v3315-2019-08-19)

### kube-state-metrics v1.8.0 ([GS v0.6.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#v060))
- Updated from upstream `kube-state-metrics` v1.7.2 - [changelog](https://github.com/kubernetes/kube-state-metrics/blob/master/CHANGELOG.md#v180--2019-10-)

### nginx-ingress-controller v0.26.1 ([GS v1.0.0](https://github.com/giantswarm/kubernetes-nginx-ingress-controller/blob/master/CHANGELOG.md#100))
- Updated from upstream `ingress-nginx` v0.25.1 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0261)
- **Note** Includes the following breaking changes
  - The variable `$the_real_ip` was removed from template and default `log_format`.
  - The default value of configmap setting `proxy-add-original-uri-header` is now `false`.
  - When the setting `proxy-add-original-uri-header` is `true`, the ingress controller adds a new header `X-Original-Uri` with the value of NGINX variable `$request_uri`. In most of the cases this is not an issue but with request with long URLs it could lead to unexpected errors in the application defined in the `Ingress` `serviceName`, like issue https://github.com/kubernetes/ingress-nginx/issues/4593.

### coreDNS v1.6.4 ([GS v0.8.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v080))
- Updated from upstream `coredns` v1.6.2 - [changelog](https://coredns.io/2019/09/27/coredns-1.6.4-release/)

### Cluster Autoscaler v1.15.2 ([GS v0.9.0](https://github.com/giantswarm/cluster-autoscaler-app/blob/master/CHANGELOG.md#v090))
- Updated from upstream `Cluster Autoscaler` v1.14.3 - [changelog](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.15.2)
