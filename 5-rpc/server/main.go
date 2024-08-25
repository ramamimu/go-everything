package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/ramamimu/go-everything/5-rpc/greeter"
	"google.golang.org/grpc"
)

type greet struct {
	greeter.UnimplementedGreeterServer
}

func (s *greet) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	log.Printf("Received in 1:1 : %v", in.GetName())
	return &greeter.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func (s *greet) SayHelloAgain(req *greeter.HelloRequest, stream greeter.Greeter_SayHelloAgainServer) error {
	log.Printf("Received in many:1 : %v", req.GetName())
	if err := stream.Send(&greeter.HelloReply{Message: "This is coming from server's stream heheheu"}); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)
	if err := stream.Send(&greeter.HelloReply{Message: "This is coming from server's stream heheheu"}); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := stream.Send(&greeter.HelloReply{Message: "This is coming from server's stream heheheu"}); err != nil {
		return err
	}
	return nil
}

func (s *greet) SayHelloReverse(stream greeter.Greeter_SayHelloReverseServer) error {
	counter := 0
	for {
		point, err := stream.Recv()
		counter++
		if err == io.EOF {
			return stream.SendAndClose(&greeter.HelloReply{Message: "Already get all data from 1:many"})
		}
		log.Printf("received from client in 1:many : %v\n has sent %d time", point.Name, counter)

		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	fmt.Println("listening to port :3000")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	var greeting greet
	greeter.RegisterGreeterServer(grpcServer, &greeting)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
