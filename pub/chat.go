package pub

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
	// i don't know what this shit means...

}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
