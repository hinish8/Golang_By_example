// package main

// import "fmt"

// func main() {
// 	areaF := getAreaFunc()
// 	print(3, 4, areaF)
// }

// func print(x, y int, area func(int, int) int) {
// 	fmt.Printf("Area is: %d\n", area(x, y))
// }

// func getAreaFunc() func(int, int) int {
// 	return func(x, y int) int {
// 		return x * y
// 	}
// }

package main

import "fmt"

func main() {
	add, subtract := getAddSubtract()
	print(3, 4, add, subtract)
}

func print(x, y int, add func(int, int) int, subtract func(int, int) int) {
	fmt.Printf("Sum is: %d\n", add(x, y))
	fmt.Printf("Difference Value is: %d\n", subtract(x, y))
}

func getAddSubtract() (func(int, int) int, func(int, int) int) {
	add := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	return add, subtract
}
