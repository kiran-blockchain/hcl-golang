package chat

import (
	"log"

	"golang.org/x/net/context"
	"grpcdemo"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *grpcdemo.Message) (*grpcdemo.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
