/**
{
	"created": "14 Feb 2019, ",
	"aim": "Using break statement in Golang",
	"codedBy": "Rishikesh Agrawani"
}
*/

package main

import "fmt"

func printNumberPairs(useBreak bool) {
	i := 1
	// Outer for loop
	for i <= 5 {
		j := 1
		// Inner for loop
		for j <= 5 {
			if useBreak {
				if j == 4 {
					break
				}
			}
			fmt.Print("(", i, j, ")", " ")

			j += 1 // += operator increments value of j by 1
		}
		fmt.Println()
		i++ // Post increment of value of i by 1
	}
}

func main() {
	userBreaks := []bool{true, false}
	var msg string = ""

	for index, useBreak := range userBreaks {
		if index == 1 {
			msg = "without "
		}

		msg += "using break statement"
		fmt.Println(msg)

		printNumberPairs(useBreak)
		fmt.Println()
	}
}
