package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create custom metrics
	requestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	responseTime := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_http_response_time_seconds",
			Help:    "Response time in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	// Register the metrics
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseTime)

	// Create an HTTP handler
	http.Handle("/metrics", promhttp.Handler())

	// Wrap your application logic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(responseTime.WithLabelValues(r.URL.Path))
		defer timer.ObserveDuration()

		// Increment request count
		requestsTotal.WithLabelValues(r.URL.Path).Inc()

		w.Write([]byte("Hello, Prometheus!"))
	})

	// Start the server
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}
}
