package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
)

// ProxyServer creates a TCP based proxy server
func ProxyServer(src net.Conn) {
	dst, err := net.Dial("tcp", "localhost:3001")
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
