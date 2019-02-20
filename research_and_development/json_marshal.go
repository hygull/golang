/**
	{
		"created": "20 Feb 2019, Wed",
		"aim": "A quick look at the working nature of Marshal() with structure",
		"codedBy": "Rishikesh Agrawani"
	}
*/

package main

import "fmt"
import "os"
import "encoding/json"

type Car struct {
	Name string
	Wheels int
}

type Truck struct {
	name string
	wheels int
}

type Bus struct {
	Name string `json:"name"`
	Wheels int	`json:"wheels"`
}

func main() {
	fmt.Println("Example #1 - Exported structure field:-")
	car := Car{Name: "Maruti", Wheels: 56}
	b, err := json.Marshal(car)
	fmt.Println("car  :", b)
	fmt.Println("error:", err)
	fmt.Println("1st form:", string(b))
	fmt.Print("2nd form: ")
	os.Stdout.Write(b) // It does not place "\n" at end like Println()

	fmt.Println("\n") // New line
	fmt.Println("Example #2 - Unexported structure field:-")

	truck := Truck{name: "TATA truck", wheels: 12}
	b, err = json.Marshal(truck)
	fmt.Println(string(b), err)

	fmt.Println() // New line
	fmt.Println("Example #3 - Exported structure field (Problem fix):-")

	// Let's fix the issue and get the desired output (with all keys in lowercase)
	bus := Bus{Name: "Mahendra bus", Wheels: 8}
	b, err = json.Marshal(bus)
	fmt.Println(string(b), err)
}