package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers.
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

// Channels as a handle on a service
// Our boring function returns a channel that lets us communicate with the boring service it provides.
// We can have more instances of the service.
func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	kit := boring("Kit")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
		fmt.Println(<-kit)
	}
	fmt.Println("You're both boring; I'm leaving.")
}