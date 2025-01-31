package main

// Unbuffered channel does not have any storage hence for an unbuffered channel
// Send on a channel is block unless there is other goroutine to receive.
// Receive is block until there is other goroutine on the other side to send.

import (
	"fmt"
	"time"
)

func receive(a chan int) {
	fmt.Println("Inside the receive function")
	time.Sleep(time.Second * 2)
	fmt.Println("The timeout has finished")
	v := <-a
	fmt.Println("value received is: ", v)
}

func send(a chan int) {
	a <- 1
	fmt.Println("value send to channel")
}

func main() {
	var a chan int
	fmt.Println(a) // nil channel by default

	a = make(chan int)
	fmt.Println(a) //the starting address is returned

	go send(a)
	go receive(a)

	time.Sleep(time.Second * 4)

}
