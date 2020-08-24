package patch

import (
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
)

var (
	testApp = v1alpha1.ReleaseSpecApp{
		ComponentVersion: "v2.3.4",
		Version:          "v1.2.3",
	}
)

func Test_App_Patch(t *testing.T) {
	patchedApp := v1alpha1.ReleaseSpecApp{
		ComponentVersion: "v2.3.5",
		Version:          "v1.2.4",
	}
	testCases := []struct {
		name     string
		base     v1alpha1.ReleaseSpecApp
		patch    AppPatch
		expected v1alpha1.ReleaseSpecApp
	}{
		{
			name: "case 0: patch all",
			base: testApp,
			patch: AppPatch{
				ComponentVersion: stringPtr(patchedApp.ComponentVersion),
				Version:          stringPtr(patchedApp.Version),
			},
			expected: patchedApp,
		},
		{
			name: "case 1: patch version only",
			base: testApp,
			patch: AppPatch{
				Version: stringPtr(patchedApp.Version),
			},
			expected: v1alpha1.ReleaseSpecApp{
				ComponentVersion: testApp.ComponentVersion,
				Version:          patchedApp.Version,
			},
		},
		{
			name:     "case 3: empty patch",
			base:     testApp,
			patch:    AppPatch{},
			expected: testApp,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			actual := patchApp(tc.base, tc.patch)
			if diff := cmp.Diff(actual, tc.expected); diff != "" {
				t.Error(diff)
			}
		})
	}
}
