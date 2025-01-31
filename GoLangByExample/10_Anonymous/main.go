package main

// A function in Go is a first-class variable so it can be used as a value as well. When using a function as a value, it is not named and can be assigned to a variable. Such a function is called anonymous functions because the function is not named.

import "fmt"

var max = func(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	res := max(2, 3)
	fmt.Println(res)

	// immediately invoked anonymous func
	func() {
		fmt.Println("From anoymous function")
	}()
}
