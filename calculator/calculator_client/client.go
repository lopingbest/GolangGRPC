package main

import (
	"context"
	"fmt"
	"github.com/lopinhbest/GolangGRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("Created client: %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Loping",
			LastName:  "Best",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
