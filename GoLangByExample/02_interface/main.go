package main

import "fmt"

type car interface {
	engineStart() bool
	engineStop() bool
}

type conveyance interface {
	transport()
}

type sedan struct {
	name string
}

func (s *sedan) engineStart() bool {
	fmt.Println("Engine has started")
	return true
}

func (s *sedan) engineStop() bool {
	fmt.Println("Engine has stopped")
	return true
}

func (s *sedan) transport() {
	fmt.Println("This Sedan can also be used for transportation")
}

func main() {
	var se sedan
	se.name = "sedan_x"

	var c car = &se

	if c.engineStart() && c.engineStop() {
		var con conveyance = &se
		con.transport()
	}

}
