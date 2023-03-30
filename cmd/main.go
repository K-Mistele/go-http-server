package main

import (
	"fmt"
	"github.com/k-mistele/go-http-server/internal/config"
	"github.com/k-mistele/go-http-server/internal/tcpserver"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var serverConfiguration *config.Config
	var port string
	var host string
	var server *tcpserver.Server
	var err error

	// Parse arguments. Check for port flag and listen host flag. if present, start a default tcpserver,
	// otherwise, start a tcpserver with the specified port and host.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-http-server port [host]")
		os.Exit(1)
	}
	port = os.Args[1]
	if len(os.Args) >= 3 {
		host = os.Args[2]
	} else {
		host = "localhost"
	}

	if port == "" {
		fmt.Println("Port not specified. Starting tcpserver with default port 8080")
		port = "8080"
	}
	if host == "" {
		fmt.Println("Host not specified. Starting tcpserver with default host localhost")
		host = "localhost"
	}
	// Convert port and create a server
	numPort, err := strconv.Atoi(port)
	check(err)
	serverConfiguration = config.NewConfigWithOptions(host, numPort)

	// Create the tcpserver
	server = tcpserver.NewServerWithOptions(serverConfiguration)
	err = server.Start()
	check(err)
}
