package main

import (
	// TODO: answer here
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//buat sebuah goroutine menggunakan anonymous function
//function print "selamat sore" dan ubah nilai called menjadi true
var _ = Describe("Goroutine", func() {
	It("can be used with anonymous function", func() {
		called := false
		// TODO: answer here
<<<<<<< HEAD
		go func() {
			called = true
			fmt.Println("selamat sore")
		}()
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		time.Sleep(10 * time.Millisecond)
		Expect(called).To(Equal(true))
	})
})
