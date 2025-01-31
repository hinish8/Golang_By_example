// package main

// import (
// 	"fmt"
// 	"runtime/debug"
// )

// func main() {
// 	a := []string{"a", "b"}
// 	checkAndPrint(a, 2)
// 	fmt.Println("Exiting normally")
// }
// func checkAndPrint(a []string, index int) {
// 	defer handleOutOfBounds()
// 	if index > (len(a) - 1) {
// 		panic("Out of bound access for slice")
// 	}
// 	fmt.Println(a[index])
// }
// func handleOutOfBounds() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovering from panic:", r)
// 		fmt.Println("Stack Trace:")
// 		debug.PrintStack() // prints the stack trace
// 	}
// }

package main

import "fmt"

func main() {
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	fmt.Println("Exiting normally")
}
func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	go checkAndPrint(a, index)
}
func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
