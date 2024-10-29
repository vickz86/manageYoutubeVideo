// main.go
package main

import (
	"fmt"
	"log"

	dataloader "github.com/vickz86/manageYoutubeVideo/dataLoader"
)

type youtubeStruct struct{
    url string
    channel string
    description string
    // 0 for no started , 1 for in progress , 2 for Done
    status int
}

// main is the entry point of the application
func main() {
    file,error := dataloader.LoadFileLines("data.txt")
    if error!=nil{
        log.Fatal("error loading file")
    }

    for _,v := range file{
        fmt.Println(v)

    }
}

// add takes two integers and returns their sum
func add(a int, b int) int {
    return a + b
}
