package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUtilsSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = Describe("Utils", func() {
	Describe("Disjointed Range", func() {
		var Rge *DisjointRange
		BeforeEach(func() {
			Rge = NewDisjointRange()
		})
		It("Should add and merge ranges correctly", func() {
			Rge.AddRange(10, 15)
			Expect(Rge.SumRanges()).To(Equal(5))
			Expect(Rge.SumInclRanges()).To(Equal(6))
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(10, 15) // duplicate range
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(11, 15) // contained range start
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(10, 12) // contained range end
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(11, 14) // fully contained range
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(9, 16) // enclosing range
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(17, 18) // touching end (should merge)
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(7, 8) // touching start (should merge)
			Expect(Rge.CountRanges()).To(Equal(1))
			Rge.AddRange(20, 25) // separate range
			Expect(Rge.CountRanges()).To(Equal(2))
			Expect(Rge.SumInclRanges()).To(Equal(6 + 12))
			Expect(Rge.SumRanges()).To(Equal(5 + 11))
			Rge.AddRange(15, 20) // bridge ranges (should merge)
			Expect(Rge.CountRanges()).To(Equal(1))
			Expect(Rge.SumInclRanges()).To(Equal(19))
			Expect(Rge.SumRanges()).To(Equal(18))
		})
		It("should bridge + overlap ranges correctly", func() {
			Rge.AddRange(1, 3)
			Rge.AddRange(5, 7)
			Rge.AddRange(9, 11)
			Expect(Rge.CountRanges()).To(Equal(3))
			Rge.AddRange(2, 10)
			Expect(Rge.CountRanges()).To(Equal(1))
		})
		It("should handle start == end", func() {
			Rge.AddRange(5, 5)
			Expect(Rge.CountRanges()).To(Equal(1))
			Expect(Rge.SumRanges()).To(Equal(0))
			Expect(Rge.SumInclRanges()).To(Equal(1))
		})
	})
})
