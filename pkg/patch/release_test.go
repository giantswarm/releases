package patch

import (
	"strconv"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func Test_Patch_Release(t *testing.T) {
	baseDate := metav1.Date(2020, 8, 23, 12, 39, 32, 0, time.UTC)
	patchDate := metav1.Date(2020, 8, 24, 12, 39, 32, 0, time.UTC)

	testCases := []struct {
		name             string
		base             v1alpha1.Release
		patch            ReleasePatch
		expectedPrevious v1alpha1.Release
		expectedPatched  v1alpha1.Release
	}{
		{
			name: "case 0: simple patch",
			base: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.2.3",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(baseDate),
					State: "active",
				},
			},
			patch: ReleasePatch{
				Date:    patchDate,
				Version: *semver.MustParse("v1.2.4"),
			},
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.2.3",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(baseDate),
					State: "deprecated",
				},
			},
			expectedPatched: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.2.4",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(patchDate),
					State: "active",
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			previous, patched := Apply(tc.base, tc.patch)

			if diff := cmp.Diff(previous, tc.expectedPrevious); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(patched, tc.expectedPatched); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_Patch_YAML(t *testing.T) {
	baseDate := metav1.Date(2020, 8, 23, 12, 39, 32, 0, time.UTC)
	patchDate := metav1.Date(2020, 8, 24, 13, 4, 31, 0, time.UTC)

	testCases := []struct {
		name             string
		base             v1alpha1.Release
		patch            string
		expectedPrevious v1alpha1.Release
		expectedPatched  v1alpha1.Release
	}{
		{
			name: "case 0: simple patch",
			base: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.0.0",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(baseDate),
					State: "active",

					Components: []v1alpha1.ReleaseSpecComponent{
						{
							Name: "test-component",

							Catalog:               testComponent.Catalog,
							Reference:             testComponent.Reference,
							ReleaseOperatorDeploy: testComponent.ReleaseOperatorDeploy,
							Version:               testComponent.Version,
						},
					},
				},
			},
			patch: `date: "2020-08-24T13:04:31Z"
version: v1.0.1
components:
- name: test-component
  change: M
  version: v1.2.4
`,
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.0.0",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(baseDate),
					State: "deprecated",

					Components: []v1alpha1.ReleaseSpecComponent{
						{
							Name: "test-component",

							Catalog:               testComponent.Catalog,
							Reference:             testComponent.Reference,
							ReleaseOperatorDeploy: testComponent.ReleaseOperatorDeploy,
							Version:               testComponent.Version,
						},
					},
				},
			},
			expectedPatched: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v1.0.1",
				},
				Spec: v1alpha1.ReleaseSpec{
					Date:  timePtr(patchDate),
					State: "active",

					Components: []v1alpha1.ReleaseSpecComponent{
						{
							Name: "test-component",

							Catalog:               testComponent.Catalog,
							Reference:             testComponent.Reference,
							ReleaseOperatorDeploy: testComponent.ReleaseOperatorDeploy,
							Version:               "v1.2.4",
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			var patch ReleasePatch
			err := yaml.Unmarshal([]byte(tc.patch), &patch)
			if err != nil {
				t.Fatal(err)
			}

			previous, patched := Apply(tc.base, patch)

			if diff := cmp.Diff(previous, tc.expectedPrevious); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(patched, tc.expectedPatched); diff != "" {
				t.Error(diff)
			}
		})
	}
}
