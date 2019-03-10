package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/pb"
)

const (
	address = "0.0.0.0:9090"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	r, err := c.GetUser(context.Background(), &pb.GetUserRequest{Id: "hhhh"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
