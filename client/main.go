package main

import (
	"log"

	pb "github.com/go-grpc-microservice/pbgo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	conn, err := grpc.Dial("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetClient(conn)

	names := &pb.NamesList{
		Names: []string{"Vivek", "Anav"},
	}

	// callSayHello(client)
	callSayHelloServerStream(client, names)
}