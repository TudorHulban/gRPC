package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	return &Message{Body: "Hi From the Server!"}, nil
}

func (s *Server) SendData(ctx context.Context, in *PersonData) (*Results, error) {
	log.Printf("Receive person data from client: %s", in.Name)

	return &Results{
		Valid: false,
	}, nil
}
