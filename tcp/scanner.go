// Package tcp contains working code to simulate Netcats 'gaping security hole' in Go code
package tcp

import (
	"fmt"
	"net"
	"sort"
)

// Try to establish a TCP connection on the port
func worker(addr string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", addr, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// PortScan performs scans a portRange for a specified web address
func PortScan(portRange int, address string) []int {

	ports := make(chan int, 100)
	results := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(address, ports, results)
	}

	go func() {
		for i := 1; i <= portRange; i++ {
			ports <- i
		}
	}()

	for i := 0; i < portRange; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
	return openports

}

// SelPortScan lets a user pass in a slice of ports to  be scanned rather than iterating from start to finish of a range.
func SelPortScan(ports []int, address string) []int {
	var openPorts []int

	for _, p := range ports {

		addr := fmt.Sprintf("%s:%d", address, p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		openPorts = append(openPorts, p)
		conn.Close()
	}
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

	return openPorts
}
