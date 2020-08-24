package patch

import (
	"fmt"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
)

type AppPatch struct {
	Change string `json:"change"`
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
		if app.Change == "A" || app.Change == "M" {
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

			if app.Change == "A" && exists {
				fmt.Println("patch adds " + app.Name + " but app already exists in base, modified existing instead")
			} else if app.Change == "M" && !exists {
				fmt.Println("patch modifies " + app.Name + " but app not found in base, added app instead")
			}
		} else if app.Change == "D" {
			deleted := false
			for j, existing := range result {
				if existing.Name == app.Name {
					result = append(result[:j], result[j+1:]...)
					deleted = true
					break
				}
			}

			if !deleted {
				fmt.Println("patch deletes " + app.Name + " but app not found in base, ignoring")
			}
		}
	}

	return result
}
