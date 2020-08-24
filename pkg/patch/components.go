package patch

import (
	"fmt"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
)

type ComponentPatch struct {
	Change string `json:"change"`
	Name   string `json:"name"`

	Catalog               *string `json:"catalog,omitempty"`
	Reference             *string `json:"reference,omitempty"`
	ReleaseOperatorDeploy *bool   `json:"releaseOperatorDeploy,omitempty"`
	Version               *string `json:"version"`
}

func patchComponent(base v1alpha1.ReleaseSpecComponent, patch ComponentPatch) v1alpha1.ReleaseSpecComponent {
	if patch.Version != nil {
		base.Version = *patch.Version
	}
	if patch.Catalog != nil {
		base.Catalog = *patch.Catalog
	}
	if patch.Reference != nil {
		base.Reference = *patch.Reference
	}
	if patch.ReleaseOperatorDeploy != nil {
		base.ReleaseOperatorDeploy = *patch.ReleaseOperatorDeploy
	}
	return base
}

func patchComponents(base []v1alpha1.ReleaseSpecComponent, patch []ComponentPatch) []v1alpha1.ReleaseSpecComponent {
	var result []v1alpha1.ReleaseSpecComponent
	if len(base) > 0 {
		result = append(result, base...)
	}

	for _, component := range patch {
		if component.Change == "A" || component.Change == "M" {
			exists := false
			for j, existing := range result {
				if existing.Name == component.Name {
					result[j] = patchComponent(existing, component)
					exists = true
					break
				}
			}

			if !exists {
				newComponentFromPatch := patchComponent(v1alpha1.ReleaseSpecComponent{}, component)
				result = append(result, newComponentFromPatch)
			}

			if component.Change == "A" && exists {
				fmt.Println("patch adds " + component.Name + " but component already exists in base, updated version instead")
			} else if component.Change == "M" && !exists {
				fmt.Println("patch modifies " + component.Name + " but component not found in base, added component instead")
			}
		} else if component.Change == "D" {
			deleted := false
			for j, existing := range result {
				if existing.Name == component.Name {
					result = append(result[:j], result[j+1:]...)
					deleted = true
					break
				}
			}

			if !deleted {
				fmt.Println("patch deletes " + component.Name + " but component not found in base, ignoring")
			}
		}
	}

	return result
}
