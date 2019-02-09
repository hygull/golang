/**
	{
		"created": "9 Feb 2019, Sat",
		"aim": "Calculating sum of all valid integers passed from command line",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"
import "os"
import "strconv"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You forgot to supply integers from command line")
	} else {
		fmt.Println("Go supplied integers as", args)
		var sum int

		for i := 0; i < len(args); i++ {
			data, err := strconv.Atoi(args[i])
			
			if(err == nil) { // Curley braces are required unike JavaScript otherwise
							 // syntax error: assignment sum += data used as value
				sum += data
			} else {
				panic(err)
			}
		} 
		fmt.Println("Sum =", sum)
	}
}