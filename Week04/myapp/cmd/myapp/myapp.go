package main

import (
	pb "github.com/kaige0207/Go-000/Week04/myapp/api/v1/user"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/pkg/configreader"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	server = configreader.GetConfig().Server
	port   = configreader.GetConfig().Port
)

func main() {
	userService := initService()
	addr := server + ":" + port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService)
	log.Println("Server Start...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
