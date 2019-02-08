/**
	{
		"created": "8 Feb 2019, Fri",
		"aim": "Working on decoding of JSON",
		"codedBy": "Rishikesh Agrawani",
		"reference": "https://golang.org/pkg/encoding/json/#MarshalIndent"
	}
*/

package main

import "fmt"
import "encoding/json"


func main() {
	// Creating a simple JSON string containing a person details
	jsonDetailsString := `{
							"fullname": "Rishikesh Agrawani",
							"age": 26,
							"profession": "Python/Django developer",
							"company": "MoneyBloom",
							"languages": [
								"Python", "JavaScript", "Golang"
							],
							"isActive": true
						}`

	// Defining a data structure to represent the decoded object
	type Person struct {
		Fullname string
		Age int8
		Profession string
		Company string 
		Languages []string
		IsActive bool
	}

	// Print values only
 	details := Person{}
	err := json.Unmarshal([]byte(jsonDetailsString), &details)

	if err != nil {
		panic(err)
	}
	fmt.Println(details, "\n")

	// Print with keys & values both 
	var detailsPretty map[string]interface{}
	err = json.Unmarshal([]byte(jsonDetailsString), &detailsPretty)

	if err != nil {
		panic(err)
	}
	fmt.Println(detailsPretty)
}