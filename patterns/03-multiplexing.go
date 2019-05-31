package main

import (
	"fmt"
	"math/rand"
	"time"
)

// docker run --interactive --tty --rm --volume $(pwd):/go golang:1.8 bash

// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers.
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

// Multiplexing
// We use a fan-in function to let whosoever is ready talk.
func fanIn(input1, input2 <-chan string) <-chan string {  // func fanIn(input1 <-chan string, input2 <-chan string) <-chan string
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

// Channels as a handle on a service
// Our boring function returns a channel that lets us communicate with the boring service it provides.
// We can have more instances of the service.
func main() {
	c := fanIn(boring("Joe"), boring("Kit"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}