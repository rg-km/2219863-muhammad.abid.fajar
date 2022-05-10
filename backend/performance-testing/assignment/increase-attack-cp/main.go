package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//kita bisa menggunakan attack yang meningkat seperti ini untuk stress/capacity test

//melakukan attack selama 4 detik
//setiap detik frequency akan naik sesuai increaseValue
//nilai yang digunakan
/*
duration := 4 * time.Second
frequency := 1
increaseValue := 2
*/
func increaseTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 4 * time.Second //durasi attack
	frequency := 4             //jumlah request
	increaseValue := 2
	rate := vegeta.Rate{Freq: frequency, Per: time.Second} //mengatur rate request
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	}) //mengatur targeter vegeta
	metrics = vegetaAttackIncreaseBySecond(targeter, rate, duration, increaseValue) //menjalankan vegeta attack
	fmt.Println(metrics.StatusCodes)                                                //menampilkan status code
	fmt.Println(metrics.Latencies.Max)

	return metrics
}

//frequency request akan naik sebesar increaseValue
//diawal tidak perlu menambahkan increaseValue
//contoh
//pertama 10 request
//kedua 15 request
//ketiga 20 request
//total 45 request

func vegetaAttackIncreaseBySecond(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration, increaseValue int) *vegeta.Metrics {
	var metrics vegeta.Metrics
	// TODO: answer here
	attacker := vegeta.NewAttacker()
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res) //menambahkan hasil attack ke dalam metrics
	} //melakukan vegeta attack
	metrics.Close()
	return &metrics
}
