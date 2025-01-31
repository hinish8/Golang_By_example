// The above error tells basically that sub-typing is not possible in GO by just using embedding.
// https://golangbyexample.com/oop-inheritance-golang-complete/
package main

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi from say function")
}

type child struct {
	base  //embedding
	style string
}

func main() {
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
}

// package main
// import "fmt"
// type iBase interface {
//     say()
// }
// type base struct {
//     color string
// }
// func (b *base) say() {
//     b.clear()
// }
// func (b *base) clear() {
//     fmt.Println("Clear from base's function")
// }
// type child struct {
//     base  //embedding
//     style string
// }
// func (b *child) clear() {
//     fmt.Println("Clear from child's function")
// }
// func check(b iBase) {
//     b.say()
// }
// func main() {
//     base := base{color: "Red"}
//     child := &child{
//         base:  base,
//         style: "somestyle",
//     }
//     child.say()
// }

//hierarchy

// package main
// import "fmt"
// type iAnimal interface {
//     breathe()
// }
// type animal struct {
// }
// func (a *animal) breathe() {
//     fmt.Println("Animal breate")
// }
// type iAquatic interface {
//     iAnimal
//     swim()
// }
// type aquatic struct {
//     animal
// }
// func (a *aquatic) swim() {
//     fmt.Println("Aquatic swim")
// }
// type iNonAquatic interface {
//     iAnimal
//     walk()
// }
// type nonAquatic struct {
//     animal
// }
// func (a *nonAquatic) walk() {
//     fmt.Println("Non-Aquatic walk")
// }
// type shark struct {
//     aquatic
// }
// type lion struct {
//     nonAquatic
// }
// func main() {
//     shark := &shark{}
//     checkAquatic(shark)
//     checkAnimal(shark)
//     lion := &lion{}
//     checkNonAquatic(lion)
//     checkAnimal(lion)
// }
// func checkAquatic(a iAquatic) {}
// func checkNonAquatic(a iNonAquatic) {}
// func checkAnimal(a iAnimal) {}
