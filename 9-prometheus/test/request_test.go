package test

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
)

func TestRequest(t *testing.T) {
	numberReq := 1000
	var wg sync.WaitGroup

	for i := 0; i < numberReq; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create a new GET request
			resp, err := http.Get("http://localhost:8088/")
			if err != nil {
				fmt.Println("Error:", err)
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			fmt.Println("Response:", string(body))
		}()
	}

	wg.Wait()
	fmt.Println("request done")
}
