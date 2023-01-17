package main

import (
	"context"
	"log"

	"github.com/TudorHulban/gRPC/pb"
)

type Server struct{}

func (s Server) SendMessage(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	return &pb.Message{Body: "Hi From the Server!"}, nil
}

func (s *Server) SendData(ctx context.Context, in *pb.PersonData) (*pb.Results, error) {
	log.Printf("Receive person data from client: %s", in.Name)

	return &pb.Results{
		Valid: false,
	}, nil
}
