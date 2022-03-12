package main

import (
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
	fmt.Printf("Created client: %f", c)

}


