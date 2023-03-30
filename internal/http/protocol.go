package http

// GetRequestFromBytes reads the byte data up from the TCP layer and assembles a TCP packet
func GetRequestFromBytes(tcpData []byte) (*HttpRequest, error) {

	httpRequest := HttpRequest{
		ClientSocket: nil,
		Headers:      nil,
		Body:         nil,
	}
	return &httpRequest, nil
}

// GetBytesFromResponse serializes the HTTP request into a byte slice
func GetBytesFromResponse(response *HttpResponse) ([]byte, error) {

	responseBytes := make([]byte, 0)
	return responseBytes, nil
}
