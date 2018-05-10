package main

import (
	"bufio"
	"io"
	"net"
	"strings"
)

// This function "handles" a TCP connection by reading in a string, then
// reversing it.
func handleConnection(connection net.Conn) {
	// This is a "logger" class I wrote. Check out logger.go briefly.
	l := randomLogger()

	l.printf("Handling connection!\n")

	// A `net.Conn` value has a method (`Read`) that lets you read in a
	// sequence of *bytes*. We'd rather read in strings. This is what a
	// buffered reader does.
	reader := bufio.NewReader(connection)

	// Here I read a line. It says: reader, read bytes until you hit a
	// newline byte. When you get there, use everything you've read as
	// characters of a string.
	s, err := reader.ReadString('\n')
	if err == io.EOF {
		// The user might not type newline before closing the connection!
		l.printf("Connection closed before newline!\n")
		return
	} else if err != nil {
		// Maybe something else terrible happens :-P
		l.printf("Unexpected error: %v", err)
		panic(err)
	}

	// This reverses the string. First it removes the newline at the end.
	// `reverse` is defined in reverse.go.
	s = strings.TrimSpace(s)
	l.printf("Reversing the string %v\n", s)
	reversedString := fastReverse(s) + "\n"

	// When writing back to the user, I won't bother with a buffered
	// writer and will convert to bytes myself. I can then write the raw
	// bytes directly. This is what a writer class would do for me.
	stringBytes := []byte(reversedString)
	numBytesWritten, err := connection.Write(stringBytes)

	if err == io.EOF {
		// User could close the connection before getting a response. They
		// may even close the connection after receiving only *some* of the
		// response.
		l.printf("Connection was closed with only %v bytes written!\n", numBytesWritten)
		return
	} else if err != nil {
		l.printf("Unexpected error: %v", err)
		panic(err)
	}

	l.printf("All %v bytes successfully written!\n", numBytesWritten)

	// Finally, we can close the connection. The client on the other side
	// will see we hung up and they will know the response is complete.
	err = connection.Close()
	if err != nil {
		l.printf("Unexpected error: %v", err)
		panic(err)
	}
	l.printf("Connection closed!\n")
}
