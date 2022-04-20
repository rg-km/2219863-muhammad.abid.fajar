package main

import "fmt"

func squareWorker(input <-chan int, output chan<- int) {
	for i := 1; i <= 10; i++ {
		//mengirim ke channel output hasil	pangkat 2 nilai dari channel input
		// TODO: answer here
		msg := <-input
		msg = i * i
		fmt.Println(msg , " - test")
		output <- msg
	}
}

//mengirim 1-n angka ke squareWorker
func numberGenerator(n int, input chan<- int) {
	for i := 1; i <= n; i++ {
		// TODO: answer here
		fmt.Println(n)
		input <- n
		fmt.Println(input)
	}
}
