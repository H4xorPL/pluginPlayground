package main

import "fmt"

type greeting struct{}

var Greeter greeting

func (greeting) Greet() {
	fmt.Println("I've changed slightly")
}
func (greeting) Reload(chan<- bool) {
	return
}
