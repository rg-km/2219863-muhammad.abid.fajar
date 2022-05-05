package main

import (
	"fmt"
	"time"
)

type character struct {
	name            string
	activity        string
	defaultActivity string
}

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//sama seperti checkpoint sebelumnya
		//tetapi sekarang menambahkan default activity yang menjalankan
		//fmt.Printf("%s %s\n", c.name, c.defaultActivity)
		//time.Sleep(100 * time.Millisecond)
		// TODO: answer here
<<<<<<< HEAD
		select {
		case c.activity = <-attackInput:
			fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		case c.activity = <-movementInput:
			fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		default:
			fmt.Printf("%s %s\n", c.name, c.defaultActivity)
			time.Sleep(100 * time.Millisecond)
		} 
=======

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}
