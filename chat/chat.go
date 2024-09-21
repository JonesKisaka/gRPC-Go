package chat

import (
	"log"

	"golang.org/x/net/context"
)

// store details concerning the Server
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, m *Message) (*Message, error) {
	log.Printf("Received the message body from client: %s", m.Body)
	return &Message{Body: "Hello from the Server!"}, nil

}
