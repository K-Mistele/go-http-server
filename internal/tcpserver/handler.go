package tcpserver

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Contains the logic for handling incoming messages from clients

// handleConnection handles TCP connections
func handleConnection(conn net.Conn) {

	defer conn.Close()
	fmt.Printf("Handling connection from %s\n", conn.RemoteAddr())

	// Set deadlines for reading and writing
	err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Println("Failed to set read deadline for connection", err)
	}
	err = conn.SetWriteDeadline(time.Now().Add(8 * time.Second))
	if err != nil {
		fmt.Println("Failed to set write deadline for connection", err)
	}

	// Read from the connection until we see a \r\n. Close connection if there's an issue reading from stream
	delimiter := "\n"
	tcpMessage, err := getDataUntilDelimiter(conn, delimiter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read the message. If there's a PROTOCOL error, return an error
	fmt.Println("got message", string(tcpMessage))
	if string(tcpMessage[0:4]) == "ping" {
		_, err = conn.Write([]byte(fmt.Sprintf("pong%s", delimiter)))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	} else {
		_, err := conn.Write([]byte(fmt.Sprintf("error%s", delimiter)))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	return
}

// getDataUntilDelimiter reads bytes from the TCP stream until we hit the delimiter. Subject to timeouts.
func getDataUntilDelimiter(conn net.Conn, delimiter string) ([]byte, error) {
	var reader = bufio.NewReader(conn)
	var data []byte

	for {
		bytes, err := reader.ReadBytes('\n') // read until a newline
		if err != nil {
			return data, err
		}
		data = append(data, bytes...)

		// If we have read in the final delimiter, break
		if len(data) >= len(delimiter) && string(data[len(data)-len(delimiter):]) == delimiter {
			break
		}
	}
	return data, nil
}
