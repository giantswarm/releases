# :zap: Giant Swarm Release v17.0.2 for AWS :zap:

This release downgrades the version of the Flatcar AMI from `3033.2.2` to `3033.2.0` due to a bug in version `3033.2.1` -> `3033.2.3` preventing successful boot on some EC2 instance type families. (Notably the `m4` instance types)

* [Giant Swarm roadmap issue](https://github.com/giantswarm/roadmap/issues/891)
* [Upstream bug issue](https://github.com/flatcar-linux/Flatcar/issues/665)

## Change details

### containerlinux [3033.2.0](https://www.flatcar-linux.org/releases/#release-3033.2.0)

