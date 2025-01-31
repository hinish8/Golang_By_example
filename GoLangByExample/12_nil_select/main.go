package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go hello(ch1)
	go world(ch2)

	for i := 0; i < 4; i++ {
		select {
		case <-ch1:
			fmt.Println("Hello")
			ch1 = nil // channel set to nil, will be ignored in further calls
		case <-ch2:
			fmt.Println("World")
		}
	}
}

func hello(ch1 chan int) {
	time.Sleep(time.Millisecond * 100)
	ch1 <- 1
	time.Sleep(time.Millisecond * 300)
	ch1 <- 1
}

func world(ch2 chan int) {
	time.Sleep(time.Millisecond * 200)
	ch2 <- 1
	time.Sleep(time.Millisecond * 400)
	ch2 <- 1
}
