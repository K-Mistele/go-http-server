package tcpserver

import (
	"errors"
	"fmt"
	"github.com/k-mistele/go-http-server/internal/config"
	"net"
)

// Server handles starting and stopping the listener using the configured options
type Server struct {
	Host     string
	Port     int
	listener *net.Listener
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

// Start launches the server and starts the tcp Listener
func (s *Server) Start() error {

	fmt.Printf("Starting server at %s:%d\n", s.Host, s.Port)
	s.listener = StartListener(s.Host, s.Port)
	return nil
}

// Stop stops the listener if it is running
func (s *Server) Stop() error {

	if s.listener == nil {
		return errors.New("Unable to stop a server that is not running!")
	} else {
		err := net.Listener.Close(*s.listener)
		return err
	}
}
