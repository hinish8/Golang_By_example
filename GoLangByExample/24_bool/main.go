package main

import (
	"fmt"
	"strconv"
)

// %t can be used to print in the boolean form
// %v will print the default string. “true” for true and “false” for false

func main() {
	input := "true"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "false"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "garbage"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	} else {
		fmt.Println("Given input is not a bool")
	}

	t := true
	f := false

	fmt.Printf("%t %t\n", t, f)
	fmt.Printf("%v %v\n", t, f)

}
