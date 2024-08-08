package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong")
	})

	http.HandleFunc("/person", func(w http.ResponseWriter, req *http.Request) {
		user := &Person{
			Name: "anonymous",
			Age:  rand.IntN(50),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Cannot encode response to JSON", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/monitor", func(w http.ResponseWriter, req *http.Request) {
		randomizer := rand.IntN(5)

		fmt.Println("got ", randomizer)
		switch randomizer {
		case 0:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		case 1:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		case 2:
			http.Error(w, "bad request", http.StatusBadRequest)
		case 3:
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "success create account")
		case 4:
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
		default:
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
		}
	})

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic("failed to create server")
	}
}
