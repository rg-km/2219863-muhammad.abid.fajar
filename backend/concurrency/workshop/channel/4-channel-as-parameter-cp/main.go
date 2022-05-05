package main

import "fmt"

// fungsi ini digunakan untuk menyapa dua nama yang ada di parameter
// kalimat sapaan yang dibuat akan dikirim melalui channel
func greet(name1, name2 string, c chan string) {
	greeting := fmt.Sprintf("halo %s dan %s", name1, name2)
	fmt.Printf("kirim %s ke channel\n", greeting)
	// TODO: answer here
<<<<<<< HEAD
	c <- "halo " + name1 + " dan " + name2
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

//gunakan channel untuk komunikasi antar goroutine
func start(name1, name2 string, output chan string) {
	c := make(chan string)
	greeting := ""
	fmt.Printf("akan menyapa %s dan %s\n", name1, name2)

	// menjalankan goroutine greet
	// TODO: answer here
<<<<<<< HEAD
	go greet(name1, name2, output)

	//menerima data dari channel
	// TODO: answer here
	greeting = <-c
=======

	//menerima data dari channel
	// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	fmt.Println(c) //agar variabel c digunakan
	fmt.Printf("fungtion greet say: %s\n", greeting)
	output <- greeting
}
