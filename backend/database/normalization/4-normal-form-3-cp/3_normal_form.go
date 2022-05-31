package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `CREATE TABLE IF NOT EXISTS rekap_3nf (
		no_bon VARCHAR(10),
		discount INTEGER,
		total INTEGER,
		bayar INTEGER,
		kembalian INTEGER,
		kode_kasir VARCHAR(30),
		tanggal VARCHAR(30),
		waktu VARCHAR(30)
	) ;` // TODO: replace this
=======
	sqlStmt := `CREATE TABLE rekap ... ` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`
	INSERT OR REPLACE INTO 
	rekap_3nf (no_bon, discount, total, bayar, kembalian, kode_kasir, tanggal, waktu)
	VALUES 
	("00001", 0, 13500, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
	("00001", 0, 36000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
	("00001", 0, 42000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
	("00001", 0, 77000, 100000, 23000, "K01", "04-05-2022", "12:00:00"),
	("00002", 0, 4500, 17500, 0, "K02", "04-05-2022", "12:00:00"),
	("00002", 0, 22000, 117500, 0, "K02", "04-05-2022", "12:00:00"),
	("00002", 0, 117500, 117500, 0, "K02", "04-05-2022", "12:00:00");`) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO rekap ... `) // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE IF NOT EXISTS rekap_detail_3nf (
		no_bon VARCHAR(10),
		kode_barang VARCHAR(30),
		harga INTEGER,
		jumlah INTEGER
	) ` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE rekap_detail ... ` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`
	INSERT OR REPLACE INTO
	rekap_detail_3nf (no_bon, kode_barang, harga, jumlah)
	VALUES
	("00001", "B001", 4500, 3),
	("00001", "B002", 22500, 1),
	("00001", "B003", 1500, 4),
	("00001", "B004", 17500, 2),
	("00002", "B001", 4500, 1),
	("00002", "B004", 17400, 1),
	("00002", "B005", 100000, 1);`) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO rekap_detail ... `) // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE IF NOT EXISTS barang_3nf (
		kode_barang VARCHAR(10),
		nama_barang VARCHAR(30),
		harga_barang INTEGER
	) ;` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE barang ... ` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`
	INSERT OR REPLACE INTO
	barang_3nf (kode_barang, nama_barang, harga_barang)
	VALUES 
	("B001", "Disket", 4500),
	("B002", "Refil Tinta", 22500),
	("B003", "CD Blank", 1500),
	("B004", "Mouse", 17500),
	("B005", "Flash Disk", 100000);`) // TODO: replace this
=======
	_, err = db.Exec(`INSEERT INTO barang ... `) // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE IF NOT EXISTS kasir_3nf (
		kode_kasir VARCHAR(10),
		nama_kasir VARCHAR(30)
	) ;` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE kasir ... ` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`
	INSERT OR REPLACE INTO
	kasir_3nf (kode_kasir, nama_kasir)
	VALUES 
	("K01", "Rosi"),
	("K02", "Dewi");`) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO kasir ... `) // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT no_bon FROM rekap_3nf WHERE no_bon = ?;` // TODO: replace this
=======
	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT no_bon FROM rekap_detail_3nf WHERE no_bon = ?;` // TODO: replace this
=======
	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT harga_barang FROM barang_3nf WHERE kode_barang = ?;` // TODO: replace this
=======
	sqlStmt := `...` // TODO: replace this
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT kode_kasir FROM kasir_3nf WHERE kode_kasir = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, kode_kasir)
	var latestId string
=======
	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
>>>>>>> 59c364d69411d5bf5e3abd9985de8b5d350c840b
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}
