package tcpserver

import (
	"fmt"
	"net"
)

func StartListener(host string, port int) *net.Listener {
	// start a TCP listener using the net package
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	check(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection:", err)
		}
		go handleConnection(conn)
	}
	return &listener
}
