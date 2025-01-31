package main

import (
	"fmt"
	"time"
)

// A buffered channel has some capacity to hold data hence for a buffered channel:
// Send on a buffer channel only blocks if the buffer is full
// Receive is only block is channel is empty
// So sending and receiving in the same goroutine is only possible for a buffered channel

// A channel to which we can only send data.
// chan<- int

// A channel from which we can only send data
// <-chan in

var total int = 0

func sum(ch chan int) {
	for val := range ch {
		total += val
	}

}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < cap(ch); i++ {
		ch <- i * 5
	}
	close(ch)
	// Also closing a already closed channel will cause a panic
	// For using a range with channel, we have to close the channel
	// Else the range would run indefinitely

	go sum(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("total of the number: ", total)
	fmt.Println("The value in the buffered channel is: ", <-ch)
	// fmt.Println("The value in the buffered channel is: ", <-ch)
	fmt.Println("capacity of the channel is: ", cap(ch))
	fmt.Println("Length of the channel is: ", len(ch))
}
