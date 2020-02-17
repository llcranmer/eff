package main

import (
	"flag"
	"fmt"

	cmd "github.com/llcranmer/eff/tcp/cmd"
	proxy "github.com/llcranmer/eff/tcp/proxy"
	tcp "github.com/llcranmer/eff/tcp/scanner"
)

func main() {

	tcpPtr := flag.String("tcp", "tcp", "tcp protocol with a range of ports to scan or a selection of ports to scan. Also supports proxy server in the form of addr:port, e.g. localhost:3001")
	addrPtr := flag.String("addr", "127.0.0.1", "the address to scan")
	numbPtr := flag.Int("portRange", 1024, "Scan from 0 up to the number inputted.")
	portsPtr := flag.String("ports", "8080,8000", "selection of ports to scan in csv format")
	hostPtr := flag.String("host", "linux", "host os to open terminal on supports windows or linux")
	// boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	if *tcpPtr == "scan" {
		fmt.Printf("Scanning ports 0:%d", *numbPtr)
		tcp.PortScan(*numbPtr, *addrPtr)
		fmt.Println("done.")
	}

	if *tcpPtr == "sscan" {
		fmt.Println("Scanning selected ports..")
		tcp.SelPortScan(*portsPtr, *addrPtr)
		fmt.Println("done.")
	}

	if *tcpPtr == "proxy" {
		fmt.Printf("Starting a proxy server on %s", *addrPtr)
		proxy.StartProxyListener(*addrPtr)
		fmt.Println("done.")
	}

	if *tcpPtr == "cmd" {
		fmt.Printf("Launching command line interface for %s", *hostPtr)
		cmd.StartExec(*hostPtr, *portsPtr)
		fmt.Println("done.")
	}
}
