package main

func squareWorker(input <-chan int, output chan<- int) {
	//lakukan for range loop
	// TODO: answer here
<<<<<<< HEAD
	// for {
	// 	num := <-input      // menerima angka
	// 	output <- num * num // mengirim hasil
	// }
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func receiver(output chan<- int) {
	input := make(chan int) // mengirim 0-10 ke squareWorker
	go squareWorker(input, output)
	for i := 0; i < 10; i++ {
<<<<<<< HEAD
		var res int
		// TODO: answer here
		res = i * i
		output <- res
	}

=======
		//kirim nilai i ke channel
		// TODO: answer here
	}
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
