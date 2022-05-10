package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Is", func() {
	// peopleAge := make([]string, 0)
	// BeforeEach(func() {
	// 	peopleAge = []string{
	// 		"0. Books  is: The Eye of the World",
	// 		"1. Books  is: The Great Hunt",
	// 		"2. Books  is: The Dragon Reborn",
	// 		"Panic error terdeteksi: runtime error: index out of range [3] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [4] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [5] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [6] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [7] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [8] with length 3",
	// 		"Panic error terdeteksi: runtime error: index out of range [9] with length 3",
	// 		"We handled the panic",
	// 	}
	// })

	Describe("test", func() {
		It("test", func() {
			r := 10
			Expect(r).To(Equal(r))
		})
	})
})
