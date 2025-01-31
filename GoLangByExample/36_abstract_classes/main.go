// // We will use a combination of a interface (abstract interface) and struct (abstract concrete type).
// // Together they can provide the functionalities of an abstract class.

// https://golangbyexample.com/go-abstract-class/

// package main

// import "fmt"

// //Abstract Interface
// type iAlpha interface {
// 	work()
// 	common()
// }

// //Abstract Concrete Type
// type alpha struct {
// 	name string
// }

// func (a *alpha) common() {
// 	fmt.Println("common called")
// }

// //Implementing Type
// type beta struct {
// 	alpha
// }

// func (b *beta) work() {
// 	fmt.Println("work called")
// 	fmt.Printf("name is %s\n", b.name)
// 	b.common()
// }

// func main() {
// 	a := alpha{
// 		name: "test",
// 	}
// 	b := &beta{
// 		alpha: a,
// 	}
// 	b.work()
// }

// // package main

// // import "fmt"

// // //Abstract Interface
// // type iAlpha interface {
// //     work()
// //     common()
// // }

// // //Abstract Concrete Type
// // type alpha struct {
// //     name string
// //     work func()
// // }

// // func (a *alpha) common() {
// //     fmt.Println("common called")
// //     a.work()
// // }

// // //Implementing Type
// // type beta struct {
// //     alpha
// // }

// // func (b *beta) work() {
// //     fmt.Println("work called")
// //     fmt.Printf("name is %s\n", b.name)
// // }

// // func main() {
// //     a := alpha{
// //         name: "test",
// //     }
// //     b := &beta{
// //         alpha: a,
// //     }
// //     b.alpha.work = b.work
// //     b.common()
// // }

package main

import "fmt"

//Abstract Interface
type iAlpha interface {
	work()
	common(iAlpha)
}

//Abstract Concrete Type
type alpha struct {
	name string
}

func (a *alpha) common(i iAlpha) {
	fmt.Println("common called")
	i.work()
}

//Implementing Type
type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("Name is %s\n", b.name)
}

func main() {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.common(b)
}
