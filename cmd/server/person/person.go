package person

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SendData(ctx context.Context, in *PersonData) (*Results, error) {
	log.Printf("Receive person data from client: %s", in.Name)

	return &Results{
		Valid: false,
	}, nil
}
