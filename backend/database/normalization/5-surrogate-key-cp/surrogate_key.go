package main

// pada contoh gambar surrogate_key, diketahui terdapat 2 table pendaftaran sekolah, dimana kedua table tersebut
// memilik kolom regirstration_no, name, dan percentage.
// gabunglah kedua table tersebut, dan buatlah surrogate key pada table yang baru.

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type SchoolA struct {
	RegistrationNo int
	Name           string
	Percentage     float64
	Grade          string
	NationalRank   int
}

type SchoolB struct {
	RegistrationNo int
	Name           string
	Percentage     float64
}

type SurrogateTable struct {
	Id             int
	RegistrationNo int
	Name           string
	Percentage     float64
	Grade          string
	NationalRank   int
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./surrogate.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS school_a_cp (
		registration_no INTEGER PRIMARY KEY,
		name TEXT,
		percentage REAL,
		grade TEXT,
		national_rank INT
	) ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	_, err = db.Exec(`
	INSERT OR REPLACE INTO 
	school_a_cp (registration_no, name, percentage, grade, national_rank)
	VALUES
	(1000, "SMA Negeri 1", 0.5, "A", 1),
	(2000, "SMA Negeri 2", 0.5, "A", 2),
	(3000, "SMA Negeri 3", 0.5, "A", 3),
	(4000, "SMA Negeri 4", 0.5, "A", 4),
	(5000, "SMA Negeri 5", 0.5, "A", 5);`) // TODO: replace this

	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS school_b_cp (
		registration_no INTEGER PRIMARY KEY,
		name TEXT,
		percentage REAL
	) ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT OR REPLACE INTO 
	school_b_cp (registration_no, name, percentage)
	VALUES
	(1000, "SMA Negeri 1", 0.5),
	(2000, "SMA Negeri 2", 0.5),
	(3000, "SMA Negeri 3", 0.5),
	(4000, "SMA Negeri 4", 0.5),
	(5000, "SMA Negeri 5", 0.5);`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS surrogate_table_cp (
		id INTEGER PRIMARY KEY,
		registration_no INTEGER,
		name TEXT,
		percentage REAL,
		grade TEXT,
		national_rank INT
	 ) ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT OR REPLACE INTO 
	surrogate_table_cp (id, registration_no, name, percentage, grade, national_rank)
	VALUES
	(1000, 1000, "SMA Negeri 1", 0.5, "A", 1),
	(2000, 2000, "SMA Negeri 2", 0.5, "A", 2),
	(3000, 3000, "SMA Negeri 3", 0.5, "A", 3),
	(4000, 4000, "SMA Negeri 4", 0.5, "A", 4),
	(5000, 5000, "SMA Negeri 5", 0.5, "A", 5);`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}