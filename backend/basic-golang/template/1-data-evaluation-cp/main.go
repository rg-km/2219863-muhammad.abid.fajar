package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, cobalah untuk mempraktikkan data evaluation pada template.
// Lengkapi function ExecuteToByteBuffer sehingga objek dari struct Account dapat ter-render ke dalam template.
// Gunakan bytes.Buffer sebagai io.Writer pada template.
// Lengkapi juga textTemplate, sehingga variabel dari objek Account dapat ter-render ke dalam template.
// Contoh:
// acc := {Name: "Tony", Number: "1002321", Balance: 1000}
// Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000.

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
<<<<<<< HEAD
	c := account
	textTemplate = "Akun {{.Name}} dengan Nomor {{.Number}} memiliki saldo sebesar ${{.Balance}}."
	tmpl, err := template.New("test").Parse(textTemplate) // ".FirstName" dan ".Age" adalah field name yang ada di struct User
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, c)
	if err != nil {
		panic(err)
	}

	return b.Bytes(), nil
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
