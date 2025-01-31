package main

import (
	"fmt"
	"time"
)

func sendOnly(a chan<- int) {
	a <- 10
	// fmt.Println(<-a)
}

func receiveOnly(a <-chan int) {
	fmt.Println(<-a)
	// a <- 10
}

func main() {
	a := make(chan int, 5)

	go sendOnly(a)
	go receiveOnly(a)

	time.Sleep(time.Second * 2)

}
