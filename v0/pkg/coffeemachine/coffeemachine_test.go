package coffeemachine_test

import (
	. "github.com/mkqavi/coffee-machine/v0/pkg/coffeemachine"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coffeemachine", func() {
	var m Machine

	BeforeEach(func() {
		m = New()
	})

	Describe("Pouring drinks", func() {
		Context("With enough cleanliness", func() {
			It("Should be poured", func() {
				Ω(m.Pour(Espresso)).Should(BeNil())
				Ω(m.Pour(Americano)).Should(BeNil())
			})

			It("Should be dirtied correctly", func() {
				c := m.Cleanliness()
				m.Pour(Espresso)
				Ω(c).Should(Equal(m.Cleanliness() + 1))
				m.Pour(Americano)
				Ω(c).Should(Equal(m.Cleanliness() + 3))
			})

			It("Should set the status correctly", func() {
				Ω(m.Status()).Should(Equal(Ready))
				m.Pour(Espresso)
				Ω(m.Status()).Should(Equal(Ready))
				m.Pour(Americano)
				Ω(m.Status()).Should(Equal(Ready))
			})
		})

		Context("With lacking cleanliness (0)", func() {
			BeforeEach(func() {
				for i := 0; i < 50; i++ {
					Ω(m.Pour(Americano)).Should(BeNil())
				}
			})

			It("Should not be poured", func() {
				Ω(m.Pour(Espresso)).ShouldNot(BeNil())
				Ω(m.Pour(Americano)).ShouldNot(BeNil())
			})

			It("Should not be dirtied below 0", func() {
				c := m.Cleanliness()
				m.Pour(Espresso)
				Ω(c).Should(Equal(m.Cleanliness()))
				m.Pour(Americano)
				Ω(c).Should(Equal(m.Cleanliness()))
				Ω(c).Should(BeNumerically(">=", 0))
			})

			It("Should set the status correctly", func() {
				Ω(m.Status()).Should(Equal(Ready))
				m.Pour(Espresso)
				Ω(m.Status()).Should(Equal(Ready))
				m.Pour(Americano)
				Ω(m.Status()).Should(Equal(Ready))
			})
		})
	})

	Describe("Cleaning the machine", func() {
		It("Should set cleanliness to 100", func() {
			m.Clean()
			Ω(m.Cleanliness()).Should(Equal(uint8(100)))
		})

		It("Should set the status correctly", func() {
			Ω(m.Status()).Should(Equal(Ready))
			m.Clean()
			Ω(m.Status()).Should(Equal(Ready))
		})
	})
})
