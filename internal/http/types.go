package http

import "net"

// HttpRequest - the request structure. Represents the type
type HttpRequest struct {
	ClientSocket *net.TCPConn
	Headers      map[string]string
	Body         []byte
}

// HttpResponse - the response structure.
type HttpResponse struct {
	ClientSocket *net.TCPConn
	StatusCode   int
	StatusText   string
	Headers      map[string]string
	Body         []byte
}
