package helper

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func AsyncHttpGets(urls []string) []*http.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond) //timeout 100ms
	defer cancel()
	ch := make(chan *http.Response)
	responses := []*http.Response{}
	// TODO: answer here
<<<<<<< HEAD
	client := http.Client{}
	for _, url := range urls {
		go func(url string) {
			// TODO: answer here
			fmt.Printf("Fetching %s \n", url)
			resp, err := client.Get(url)
			ch <- &http.Response{}
			if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
			}
=======
	for _, url := range urls {
		go func(url string) {
			// TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		}(url)
	}

	for {
		select {
		case resp := <-ch:
			// fmt.Printf("%s was fetched\n", resp.Request.URL)
			responses = append(responses, resp)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		case <-ctx.Done():
			return responses
		}
	}
}

func SyncHttpGets(urls []string) []*http.Response {
	responses := []*http.Response{}
	client := http.Client{}

	for _, url := range urls {
		func(url string) {
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("error", err)
			}
			defer resp.Body.Close()
			// fmt.Printf("%s was fetched\n", resp.Request.URL)
			responses = append(responses, resp)
		}(url)
	}
	return responses
}

//reference: https://matt.aimonetti.net/posts/2012-11-real-life-concurrency-in-go/
