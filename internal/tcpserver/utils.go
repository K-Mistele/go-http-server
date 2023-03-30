package tcpserver

// Check for an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}
