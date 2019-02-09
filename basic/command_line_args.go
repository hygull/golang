/**
	{
		"created": "9 Feb 2019, Sat",
		"aim": "Working with command line arguments",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"
import "os"

func main() {
	// Getting raw command line arguments
	rawArguments := os.Args
	fmt.Println(rawArguments)

	// Getting list of supplied arguments from command line
	suppliedArguments := rawArguments[1:]
	fmt.Println(suppliedArguments)
}