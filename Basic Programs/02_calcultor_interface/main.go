package main

import "fmt"

type Calculator interface {
	Compute(a, b int) int
}

// all 4 types of struct satisfy the interface
type Add struct{}
type Subtract struct{}
type Multiply struct{}
type Divide struct{}

func (Add) Compute(a, b int) int {
	return a + b
}

func (Subtract) Compute(a, b int) int {
	return a - b
}

func (Multiply) Compute(a, b int) int {
	return a * b
}

func (Divide) Compute(a, b int) int {
	if b != 0 {
		return a / b
	} else {
		return 0
	}
}

func main() {
	var a, b, op int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)

	fmt.Println("Choose operation: \n1.Add \n2.Subtract \n3.Multiply \n4.Divide")
	fmt.Scan(&op)

	var calc Calculator

	switch op {
	case 1:
		calc = Add{}
	case 2:
		calc = Subtract{}
	case 3:
		calc = Multiply{}
	case 4:
		calc = Divide{}
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println("Result:", calc.Compute(a, b))
}
