package main

import (
	"fmt"
)

type MyError struct {
	// fields
	x uint32
	y uint64
}

type Set map[string]struct{}

func (e *MyError) Error() string {
	return "this is my error"
}

func doSomething() error {
	// logic...
	if true {
		var m *MyError
		return m
	}
	return nil
}

func main() {
	var err error
	err = doSomething()
	fmt.Println(err == nil)
}
