package model

type YoutubeStruct struct {
	Url         string
	Channel     string
	Description string
	// 0 for no started , 1 for in progress , 2 for Done
	Status int
}
