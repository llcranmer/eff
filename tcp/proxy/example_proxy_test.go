package tcp

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestProxyServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}
	fmt.Println("Listening on 127.0.0.1:3001")
	for {
		conn, err := listener.Accept()
		// log.Println("Received connection")
		if err != nil {
			log.Println("Unable to accept connection")
			break
		}
		go ProxyServer(conn)

	}
	fmt.Println("Successfully connected.")

	// Output:
	// Listening on 127.0.0.1:3001
	// Successfully connected.
	// Unable to connect to our unreachable host.
}
