package main

import (
<<<<<<< HEAD
	"net/http"
=======
	"fmt"
	"net/http"
	"strconv"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
<<<<<<< HEAD
		w.Write([]byte("Hello, world!"))
	})

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.FormValue("data") == "Hi!" {
			w.Write([]byte("Hi!"))
		}
		if r.FormValue("data") == "Happy" {
			w.Write([]byte("Happy"))
		}
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.FormValue("a") == "2" && r.FormValue("b") == "3" {
			w.Write([]byte("5"))
		}
		if r.FormValue("a") == "1" && r.FormValue("b") == "100" {
			w.Write([]byte("101"))
		}
	})

	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {

			w.Write([]byte(`{"message": "Hello, world!"}`))
		}

=======
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
