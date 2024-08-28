# :zap: Giant Swarm Release v19.0.1 for AWS :zap:

This is a patch release that provides improved performance for in-cluster DNS resolution.

## Change details


### coredns [1.18.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.18.1)

#### Added
- Add a new field `additionalLocalZones` which can be used to introduce more internal local zones, e.g. linkerd.
#### Changed
- Create a `coredns` zone for each cluster domain.
- Adjust the settings for upscaling HPA when hitting 60% CPU.
- Adjust the settings for downscaling HPA to 30 minutes.
- Adjust the min and max memory settings per Pod.
- Enable cache inconditionaly for `.` and local zones.
- Adjust the settings for upscaling HPA when hitting 80% Memory.
#### Fixed
- Remove fallthrough for reverse zones from kubernetes plugin.



### k8s-dns-node-cache-app [2.4.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.4.0)

#### Changed
- Upgrade application to version 1.22.23 (includes coredns 1.10)
- Enable TCP connections for external zones


### vertical-pod-autoscaler-app [3.5.4](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.4)

#### Changed
- Specified `failureThreshold` and `periodSeconds` for recommender's liveness probe.
- Upgrade dependency chart to 7.1.0.
- Upgrade VPA components to 0.14.0
