package main

import (
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {
	file, err := os.Create("write.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	//menutup file sebelum fungsi selesai dijalankan
	defer file.Close()

	file.WriteString("Write anything here\n" +
		"This program demonstrates reading and writing\n" +
		"operations to a file in Go")

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}
