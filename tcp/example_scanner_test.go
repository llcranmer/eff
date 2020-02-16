package tcp_test

import tcp "github.com/llcranmer/eff/tcp"

func Example_portscan() {
	tcp.PortScan(1024, "127.0.0.1")

	// Output:
	//
}

func Example_selportscan() {
	ports := []int{80}
	tcp.SelPortScan(ports, "scanme.nmap.org")

	// Output:
	// [80]
}
