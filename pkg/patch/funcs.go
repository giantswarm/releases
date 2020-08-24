package patch

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

func timePtr(t metav1.Time) *metav1.Time {
	return &t
}
