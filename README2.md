source: https://golangbot.com/channels/

chan T is a channel of type T

channel has to be defined using make

short hand declaration
```go
a := make(chan int)
```
```go
data := <- a // read from channel a  
a <- data // write to channel a  
```

Sends and receives are blocking by default

When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. 

Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.

```go
package main

import "fmt"

func sendData(sendch chan<- int) { // directional chan 
    sendch <- 10
}

func main() {  
    chnl := make(chan int) // bidirectional chan
    go sendData(chnl) // bidirectional to unidirectional conversion
    fmt.Println(<-chnl)
}
```

## Closing channels and for range loops on channels
```go
v, ok := <- ch  
```
