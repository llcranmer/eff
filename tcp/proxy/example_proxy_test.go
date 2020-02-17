package tcp

import (
	"testing"
)

func TestProxyServer(t *testing.T) {
	// setting up a listener
	addrAndPort := "localhost:3001"
	StartProxyListener(addrAndPort)

	// Output:
	// Listening on 127.0.0.1:3001
	// Unable to connect to our unreachable host.
}
