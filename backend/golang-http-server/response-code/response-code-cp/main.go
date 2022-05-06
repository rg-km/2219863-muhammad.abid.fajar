package main

import (
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "POST" {
			if r.FormValue("name") == "Roger" {
				w.WriteHeader(405)
				return
			}
		} else if r.Method == "GET" {
			if r.FormValue("name") == "Roger" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(404)
			}
		}

	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/", GetNameHandler())
	return mux
}
