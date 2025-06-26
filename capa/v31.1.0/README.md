# :zap: Giant Swarm Release v31.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v31.0.0

### Components

- Flatcar from v4152.2.3 to [v4230.2.0](https://www.flatcar-linux.org/releases/#release-4230.2.0)
- Kubernetes from v1.31.9 to [v1.31.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.31.md#v1.31.10)

### Apps

- coredns from v1.25.0 to v1.26.0
- k8s-dns-node-cache from v2.8.1 to v2.9.0


### coredns [v1.25.0...v1.26.0](https://github.com/giantswarm/coredns-app/compare/v1.25.0...v1.26.0)

#### Changed

- Update `coredns` image to [1.12.2](https://github.com/coredns/coredns/releases/tag/v1.12.2).

### k8s-dns-node-cache [v2.8.1...v2.9.0](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.8.1...v2.9.0)

#### Changed

- Upgrade application to version 1.26.4 (includes coredns 1.11.3)
- Increase ServiceMonitor's scrapping interval to 1m.
- Remove obsolete PSPs
