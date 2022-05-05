package main

import "fmt"

type character struct {
	name            string
	activity        string
	defaultActivity string
}

//pada program ini dibuat character yang bernama john dengan defaulActivity mengawasi area dan activity kosong
//kita dapat mengubah activity dari john dengan cara mengirim ke channel
//untuk bergerak kita kirim ke movementInput, lalu untuk menyerang kita kirim ke attackInput

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//melakukan select
		//ketika menerima dari attackInput maka jalankan
		//fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		//ketika menerima dari movementInput maka jalankan
		//fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		// TODO: answer here
<<<<<<< HEAD
		select {
		case c.activity = <-attackInput:
			fmt.Printf("%s melakukan serangan %s\n", c.name, c.activity)
		case c.activity = <-movementInput:
			fmt.Printf("%s bergerak ke %s\n", c.name, c.activity)
		}
=======

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
}
