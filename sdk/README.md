# Releases SDK

Releases SDK is a Go module for working with Giant Swarm releases.

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
