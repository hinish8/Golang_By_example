package operation

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		fmt.Println("Divisor Cannot be Zero")
		return 0
	}
	return a / b
}
