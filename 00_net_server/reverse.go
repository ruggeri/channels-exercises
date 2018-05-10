package main

import (
	"unicode/utf8"
)

// O(n**2) time.
func slowReverse(s string) string {
	result := ""
	for _, char := range s {
		// + creates copies of strings. This is an O(n) operation in the
		// inner loop. The problem is that strings are not mutable, so there
		// is no `concat` method like in Ruby.
		result = string(char) + result
	}

	return result
}

// DO NOT READ THIS.

func fastReverse(input string) string {
	// Treats the string as the sequence of underlying bytes. In ASCII,
	// each byte would be a character.
	inputBytes := []byte(input)
	inputBytesLen := len(inputBytes)

	// But all strings in Golang are *unicode*. Which means the decoding
	// is more complicated because of the possibility of multibyte
	// characters.

	// First, let's make space for the result. We know that the reverse
	// string will have the same number of bytes as the original. So I
	// make a slice of the correct length.
	resultBytes := make([]byte, inputBytesLen)

	totalBytesWritten := 0
	for totalBytesWritten < inputBytesLen {
		// Decodes the first unicode character and tells us how many bytes
		// long it was.
		//
		// First return value would be the rune value, but I'll actually
		// ignore that!
		_, numRuneBytes := utf8.DecodeRune(inputBytes)

		// Get the bytes for the rune by slicing the input. These are the
		// bytes to write.
		runeBytes := inputBytes[totalBytesWritten : totalBytesWritten+numRuneBytes]

		// Where to start writing the bytes. We need to write the bytes *in
		// order*. We're not reversing the order of the bytes inside an
		// individual rune.
		writeStartIdx := inputBytesLen - totalBytesWritten - numRuneBytes
		for runeByteIdx := 0; runeByteIdx < numRuneBytes; runeByteIdx++ {
			writeIdx := writeStartIdx + runeByteIdx
			resultBytes[writeIdx] = runeBytes[runeByteIdx]
			totalBytesWritten++
		}

		// Note could also have used the `copy` function.
	}

	// Treat the bytes as a unicode string now! All done!
	return string(resultBytes)
}
