package main

import (
	"fmt"
	"runtime"
	"time"
)

var total int = 0

func add() {
	total += 1

}

func display(val int) {
	fmt.Println(val)
}

func main() {

	for i := 0; i < 10000; i++ {
		go add()
	}
	fmt.Println("Number of running goroutines, ", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go display(i)
	}
	fmt.Println("The total sum by 10000 goroutines is: ", total)
	fmt.Println("Number of CPU cores:", runtime.NumCPU())

	time.Sleep(time.Second * 3)

}
