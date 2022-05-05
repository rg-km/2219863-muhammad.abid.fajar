package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	// TODO: answer here
<<<<<<< HEAD
	Name         string
	ScoreAverage float64
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
<<<<<<< HEAD
	std := Student{Name: "Rogu", ScoreAverage: 7.8}

	// "New" membuat template baru
	// dengan nama Template_1
	tmp1 := template.New("Template_1")

	// "Parse" parses string dan set action untuk mengambil nilai dari Field Name dan Age
	tmp1, err := tmp1.Parse("Hello {{.Name}}, Nilai rata rata kamu {{.ScoreAverage}}!")
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
