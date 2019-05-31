# Source:
## Go by Example
## by Mark McGranaghan
https://gobyexample.com/

[additional info](./README2.md)

docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go run app.go
docker run --interactive --tty --rm --volume $(pwd):/go golang:1.8 bash

## Channels
#### https://talks.golang.org/2012/concurrency.slide#19
#### A channel in Go provides a connection between two goroutines, allowing them to communicate.
```go
// Declaring and initializing.
var c chan int
c = make(chan int)
// or
c := make(chan int)
```
```go
// Sending on a channel.
c <- 1
```
```go
// Receiving from a channel.
// The "arrow" indicates the direction of data flow.
value = <-c
```

## Using channels
#### A channel connects the main and boring goroutines so they can communicate.
```go
func main() {
    c := make(chan string)
    go boring("boring!", c)
    for i := 0; i < 5; i++ {
        fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
    }
    fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan string) {
    for i := 0; ; i++ {
        c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
        time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
}
```
Synchronization

When the main function executes <–c, it will wait for a value to be sent.
Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.
Thus channels both communicate and synchronize.

Go channels can also be created with a buffer.

Buffering removes synchronization.