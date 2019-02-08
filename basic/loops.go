/**
	{
		"created": "7 Feb 2019, Thu",
		"aim": "Working on loops (3 basic forms)",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"

func main() {
	// for loop - 1st form
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println() // new line

	// for loop - 2nd form
	j := 1
	for j <= 3 {
		fmt.Println(j + 2)
		j++
	}

	fmt.Println() // new line

	// for loop - 3rd form
	arr := [5]int{1, 4, 5, 6, 67} // an array of integers of 5 elements 
	for index, item := range arr {
		fmt.Println(index, item) 
	}

	fmt.Println() // new line

	// Nested for loop
	for i := 1; i <= 4; i = i + 2 {  // i => 1, 3
		for j := 1; j <= 4; j += 2 { // j => 1, 3
			fmt.Println(i + j)       // 1 + 1, 1 + 3 | 3 + 1, 3 + 3 => 2, 4 | 4, 6
		}
	} 
}