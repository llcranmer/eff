package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// ProxyServer creates a TCP based proxy server
// to interact with the server locally run
// curl -i -X get localhost:3001 which will return an
// empty reply  from the server
func ProxyServer(src net.Conn, addrAndPort string) {
	dst, err := net.Dial("tcp", addrAndPort)
	if err != nil {
		fmt.Println("Unable to connect to our unreachable host.")
	}
	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

// StartListener creates a listener on the specified port in the form of
// :port, e.g. :3001
func StartListener(addrAndPort string) {
	port := strings.Split(addrAndPort, ":")
	listener, err := net.Listen("tcp", ":"+port[1])
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}
	fmt.Printf("Listening on %s", addrAndPort)
	// continously listen and accetp connections..
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Unable to accept connection")
			break
		}
		go ProxyServer(conn, addrAndPort)
	}

}
