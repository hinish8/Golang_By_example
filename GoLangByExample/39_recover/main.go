// We already learn above that defer function is the only function that is called after the panic.
// So it makes sense to put the recover function in the defer function only.
// If the recover function is not within the defer function then it will not stop panic.

package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint(a []string, index int) {
	defer handleoutOfBounds()
	if index > (len(a) - 1) {
		panic("Index out of bounds")
	}
	fmt.Println(a[index])
}

// if the recover had not been in the defer, then we would not be able to recover
// The recover function will catch the panic and we can also print the message from the panic.
func handleoutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic: ", r)
	}
}
