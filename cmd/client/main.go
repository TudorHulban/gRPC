package main

// client implementation

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/TudorHulban/gRPC/pb"
	"google.golang.org/grpc"
)

func chatSend(wg *sync.WaitGroup, msg string, c pb.ChatServiceClient) error {
	defer wg.Done()

	resp, err := c.SendMessage(context.Background(), &pb.Message{Body: msg})
	if err != nil {
		log.Fatalf("Error when calling chat: %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)

	return nil
}

func personSend(wg *sync.WaitGroup, p *pb.PersonData, serv pb.PersonServiceClient) error {
	defer wg.Done()

	resp, err := serv.SendData(context.Background(), p)
	if err != nil {
		log.Fatalf("Error when calling person: %t", err)
	}

	log.Printf("Response from server: %t", resp.Valid)

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

	gClientChat := pb.NewChatServiceClient(conn)
	wg.Add(1)
	go chatSend(&wg, fmt.Sprintf("time is: %d", int(time.Now().Unix())), gClientChat)

	gClientPerson := pb.NewPersonServiceClient(conn)
	wg.Add(1)
	go personSend(&wg, &pb.PersonData{
		Name: "John",
		Age:  32,
	}, gClientPerson)

	wg.Wait()
}
