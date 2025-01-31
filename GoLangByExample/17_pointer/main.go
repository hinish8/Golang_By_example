package main

import "fmt"

func main() {

	// a == *&a
	// b == *&b == &*b
	// Default zero value of a pointer is nil.

	//Declare
	var b *int
	a := 2
	b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(*b)
	a = 500
	fmt.Println(a)
	fmt.Println(*b)

	b = new(int)
	*b = 10
	fmt.Println(*b)

	var h *int
	fmt.Print("Default Zero Value of a pointer: ")
	fmt.Println(h)

}
