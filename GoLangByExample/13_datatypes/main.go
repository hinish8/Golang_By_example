package main

// byte  --- char -- ASCII -- 8 bits
// rune ---- unicode --- 32 bits

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	b := 2
	fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))

	var c int8 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(c))

	sizeOfuintInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfuintInBits)

	var d uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(d))

}
