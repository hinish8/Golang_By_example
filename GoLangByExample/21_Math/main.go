package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ceil function")
	res := math.Ceil(1.6)
	fmt.Println(res)
	res = math.Ceil(-1.6)
	fmt.Println(res)
	res = math.Ceil(1)
	fmt.Println(res)

	fmt.Println("Floor function")
	res1 := math.Floor(1.6)
	fmt.Println(res1)
	res1 = math.Floor(-1.6)
	fmt.Println(res1)
	res1 = math.Floor(1)
	fmt.Println(res1)

	fmt.Println("Trunc Function-- float to int")
	res3 := math.Trunc(1.6)
	fmt.Println(res3)
	res3 = math.Trunc(-1.6)
	fmt.Println(res3)
	res3 = math.Trunc(1)
	fmt.Println(res3)

	fmt.Println("Round Function")
	res = math.Round(1.6)
	fmt.Println(res)
	res = math.Round(1.5)
	fmt.Println(res)
	res = math.Round(1.4)
	fmt.Println(res)
	res = math.Round(-1.6)
	fmt.Println(res)
	res = math.Round(-1.5)
	fmt.Println(res)
	res = math.Round(-1.4)
	fmt.Println(res)
	res = math.Round(1)
	fmt.Println(res)

	fmt.Println("Round to Even")
	res = math.RoundToEven(0.5)
	fmt.Println(res)
	res = math.RoundToEven(1.5)
	fmt.Println(res)
	res = math.RoundToEven(2.5)
	fmt.Println(res)
	res = math.RoundToEven(3.5)
	fmt.Println(res)
	res = math.RoundToEven(4.5)
	fmt.Println(res)

	fmt.Println("Get absolute value")
	res = math.Abs(-2)
	fmt.Println(res)
	res = math.Abs(2)
	fmt.Println(res)
	res = math.Abs(3.5)
	fmt.Println(res)
	res = math.Abs(-3.5)
	fmt.Println(res)

	fmt.Println("Pi value is a constant in go's math package")
	pi := math.Pi
	fmt.Println(pi)

	fmt.Println("Square root of a number ")
	res = math.Sqrt(4)
	fmt.Println(res)
	res = math.Sqrt(9)
	fmt.Println(res)
	res = math.Sqrt(30.33)
	fmt.Println(res)
	res = math.Sqrt(-9)
	fmt.Println(res)

	fmt.Println("Cube root")
	res = math.Cbrt(8)
	fmt.Println(res)
	res = math.Cbrt(27)
	fmt.Println(res)
	res = math.Cbrt(30.33)
	fmt.Println(res)

	fmt.Println("Break number into integer and fractional part")
	integer, fraction := math.Modf(2.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)
	integer, fraction = math.Modf(2)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)
	integer, fraction = math.Modf(-2.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)
	integer, fraction = math.Modf(0.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)

	fmt.Println("Pow function")
	res = math.Pow(2, 10)
	fmt.Println(res)
	res = math.Pow(1.5, 2)
	fmt.Println(res)
	res = math.Pow(3, 0)
	fmt.Println(res)

	fmt.Println("Check if number is +ve or -ve")
	b := math.Signbit(4)
	fmt.Println(b)
	b = math.Signbit(-4)
	fmt.Println(b)
	b = math.Signbit(0)
	fmt.Println(b)
	b = math.Signbit(-0)
	fmt.Println(b)

	fmt.Println("Minimum of two numbers")
	min := math.Min(2, 3)
	fmt.Println(min)
	min = math.Min(-2.1, -3.3)
	fmt.Println(min)

	fmt.Println("Maximum of twonumbers ")
	max := math.Max(2, 3)
	fmt.Println(max)
	max = math.Max(-2.1, -3.3)
	fmt.Println(max)

	fmt.Println("Natural log")
	res = math.Log(4)
	fmt.Println(res)

	fmt.Println("Log to the base e")
	res = math.Logb(4)
	fmt.Println(res)

	fmt.Println("Log to the base 2")
	res = math.Log2(4)
	fmt.Println(res)

	fmt.Println("Log to the base 10")
	res = math.Log10(100)
	fmt.Println(res)

	fmt.Println("Simple mod operator")
	res = 4 % 2
	fmt.Println(res)
	res = 5 % 2
	fmt.Println(res)
	res = 8 % 3
	fmt.Println(res)

	fmt.Println("Mod function for float numbers ")
	res = math.Mod(4, 2)
	fmt.Println(res)
	res = math.Mod(4.2, 2)
	fmt.Println(res)
	res = math.Mod(5, 2)
	fmt.Println(res)
	res = math.Mod(-5, 2)
	fmt.Println(res)

	fmt.Println("Remainder -- IEEE754")
	res = math.Remainder(4, 2)
	fmt.Println(res)
	res = math.Remainder(5, 2)
	fmt.Println(res)
	res = math.Remainder(5.5, 2)
	fmt.Println(res)
	res = math.Remainder(5.5, 1.5)
	fmt.Println(res)

}
