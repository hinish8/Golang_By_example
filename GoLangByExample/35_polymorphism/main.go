// Go doesn’t directly support method/function/operator overloading but
// variadic function provides a way of achieving the same with increased code complexity.

// Go implements run time polymorphism by using interfaces
package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}
type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}
func main() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}
	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax is %d\n", totalTax)
}
func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax() //This is where runtime polymorphism happens
	}
	return totalTax
}

// method overloading can be achieved by using, vardiac functions and interfaces
