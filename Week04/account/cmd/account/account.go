package main

import (
	"context"
	pb "github.com/kaige0207/Go-000/Week04/myapp/api/v1/user"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/pkg/configreader"
	"golang.org/x/sync/errgroup"
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
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, &userService)
		log.Println("Server Start...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			cancel()
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		log.Fatalf("error in serve: %v", err)
	}

}
