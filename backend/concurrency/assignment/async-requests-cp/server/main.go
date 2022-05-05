package main

import (
	"fmt"
	"log"
	"net/http"

<<<<<<< HEAD
	"github.com/ruang-guru/playground/backend/concurrency/exercise/async-requests-cp/server/handler"
=======
	"github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/server/handler"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

func main() {
	http.HandleFunc("/", handler.GetMessage)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
