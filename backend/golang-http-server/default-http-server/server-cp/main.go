package main

import (
	"log"
	"net/http"
)

// Dari contoh yang diberikan, buatlah implementasi server dengan menggunakan struct Server dari package http/net/
// Buatlah server dengan address localhost dan di port 3000

func main() {
	// TODO: answer here
	server := http.Server{
		Addr: "localhost:3000",
	}

	log.Println("Server running on port 3000")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
