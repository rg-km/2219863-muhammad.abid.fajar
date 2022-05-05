package main

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

// Buat string JSON dengan hasil seperti berikut
// [{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]

type Meja struct {
	// TODO: answer here
<<<<<<< HEAD
	Jenis     string `json:"jenis"`
	Warna     string `json:"warna"`
	Jumlah    int    `json:"jumlah"`
	Deskripsi string `json:"deskripsi"`
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
<<<<<<< HEAD

	a, err := json.Marshal(m.MejaMeja)
	if err != nil {
		return err.Error()
	} else {
		return string(a)
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func NewMeja(m Items) Items {
	return m
}
