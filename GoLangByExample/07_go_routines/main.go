package main

import (
	"fmt"
	"time"
)

// any return value from the goroutine will be ignored.
// All  goroutines are started from the main goroutine. These goroutines can then start multiple other goroutine and so on.
// The main goroutine represents the main program. Once it exits then it means that the program has exited.

func something() {
	fmt.Println("Inside a goroutine")
	go something2()
}
func something2() {
	fmt.Println("inside the 2nd goroutine")

}

func main() {

	go something()
	fmt.Println("Welcome to go routines")
	time.Sleep(time.Second * 2)
}
