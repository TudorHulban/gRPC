package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SendMessage(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	return &Message{Body: "Hi From the Server!"}, nil
}
