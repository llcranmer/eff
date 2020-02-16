package main

import (
	"fmt"

	"github.com/llcranmer/eff/tcp"
)

func main() {
	fmt.Println("hello from the main file")
	tcp.PortScan(1024, "scanme.nmap.org")
	fmt.Println("done with  ports scan")
	ports := []int{80, 20, 3001}
	tcp.SelPortScan(ports, "scanme.nmap.org")
	fmt.Println("done with sel port scan")
}
