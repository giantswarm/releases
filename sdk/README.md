# Releases SDK

Releases SDK is a Go module for working with new Giant Swarm workload cluster releases.

> Note: The SDK can be used only for the workload cluster releases for new product based on Cluster API project, so it
> does not support working with the old vintage releases.

## 1. Using SDK

Add SDK to your project:

```shell
go get github.com/giantswarm/releases/sdk
```

## 2. CustomResourceDefinitions

### 2.1. `v1alpha1` API types

- `Release` represents a Giant Swarm workload cluster release.

Usage:
```go
package mypackage

import (
    rsdk "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

func main() {
    var release rsdk.Release
}
```
