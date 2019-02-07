/**
	{
		"created": "7 Feb 2019, Thu",
		"aim": "Working with if else statements",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"

func main() {
	var age = 26
	var maxAge int = 100
	var firstName string = "Rishikesh"

	// If statement
	if( age < maxAge ) {
		fmt.Printf("I am still young as my age is %d and it is <= %d\n", age, maxAge)
	}

	// If else statement
	length := len(firstName)
	message := "My first name %s contains"
	message2 := "than %d characters"

	if( length > 10 ) {
		fmt.Printf(fmt.Sprintf("%s greater %s\n", message, message2), firstName, length)
	} else {
		fmt.Printf(fmt.Sprintf("%s less %s\n", message, message2), firstName, length)
	}

	// If else...if else statement
	var a, b, c = 12, 67, 87

	if( a > b ) { // Evaluates to false
		fmt.Printf("I am in 2nd floor")
	} else if( a > c ) { // Evaluates to false
		fmt.Printf("I am in 1st floor")
	} else {
		fmt.Printf("I am in ground floor") // So, it will be executed
	}

	fmt.Printf("\n")

}