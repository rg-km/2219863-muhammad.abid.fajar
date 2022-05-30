package main

// perhatikan gambar unormal-example.png
// identifikasikan-lah semua attribut yang ada, lalu buatlah table seperti pada gambar  dengan nama table "unormal"
// hint : ada sekitar 13 attribut yg bisa dijadikan kolom

import (
	"database/sql"
	"fmt"
	"strconv"

	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Unormal struct {
	NoBon      int
	NamaBarang string
	Harga      int
	Jumlah     int
	Biaya      int
	SubTotal   int
	Discount   int
	Total      int
	Bayar      int
	Kembalian  int
	Kasir      string
	Tanggal    string
	Waktu      string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS unormal (
		no_bon VARCHAR(10),
		nama_produk VARCHAR(30),
		harga_produk INTEGER,
		jumlah_produk INTEGER,
		harga_total INTEGER,
		sub_total INTEGER,
		discount INTEGER,
		total INTEGER,
		bayar INTEGER,
		kembalian INTEGER,
		nama_kasir VARCHAR(30),
		tanggal VARCHAR(30),
		waktu VARCHAR(30)
	) ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		INSERT OR REPLACE INTO 
		unormal (no_bon, nama_produk, harga_produk, jumlah_produk, harga_total, sub_total, discount, total, bayar, kembalian, nama_kasir, tanggal, waktu)
		VALUES 
		("00001", "Disket", 4500, 3, 13500, 13500, 0, 0, 0, 0, "Rosi", "04-05-2022", "12:00:00"),
		("", "Refil Tinta", 22500, 1, 22500, 36000, 0,0,0,0, "", "", ""),
		("", "CD Blank", 1500, 4, 6000, 42000, 0,0,0,0, "", "", ""),
		("", "Mouse", 17500, 2, 35000, 77000, 0,77000,100000,23000, "", "", ""),
		("00002", "Disket", 4500, 1, 4500, 4500, 0,0,0,0, "Dewi", "04-05-2022", "12:00:00"),
		("", "Mouse", 17500, 1, 17500, 22000, 0,0,0,0, "", "", ""),
		("", "Flash Disk", 100000, 1, 100000, 117500, 0,117500,117500,0, "", "", "");`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestId:
func checkLatestId(id int) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}
	f := "0000" + strconv.Itoa(id)
	fmt.Print(f)
	sqlStmt := `SELECT no_bon FROM unormal WHERE no_bon = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, f)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// insert value hint
// ("00001", "Disket", 4500, 3, 13500, 13500, 0, 0, 0, 0, "Rosi", "04-05-2022", "12:00:00"),
// ("", "Refil Tinta", 22500, 1, 22500, 36000, 0,0,0,0, "", "", ""),
// ("", "CD Blank", 1500, 4, 6000, 42000, 0,0,0,0, "", "", ""),
// ("", "Mouse", 17500, 2, 35000, 77000, 0,77000,100000,23000, "", "", ""),
// ("00002", "Disket", 4500, 1, 4500, 4500, 0,0,0,0, "Dewi", "04-05-2022", "12:00:00"),
// ("", "Mouse", 17500, 1, 17500, 22000, 0,0,0,0, "", "", ""),
// ("", "Flash Disk", 100000, 1, 100000, 117500, 0,117500,117500,0, "", "", "")
