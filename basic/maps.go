/**
	{
		"created": "12 Feb 2019, Tue",
		"aim": "Working with maps in Golang",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"

func main() {
	// Creating a map (string as keys, integers as values)

	/*** 1st way ***/
	numbers := make(map[string]int) // Unordered data structure

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println(numbers)

	// Deleting 3rd entry from the map
	delete(numbers, "ten")
	fmt.Println(numbers)

	// Accessing entries
	fmt.Println(numbers["one"])
	val, exists := numbers["two"] // Here, return arguments `exists` is optional 
	fmt.Println(val, exists)

	/*** 2nd way ***/
	cities := map[string]string{"city1": "Bangalore", "city2": "Raipur", "city3": "Kondagaon"}
	fmt.Println(cities)
	fmt.Println(cities["city1"])
}