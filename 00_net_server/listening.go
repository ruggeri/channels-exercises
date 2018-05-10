package main

import (
	"fmt"
	"net"
)

func listenForConnection(listener net.Listener) {
	fmt.Println("Beginning to listen!")
	c, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	handleConnection(c)
}

func listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		listenForConnection(listener)
	}
}
