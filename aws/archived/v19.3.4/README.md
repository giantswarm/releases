# :zap: Giant Swarm Release v19.3.4 for AWS :zap:

This is a security release featuring latest version of Flatcar Container Linux.

## Change details


### containerlinux [3815.2.2](https://www.flatcar-linux.org/releases/#release-3815.2.2)

 _Changes since **Stable 3815.2.1**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-28746](https://nvd.nist.gov/vuln/detail/CVE-2023-28746), [CVE-2023-47233](https://nvd.nist.gov/vuln/detail/CVE-2023-47233), [CVE-2023-52639](https://nvd.nist.gov/vuln/detail/CVE-2023-52639), [CVE-2023-6270](https://nvd.nist.gov/vuln/detail/CVE-2023-6270), [CVE-2023-7042](https://nvd.nist.gov/vuln/detail/CVE-2023-7042), [CVE-2024-22099](https://nvd.nist.gov/vuln/detail/CVE-2024-22099), [CVE-2024-23307](https://nvd.nist.gov/vuln/detail/CVE-2024-23307), [CVE-2024-24861](https://nvd.nist.gov/vuln/detail/CVE-2024-24861), [CVE-2024-26584](https://nvd.nist.gov/vuln/detail/CVE-2024-26584), [CVE-2024-26585](https://nvd.nist.gov/vuln/detail/CVE-2024-26585), [CVE-2024-26642](https://nvd.nist.gov/vuln/detail/CVE-2024-26642), [CVE-2024-26651](https://nvd.nist.gov/vuln/detail/CVE-2024-26651), [CVE-2024-26654](https://nvd.nist.gov/vuln/detail/CVE-2024-26654), [CVE-2024-26659](https://nvd.nist.gov/vuln/detail/CVE-2024-26659), [CVE-2024-26686](https://nvd.nist.gov/vuln/detail/CVE-2024-26686), [CVE-2024-26700](https://nvd.nist.gov/vuln/detail/CVE-2024-26700), [CVE-2024-26809](https://nvd.nist.gov/vuln/detail/CVE-2024-26809))
 - Downgraded xz-utils to 5.4.2 as precaution even though Flatcar is not affected of the SSH backdoor ([CVE-2024-3094](https://nvd.nist.gov/vuln/detail/CVE-2024-3094))
 - openssh ([CVE-2023-48795](https://nvd.nist.gov/vuln/detail/CVE-2023-48795), [CVE-2023-51384](https://nvd.nist.gov/vuln/detail/CVE-2023-51384), [CVE-2023-51385](https://nvd.nist.gov/vuln/detail/CVE-2023-51385))
 
 #### Bug fixes:
 
 - Disabled user-configdrive.service on OpenStack when config drive is used, which caused the hostname to be overwritten. The coreos-cloudinit.service unit already runs on OpenStack if the system is not configured via ignition. ([Flatcar#1385](https://github.com/flatcar/Flatcar/issues/1385))
 - Fixed `toolbox` to prevent mounted `ctr` snapshots from being garbage-collected ([toolbox#9](https://github.com/flatcar/toolbox/pull/9))
 
 #### Changes:
 
 - Disabled real-time priority for multipathd as it prevents the cgroups2 cpu controller from working. ([scripts#1771](https://github.com/flatcar/scripts/pull/1771))
 - SDK: Unified qemu image formats, so that the `qemu_uefi` build target provides the regular `qemu` and the `qemu_uefi_secure` artifacts ([scripts#1847](https://github.com/flatcar/scripts/pull/1847))
 
 #### Updates:
 
 - Linux ([6.1.85](https://lwn.net/Articles/969355) (includes [6.1.84](https://lwn.net/Articles/968254), [6.1.83](https://lwn.net/Articles/966759), [6.1.82](https://lwn.net/Articles/965607)))
 - ca-certificates ([3.99](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_99.html))
 - openssh ([9.6p1](https://www.openssh.com/releasenotes.html#9.6p1))



