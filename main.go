// main.go
package main

import (
	"fmt"

	dataloader "github.com/vickz86/manageYoutubeVideo/dataLoader"
)

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

