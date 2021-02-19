package main

// client implementation

import (
	"context"
	"log"

	"github.com/TudorHulban/gRPC/cmd/client/chat"
	"github.com/TudorHulban/gRPC/cmd/client/person"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	resp1, err := c.SayHello(context.Background(), &chat.Message{Body: "Hi From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", resp1.Body)

	p := person.NewPersonServiceClient(conn)

	resp2, err := p.SendData(context.Background(), &person.PersonData{
		Name: "John",
		Age:  32,
	})
	if err != nil {
		log.Fatalf("Error when calling SendData: %s", err)
	}

	log.Printf("Response from server: %t", resp2.Valid)
}
