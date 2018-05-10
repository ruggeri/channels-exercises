package main

import (
	"fmt"
	"net"
)

// Listens for ("accepts") an incoming TCP connection. "Blocks": if no
// TCP connection is incoming, will just sit and do nothing.
func listenForConnection(listener net.Listener) {
	fmt.Println("Beginning to listen!")
	// The Accept method is the one that blocks.
	c, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	// See how connections are handled in handling.go.
	handleConnection(c)
}

// Starts listening for TCP connections on port 8080. Will run a service
// that allows a user to upload a string, and we will write back the
// string in reverse order.
func listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		listenForConnection(listener)
	}
}
