package main

import (
	"fmt"
	"sync"
)

const nWorkers = 3
const nRequesters = 5

type Work struct {
	data string
}

var mu = &sync.Mutex{}

var data = map[string]bool{}

func worker(in <-chan *Work, out chan<- *Work, number int) {
	for w := range in {
		w.data = fmt.Sprintf("worker %d accepted %s", number, w.data)
		// TODO: answer here
<<<<<<< HEAD
		// fmt.Printf(w.data)
		out <- &Work{w.data}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}

func createRequest(in chan<- *Work, number int) {
	for {
		// TODO: answer here
<<<<<<< HEAD
		in <- &Work{data: fmt.Sprintf("request from client %d", number)}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}

func receiver(out <-chan *Work) {
	for {
		//pada bagian ini masukkan data dari channel out
		//ke dalam map `data`
		//gunakan mutex untuk mengamankan penulisan ke map secara concurrent

<<<<<<< HEAD
		// <-out // TODO: replace this

		// TODO: answer here

		w := <-out
		mu.Lock()
		data[w.data] = true
		mu.Unlock()
		<-out
=======
		<-out // TODO: replace this
		<-out
		// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}
