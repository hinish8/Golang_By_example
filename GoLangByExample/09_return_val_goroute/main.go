package main

import (
	"fmt"
)

type res struct {
	add int
	sub int
}

func main() {
	result := make(chan int, 1)
	go sum(2, 3, result)
	value := <-result // buffered channel blocked as receiving
	fmt.Printf("Value: %d\n", value)
	close(result)

	ch := make(chan res)
	go op(5, 10, ch)
	val := <-ch
	fmt.Printf("Value: %d\nValue: %d\n", val.add, val.sub)
	close(ch)

}

func op(a, b int, ch chan res) {
	res := res{a + b, a - b}
	ch <- res
}

func sum(a, b int, result chan int) {
	sumValue := a + b
	result <- sumValue
	// return
}
