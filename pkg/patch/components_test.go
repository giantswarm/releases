package patch

import (
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
)

var (
	testComponent = v1alpha1.ReleaseSpecComponent{
		Catalog:               "catalog-1",
		Reference:             "v1.2.3-0",
		ReleaseOperatorDeploy: true,
		Version:               "v1.2.3",
	}
)

func Test_Component_Patch(t *testing.T) {
	patchedComponent := v1alpha1.ReleaseSpecComponent{
		Catalog:               "catalog-2",
		Reference:             "v1.2.4-0",
		ReleaseOperatorDeploy: false,
		Version:               "v1.2.4",
	}
	testCases := []struct {
		name     string
		base     v1alpha1.ReleaseSpecComponent
		patch    ComponentPatch
		expected v1alpha1.ReleaseSpecComponent
	}{
		{
			name: "case 0: patch all",
			base: testComponent,
			patch: ComponentPatch{
				Catalog:               stringPtr(patchedComponent.Catalog),
				Reference:             stringPtr(patchedComponent.Reference),
				ReleaseOperatorDeploy: boolPtr(patchedComponent.ReleaseOperatorDeploy),
				Version:               stringPtr(patchedComponent.Version),
			},
			expected: patchedComponent,
		},
		{
			name: "case 1: patch version only",
			base: testComponent,
			patch: ComponentPatch{
				Version: stringPtr(patchedComponent.Version),
			},
			expected: v1alpha1.ReleaseSpecComponent{
				Catalog:               testComponent.Catalog,
				Reference:             testComponent.Reference,
				ReleaseOperatorDeploy: testComponent.ReleaseOperatorDeploy,
				Version:               patchedComponent.Version,
			},
		},
		{
			name:     "case 3: empty patch",
			base:     testComponent,
			patch:    ComponentPatch{},
			expected: testComponent,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			actual := patchComponent(tc.base, tc.patch)
			if diff := cmp.Diff(actual, tc.expected); diff != "" {
				t.Error(diff)
			}
		})
	}
}
