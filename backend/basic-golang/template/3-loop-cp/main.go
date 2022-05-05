package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
<<<<<<< HEAD
	Name string
	Rank int
=======
	Name  string
	Email string
	Rank  int
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

type Leaderboard struct {
	Users []*UserRank
}

<<<<<<< HEAD
// Peringkat ke-1: RogerPeringkat ke-2: TonyPeringkat ke-3: BrucePeringkat ke-4: NatashaPeringkat ke-5: Clint
func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	textTemplate := `{{range . }}Peringkat ke-{{ .Rank }}: {{ .Name }}{{else}} Invalid "struct" Users harus berupa array!{{end}}`
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		panic(err)
	}
	var a bytes.Buffer
	err = tmpl.Execute(&a, leaderboard.Users)
	if err != nil {
		panic(err)
	}

	return a.Bytes(), nil
=======
func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
