package main

import (
	"context"
	"io"
	"log"

	pb "github.com/go-grpc-microservice/pbgo/proto"
)

func callSayHelloServerStream(client pb.GreetClient, names *pb.NamesList) {
	log.Printf("streaming started..")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names: %v", names)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Println(message)
	}

	log.Printf("streaming finished.")
}