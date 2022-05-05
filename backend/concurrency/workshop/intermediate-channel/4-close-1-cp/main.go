package main

//numberWorker digunakan untuk mengirim data ke channel input
func numberWorker(input chan int) {
	for i := 0; i < 10; i++ {
		input <- i
	}
	// TODO: answer here
<<<<<<< HEAD
	close(input)
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func receiver(output chan int) {
	input := make(chan int, 10)
	go numberWorker(input)
	for number := range input {
		output <- number
	}
	// TODO: answer here
<<<<<<< HEAD
	close(output)
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
