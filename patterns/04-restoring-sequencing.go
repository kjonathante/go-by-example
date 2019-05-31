package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Restoring sequencing
// Send a channel on a channel, making goroutine wait its turn.
// Receive all messages, then enable them again by sending on a private channel.
// First we define a message type that contains a channel for the reply.

type Message struct {
	str string
	wait chan bool
}

func main() {
	c := fanIn(boring("Joe"), boring("Kit"))
	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan Message { // Returns receive-only channel of strings.
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn(inputs ... <-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() { for { c <- <-input } }()
	}
	return c
}