package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	// TODO: answer here
<<<<<<< HEAD
	c := make(chan int)
=======

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
<<<<<<< HEAD
			defer wg.Done()
			//kirim 1 ke channel
			// TODO: answer here
			c <- 1
=======
			//kirim 1 ke channel
			// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		}()
	}
	//mengubah nilai count menggunakan data dari channel

	// TODO: answer here
<<<<<<< HEAD
	go func() {
		for {
			count += <-c
		}
	}()
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count
}
