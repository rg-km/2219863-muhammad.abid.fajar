package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//menjalankan vegeta attack dengan method GET
//durasi yang digunakan 2 detik dan rate/frequency 10
//target didapatkan dari parameter
func vegetaGet(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 2 * time.Second
	frequency := 10
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics = vegetaAttack(targeter, rate, duration)
	return metrics
}

//menjalankan vegeta attack dengan method POST
//durasi yang digunakan 2 detik dan rate/frequency 15
//target didapatkan dari parameter fungsi ini
func vegetaPost(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 2 * time.Second
	frequency := 15
	body := "hello"
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   []byte(body),
	})
	metrics = vegetaAttack(targeter, rate, duration)
	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration) *vegeta.Metrics {
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
