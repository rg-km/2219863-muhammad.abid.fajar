package main

import (
	"fmt"
)

func communicate(output chan string) {
	c := make(chan string)
	waitGoroutine := make(chan struct{}) //hanya digunakan untuk menunggu
	greetSteve := ""
	greet := func(name string) {
		fmt.Println("send to channel c")
		//kirim "hello"+name ke channel
		// TODO: answer here
<<<<<<< HEAD
		go func() {
			c <- "hello " + name
		}()
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
	go greet("steve")

	receiveGreet := func() {
		// TODO: answer here
<<<<<<< HEAD
		greetSteve = <-c
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		waitGoroutine <- struct{}{}
	}
	go receiveGreet()

	<-waitGoroutine
	output <- greetSteve
	fmt.Println(c) //agar variabel c digunakan
}
