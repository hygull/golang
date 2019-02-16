/**
	{
		"created": "16 Feb 2019, Fri",
		"aim": "Working with structure in Golang",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"
import "reflect"

// Defining a structure named Car to define the structure of car object

// Note: Make sure you have already capitalized struct filed names
// It's for iterating over it, and using Interface() method will cause error, if names are not capitalized
// eg. panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
// (If we used 'name' in place of 'Name')
type Car struct {
	Name string				
	Wheels int8				
	Price float64
	Color string
	Models []string
	PriceOptions []float64
	OfficialWebsite string
}

func main() {
	// - Create a structure object
	// - Even the data are not in order (below) as they appear (above)
	// but final output will follow the order defined its structure  
	bugatti := Car{
		Name: "Bugatti",
		Wheels: 4, // 4 will be of type int8 (here)
		Price: 40000000.58,
		PriceOptions: []float64{120, 130, 500},
		Color: "Red",
		Models: []string{"Type 57 SC Atlantic", "Type 41 Royale", "Vision Gran Turismo", "Chiron"},
		OfficialWebsite: "https://www.bugatti.com/home/",
	}

	// Printing structure
	fmt.Println(bugatti)

	// Accessing structure
	fmt.Println(bugatti.Name) // Bugatti

	// Checcking type
	fmt.Println(reflect.TypeOf(bugatti.Wheels)) // int8

	// Iterating over field of struct in Golang
	elem := reflect.ValueOf(&bugatti).Elem()
	fmt.Println("Elem: ", elem, "\n")

	for i := 0; i < elem.NumField(); i++ {
		name := elem.Type().Field(i).Name 
		typ := elem.Type().Field(i).Type 
		value := elem.Field(i).Interface()

		fmt.Println("Name:", name, "Type:", typ, "Value:", value)
	}
}