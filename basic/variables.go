/**
	{
		"created": "5 Feb 2019, Tue",
		"aim": "creating variables in Golang (using 4 commnly used styles)",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"

func main() {
	// 1st
	// defining type of variable a as int
	var a int 
	// initialization of variable a (assignment)
	a = 1729

	// 2nd
	// creating 2/more variables in a single statement (using above style/approach)
	var c, d float64 = 34.56, 3.14

	// 3rd
	// creating variable b in a single statement
	b := 67

	// 4th
	// creating 2/more variables in a single statement (using 3rd approach)
	fullname, age := "Rishikesh Agrawani", 26

	// priting values of all variables
	fmt.Println(a, b, c, d) // 1729 67 34.56 3.14
	fmt.Printf("%s, %d\n", fullname, age) // Rishikesh Agrawani, 26
}