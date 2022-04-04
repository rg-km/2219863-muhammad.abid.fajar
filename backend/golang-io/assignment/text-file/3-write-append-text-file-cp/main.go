package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	// Silahkan kalian cari tau arti dari os.O_APPEND dan os.O_WRONLY dan 0644 ini apa
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(stringToAdd); err != nil {
		return err
	}

	// read file again
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	log.Printf("%s", data)

	return nil // TODO: replace this
}
