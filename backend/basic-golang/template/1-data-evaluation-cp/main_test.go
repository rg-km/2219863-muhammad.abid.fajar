package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Describe("ExecuteToByteBuffer", func() {
		It("returns slice of bytes", func() {
			account := Account{
				Name:    "Tony",
				Number:  "1002321",
				Balance: 1000,
			}
			b, err := ExecuteToByteBuffer(account)
			Expect(err).ShouldNot(HaveOccurred())
<<<<<<< HEAD
			// fmt.Println(string(b))
			// Expect(string(b)).To(Equal(string(b)))
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			Expect(string(b)).To(Equal("Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."))
		})
	})
})
