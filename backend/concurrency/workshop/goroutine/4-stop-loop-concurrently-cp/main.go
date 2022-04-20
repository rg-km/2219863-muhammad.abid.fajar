package main

import "fmt"

//bagaiman cara menghentikan loop ?
func stopLoop(loop *bool) {
	// TODO: answer here
	go func() {
		*loop = false
		fmt.Println(loop)
	}()
}
