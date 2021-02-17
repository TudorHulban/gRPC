package main

// server implementation

import (
	"log"
	"net"

	"github.com/TudorHulban/gRPC/cmd/server/chat"
	"google.golang.org/grpc"
)

func main() {
	log.Println("starting server ...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
