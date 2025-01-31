package main

import "fmt"

type Size uint8

const (
	a, d = iota + 10, iota
	b, e
	c, f
)

const (
	m = iota
	n
	_ // increment skipped
	o
)

const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(m, n, o)
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}
