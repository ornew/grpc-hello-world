package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/ornew/grpc-hello-world/api"
	"google.golang.org/grpc"
)

var (
	address = flag.String("address", "localhost:8081", "server address")
	name    = flag.String("name", "john", "your name")
)

func init() {
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := api.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("greeting: %s", r.GetMessage())
}
