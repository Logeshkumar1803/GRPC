package main

import (
	"context"
	"grpc_learn/cobra.com/learn/gogrpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received : %s", req.Name)
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":28080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Println("gRPC Server is running on port 28080...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
