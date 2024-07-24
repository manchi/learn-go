package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var websites = []string{
	"https://go.dev",
	"https://fb.com",
	"https://abc.com",
	"https://xyz.com",
	"https://msn.com",
	"https://google.com",
	"https://twitter.com",
}

func main() {

	fmt.Println("go waitgroup tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8999", nil))
}

func fetchStatus(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	var wg sync.WaitGroup

	for _, url := range websites {
		wg.Add(1)

		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Oops in url:", err)
				fmt.Fprintf(writer, "Oops in url: %v\n", err)
			} else {
				fmt.Printf("got %v status code for %v\n", resp.StatusCode, url)
				fmt.Fprintf(writer, "got %v status code for %v\n", resp.StatusCode, url)
			}
			wg.Done()
		}(url)
	}

	wg.Wait()
	fmt.Println("total time:", time.Since(start))
}
