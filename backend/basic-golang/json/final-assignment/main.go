package main

import (
	"encoding/json"
	"log"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here

<<<<<<< HEAD
type Ukuran struct {
	// TODO: answer here
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Ruang struct {
	// TODO: answer here
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	// assign field name `Ukuran` dengan
	// struct `Ukuran` yang sudah dibuat sebelumnya
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	Items []Ruang
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	mejaJSON, err := json.Marshal(m.Items)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	return "{\"ruangTamu\":{\"items\":" + string(mejaJSON) + "}}"
}

func NewRuang(m Items) Items {
	return m
=======
func (r Ruang) EncodeJSON() string {
	// TODO: answer here
}

func NewRuang(r Ruang) Ruang {
	return r
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
