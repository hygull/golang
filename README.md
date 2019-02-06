# Golang

+ is a general purpose, statically typed, strongly typed programming language.

+ was initially developed at Google in year 2007 by Robert Griesemer, Rob Pike & Ken Thompson.

+ provides inbuilt support for Garbage collection and supports concurrent programming.


<h2 id='examples'> Examples</h2>

Have a look at the following examples one by one. I think it would be very easy for you to understand as it is very simple and most of the lines are commented.

<h3 id='1'>Writing a simple hello world program</h3>

```go
/**
    {
        "created": "5 Feb 2019, Tue",
        "aim": "Writing a simple hello world program",
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

```

<h3 id='2'>Creating variables in Golang (using 4 commnly used styles)</h3>

```go
/**
    {
        "created": "5 Feb 2019, Tue",
        "aim": "creating variables in Golang (using 4 commnly used styles)",
        "codedBy": "Rishikesh Agrawani"
    }
*/

package main

import "fmt"

func main() {
    // 1st
    // defining type of variable a as int
    var a int 
    // initialization of variable a (assignment)
    a = 1729

    // 2nd
    // creating 2/more variables in a single statement (using above style/approach)
    var c, d float64 = 34.56, 3.14

    // 3rd
    // creating variable b in a single statement
    b := 67

    // 4th
    // creating 2/more variables in a single statement (using 3rd approach)
    fullname, age := "Rishikesh Agrawani", 26

    // printing values of all variables
    fmt.Println(a, b, c, d) // 1729 67 34.56 3.14
    fmt.Printf("%s, %d\n", fullname, age) // Rishikesh Agrawani, 26
}

```

> Output

```bash
Rishikeshs-MacBook-Air:basic hygull$ go run variables.go 
1729 67 34.56 3.14
Rishikesh Agrawani, 26

```


<h3 id='references'>References</h3>

The following are the list of urls of websites where you can go and try to learn & clear concepts of Golang.

These are really excellent to teach and make you understand Go as I have experienced and still I'm experiencing.

- [x] A tour of Golang - [https://tour.golang.org/welcome/1](https://tour.golang.org/welcome/1)

- [x] Dotnetperls - [https://www.dotnetperls.com/#go](https://www.dotnetperls.com/#go)

- [x] Tutorialspoint - [https://www.tutorialspoint.com/go/index.htm](https://www.tutorialspoint.com/go/index.htm)

- [x] GoByExample - [https://gobyexample.com](https://gobyexample.com)

- [x] How to write Go code - [https://golang.org/doc/code.html](https://golang.org/doc/code.html)

- [x] Code organization - [https://golang.org/doc/code.html#Organization](https://golang.org/doc/code.html#Organization)

- [x] GOPATH setting - [https://github.com/golang/go/wiki/SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
