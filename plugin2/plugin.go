package main

import (
	"fmt"
)

type greeting struct{}

var Greeter greeting

func (greeting) Greet() {
	fmt.Println("Hello from plugin 2")
}

func (greeting) Reload(chan<- bool) {
	return
}
