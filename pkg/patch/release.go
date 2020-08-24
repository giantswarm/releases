package patch

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReleasePatch struct {
	Date    metav1.Time    `json:"date"`
	Version semver.Version `json:"version"`

	Apps       []AppPatch       `json:"apps"`
	Components []ComponentPatch `json:"components"`
}

// Apply patch to the given base release and return the resulting previous and patched releases.
func Apply(base v1alpha1.Release, patch ReleasePatch) (previous v1alpha1.Release, patched v1alpha1.Release) {
	previous = base
	previous.Spec.State = "deprecated"

	patched = v1alpha1.Release{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("v%s", patch.Version.String()),
		},
		Spec: v1alpha1.ReleaseSpec{
			Apps:       patchApps(base.Spec.Apps, patch.Apps),
			Components: patchComponents(base.Spec.Components, patch.Components),
			Date:       timePtr(patch.Date),
			State:      "active",
		},
		Status: v1alpha1.ReleaseStatus{},
	}

	return
}
