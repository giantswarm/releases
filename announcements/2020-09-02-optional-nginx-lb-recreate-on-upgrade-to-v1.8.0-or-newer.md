# Known Issue: Potential ingress downtime when upgrading to NGINX IC v1.8.0+

We'd like to inform you about a known issue when upgrading NGINX ingress controller (IC) optional app to v1.8.0+ on **AWS** and **Azure** clusters - NGINX IC LoadBalancer Service may get renamed/recreated with the app upgrade.

In older releases the Service name was hardcoded to `nginx-ingress-controller`. As of v1.8.0, to ensure its uniqueness for multiple NGINX ICs per cluster support, the LoadBalancer Service name was made to be dynamic, derived from Helm release i.e. App Custom Resource (CR) name.
Therefore, NGINX IC LoadBalancer Service will be replaced by a new one for every NGINC IC App CR whose name is not `nginx-ingress-controller`.

When NGINX IC LoadBalancer Service gets recreated, cloud service provider (CSP) load balancer behind it gets recycled as well.
It can take minute or so for ingress DNS records to be updated by external-dns and change propagated to clients.
During that time there's ingress traffic downtime, since clients still resolve old no longer present CSP load balancer.

**Please take the potential ingress downtime into consideration when planning the NGINX IC App upgrade from older to v1.8.0+.**

To make sure the downtime is shortest possible, external-dns availability is important precondition.
In recent platform releases (Azure v12.0.2, and AWS v12.1.4 and v11.5.4) we've improved external-dns monitoring and alerting.

**tl;dr: Therefore, before upgrading NGINX IC optional app to v1.8.0+, please make sure that your cluster has been upgraded to the latest platform release.**
If you have any further questions please do not hesitate to reach out for support.
