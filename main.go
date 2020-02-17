package main

import (
	"flag"
	"fmt"

	"github.com/llcranmer/eff/tcp"
)

func main() {

	tcpPtr := flag.String("tcp", "tcp", "tcp protocol with a range of ports to scan or a selection of ports to scan")
	addrPtr := flag.String("addr", "127.0.0.1", "the address to scan")
	numbPtr := flag.Int("portRange", 1024, "Scan from 0 up to the number inputted.")
	portsPtr := flag.String("ports", "8080,8000", "selection of ports to scan in csv format")
	// boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	if *tcpPtr == "scan" {
		fmt.Println("Scanning ports..")
		tcp.PortScan(*numbPtr, *addrPtr)
		fmt.Println("done.")
	}

	if *tcpPtr == "sscan" {
		fmt.Println("Scanning selected ports..")
		tcp.SelPortScan(*portsPtr, *addrPtr)
		fmt.Println("done.")
	}

}
