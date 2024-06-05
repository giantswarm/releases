package sdk_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSdk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sdk Suite")
}
