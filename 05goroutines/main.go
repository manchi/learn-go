package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// create a wait group
	var wg sync.WaitGroup

	websites := []string{
		"https://go.dev",
		"https://fb.com",
		"https://abc.com",
		"https://xyz.com",
		"https://msn.com",
		"https://google.com",
		"https://twitter.com",
	}

	start := time.Now()
	// until the job in getStatusCode is done, wait for it
	for _, web := range websites {
		go getStatusCode(web, &wg)
		wg.Add(1)
	}

	// wait for the signal from the goroutine
	wg.Wait()
	fmt.Println("total time:", time.Since(start))
}

func getStatusCode(url string, wg *sync.WaitGroup) {
	// mark the done for wait group
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Oops in url:", err)
	} else {
		fmt.Printf("got %v status code for %v\n", res.StatusCode, url)
	}

}
