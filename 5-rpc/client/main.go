package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/ramamimu/go-everything/5-rpc/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func simpleRPC(client greeter.GreeterClient) {
	// 1:1
	resp, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "Rams"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())
}

func serverSideStream(client greeter.GreeterClient, wg *sync.WaitGroup) {
	// rcv stream from server
	stream, err := client.SayHelloAgain(context.Background(), &greeter.HelloRequest{
		Name: "mamimu",
	})
	if err != nil {
		log.Println("failed trigger server's stream")
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			} else {
				valStr := fmt.Sprintf("Response msg: %s", resp.Message)
				log.Println(valStr)
			}

			if err != nil {
				log.Println("failed to read server's stream")
			}
		}

	}()
}

func clientSideStream(client greeter.GreeterClient, wg *sync.WaitGroup) {
	// 1:many
	streamClient, err := client.SayHelloReverse(context.Background())
	if err != nil {
		log.Fatalf("could not initiate 1:many to server: %v", err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			if err := streamClient.Send(&greeter.HelloRequest{Name: fmt.Sprintf("double double you %d", i)}); err != nil {
				log.Fatalf("could not send clientSideStream to server: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		reply, err := streamClient.CloseAndRecv()
		if err != nil {
			log.Fatalf("client.SayHelloReverse failed: %v", err)
		}
		log.Printf("Stream summary: %v", reply)
		wg.Done()
	}()
}

func bothSideStream(client greeter.GreeterClient, wg *sync.WaitGroup) {
	// many:many
	stream, err := client.SayHelloBidirectional(context.Background())
	if err != nil {
		log.Fatalf("could not initiate many:many to server: %v", err)
	}

	go func() {
		for i := 0; i < 40; i++ {
			if err := stream.Send(&greeter.HelloRequest{Name: fmt.Sprintf("double double you %d", i)}); err != nil {
				log.Fatalf("could not send stream many:many to server: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		if err := stream.CloseSend(); err != nil {
			log.Printf("could not close send stream: %v", err)
		}
		wg.Done()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server closed the stream")
				wg.Done()
				return
			}

			if err != nil {
				log.Println("failed to read bidirectional server's stream")
			} else {
				valStr := fmt.Sprintf("Response many:many msg: %s", resp.Message)
				log.Println(valStr)
			}
		}
	}()
}

func main() {
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	var wg sync.WaitGroup

	client := greeter.NewGreeterClient(conn)
	simpleRPC(client)
	// wg.Add(1)
	// serverSideStream(client, &wg)
	// wg.Add(1)
	// clientSideStream(client, &wg)
	wg.Add(2)
	bothSideStream(client, &wg)

	wg.Wait()
}
