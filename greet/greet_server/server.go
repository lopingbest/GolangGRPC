package main

import (
	"fmt"
	"github.com/lopinhbest/GolangGRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func main() {
	fmt.Println("Hello World")

	// listen on port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	// check error
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create a new gRPC server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	//display eror message
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
