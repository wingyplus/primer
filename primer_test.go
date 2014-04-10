package primer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/primer"
	"testing"
)

func TestPrimer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Primer Spec")
}

var _ = Describe("Primer", func() {
	It("should be return [1] when number is 1", func() {
		Expect(primer.PrimeFactor(1)).To(Equal([]int{}))
	})
	It("should be return [2] when number is 2", func() {
		Expect(primer.PrimeFactor(2)).To(Equal([]int{2}))
	})
	It("should be return [3] when number is 3", func() {
		Expect(primer.PrimeFactor(3)).To(Equal([]int{3}))
	})
	It("should be return [2,2] when number is 4", func() {
		Expect(primer.PrimeFactor(4)).To(Equal([]int{2, 2}))
	})
	It("should be return [5] when number is 5", func() {
		Expect(primer.PrimeFactor(5)).To(Equal([]int{5}))
	})
	It("should be return [2,3] when number is 6", func() {
		Expect(primer.PrimeFactor(6)).To(Equal([]int{2, 3}))
	})
	It("should be return [7] when number is 7", func() {
		Expect(primer.PrimeFactor(7)).To(Equal([]int{7}))
	})
	It("should be return [2,2,2] when number is 8", func() {
		Expect(primer.PrimeFactor(8)).To(Equal([]int{2, 2, 2}))
	})
	It("should be return [3,3] when number is 9", func() {
		Expect(primer.PrimeFactor(9)).To(Equal([]int{3, 3}))
	})
	It("should be return [2,5] when number is 10", func() {
		Expect(primer.PrimeFactor(10)).To(Equal([]int{2, 5}))
	})
	It("should be return [11] when number is 11", func() {
		Expect(primer.PrimeFactor(11)).To(Equal([]int{11}))
	})
	It("should be return [2,2,3] when number is 12", func() {
		Expect(primer.PrimeFactor(12)).To(Equal([]int{2, 2, 3}))
	})
	It("should be return [13] when number is 13", func() {
		Expect(primer.PrimeFactor(13)).To(Equal([]int{13}))
	})
})
