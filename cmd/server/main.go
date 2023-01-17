package main

// server implementation

import (
	"log"
	"net"

	"github.com/TudorHulban/gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("starting server ...")

	listener, errListen := net.Listen("tcp", ":9000")
	if errListen != nil {
		log.Fatalf("failed to listen: %v", errListen)
	}

	var server Server

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &server)
	pb.RegisterPersonServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
