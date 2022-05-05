package main

import "fmt"

func receiveBlock(output chan int) {
	c := make(chan int)
	result := 0
	go func() {
		fmt.Println("send to channel c")
		//kirim 1 ke channel c
		// TODO: answer here
<<<<<<< HEAD
		c <- 1
	}()

	//result menerima data dari channel c
	result = <-c
	// TODO: answer here
	output <- result
	fmt.Println(c) //agar variabel c digunakan
}
=======
	}()

	//result menerima data dari channel c
	// TODO: answer here
	output <- result
	fmt.Println(c) //agar variabel c digunakan
}
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
