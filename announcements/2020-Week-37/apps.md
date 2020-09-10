# :zap: This Week in Provider-Independent Announcements (September 10, 2020):zap:

**Security Announcements**

1. [CVE-2020-14386](https://docs.google.com/document/d/1WJ_T9M9orEb0_bIQo4r827GNUm5GgK1Uz6Xbe9Ynaaw/edit?usp=sharing) affecting the Linux kernel is mitigated by the default Giant Swarm configuration. If you run privileged pods or pods with the `CAP_NET_RAW` capability enabled, please review this advisory.

**Giant Swarm Catalog (Managed Apps)**

1. [Kong app v0.9.0](https://github.com/giantswarm/kong-app/blob/master/CHANGELOG.md#090---2020-08-25) upgrades default Kong version to 2.1.3, updates other dependencies, syncs with upstream chart 1.8.0.
2. [Kong app v0.9.1](https://github.com/giantswarm/kong-app/blob/master/CHANGELOG.md#091---2020-09-07) splits image repository value to allow switching registry i.e. instead of default Quay use Alibaba by default in China.
3. [NGINX IC app v1.9.2](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#192---2020-09-02) upgrades ingress-nginx from v0.34.1 to v0.35.0.
4. [EFK Stack app v0.3.0](https://github.com/giantswarm/efk-stack-app/blob/master/CHANGELOG.md#030---2020-09-04) upgrades to Open Distro v1.9.0.

**Playground Catalog (Playground Apps)**

1. [Strimzi Kafka Operator app v0.1.0](https://github.com/giantswarm/strimzi-kafka-operator-app/blob/master/CHANGELOG.md#010---2020-08-14) upgrades upstream strimzi-kafka-operator from v0.18.0 to v0.19.0.
2. [Jaeger Operator app v0.2.0](https://github.com/giantswarm/jaeger-operator-app/blob/master/CHANGELOG.md#020---2020-09-01) upgrades jaeger-operator version to v1.18.0.


---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
