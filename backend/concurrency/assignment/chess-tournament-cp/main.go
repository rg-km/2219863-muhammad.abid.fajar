package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}
func main() {
	startTournament()
}

func playMatch(i int) {
	fmt.Printf("match #%d started\n", i+1)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("match #%d finished\n", i+1)
}

func startTournament() {
	// TODO: answer here
<<<<<<< HEAD
	for i := 0; i < 10; i++ {
		fmt.Printf("match #%d started\n", i+1)
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("match #%d finished\n", i+1)
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
