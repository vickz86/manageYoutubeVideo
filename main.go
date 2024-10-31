// main.go
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	browser "github.com/toqueteos/webbrowser"
	stringnumber "github.com/vickz86/GoFunctions/stringNumber"
	timeUti "github.com/vickz86/GoFunctions/timeUti"
	dataloader "github.com/vickz86/manageYoutubeVideo/dataLoader"
	"github.com/vickz86/manageYoutubeVideo/model"
)

// main is the entry point of the application
func main() {

	//load the slice of youtubeStruct
	youStruc := dataloader.CreateSliceYoutube("data.txt")

	// LOGIC
    question := "toDo : 0-exit,1-print all links,2-open link,3-change time of link"

    for {
        //Get int result from what to do
        toDo := stringnumber.IntFromString(question,0,3)

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
        case 3:
            youStruc =SetTimeStruct(youStruc)

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

// split the string into 2 string at & symbol
func SplitStrinInTwo(theString string)(string,string,error){
    
    string1,string2,isHere := strings.Cut(theString,"&")
    if !isHere{
        return "","",errors.New("ERROR , string dont contain &")
    }


    return string1,string2,nil


}

func SetTimeStruct(allLinks []model.YoutubeStruct)[]model.YoutubeStruct{
        //get the max value
        maxVal := len(allLinks)-1

        //question to ask
        question := "type index of link to change time"
    
        // get the index to change time
        indexNb := stringnumber.IntFromString(question,0,maxVal)


        //create the first part of the new string for the url
        firstPart,_,err := SplitStrinInTwo(allLinks[indexNb].Url)
        if err!=nil{
            fmt.Println(err)
        }
        
        
        //create the second part of the new url
        var timeString string
        
        fmt.Println("type the new time for this video : hh.mm.ss")

        _,err2:= fmt.Scan(&timeString)
        if err2!=nil{
            fmt.Println(err)
        }
        
        // convert to string : &t=60s
        newTime,_,err3 := timeUti.StringToSecond(timeString)
        if err3!=nil{
            fmt.Println(err)
        }

        newTimeString := firstPart+"&t="+newTime

        //set the new URL
        allLinks[indexNb].Url = newTimeString

        //if set to not started set to inpprogress
        if allLinks[indexNb].Status ==0{
            allLinks[indexNb].Status =1
        }
        
        return allLinks


}
