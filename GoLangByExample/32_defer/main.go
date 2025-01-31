// deferred function will also be executed if the enclosing function terminates abruptly. For example in case of a panic
// defer function can read as well as modified those named return values. If the defer function modifies the name return value then that modified value will  be returned
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := writeToTempFile("Some text")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Write to file succesful")

	defer func() { fmt.Println("In inline defer") }()
	fmt.Println("Executed")

	// Important output
	i := 0
	i = 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)

	sample := "abc"
	defer fmt.Printf("In defer sample is: %s\n", sample)
	sample = "xyz"
}

func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteString(text)
	if err != nil {
		return err
	}
	fmt.Printf("Number of bytes written: %d", n)
	return nil
}
