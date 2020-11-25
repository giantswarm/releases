## ⚡️ Tenant Cluster Release 11.2.3 for Azure ⚡️

This release improves the way upgrades handle Azure VMSS Instances to make them more reliable in the Germany West Central location. This does not affect any existing functionality.

To upgrade, please contact your Solution Engineer.

Below, you can find more details on components that were changed with this release.

### azure-operator v3.0.6

- Removed usage of LastModelApplied field of the VMSS Instance type
