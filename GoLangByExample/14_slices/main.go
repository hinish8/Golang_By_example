package main

import "fmt"

func main() {
	s := make([]string, 2, 3)
	fmt.Println(s)

	p := []string{"a", "b", "c"}
	fmt.Println(p)

	p = append(p, "d")
	fmt.Println(p)

	//Iterate over a slcie
	for _, val := range p {
		fmt.Println(val)
	}
	fmt.Println("Capacity: ", cap(p))
	fmt.Println("Length: ", len(p))

	// MAP

	var employeeSalary map[string]int
	fmt.Println(employeeSalary)

	employeeSalary2 := make(map[string]int)
	fmt.Println(employeeSalary2)

	employeeSalary3 := map[string]int{
		"John": 1000,
		"Sam":  1200,
	}
	fmt.Println(employeeSalary3)

	employeeSalary3["Carl"] = 1500
	fmt.Printf("John salary is %d\n", employeeSalary3["John"])
	delete(employeeSalary3, "Carl")
	fmt.Println("\nPrinting employeeSalary3 map")
	fmt.Println(employeeSalary3)
}
