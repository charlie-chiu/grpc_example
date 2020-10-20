package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc"
)

const addr = "127.0.0.1:9999"

type myServer struct{}

func (myServer) SayHello(_ context.Context, request *greeting.HelloRequest) (*greeting.HelloReply, error) {
	log.Printf("got request %v", request.GetName())
	return &greeting.HelloReply{
		Greeter: "go server",
		Message: "hello, " + request.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net.listen error %v", err)
	}
	log.Printf("server listening on %v", addr)
	gRPCServer := grpc.NewServer()
	greeting.RegisterGreeterServer(gRPCServer, &myServer{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve, %v", err)
	}
}
