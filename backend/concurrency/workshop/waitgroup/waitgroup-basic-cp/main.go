package main

import "sync"

func testWG(output chan<- []bool) {
	// TODO: answer here
<<<<<<< HEAD
	var wg sync.WaitGroup
=======

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	done := make([]bool, 5)

	for i := 0; i < 5; i++ {
		// TODO: answer here
<<<<<<< HEAD
		wg.Add(1)
		go func(i int) {
			// TODO: answer here
			defer wg.Done()
=======
		go func(i int) {
			// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			done[i] = true
		}(i)
	}

	// TODO: answer here
<<<<<<< HEAD
	wg.Wait()
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	output <- done
}
