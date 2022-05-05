package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat function untuk menghitung average score siswa
//panggil function didalam template

type Student struct {
	Name   string
	Scores []float64
}

func (s Student) CalculateScore(scores []float64) float64 {
	// TODO: answer here
<<<<<<< HEAD
	var b float64
	// b = s.Scores[0] + s.Scores[1] + s.Scores[2]
	for _, i := range s.Scores {
		b += i
	}
	return b / 3
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (s Student) GenerateStudentTemplate() string {
	buff := new(bytes.Buffer)
	// TODO: answer here
<<<<<<< HEAD
	tmp1 := template.New("Template_1")

	// "Parse" parses string dan set action untuk mengambil nilai dari Field Name dan Age
	tmp1, err := tmp1.Parse("Hello {{.Name}}, Nilai rata-rata kamu 11")
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}

	if err := tmp1.Execute(buff, s); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	return buff.String()
}

func NewStudent(name string, scores []float64) Student {
	return Student{name, scores}
}

// main function
func main() {
	std := NewStudent("Rogu", []float64{10, 11, 12})
	fmt.Println(std.GenerateStudentTemplate())
}
