package tcp

import (
	"log"
	"net"
	"testing"
)

func TestWindowsExec(t *testing.T) {
	host := "windows"
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go Exec(conn, host)
	}
}

func TestLinuxExec(t *testing.T) {
	host := "linux"
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go Exec(conn, host)
	}

}
