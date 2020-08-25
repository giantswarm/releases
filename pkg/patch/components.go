package patch

import (
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
)

type ComponentPatch struct {
	Change Change `json:"change"`
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
		if component.Change.IsAddOrModify() {
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
		} else if component.Change == ChangeDelete {
			for j, existing := range result {
				if existing.Name == component.Name {
					result = append(result[:j], result[j+1:]...)
					break
				}
			}
		}
	}

	return result
}
