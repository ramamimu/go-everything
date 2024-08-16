package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	sloglogstash "github.com/samber/slog-logstash/v2"
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
	logstashAddr := "localhost:5000"
	conn, err := net.Dial("tcp", logstashAddr)
	if err != nil {
		log.Fatalf("could not connect to Logstash: %v", err)
	} else {
		println("connected to logstash successfully")
	}
	defer conn.Close()
	_ = sloglogstash.Option{Level: slog.LevelDebug, Conn: conn}.NewLogstashHandler()

	logHandler := slog.NewJSONHandler(conn, &slog.HandlerOptions{Level: slog.LevelDebug})

	log.Println("server start")
	prometheus.MustRegister(opsProcessed)
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	logger.Info("starting http server...")

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
			logger.With("method", "GET", "product", "a").Error("method not allowed")
		case 1:
			opsProcessed.WithLabelValues("GET", "/monitor", "500").Inc()
			http.Error(w, "internal server error", http.StatusInternalServerError)
			logger.With("method", "GET", "product", "b").Error("internal server error")
		case 2:
			opsProcessed.WithLabelValues("GET", "/monitor", "400").Inc()
			http.Error(w, "bad request", http.StatusBadRequest)
			logger.With("method", "GET", "product", "c").Error("bad request")
		case 3:
			opsProcessed.WithLabelValues("GET", "/monitor", "201").Inc()
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "success create account")
			logger.With("method", "GET", "product", "d").Info("success create account")
		case 4:
			opsProcessed.WithLabelValues("GET", "/monitor", "200").Inc()
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
			logger.With("method", "GET", "product", "e").Info("success get data")
		default:
			opsProcessed.WithLabelValues("GET", "/monitor", "200").Inc()
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success fetch the data")
			logger.With("method", "GET", "product", "f").Info("success get data")
		}
	})

	err = http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Panic("failed to run server")
	}
	log.Println("server end")
}

type LogstashHandler struct {
	conn net.Conn
}

func NewLogstashHandler(address string) (*LogstashHandler, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &LogstashHandler{conn: conn}, nil
}

func (h *LogstashHandler) Handle(ctx context.Context, record slog.Record) error {
	timestamp := time.Now().Format(time.RFC3339)
	message := record.Message
	logLine := timestamp + " " + record.Level.String() + " " + message + "\n"
	_, err := h.conn.Write([]byte(logLine))
	return err
}

func (h *LogstashHandler) Enabled(ctx context.Context, level slog.Level) {
	// return
}

func (h *LogstashHandler) Close() error {
	return h.conn.Close()
}
