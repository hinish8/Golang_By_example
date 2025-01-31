package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Printf("%U\n", []rune("0b£"))

	r := 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))
	fmt.Printf("Unicode CodePoint: %U\n", r)
	fmt.Printf("Character: %c\n", r)
	s := "0b£"
	fmt.Printf("%U\n", []rune(s))
	fmt.Println([]rune(s))

	runeArray := []rune{'a', 'b', '£'}
	str := string(runeArray)
	fmt.Println(str)

}
