package edgp_test

import (
	"crypto/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ursachec/edgp/edgp"
)

var _ = Describe("edgp", func() {
	Describe("Generating throws", func() {
		entries := map[int64]string{
			11: "one",
			12: "two",
			21: "three",
			22: "four",
		}

		dicePerThrow := 2
		maxDiceValue := 2
		t := edgp.NewThrower(rand.Reader, entries, dicePerThrow, maxDiceValue)
		numThrows := 6

		Context("Simple throw", func() {
			It("Should return a slice with numThrows entries", func() {
				entries, err := t.Throw(numThrows)
				Expect(err).To(BeNil())
				Expect(len(entries)).To(Equal(numThrows))
			})
		})
	})
})
