package main

import (
	"bufio"
	"io"
	"net"
	"strings"
)

func handleConnection(connection net.Conn) {
	l := randomLogger()

	l.printf("Handling connection!\n")
	reader := bufio.NewReader(connection)

	s, err := reader.ReadString('\n')
	if err == io.EOF {
		l.printf("Connection closed before newline!\n")
		return
	} else if err != nil {
		l.printf("Unexpected error: %v", err)
		panic(err)
	}

	s = strings.TrimSpace(s)
	l.printf("Reversing the string %v\n", s)
	reversedString := reverse(s) + "\n"
	stringBytes := []byte(reversedString)
	numBytesWritten, err := connection.Write(stringBytes)

	if err == io.EOF {
		l.printf("Connection was closed before all %v bytes written!\n", numBytesWritten)
		return
	} else if err != nil {
		l.printf("Unexpected error: %v", err)
		panic(err)
	}

	l.printf("All %v bytes successfully written!\n", numBytesWritten)

	err = connection.Close()
	if err != nil {
		l.printf("Unexpected error: %v", err)
		panic(err)
	}
	l.printf("Connection closed!\n")
}
