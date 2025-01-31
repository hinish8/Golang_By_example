package main

// Func vs method
// it is clear that method has a receiver argument. A receiver can be a struct or any other type. The method will have access to the properties of the receiver and can call the receiver’s other methods.
// This is the only difference between function and method, but due to it they differ in terms of functionality they offer

// A function can be used as first-order objects and can be passed around while methods cannot.
// Methods can be used for chaining on the receiver while function cannot be used for the same.
// There can exist different methods with the same name with a different receiver, but there cannot exist two different functions with the same name in the same package.

// function as a type
//In go function is also a type. Two function will be of same type if

// They have the same number of arguments with each argument is of the same type
// They have same number of return values and each return value is of same type.

import "fmt"

func main() {

	// Immediately invoked functions
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)

	var shapes []shape
	s := &square{side: 2}
	shapes = append(shapes, s)
	r := &rectangle{length: 2, breath: 3}
	shapes = append(shapes, r)
	for _, shape := range shapes {
		fmt.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
	}
}

type shape interface {
	area() int
	getType() string
}
type rectangle struct {
	length int
	breath int
}

func (r *rectangle) area() int {
	return r.length * r.breath
}
func (r *rectangle) getType() string {
	return "rectangle"
}

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}
func (s *square) getType() string {
	return "square"
}
