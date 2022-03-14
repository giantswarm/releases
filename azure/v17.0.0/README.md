# :zap: Giant Swarm Release v17.0.0 for Azure :zap:

<< Add description here >>

## Change details

### containerlinux [3033.2.3](https://www.flatcar.org/releases/#release-3033.2.3)

Changes since Stable 3033.2.2
***Security fixes***
- Linux (CVE-2022-24448, CVE-2022-0617, CVE-2022-24959, CVE-2022-0492, CVE-2022-0516, CVE-2022-0435, CVE-2022-0487, CVE-2022-25375, CVE-2022-25258, CVE-2022-0847)
- go (CVE-2022-23806, CVE-2022-23772, CVE-2022-23773)
- ignition (CVE-2020-14040)
- expat (CVE-2022-25235, CVE-2022-25236, CVE-2022-25313, CVE-2022-25314, CVE-2022-25315)

***Bug fixes***

- Disabled the systemd-networkd settings ManageForeignRoutes and ManageForeignRoutingPolicyRules by default to ensure that CNIs like Cilium don’t get their routes or routing policy rules discarded on network reconfiguration events (Flatcar#620).
- Prevented hitting races when creating filesystems in Ignition, these races caused boot failures like fsck[1343]: Failed to stat /dev/disk/by-label/ROOT: No such file or directory when creating a btrfs root filesystem (ignition#35)
- Reverted the Linux kernel change to forbid xfrm id 0 for IPSec state because it broke Cilium (Flatcar#626, coreos-overlay#1682)

***Updates***

- Linux (5.10.102) (from 5.10.96)
- Go (1.17.7 (includes 1.17.6))
- expat (2.4.6)
- ca-certificates (3.75)


### cluster-autoscaler [1.22.2-gs4](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs4)

***Fixed***

- Updated to correct cluster-autoscaler version
- Use GS-built 1.22 image to deliver upstream unreleased fix kubernetes/autoscaler#4600
 
***Added***

- Added support for specifying balance-similar-node-groups flag
          


