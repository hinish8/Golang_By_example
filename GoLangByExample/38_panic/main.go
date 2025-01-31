// Panic can occur in a program in two ways
// Runtime error in the program
// By calling the panic function explicitly. This can be called by the programmer when the program cannot continue and it has to exit

// When a panic happens in a  program it outputs two things
// The error message that is passed to the panic function as an argument
// Stack trace of where the panic happened

// The defered statements would be executed after the panic has occurred

package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	// print(a, 2)

	// fmt.Println(a[3])
	defer fmt.Println("Hello, world!", a)
	panic("Panic statement after defer")
	// fmt.Println("Print statement after panic")  // This will not be printed as it is unreachable code

}

func print(a []string, index int) {
	fmt.Println(a[index])
}
