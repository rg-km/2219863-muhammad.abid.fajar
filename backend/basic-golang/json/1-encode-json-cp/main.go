package main

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
<<<<<<< HEAD
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
<<<<<<< HEAD
	mejaJSON, _ := json.Marshal(m)
	return string(mejaJSON)

}
=======
}

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
func NewMeja(m Meja) Meja {
	return m
}
