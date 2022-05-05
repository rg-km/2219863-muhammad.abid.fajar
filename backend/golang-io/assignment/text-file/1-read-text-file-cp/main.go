package main

import (
	"fmt"
	"io/ioutil"
<<<<<<< HEAD
	"log"
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
<<<<<<< HEAD
	fileName := "./read.txt"
	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	f := FileData{
		Name: fileName,
		Size: len(data),
		Data: data,
	}

	return f, nil // TODO: replace this
=======
	return FileData{}, nil // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
