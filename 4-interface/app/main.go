package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	opsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "response_variance",
			Namespace: "http_request_count",
		},
		[]string{"method", "path", "code"},
	)
)

func main() {
	prometheus.MustRegister(opsProcessed)

	// pull method
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
			opsProcessed.WithLabelValues("GET", "/monitor", "405").Inc()
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		case 1:
			opsProcessed.WithLabelValues("GET", "/monitor", "500").Inc()
			http.Error(w, "internal server error", http.StatusInternalServerError)
		case 2:
			opsProcessed.WithLabelValues("GET", "/monitor", "400").Inc()
			http.Error(w, "bad request", http.StatusBadRequest)
		case 3:
			opsProcessed.WithLabelValues("GET", "/monitor", "201").Inc()
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "success create account")
		case 4:
			opsProcessed.WithLabelValues("GET", "/monitor", "200").Inc()
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
		default:
			opsProcessed.WithLabelValues("GET", "/monitor", "200").Inc()
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
		}
	})

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic("failed to create server")
	}
}
