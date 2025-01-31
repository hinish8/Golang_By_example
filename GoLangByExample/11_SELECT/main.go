package main

import (
	"fmt"
	"time"
)

// The select blocks until any of the case statements are ready.
// If multiple case statements are ready then it selects one at random and proceeds
// that select can block if any of the case statement is not ready.

func main() {

	// the empty select statement will block the main thread
	// Resulting in a deadlock
	// select {}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)
	time.Sleep(time.Second * 5)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout happened and had to execute this")
	default:
		time.Sleep(time.Second * 2)
		fmt.Println("The select was about to go into blocking")
	}
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "from goTwo goroutine "
}

// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go goOne(ch1)
// 	go goTwo(ch2)

// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println(msg1)
// 	case ch2 <- "To goTwo goroutine":
// 	}
// }

// func goOne(ch chan string) {
// 	ch <- "From goOne goroutine"
// }

// func goTwo(ch chan string) {
// 	msg := <-ch
// 	fmt.Println(msg)
// }
