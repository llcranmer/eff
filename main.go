package main

import (
	"fmt"

	"github.com/llcranmer/eff/tcp"
)

func main() {
	fmt.Println("hello from the main file")
	tcp.PortScan(1024, "127.0.01")
	fmt.Println("done with scan")
}
