package main

import (
	"fmt"
	"os"
	"plugin"
	"time"
)

type Greeter interface {
	Greet()
	Reload(chan<- bool)
}

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	symGreeter, err := p.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// go func() {
	// 	for {
	// 		<-time.After(2 * time.Second)
	// 		greeter.Greet()
	// 	}
	// }()

	greeter.Greet()
	reloadChan := make(chan bool)
	<-time.After(5 * time.Second)
	func() {
		go greeter.Reload(reloadChan)
	}()

	if <-reloadChan {
		fmt.Println("reloaded")
		p, err = plugin.Open("plugin.so")
		if err != nil {
			panic(err)
		}

		symGreeter, err = p.Lookup("Greeter")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		greeter, ok = symGreeter.(Greeter)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		fmt.Print("Greet after Reload: ")
		greeter.Greet()
	}

	time.Sleep(time.Second * 10)
}
