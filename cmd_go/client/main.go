package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc"
)

func main() {
	log.Println("gRPC client in golang")
	const addr = "127.0.0.1:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("grpc.Dial error:", err)
	}
	defer conn.Close()

	client := greeting.NewGreeterClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	name := "Charlie the Gopher"
	log.Printf("say hello to server with name %q\n", name)
	helloReply, err := client.SayHello(ctx, &greeting.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("sayHello failed,", err)
	}
	log.Printf("Response from greeter %q, msg: %q\n", helloReply.Greeter, helloReply.Message)
}
