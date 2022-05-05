package main

<<<<<<< HEAD
import "fmt"

=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
func squareWorker(input <-chan int, output chan<- int) {
	for i := 1; i <= 10; i++ {
		//mengirim ke channel output hasil	pangkat 2 nilai dari channel input
		// TODO: answer here
<<<<<<< HEAD
		msg := <-input
		msg = i * i
		fmt.Println(msg , " - test")
		output <- msg
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}

//mengirim 1-n angka ke squareWorker
func numberGenerator(n int, input chan<- int) {
	for i := 1; i <= n; i++ {
		// TODO: answer here
<<<<<<< HEAD
		fmt.Println(n)
		input <- n
		fmt.Println(input)
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}
