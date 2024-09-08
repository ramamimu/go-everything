package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func callFirstApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 2000)
	fmt.Printf("from first api %v\n", fmt.Sprintf("%v", time.Now()))
}

func callSecondApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 2500)
	fmt.Printf("from second api %v\n", fmt.Sprintf("%v", time.Now()))
}

func callThirdApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 1000)
	fmt.Printf("from third api %v\n", fmt.Sprintf("%v", time.Now()))
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong")
	})

	http.HandleFunc("/multiple-context", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
		defer cancel()

		go callFirstApi(ctx)
		go callSecondApi(ctx)
		go callThirdApi(ctx)

		fmt.Println("I am from called func")
		io.WriteString(w, fmt.Sprintf("%v", time.Now()))
	})

	log.Println("running server on 8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Panicf("failed to run server due to %v", err)
	}
}
