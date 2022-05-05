package main

import (
	"fmt"
<<<<<<< HEAD
	"sync"
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	"time"
)

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

func newSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}

//kita bisa lihat hanya ada 10 goroutine yang berjalan dalam satu waktu
func main() {
<<<<<<< HEAD
	// mu := &sync.Mutex{}
	var wg sync.WaitGroup
	semaphore := newSemaphore(10) // kita ingin hanya ada 10 akses dalam satu waktu
	for i := 1; i <= 30; i++ {
		semaphore.Acquire()
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
=======
	semaphore := newSemaphore(10) // kita ingin hanya ada 10 akses dalam satu waktu
	for i := 1; i <= 30; i++ {
		semaphore.Acquire()
		go func(i int) {
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
			defer semaphore.Release()
			longRunningProcess(i)
		}(i)
	}
	//kapan terjadi blocking pada program ini ?
	// TODO: answer here
<<<<<<< HEAD
	wg.Wait()

=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(1 * time.Second) // melakukan sesuatu
}
