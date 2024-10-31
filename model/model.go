package model

import "fmt"

type YoutubeStruct struct {
	Index       int
	Url         string
	Channel     string
	Description string
	// 0 for no started , 1 for in progress , 2 for Done
	Status int
}


// print index, description and status of a YoutubeStruct
func (y YoutubeStruct) Describe() {
	fmt.Printf("Index:%v\nDescription:\n%s\nurl %s: ",y.Index,y.Description,y.Url)
	//TODO , print status based on int value
	switch y.Status{
	case 0 :
		fmt.Println("Status : not started")
	case 1 :
		fmt.Println("Status : in progress")
	case 2 :
		fmt.Println("Status : done")

	}
	fmt.Println("")

}