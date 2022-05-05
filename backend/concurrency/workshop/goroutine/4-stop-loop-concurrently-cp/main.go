package main

<<<<<<< HEAD
import "fmt"

//bagaiman cara menghentikan loop ?
func stopLoop(loop *bool) {
	// TODO: answer here
	go func() {
		*loop = false
		fmt.Println(loop)
	}()
=======
//bagaiman cara menghentikan loop ?
func stopLoop(loop *bool) {
	// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
