package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/llcranmer/eff/shodan"
	"github.com/llcranmer/eff/tcp"
)

func main() {

	tcpPtr := flag.String("tcp", "tcp", "Tcp protocol with a range of ports to scan or a selection of ports to scan")
	addrPtr := flag.String("addr", "127.0.0.1", "the address to scan")
	numbPtr := flag.Int("portRange", 1024, "Scan from 0 up to the number inputted.")
	portsPtr := flag.String("ports", "8080,8000", "selection of ports to scan in csv format")
	shodPtr := flag.String("shod", "uinfo", "Interact with shodan.io")
	qPtr := flag.String("q", "localhost", "Query string to pass to search flag")


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

	if *shodPtr == "uinfo" {
		fmt.Println("Showing user account info...assuming SHODAN_API_KEY is set.")
		apiKey := os.Getenv("SHODAN_API_KEY")
		s := shodan.New(apiKey)
		info, err := s.APIInfo()
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf(
			"Query Credits: %d\nScan Credits: %d\n\n",
			info.QueryCredits,
			info.ScanCredits)
	}

	if *shodPtr == "search" {
		apiKey := os.Getenv("SHODAN_API_KEY")
		s := shodan.New(apiKey)
		hostSearch, err := s.HostSearch(*qPtr)
		if err != nil {
			log.Panicln(err)
		}
		for _, host := range hostSearch.Matches {
			fmt.Printf("%18s%8d\n", host.IPString, host.Port)
		}
	}
}
