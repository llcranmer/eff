package main

import (
	"flag"
	"fmt"

	"github.com/llcranmer/eff/tcp"
)

func main() {
	// tcp.PortScan(1024, "127.0.0.1")
	// fmt.Println("done with  ports scan")
	ports := []int{80, 20, 3001}
	tcp.SelPortScan(ports, "127.0.0.1")
	fmt.Println("done with sel port scan")

	wordPtr := flag.String("tcp", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	// boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	// fmt.Println("tcp:", *wordPtr)
	// fmt.Println("numb:", *numbPtr)
	// fmt.Println("fork:", *boolPtr)
	// fmt.Println("svar:", svar)
	// fmt.Println("tail:", flag.Args())

	if *wordPtr == "scan" {
		fmt.Println("Scanning ports..")
		tcp.PortScan(*numbPtr, "127.0.0.1")
		fmt.Println("done.")
	}
}
