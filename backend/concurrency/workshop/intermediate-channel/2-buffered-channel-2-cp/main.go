package main

import (
	"fmt"
	"time"
)

var person = "andi"
var names = []string{"budi", "toni", "adi", "ado", "alif", "yudi"}

//mengembalikan string, dimana `name` menyapa semua `names`
func greetAll(person string, names []string, output chan<- string) {
	// TODO: answer here
<<<<<<< HEAD
	for _, g := range names {
		output <- person + " say hello to " + g
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	fmt.Println("selesai mengirim")

}

// buat size buffered channel sesuai jumlah names
func testBufferedChannel(result chan<- string) {
	output := make(chan string) // TODO: replace this

	go greetAll(person, names, output)
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 6; i++ {
		greeting := <-output
		result <- greeting
	}
}

//goroutine greetAll dapat mengirim ke goroutine testBufferedChannel walaupun channel belum siap menerima
