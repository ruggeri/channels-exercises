package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// A "logger" is really just a string. The string is a (somewhat) unique
// identifier. All messages logged by this logger will be prefixed by
// the identifier.
type logger string

// This function builds a logger for you with a random hex identifier.
func randomLogger() logger {
	// Samples a random 64 bit non-negative integer. The 64th bit is used
	// for sign (positive or negative).
	randomNumber := rand.Int63()
	// Converts the number to a hexadecimal string.
	randomIdentifier := strconv.FormatInt(randomNumber, 16)
	// Even though a string *is* the same as a logger, I want to tell
	// golang to treat this string *as* a logger. This lets me define and
	// call methods on the logger.
	l := logger(randomIdentifier)

	return l
}

// My own fancy printf function. Notice how fmt.Printf takes in a
// variable number of arguments? We call a function like that
// *variadic*. By using `...`, the variable number of args is captured
// in a slice.
//
// This function works just like fmt.Printf, except it always tacks on
// the identifier at the start of the line.
func (l logger) printf(s string, args ...interface{}) {
	// Makes a new slice, with the logger's identifier at the front.
	newArgs := make([]interface{}, len(args)+1)
	newArgs[0] = string(l)
	// Copies over the old arguments into the newArgs slice. The first
	// argument of `copy` is the destination: this follows traditional C
	// convention.
	copy(newArgs[1:], args)

	// Piggyback on the original Printf function. See how I unrolled the
	// newArgs with a `...`. Golang doesn't let you do something like
	// f(newArg, oldArgs...). All arguments need to be in the slice if
	// you're going to use `...`.
	fmt.Printf("%v: "+s, newArgs...)
}
