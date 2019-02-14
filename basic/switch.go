/**
	{
		"created": "14 Feb 2019, Thu",
		"aim": "Using select statement in Go",
		"codedBy": "Rishikesh Agrawani"
        "reference": "https://golangbot.com/types/"
	}
*/

package main

import "fmt"

func checkType(data interface{}) {
    switch typ := data.(type) {
        case int:
            fmt.Println("INT:", data)
        case float32:
            fmt.Println("FLOAT32:", data)
        case float64:
            fmt.Println("FLOAT64:", data)
        case complex64:
            fmt.Println("COMPLEX64:", data)
        case complex128:
            fmt.Println("COMPLEX128:", data)
        case string:
            fmt.Println("STRING:", data)
        case byte:
            fmt.Println("BYTE:", data)
        case bool:
            fmt.Println("BOOL:", data)
        case map[string]int:
            fmt.Println("MAP(string, int):", data)
        case []string:
            fmt.Println("SLICE(string):", data)
        case rune: // An alias for int32
            fmt.Println("RUNE/INT32: ", data)
        default:
            fmt.Println("Got it, but I don't know about it,", data, "is of type ", typ)
    }
}

func main() {
    checkType(12)
    checkType("Python & Golang are my favourite")
    checkType(3.14)
    checkType(23 + 8i)
    checkType(3 == 4)
    checkType([]string{"Go", "Python", "Node", "JavaScript"})
    checkType(map[string]int{"nice": 1, "one": 2, "and": 3, "then": 4})
    checkType([]int{1729, 67, 92, 2019})

    var c complex64 = 6 + 7i
    var f float32 = 92.67
    var r rune = 1992

    checkType(c)
    checkType(f)
    checkType(r)
}