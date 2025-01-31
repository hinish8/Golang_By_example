// The reason we compare it with nil because the default zero value of interface is nil and since error is also an interface its default zero value is also nil./

package main

import (
	"errors"
	"fmt"
	"os"
)

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func main() {
	err := validate("", "")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}

	// using type that implements error interface
	// here the Open method return error
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}

	// Ways of creating error in Go
	sampleErr := errors.New("error occured")
	fmt.Println(sampleErr)

	sampleErr = fmt.Errorf("Err is: %s", "database connection issue")
	fmt.Println(sampleErr)

}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingField: "gender"}
	}
	return nil
}
