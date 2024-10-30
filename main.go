// main.go
package main

import (
	"fmt"

	dataloader "github.com/vickz86/manageYoutubeVideo/dataLoader"
)

type YoutubeStruct struct{
    url string
    channel string
    description string
    // 0 for no started , 1 for in progress , 2 for Done
    status int
}

// main is the entry point of the application
func main() {


    file := dataloader.CreateSliceYoutube("data.txt")
    fmt.Println(file)
    // file := dataloader.SplitEachLine("data.txt")
    

    // for _,v := range file{
    //     fmt.Println(v)
    //     fmt.Println(len(v))


    // }
}

