package patch

import (
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
)

type AppPatch struct {
	Change Change `json:"change"`
	Name   string `json:"name"`

	ComponentVersion *string `json:"componentVersion,omitempty"`
	Version          *string `json:"version"`
}

func patchApp(base v1alpha1.ReleaseSpecApp, patch AppPatch) v1alpha1.ReleaseSpecApp {
	if patch.Version != nil {
		base.Version = *patch.Version
	}
	if patch.ComponentVersion != nil {
		base.ComponentVersion = *patch.ComponentVersion
	}
	return base
}

func patchApps(base []v1alpha1.ReleaseSpecApp, patch []AppPatch) []v1alpha1.ReleaseSpecApp {
	var result []v1alpha1.ReleaseSpecApp
	if len(base) > 0 {
		result = append(result, base...)
	}

	for _, app := range patch {
		if app.Change.IsAddOrModify() {
			exists := false
			for j, existing := range result {
				if existing.Name == app.Name {
					result[j] = patchApp(existing, app)
					exists = true
					break
				}
			}

			if !exists {
				newAppFromPatch := patchApp(v1alpha1.ReleaseSpecApp{}, app)
				result = append(result, newAppFromPatch)
			}
		} else if app.Change == ChangeDelete {
			for j, existing := range result {
				if existing.Name == app.Name {
					result = append(result[:j], result[j+1:]...)
					break
				}
			}
		}
	}

	return result
}
