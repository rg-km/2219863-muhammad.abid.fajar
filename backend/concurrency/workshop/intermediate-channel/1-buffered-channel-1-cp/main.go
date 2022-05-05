package main

//buffered channel tidak blocking selama buffer belum full
func testBuffer(output chan int) {
	input := make(chan int, 4)
	input <- 0
	input <- 1
	input <- 2
	input <- 3

	sum := 0
	for i := 0; i < cap(input); i++ {
		// TODO: answer here
<<<<<<< HEAD
		sum += i
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}

	output <- sum
}
