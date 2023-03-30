package server

import (
	"github.com/k-mistele/go-http-server/internal/config"
)

// Server handles starting and stopping the listener using the configured options
type Server struct {
	Host string
	Port int
}

// NewServer creates a new server using defaults
func NewServer() *Server {

	serverConfig := config.NewConfig()
	return &Server{
		Host: serverConfig.ListenHost,
		Port: serverConfig.ListenPort,
	}
}

// NewServerWithOptions creates a server using a configuration object
func NewServerWithOptions(configuration config.Config) *Server {
	return &Server{
		Host: configuration.ListenHost,
		Port: configuration.ListenPort,
	}
}
