// Golang provides encapsulation at the package level.
// Go doesn’t have any public,  private or protected keyword.
// The only mechanism to control the visibility is using the capitalized and non-capitalized formats

// There are five kinds of identifier which can be exported or non-exported/
// Structure
// Structure’s Method
// Structure’s Field
// Function
// Variable

package main

import "fmt"

//Test function
func Test() {
	//STRUCTURE IDENTIFIER
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//FUNCTION
	person2 := GetPerson()
	fmt.Println(person2)
	companyName := getCompanyName()
	fmt.Println(companyName)

	//VARIBLES
	fmt.Println(companyLocation)
	fmt.Println(CompanyName)
}
