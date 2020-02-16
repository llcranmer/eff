package tcp_test

import (
	"fmt"

	tcp "github.com/llcranmer/eff/tcp"
)

func Example_portscan() {
	localhost := tcp.PortScan(1024, "127.0.0.1")
	if localhost != nil {
		fmt.Println("here are the ports")
	}
	// Output:
	//
}

func Example_selportscan() {
	ports := []int{80, 22, 3001}
	tcp.SelPortScan(ports, "scanme.nmap.org")

	// Output:
	// 80 open
	// 22 open
}
