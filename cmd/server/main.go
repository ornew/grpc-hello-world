package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/ornew/grpc-hello-world/api"
)

type greeter struct {
	api.UnimplementedGreeterServer
}

var _ api.GreeterServer = (*greeter)(nil)

func (a *greeter) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("request: %v", req)
	if req.Name == "arata" {
		return nil, status.Error(codes.InvalidArgument, "you are not arata!")
	}
	if req.Name == "furukawa" {
		return &api.HelloReply{Message: "you are also furukawa?"}, nil
	}
	return &api.HelloReply{Message: "hello world."}, nil
}

const (
	port = ":8081"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &greeter{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
