// main.go
package main

import (
	"fmt"
	"os"

	browser "github.com/toqueteos/webbrowser"
	stringnumber "github.com/vickz86/GoFunctions/stringNumber"
	dataloader "github.com/vickz86/manageYoutubeVideo/dataLoader"
	"github.com/vickz86/manageYoutubeVideo/model"
)

// main is the entry point of the application
func main() {

	//load the slice of youtubeStruct
	youStruc := dataloader.CreateSliceYoutube("data.txt")

	// LOGIC
    question := "toDo : 0-exit,1-print all links,2-open links"

    for {
        //Get int result from what to do
        toDo := stringnumber.IntFromString(question,0,2)

        switch toDo{
        case 0 :
            os.Exit(0)
        case 1:
            //print a description of all link
            for _,stru := range youStruc{
                stru.Describe()
            }
        case 2:
            OpenLink(youStruc)


        }

    }
    
	// }
}

// - - - - FUNCTION - - - -
// open a link based on index
func OpenLink(allLinks []model.YoutubeStruct){
    //get the max value
    maxVal := len(allLinks)-1

    //question to ask
    question := "type index of link to open"

    // get the index to open
    indexNb := stringnumber.IntFromString(question,0,maxVal)

    for nb,st := range allLinks{
        if nb == indexNb{
            err := browser.Open(st.Url)
            if err != nil{
                fmt.Println(err)
            }
        }

    }





}
