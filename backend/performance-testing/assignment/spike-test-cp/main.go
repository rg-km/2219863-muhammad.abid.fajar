package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//kita bisa menggunakan attack yang meningkat seperti ini untuk stress/capacity test

//melakukan attack selama 4 detik
//setiap detik frequency akan naik sesuai increaseValue
//nilai yang digunakan
/*
duration := 4 * time.Second
frequency := 5
spikeValue := 20
*/

func spikeTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second                            //durasi attack
	frequency := 50                                        //jumlah request
	rate := vegeta.Rate{Freq: frequency, Per: time.Second} //mengatur rate request
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	}) //mengatur targeter vegeta
	metrics = vegetaAttackSpike(targeter, rate, duration,20) //menjalankan vegeta attack
	// fmt.Println(metrics.StatusCodes)                  //menampilkan status code
	// fmt.Println(metrics.Latencies.Max) 
	return metrics
}

//setiap detik ke 2 request akan melonjak
//contoh
//detik 1 request = 5
//detik 2 request = 20
//detik 3 request = 5
//detik 4 request = 20

func vegetaAttackSpike(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration, spikeValue int) *vegeta.Metrics {
	var metrics vegeta.Metrics
	// TODO: answer here
	attacker := vegeta.NewAttacker() //membuat attacker baru
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res) //menambahkan hasil attack ke dalam metrics
	} //melakukan vegeta attack
	metrics.Close()
	return &metrics
}
