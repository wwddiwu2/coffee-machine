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

	Describe("Brewing drinks", func() {
		Context("With enough cleanliness", func() {
			It("Should be brewed", func() {
				Ω(m.Brew(Espresso)).Should(BeNil())
				Ω(m.Brew(Americano)).Should(BeNil())
			})

			It("Should be dirtied correctly", func() {
				c := m.Cleanliness()
				m.Brew(Espresso)
				Ω(c).Should(Equal(m.Cleanliness() + 1))
				m.Brew(Americano)
				Ω(c).Should(Equal(m.Cleanliness() + 3))
			})

			It("Should set the status correctly", func() {
				Ω(m.Status()).Should(Equal(Ready))
				m.Brew(Espresso)
				Ω(m.Status()).Should(Equal(Ready))
				m.Brew(Americano)
				Ω(m.Status()).Should(Equal(Ready))
			})
		})

		Context("With lacking cleanliness (0)", func() {
			BeforeEach(func() {
				for i := 0; i < 5; i++ {
					Ω(m.Brew(Americano)).Should(BeNil())
				}
			})

			It("Should not be brewed", func() {
				Ω(m.Brew(Espresso)).ShouldNot(BeNil())
				Ω(m.Brew(Americano)).ShouldNot(BeNil())
			})

			It("Should not be dirtied below 0", func() {
				c := m.Cleanliness()
				m.Brew(Espresso)
				Ω(c).Should(Equal(m.Cleanliness()))
				m.Brew(Americano)
				Ω(c).Should(Equal(m.Cleanliness()))
				Ω(c).Should(BeNumerically(">=", 0))
			})

			It("Should set the status correctly", func() {
				Ω(m.Status()).Should(Equal(Ready))
				m.Brew(Espresso)
				Ω(m.Status()).Should(Equal(Ready))
				m.Brew(Americano)
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
