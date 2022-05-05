package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {
	// TODO: answer here
<<<<<<< HEAD
	mu := &sync.Mutex{}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: answer here
			//kirim 1 ke channel
<<<<<<< HEAD
			mu.Lock()
			count++
			// TODO: answer here
			mu.Unlock()
=======
			count++
			// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		}()
	}
	wg.Wait()
	//mengubah nilai count menggunakan data dari channel
	output <- count
}
