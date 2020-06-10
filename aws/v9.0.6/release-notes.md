## :zap: Giant Swarm Release 9.0.6 for AWS :zap:

This release [replaces CoreOS with Flatcar Container Linux](https://www.giantswarm.io/blog/time-to-catch-a-new-train-flatcar-linux).
CoreOS has gone [end-of-life](https://coreos.com/os/eol/) and is being rapidly phased out.
Flatcar is a compatible fork of CoreOS which receives ongoing support.
To continue receiving security updates and to minimize the effort needed to migrate in the future, we recommend upgrading to this release.

**Note to SEs when upgrading from 8.5.0 or 9.0.0:** Existing customer automation or processes that manage the configuration of coredns, nginx-ingress-controller, or cluster-autoscaler must be modified in order to work with the changed location and format of the *-user-values configmaps. Please see our docs on [Giant Swarm Release Versions: Versions that use the App Platform](https://docs.giantswarm.io/reference/release-versions/#versions-that-use-the-app-platform) for more details.

**Note for future 9.0.x releases:** Please include this note and the one above in all future 9.0.x releases.

Below, you can find more details on components that were changed with this release.

### aws-operator [v5.5.3](https://github.com/giantswarm/aws-operator/releases/tag/v5.5.3)
- Fixed issue with kube-proxy's `conntrackMaxPerCore` parameter not taking effect.

### Calico [v3.9.6](https://docs.projectcalico.org/archive/v3.9/release-notes/)
- Updated from Calico v3.9.1.
- Fixed IPv6 rogue router advertisement vulnerability CVE-2020-13597. See [security bulletin](https://www.projectcalico.org/security-bulletins/) for more information.

### Flatcar Container Linux [v2512.2.0](https://www.flatcar-linux.org/releases/#release-2512.2.0)
- Updated from CoreOS Container Linux 2191.5.0.
- Updated Linux Kernel to 4.19.124.

### Kubernetes [v1.15.12](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.15.md#changelog-since-v11511)
- Updated from v1.15.11.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.12](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v1612-2020-06-04))

- Made healthcheck probes configurable.
- Made liveness probe more resilient.
- Aligned labels using `app.kubernetes.io/name` instead of `k8s-app` where possible. `k8s-app` remains to be used for compatibility reasons as selectors are not modifiable without recreating the Deployment.
