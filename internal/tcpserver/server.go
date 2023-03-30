package tcpserver

import (
	"fmt"
	"github.com/k-mistele/go-http-server/internal/config"
)

// Server handles starting and stopping the listener using the configured options
type Server struct {
	Host string
	Port int
}

// NewServer creates a new tcpserver using defaults
func NewServer() *Server {

	serverConfig := config.NewConfig()
	return &Server{
		Host: serverConfig.ListenHost,
		Port: serverConfig.ListenPort,
	}
}

// NewServerWithOptions creates a tcpserver using a configuration object
func NewServerWithOptions(configuration *config.Config) *Server {
	return &Server{
		Host: configuration.ListenHost,
		Port: configuration.ListenPort,
	}
}

func (s *Server) Start() error {

	fmt.Printf("Starting server at %s:%d\n", s.Host, s.Port)
	return nil
}
