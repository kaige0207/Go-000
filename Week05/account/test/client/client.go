package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kaige0207/Go-000/Week05/account/api/v1/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.UserRequest{Username: "jack", Password: "jack"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Login message: %s", r.GetMessage())

	r, err = c.Register(ctx, &pb.UserRequest{Username: "jack", Password: "jack"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Register message: %s", r.GetMessage())
}
