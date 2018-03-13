package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func init() {
	fmt.Println("Hello from init")
}

type greeting struct{}

var Greeter greeting

var greetFunction = func() { fmt.Println("hello") }

var message = func() { fmt.Println("do I terminate?") }

func (greeting) Reload(c chan<- bool) {
	f, err := os.Create("plugin/plugin.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(`package main

import "fmt"

type greeting struct{}

var Greeter greeting

func (greeting) Greet() {
	fmt.Println("I've changed slightly")
}
func (greeting) Reload(chan<- bool) {
	return
}
`)
	if err != nil {
		panic(err)
	}

	w.Flush()

	command := exec.Command("go", "build", "-buildmode=plugin", "-o", "plugin.so", "plugin/plugin.go")

	// set var to get the output
	var out bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	err = command.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(out.String())
	c <- true
}

func (greeting) Greet() {
	greetFunction()

	go func() {
		for {
			<-time.After(1 * time.Second)
			message()
		}
	}()
}
