package main

import (
	"Backend/cisco"
	"Backend/resources"
	"Backend/server"
	"fmt"
	"net/http"
)

var (
	httpSrv *http.Server
)

func main() {
	fmt.Println("Starting")

	// Create a server instance
	srv := server.New("0.0.0.0:162")
	// Start the trap listener in another thread
	go startTrapListener(srv)
	// Start the file server (Serving on port 80) in another thread
	go makeAndServeFileServer()
	// Start the API server in the main thread
	makeAndStartHTTPServer(srv)
}
//startTrapListener starts the trap listener
// TODO update the credentials - not quite sure they're even needed anymore?..
func startTrapListener(srv *server.Server) {
	cisco.SnmpTrapListener(resources.SnmpCredentials{
		Host:           "192.168.1.5",
		User:           "MYUSER",
		Authentication: "MYPASS123",
		Privacy:        "MYKEY123",
	}, srv)
}
