/**
{
	"created": "15 Feb 2019, Fri",
	"aim": "Working with arrays in Go",
	"codedBy": "Rishikesh Agrawani",
	"reference": "https://tour.golang.org/moretypes/6"
}
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Creating an array of 5 integers (Note: Array is of fixed size unlike slices)
	var numbers [5]int

	numbers[0] = 0
	numbers[1] = -67
	numbers[2] = 43

	// Printing the above array
	fmt.Println(numbers) // Last 2 array positions (3 & 4)
	// will have the zero value of integer i.e. 0 (again)

	// Creating an array of 3 strings by initializing them in the same statement
	scientists := [3]string{"Albert Einstein", "Dr. Homi Jenhagir Bhabha", "Sir Issac Newton"}

	// Printing the above array of strings
	fmt.Println(scientists)

	// Accessing contents of arrays
	fmt.Println(numbers[2])    // 43
	fmt.Println(scientists[1]) // Dr. Homi Jenhagir Bhabha

	// Calculating length of arrays
	fmt.Println("len(scientists) = ", len(scientists))
	fmt.Println("len(numbers) = ", len(numbers))

	// Checking type of arrays
	fmt.Println(reflect.TypeOf(scientists)) // [3]string
	fmt.Println(reflect.TypeOf(numbers))    // [5]int

	// Slicing arrays
	slice1 := scientists[1:len(scientists)] // 1:3 => 1, 2 (indices)
	slice2 := numbers[2:5]                  // 2:5 => 2, 3, 4 (indices)

	// Printing created slices
	fmt.Println(slice1)
	fmt.Println(slice2)

	// Checking type of slices
	fmt.Println(reflect.TypeOf(slice1)) // []string
	fmt.Println(reflect.TypeOf(slice2)) // []int
}
