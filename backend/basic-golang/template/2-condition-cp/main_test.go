package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Describe("ExecuteToByteBuffer", func() {
		It("returns slice of bytes with correct wording", func() {
			account := Account{
				Name:    "Tony",
				Number:  "1002321",
				Balance: 1000,
			}
			b, err := ExecuteToByteBuffer(account)
<<<<<<< HEAD
			// fmt.Println(err)
			Expect(err).ShouldNot(HaveOccurred())
			// c := string(b)
			// fmt.Println(c)
			// Expect(string(b)).To(Equal(c))
=======
			Expect(err).ShouldNot(HaveOccurred())
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			Expect(string(b)).To(Equal("Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."))
		})
		It("return slice of bytes with empty state wording", func() {
			account := Account{
				Name:    "Tony",
				Number:  "1002321",
				Balance: 0,
			}
			b, err := ExecuteToByteBuffer(account)
<<<<<<< HEAD
			// fmt.Println(err)
			Expect(err).ShouldNot(HaveOccurred())

			// c := string(b)
			// fmt.Println(c)
			// Expect(string(b)).To(Equal(c))
=======
			Expect(err).ShouldNot(HaveOccurred())
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			Expect(string(b)).To(Equal("Akun Tony dengan Nomor 1002321 tidak memiliki saldo."))
		})
	})
})
