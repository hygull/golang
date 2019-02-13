/**
{
	"created": "13 Feb 2019, Wed",
	"aim": "Use of slice data structure in Golang",
	"codedBy": "Rishikesh Agrawani"
	"references": [
		"https://blog.golang.org/go-slices-usage-and-internals",
		"https://flaviocopes.com/go-empty-interface/"
	]
}
*/

package main

import "fmt"

func printCapacityAndLength(cities interface{}) {
	// Printing capacity & length of slice
	switch cities.(type) {
	case []int:
		fmt.Println("Case 1, []int")
		v, _ := cities.([]int)
		fmt.Println(cap(v), len(v))
	case []string:
		fmt.Println("Case 2, []string")
		v, _ := cities.([]string)
		fmt.Println(cap(v), len(v))
	default:
		fmt.Println("Catch does not match")
	}
}

func main() {
	// Creating slice of strings
	cities := []string{
		"Bangalore", "Raipur",
		"Kondagaon", "Bilaspur",
	}

	// Printing slice
	fmt.Println(cities)

	// Calculating total number of cities (strings) in slice
	fmt.Println("Total cities: ", len(cities))
	printCapacityAndLength(cities)

	// Appending 1|more cit(y|ies) to the slice,
	// printing length & capacity after each append operation
	cities = append(cities, "Gurgaon")
	printCapacityAndLength(cities)

	cities = append(cities, "Silicon Valley", "Hyderabad")
	fmt.Println("After addition of 3 more cities: ", cities)
	printCapacityAndLength(cities)

	// Creating another slice of villages
	villages := []string{"Bedagaon", "Badedongar"}

	// Appending all villages to cities (assume these villages turned to cities)
	cities = append(cities, villages...)
	fmt.Println("After addition of 2 more villages(in cities): ", cities)
	printCapacityAndLength(cities)

	// Slicing
	first3cities := cities[:3]
	fmt.Println("First 3 cities: ", first3cities)

	last3cities := cities[len(cities)-3 : len(cities)]
	fmt.Println("Last 3 cities: ", last3cities)

	fmt.Println("Working on slices of integers")
	// Using copy() function (evens & odds, the slices of integers)
	evens := make([]int, 3, 5)
	odds := []int{12, 45, 67, 89, 21}
	integers := make([]int, len(evens)+len(odds))
	printCapacityAndLength(evens)
	printCapacityAndLength(odds)
	printCapacityAndLength(integers)
	/*
		func copy(dst, src []T) int
		---
		ret == number of elements copied
	*/
	ret := copy(integers, odds)
	fmt.Println(integers)
	ret = copy(integers, evens)
	fmt.Println(integers, ret)
}
