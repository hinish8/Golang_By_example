// This is an unsigned integer type which is large enough to hold any pointer address.
// Therefore its size is platform dependent.

// Arithmetic can be performed on the uintptr.
// Do note here arithmetic cannot be performed in a pointer in Go or unsafe.Pointer in Go.

// uintptr even though it holds a pointer address, is just a value and does not reference any object. Therefore
// Its value will not be updated if the corresponding object moves. Eg When goroutine stack changes
// The corresponding object can be garbage collected. The GC does not consider uintptr as live references and hence they can be garbage collected.

package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{a: 1, b: "test"}

	//Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	//Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))
}
