package main

import (

	// TODO: answer here

	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

//Baca README untuk tau jumlah request yang perlu dilakukan dan targetnya
//untuk durasi cukup gunakan satu detik

//menambahkan movie baru
//untuk data yang dikirim adalah JSON
//gunakan struct Movie diatas, cukup gunakan field episode dan name
//ID sudah auto increment
func addMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second
	frequency := 10
	body := "hello"
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   []byte(body), //convert string ke byte agar bisa dikirim
	})
	metrics = vegetaAttack(targeter, frequency, duration)
	// fmt.Println(metrics.StatusCodes)
	// fmt.Println(metrics.Latencies.Max)
	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second                            //durasi attack
	frequency := 25                                        //jumlah request
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	}) //mengatur targeter vegeta
	metrics = vegetaAttack(targeter, frequency, duration) //menjalankan vegeta attack
	// fmt.Println(metrics.StatusCodes)                  //menampilkan status code
	// fmt.Println(metrics.Latencies.Max)  
	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second                            //durasi attack
	frequency := 20                                        //jumlah request
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	}) //mengatur targeter vegeta
	metrics = vegetaAttack(targeter, frequency, duration) //menjalankan vegeta attack
	// fmt.Println(metrics.StatusCodes)                  //menampilkan status code
	// fmt.Println(metrics.Latencies.Max)    
	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, frequency int, duration time.Duration) *vegeta.Metrics {
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
