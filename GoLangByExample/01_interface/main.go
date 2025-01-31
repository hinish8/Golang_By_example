package main

import (
	"fmt"
)

type animal interface {
	walk() string
	breathe() string
	age() int
}

type cat struct {
	name string
}

func (c cat) walk() string {
	return c.name + "cat is walking"
}

func (c cat) breathe() string {
	return c.name + "is breathing"
}

func (c cat) age() int {
	return -1
}

// passing an empty inteface to a func
func hello(v interface{}) {
	fmt.Println("Inteface as an argument to a function")
	fmt.Printf("%T , %v \n", v, v)
}

func main() {

	var a animal
	c := cat{name: "tim"}
	a = c
	fmt.Println(a.walk())
	fmt.Println(a.breathe())
	fmt.Println("Age of the cat is: ", a.age())
	hello("hi")
	hello(true)
}
