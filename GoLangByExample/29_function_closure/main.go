package main

import (
	"fmt"
)

// The getModulus function returns a closure. It is assigned to a variable modulus
// This closure function can access the count variable defined outside its body.
// The value of the count variable is retained between different function calls of modulus function
func main() {
	modulus := getModulus()
	modulus(-1)
	modulus(2)
	modulus(-5)
}

func getModulus() func(int) int {
	count := 0
	return func(x int) int {
		count = count + 1
		fmt.Printf("modulus function called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}

// Function closures are nothing but an anonymous function
// that can access variables declared outside the function
// and also retain the current value of those variables between different function calls.
// Anonymous functions are functions that are not named.

// A closure happens when a function is defined within a different function
//  and the inner function can access the variable of the outer function
