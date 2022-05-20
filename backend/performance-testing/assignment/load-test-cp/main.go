package main

import (

	// TODO: answer here

	"encoding/json"
	"strconv"
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
	payload, _ := json.Marshal(Movie{
		Episode: 1,
		Name:    "baru",
	})
	duration := 1 * time.Second
	frequency := 10
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target + "/movie",
		Body:   payload,
	})
	metrics = vegetaAttack(targeter, frequency, duration)
	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second
	frequency := 25
	targets := []vegeta.Target{}
	for i := 0; i < 25; i++ {
		numberString := strconv.Itoa(i + 1)
		targeter := vegeta.Target{
			Method: "GET",
			URL:    target + "/movie/" + numberString,
		}
		targets = append(targets, targeter)
	}
	targeters := vegeta.NewStaticTargeter(targets...)
	metrics = vegetaAttack(targeters, frequency, duration)
	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 1 * time.Second
	frequency := 20
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target + "/movies/",
	})
	metrics = vegetaAttack(targeter, frequency, duration)
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
