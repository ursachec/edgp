package edgp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEdgp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Edgp Suite")
}
