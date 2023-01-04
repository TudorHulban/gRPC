package main

// server implementation

import (
	"log"
	"net"

	"github.com/TudorHulban/gRPC/cmd/server/chat"
	"github.com/TudorHulban/gRPC/cmd/server/person"

	"google.golang.org/grpc"
)

func main() {
	log.Println("starting server ...")

	lis, errListen := net.Listen("tcp", ":9000")
	if errListen != nil {
		log.Fatalf("failed to listen: %v", errListen)
	}

	var serverChat chat.Server
	var serverPerson person.Server

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &serverChat)
	person.RegisterPersonServiceServer(grpcServer, &serverPerson)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
