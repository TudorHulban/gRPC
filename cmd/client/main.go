package main

// client implementation

import (
	"context"
	"log"
	"sync"

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

	var wg sync.WaitGroup

	c := chat.NewChatServiceClient(conn)
	wg.Add(1) // before the actual go routine call to prevent race conditions
	go chatSend(&wg, "xxxxxx", c)

	p := person.NewPersonServiceClient(conn)
	wg.Add(1)
	go personSend(&wg, person.PersonData{
		Name: "John",
		Age:  32,
	}, p)

	wg.Wait()
}

func chatSend(wg *sync.WaitGroup, msg string, c chat.ChatServiceClient) error {
	defer wg.Done()

	resp, err := c.SayHello(context.Background(), &chat.Message{Body: msg})
	if err != nil {
		log.Fatalf("Error when calling chat: %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)

	return nil
}

func personSend(wg *sync.WaitGroup, p person.PersonData, serv person.PersonServiceClient) error {
	defer wg.Done()

	resp, err := serv.SendData(context.Background(), &p)
	if err != nil {
		log.Fatalf("Error when calling person: %t", err)
	}

	log.Printf("Response from server: %t", resp.Valid)

	return nil
}
