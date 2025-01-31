package main

import (
	"fmt"

	"calc/operation" // Import the operation package
)

func main() {
	var a int
	var b int
	var op int
	fmt.Println("Enter the first number")
	fmt.Scan(&a)
	fmt.Println("Enter the second number")
	fmt.Scan(&b)
	fmt.Println("Enter the number according to the operation")
	fmt.Println("Enter 1 for additon")
	fmt.Println("Enter 2 for Subtraction")
	fmt.Println("Enter 3 for Multiplication")
	fmt.Println("Enter 4 for Division")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println("Result of addition is: ", operation.Add(a, b))
	case 2:
		fmt.Println("Result of subtraction is: ", operation.Sub(a, b))
	case 3:
		fmt.Println("Result of multiplication is: ", operation.Mul(a, b))
	case 4:
		fmt.Println("Result of division is: ", operation.Div(a, b))
	default:
		fmt.Println("Invalid operation")
	}

}
