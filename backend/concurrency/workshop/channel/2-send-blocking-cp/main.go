package main

import "fmt"

func sendBlock(output chan bool) {
	c := make(chan bool)
	called := false
	go func() {
		fmt.Println("receive from main")
		//memberi called nilai dari channel c
		// TODO: answer here
<<<<<<< HEAD
		called = <-c
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}()

	//mengirim bool value true ke channel c
	// TODO: answer here
<<<<<<< HEAD
	c <- true
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	output <- called
	fmt.Println(c) //agar variabel c digunakan
}
