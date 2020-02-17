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
	s := "8080,3001,3000"
	address := "127.0.0.1"
	fmt.Println(s)
	have := SelPortScan(s, address)
	if have == nil {
		fmt.Println("error")
	}
}
