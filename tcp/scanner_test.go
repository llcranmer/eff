package tcp

import (
	"fmt"
	"testing"
)

func TestPortScan(t *testing.T) {
	have := PortScan(1024, "127.0.0.1")
	if have != nil {
		fmt.Printf("have := %d, want := nil", have)
	}
}

func TestSelPortScan(t *testing.T) {
	ports := []int{8080, 3000, 3001}
	address := "127.0.0.1"
	fmt.Println(ports)
	have := SelPortScan(ports, address)
	if have == nil {
		fmt.Println("error")
	}
}
