package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ramamimu/go-everything/5-rpc/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	client := greeter.NewGreeterClient(conn)

	// 1:1
	resp, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "Rams"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())

	// 1:many
	streamClient, err := client.SayHelloReverse(context.Background())
	if err != nil {
		log.Fatalf("could not initiate 1:many to server: %v", err)
	}

	for i := 0; i < 3; i++ {
		if err := streamClient.Send(&greeter.HelloRequest{Name: fmt.Sprintf("double double you %d", i)}); err != nil {
			log.Fatalf("could not send stream to server: %v", err)
			time.Sleep(500 * time.Millisecond)
		}
	}
	reply, err := streamClient.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.SayHelloReverse failed: %v", err)
	}
	log.Printf("Stream summary: %v", reply)

	// rcv stream from server
	stream, err := client.SayHelloAgain(context.Background(), &greeter.HelloRequest{
		Name: "mamimu",
	})
	if err != nil {
		log.Println("failed trigger server's stream")
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		} else {
			valStr := fmt.Sprintf("Response msg: %s", resp.Message)
			log.Println(valStr)
		}

		if err != nil {
			log.Println("faied to read server's stream")
		}
	}
}
