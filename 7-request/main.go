package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numRequests := 1000
	url := "http://localhost:9999/ping"

	success := 0
	failed := 0

	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		// Start each request in a new goroutine
		go func(requestID int) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Request #%d failed: %v", requestID, err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Request #%d failed to read response: %v", requestID, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				failed++
			} else {
				success++
			}

			fmt.Printf("Response from request #%d: %s\n", requestID, body)
		}(i + 1) // Pass request ID for logging
	}

	// Wait for all requests to complete
	wg.Wait()
	fmt.Println("All requests completed.")
	fmt.Println("success: ", success)
	fmt.Println("Failed: ", failed)
}
