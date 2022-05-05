package main

import (
	"fmt"
<<<<<<< HEAD
	"time"
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	// TODO: answer here
)

func multiply(a, b int, called *bool) {
	fmt.Printf("hasil perkalian %d dan %d: %d\n", a, b, a*b)
	*called = true
}

//jalankan fungsi multiply dengan goroutine
//hint: gunakan sesuatu yang blocking
func start(multiplyCalled *bool) {
	go multiply(4, 5, multiplyCalled)
	// TODO: answer here
<<<<<<< HEAD
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main stop")
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
