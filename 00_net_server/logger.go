package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type logger string

func (l logger) printf(s string, args ...interface{}) {
	newArgs := make([]interface{}, len(args)+1)
	newArgs[0] = l
	copy(newArgs[1:], args)
	fmt.Printf("%v: "+s, newArgs...)
}

func randomLogger() logger {
	randomNumber := rand.Int31()
	randomIdentifier := strconv.FormatInt(int64(randomNumber), 16)
	l := logger(randomIdentifier)

	return l
}
