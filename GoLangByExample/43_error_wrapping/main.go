package main

import (
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}

func main() {

	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(e2)
	fmt.Println(e3)

}
