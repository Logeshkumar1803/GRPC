package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc_learn/cobra.com/learn/gogrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:28080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewGreetServiceClient(conn)

	// Call SayHello RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Logesh"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	fmt.Println("Server Response:", res.Message)
}
