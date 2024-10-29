// main.go
package main

import "fmt"

type youtubeStruct struct{
    url string
    channel string
    description string
    // 0 for no started , 1 for in progress , 2 for Done
    status int
}

// main is the entry point of the application
func main() {
    result := add(3, 5)
    fmt.Printf("The result of adding is: %d\n", result)
}

// add takes two integers and returns their sum
func add(a int, b int) int {
    return a + b
}
