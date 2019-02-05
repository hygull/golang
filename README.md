# Golang

<h3 id='1'>A simple hello world program</h3>

```go
/**
	{
		"created": "5 Feb 2019, Tue",
		"aim": "Writing a simple hello world program"
		"codedBy": "Rishikesh Agrawani"
	}
*/

// package declaration (required)
package main

// import statement (importing "fmt" package)
import "fmt"

// main() function definition
func main() {
	// Writing message to console using "Println()" function defined in "fmt" package
	fmt.Println("Hello programmers, this is really a great chance to learn Golang")
}
```

> **Output**

```bash
Rishikeshs-MacBook-Air:basic hygull$ pwd
/Users/hygull/Projects/Golang/golang/basic
Rishikeshs-MacBook-Air:basic hygull$ 
Rishikeshs-MacBook-Air:basic hygull$ ls
helloworld.go
Rishikeshs-MacBook-Air:basic hygull$ go run helloworld.go 
Hello programmers, this is really a great chance to learn Golang
Rishikeshs-MacBook-Air:basic hygull$ 
```

