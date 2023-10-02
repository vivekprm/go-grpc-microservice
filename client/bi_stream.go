package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/go-grpc-microservice/pbgo/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started.")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Couldn't send names")
	}

	waitc := make(chan struct{})

	go func ()  {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}
			log.Printf(message.Message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest {
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi-directional streaming finished.")
}