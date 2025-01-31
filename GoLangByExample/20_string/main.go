package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Strings in back quotes are raw strings do not recognise any escape sequences
	multiline := `hello
world
this is a multiline string`
	fmt.Println(multiline)

	a := "  hello World  "
	num := "100"

	res, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Converted string is:", res)

	replace_all := strings.ReplaceAll(a, " ", "")
	fmt.Println("String after replacing all spaces with nothing is:", replace_all)

	// Comparison is done lexographically
	res_com := strings.Compare("bc", "abc")
	fmt.Println("The result of comparison is: ", res_com)

	present := strings.Contains("hello world", " w")
	fmt.Println("The result of contains is: ", present)

	// Stings separator function
	res1 := strings.Split("ab$cd$ef", "$")
	fmt.Println(res1)

	res2 := strings.Split("ab$cd$ef", "-")
	fmt.Println(res2)

	res3 := strings.Split("ab$cd$ef", "")
	fmt.Println(res3)

	res4 := strings.Split("", "")
	fmt.Println(res4)

	//Case 1 Input string contains the substring
	res5 := strings.Count("abcdabcd", "ab")
	fmt.Println(res5)

	//Case 1 Input string doesn't contains the substring
	res6 := strings.Count("abcdabcd", "xy")
	fmt.Println(res6)

	//Case 1 Substring supplied is an empty string
	res7 := strings.Count("abcdabcd", "")
	fmt.Println(res7)

	res8 := strings.Replace("abcdabxyabr", "ab", "12", 1)
	fmt.Println(res8)

	res9 := strings.Replace("abcdabxyabr", "ab", "12", -1)
	fmt.Println(res9)

	sample := "ab£d"
	r := []rune(sample)

	fmt.Printf("Before %s\n", string(r))
	r[2], r[3] = r[3], r[2]

	fmt.Printf("After %s\n", string(r))
}
