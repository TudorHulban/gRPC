package main

// client implementation

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/TudorHulban/gRPC/cmd/client/chat"
	"github.com/TudorHulban/gRPC/cmd/client/person"
	"google.golang.org/grpc"
)

func chatSend(wg *sync.WaitGroup, msg string, c chat.ChatServiceClient) error {
	defer wg.Done()

	resp, err := c.SendMessage(context.Background(), &chat.Message{Body: msg})
	if err != nil {
		log.Fatalf("Error when calling chat: %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)

	return nil
}

func main() {
	log.Println("starting client ...")

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	var wg sync.WaitGroup

	gClientChat := chat.NewChatServiceClient(conn)
	wg.Add(1)
	go chatSend(&wg, fmt.Sprintf("time is: %d", int(time.Now().Unix())), gClientChat)

	gClientPerson := person.NewPersonServiceClient(conn)
	wg.Add(1)
	go personSend(&wg, person.PersonData{
		Name: "John",
		Age:  32,
	}, gClientPerson)

	wg.Wait()
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
