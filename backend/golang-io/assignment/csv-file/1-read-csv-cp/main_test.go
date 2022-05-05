package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("CSVToMap", func() {
		It("Should return a map with structure like this `question:answer`", func() {
			data := make(map[string]string)
			data, err := CSVToMap(data, "questions.csv")
			Expect(err).To(BeNil())
<<<<<<< HEAD
			// data["Nama ibukota indonesia?"] = "Jakarta"
			// fmt.Printf(data["Nama ibukota indonesia?"])
=======

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			Expect(data["Nama ibukota indonesia?"]).To(Equal("Jakarta"))
			Expect(data["Siapakah presiden pertama Indonesia?"]).To(Equal("Soekarno"))
		})
	})
})
