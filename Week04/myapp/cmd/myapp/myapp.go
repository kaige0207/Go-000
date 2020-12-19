package main

import (
	pb "github.com/kaige0207/Go-000/Week04/myapp/api/v1/user"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/pkg/configreader"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = configreader.GetConfig().Port

func main() {
	userService := initService()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
