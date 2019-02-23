/**
	{
		"created": "23 Feb 2019, Sat",
		"aim": "Working with functions in Golang",
		"codedBy": "Rishikesh Agrawani",
        "reference": "https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line"
	}
*/

package main

import "fmt"
import "bufio"
import "os"

/* A function which asks user to enter data */
func printMessage(message string) {
    fmt.Print("Enter your " + message + ": ") // Display message to console output (screen)
}

/* A function which takes data (fullname) from user and returns to caller */
func getFullname() string {
    var fullname string

    // Using fmt.Scan(&fullname) only takes first word from a multi-word string
    reader := bufio.NewReader(os.Stdin)
    printMessage("fullname") // Call printMessage() function
    fullname, _ = reader.ReadString('\n') // Take fullname from console input (keyboard)

    return fullname // Return the value entered
}

/* A function which takes data (age) from user and returns to caller */
func getAge() int8 {
    var age int8

    printMessage("age") // Call printMessage() function
    fmt.Scanln(&age)     // Take age from console input (keyboard)

    return age          // Return the value entered
}

/* A function which prints data entered by user */
func printDetails(fullname string, age int8) {
    fmt.Println("\nDetails:-")
    fmt.Print("Full name: ", fullname)
    fmt.Println("Age.     : ", age)
}

/* 
    - main() function, the only required function in Golang program
    - A starter (which initiates the work)
*/
func main() {
    var fullname string = getFullname() // Call getFullname() function
    age := getAge() // Call getAge() function

    printDetails(fullname, age)  // Call printDetails() function
}  