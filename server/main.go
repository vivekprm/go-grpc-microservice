package main

import (
	"log"
	"net"

	pb "github.com/go-grpc-microservice/pbgo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

type helloServer struct {
	pb.GreetServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServer(grpcServer, &helloServer{})

	log.Printf("Server started at: %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}