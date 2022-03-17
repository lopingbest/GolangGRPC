package main

import (
	"context"
	"fmt"
	"github.com/lopinhbest/GolangGRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManytimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

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
