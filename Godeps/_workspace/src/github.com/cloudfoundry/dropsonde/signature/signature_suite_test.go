package signature_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnmarshaller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Signature Suite")
}
