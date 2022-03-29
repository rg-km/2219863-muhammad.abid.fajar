package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Buatlah dua data makanan dan minuman yaitu ayam goreng dan cola yang memiliki atribut
// Nama, Jenis, dan Harga.
// Ayam Goreng, Cepat saji, 20000
// Cola, Minuman, 7000

func GetMenu() []map[string]interface{} {
	var menu []map[string]interface{}

	makanan := make(map[string]interface{})
	makanan["Nama"] = "Ayam Goreng"
	makanan["Jenis"] = "Cepat saji"
	makanan["Harga"] = 20_000

	minuman := make(map[string]interface{})
	minuman["Nama"] = "Cola"
	minuman["Jenis"] = "Minuman"
	minuman["Harga"] = 7_000

	// type makanan struct {
	// 	Nama  string
	// 	Jenis string
	// 	Harga string
	// }

	// type minuman struct {
	// 	Nama  string
	// 	Jenis string
	// 	Harga string
	// }

	data := make(map[string]interface{})
	for k, v := range makanan {
		if _, ok := makanan[k]; ok {
			data[k] = v
		}
	}
	for k, v := range minuman {
		if _, ok := minuman[k]; ok {
			data[k] = v
		}
	}

	menu = append(menu, makanan, minuman)
	return menu
}

func main() {
	menu := GetMenu()

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
