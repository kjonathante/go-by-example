// Select
// A control structure unique to concurrency.
// The reason channels and goroutines are built into the language.
// The select statement provides another way to handle multiple channels. 
// It's like a switch, but each case is a communication: 
// - All channels are evaluated. 
// - Selection blocks until one communication can proceed, which then does. 
// - If multiple can proceed, select chooses pseudo-randomly. 
// - A default clause, if present, executes immediately if no channel is ready.
// select {
// case v1 := <-c1:
// 		fmt.Printf("received %v from c1\n", v1)
// case v2 := <-c2:
// 		fmt.Printf("received %v from c2\n", v1)
// case c3 <- 23:
// 		fmt.Printf("sent %v to c3\n", 23)
// default:
// 		fmt.Printf("no one was ready to communicate\n")
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:  c <- s
			case s := <-input2:  c <- s
			}
		}
	}()
	return c
}

// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c := make(chan string)
// 	go func() { for { c <- <-input1 } }()
// 	go func() { for { c <- <-input2 } }()
// 	return c
// }