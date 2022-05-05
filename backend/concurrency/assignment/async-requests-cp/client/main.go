package main

import (
	"fmt"
	"time"

<<<<<<< HEAD
	helper "github.com/ruang-guru/playground/backend/concurrency/exercise/async-requests-cp/client/request"
=======
	helper "github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/client/request"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	baseURL := "http://localhost:8080/"
	urls := []string{}
	for i := 0; i < 10; i++ {
		urls = append(urls, baseURL)
	}
	results := helper.AsyncHttpGets(urls)
	for _, result := range results {
		if result != nil && result.Body != nil {
			fmt.Printf("%s. Waktu %v\n", result.Status, time.Since(start))
		}
	}
}

//reference: https://matt.aimonetti.net/posts/2012-11-real-life-concurrency-in-go/
