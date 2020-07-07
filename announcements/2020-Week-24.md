# This Week in Releases (June 12, 2020)

1. **CoreOS is replaced with Flatcar Container Linux** in [AWS v9.0.6](https://github.com/giantswarm/releases/tree/master/aws/v9.0.6)

2. Reliability of NGINX Ingress Controller is improved to **speed up recovery when NGINX gets overloaded** in [AWS v9.3.3](https://github.com/giantswarm/releases/tree/master/aws/v9.3.3), [AWS v9.0.6](https://github.com/giantswarm/releases/tree/master/aws/v9.0.6), [Azure v11.3.2](https://github.com/giantswarm/releases/tree/master/azure/v11.3.2), [KVM v11.3.2](https://github.com/giantswarm/releases/tree/master/kvm/v11.3.2), and [NGINX IC v1.6.12](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v1612-2020-06-04) (app can be installed from the Giant Swarm App Catalog)

3. **An issue causing cert-manager to be killed** for exceeding its memory limit, stopping SSL certificates from being automatically renewed **is fixed** in [AWS v11.3.3](https://github.com/giantswarm/releases/tree/master/aws/v11.3.3)

4. **Kong is upgraded** to the latest upstream Kong v1.6.1 and Kong Ingress Controller v0.9.0 in [Kong v0.8.1](https://github.com/giantswarm/kong-app/blob/master/CHANGELOG.md#v081---2020-06-02) (app can be installed from the Giant Swarm App Catalog)

5. [**Kubernetes GPU**](https://github.com/giantswarm/kubernetes-gpu/blob/master/README.md) can now be installed as an app from the Playground Catalog. For users that want to run GPU workloads, this app makes sure required special drivers (CUDA) are present and installed in the right version in the GPU machines.

---

Please help us improve our communication. **Was this helpful?** :thumbsup: :thumbsdown:
