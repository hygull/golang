/**
	{
		"created": "8 Feb 2019, Fri",
		"aim": "Working with encoding/decoding to/from JSON",
		"codedBy": "Rishikesh Agrawani"
		"reference": "https://golang.org/pkg/encoding/json/#MarshalIndent"
	}
*/

package main

import "fmt"
import "encoding/json"	

func main() {
	// Defining a struture
	type Person struct {
		Name string			`json:"name"`
		Age int8			`json:"age"`
		Profession string	`json:"profession"`
		Address string		`json:"address"`
		Languages []string	`json:"languages"`
	}

	// Initializing structure
	rishikesh := Person{
					Name: "Rishikesh",
					Age: 26, Profession: "Python/Django developer",
					Address: "Bangalore, India",
					Languages: []string{"Python", "JavaScript", "Golang"}, // , is must here
				}

	details, _ := json.MarshalIndent(&rishikesh, "", "\t") // pointer to struct, prefix, indent
	fmt.Println(string(details))
}