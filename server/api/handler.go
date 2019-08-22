package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

func (s *Server) HealthCheck(ctx context.Context, in *HealthMessage) (*HealthMessage, error) {
	return &HealthMessage{Health: "Health is good ! ", Status: true}, nil
}
